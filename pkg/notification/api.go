package notification

import (
	"context"
	"time"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

// Type aliases for convenience
type Notification = tensorleapapi.Notification
type MessageLevel = tensorleapapi.MessageLevel

// MessageLevel constants
const (
	MessageLevelError   = tensorleapapi.MESSAGELEVEL_ERROR
	MessageLevelWarning = tensorleapapi.MESSAGELEVEL_WARNING
	MessageLevelInfo    = tensorleapapi.MESSAGELEVEL_INFO
	MessageLevelVerbose = tensorleapapi.MESSAGELEVEL_VERBOSE
)

// NotificationFilter holds filter options for fetching notifications
type NotificationFilter struct {
	JobId             *string
	FromDate          *time.Time
	NotificationTypes []MessageLevel
}

// GetNotifications fetches all notifications for the current user
func GetNotifications(ctx context.Context) ([]Notification, error) {
	res, response, err := api.ApiClient.GetNotifications(ctx).Execute()
	if err := api.CheckRes(response, err); err != nil {
		return nil, err
	}
	return res.GetNotifications(), nil
}

// GetNotificationsByFilter fetches notifications with optional filters
func GetNotificationsByFilter(ctx context.Context, filter NotificationFilter) ([]Notification, error) {
	params := tensorleapapi.NewGetNotificationsByFilterParams()

	if filter.JobId != nil {
		params.SetJobId(*filter.JobId)
	}

	if filter.FromDate != nil {
		params.SetFromDate(filter.FromDate.Format(time.RFC3339))
	}

	if len(filter.NotificationTypes) > 0 {
		params.SetNotificationTypes(filter.NotificationTypes)
	}

	res, response, err := api.ApiClient.GetNotificationsByFilter(ctx).
		GetNotificationsByFilterParams(*params).
		Execute()
	if err := api.CheckRes(response, err); err != nil {
		return nil, err
	}
	return res.GetNotifications(), nil
}

// GetNotificationsSince fetches notifications created on or after the specified date
func GetNotificationsSince(ctx context.Context, since time.Time) ([]Notification, error) {
	return GetNotificationsByFilter(ctx, NotificationFilter{
		FromDate: &since,
	})
}

// GetNotificationsByJobId fetches notifications for a specific job
func GetNotificationsByJobId(ctx context.Context, jobId string) ([]Notification, error) {
	return GetNotificationsByFilter(ctx, NotificationFilter{
		JobId: &jobId,
	})
}

// GetNotificationsByType fetches notifications of specific types
func GetNotificationsByType(ctx context.Context, types []MessageLevel) ([]Notification, error) {
	return GetNotificationsByFilter(ctx, NotificationFilter{
		NotificationTypes: types,
	})
}

// GetNotificationsByJobIdSince fetches notifications for a specific job since a date
func GetNotificationsByJobIdSince(ctx context.Context, jobId string, since time.Time) ([]Notification, error) {
	return GetNotificationsByFilter(ctx, NotificationFilter{
		JobId:    &jobId,
		FromDate: &since,
	})
}

// GetErrorNotifications fetches only error notifications
func GetErrorNotifications(ctx context.Context) ([]Notification, error) {
	return GetNotificationsByType(ctx, []MessageLevel{MessageLevelError})
}

// GetErrorNotificationsSince fetches error notifications since a date
func GetErrorNotificationsSince(ctx context.Context, since time.Time) ([]Notification, error) {
	return GetNotificationsByFilter(ctx, NotificationFilter{
		FromDate:          &since,
		NotificationTypes: []MessageLevel{MessageLevelError},
	})
}

// GetWarningAndErrorNotifications fetches warning and error notifications
func GetWarningAndErrorNotifications(ctx context.Context) ([]Notification, error) {
	return GetNotificationsByType(ctx, []MessageLevel{MessageLevelError, MessageLevelWarning})
}
