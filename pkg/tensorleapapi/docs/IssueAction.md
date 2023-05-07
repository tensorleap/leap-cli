# IssueAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignment** | Pointer to [**ActionAssignmentElement**](ActionAssignmentElement.md) |  | [optional] 
**TagsAdded** | Pointer to [**ActionTagElement**](ActionTagElement.md) |  | [optional] 
**Comment** | Pointer to [**ActionCommentElement**](ActionCommentElement.md) |  | [optional] 

## Methods

### NewIssueAction

`func NewIssueAction() *IssueAction`

NewIssueAction instantiates a new IssueAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueActionWithDefaults

`func NewIssueActionWithDefaults() *IssueAction`

NewIssueActionWithDefaults instantiates a new IssueAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignment

`func (o *IssueAction) GetAssignment() ActionAssignmentElement`

GetAssignment returns the Assignment field if non-nil, zero value otherwise.

### GetAssignmentOk

`func (o *IssueAction) GetAssignmentOk() (*ActionAssignmentElement, bool)`

GetAssignmentOk returns a tuple with the Assignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignment

`func (o *IssueAction) SetAssignment(v ActionAssignmentElement)`

SetAssignment sets Assignment field to given value.

### HasAssignment

`func (o *IssueAction) HasAssignment() bool`

HasAssignment returns a boolean if a field has been set.

### GetTagsAdded

`func (o *IssueAction) GetTagsAdded() ActionTagElement`

GetTagsAdded returns the TagsAdded field if non-nil, zero value otherwise.

### GetTagsAddedOk

`func (o *IssueAction) GetTagsAddedOk() (*ActionTagElement, bool)`

GetTagsAddedOk returns a tuple with the TagsAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsAdded

`func (o *IssueAction) SetTagsAdded(v ActionTagElement)`

SetTagsAdded sets TagsAdded field to given value.

### HasTagsAdded

`func (o *IssueAction) HasTagsAdded() bool`

HasTagsAdded returns a boolean if a field has been set.

### GetComment

`func (o *IssueAction) GetComment() ActionCommentElement`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IssueAction) GetCommentOk() (*ActionCommentElement, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IssueAction) SetComment(v ActionCommentElement)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IssueAction) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


