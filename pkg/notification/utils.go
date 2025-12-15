package notification

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/samber/lo"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

// GetJobFailureNotifications fetches and prints error and warning notifications
// for a specific job since the given start time
func GetJobFailureNotifications(ctx context.Context, jobId string, since time.Time) (*JobFailureNotificationReport, error) {
	report := JobFailureNotificationReport{
		JobId:         jobId,
		Notifications: []NotificationReport{},
	}
	notifications, err := GetNotificationsByFilter(ctx, NotificationFilter{
		JobId:             &jobId,
		FromDate:          &since,
		NotificationTypes: []MessageLevel{MessageLevelError, MessageLevelWarning},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch job notifications: %w", err)
	}

	filteredNotifications := FilterJobStatusNotifications(notifications)

	report.Notifications = lo.Map(filteredNotifications, func(notification Notification, _ int) NotificationReport {
		return ToNotificationReport(notification)
	})
	return &report, nil
}

func FilterJobStatusNotifications(notifications []Notification) []Notification {
	return lo.Filter(notifications, func(notification Notification, _ int) bool {
		return notification.MessageData.Params.JobMessageParams == nil
	})
}

func levelWithColor(level MessageLevel) string {
	switch level {
	case MessageLevelError:
		return text.FgRed.Sprint("Error")
	case MessageLevelWarning:
		return text.FgYellow.Sprint("Warning")
	case MessageLevelInfo:
		return text.FgGreen.Sprint("Info")
	case MessageLevelVerbose:
		return text.FgHiBlack.Sprint("Verbose")
	}
	return "Unknown"
}

// normalizeMessage replaces escaped newlines with actual newlines
// and trims leading/trailing empty lines
func normalizeMessage(msg string) string {
	if msg == "" {
		return ""
	}
	// Replace escaped newlines with actual newlines
	replaced := strings.ReplaceAll(msg, "\\n", "\n")
	// Trim only leading/trailing newlines (not spaces)
	return strings.Trim(replaced, "\n\r")
}

// getNotificationMessage extracts the message from a notification.
// It checks multiple possible locations in order of priority:
// 1. JobMessageParams.FailureInfo.Message (job failure notifications)
// 2. ModuleMessageData.GeneralMessageParams.Message (general notifications)
// 3. ModuleMessageData.GeneralMessageParams.MessageCode (code snippet for general notifications)
func getNotificationMessage(notification tensorleapapi.Notification) (string, string) {
	params := notification.MessageData.Params

	// Check for job message params with failure info
	if params.JobMessageParams != nil {
		if params.JobMessageParams.HasFailureInfo() {
			failureInfo := params.JobMessageParams.GetFailureInfo()
			message := normalizeMessage(failureInfo.Message)
			extraMessage := ""
			if failureInfo.ExtraFailureData != nil {
				extraMessage = normalizeMessage(*failureInfo.ExtraFailureData)
			}
			return message, extraMessage
		}
	}

	// Check for module message data with general message params
	if params.ModuleMessageData != nil {
		if params.ModuleMessageData.GeneralMessageParams != nil {
			generalParams := params.ModuleMessageData.GeneralMessageParams

			// Build message from available fields
			message := ""
			extraMessage := ""

			if generalParams.Message != "" {
				message = normalizeMessage(generalParams.Message)
			}

			if generalParams.MessageCode != nil && *generalParams.MessageCode != "" {
				extraMessage = normalizeMessage(*generalParams.MessageCode)
			}

			return message, extraMessage
		}
	}

	return "--no message--", ""
}
