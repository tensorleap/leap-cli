package notification

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type JobFailureNotificationReport struct {
	JobId         string
	Notifications []NotificationReport
}

func (r *JobFailureNotificationReport) View() string {
	view := ""
	for _, notification := range r.Notifications {
		view += notification.View() + "\n"
	}
	return view
}

type NotificationReport struct {
	Level        MessageLevel
	Title        string
	Message      string
	ExtraMessage string
}

func ToNotificationReport(notification Notification) NotificationReport {
	message, extraMessage := getNotificationMessage(notification)
	return NotificationReport{
		Level:        notification.MessageData.Level,
		Title:        notification.Title,
		Message:      message,
		ExtraMessage: extraMessage,
	}
}

func (report NotificationReport) View() string {
	msg := fmt.Sprintf("  %s: %s - %s", levelWithColor(report.Level), report.Title, report.Message)
	return msg
}

// ---------------------------
// Page Content Styles
// ---------------------------
var (
	pageErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5555")).
			Bold(true)

	pageWarningStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFAA00"))

	pageInfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#88C0D0"))

	pageLabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A3BE8C")).
			Bold(true)

	// Style for main notification message - more prominent
	pageMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ECEFF4")).
				Bold(true)

	// Style for extra/secondary message - less prominent
	pageDimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6C6C6C")).
			Italic(true)

	pageSectionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#81A1C1")).
				Bold(true).
				Underline(true)
)

// ---------------------------
// Notifications Page Content
// ---------------------------
type NotificationsPageContent struct {
	Notifications *JobFailureNotificationReport
}

func (n *NotificationsPageContent) Init() tea.Cmd { return nil }

func (n *NotificationsPageContent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return n, nil
}

func (n *NotificationsPageContent) View() string {
	if n.Notifications == nil || len(n.Notifications.Notifications) == 0 {
		return pageDimStyle.Render("No error notifications")
	}

	var b strings.Builder
	total := len(n.Notifications.Notifications)

	if total > 1 {
		b.WriteString(formatNotificationSummary(n.Notifications.Notifications))
		b.WriteString("\n\n")
	}

	for i, notif := range n.Notifications.Notifications {
		levelStr := getLevelStyle(notif.Level).Render(fmt.Sprintf("[%s]", notif.Level))
		title := pageLabelStyle.Render(notif.Title)

		if total > 1 {
			b.WriteString(fmt.Sprintf("[%d/%d] %s %s\n", i+1, total, levelStr, title))
		} else {
			b.WriteString(fmt.Sprintf("• %s %s\n", levelStr, title))
		}
		if notif.Message != "" {
			b.WriteString(fmt.Sprintf("  %s\n", pageMessageStyle.Render(notif.Message)))
		}
		if notif.ExtraMessage != "" {
			b.WriteString(fmt.Sprintf("  %s\n", pageDimStyle.Render(notif.ExtraMessage)))
		}
		if i < total-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

func formatNotificationSummary(notifications []NotificationReport) string {
	errorCount := 0
	warningCount := 0
	infoCount := 0

	for _, notif := range notifications {
		switch notif.Level {
		case MessageLevelError:
			errorCount++
		case MessageLevelWarning:
			warningCount++
		default:
			infoCount++
		}
	}

	var parts []string
	if errorCount > 0 {
		parts = append(parts, pageErrorStyle.Render(fmt.Sprintf("%d Error", errorCount)))
	}
	if warningCount > 0 {
		parts = append(parts, pageWarningStyle.Render(fmt.Sprintf("%d Warning", warningCount)))
	}
	if infoCount > 0 {
		parts = append(parts, pageInfoStyle.Render(fmt.Sprintf("%d Info", infoCount)))
	}

	return strings.Join(parts, " ")
}

func getLevelStyle(level MessageLevel) lipgloss.Style {
	switch level {
	case MessageLevelError:
		return pageErrorStyle
	case MessageLevelWarning:
		return pageWarningStyle
	default:
		return pageInfoStyle
	}
}
