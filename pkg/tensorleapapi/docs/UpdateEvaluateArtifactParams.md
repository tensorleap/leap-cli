# UpdateEvaluateArtifactParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**ProjectId** | **string** |  | 
**UpdateActions** | [**[]UpdateAction**](UpdateAction.md) |  | 
**CopyMode** | Pointer to **bool** |  | [optional] 
**SourceVersionId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateEvaluateArtifactParams

`func NewUpdateEvaluateArtifactParams(versionId string, projectId string, updateActions []UpdateAction, ) *UpdateEvaluateArtifactParams`

NewUpdateEvaluateArtifactParams instantiates a new UpdateEvaluateArtifactParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEvaluateArtifactParamsWithDefaults

`func NewUpdateEvaluateArtifactParamsWithDefaults() *UpdateEvaluateArtifactParams`

NewUpdateEvaluateArtifactParamsWithDefaults instantiates a new UpdateEvaluateArtifactParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *UpdateEvaluateArtifactParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *UpdateEvaluateArtifactParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *UpdateEvaluateArtifactParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetProjectId

`func (o *UpdateEvaluateArtifactParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateEvaluateArtifactParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateEvaluateArtifactParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetUpdateActions

`func (o *UpdateEvaluateArtifactParams) GetUpdateActions() []UpdateAction`

GetUpdateActions returns the UpdateActions field if non-nil, zero value otherwise.

### GetUpdateActionsOk

`func (o *UpdateEvaluateArtifactParams) GetUpdateActionsOk() (*[]UpdateAction, bool)`

GetUpdateActionsOk returns a tuple with the UpdateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateActions

`func (o *UpdateEvaluateArtifactParams) SetUpdateActions(v []UpdateAction)`

SetUpdateActions sets UpdateActions field to given value.


### GetCopyMode

`func (o *UpdateEvaluateArtifactParams) GetCopyMode() bool`

GetCopyMode returns the CopyMode field if non-nil, zero value otherwise.

### GetCopyModeOk

`func (o *UpdateEvaluateArtifactParams) GetCopyModeOk() (*bool, bool)`

GetCopyModeOk returns a tuple with the CopyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyMode

`func (o *UpdateEvaluateArtifactParams) SetCopyMode(v bool)`

SetCopyMode sets CopyMode field to given value.

### HasCopyMode

`func (o *UpdateEvaluateArtifactParams) HasCopyMode() bool`

HasCopyMode returns a boolean if a field has been set.

### GetSourceVersionId

`func (o *UpdateEvaluateArtifactParams) GetSourceVersionId() string`

GetSourceVersionId returns the SourceVersionId field if non-nil, zero value otherwise.

### GetSourceVersionIdOk

`func (o *UpdateEvaluateArtifactParams) GetSourceVersionIdOk() (*string, bool)`

GetSourceVersionIdOk returns a tuple with the SourceVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersionId

`func (o *UpdateEvaluateArtifactParams) SetSourceVersionId(v string)`

SetSourceVersionId sets SourceVersionId field to given value.

### HasSourceVersionId

`func (o *UpdateEvaluateArtifactParams) HasSourceVersionId() bool`

HasSourceVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


