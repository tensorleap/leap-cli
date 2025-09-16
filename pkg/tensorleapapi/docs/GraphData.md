# GraphData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **[][]float64** |  | 
**Heatmap** | Pointer to [**Heatmap**](Heatmap.md) |  | [optional] 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 
**XLabel** | Pointer to **string** |  | [optional] 
**YLabel** | Pointer to **string** |  | [optional] 
**XRange** | Pointer to **[]float64** |  | [optional] 
**Legend** | Pointer to **[]string** |  | [optional] 

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


### GetXLabel

`func (o *GraphData) GetXLabel() string`

GetXLabel returns the XLabel field if non-nil, zero value otherwise.

### GetXLabelOk

`func (o *GraphData) GetXLabelOk() (*string, bool)`

GetXLabelOk returns a tuple with the XLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLabel

`func (o *GraphData) SetXLabel(v string)`

SetXLabel sets XLabel field to given value.

### HasXLabel

`func (o *GraphData) HasXLabel() bool`

HasXLabel returns a boolean if a field has been set.

### GetYLabel

`func (o *GraphData) GetYLabel() string`

GetYLabel returns the YLabel field if non-nil, zero value otherwise.

### GetYLabelOk

`func (o *GraphData) GetYLabelOk() (*string, bool)`

GetYLabelOk returns a tuple with the YLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYLabel

`func (o *GraphData) SetYLabel(v string)`

SetYLabel sets YLabel field to given value.

### HasYLabel

`func (o *GraphData) HasYLabel() bool`

HasYLabel returns a boolean if a field has been set.

### GetXRange

`func (o *GraphData) GetXRange() []float64`

GetXRange returns the XRange field if non-nil, zero value otherwise.

### GetXRangeOk

`func (o *GraphData) GetXRangeOk() (*[]float64, bool)`

GetXRangeOk returns a tuple with the XRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXRange

`func (o *GraphData) SetXRange(v []float64)`

SetXRange sets XRange field to given value.

### HasXRange

`func (o *GraphData) HasXRange() bool`

HasXRange returns a boolean if a field has been set.

### GetLegend

`func (o *GraphData) GetLegend() []string`

GetLegend returns the Legend field if non-nil, zero value otherwise.

### GetLegendOk

`func (o *GraphData) GetLegendOk() (*[]string, bool)`

GetLegendOk returns a tuple with the Legend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegend

`func (o *GraphData) SetLegend(v []string)`

SetLegend sets Legend field to given value.

### HasLegend

`func (o *GraphData) HasLegend() bool`

HasLegend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


