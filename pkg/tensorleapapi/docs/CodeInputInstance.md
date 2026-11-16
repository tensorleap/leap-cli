# CodeInputInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Shape** | **[]float64** |  | 
**ChannelDim** | Pointer to **float64** |  | [optional] 
**Foundation** | Pointer to [**FoundationInputInstance**](FoundationInputInstance.md) |  | [optional] 
**FrameDim** | Pointer to **float64** |  | [optional] 

## Methods

### NewCodeInputInstance

`func NewCodeInputInstance(name string, shape []float64, ) *CodeInputInstance`

NewCodeInputInstance instantiates a new CodeInputInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeInputInstanceWithDefaults

`func NewCodeInputInstanceWithDefaults() *CodeInputInstance`

NewCodeInputInstanceWithDefaults instantiates a new CodeInputInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CodeInputInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodeInputInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodeInputInstance) SetName(v string)`

SetName sets Name field to given value.


### GetShape

`func (o *CodeInputInstance) GetShape() []float64`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *CodeInputInstance) GetShapeOk() (*[]float64, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *CodeInputInstance) SetShape(v []float64)`

SetShape sets Shape field to given value.


### GetChannelDim

`func (o *CodeInputInstance) GetChannelDim() float64`

GetChannelDim returns the ChannelDim field if non-nil, zero value otherwise.

### GetChannelDimOk

`func (o *CodeInputInstance) GetChannelDimOk() (*float64, bool)`

GetChannelDimOk returns a tuple with the ChannelDim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelDim

`func (o *CodeInputInstance) SetChannelDim(v float64)`

SetChannelDim sets ChannelDim field to given value.

### HasChannelDim

`func (o *CodeInputInstance) HasChannelDim() bool`

HasChannelDim returns a boolean if a field has been set.

### GetFoundation

`func (o *CodeInputInstance) GetFoundation() FoundationInputInstance`

GetFoundation returns the Foundation field if non-nil, zero value otherwise.

### GetFoundationOk

`func (o *CodeInputInstance) GetFoundationOk() (*FoundationInputInstance, bool)`

GetFoundationOk returns a tuple with the Foundation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundation

`func (o *CodeInputInstance) SetFoundation(v FoundationInputInstance)`

SetFoundation sets Foundation field to given value.

### HasFoundation

`func (o *CodeInputInstance) HasFoundation() bool`

HasFoundation returns a boolean if a field has been set.

### GetFrameDim

`func (o *CodeInputInstance) GetFrameDim() float64`

GetFrameDim returns the FrameDim field if non-nil, zero value otherwise.

### GetFrameDimOk

`func (o *CodeInputInstance) GetFrameDimOk() (*float64, bool)`

GetFrameDimOk returns a tuple with the FrameDim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameDim

`func (o *CodeInputInstance) SetFrameDim(v float64)`

SetFrameDim sets FrameDim field to given value.

### HasFrameDim

`func (o *CodeInputInstance) HasFrameDim() bool`

HasFrameDim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


