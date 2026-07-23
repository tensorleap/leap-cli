# AudioDataVisual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blob** | **string** |  | 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 
**Body** | **[][]float64** |  | 
**Heatmap** | Pointer to [**Heatmap**](Heatmap.md) |  | [optional] 
**XLabel** | Pointer to **string** |  | [optional] 
**YLabel** | Pointer to **string** |  | [optional] 
**XRange** | Pointer to **[]float64** |  | [optional] 
**Legend** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAudioDataVisual

`func NewAudioDataVisual(blob string, type_ DataTypeEnum, body [][]float64, ) *AudioDataVisual`

NewAudioDataVisual instantiates a new AudioDataVisual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioDataVisualWithDefaults

`func NewAudioDataVisualWithDefaults() *AudioDataVisual`

NewAudioDataVisualWithDefaults instantiates a new AudioDataVisual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlob

`func (o *AudioDataVisual) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *AudioDataVisual) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *AudioDataVisual) SetBlob(v string)`

SetBlob sets Blob field to given value.


### GetType

`func (o *AudioDataVisual) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AudioDataVisual) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AudioDataVisual) SetType(v DataTypeEnum)`

SetType sets Type field to given value.


### GetBody

`func (o *AudioDataVisual) GetBody() [][]float64`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AudioDataVisual) GetBodyOk() (*[][]float64, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AudioDataVisual) SetBody(v [][]float64)`

SetBody sets Body field to given value.


### GetHeatmap

`func (o *AudioDataVisual) GetHeatmap() Heatmap`

GetHeatmap returns the Heatmap field if non-nil, zero value otherwise.

### GetHeatmapOk

`func (o *AudioDataVisual) GetHeatmapOk() (*Heatmap, bool)`

GetHeatmapOk returns a tuple with the Heatmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmap

`func (o *AudioDataVisual) SetHeatmap(v Heatmap)`

SetHeatmap sets Heatmap field to given value.

### HasHeatmap

`func (o *AudioDataVisual) HasHeatmap() bool`

HasHeatmap returns a boolean if a field has been set.

### GetXLabel

`func (o *AudioDataVisual) GetXLabel() string`

GetXLabel returns the XLabel field if non-nil, zero value otherwise.

### GetXLabelOk

`func (o *AudioDataVisual) GetXLabelOk() (*string, bool)`

GetXLabelOk returns a tuple with the XLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLabel

`func (o *AudioDataVisual) SetXLabel(v string)`

SetXLabel sets XLabel field to given value.

### HasXLabel

`func (o *AudioDataVisual) HasXLabel() bool`

HasXLabel returns a boolean if a field has been set.

### GetYLabel

`func (o *AudioDataVisual) GetYLabel() string`

GetYLabel returns the YLabel field if non-nil, zero value otherwise.

### GetYLabelOk

`func (o *AudioDataVisual) GetYLabelOk() (*string, bool)`

GetYLabelOk returns a tuple with the YLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYLabel

`func (o *AudioDataVisual) SetYLabel(v string)`

SetYLabel sets YLabel field to given value.

### HasYLabel

`func (o *AudioDataVisual) HasYLabel() bool`

HasYLabel returns a boolean if a field has been set.

### GetXRange

`func (o *AudioDataVisual) GetXRange() []float64`

GetXRange returns the XRange field if non-nil, zero value otherwise.

### GetXRangeOk

`func (o *AudioDataVisual) GetXRangeOk() (*[]float64, bool)`

GetXRangeOk returns a tuple with the XRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXRange

`func (o *AudioDataVisual) SetXRange(v []float64)`

SetXRange sets XRange field to given value.

### HasXRange

`func (o *AudioDataVisual) HasXRange() bool`

HasXRange returns a boolean if a field has been set.

### GetLegend

`func (o *AudioDataVisual) GetLegend() []string`

GetLegend returns the Legend field if non-nil, zero value otherwise.

### GetLegendOk

`func (o *AudioDataVisual) GetLegendOk() (*[]string, bool)`

GetLegendOk returns a tuple with the Legend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegend

`func (o *AudioDataVisual) SetLegend(v []string)`

SetLegend sets Legend field to given value.

### HasLegend

`func (o *AudioDataVisual) HasLegend() bool`

HasLegend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


