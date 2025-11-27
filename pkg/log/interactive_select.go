package log

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ---------------------------
// Callback interface (same)
// ---------------------------
type OnInputChange func(cursor int, filter string) (title string, options []string, newCursor int, inputErr string, hint string)
type OnEnter func(options []string, cursor int, filter string) (selectedOption *string, newFilter string, inputErr string, hint string)
type OnTabKey func(options []string, cursor int, filter string) (newFilter string)

// ---------------------------
// Model
// ---------------------------
type interactiveModel struct {
	list          list.Model
	input         textinput.Model
	filter        string
	title         string
	options       []string
	cursor        int
	onInputChange OnInputChange
	onEnter       OnEnter
	onTabKey      OnTabKey
	selected      *string
	width, height int
	quitting      bool
	err           error
	inputErr      string
	hint          string
}

// Item
type listItem string

func (i listItem) Title() string       { return string(i) }
func (i listItem) Description() string { return "" }
func (i listItem) FilterValue() string { return string(i) }

// ---------------------------
// Styles
// ---------------------------
var (
	colorBlue     = lipgloss.Color("#00BFFF") // survey-like blue
	colorGray     = lipgloss.Color("#A0A0A0")
	styleTitle    = lipgloss.NewStyle().Bold(true).Foreground(colorBlue).Transform(func(s string) string { return fmt.Sprintf("? %s", s) })
	styleError    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true)
	subtitleStyle = lipgloss.NewStyle().Foreground(colorGray).Bold(true)
)

// ---------------------------
// Init
// ---------------------------
func newInteractiveModel(onInputChange OnInputChange, onEnter OnEnter, onTabKey OnTabKey) interactiveModel {
	title, options, cursor, inputErr, hint := onInputChange(0, "")
	items := make([]list.Item, len(options))
	for i, opt := range options {
		items[i] = listItem(opt)
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false // no double-spacing
	delegate.SetSpacing(0)           // remove spacing between items
	delegate.Styles.SelectedTitle.Foreground(colorBlue)
	delegate.Styles.SelectedTitle = lipgloss.NewStyle().
		Foreground(colorBlue). // survey blue
		Bold(true).
		Transform(func(title string) string { return fmt.Sprintf("> %s", title) })

	l := list.New(items, delegate, 0, 0)
	l.SetShowHelp(false)
	l.SetFilteringEnabled(false)
	l.Title = title
	l.Styles.Title = styleTitle
	l.SetShowPagination(false)
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)
	l.SetShowTitle(false)

	ti := textinput.New()
	ti.Placeholder = ""
	ti.Prompt = lipgloss.NewStyle().
		Foreground(colorBlue).
		Bold(true).
		Render("❯ ") // survey-style prompt
	ti.Focus()
	ti.CharLimit = 256

	return interactiveModel{
		list:          l,
		title:         title,
		options:       options,
		input:         ti,
		cursor:        cursor,
		onInputChange: onInputChange,
		onEnter:       onEnter,
		onTabKey:      onTabKey,
		inputErr:      inputErr,
		hint:          hint,
	}
}

// ---------------------------
// Bubble Tea
// ---------------------------
func (m interactiveModel) Init() tea.Cmd { return nil }

func (m interactiveModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	updateFilter := func(newFilter string) {
		if newFilter != m.filter {
			m.filter = newFilter
			m.input.SetValue(m.filter)
			m.input.CursorEnd()
			m.refreshOptions()
		}
	}

	updateErrorAndHint := func(inputErr string, hint string) {
		m.inputErr = inputErr
		m.hint = hint
	}

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.list.SetSize(m.width, m.height-6)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			// ENTER only applies when list is active
			selIdx := -1
			if len(m.options) > 0 {
				selIdx = m.list.Index()
			}
			opt, newFilter, err, hint := m.onEnter(m.options, selIdx, m.filter)
			if opt != nil {
				m.selected = opt
				m.quitting = true
				return m, tea.Quit
			}
			updateFilter(newFilter)
			updateErrorAndHint(err, hint)
		case "tab":
			selIdx := -1
			if len(m.options) > 0 {
				selIdx = m.list.Index()
			}
			newFilter := m.onTabKey(m.options, selIdx, m.filter)
			updateFilter(newFilter)

		case "up", "down":
			m.list, cmd = m.list.Update(msg)
			return m, cmd

		case "left", "right":
			m.input, cmd = m.input.Update(msg)
			return m, cmd

		case "backspace":
			if m.input.Focused() && len(m.filter) > 0 {
				m.filter = m.filter[:len(m.filter)-1]
				m.input.SetValue(m.filter)
				m.refreshOptions()
			}

		default:
			// typing always focuses input
			if len(msg.Runes) > 0 && msg.Runes[0] >= 32 && msg.Runes[0] <= 126 {
				if !m.input.Focused() {
					m.input.Focus()
				}
				m.input, _ = m.input.Update(msg)
				m.filter = m.input.Value()
				m.refreshOptions()
				return m, nil
			}
		}
	}

	// update input always (so it re-renders properly)
	// var inputCmd tea.Cmd
	// m.input, inputCmd = m.input.Update(msg)

	return m, cmd
}

func (m *interactiveModel) refreshOptions() {
	title, options, cursor, inputErr, hint := m.onInputChange(m.cursor, m.filter)
	m.title = title
	m.options = options
	m.cursor = cursor
	m.inputErr = inputErr
	m.hint = hint

	items := make([]list.Item, len(options))
	for i, opt := range options {
		items[i] = listItem(opt)
	}
	m.list.SetItems(items)
	m.list.Title = styleTitle.Render(m.title)
	if cursor >= 0 && cursor < len(items) {
		m.list.Select(cursor)
	}
}

func (m interactiveModel) View() string {
	if m.quitting {
		if m.err != nil {
			return fmt.Sprintf("Error: %v\n", m.err)
		}
		if m.selected != nil {
			return ""
		}
		return ""
	}

	// footer with filter and navigation hint
	statusLine := "↑↓ move | ENTER select"
	errorOrHint := ""
	if m.inputErr != "" {
		errorOrHint = styleError.Render(m.inputErr)
	} else {
		if m.hint != "" {
			errorOrHint = fmt.Sprintf("%s | %s", statusLine, m.hint)
		}
		errorOrHint = subtitleStyle.Render(errorOrHint)
	}

	title := styleTitle.Render(m.title + ":")

	header := fmt.Sprintf("%s %s\n%s",
		title,
		errorOrHint,
		m.input.View(),
	)

	return fmt.Sprintf("%s\n%s", header, m.list.View())
}

// ---------------------------
// Public API
// ---------------------------
func InteractiveSelectTea(onInputChange OnInputChange, onEnter OnEnter, onTabKey OnTabKey) (string, error) {
	p := tea.NewProgram(newInteractiveModel(onInputChange, onEnter, onTabKey))

	m, err := p.Run()
	if err != nil {
		return "", err
	}

	fm := m.(interactiveModel)
	if fm.selected != nil {
		return *fm.selected, nil
	}
	return "", fmt.Errorf("cancelled")
}
