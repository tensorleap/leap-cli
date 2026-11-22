package notification

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/samber/lo"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

// GetJobFailureNotifications fetches what explains a job's failure since the given start time:
// its error and warning notifications, and the reason attached to its status
func GetJobFailureNotifications(ctx context.Context, jobId string, since time.Time) (*JobFailureNotificationReport, error) {
	report := JobFailureNotificationReport{
		JobId:         jobId,
		Notifications: []NotificationReport{},
	}
	// Every level: explainsFailure decides, and a dead pod's reason is stored at Info
	notifications, err := GetNotificationsByFilter(ctx, NotificationFilter{
		JobId:    &jobId,
		FromDate: &since,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch job notifications: %w", err)
	}

	report.Notifications = lo.Map(lo.Filter(notifications, explainsFailure), func(notification Notification, _ int) NotificationReport {
		return ToNotificationReport(notification)
	})
	return &report, nil
}

// A job-status notification describes the status it was raised for, so only a failed one counts
// (not the queued notice), and only with its reason: node-server attaches one when a pod dies
// (out of memory, image pull, crash loop)
func explainsFailure(notification Notification, _ int) bool {
	if jobStatus := notification.MessageData.Params.JobMessageParams; jobStatus != nil {
		return api.IsJobFailed(jobStatus.JobStatus) && jobStatus.HasFailureInfo()
	}
	level := notification.MessageData.Level
	return level == MessageLevelError || level == MessageLevelWarning
}

// A failed status is an error whatever level it was stored at
func reportLevel(notification Notification) MessageLevel {
	if jobStatus := notification.MessageData.Params.JobMessageParams; jobStatus != nil && api.IsJobFailed(jobStatus.JobStatus) {
		return MessageLevelError
	}
	return notification.MessageData.Level
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
