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
	Stop()
}

// -----------------------------
// Adaptive public wrapper
// -----------------------------

type Renderer struct {
	impl rendererImpl
}

// NewRenderer automatically picks TTY or log mode.
func NewRenderer() *Renderer {
	isTTY := term.IsTerminal(int(os.Stdout.Fd())) || os.Getenv("FORCE_TTY") == "1"
	if isTTY {
		return &Renderer{impl: newTTYRenderer()}
	}
	return &Renderer{impl: newLogRenderer()}
}

func (r *Renderer) Start()              { r.impl.Start() }
func (r *Renderer) Update(steps []Step) { r.impl.Update(steps) }
func (r *Renderer) Stop()               { r.impl.Stop() }

// ====================================================================
// TTY RENDERER (interactive spinner + in-place updates)
// ====================================================================

type ttyRenderer struct {
	mu      sync.Mutex
	writer  *blockWriter
	steps   []Step
	stepMap map[string]*Step
	stopCh  chan struct{}
	done    bool
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

func (r *ttyRenderer) redraw(spin string) {
	lines := r.renderLines(spin)
	r.writer.Render(lines)
}

func (r *ttyRenderer) renderLines(spin string) []string {
	lines := []string{}

	for _, s := range r.steps {
		icon := diffIcon(s.Status)
		if s.Status == StepStatusRunning && spin != "" {
			icon = text.FgHiBlue.Sprint(spin)
		}
		lines = append(lines, printStepStatus(s, icon))
	}
	return lines
}

// --------------------------------------------------------------------
// blockWriter (cursor-anchored block clearing for TTY)
// --------------------------------------------------------------------

type blockWriter struct {
	mu        sync.Mutex
	lineCount int
}

func newBlockWriter() *blockWriter { return &blockWriter{} }

func (bw *blockWriter) Render(lines []string) {
	bw.mu.Lock()
	defer bw.mu.Unlock()

	if bw.lineCount > 0 {
		clearPreviousLines(bw.lineCount)
	}

	for _, ln := range lines {
		fmt.Println(ln)
	}
	bw.lineCount = len(lines)
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
}

func newLogRenderer() *logRenderer {
	return &logRenderer{
		lastPrinted: make(map[string]Step),
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
)

func diffIcon(status StepStatus) string {
	switch status {
	case StepStatusDone:
		return text.FgGreen.Sprint("✔")
	case StepStatusFailed:
		return text.FgRed.Sprint("✖")
	case StepStatusRunning:
		return text.FgHiBlue.Sprint("▶")
	default:
		return text.FgHiBlack.Sprint("•")
	}
}
