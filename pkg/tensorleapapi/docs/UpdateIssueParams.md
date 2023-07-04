# UpdateIssueParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**ProjectId** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**IssueStatus**](IssueStatus.md) |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Assignment** | Pointer to **string** |  | [optional] 
**Activities** | Pointer to [**[]IssueActivity**](IssueActivity.md) |  | [optional] 

## Methods

### NewUpdateIssueParams

`func NewUpdateIssueParams(cid string, projectId string, ) *UpdateIssueParams`

NewUpdateIssueParams instantiates a new UpdateIssueParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIssueParamsWithDefaults

`func NewUpdateIssueParamsWithDefaults() *UpdateIssueParams`

NewUpdateIssueParamsWithDefaults instantiates a new UpdateIssueParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *UpdateIssueParams) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *UpdateIssueParams) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *UpdateIssueParams) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProjectId

`func (o *UpdateIssueParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateIssueParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateIssueParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTitle

`func (o *UpdateIssueParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateIssueParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateIssueParams) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateIssueParams) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateIssueParams) GetStatus() IssueStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateIssueParams) GetStatusOk() (*IssueStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateIssueParams) SetStatus(v IssueStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateIssueParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBranch

`func (o *UpdateIssueParams) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *UpdateIssueParams) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *UpdateIssueParams) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *UpdateIssueParams) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetTags

`func (o *UpdateIssueParams) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateIssueParams) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateIssueParams) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateIssueParams) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAssignment

`func (o *UpdateIssueParams) GetAssignment() string`

GetAssignment returns the Assignment field if non-nil, zero value otherwise.

### GetAssignmentOk

`func (o *UpdateIssueParams) GetAssignmentOk() (*string, bool)`

GetAssignmentOk returns a tuple with the Assignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignment

`func (o *UpdateIssueParams) SetAssignment(v string)`

SetAssignment sets Assignment field to given value.

### HasAssignment

`func (o *UpdateIssueParams) HasAssignment() bool`

HasAssignment returns a boolean if a field has been set.

### GetActivities

`func (o *UpdateIssueParams) GetActivities() []IssueActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *UpdateIssueParams) GetActivitiesOk() (*[]IssueActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *UpdateIssueParams) SetActivities(v []IssueActivity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *UpdateIssueParams) HasActivities() bool`

HasActivities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


