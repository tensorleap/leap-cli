package model

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tensorleap/leap-cli/pkg/interactive_pages"
	"github.com/tensorleap/leap-cli/pkg/notification"
	"github.com/tensorleap/leap-cli/pkg/run"
)

// ---------------------------
// Report Types
// ---------------------------

type ValidateAssetError struct {
	Error          string
	Name           string
	NodeId         string
	ConnectionName string
	Type           string
}

type ValidateAssetReport struct {
	GeneralError string
	Nodes        []ValidateAssetError
}

type ImportModelErrorReport struct {
	Notifications       *notification.JobFailureNotificationReport
	ValidateAssetReport *ValidateAssetReport
	TopLogs             []string
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

	pageLabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A3BE8C")).
			Bold(true)

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
// Show Methods
// ---------------------------

func (r *ValidateAssetReport) Show() {
	for _, node := range r.Nodes {
		fmt.Printf("Error: %s\n", node.Error)
		fmt.Printf("Name: %s\n", node.Name)
		fmt.Printf("Node ID: %s\n", node.NodeId)
		fmt.Printf("Connection Name: %s\n", node.ConnectionName)
		fmt.Printf("Type: %s\n", node.Type)
	}
	if r.GeneralError != "" {
		fmt.Printf("General Error: %s\n", r.GeneralError)
	}
}

func (r *ImportModelErrorReport) Show() {
	if err := r.ShowInteractive(); err != nil {
		// Fallback to simple output if interactive fails
		if r.ValidateAssetReport != nil {
			r.ValidateAssetReport.Show()
		}
		if r.TopLogs != nil {
			for _, logEntry := range r.TopLogs {
				fmt.Printf("Log: %s\n", logEntry)
			}
		}
	}
}

// ShowInteractive displays the report in an interactive tabbed interface
func (r *ImportModelErrorReport) ShowInteractive() error {
	reportPages := r.ToReportPages()
	return interactive_pages.RunInteractivePages(reportPages)
}

// ---------------------------
// Validation Errors Page Content
// ---------------------------
type ValidationErrorsPageContent struct {
	report *ValidateAssetReport
}

func (v *ValidationErrorsPageContent) Init() tea.Cmd { return nil }

func (v *ValidationErrorsPageContent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return v, nil
}

func (v *ValidationErrorsPageContent) View() string {
	if v.report == nil {
		return pageDimStyle.Render("No validation errors")
	}

	var b strings.Builder

	if v.report.GeneralError != "" {
		b.WriteString(pageSectionStyle.Render("General Error"))
		b.WriteString("\n")
		b.WriteString(pageErrorStyle.Render(v.report.GeneralError))
		b.WriteString("\n\n")
	}

	if len(v.report.Nodes) > 0 {
		// Group errors by type
		errorsByType := make(map[string][]ValidateAssetError)
		for _, node := range v.report.Nodes {
			errorsByType[node.Type] = append(errorsByType[node.Type], node)
		}

		for nodeType, errors := range errorsByType {
			b.WriteString(pageSectionStyle.Render(fmt.Sprintf("%s (%d errors)", nodeType, len(errors))))
			b.WriteString("\n\n")

			for i, node := range errors {
				b.WriteString(pageLabelStyle.Render(fmt.Sprintf("• %s", node.Name)))
				b.WriteString("\n")
				b.WriteString(fmt.Sprintf("  %s: %s\n",
					pageInfoStyle.Render("Error"),
					pageErrorStyle.Render(node.Error),
				))
				if node.NodeId != "" {
					b.WriteString(fmt.Sprintf("  %s: %s\n",
						pageInfoStyle.Render("Node ID"),
						pageValueStyle.Render(node.NodeId),
					))
				}
				if node.ConnectionName != "" {
					b.WriteString(fmt.Sprintf("  %s: %s\n",
						pageInfoStyle.Render("Connection"),
						pageValueStyle.Render(node.ConnectionName),
					))
				}
				if i < len(errors)-1 {
					b.WriteString("\n")
				}
			}
			b.WriteString("\n")
		}
	}

	if b.Len() == 0 {
		return pageDimStyle.Render("No validation errors")
	}

	return b.String()
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

// ---------------------------
// Conversion Function
// ---------------------------

// ToReportPages converts an ImportModelErrorReport to an interactive ReportPages
// for displaying in a tabbed interface
func (r *ImportModelErrorReport) ToReportPages() *interactive_pages.ReportPages {
	pages := make([]interactive_pages.InteractivePage, 0)

	if r.Notifications != nil && len(r.Notifications.Notifications) > 0 {
		pages = append(pages, interactive_pages.InteractivePage{
			ID:      "notifications",
			Name:    fmt.Sprintf("Error Notifications (%d)", len(r.Notifications.Notifications)),
			Content: &notification.NotificationsPageContent{Notifications: r.Notifications},
		})
	}

	// Add validation errors page if present
	if r.ValidateAssetReport != nil {
		hasContent := r.ValidateAssetReport.GeneralError != "" || len(r.ValidateAssetReport.Nodes) > 0

		if hasContent {
			errCount := len(r.ValidateAssetReport.Nodes)
			if r.ValidateAssetReport.GeneralError != "" {
				errCount++
			}
			pages = append(pages, interactive_pages.InteractivePage{
				ID:      "validation-errors",
				Name:    fmt.Sprintf("Validation Errors (%d)", errCount),
				Content: &ValidationErrorsPageContent{report: r.ValidateAssetReport},
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
		Title:    "Import Model Error Report",
		Subtitle: "Review the errors and warnings from the model import process",
		Helper:   "Use arrow keys to navigate tabs, scroll to view content, press 'c' to copy errors",
		Pages:    pages,
	}
}
