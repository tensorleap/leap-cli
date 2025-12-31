package interactive_pages

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	helpTextHeight   = 2
	tabBarHeight     = 2
	separatorHeight  = 1
	contentPaddingX  = 6
	minContentHeight = 3
	minContentWidth  = 20
	defaultWidth     = 60
)

var (
	colorBlue      = lipgloss.Color("#4A90D9")
	colorLightBlue = lipgloss.Color("#7EB6FF")
	colorCyan      = lipgloss.Color("#5BC0DE")
	colorGray      = lipgloss.Color("#6C757D")
	colorLightGray = lipgloss.Color("#ADB5BD")
	colorDarkGray  = lipgloss.Color("#495057")
	colorWhite     = lipgloss.Color("#F8F9FA")

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorLightBlue)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(colorCyan)

	helpStyle = lipgloss.NewStyle().
			Foreground(colorGray).
			MarginTop(1)

	activeTabStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorWhite).
			Background(colorBlue).
			Padding(0, 2)

	inactiveTabStyle = lipgloss.NewStyle().
				Foreground(colorLightGray).
				Background(colorDarkGray).
				Padding(0, 2)

	tabBarStyle = lipgloss.NewStyle().
			MarginTop(1).
			MarginBottom(0)

	contentStyle = lipgloss.NewStyle().
			Padding(0, 2).
			MarginTop(0)

	separatorStyle = lipgloss.NewStyle().
			Foreground(colorDarkGray)

	pageHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorBlue)

	tabArrowStyle = lipgloss.NewStyle().
			Foreground(colorGray)
)

type WrappedLine struct {
	Original     string
	Segments     []string
	SegmentStart int
}

type ScrollableContent struct {
	innerContent  tea.Model
	wrappedLines  []string
	lineMapping   []WrappedLine
	scrollOffset  int
	viewWidth     int
	viewHeight    int
	isInitialized bool
}

func NewScrollableContent(content tea.Model) *ScrollableContent {
	return &ScrollableContent{
		innerContent: content,
		scrollOffset: 0,
	}
}

func (s *ScrollableContent) Init() tea.Cmd {
	if s.innerContent != nil {
		return s.innerContent.Init()
	}
	return nil
}

func (s *ScrollableContent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return s.handleKeyPress(msg)
	case tea.MouseMsg:
		return s.handleMouseEvent(msg)
	}

	if s.innerContent != nil {
		var cmd tea.Cmd
		s.innerContent, cmd = s.innerContent.Update(msg)
		s.rebuildWrappedContent()
		return s, cmd
	}

	return s, nil
}

func (s *ScrollableContent) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		s.scrollUp(1)
	case "down", "j":
		s.scrollDown(1)
	case "pgup":
		s.scrollUp(s.viewHeight / 2)
	case "pgdown":
		s.scrollDown(s.viewHeight / 2)
	case "home", "g":
		s.scrollToTop()
	case "end", "G":
		s.scrollToBottom()
	}
	return s, nil
}

func (s *ScrollableContent) handleMouseEvent(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	switch msg.Button {
	case tea.MouseButtonWheelUp:
		s.scrollUp(3)
	case tea.MouseButtonWheelDown:
		s.scrollDown(3)
	}
	return s, nil
}

func (s *ScrollableContent) scrollUp(lines int) {
	s.scrollOffset -= lines
	if s.scrollOffset < 0 {
		s.scrollOffset = 0
	}
}

func (s *ScrollableContent) scrollDown(lines int) {
	maxOffset := s.maxScrollOffset()
	s.scrollOffset += lines
	if s.scrollOffset > maxOffset {
		s.scrollOffset = maxOffset
	}
}

func (s *ScrollableContent) scrollToTop() {
	s.scrollOffset = 0
}

func (s *ScrollableContent) scrollToBottom() {
	s.scrollOffset = s.maxScrollOffset()
}

func (s *ScrollableContent) maxScrollOffset() int {
	max := len(s.wrappedLines) - s.viewHeight
	if max < 0 {
		return 0
	}
	return max
}

