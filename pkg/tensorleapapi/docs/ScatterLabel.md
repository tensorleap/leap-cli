# ScatterLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**TextVisualizations** | [**[]TextViz**](TextViz.md) |  | 
**InputName** | Pointer to **string** |  | [optional] 
**Data** | **[][]float64** |  | 
**Src** | **string** |  | 
**Blob** | **string** |  | 
**GraphVisualizations** | [**[]GraphViz**](GraphViz.md) |  | 
**HorizontalBarVisualizations** | [**[]HorizontalBarViz**](HorizontalBarViz.md) |  | 
**BoundingBoxes** | [**[][]BoundingBox**]([]BoundingBox.md) |  | 
**MaskSrc** | **string** |  | 
**MaskBlob** | **string** |  | 
**Labels** | **[]string** |  | 
**Mask** | **[][]float64** |  | 

## Methods

### NewScatterLabel

`func NewScatterLabel(type_ string, textVisualizations []TextViz, data [][]float64, src string, blob string, graphVisualizations []GraphViz, horizontalBarVisualizations []HorizontalBarViz, boundingBoxes [][]BoundingBox, maskSrc string, maskBlob string, labels []string, mask [][]float64, ) *ScatterLabel`

NewScatterLabel instantiates a new ScatterLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScatterLabelWithDefaults

`func NewScatterLabelWithDefaults() *ScatterLabel`

NewScatterLabelWithDefaults instantiates a new ScatterLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScatterLabel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScatterLabel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScatterLabel) SetType(v string)`

SetType sets Type field to given value.


### GetTextVisualizations

`func (o *ScatterLabel) GetTextVisualizations() []TextViz`

GetTextVisualizations returns the TextVisualizations field if non-nil, zero value otherwise.

### GetTextVisualizationsOk

`func (o *ScatterLabel) GetTextVisualizationsOk() (*[]TextViz, bool)`

GetTextVisualizationsOk returns a tuple with the TextVisualizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextVisualizations

`func (o *ScatterLabel) SetTextVisualizations(v []TextViz)`

SetTextVisualizations sets TextVisualizations field to given value.


### GetInputName

`func (o *ScatterLabel) GetInputName() string`

GetInputName returns the InputName field if non-nil, zero value otherwise.

### GetInputNameOk

`func (o *ScatterLabel) GetInputNameOk() (*string, bool)`

GetInputNameOk returns a tuple with the InputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputName

`func (o *ScatterLabel) SetInputName(v string)`

SetInputName sets InputName field to given value.

### HasInputName

`func (o *ScatterLabel) HasInputName() bool`

HasInputName returns a boolean if a field has been set.

### GetData

`func (o *ScatterLabel) GetData() [][]float64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ScatterLabel) GetDataOk() (*[][]float64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ScatterLabel) SetData(v [][]float64)`

SetData sets Data field to given value.


### GetSrc

`func (o *ScatterLabel) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *ScatterLabel) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *ScatterLabel) SetSrc(v string)`

SetSrc sets Src field to given value.


### GetBlob

`func (o *ScatterLabel) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *ScatterLabel) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *ScatterLabel) SetBlob(v string)`

SetBlob sets Blob field to given value.


### GetGraphVisualizations

`func (o *ScatterLabel) GetGraphVisualizations() []GraphViz`

GetGraphVisualizations returns the GraphVisualizations field if non-nil, zero value otherwise.

### GetGraphVisualizationsOk

`func (o *ScatterLabel) GetGraphVisualizationsOk() (*[]GraphViz, bool)`

GetGraphVisualizationsOk returns a tuple with the GraphVisualizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphVisualizations

`func (o *ScatterLabel) SetGraphVisualizations(v []GraphViz)`

SetGraphVisualizations sets GraphVisualizations field to given value.


### GetHorizontalBarVisualizations

`func (o *ScatterLabel) GetHorizontalBarVisualizations() []HorizontalBarViz`

GetHorizontalBarVisualizations returns the HorizontalBarVisualizations field if non-nil, zero value otherwise.

### GetHorizontalBarVisualizationsOk

`func (o *ScatterLabel) GetHorizontalBarVisualizationsOk() (*[]HorizontalBarViz, bool)`

GetHorizontalBarVisualizationsOk returns a tuple with the HorizontalBarVisualizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalBarVisualizations

`func (o *ScatterLabel) SetHorizontalBarVisualizations(v []HorizontalBarViz)`

SetHorizontalBarVisualizations sets HorizontalBarVisualizations field to given value.


### GetBoundingBoxes

`func (o *ScatterLabel) GetBoundingBoxes() [][]BoundingBox`

GetBoundingBoxes returns the BoundingBoxes field if non-nil, zero value otherwise.

### GetBoundingBoxesOk

`func (o *ScatterLabel) GetBoundingBoxesOk() (*[][]BoundingBox, bool)`

GetBoundingBoxesOk returns a tuple with the BoundingBoxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingBoxes

`func (o *ScatterLabel) SetBoundingBoxes(v [][]BoundingBox)`

SetBoundingBoxes sets BoundingBoxes field to given value.


### GetMaskSrc

`func (o *ScatterLabel) GetMaskSrc() string`

GetMaskSrc returns the MaskSrc field if non-nil, zero value otherwise.

### GetMaskSrcOk

`func (o *ScatterLabel) GetMaskSrcOk() (*string, bool)`

GetMaskSrcOk returns a tuple with the MaskSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskSrc

`func (o *ScatterLabel) SetMaskSrc(v string)`

SetMaskSrc sets MaskSrc field to given value.


### GetMaskBlob

`func (o *ScatterLabel) GetMaskBlob() string`

GetMaskBlob returns the MaskBlob field if non-nil, zero value otherwise.

### GetMaskBlobOk

`func (o *ScatterLabel) GetMaskBlobOk() (*string, bool)`

GetMaskBlobOk returns a tuple with the MaskBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBlob

`func (o *ScatterLabel) SetMaskBlob(v string)`

SetMaskBlob sets MaskBlob field to given value.


### GetLabels

`func (o *ScatterLabel) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ScatterLabel) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ScatterLabel) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetMask

`func (o *ScatterLabel) GetMask() [][]float64`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *ScatterLabel) GetMaskOk() (*[][]float64, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *ScatterLabel) SetMask(v [][]float64)`

SetMask sets Mask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


