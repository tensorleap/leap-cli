package log

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
)

// The invariant blockWriter depends on: every string it prints occupies exactly
// one terminal row. If a row is wider than the terminal it wraps, and
// clearPreviousLines then erases too few rows and strands output every frame.
func assertFitsWidth(t *testing.T, rows []string, width int) {
	t.Helper()
	for i, row := range rows {
		if w := ansi.StringWidth(row); w > width {
			t.Errorf("row %d is %d cells wide, exceeds width %d: %q", i, w, width, row)
		}
	}
}

func TestToRowsWrapsToWidth(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		width    int
		wantRows int
	}{
		{"short line stays one row", "Collecting torch", 40, 1},
		{"exact width stays one row", strings.Repeat("a", 40), 40, 1},
		{"one over wraps to two", strings.Repeat("a", 41), 40, 2},
		{"three times width", strings.Repeat("a", 120), 40, 3},
		{"empty line is one blank row", "", 40, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := toRows(tt.line, tt.width)
			if len(rows) != tt.wantRows {
				t.Errorf("got %d rows, want %d: %q", len(rows), tt.wantRows, rows)
			}
			assertFitsWidth(t, rows, tt.width)
		})
	}
}

func TestToRowsNeutralizesControlSequences(t *testing.T) {
	tests := []struct {
		name string
		line string
		want string
	}{
		{
			// Would move the cursor two rows up and corrupt the whole block.
			name: "cursor movement stripped",
			line: "before\x1b[2Aafter",
			want: "beforeafter",
		},
		{
			name: "colour stripped",
			line: "\x1b[33mwarning\x1b[0m",
			want: "warning",
		},
		{
			// Would jump to column 0 and overwrite the row in place.
			name: "carriage return dropped",
			line: "progress 50%\rprogress 100%",
			want: "progress 50%progress 100%",
		},
		{
			name: "tab expanded to next stop",
			line: "ab\tc",
			want: "ab      c",
		},
		{
			name: "indentation preserved",
			line: "  Downloading torch",
			want: "  Downloading torch",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := toRows(tt.line, 200)
			got := strings.Join(rows, "")
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestToRowsWideCharactersFitWidth(t *testing.T) {
	// Full-width runes are one rune but two cells; a rune-counting wrapper
	// would emit rows twice as wide as the terminal.
	rows := toRows(strings.Repeat("生", 40), 20)
	assertFitsWidth(t, rows, 20)
}

func TestFixedTailAlwaysReturnsNRows(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		n     int
	}{
		{"fewer lines than rows", []string{"a", "b"}, 5},
		{"exactly n lines", []string{"a", "b", "c", "d", "e"}, 5},
		{"more lines than rows", []string{"a", "b", "c", "d", "e", "f", "g"}, 5},
		{"no lines", nil, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fixedTail(tt.lines, tt.n, 40)
			if len(got) != tt.n {
				t.Errorf("got %d rows, want %d", len(got), tt.n)
			}
			assertFitsWidth(t, got, 40)
		})
	}
}

func TestFixedTailKeepsMostRecent(t *testing.T) {
	got := fixedTail([]string{"a", "b", "c", "d", "e", "f", "g"}, 3, 40)
	want := []string{"e", "f", "g"}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("row %d: got %q, want %q", i, got[i], want[i])
		}
	}
}

func TestFixedTailLongLineConsumesMultipleRows(t *testing.T) {
	// A line three widths long should occupy three of the five slots, leaving
	// room for only the two most recent short lines before it.
	long := strings.Repeat("x", 120)
	got := fixedTail([]string{"a", "b", "c", long}, 5, 40)

	if len(got) != 5 {
		t.Fatalf("got %d rows, want 5", len(got))
	}
	if got[0] != "b" || got[1] != "c" {
		t.Errorf("expected the two newest short lines before the long one, got %q, %q", got[0], got[1])
	}
	for i := 2; i < 5; i++ {
		if got[i] != strings.Repeat("x", 40) {
			t.Errorf("row %d: expected a full-width chunk of the long line, got %q", i, got[i])
		}
	}
}
