# IssueActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**ActionType** | [**IssueActionType**](IssueActionType.md) |  | 
**Action** | [**IssueAction**](IssueAction.md) |  | 

## Methods

### NewIssueActivity

`func NewIssueActivity(createdAt time.Time, updatedAt time.Time, createdBy string, actionType IssueActionType, action IssueAction, ) *IssueActivity`

NewIssueActivity instantiates a new IssueActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueActivityWithDefaults

`func NewIssueActivityWithDefaults() *IssueActivity`

NewIssueActivityWithDefaults instantiates a new IssueActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IssueActivity) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueActivity) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueActivity) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *IssueActivity) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueActivity) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueActivity) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedBy

`func (o *IssueActivity) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *IssueActivity) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *IssueActivity) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetActionType

`func (o *IssueActivity) GetActionType() IssueActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *IssueActivity) GetActionTypeOk() (*IssueActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *IssueActivity) SetActionType(v IssueActionType)`

SetActionType sets ActionType field to given value.


### GetAction

`func (o *IssueActivity) GetAction() IssueAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IssueActivity) GetActionOk() (*IssueAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IssueActivity) SetAction(v IssueAction)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


