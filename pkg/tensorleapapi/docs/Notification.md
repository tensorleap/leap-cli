# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** |  | [optional] 
**Cid** | **string** |  | 
**TeamId** | **string** |  | 
**User** | **string** |  | 
**Title** | **string** |  | 
**MessageData** | [**CustomMessageData**](CustomMessageData.md) |  | 
**Identifier** | **string** |  | 
**Context** | [**JobNotificationContext**](JobNotificationContext.md) |  | 
**CreatedAt** | **time.Time** |  | 
**IsRead** | **bool** |  | 

## Methods

### NewNotification

`func NewNotification(cid string, teamId string, user string, title string, messageData CustomMessageData, identifier string, context JobNotificationContext, createdAt time.Time, isRead bool, ) *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *Notification) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Notification) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Notification) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Notification) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetCid

`func (o *Notification) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *Notification) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *Notification) SetCid(v string)`

SetCid sets Cid field to given value.


### GetTeamId

`func (o *Notification) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Notification) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Notification) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetUser

`func (o *Notification) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Notification) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Notification) SetUser(v string)`

SetUser sets User field to given value.


### GetTitle

`func (o *Notification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Notification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Notification) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMessageData

`func (o *Notification) GetMessageData() CustomMessageData`

GetMessageData returns the MessageData field if non-nil, zero value otherwise.

### GetMessageDataOk

`func (o *Notification) GetMessageDataOk() (*CustomMessageData, bool)`

GetMessageDataOk returns a tuple with the MessageData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageData

`func (o *Notification) SetMessageData(v CustomMessageData)`

SetMessageData sets MessageData field to given value.


### GetIdentifier

`func (o *Notification) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Notification) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Notification) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetContext

`func (o *Notification) GetContext() JobNotificationContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Notification) GetContextOk() (*JobNotificationContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Notification) SetContext(v JobNotificationContext)`

SetContext sets Context field to given value.


### GetCreatedAt

`func (o *Notification) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Notification) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Notification) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIsRead

`func (o *Notification) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *Notification) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *Notification) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


