# ImportModelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** |  | 
**ModelType** | [**ImportModelType**](ImportModelType.md) |  | 
**TransformInputs** | Pointer to **bool** |  | [optional] 

## Methods

### NewImportModelInfo

`func NewImportModelInfo(fileName string, modelType ImportModelType, ) *ImportModelInfo`

NewImportModelInfo instantiates a new ImportModelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportModelInfoWithDefaults

`func NewImportModelInfoWithDefaults() *ImportModelInfo`

NewImportModelInfoWithDefaults instantiates a new ImportModelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *ImportModelInfo) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ImportModelInfo) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ImportModelInfo) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetModelType

`func (o *ImportModelInfo) GetModelType() ImportModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *ImportModelInfo) GetModelTypeOk() (*ImportModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *ImportModelInfo) SetModelType(v ImportModelType)`

SetModelType sets ModelType field to given value.


### GetTransformInputs

`func (o *ImportModelInfo) GetTransformInputs() bool`

GetTransformInputs returns the TransformInputs field if non-nil, zero value otherwise.

### GetTransformInputsOk

`func (o *ImportModelInfo) GetTransformInputsOk() (*bool, bool)`

GetTransformInputsOk returns a tuple with the TransformInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformInputs

`func (o *ImportModelInfo) SetTransformInputs(v bool)`

SetTransformInputs sets TransformInputs field to given value.

### HasTransformInputs

`func (o *ImportModelInfo) HasTransformInputs() bool`

HasTransformInputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


