# TextMaskScatterLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**TextVisualizations** | [**[]TextViz**](TextViz.md) |  | 
**Mask** | **[][]float64** |  | 
**Labels** | **[]string** |  | 
**InputName** | Pointer to **string** |  | [optional] 

## Methods

### NewTextMaskScatterLabel

`func NewTextMaskScatterLabel(type_ string, textVisualizations []TextViz, mask [][]float64, labels []string, ) *TextMaskScatterLabel`

NewTextMaskScatterLabel instantiates a new TextMaskScatterLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextMaskScatterLabelWithDefaults

`func NewTextMaskScatterLabelWithDefaults() *TextMaskScatterLabel`

NewTextMaskScatterLabelWithDefaults instantiates a new TextMaskScatterLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TextMaskScatterLabel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextMaskScatterLabel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextMaskScatterLabel) SetType(v string)`

SetType sets Type field to given value.


### GetTextVisualizations

`func (o *TextMaskScatterLabel) GetTextVisualizations() []TextViz`

GetTextVisualizations returns the TextVisualizations field if non-nil, zero value otherwise.

### GetTextVisualizationsOk

`func (o *TextMaskScatterLabel) GetTextVisualizationsOk() (*[]TextViz, bool)`

GetTextVisualizationsOk returns a tuple with the TextVisualizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextVisualizations

`func (o *TextMaskScatterLabel) SetTextVisualizations(v []TextViz)`

SetTextVisualizations sets TextVisualizations field to given value.


### GetMask

`func (o *TextMaskScatterLabel) GetMask() [][]float64`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *TextMaskScatterLabel) GetMaskOk() (*[][]float64, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *TextMaskScatterLabel) SetMask(v [][]float64)`

SetMask sets Mask field to given value.


### GetLabels

`func (o *TextMaskScatterLabel) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *TextMaskScatterLabel) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *TextMaskScatterLabel) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetInputName

`func (o *TextMaskScatterLabel) GetInputName() string`

GetInputName returns the InputName field if non-nil, zero value otherwise.

### GetInputNameOk

`func (o *TextMaskScatterLabel) GetInputNameOk() (*string, bool)`

GetInputNameOk returns a tuple with the InputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputName

`func (o *TextMaskScatterLabel) SetInputName(v string)`

SetInputName sets InputName field to given value.

### HasInputName

`func (o *TextMaskScatterLabel) HasInputName() bool`

HasInputName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


