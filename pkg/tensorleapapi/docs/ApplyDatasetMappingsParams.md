# ApplyDatasetMappingsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelGraph** | [**ModelGraph**](ModelGraph.md) |  | 
**Yaml** | **string** |  | 
**DatasetVersionId** | Pointer to **string** |  | [optional] 

## Methods

### NewApplyDatasetMappingsParams

`func NewApplyDatasetMappingsParams(modelGraph ModelGraph, yaml string, ) *ApplyDatasetMappingsParams`

NewApplyDatasetMappingsParams instantiates a new ApplyDatasetMappingsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyDatasetMappingsParamsWithDefaults

`func NewApplyDatasetMappingsParamsWithDefaults() *ApplyDatasetMappingsParams`

NewApplyDatasetMappingsParamsWithDefaults instantiates a new ApplyDatasetMappingsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelGraph

`func (o *ApplyDatasetMappingsParams) GetModelGraph() ModelGraph`

GetModelGraph returns the ModelGraph field if non-nil, zero value otherwise.

### GetModelGraphOk

`func (o *ApplyDatasetMappingsParams) GetModelGraphOk() (*ModelGraph, bool)`

GetModelGraphOk returns a tuple with the ModelGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelGraph

`func (o *ApplyDatasetMappingsParams) SetModelGraph(v ModelGraph)`

SetModelGraph sets ModelGraph field to given value.


### GetYaml

`func (o *ApplyDatasetMappingsParams) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *ApplyDatasetMappingsParams) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *ApplyDatasetMappingsParams) SetYaml(v string)`

SetYaml sets Yaml field to given value.


### GetDatasetVersionId

`func (o *ApplyDatasetMappingsParams) GetDatasetVersionId() string`

GetDatasetVersionId returns the DatasetVersionId field if non-nil, zero value otherwise.

### GetDatasetVersionIdOk

`func (o *ApplyDatasetMappingsParams) GetDatasetVersionIdOk() (*string, bool)`

GetDatasetVersionIdOk returns a tuple with the DatasetVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetVersionId

`func (o *ApplyDatasetMappingsParams) SetDatasetVersionId(v string)`

SetDatasetVersionId sets DatasetVersionId field to given value.

### HasDatasetVersionId

`func (o *ApplyDatasetMappingsParams) HasDatasetVersionId() bool`

HasDatasetVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


