# VisualizerInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**LeapDataType**](LeapDataType.md) |  | 
**ArgNames** | **[]string** |  | 

## Methods

### NewVisualizerInstance

`func NewVisualizerInstance(name string, type_ LeapDataType, argNames []string, ) *VisualizerInstance`

NewVisualizerInstance instantiates a new VisualizerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualizerInstanceWithDefaults

`func NewVisualizerInstanceWithDefaults() *VisualizerInstance`

NewVisualizerInstanceWithDefaults instantiates a new VisualizerInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VisualizerInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VisualizerInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VisualizerInstance) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *VisualizerInstance) GetType() LeapDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VisualizerInstance) GetTypeOk() (*LeapDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VisualizerInstance) SetType(v LeapDataType)`

SetType sets Type field to given value.


### GetArgNames

`func (o *VisualizerInstance) GetArgNames() []string`

GetArgNames returns the ArgNames field if non-nil, zero value otherwise.

### GetArgNamesOk

`func (o *VisualizerInstance) GetArgNamesOk() (*[]string, bool)`

GetArgNamesOk returns a tuple with the ArgNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgNames

`func (o *VisualizerInstance) SetArgNames(v []string)`

SetArgNames sets ArgNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


