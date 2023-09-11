# ApplyMappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MappingErrors** | Pointer to [**[]MappingError**](MappingError.md) |  | [optional] 
**ModelGraph** | Pointer to [**ModelGraph**](ModelGraph.md) |  | [optional] 

## Methods

### NewApplyMappingResponse

`func NewApplyMappingResponse() *ApplyMappingResponse`

NewApplyMappingResponse instantiates a new ApplyMappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyMappingResponseWithDefaults

`func NewApplyMappingResponseWithDefaults() *ApplyMappingResponse`

NewApplyMappingResponseWithDefaults instantiates a new ApplyMappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappingErrors

`func (o *ApplyMappingResponse) GetMappingErrors() []MappingError`

GetMappingErrors returns the MappingErrors field if non-nil, zero value otherwise.

### GetMappingErrorsOk

`func (o *ApplyMappingResponse) GetMappingErrorsOk() (*[]MappingError, bool)`

GetMappingErrorsOk returns a tuple with the MappingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingErrors

`func (o *ApplyMappingResponse) SetMappingErrors(v []MappingError)`

SetMappingErrors sets MappingErrors field to given value.

### HasMappingErrors

`func (o *ApplyMappingResponse) HasMappingErrors() bool`

HasMappingErrors returns a boolean if a field has been set.

### GetModelGraph

`func (o *ApplyMappingResponse) GetModelGraph() ModelGraph`

GetModelGraph returns the ModelGraph field if non-nil, zero value otherwise.

### GetModelGraphOk

`func (o *ApplyMappingResponse) GetModelGraphOk() (*ModelGraph, bool)`

GetModelGraphOk returns a tuple with the ModelGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelGraph

`func (o *ApplyMappingResponse) SetModelGraph(v ModelGraph)`

SetModelGraph sets ModelGraph field to given value.

### HasModelGraph

`func (o *ApplyMappingResponse) HasModelGraph() bool`

HasModelGraph returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


