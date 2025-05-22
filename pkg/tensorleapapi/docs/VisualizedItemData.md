# VisualizedItemData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blob** | **string** |  | 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 
**HeatmapBlob** | **string** |  | 
**Body** | **[][]float64** |  | 
**Heatmap** | Pointer to [**Heatmap**](Heatmap.md) |  | [optional] 
**Labels** | **[]string** |  | 
**Gt** | Pointer to **[]float64** |  | [optional] 
**XLabel** | Pointer to **string** |  | [optional] 
**YLabel** | Pointer to **string** |  | [optional] 
**XRange** | Pointer to **[]float64** |  | [optional] 
**BoundingBox** | [**[]BoundingBox**](BoundingBox.md) |  | 
**MaskBlob** | **string** |  | 
**Text** | **[]string** |  | 
**Mask** | **[]float64** |  | 
**VideoBlob** | **string** |  | 
**Data** | [**[]CompositeVizItem**](CompositeVizItem.md) |  | 
**Algo** | Pointer to [**SampleAnalysisAlgo**](SampleAnalysisAlgo.md) |  | [optional] 

## Methods

### NewVisualizedItemData

`func NewVisualizedItemData(blob string, type_ DataTypeEnum, heatmapBlob string, body [][]float64, labels []string, boundingBox []BoundingBox, maskBlob string, text []string, mask []float64, videoBlob string, data []CompositeVizItem, ) *VisualizedItemData`

NewVisualizedItemData instantiates a new VisualizedItemData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualizedItemDataWithDefaults

`func NewVisualizedItemDataWithDefaults() *VisualizedItemData`

NewVisualizedItemDataWithDefaults instantiates a new VisualizedItemData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlob

`func (o *VisualizedItemData) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *VisualizedItemData) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *VisualizedItemData) SetBlob(v string)`

SetBlob sets Blob field to given value.


### GetType

`func (o *VisualizedItemData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VisualizedItemData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VisualizedItemData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.


### GetHeatmapBlob

`func (o *VisualizedItemData) GetHeatmapBlob() string`

GetHeatmapBlob returns the HeatmapBlob field if non-nil, zero value otherwise.

### GetHeatmapBlobOk

`func (o *VisualizedItemData) GetHeatmapBlobOk() (*string, bool)`

GetHeatmapBlobOk returns a tuple with the HeatmapBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmapBlob

`func (o *VisualizedItemData) SetHeatmapBlob(v string)`

SetHeatmapBlob sets HeatmapBlob field to given value.


### GetBody

`func (o *VisualizedItemData) GetBody() [][]float64`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *VisualizedItemData) GetBodyOk() (*[][]float64, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *VisualizedItemData) SetBody(v [][]float64)`

SetBody sets Body field to given value.


### GetHeatmap

`func (o *VisualizedItemData) GetHeatmap() Heatmap`

GetHeatmap returns the Heatmap field if non-nil, zero value otherwise.

### GetHeatmapOk

`func (o *VisualizedItemData) GetHeatmapOk() (*Heatmap, bool)`

GetHeatmapOk returns a tuple with the Heatmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmap

`func (o *VisualizedItemData) SetHeatmap(v Heatmap)`

SetHeatmap sets Heatmap field to given value.

### HasHeatmap

`func (o *VisualizedItemData) HasHeatmap() bool`

HasHeatmap returns a boolean if a field has been set.

### GetLabels

`func (o *VisualizedItemData) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *VisualizedItemData) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *VisualizedItemData) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetGt

`func (o *VisualizedItemData) GetGt() []float64`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *VisualizedItemData) GetGtOk() (*[]float64, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *VisualizedItemData) SetGt(v []float64)`

SetGt sets Gt field to given value.

### HasGt

