package run

import (
	"fmt"
	"strings"

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
	Logs []LogLine
}

func (l *LogsPageContent) View() string {
	if len(l.Logs) == 0 {
		return pageDimStyle.Render("No error logs available")
	}

	var b strings.Builder

	for i, logEntry := range l.Logs {
		b.WriteString(fmt.Sprintf("• %s\n", pageValueStyle.Render(string(logEntry))))
		if i < len(l.Logs)-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}
