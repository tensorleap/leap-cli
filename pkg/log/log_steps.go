package log

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jedib0t/go-pretty/v6/text"
	"golang.org/x/term"
)

// -----------------------------
// Public types
// -----------------------------

type Step struct {
	// ID is the server-side event id (e.g. "build_dependencies"). Stable across
	// wording changes to Name, so behaviour keyed off a specific step matches on
	// this rather than on the display label.
	ID      string
	Name    string
	Status  StepStatus
	Current float64
	Total   float64
}

type Job struct {
	ID    string
	Steps []Step
}

// -----------------------------
// Shared interface
// -----------------------------

type rendererImpl interface {
	Start()
	Update(steps []Step)
	UpdateLogs(lines []string)
	Stop()
}

// -----------------------------
// Adaptive public wrapper
// -----------------------------

type Renderer struct {
	impl  rendererImpl
	IsTTY bool
}

// NewRenderer automatically picks TTY or log mode.
func NewRenderer() *Renderer {
	isTTY := term.IsTerminal(int(os.Stdout.Fd())) || os.Getenv("FORCE_TTY") == "1"
	if isTTY {
		return &Renderer{impl: newTTYRenderer(), IsTTY: true}
	}
	return &Renderer{impl: newLogRenderer(), IsTTY: false}
}

func (r *Renderer) Start()                    { r.impl.Start() }
func (r *Renderer) Update(steps []Step)       { r.impl.Update(steps) }
func (r *Renderer) UpdateLogs(lines []string) { r.impl.UpdateLogs(lines) }
func (r *Renderer) Stop()                     { r.impl.Stop() }

// ====================================================================
// TTY RENDERER (interactive spinner + in-place updates)
// ====================================================================

type ttyRenderer struct {
	mu       sync.Mutex
	writer   *blockWriter
	steps    []Step
	stepMap  map[string]*Step
	logLines []string
	stopCh   chan struct{}
	done     bool
}

func newTTYRenderer() *ttyRenderer {
	return &ttyRenderer{
		writer:  newBlockWriter(),
		stepMap: make(map[string]*Step),
		stopCh:  make(chan struct{}),
	}
}

const FrameDuration = time.Millisecond * 100 // 10 frame per second
var spin = spinner.CharSets[14]

func (r *ttyRenderer) Start() {
	frame := 0

	go func() {
		t := time.NewTicker(FrameDuration)
		defer t.Stop()
		for {
			select {
			case <-r.stopCh:
				return
			case <-t.C:
				r.mu.Lock()
				r.redraw(spin[frame])
				r.mu.Unlock()
				frame = (frame + 1) % len(spin)
			}
		}
	}()
}

func (r *ttyRenderer) Stop() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.done {
		return
	}
	r.done = true
	close(r.stopCh)
}

func (r *ttyRenderer) Update(steps []Step) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var newSteps []Step
	newMap := make(map[string]*Step)
	for _, s := range steps {
		if existing, ok := r.stepMap[s.Name]; ok {
			existing.Status = s.Status
			existing.Current = s.Current
			existing.Total = s.Total
			newSteps = append(newSteps, *existing)
			copy := *existing
			newMap[s.Name] = &copy
		} else {
			newStep := s
			newSteps = append(newSteps, newStep)
			newMap[s.Name] = &newStep
		}
	}
	r.steps = newSteps
	r.stepMap = newMap
}

// UpdateLogs sets the rolling log tail shown beneath the steps. Pass nil to
// remove it.
func (r *ttyRenderer) UpdateLogs(lines []string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.logLines = lines
}

func (r *ttyRenderer) redraw(spin string) {
	width := getTerminalWidth()
	lines := r.renderLines(spin, width)
	r.writer.Render(lines, width)
}

// logGutter keeps the tail visually subordinate to the steps above it.
var logGutter = text.FgHiBlack.Sprint("│")

func (r *ttyRenderer) renderLines(spin string, width int) []string {
	lines := []string{}

	for _, s := range r.steps {
		icon := diffIcon(s.Status)
		if s.Status == StepStatusRunning && spin != "" {
			icon = text.FgHiBlue.Sprint(spin)
		}
		lines = append(lines, printStepStatus(s, icon))
	}

	if len(r.logLines) == 0 {
		return lines
	}

	// The whole block is cursor-addressed, so it has to fit on screen —
	// clearPreviousLines can't reach rows that have scrolled off the top.
	height := LogTailHeight
	if max := getTerminalHeight() / 3; height > max {
		height = max
	}
	// "  │ " prefix; the tail wraps to the remaining columns.
	const gutterWidth = 4
	tailWidth := width - gutterWidth
	if height < 1 || tailWidth < 1 {
		return lines
	}

	lines = append(lines, "")
	for _, row := range fixedTail(r.logLines, height, tailWidth) {
		lines = append(lines, "  "+logGutter+" "+row)
	}
	return lines
}

