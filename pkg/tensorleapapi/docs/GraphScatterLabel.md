# GraphScatterLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**GraphVisualizations** | [**[]GraphViz**](GraphViz.md) |  | 
**InputName** | Pointer to **string** |  | [optional] 

## Methods

### NewGraphScatterLabel

`func NewGraphScatterLabel(type_ string, graphVisualizations []GraphViz, ) *GraphScatterLabel`

NewGraphScatterLabel instantiates a new GraphScatterLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphScatterLabelWithDefaults

`func NewGraphScatterLabelWithDefaults() *GraphScatterLabel`

NewGraphScatterLabelWithDefaults instantiates a new GraphScatterLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphScatterLabel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphScatterLabel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphScatterLabel) SetType(v string)`

SetType sets Type field to given value.


### GetGraphVisualizations

`func (o *GraphScatterLabel) GetGraphVisualizations() []GraphViz`

GetGraphVisualizations returns the GraphVisualizations field if non-nil, zero value otherwise.

### GetGraphVisualizationsOk

`func (o *GraphScatterLabel) GetGraphVisualizationsOk() (*[]GraphViz, bool)`

GetGraphVisualizationsOk returns a tuple with the GraphVisualizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphVisualizations

`func (o *GraphScatterLabel) SetGraphVisualizations(v []GraphViz)`

SetGraphVisualizations sets GraphVisualizations field to given value.


### GetInputName

`func (o *GraphScatterLabel) GetInputName() string`

GetInputName returns the InputName field if non-nil, zero value otherwise.

### GetInputNameOk

`func (o *GraphScatterLabel) GetInputNameOk() (*string, bool)`

GetInputNameOk returns a tuple with the InputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputName

`func (o *GraphScatterLabel) SetInputName(v string)`

SetInputName sets InputName field to given value.

### HasInputName

`func (o *GraphScatterLabel) HasInputName() bool`

HasInputName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


