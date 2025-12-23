package run

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	pageValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#D8DEE9"))

	pageDimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6C6C6C")).
			Italic(true)
)

type LogsPageContent struct {
	Logs []string
}

func (l *LogsPageContent) Init() tea.Cmd { return nil }

func (l *LogsPageContent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return l, nil
}

func (l *LogsPageContent) View() string {
	if len(l.Logs) == 0 {
		return pageDimStyle.Render("No error logs available")
	}

	var b strings.Builder

	for i, logEntry := range l.Logs {
		b.WriteString(fmt.Sprintf("• %s\n", pageValueStyle.Render(logEntry)))
		if i < len(l.Logs)-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}
