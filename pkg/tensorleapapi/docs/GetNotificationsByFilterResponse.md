# GetNotificationsByFilterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notifications** | [**[]Notification**](Notification.md) |  | 

## Methods

### NewGetNotificationsByFilterResponse

`func NewGetNotificationsByFilterResponse(notifications []Notification, ) *GetNotificationsByFilterResponse`

NewGetNotificationsByFilterResponse instantiates a new GetNotificationsByFilterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationsByFilterResponseWithDefaults

`func NewGetNotificationsByFilterResponseWithDefaults() *GetNotificationsByFilterResponse`

NewGetNotificationsByFilterResponseWithDefaults instantiates a new GetNotificationsByFilterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifications

`func (o *GetNotificationsByFilterResponse) GetNotifications() []Notification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *GetNotificationsByFilterResponse) GetNotificationsOk() (*[]Notification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *GetNotificationsByFilterResponse) SetNotifications(v []Notification)`

SetNotifications sets Notifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


