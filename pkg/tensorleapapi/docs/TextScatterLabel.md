# TextScatterLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**TextVisualizations** | [**[]TextViz**](TextViz.md) |  | 
**InputName** | Pointer to **string** |  | [optional] 

## Methods

### NewTextScatterLabel

`func NewTextScatterLabel(type_ string, textVisualizations []TextViz, ) *TextScatterLabel`

NewTextScatterLabel instantiates a new TextScatterLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextScatterLabelWithDefaults

`func NewTextScatterLabelWithDefaults() *TextScatterLabel`

NewTextScatterLabelWithDefaults instantiates a new TextScatterLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TextScatterLabel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextScatterLabel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextScatterLabel) SetType(v string)`

SetType sets Type field to given value.


### GetTextVisualizations

`func (o *TextScatterLabel) GetTextVisualizations() []TextViz`

GetTextVisualizations returns the TextVisualizations field if non-nil, zero value otherwise.

### GetTextVisualizationsOk

`func (o *TextScatterLabel) GetTextVisualizationsOk() (*[]TextViz, bool)`

GetTextVisualizationsOk returns a tuple with the TextVisualizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextVisualizations

`func (o *TextScatterLabel) SetTextVisualizations(v []TextViz)`

SetTextVisualizations sets TextVisualizations field to given value.


### GetInputName

`func (o *TextScatterLabel) GetInputName() string`

GetInputName returns the InputName field if non-nil, zero value otherwise.

### GetInputNameOk

`func (o *TextScatterLabel) GetInputNameOk() (*string, bool)`

GetInputNameOk returns a tuple with the InputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputName

`func (o *TextScatterLabel) SetInputName(v string)`

SetInputName sets InputName field to given value.

### HasInputName

`func (o *TextScatterLabel) HasInputName() bool`

HasInputName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


