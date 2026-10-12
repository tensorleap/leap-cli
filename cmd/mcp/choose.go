package mcp

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var errCancelled = errors.New("cancelled")

type choice struct {
	label, detail, badge string
}

// asked counts questions so far; every question after the first gets a blank line above it
var asked int

type chooser struct {
	lead     string
	question string
	choices  []choice
	cursor   int
	picked   bool
	quit     bool
}

func (m *chooser) Init() tea.Cmd { return nil }

func (m *chooser) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if k, ok := msg.(tea.KeyPressMsg); ok {
		switch k.String() {
		case "up", "k", "shift+tab":
			m.cursor = (m.cursor + len(m.choices) - 1) % len(m.choices)
		case "down", "j", "tab":
			m.cursor = (m.cursor + 1) % len(m.choices)
		case "enter", " ":
			m.picked = true
			return m, tea.Quit
		case "esc", "ctrl+c", "q":
			m.quit = true
			return m, tea.Quit
		default:
			if len(k.String()) == 1 && k.String()[0] >= '1' && int(k.String()[0]-'1') < len(m.choices) {
				m.cursor = int(k.String()[0] - '1')
				m.picked = true
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m *chooser) View() tea.View {
	if m.picked || m.quit {
		return tea.NewView("")
	}
	width := 0
	for _, c := range m.choices {
		width = max(width, lipgloss.Width(c.label))
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%s%s %s\n\n", m.lead, codeStyle.Render("?"), titleStyle.Render(m.question))
	for i, c := range m.choices {
		pointer, label := "  ", c.label
		if i == m.cursor {
			pointer = codeStyle.Render("❯ ")
			label = titleStyle.Render(label)
		}
		pad := strings.Repeat(" ", width-lipgloss.Width(c.label)+3)
		line := fmt.Sprintf("  %s%s%s%s", pointer, label, pad, dimStyle.Render(c.detail))
		if c.badge != "" {
			line += "  " + okStyle.Render(c.badge)
		}
		b.WriteString(line + "\n")
	}
	b.WriteString("\n  " + dimStyle.Render("↑/↓ move · enter select · esc cancel") + "\n")
	return tea.NewView(b.String())
}

// choose asks one question in the terminal and leaves a one-line answer behind
func choose(out ui, question, answerLabel string, choices []choice, def int) (int, error) {
	m := &chooser{question: question, choices: choices, cursor: def}
	if asked > 0 {
		m.lead = "\n"
	}
	asked++
	if _, err := tea.NewProgram(m, tea.WithOutput(os.Stdout)).Run(); err != nil {
		return 0, err
	}
	if !m.picked {
		return 0, errCancelled
	}
	if answerLabel == "" {
		out.ok("%s", choices[m.cursor].label)
	} else {
		out.ok("%s: %s", answerLabel, titleStyle.Render(choices[m.cursor].label))
	}
	return m.cursor, nil
}

func interactive() bool {
	return term.IsTerminal(int(os.Stdin.Fd())) && term.IsTerminal(int(os.Stdout.Fd()))
}

// spin shows a spinner beside msg while work runs; the caller prints the outcome line
func spin(w io.Writer, msg string, work func()) {
	if !interactive() {
		work()
		return
	}
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		for i := 0; ; i++ {
			fmt.Fprintf(w, "\r  %s %s", codeStyle.Render(frames[i%len(frames)]), dimStyle.Render(msg))
			select {
			case <-done:
				fmt.Fprint(w, "\r\033[2K")
				return
			case <-time.After(80 * time.Millisecond):
			}
		}
	}()
	work()
	close(done)
	wg.Wait()
}
