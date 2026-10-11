package interactive_pages

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
)

type staticContent string

func (s staticContent) View() string { return string(s) }

func TestTabbedReportKeys(t *testing.T) {
	r := &TabbedReport{Pages: []Page{
		{Name: "First", Content: staticContent(strings.Repeat("line\n", 100))},
		{Name: "Second", Content: staticContent("short")},
	}}
	r.Init()
	r.Update(tea.WindowSizeMsg{Width: 80, Height: 20})

	r.Update(tea.KeyPressMsg{Code: 'j', Text: "j"})
	if got := r.getActiveScrollable().scrollOffset; got != 1 {
		t.Fatalf("scroll offset after j = %d, want 1", got)
	}
	r.Update(tea.KeyPressMsg{Code: 'G', Text: "G"})
	if s := r.getActiveScrollable(); s.scrollOffset != s.maxScrollOffset() {
		t.Fatalf("G did not scroll to bottom")
	}

	r.Update(tea.KeyPressMsg{Code: tea.KeyRight})
	if r.activeTabIndex != 1 {
		t.Fatalf("active tab after right = %d, want 1", r.activeTabIndex)
	}

	if _, cmd := r.Update(tea.KeyPressMsg{Code: 'q', Text: "q"}); cmd == nil {
		t.Fatal("q did not quit")
	} else if _, ok := cmd().(tea.QuitMsg); !ok {
		t.Fatal("q did not return tea.Quit")
	}
	if !r.View().AltScreen {
		t.Fatal("report should render in the alt screen")
	}
}
