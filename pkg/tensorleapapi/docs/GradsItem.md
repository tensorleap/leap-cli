# GradsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeName** | **string** |  | 
**Label** | **string** |  | 
**Grads** | [**VisData**](VisData.md) |  | 
**VisualizerName** | **string** |  | 
**EncoderNames** | **[]string** |  | 
**ConnectionName** | **string** |  | 
**ChannelImportance** | Pointer to [**HorizontalBarData**](HorizontalBarData.md) |  | [optional] 

## Methods

### NewGradsItem

`func NewGradsItem(nodeName string, label string, grads VisData, visualizerName string, encoderNames []string, connectionName string, ) *GradsItem`

NewGradsItem instantiates a new GradsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGradsItemWithDefaults

`func NewGradsItemWithDefaults() *GradsItem`

NewGradsItemWithDefaults instantiates a new GradsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeName

`func (o *GradsItem) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *GradsItem) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *GradsItem) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetLabel

`func (o *GradsItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GradsItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GradsItem) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetGrads

`func (o *GradsItem) GetGrads() VisData`

GetGrads returns the Grads field if non-nil, zero value otherwise.

### GetGradsOk

`func (o *GradsItem) GetGradsOk() (*VisData, bool)`

GetGradsOk returns a tuple with the Grads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrads

`func (o *GradsItem) SetGrads(v VisData)`

SetGrads sets Grads field to given value.


### GetVisualizerName

`func (o *GradsItem) GetVisualizerName() string`

GetVisualizerName returns the VisualizerName field if non-nil, zero value otherwise.

### GetVisualizerNameOk

`func (o *GradsItem) GetVisualizerNameOk() (*string, bool)`

GetVisualizerNameOk returns a tuple with the VisualizerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizerName

`func (o *GradsItem) SetVisualizerName(v string)`

SetVisualizerName sets VisualizerName field to given value.


### GetEncoderNames

`func (o *GradsItem) GetEncoderNames() []string`

GetEncoderNames returns the EncoderNames field if non-nil, zero value otherwise.

### GetEncoderNamesOk

`func (o *GradsItem) GetEncoderNamesOk() (*[]string, bool)`

GetEncoderNamesOk returns a tuple with the EncoderNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderNames

`func (o *GradsItem) SetEncoderNames(v []string)`

SetEncoderNames sets EncoderNames field to given value.


### GetConnectionName

`func (o *GradsItem) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *GradsItem) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *GradsItem) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetChannelImportance

`func (o *GradsItem) GetChannelImportance() HorizontalBarData`

GetChannelImportance returns the ChannelImportance field if non-nil, zero value otherwise.

### GetChannelImportanceOk

`func (o *GradsItem) GetChannelImportanceOk() (*HorizontalBarData, bool)`

GetChannelImportanceOk returns a tuple with the ChannelImportance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelImportance

`func (o *GradsItem) SetChannelImportance(v HorizontalBarData)`

SetChannelImportance sets ChannelImportance field to given value.

### HasChannelImportance

`func (o *GradsItem) HasChannelImportance() bool`

HasChannelImportance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


