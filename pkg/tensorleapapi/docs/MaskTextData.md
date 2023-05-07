# MaskTextData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **[]string** |  | 
**Mask** | **[]float64** |  | 
**Labels** | **[]string** |  | 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 

## Methods

### NewMaskTextData

`func NewMaskTextData(text []string, mask []float64, labels []string, type_ DataTypeEnum, ) *MaskTextData`

NewMaskTextData instantiates a new MaskTextData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaskTextDataWithDefaults

`func NewMaskTextDataWithDefaults() *MaskTextData`

NewMaskTextDataWithDefaults instantiates a new MaskTextData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *MaskTextData) GetText() []string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MaskTextData) GetTextOk() (*[]string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MaskTextData) SetText(v []string)`

SetText sets Text field to given value.


### GetMask

`func (o *MaskTextData) GetMask() []float64`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *MaskTextData) GetMaskOk() (*[]float64, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *MaskTextData) SetMask(v []float64)`

SetMask sets Mask field to given value.


### GetLabels

`func (o *MaskTextData) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MaskTextData) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MaskTextData) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetType

`func (o *MaskTextData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MaskTextData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MaskTextData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


