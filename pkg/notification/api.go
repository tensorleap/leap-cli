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