// --------------------------------------------------------------------
// blockWriter (cursor-anchored block clearing for TTY)
// --------------------------------------------------------------------

type blockWriter struct {
	mu        sync.Mutex
	lineCount int
	lastWidth int
}

func newBlockWriter() *blockWriter { return &blockWriter{} }

// Render repaints the block in place. Callers must pass only lines that fit
// within width — lineCount counts strings while clearPreviousLines erases rows,
// so a wrapped line would leave debris behind on every frame.
func (bw *blockWriter) Render(lines []string, width int) {
	bw.mu.Lock()
	defer bw.mu.Unlock()

	// On resize the terminal re-flows rows we already printed, so lineCount no
	// longer describes what's on screen and erasing would eat the wrong rows.
	// Leaving one stale block behind beats accumulating debris every frame.
	resized := bw.lastWidth != 0 && bw.lastWidth != width

	if bw.lineCount > 0 && !resized {
		clearPreviousLines(bw.lineCount)
	}

	for _, ln := range lines {
		fmt.Println(ln)
	}
	bw.lineCount = len(lines)
	bw.lastWidth = width
}

func clearPreviousLines(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("\033[1A") // move cursor up one line
		fmt.Print("\033[2K") // clear the entire line
	}
}

// ====================================================================
// LOG RENDERER (non-TTY incremental diff printer)
// ====================================================================

type logRenderer struct {
	mu          sync.Mutex
	lastPrinted map[string]Step
	steps       []Step
	printedLogs map[string]bool
}

func newLogRenderer() *logRenderer {
	return &logRenderer{
		lastPrinted: make(map[string]Step),
		printedLogs: make(map[string]bool),
	}
}

func (r *logRenderer) Start() {
	// no spinner in log mode
}

func (r *logRenderer) Stop() {
	// no stopping in log mode
}

func (r *logRenderer) Update(steps []Step) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.steps = steps
	for _, s := range steps {
		prev, ok := r.lastPrinted[s.Name]
		if !ok || stepChanged(prev, s) {
			fmt.Println(printStepStatus(s, diffIcon(s.Status)))
			r.lastPrinted[s.Name] = s
		}
	}
}

// UpdateLogs appends log lines not yet seen. A rolling window is meaningless
// when the output is a file or CI log, so each line is printed once instead.
func (r *logRenderer) UpdateLogs(lines []string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, line := range lines {
		clean := sanitizeLine(line)
		if clean == "" || r.printedLogs[clean] {
			continue
		}
		r.printedLogs[clean] = true
		fmt.Printf("    | %s\n", clean)
	}
}

// shared helpers
func stepChanged(prev, curr Step) bool {
	return prev.Status != curr.Status ||
		prev.Current != curr.Current ||
		prev.Total != curr.Total
}

func printStepStatus(s Step, icon string) string {
	if s.Total > 0 {
		return fmt.Sprintf("  %s %s (%d/%d)", icon, s.Name, int(s.Current), int(s.Total))
	} else {
		return fmt.Sprintf("  %s %s", icon, s.Name)
	}
}

type StepStatus string

const (
	StepStatusDone    StepStatus = "DONE"
	StepStatusFailed  StepStatus = "FAILED"
	StepStatusRunning StepStatus = "RUNNING"
	StepStatusPending StepStatus = "PENDING"
	StepStatusWaiting StepStatus = "WAITING"
	// Step's work wasn't needed for this run (engine marks events SKIPPED
	// when the caller didn't request a given step or its inputs are
	// unchanged). Terminal status, treat alongside DONE.
	StepStatusSkipped StepStatus = "SKIPPED"
)

func diffIcon(status StepStatus) string {
	switch status {
	case StepStatusDone:
		return text.FgGreen.Sprint("✔")
	case StepStatusFailed:
		return text.FgRed.Sprint("✖")
	case StepStatusRunning:
		return text.FgHiBlue.Sprint("▶")
	case StepStatusSkipped:
		return text.FgCyan.Sprint("—")
	default:
		return text.FgHiBlack.Sprint("•")
	}
}