func (s *ScrollableContent) rebuildWrappedContent() {
	if s.innerContent == nil {
		return
	}

	rawContent := s.innerContent.View()
	rawLines := strings.Split(rawContent, "\n")
	collapsedLines := collapseConsecutiveEmptyLines(rawLines, 2)

	s.wrappedLines = nil
	s.lineMapping = nil
	segmentIndex := 0

	for _, line := range collapsedLines {
		segments := wrapLinePreservingAnsi(line, s.viewWidth)
		mapping := WrappedLine{
			Original:     line,
			Segments:     segments,
			SegmentStart: segmentIndex,
		}
		s.lineMapping = append(s.lineMapping, mapping)
		s.wrappedLines = append(s.wrappedLines, segments...)
		segmentIndex += len(segments)
	}
}

func wrapLinePreservingAnsi(line string, maxWidth int) []string {
	if maxWidth <= 0 {
		return []string{line}
	}

	visibleWidth := lipgloss.Width(line)
	if visibleWidth <= maxWidth {
		return []string{line}
	}

	var segments []string
	var currentSegment strings.Builder
	var escapeSequence strings.Builder
	var activeAnsiStyle string
	currentWidth := 0
	insideEscape := false

	for i := 0; i < len(line); i++ {
		char := line[i]

		if char == '\x1b' {
			insideEscape = true
			escapeSequence.Reset()
			escapeSequence.WriteByte(char)
			currentSegment.WriteByte(char)
			continue
		}

		if insideEscape {
			escapeSequence.WriteByte(char)
			currentSegment.WriteByte(char)
			if isAnsiTerminator(char) {
				insideEscape = false
				activeAnsiStyle = updateActiveStyle(activeAnsiStyle, escapeSequence.String())
			}
			continue
		}

		currentWidth++
		currentSegment.WriteByte(char)

		if currentWidth >= maxWidth {
			if activeAnsiStyle != "" {
				currentSegment.WriteString("\x1b[0m")
			}
			segments = append(segments, currentSegment.String())
			currentSegment.Reset()
			currentWidth = 0
			if activeAnsiStyle != "" {
				currentSegment.WriteString(activeAnsiStyle)
			}
		}
	}

	if currentSegment.Len() > 0 {
		segments = append(segments, currentSegment.String())
	}

	if len(segments) == 0 {
		return []string{line}
	}

	return segments
}

func isAnsiTerminator(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func updateActiveStyle(current string, escapeSequence string) string {
	if escapeSequence == "\x1b[0m" || escapeSequence == "\x1b[m" {
		return ""
	}
	return escapeSequence
}

func collapseConsecutiveEmptyLines(lines []string, threshold int) []string {
	if threshold < 1 {
		threshold = 2
	}

	var result []string
	consecutiveEmptyCount := 0

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			consecutiveEmptyCount++
			continue
		}

		if consecutiveEmptyCount > threshold {
			result = append(result, "")
			result = append(result, fmt.Sprintf("--- %d empty lines ---", consecutiveEmptyCount))
			result = append(result, "")
		} else {
			for i := 0; i < consecutiveEmptyCount; i++ {
				result = append(result, "")
			}
		}
		consecutiveEmptyCount = 0
		result = append(result, line)
	}

	if consecutiveEmptyCount > threshold {
		result = append(result, "")
		result = append(result, fmt.Sprintf("--- %d empty lines ---", consecutiveEmptyCount))
		result = append(result, "")
	} else {
		for i := 0; i < consecutiveEmptyCount; i++ {
			result = append(result, "")
		}
	}

	return result
}

func (s *ScrollableContent) View() string {
	if len(s.wrappedLines) == 0 {
		return ""
	}

	if s.viewHeight <= 0 {
		return strings.Join(s.wrappedLines, "\n")
	}

	visibleStart := s.clampIndex(s.scrollOffset)
	visibleEnd := s.clampIndex(visibleStart + s.viewHeight)

	return strings.Join(s.wrappedLines[visibleStart:visibleEnd], "\n")
}

func (s *ScrollableContent) clampIndex(index int) int {
	if index < 0 {
		return 0
	}
	if index > len(s.wrappedLines) {
		return len(s.wrappedLines)
	}
	return index
}

