# ImportNewModelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**FileName** | **string** |  | 
**DatasetId** | Pointer to **string** |  | [optional] 
**ModelName** | **string** |  | 
**VersionName** | **string** |  | 
**BranchName** | Pointer to **string** |  | [optional] 
**ModelType** | [**ImportModelType**](ImportModelType.md) |  | 
**TransformInputs** | Pointer to **bool** |  | [optional] 
**MappingYaml** | Pointer to **string** |  | [optional] 

## Methods

### NewImportNewModelParams

`func NewImportNewModelParams(projectId string, fileName string, modelName string, versionName string, modelType ImportModelType, ) *ImportNewModelParams`

NewImportNewModelParams instantiates a new ImportNewModelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportNewModelParamsWithDefaults

`func NewImportNewModelParamsWithDefaults() *ImportNewModelParams`

NewImportNewModelParamsWithDefaults instantiates a new ImportNewModelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ImportNewModelParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ImportNewModelParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ImportNewModelParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetFileName

`func (o *ImportNewModelParams) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ImportNewModelParams) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ImportNewModelParams) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetDatasetId

`func (o *ImportNewModelParams) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *ImportNewModelParams) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *ImportNewModelParams) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *ImportNewModelParams) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### GetModelName

`func (o *ImportNewModelParams) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ImportNewModelParams) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ImportNewModelParams) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetVersionName

`func (o *ImportNewModelParams) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *ImportNewModelParams) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *ImportNewModelParams) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.


### GetBranchName

`func (o *ImportNewModelParams) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *ImportNewModelParams) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *ImportNewModelParams) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *ImportNewModelParams) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetModelType

`func (o *ImportNewModelParams) GetModelType() ImportModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *ImportNewModelParams) GetModelTypeOk() (*ImportModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *ImportNewModelParams) SetModelType(v ImportModelType)`

SetModelType sets ModelType field to given value.


### GetTransformInputs

`func (o *ImportNewModelParams) GetTransformInputs() bool`

GetTransformInputs returns the TransformInputs field if non-nil, zero value otherwise.

### GetTransformInputsOk

`func (o *ImportNewModelParams) GetTransformInputsOk() (*bool, bool)`

GetTransformInputsOk returns a tuple with the TransformInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformInputs

`func (o *ImportNewModelParams) SetTransformInputs(v bool)`

SetTransformInputs sets TransformInputs field to given value.

### HasTransformInputs

`func (o *ImportNewModelParams) HasTransformInputs() bool`

HasTransformInputs returns a boolean if a field has been set.

### GetMappingYaml

`func (o *ImportNewModelParams) GetMappingYaml() string`

GetMappingYaml returns the MappingYaml field if non-nil, zero value otherwise.

### GetMappingYamlOk

`func (o *ImportNewModelParams) GetMappingYamlOk() (*string, bool)`

GetMappingYamlOk returns a tuple with the MappingYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingYaml

`func (o *ImportNewModelParams) SetMappingYaml(v string)`

SetMappingYaml sets MappingYaml field to given value.

### HasMappingYaml

`func (o *ImportNewModelParams) HasMappingYaml() bool`

HasMappingYaml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


