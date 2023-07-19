# VisualizedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisualizerName** | **string** |  | 
**Data** | [**VisData**](VisData.md) |  | 
**EncoderNames** | **[]string** |  | 
**ConnectionName** | **string** |  | 
**NodeId** | Pointer to **string** |  | [optional] 
**HasError** | Pointer to **bool** |  | [optional] 

## Methods

### NewVisualizedItem

`func NewVisualizedItem(visualizerName string, data VisData, encoderNames []string, connectionName string, ) *VisualizedItem`

NewVisualizedItem instantiates a new VisualizedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualizedItemWithDefaults

`func NewVisualizedItemWithDefaults() *VisualizedItem`

NewVisualizedItemWithDefaults instantiates a new VisualizedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisualizerName

`func (o *VisualizedItem) GetVisualizerName() string`

GetVisualizerName returns the VisualizerName field if non-nil, zero value otherwise.

### GetVisualizerNameOk

`func (o *VisualizedItem) GetVisualizerNameOk() (*string, bool)`

GetVisualizerNameOk returns a tuple with the VisualizerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizerName

`func (o *VisualizedItem) SetVisualizerName(v string)`

SetVisualizerName sets VisualizerName field to given value.


### GetData

`func (o *VisualizedItem) GetData() VisData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VisualizedItem) GetDataOk() (*VisData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VisualizedItem) SetData(v VisData)`

SetData sets Data field to given value.


### GetEncoderNames

`func (o *VisualizedItem) GetEncoderNames() []string`

GetEncoderNames returns the EncoderNames field if non-nil, zero value otherwise.

### GetEncoderNamesOk

`func (o *VisualizedItem) GetEncoderNamesOk() (*[]string, bool)`

GetEncoderNamesOk returns a tuple with the EncoderNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderNames

`func (o *VisualizedItem) SetEncoderNames(v []string)`

SetEncoderNames sets EncoderNames field to given value.


### GetConnectionName

`func (o *VisualizedItem) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *VisualizedItem) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *VisualizedItem) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetNodeId

`func (o *VisualizedItem) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *VisualizedItem) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *VisualizedItem) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *VisualizedItem) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetHasError

`func (o *VisualizedItem) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *VisualizedItem) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *VisualizedItem) SetHasError(v bool)`

SetHasError sets HasError field to given value.

### HasHasError

`func (o *VisualizedItem) HasHasError() bool`

HasHasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