func (s *ScrollableContent) OriginalContent() string {
	if len(s.lineMapping) == 0 {
		return strings.Join(s.wrappedLines, "\n")
	}

	var originals []string
	for _, mapping := range s.lineMapping {
		originals = append(originals, mapping.Original)
	}
	return strings.Join(originals, "\n")
}

func (s *ScrollableContent) ScrollPercent() float64 {
	maxOffset := s.maxScrollOffset()
	if maxOffset <= 0 {
		return 0
	}
	return float64(s.scrollOffset) / float64(maxOffset)
}

func (s *ScrollableContent) SetDimensions(width, height int) {
	s.viewWidth = width
	s.viewHeight = height
	s.isInitialized = true
	s.rebuildWrappedContent()
	s.clampScrollOffset()
}

func (s *ScrollableContent) clampScrollOffset() {
	maxOffset := s.maxScrollOffset()
	if s.scrollOffset > maxOffset {
		s.scrollOffset = maxOffset
	}
}

type Page struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content tea.Model
}

type TabbedReport struct {
	Title           string `json:"title"`
	Subtitle        string `json:"subtitle"`
	Helper          string `json:"helper"`
	Pages           []Page `json:"pages"`
	activeTabIndex  int
	windowWidth     int
	windowHeight    int
	isQuitting      bool
	isReady         bool
	hasCopied       bool
	scrollablePages []*ScrollableContent
}

func (r *TabbedReport) Init() tea.Cmd {
	r.initializeScrollablePages()
	var cmds []tea.Cmd
	for _, page := range r.scrollablePages {
		if page != nil {
			cmds = append(cmds, page.Init())
		}
	}
	return tea.Batch(cmds...)
}

func (r *TabbedReport) initializeScrollablePages() {
	r.scrollablePages = make([]*ScrollableContent, len(r.Pages))
	for i, page := range r.Pages {
		r.scrollablePages[i] = NewScrollableContent(page.Content)
	}
}

func (r *TabbedReport) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return r.handleWindowResize(msg)
	case tea.MouseMsg:
		return r.forwardToActiveScrollable(msg)
	case tea.KeyMsg:
		return r.handleKeyPress(msg)
	}

	return r.forwardToActiveScrollable(msg)
}

func (r *TabbedReport) handleWindowResize(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	r.windowWidth = msg.Width
	r.windowHeight = msg.Height
	r.isReady = true

	contentHeight := r.windowHeight - r.fixedHeaderHeight() - separatorHeight
	contentWidth := r.windowWidth - contentPaddingX

	if contentHeight < minContentHeight {
		contentHeight = minContentHeight
	}
	if contentWidth < minContentWidth {
		contentWidth = defaultWidth
	}

	var cmds []tea.Cmd
	for i, scrollable := range r.scrollablePages {
		if scrollable != nil {
			scrollable.SetDimensions(contentWidth, contentHeight)
			if scrollable.innerContent != nil {
				var cmd tea.Cmd
				scrollable.innerContent, cmd = scrollable.innerContent.Update(msg)
				cmds = append(cmds, cmd)
			}
			r.scrollablePages[i] = scrollable
		}
	}

	return r, tea.Batch(cmds...)
}

func (r *TabbedReport) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q", "esc":
		r.isQuitting = true
		return r, tea.Quit
	case "left":
		r.navigateToPreviousTab()
		return r, nil
	case "right":
		r.navigateToNextTab()
		return r, nil
	case "tab":
		r.navigateToNextTabCyclic()
		return r, nil
	case "shift+tab":
		r.navigateToPreviousTabCyclic()
		return r, nil
	case "c":
		r.copyActivePageContent()
		return r, nil
	}

	return r.forwardToActiveScrollable(msg)
}

func (r *TabbedReport) navigateToPreviousTab() {
	if r.activeTabIndex > 0 {
		r.activeTabIndex--
		r.hasCopied = false
	}
}

func (r *TabbedReport) navigateToNextTab() {
	if r.activeTabIndex < len(r.Pages)-1 {
		r.activeTabIndex++
		r.hasCopied = false
	}
}

