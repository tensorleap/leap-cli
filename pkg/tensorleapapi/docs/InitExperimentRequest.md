# InitExperimentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**CodeIntegrationName** | Pointer to **string** |  | [optional] 
**ExperimentName** | **string** |  | 
**RemoveUntaggedUploadedModels** | Pointer to **bool** |  | [optional] 
**CodeIntegrationId** | Pointer to **string** |  | [optional] 
**CodeIntegrationBranch** | Pointer to **string** |  | [optional] 

## Methods

### NewInitExperimentRequest

`func NewInitExperimentRequest(experimentName string, ) *InitExperimentRequest`

NewInitExperimentRequest instantiates a new InitExperimentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitExperimentRequestWithDefaults

`func NewInitExperimentRequestWithDefaults() *InitExperimentRequest`

NewInitExperimentRequestWithDefaults instantiates a new InitExperimentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *InitExperimentRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *InitExperimentRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *InitExperimentRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *InitExperimentRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *InitExperimentRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *InitExperimentRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *InitExperimentRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *InitExperimentRequest) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetCodeIntegrationName

`func (o *InitExperimentRequest) GetCodeIntegrationName() string`

GetCodeIntegrationName returns the CodeIntegrationName field if non-nil, zero value otherwise.

### GetCodeIntegrationNameOk

`func (o *InitExperimentRequest) GetCodeIntegrationNameOk() (*string, bool)`

GetCodeIntegrationNameOk returns a tuple with the CodeIntegrationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIntegrationName

`func (o *InitExperimentRequest) SetCodeIntegrationName(v string)`

SetCodeIntegrationName sets CodeIntegrationName field to given value.

### HasCodeIntegrationName

`func (o *InitExperimentRequest) HasCodeIntegrationName() bool`

HasCodeIntegrationName returns a boolean if a field has been set.

### GetExperimentName

`func (o *InitExperimentRequest) GetExperimentName() string`

GetExperimentName returns the ExperimentName field if non-nil, zero value otherwise.

### GetExperimentNameOk

`func (o *InitExperimentRequest) GetExperimentNameOk() (*string, bool)`

GetExperimentNameOk returns a tuple with the ExperimentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentName

`func (o *InitExperimentRequest) SetExperimentName(v string)`

SetExperimentName sets ExperimentName field to given value.


### GetRemoveUntaggedUploadedModels

`func (o *InitExperimentRequest) GetRemoveUntaggedUploadedModels() bool`

GetRemoveUntaggedUploadedModels returns the RemoveUntaggedUploadedModels field if non-nil, zero value otherwise.

### GetRemoveUntaggedUploadedModelsOk

`func (o *InitExperimentRequest) GetRemoveUntaggedUploadedModelsOk() (*bool, bool)`

GetRemoveUntaggedUploadedModelsOk returns a tuple with the RemoveUntaggedUploadedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveUntaggedUploadedModels

`func (o *InitExperimentRequest) SetRemoveUntaggedUploadedModels(v bool)`

SetRemoveUntaggedUploadedModels sets RemoveUntaggedUploadedModels field to given value.

### HasRemoveUntaggedUploadedModels

`func (o *InitExperimentRequest) HasRemoveUntaggedUploadedModels() bool`

HasRemoveUntaggedUploadedModels returns a boolean if a field has been set.

### GetCodeIntegrationId

`func (o *InitExperimentRequest) GetCodeIntegrationId() string`

GetCodeIntegrationId returns the CodeIntegrationId field if non-nil, zero value otherwise.

### GetCodeIntegrationIdOk

`func (o *InitExperimentRequest) GetCodeIntegrationIdOk() (*string, bool)`

GetCodeIntegrationIdOk returns a tuple with the CodeIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIntegrationId

`func (o *InitExperimentRequest) SetCodeIntegrationId(v string)`

SetCodeIntegrationId sets CodeIntegrationId field to given value.

### HasCodeIntegrationId

`func (o *InitExperimentRequest) HasCodeIntegrationId() bool`

HasCodeIntegrationId returns a boolean if a field has been set.

### GetCodeIntegrationBranch

`func (o *InitExperimentRequest) GetCodeIntegrationBranch() string`

GetCodeIntegrationBranch returns the CodeIntegrationBranch field if non-nil, zero value otherwise.

### GetCodeIntegrationBranchOk

`func (o *InitExperimentRequest) GetCodeIntegrationBranchOk() (*string, bool)`

GetCodeIntegrationBranchOk returns a tuple with the CodeIntegrationBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIntegrationBranch

`func (o *InitExperimentRequest) SetCodeIntegrationBranch(v string)`

SetCodeIntegrationBranch sets CodeIntegrationBranch field to given value.

### HasCodeIntegrationBranch

`func (o *InitExperimentRequest) HasCodeIntegrationBranch() bool`

HasCodeIntegrationBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


