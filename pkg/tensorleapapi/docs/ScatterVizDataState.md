# ScatterVizDataState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataState** | Pointer to [**DataStateType**](DataStateType.md) |  | [optional] 
**ScatterData** | **[][]float64** |  | 
**Labels** | [**[]ScatterLabel**](ScatterLabel.md) |  | 
**Samples** | [**[]SampleIdentity**](SampleIdentity.md) |  | 
**Metadata** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 
**MiByCluster** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewScatterVizDataState

`func NewScatterVizDataState(scatterData [][]float64, labels []ScatterLabel, samples []SampleIdentity, metadata map[string]interface{}, ) *ScatterVizDataState`

NewScatterVizDataState instantiates a new ScatterVizDataState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScatterVizDataStateWithDefaults

`func NewScatterVizDataStateWithDefaults() *ScatterVizDataState`

NewScatterVizDataStateWithDefaults instantiates a new ScatterVizDataState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataState

`func (o *ScatterVizDataState) GetDataState() DataStateType`

GetDataState returns the DataState field if non-nil, zero value otherwise.

### GetDataStateOk

`func (o *ScatterVizDataState) GetDataStateOk() (*DataStateType, bool)`

GetDataStateOk returns a tuple with the DataState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataState

`func (o *ScatterVizDataState) SetDataState(v DataStateType)`

SetDataState sets DataState field to given value.

### HasDataState

`func (o *ScatterVizDataState) HasDataState() bool`

HasDataState returns a boolean if a field has been set.

### GetScatterData

`func (o *ScatterVizDataState) GetScatterData() [][]float64`

GetScatterData returns the ScatterData field if non-nil, zero value otherwise.

### GetScatterDataOk

`func (o *ScatterVizDataState) GetScatterDataOk() (*[][]float64, bool)`

GetScatterDataOk returns a tuple with the ScatterData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScatterData

`func (o *ScatterVizDataState) SetScatterData(v [][]float64)`

SetScatterData sets ScatterData field to given value.


### GetLabels

`func (o *ScatterVizDataState) GetLabels() []ScatterLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ScatterVizDataState) GetLabelsOk() (*[]ScatterLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ScatterVizDataState) SetLabels(v []ScatterLabel)`

SetLabels sets Labels field to given value.


### GetSamples

`func (o *ScatterVizDataState) GetSamples() []SampleIdentity`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ScatterVizDataState) GetSamplesOk() (*[]SampleIdentity, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ScatterVizDataState) SetSamples(v []SampleIdentity)`

SetSamples sets Samples field to given value.


### GetMetadata

`func (o *ScatterVizDataState) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ScatterVizDataState) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ScatterVizDataState) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetMiByCluster

`func (o *ScatterVizDataState) GetMiByCluster() map[string]interface{}`

GetMiByCluster returns the MiByCluster field if non-nil, zero value otherwise.

### GetMiByClusterOk

`func (o *ScatterVizDataState) GetMiByClusterOk() (*map[string]interface{}, bool)`

GetMiByClusterOk returns a tuple with the MiByCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiByCluster

`func (o *ScatterVizDataState) SetMiByCluster(v map[string]interface{})`

SetMiByCluster sets MiByCluster field to given value.

### HasMiByCluster

`func (o *ScatterVizDataState) HasMiByCluster() bool`

HasMiByCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


