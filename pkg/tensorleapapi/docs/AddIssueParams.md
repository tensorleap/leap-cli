# AddIssueParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**Title** | **string** |  | 
**Status** | [**IssueStatus**](IssueStatus.md) |  | 
**Branch** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Assignment** | Pointer to **string** |  | [optional] 
**Activities** | Pointer to [**[]IssueActivity**](IssueActivity.md) |  | [optional] 

## Methods

### NewAddIssueParams

`func NewAddIssueParams(projectId string, title string, status IssueStatus, ) *AddIssueParams`

NewAddIssueParams instantiates a new AddIssueParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIssueParamsWithDefaults

`func NewAddIssueParamsWithDefaults() *AddIssueParams`

NewAddIssueParamsWithDefaults instantiates a new AddIssueParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *AddIssueParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AddIssueParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AddIssueParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTitle

`func (o *AddIssueParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AddIssueParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AddIssueParams) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *AddIssueParams) GetStatus() IssueStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddIssueParams) GetStatusOk() (*IssueStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddIssueParams) SetStatus(v IssueStatus)`

SetStatus sets Status field to given value.


### GetBranch

`func (o *AddIssueParams) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *AddIssueParams) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *AddIssueParams) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *AddIssueParams) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetTags

`func (o *AddIssueParams) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddIssueParams) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddIssueParams) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AddIssueParams) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAssignment

`func (o *AddIssueParams) GetAssignment() string`

GetAssignment returns the Assignment field if non-nil, zero value otherwise.

### GetAssignmentOk

`func (o *AddIssueParams) GetAssignmentOk() (*string, bool)`

GetAssignmentOk returns a tuple with the Assignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignment

`func (o *AddIssueParams) SetAssignment(v string)`

SetAssignment sets Assignment field to given value.

### HasAssignment

`func (o *AddIssueParams) HasAssignment() bool`

HasAssignment returns a boolean if a field has been set.

### GetActivities

`func (o *AddIssueParams) GetActivities() []IssueActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *AddIssueParams) GetActivitiesOk() (*[]IssueActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *AddIssueParams) SetActivities(v []IssueActivity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *AddIssueParams) HasActivities() bool`

HasActivities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


