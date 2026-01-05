package code

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
	"github.com/tensorleap/leap-cli/pkg/interactive_pages"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/notification"
	"github.com/tensorleap/leap-cli/pkg/run"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func CollectErrorsOnCodeParseFailed(ctx context.Context, jobId string, codeSnapshot *tensorleapapi.CodeSnapshot, commandStartTime time.Time) (*CodeSnapshotParserErr, error) {
	report := CodeSnapshotParserErr{}
	var err error
	if jobId != "" {
		report.notifications, err = notification.GetJobFailureNotifications(ctx, jobId, commandStartTime)
		if err != nil {
			log.Warnf("Failed to get job failure notifications: %v", err)
		}
	}

	report.codeSnapshotParserErrs = CollectCodeSnapshotParserErr(codeSnapshot)

	report.TopLogs, err = GetLastEngineErrors(ctx, jobId, 3)
	if err != nil {
		log.Warnf("Failed to print last engine error: %v", err)

	}

	return &report, nil
}

func GetLastEngineErrors(ctx context.Context, jobId string, count int) ([]string, error) {
	logs, err := run.GetRunLogs(ctx, jobId)
	if err != nil {
		return nil, err
	}
	if len(logs) == 0 {
		return nil, nil
	}
	errorPattern := regexp.MustCompile(`(?i)"levelname":\s*"ERROR"`)
	topLogs := run.GetTopLogs(logs, errorPattern, count)
	return topLogs, nil
}

type CodeSnapshotParserErr struct {
	notifications          *notification.JobFailureNotificationReport
	codeSnapshotParserErrs *CodeSnapshotParserErrs
	TopLogs                []string
}

func (r *CodeSnapshotParserErr) Init() tea.Cmd {
	return nil
}

func (r *CodeSnapshotParserErr) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return r, nil
}

func (r *CodeSnapshotParserErr) View() string {
	view := ""
	if r.notifications != nil {
		view += r.notifications.View()
	}
	if r.codeSnapshotParserErrs != nil {
		view += r.codeSnapshotParserErrs.View()
	}
	if r.TopLogs != nil {
		for _, log := range r.TopLogs {
			view += log + "\n"
		}
	}
	return view
}

type BinderError struct {
	name    string
	display map[string]interface{}
}
type CodeSnapshotParserErrs struct {
	generalError  string
	printLog      string
	bindersErrors []BinderError
}

func (r *CodeSnapshotParserErrs) View() string {
	view := ""
	if r.generalError != "" {
		view += r.generalError + "\n"
	}
	if r.printLog != "" {
		view += r.printLog + "\n"
	}
	for _, binderError := range r.bindersErrors {
		view += binderError.name + "\n"
		display := ""
		for key, value := range binderError.display {
			display += fmt.Sprintf("%s: %v\n", key, value)
		}
		view += display + "\n"
	}
	return view
}

func CollectCodeSnapshotParserErr(civ *CodeSnapshot) *CodeSnapshotParserErrs {
	if civ == nil || civ.ParseResult == nil {
		return nil
	}
	errs := &CodeSnapshotParserErrs{}

	if civ.ParseResult.SetupStatus.GeneralError != nil {
		errs.generalError = *civ.ParseResult.SetupStatus.GeneralError
	}

	// Extract PrintLog if present
	if civ.ParseResult.SetupStatus.HasPrintLog() {
		errs.printLog = civ.ParseResult.SetupStatus.GetPrintLog()
	}

	hasErrors := lo.SomeBy(civ.ParseResult.SetupStatus.BindersStatus, func(binderStatus tensorleapapi.DatasetTestResultPayload) bool {
		return !binderStatus.IsPassed && len(binderStatus.Display) > 0
	})

	if hasErrors {
		for _, binderStatus := range civ.ParseResult.SetupStatus.BindersStatus {
			if !binderStatus.IsPassed && len(binderStatus.Display) > 0 {
				errs.bindersErrors = append(errs.bindersErrors, BinderError{
					name:    binderStatus.Name,
					display: binderStatus.Display,
				})
			}
		}
	}

	return errs
}

func formatBinderLine(err BinderError, index, total int, useErrorStyle bool) string {
	var text string
	if index > 0 && total > 1 {
		text = fmt.Sprintf("[%d/%d] %s", index, total, err.name)
	} else {
		text = err.name
	}

	if useErrorStyle {
		return pageErrorStyle.Render(text)
	}
	return pageDimStyle.Render(text)
}

// ---------------------------
// Page Content Styles
// ---------------------------
var (
	pageErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5555")).
			Bold(true)

	pageInfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#88C0D0"))

	pageValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#D8DEE9"))

	pageDimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6C6C6C")).
			Italic(true)

	pageSectionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#81A1C1")).
				Bold(true).
				Underline(true)
)

// ---------------------------
// Parser Errors Page Content
// ---------------------------
type ParserErrorsPageContent struct {
	parserErrs *CodeSnapshotParserErrs
}

