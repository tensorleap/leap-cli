# FoundationInputInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RgbRange** | **[][]float64** |  | 
**RgbChannels** | **[]float64** |  | 
**Dtype** | **string** |  | 
**WiredToModel** | **bool** |  | 
**MasksEncoder** | Pointer to **string** |  | [optional] 

## Methods

### NewFoundationInputInstance

`func NewFoundationInputInstance(rgbRange [][]float64, rgbChannels []float64, dtype string, wiredToModel bool, ) *FoundationInputInstance`

NewFoundationInputInstance instantiates a new FoundationInputInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoundationInputInstanceWithDefaults

`func NewFoundationInputInstanceWithDefaults() *FoundationInputInstance`

NewFoundationInputInstanceWithDefaults instantiates a new FoundationInputInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRgbRange

`func (o *FoundationInputInstance) GetRgbRange() [][]float64`

GetRgbRange returns the RgbRange field if non-nil, zero value otherwise.

### GetRgbRangeOk

`func (o *FoundationInputInstance) GetRgbRangeOk() (*[][]float64, bool)`

GetRgbRangeOk returns a tuple with the RgbRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRgbRange

`func (o *FoundationInputInstance) SetRgbRange(v [][]float64)`

SetRgbRange sets RgbRange field to given value.


### GetRgbChannels

`func (o *FoundationInputInstance) GetRgbChannels() []float64`

GetRgbChannels returns the RgbChannels field if non-nil, zero value otherwise.

### GetRgbChannelsOk

`func (o *FoundationInputInstance) GetRgbChannelsOk() (*[]float64, bool)`

GetRgbChannelsOk returns a tuple with the RgbChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRgbChannels

`func (o *FoundationInputInstance) SetRgbChannels(v []float64)`

SetRgbChannels sets RgbChannels field to given value.


### GetDtype

`func (o *FoundationInputInstance) GetDtype() string`

GetDtype returns the Dtype field if non-nil, zero value otherwise.

### GetDtypeOk

`func (o *FoundationInputInstance) GetDtypeOk() (*string, bool)`

GetDtypeOk returns a tuple with the Dtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtype

`func (o *FoundationInputInstance) SetDtype(v string)`

SetDtype sets Dtype field to given value.


### GetWiredToModel

`func (o *FoundationInputInstance) GetWiredToModel() bool`

GetWiredToModel returns the WiredToModel field if non-nil, zero value otherwise.

### GetWiredToModelOk

`func (o *FoundationInputInstance) GetWiredToModelOk() (*bool, bool)`

GetWiredToModelOk returns a tuple with the WiredToModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredToModel

`func (o *FoundationInputInstance) SetWiredToModel(v bool)`

SetWiredToModel sets WiredToModel field to given value.


### GetMasksEncoder

`func (o *FoundationInputInstance) GetMasksEncoder() string`

GetMasksEncoder returns the MasksEncoder field if non-nil, zero value otherwise.

### GetMasksEncoderOk

`func (o *FoundationInputInstance) GetMasksEncoderOk() (*string, bool)`

GetMasksEncoderOk returns a tuple with the MasksEncoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasksEncoder

`func (o *FoundationInputInstance) SetMasksEncoder(v string)`

SetMasksEncoder sets MasksEncoder field to given value.

### HasMasksEncoder

`func (o *FoundationInputInstance) HasMasksEncoder() bool`

HasMasksEncoder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


