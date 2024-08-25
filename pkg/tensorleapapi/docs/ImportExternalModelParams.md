# ImportExternalModelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SessionId** | **string** |  | 
**Epoch** | **float64** |  | 
**TransformInputs** | Pointer to **bool** |  | [optional] 

## Methods

### NewImportExternalModelParams

`func NewImportExternalModelParams(projectId string, sessionId string, epoch float64, ) *ImportExternalModelParams`

NewImportExternalModelParams instantiates a new ImportExternalModelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportExternalModelParamsWithDefaults

`func NewImportExternalModelParamsWithDefaults() *ImportExternalModelParams`

NewImportExternalModelParamsWithDefaults instantiates a new ImportExternalModelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ImportExternalModelParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ImportExternalModelParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ImportExternalModelParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSessionId

`func (o *ImportExternalModelParams) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ImportExternalModelParams) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ImportExternalModelParams) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetEpoch

`func (o *ImportExternalModelParams) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *ImportExternalModelParams) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *ImportExternalModelParams) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetTransformInputs

`func (o *ImportExternalModelParams) GetTransformInputs() bool`

GetTransformInputs returns the TransformInputs field if non-nil, zero value otherwise.

### GetTransformInputsOk

`func (o *ImportExternalModelParams) GetTransformInputsOk() (*bool, bool)`

GetTransformInputsOk returns a tuple with the TransformInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformInputs

`func (o *ImportExternalModelParams) SetTransformInputs(v bool)`

SetTransformInputs sets TransformInputs field to given value.

### HasTransformInputs

`func (o *ImportExternalModelParams) HasTransformInputs() bool`

HasTransformInputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


