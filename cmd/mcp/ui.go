package mcp

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// the colors read on light and dark terminals alike, so the renderer never queries the terminal's
// background (that query prints stray escape codes in some terminals)
var renderer = func() *lipgloss.Renderer {
	r := lipgloss.NewRenderer(os.Stdout)
	r.SetHasDarkBackground(true)
	return r
}()

var (
	okStyle    = renderer.NewStyle().Foreground(lipgloss.Color("#3FA34D")).Bold(true)
	warnStyle  = renderer.NewStyle().Foreground(lipgloss.Color("#D08C00")).Bold(true)
	failStyle  = renderer.NewStyle().Foreground(lipgloss.Color("#E5484D")).Bold(true)
	titleStyle = renderer.NewStyle().Bold(true)
	dimStyle   = renderer.NewStyle().Foreground(lipgloss.Color("#7B8794"))
	codeStyle  = renderer.NewStyle().Foreground(lipgloss.Color("#2F80C9"))
)

// ui prints the setup as a short checklist: one line per step, details indented under it
type ui struct{ w io.Writer }

func (u ui) title(s string)               { fmt.Fprintf(u.w, "\n%s\n\n", titleStyle.Render(s)) }
func (u ui) ok(format string, a ...any)   { u.line(okStyle.Render("✓"), format, a...) }
func (u ui) warn(format string, a ...any) { u.line(warnStyle.Render("!"), format, a...) }
func (u ui) fail(format string, a ...any) { u.line(failStyle.Render("✗"), format, a...) }
func (u ui) skip(format string, a ...any) { u.line(dimStyle.Render("–"), format, a...) }
func (u ui) detail(format string, a ...any) {
	for _, l := range strings.Split(fmt.Sprintf(format, a...), "\n") {
		fmt.Fprintf(u.w, "    %s\n", dimStyle.Render(l))
	}
}
func (u ui) command(s string) { fmt.Fprintf(u.w, "    %s\n", codeStyle.Render(s)) }
func (u ui) text(format string, a ...any) {
	fmt.Fprintf(u.w, format+"\n", a...)
}

func (u ui) line(mark, format string, a ...any) {
	fmt.Fprintf(u.w, "  %s %s\n", mark, fmt.Sprintf(format, a...))
}