func (r *TabbedReport) navigateToNextTabCyclic() {
	r.activeTabIndex = (r.activeTabIndex + 1) % len(r.Pages)
	r.hasCopied = false
}

func (r *TabbedReport) navigateToPreviousTabCyclic() {
	r.activeTabIndex = (r.activeTabIndex - 1 + len(r.Pages)) % len(r.Pages)
	r.hasCopied = false
}

func (r *TabbedReport) copyActivePageContent() {
	activeScrollable := r.getActiveScrollable()
	if activeScrollable == nil {
		return
	}

	content := activeScrollable.OriginalContent()
	cleanContent := removeAnsiCodes(content)
	if err := clipboard.WriteAll(cleanContent); err == nil {
		r.hasCopied = true
	}
}

func (r *TabbedReport) forwardToActiveScrollable(msg tea.Msg) (tea.Model, tea.Cmd) {
	if len(r.scrollablePages) == 0 || r.activeTabIndex >= len(r.scrollablePages) {
		return r, nil
	}

	activeScrollable := r.scrollablePages[r.activeTabIndex]
	if activeScrollable == nil {
		return r, nil
	}

	updatedModel, cmd := activeScrollable.Update(msg)
	r.scrollablePages[r.activeTabIndex] = updatedModel.(*ScrollableContent)
	return r, cmd
}

func (r *TabbedReport) getActiveScrollable() *ScrollableContent {
	if len(r.scrollablePages) > 0 && r.activeTabIndex < len(r.scrollablePages) {
		return r.scrollablePages[r.activeTabIndex]
	}
	return nil
}

func (r *TabbedReport) fixedHeaderHeight() int {
	height := 0
	if r.Title != "" {
		height++
	}
	if r.Subtitle != "" {
		height++
	}
	if r.Helper != "" {
		height++
	}
	height += helpTextHeight
	height += tabBarHeight
	return height
}

func (r *TabbedReport) View() string {
	if r.isQuitting {
		return ""
	}

	var output strings.Builder

	r.renderTitle(&output)
	r.renderSubtitle(&output)
	r.renderHelper(&output)
	r.renderHelpText(&output)
	r.renderTabBar(&output)
	r.renderSeparator(&output)
	r.renderActivePageContent(&output)

	return output.String()
}

func (r *TabbedReport) renderTitle(output *strings.Builder) {
	if r.Title != "" {
		output.WriteString(titleStyle.Render(r.Title))
		output.WriteString("\n")
	}
}

func (r *TabbedReport) renderSubtitle(output *strings.Builder) {
	if r.Subtitle != "" {
		output.WriteString(subtitleStyle.Render(r.Subtitle))
		output.WriteString("\n")
	}
}

func (r *TabbedReport) renderHelper(output *strings.Builder) {
	if r.Helper != "" {
		output.WriteString(subtitleStyle.Render(r.Helper))
		output.WriteString("\n")
	}
}

func (r *TabbedReport) renderHelpText(output *strings.Builder) {
	scrollInfo := r.formatScrollInfo()
	copyStatus := r.formatCopyStatus()
	helpText := fmt.Sprintf("← → tabs • ↑↓ scroll • 'c' copy • q quit%s%s", scrollInfo, copyStatus)
	output.WriteString(helpStyle.Render(helpText))
	output.WriteString("\n")
}

func (r *TabbedReport) formatScrollInfo() string {
	activeScrollable := r.getActiveScrollable()
	if activeScrollable == nil {
		return ""
	}
	scrollPercent := activeScrollable.ScrollPercent() * 100
	return fmt.Sprintf(" • %.0f%%", scrollPercent)
}

func (r *TabbedReport) formatCopyStatus() string {
	if r.hasCopied {
		return " • ✓ Copied!"
	}
	return ""
}