func (p *ParserErrorsPageContent) Init() tea.Cmd { return nil }

func (p *ParserErrorsPageContent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}

func (p *ParserErrorsPageContent) View() string {
	if p.parserErrs == nil {
		return pageDimStyle.Render("No parser errors")
	}

	var b strings.Builder

	if p.parserErrs.generalError != "" {
		b.WriteString(pageSectionStyle.Render("General Error"))
		b.WriteString("\n")
		b.WriteString(pageErrorStyle.Render(p.parserErrs.generalError))
		b.WriteString("\n\n")
	}

	if p.parserErrs.printLog != "" {
		b.WriteString(pageSectionStyle.Render("Log Output"))
		b.WriteString("\n")
		b.WriteString(pageValueStyle.Render(strings.TrimSpace(p.parserErrs.printLog)))
		b.WriteString("\n")
	}

	if len(p.parserErrs.bindersErrors) > 0 {
		totalErrors := len(p.parserErrs.bindersErrors)
		b.WriteString(pageSectionStyle.Render(fmt.Sprintf("Binder Errors (%d)", totalErrors)))
		b.WriteString("\n")

		// Summary of failed binders - each on its own line with type
		if totalErrors > 1 {
			for _, err := range p.parserErrs.bindersErrors {
				b.WriteString(fmt.Sprintf("  • %s\n", formatBinderLine(err, 0, 0, false)))
			}
		}
		b.WriteString("\n")

		for i, binderErr := range p.parserErrs.bindersErrors {
			b.WriteString(formatBinderLine(binderErr, i+1, totalErrors, true))
			b.WriteString("\n")

			for key, value := range binderErr.display {
				b.WriteString(fmt.Sprintf("  %s: %s\n",
					pageInfoStyle.Render(key),
					pageValueStyle.Render(fmt.Sprintf("%v", value)),
				))
			}

			if i < totalErrors-1 {
				b.WriteString("\n")
			}
		}
	}

	if b.Len() == 0 {
		return pageDimStyle.Render("No parser errors")
	}

	return b.String()
}

// ---------------------------
// Conversion Function
// ---------------------------

// ToReportPages converts a CodeSnapshotParserErr to an interactive ReportPages
// for displaying in a tabbed interface
func (r *CodeSnapshotParserErr) ToReportPages() *interactive_pages.ReportPages {
	pages := make([]interactive_pages.InteractivePage, 0)

	if r.notifications != nil && len(r.notifications.Notifications) > 0 {
		pages = append(pages, interactive_pages.InteractivePage{
			ID:      "notifications",
			Name:    fmt.Sprintf("Error Notifications (%d)", len(r.notifications.Notifications)),
			Content: &notification.NotificationsPageContent{Notifications: r.notifications},
		})
	}

	// Add parser errors page if present
	if r.codeSnapshotParserErrs != nil {
		hasContent := r.codeSnapshotParserErrs.generalError != "" ||
			r.codeSnapshotParserErrs.printLog != "" ||
			len(r.codeSnapshotParserErrs.bindersErrors) > 0

		if hasContent {
			errCount := len(r.codeSnapshotParserErrs.bindersErrors)
			if r.codeSnapshotParserErrs.generalError != "" {
				errCount++
			}
			pages = append(pages, interactive_pages.InteractivePage{
				ID:      "parser-errors",
				Name:    fmt.Sprintf("Parser Errors (%d)", errCount),
				Content: &ParserErrorsPageContent{parserErrs: r.codeSnapshotParserErrs},
			})
		}
	}

	if len(r.TopLogs) > 0 {
		pages = append(pages, interactive_pages.InteractivePage{
			ID:      "logs",
			Name:    fmt.Sprintf("Error Logs (%d)", len(r.TopLogs)),
			Content: &run.LogsPageContent{Logs: r.TopLogs},
		})
	}

	// If no pages, add an empty state page
	if len(pages) == 0 {
		pages = append(pages, interactive_pages.InteractivePage{
			ID:      "empty",
			Name:    "No Errors",
			Content: &EmptyPageContent{message: "No errors or warnings were found."},
		})
	}

	return &interactive_pages.ReportPages{
		Title:    "Code Parser Error Report",
		Subtitle: "Review the errors and warnings from the code parsing process",
		Helper:   "Use arrow keys to navigate tabs, scroll to view content, press 'c' to copy errors",
		Pages:    pages,
	}
}

// ---------------------------
// Empty Page Content
// ---------------------------
type EmptyPageContent struct {
	message string
}

func (e *EmptyPageContent) Init() tea.Cmd { return nil }

func (e *EmptyPageContent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return e, nil
}

func (e *EmptyPageContent) View() string {
	return pageDimStyle.Render(e.message)
}

// ShowReportInteractive displays the report in an interactive tabbed interface
func ShowReportInteractive(report *CodeSnapshotParserErr) error {
	reportPages := report.ToReportPages()
	return interactive_pages.RunInteractivePages(reportPages)
}
