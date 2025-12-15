package notification

import "fmt"


type JobFailureNotificationReport struct {
	JobId string
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
