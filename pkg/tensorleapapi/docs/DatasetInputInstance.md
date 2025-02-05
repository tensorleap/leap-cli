# DatasetInputInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Shape** | **[]float64** |  | 
**ChannelDim** | Pointer to **float64** |  | [optional] 

## Methods

### NewDatasetInputInstance

`func NewDatasetInputInstance(name string, shape []float64, ) *DatasetInputInstance`

NewDatasetInputInstance instantiates a new DatasetInputInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetInputInstanceWithDefaults

`func NewDatasetInputInstanceWithDefaults() *DatasetInputInstance`

NewDatasetInputInstanceWithDefaults instantiates a new DatasetInputInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatasetInputInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasetInputInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasetInputInstance) SetName(v string)`

SetName sets Name field to given value.


### GetShape

`func (o *DatasetInputInstance) GetShape() []float64`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *DatasetInputInstance) GetShapeOk() (*[]float64, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *DatasetInputInstance) SetShape(v []float64)`

SetShape sets Shape field to given value.


### GetChannelDim

`func (o *DatasetInputInstance) GetChannelDim() float64`

GetChannelDim returns the ChannelDim field if non-nil, zero value otherwise.

### GetChannelDimOk

`func (o *DatasetInputInstance) GetChannelDimOk() (*float64, bool)`

GetChannelDimOk returns a tuple with the ChannelDim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelDim

`func (o *DatasetInputInstance) SetChannelDim(v float64)`

SetChannelDim sets ChannelDim field to given value.

### HasChannelDim

`func (o *DatasetInputInstance) HasChannelDim() bool`

HasChannelDim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