func (r *TabbedReport) renderTabBar(output *strings.Builder) {
	if len(r.Pages) == 0 {
		return
	}

	var tabs []string

	if r.activeTabIndex > 0 {
		tabs = append(tabs, tabArrowStyle.Render("◀ "))
	} else {
		tabs = append(tabs, "  ")
	}

	for i, page := range r.Pages {
		tabName := page.Name
		if tabName == "" {
			tabName = fmt.Sprintf("Tab %d", i+1)
		}
		if i == r.activeTabIndex {
			tabs = append(tabs, activeTabStyle.Render(tabName))
		} else {
			tabs = append(tabs, inactiveTabStyle.Render(tabName))
		}
		if i < len(r.Pages)-1 {
			tabs = append(tabs, " ")
		}
	}

	if r.activeTabIndex < len(r.Pages)-1 {
		tabs = append(tabs, tabArrowStyle.Render(" ▶"))
	}

	tabBar := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
	output.WriteString(tabBarStyle.Render(tabBar))
	output.WriteString("\n")
}

func (r *TabbedReport) renderSeparator(output *strings.Builder) {
	width := r.windowWidth - 2
	if width < minContentWidth {
		width = defaultWidth
	}
	line := strings.Repeat("─", width)
	output.WriteString(separatorStyle.Render(line))
	output.WriteString("\n")
}

func (r *TabbedReport) renderActivePageContent(output *strings.Builder) {
	if len(r.Pages) == 0 {
		output.WriteString(lipgloss.NewStyle().
			Foreground(colorGray).
			Italic(true).
			Render("No pages available"))
		return
	}

	activeScrollable := r.getActiveScrollable()
	if activeScrollable != nil {
		styledContent := contentStyle.Render(activeScrollable.View())
		output.WriteString(styledContent)
	}
}

func (r *TabbedReport) RenderFullContent() string {
	var output strings.Builder

	r.renderTitle(&output)
	r.renderSubtitle(&output)
	r.renderHelper(&output)
	output.WriteString("\n")

	for i, page := range r.Pages {
		r.renderPageHeader(&output, page, i)
		r.renderPageContent(&output, page)
	}

	return output.String()
}

func (r *TabbedReport) renderPageHeader(output *strings.Builder, page Page, index int) {
	pageName := page.Name
	if pageName == "" {
		pageName = fmt.Sprintf("Page %d", index+1)
	}

	boxWidth := 40
	separator := strings.Repeat("─", boxWidth)
	padding := boxWidth - len(pageName) - 2
	if padding < 0 {
		padding = 0
	}
	paddedName := pageName + strings.Repeat(" ", padding)

	output.WriteString(pageHeaderStyle.Render(fmt.Sprintf("┌%s┐", separator)))
	output.WriteString("\n")
	output.WriteString(pageHeaderStyle.Render(fmt.Sprintf("│ %s │", paddedName)))
	output.WriteString("\n")
	output.WriteString(pageHeaderStyle.Render(fmt.Sprintf("└%s┘", separator)))
	output.WriteString("\n\n")
}

func (r *TabbedReport) renderPageContent(output *strings.Builder, page Page) {
	if page.Content != nil {
		output.WriteString(page.Content.View())
		output.WriteString("\n\n")
	}
}

func removeAnsiCodes(text string) string {
	var result strings.Builder
	insideEscape := false

	for i := 0; i < len(text); i++ {
		if text[i] == '\x1b' {
			insideEscape = true
			continue
		}
		if insideEscape {
			if isAnsiTerminator(text[i]) {
				insideEscape = false
			}
			continue
		}
		result.WriteByte(text[i])
	}

	return result.String()
}

type ReportPages = TabbedReport
type InteractivePage = Page

func RunInteractivePages(report *TabbedReport) error {
	program := tea.NewProgram(
		report,
		tea.WithAltScreen(),
	)
	_, err := program.Run()
	if err != nil {
		return fmt.Errorf("cancelled")
	}

	fmt.Print("\n")
	fmt.Print(report.RenderFullContent())
	return nil
}

func RunInteractivePagesWithoutPrint(report *TabbedReport) error {
	program := tea.NewProgram(
		report,
		tea.WithAltScreen(),
	)
	_, err := program.Run()
	if err != nil {
		return fmt.Errorf("cancelled")
	}
	return nil
}

func PrintOnly(report *TabbedReport) {
	fmt.Print(report.RenderFullContent())
}
