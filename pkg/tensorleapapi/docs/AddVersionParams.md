# AddVersionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**ModelGraph** | [**ModelGraph**](ModelGraph.md) |  | 
**BranchName** | **string** |  | 
**Description** | **string** |  | 
**CodeIntegrationVersionId** | Pointer to **string** |  | [optional] 
**DatasetSetup** | [**DatasetSetup**](DatasetSetup.md) |  | 
**Hash** | Pointer to **string** |  | [optional] 

## Methods

### NewAddVersionParams

`func NewAddVersionParams(projectId string, modelGraph ModelGraph, branchName string, description string, datasetSetup DatasetSetup, ) *AddVersionParams`

NewAddVersionParams instantiates a new AddVersionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVersionParamsWithDefaults

`func NewAddVersionParamsWithDefaults() *AddVersionParams`

NewAddVersionParamsWithDefaults instantiates a new AddVersionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *AddVersionParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AddVersionParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AddVersionParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetModelGraph

`func (o *AddVersionParams) GetModelGraph() ModelGraph`

GetModelGraph returns the ModelGraph field if non-nil, zero value otherwise.

### GetModelGraphOk

`func (o *AddVersionParams) GetModelGraphOk() (*ModelGraph, bool)`

GetModelGraphOk returns a tuple with the ModelGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelGraph

`func (o *AddVersionParams) SetModelGraph(v ModelGraph)`

SetModelGraph sets ModelGraph field to given value.


### GetBranchName

`func (o *AddVersionParams) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *AddVersionParams) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *AddVersionParams) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetDescription

`func (o *AddVersionParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddVersionParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddVersionParams) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCodeIntegrationVersionId

`func (o *AddVersionParams) GetCodeIntegrationVersionId() string`

GetCodeIntegrationVersionId returns the CodeIntegrationVersionId field if non-nil, zero value otherwise.

### GetCodeIntegrationVersionIdOk

`func (o *AddVersionParams) GetCodeIntegrationVersionIdOk() (*string, bool)`

GetCodeIntegrationVersionIdOk returns a tuple with the CodeIntegrationVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIntegrationVersionId

`func (o *AddVersionParams) SetCodeIntegrationVersionId(v string)`

SetCodeIntegrationVersionId sets CodeIntegrationVersionId field to given value.

### HasCodeIntegrationVersionId

`func (o *AddVersionParams) HasCodeIntegrationVersionId() bool`

HasCodeIntegrationVersionId returns a boolean if a field has been set.

### GetDatasetSetup

`func (o *AddVersionParams) GetDatasetSetup() DatasetSetup`

GetDatasetSetup returns the DatasetSetup field if non-nil, zero value otherwise.

### GetDatasetSetupOk

`func (o *AddVersionParams) GetDatasetSetupOk() (*DatasetSetup, bool)`

GetDatasetSetupOk returns a tuple with the DatasetSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetSetup

`func (o *AddVersionParams) SetDatasetSetup(v DatasetSetup)`

SetDatasetSetup sets DatasetSetup field to given value.


### GetHash

`func (o *AddVersionParams) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *AddVersionParams) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *AddVersionParams) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *AddVersionParams) HasHash() bool`

HasHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


