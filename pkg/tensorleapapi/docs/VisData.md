# VisData

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

## Methods

### NewVisData

`func NewVisData(blob string, type_ DataTypeEnum, heatmapBlob string, body [][]float64, labels []string, boundingBox []BoundingBox, maskBlob string, text []string, mask []float64, ) *VisData`

NewVisData instantiates a new VisData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisDataWithDefaults

`func NewVisDataWithDefaults() *VisData`

NewVisDataWithDefaults instantiates a new VisData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlob

`func (o *VisData) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *VisData) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *VisData) SetBlob(v string)`

SetBlob sets Blob field to given value.


### GetType

`func (o *VisData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VisData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VisData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.


### GetHeatmapBlob

`func (o *VisData) GetHeatmapBlob() string`

GetHeatmapBlob returns the HeatmapBlob field if non-nil, zero value otherwise.

### GetHeatmapBlobOk

`func (o *VisData) GetHeatmapBlobOk() (*string, bool)`

GetHeatmapBlobOk returns a tuple with the HeatmapBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmapBlob

`func (o *VisData) SetHeatmapBlob(v string)`

SetHeatmapBlob sets HeatmapBlob field to given value.


### GetBody

`func (o *VisData) GetBody() [][]float64`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *VisData) GetBodyOk() (*[][]float64, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *VisData) SetBody(v [][]float64)`

SetBody sets Body field to given value.


### GetHeatmap

`func (o *VisData) GetHeatmap() Heatmap`

GetHeatmap returns the Heatmap field if non-nil, zero value otherwise.

### GetHeatmapOk

`func (o *VisData) GetHeatmapOk() (*Heatmap, bool)`

GetHeatmapOk returns a tuple with the Heatmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmap

`func (o *VisData) SetHeatmap(v Heatmap)`

SetHeatmap sets Heatmap field to given value.

### HasHeatmap

`func (o *VisData) HasHeatmap() bool`

HasHeatmap returns a boolean if a field has been set.

### GetLabels

`func (o *VisData) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *VisData) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *VisData) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetGt

`func (o *VisData) GetGt() []float64`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *VisData) GetGtOk() (*[]float64, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *VisData) SetGt(v []float64)`

SetGt sets Gt field to given value.

### HasGt

`func (o *VisData) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetXLabel

`func (o *VisData) GetXLabel() string`

GetXLabel returns the XLabel field if non-nil, zero value otherwise.

### GetXLabelOk

`func (o *VisData) GetXLabelOk() (*string, bool)`

GetXLabelOk returns a tuple with the XLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXLabel

`func (o *VisData) SetXLabel(v string)`

SetXLabel sets XLabel field to given value.

### HasXLabel

`func (o *VisData) HasXLabel() bool`

HasXLabel returns a boolean if a field has been set.

### GetYLabel

`func (o *VisData) GetYLabel() string`

GetYLabel returns the YLabel field if non-nil, zero value otherwise.

### GetYLabelOk

`func (o *VisData) GetYLabelOk() (*string, bool)`

GetYLabelOk returns a tuple with the YLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYLabel

`func (o *VisData) SetYLabel(v string)`

SetYLabel sets YLabel field to given value.

### HasYLabel

`func (o *VisData) HasYLabel() bool`

HasYLabel returns a boolean if a field has been set.

### GetXRange

`func (o *VisData) GetXRange() []float64`

GetXRange returns the XRange field if non-nil, zero value otherwise.

### GetXRangeOk

`func (o *VisData) GetXRangeOk() (*[]float64, bool)`

GetXRangeOk returns a tuple with the XRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXRange

`func (o *VisData) SetXRange(v []float64)`

SetXRange sets XRange field to given value.

### HasXRange

`func (o *VisData) HasXRange() bool`

HasXRange returns a boolean if a field has been set.

### GetBoundingBox

`func (o *VisData) GetBoundingBox() []BoundingBox`

GetBoundingBox returns the BoundingBox field if non-nil, zero value otherwise.

### GetBoundingBoxOk

`func (o *VisData) GetBoundingBoxOk() (*[]BoundingBox, bool)`

GetBoundingBoxOk returns a tuple with the BoundingBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingBox

`func (o *VisData) SetBoundingBox(v []BoundingBox)`

SetBoundingBox sets BoundingBox field to given value.


### GetMaskBlob

`func (o *VisData) GetMaskBlob() string`

GetMaskBlob returns the MaskBlob field if non-nil, zero value otherwise.

### GetMaskBlobOk

`func (o *VisData) GetMaskBlobOk() (*string, bool)`

GetMaskBlobOk returns a tuple with the MaskBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBlob

`func (o *VisData) SetMaskBlob(v string)`

SetMaskBlob sets MaskBlob field to given value.


### GetText

`func (o *VisData) GetText() []string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *VisData) GetTextOk() (*[]string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *VisData) SetText(v []string)`

SetText sets Text field to given value.


### GetMask

`func (o *VisData) GetMask() []float64`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *VisData) GetMaskOk() (*[]float64, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *VisData) SetMask(v []float64)`

SetMask sets Mask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