`func (o *VisualizedItemData) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetXLabel

`func (o *VisualizedItemData) GetXLabel() string`

GetXLabel returns the XLabel field if non-nil, zero value otherwise.

### GetXLabelOk

`func (o *VisualizedItemData) GetXLabelOk() (*string, bool)`

GetXLabelOk returns a tuple with the XLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLabel

`func (o *VisualizedItemData) SetXLabel(v string)`

SetXLabel sets XLabel field to given value.

### HasXLabel

`func (o *VisualizedItemData) HasXLabel() bool`

HasXLabel returns a boolean if a field has been set.

### GetYLabel

`func (o *VisualizedItemData) GetYLabel() string`

GetYLabel returns the YLabel field if non-nil, zero value otherwise.

### GetYLabelOk

`func (o *VisualizedItemData) GetYLabelOk() (*string, bool)`

GetYLabelOk returns a tuple with the YLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYLabel

`func (o *VisualizedItemData) SetYLabel(v string)`

SetYLabel sets YLabel field to given value.

### HasYLabel

`func (o *VisualizedItemData) HasYLabel() bool`

HasYLabel returns a boolean if a field has been set.

### GetXRange

`func (o *VisualizedItemData) GetXRange() []float64`

GetXRange returns the XRange field if non-nil, zero value otherwise.

### GetXRangeOk

`func (o *VisualizedItemData) GetXRangeOk() (*[]float64, bool)`

GetXRangeOk returns a tuple with the XRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXRange

`func (o *VisualizedItemData) SetXRange(v []float64)`

SetXRange sets XRange field to given value.

### HasXRange

`func (o *VisualizedItemData) HasXRange() bool`

HasXRange returns a boolean if a field has been set.

### GetBoundingBox

`func (o *VisualizedItemData) GetBoundingBox() []BoundingBox`

GetBoundingBox returns the BoundingBox field if non-nil, zero value otherwise.

### GetBoundingBoxOk

`func (o *VisualizedItemData) GetBoundingBoxOk() (*[]BoundingBox, bool)`

GetBoundingBoxOk returns a tuple with the BoundingBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingBox

`func (o *VisualizedItemData) SetBoundingBox(v []BoundingBox)`

SetBoundingBox sets BoundingBox field to given value.


### GetMaskBlob

`func (o *VisualizedItemData) GetMaskBlob() string`

GetMaskBlob returns the MaskBlob field if non-nil, zero value otherwise.

### GetMaskBlobOk

`func (o *VisualizedItemData) GetMaskBlobOk() (*string, bool)`

GetMaskBlobOk returns a tuple with the MaskBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBlob

`func (o *VisualizedItemData) SetMaskBlob(v string)`

SetMaskBlob sets MaskBlob field to given value.


### GetText

`func (o *VisualizedItemData) GetText() []string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *VisualizedItemData) GetTextOk() (*[]string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *VisualizedItemData) SetText(v []string)`

SetText sets Text field to given value.


### GetMask

`func (o *VisualizedItemData) GetMask() []float64`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *VisualizedItemData) GetMaskOk() (*[]float64, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *VisualizedItemData) SetMask(v []float64)`

SetMask sets Mask field to given value.


### GetVideoBlob

`func (o *VisualizedItemData) GetVideoBlob() string`

GetVideoBlob returns the VideoBlob field if non-nil, zero value otherwise.

### GetVideoBlobOk

`func (o *VisualizedItemData) GetVideoBlobOk() (*string, bool)`

GetVideoBlobOk returns a tuple with the VideoBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoBlob

`func (o *VisualizedItemData) SetVideoBlob(v string)`

SetVideoBlob sets VideoBlob field to given value.


### GetData

`func (o *VisualizedItemData) GetData() []CompositeVizItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VisualizedItemData) GetDataOk() (*[]CompositeVizItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VisualizedItemData) SetData(v []CompositeVizItem)`

SetData sets Data field to given value.


### GetAlgo

`func (o *VisualizedItemData) GetAlgo() SampleAnalysisAlgo`

GetAlgo returns the Algo field if non-nil, zero value otherwise.

### GetAlgoOk

`func (o *VisualizedItemData) GetAlgoOk() (*SampleAnalysisAlgo, bool)`

GetAlgoOk returns a tuple with the Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgo

`func (o *VisualizedItemData) SetAlgo(v SampleAnalysisAlgo)`

SetAlgo sets Algo field to given value.

### HasAlgo

`func (o *VisualizedItemData) HasAlgo() bool`

HasAlgo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


