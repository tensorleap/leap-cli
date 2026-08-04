package log

import (
	"os"
	"regexp"
	"strings"

	"github.com/charmbracelet/x/ansi"
	"golang.org/x/term"
)

// LogTailHeight is how many terminal rows the tail occupies. Rows, not log
// lines — a long line wraps and spends more than one of them.
const LogTailHeight = 5

const defaultTerminalWidth = 80

// tabWidth matches the near-universal terminal default. Tabs are expanded here
// because a wrapper counting cells can't know the terminal's real tab stops.
const tabWidth = 8

// csiSequence matches CSI escape sequences (colour, but also cursor movement
// and erase). Log content is echoed into the user's terminal, so sequences that
// move the cursor would break blockWriter's row arithmetic — strip all of them
// rather than trying to allow-list the harmless ones.
var csiSequence = regexp.MustCompile(`\x1b\[[0-9;?]*[ -/]*[@-~]`)

// otherEscape catches the non-CSI escape forms (OSC, single-char escapes) left
// over after csiSequence.
var otherEscape = regexp.MustCompile(`\x1b[]PX^_][^\x1b]*(?:\x1b\\|\x07)?|\x1b.`)

func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width <= 0 {
		return defaultTerminalWidth
	}
	return width
}

// sanitizeLine makes a raw log line safe to print inside a cursor-addressed
// block: no escape sequences the terminal would act on, no carriage returns
// jumping to column 0, and tabs expanded so width measurement is honest.
func sanitizeLine(line string) string {
	line = csiSequence.ReplaceAllString(line, "")
	line = otherEscape.ReplaceAllString(line, "")

	var b strings.Builder
	col := 0
	for _, r := range line {
		switch {
		case r == '\t':
			pad := tabWidth - (col % tabWidth)
			b.WriteString(strings.Repeat(" ", pad))
			col += pad
		case r == '\r' || r == '\n':
			// Dropped, not replaced: \r would jump the cursor to column 0 and
			// overwrite the row; \n would desync lineCount from rows printed.
		case r < 0x20 || r == 0x7f:
			// Remaining C0 controls and DEL — no width, unpredictable effect.
		default:
			b.WriteRune(r)
			col++
		}
	}
	return b.String()
}

// toRows sanitizes and hard-wraps a log line into one string per terminal row.
// Every returned string is guaranteed to fit the width, which is what keeps
// blockWriter's "one string == one row" assumption true.
func toRows(line string, width int) []string {
	if width <= 0 {
		width = defaultTerminalWidth
	}
	clean := sanitizeLine(line)
	if clean == "" {
		return []string{""}
	}
	return strings.Split(ansi.Hardwrap(clean, width, true), "\n")
}

// fixedTail renders log lines as exactly n terminal rows, showing the most
// recent output and padding with blanks so the block height never changes
// between frames (a block that changes height visibly jitters).
func fixedTail(lines []string, n, width int) []string {
	if n <= 0 {
		return nil
	}

	rows := make([]string, 0, n)
	for _, line := range lines {
		rows = append(rows, toRows(line, width)...)
	}

	out := make([]string, n)
	start := len(rows) - n
	if start < 0 {
		start = 0
	}
	copy(out, rows[start:])
	return out
}
