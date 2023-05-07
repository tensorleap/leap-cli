# GraphData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **[][]float64** |  | 
**Heatmap** | Pointer to [**Heatmap**](Heatmap.md) |  | [optional] 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 

## Methods

### NewGraphData

`func NewGraphData(body [][]float64, type_ DataTypeEnum, ) *GraphData`

NewGraphData instantiates a new GraphData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphDataWithDefaults

`func NewGraphDataWithDefaults() *GraphData`

NewGraphDataWithDefaults instantiates a new GraphData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *GraphData) GetBody() [][]float64`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GraphData) GetBodyOk() (*[][]float64, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GraphData) SetBody(v [][]float64)`

SetBody sets Body field to given value.


### GetHeatmap

`func (o *GraphData) GetHeatmap() Heatmap`

GetHeatmap returns the Heatmap field if non-nil, zero value otherwise.

### GetHeatmapOk

`func (o *GraphData) GetHeatmapOk() (*Heatmap, bool)`

GetHeatmapOk returns a tuple with the Heatmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmap

`func (o *GraphData) SetHeatmap(v Heatmap)`

SetHeatmap sets Heatmap field to given value.

### HasHeatmap

`func (o *GraphData) HasHeatmap() bool`

HasHeatmap returns a boolean if a field has been set.

### GetType

`func (o *GraphData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


