# ModuleMessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**MessageCode** | Pointer to **string** |  | [optional] 
**Module** | [**Module**](Module.md) |  | 
**LayerName** | **string** |  | 
**LayerType** | [**NodeType**](NodeType.md) |  | 
**VisualizerName** | **string** |  | 
**NodeId** | **string** |  | 
**LineNumber** | Pointer to **float64** |  | [optional] 
**Line** | **float64** |  | 
**FuncName** | **string** |  | 
**LossName** | **string** |  | 
**MetricName** | **string** |  | 
**LossNodeId** | Pointer to **string** |  | [optional] 

## Methods

### NewModuleMessageData

`func NewModuleMessageData(message string, module Module, layerName string, layerType NodeType, visualizerName string, nodeId string, line float64, funcName string, lossName string, metricName string, ) *ModuleMessageData`

NewModuleMessageData instantiates a new ModuleMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleMessageDataWithDefaults

`func NewModuleMessageDataWithDefaults() *ModuleMessageData`

NewModuleMessageDataWithDefaults instantiates a new ModuleMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ModuleMessageData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModuleMessageData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModuleMessageData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessageCode

`func (o *ModuleMessageData) GetMessageCode() string`

GetMessageCode returns the MessageCode field if non-nil, zero value otherwise.

### GetMessageCodeOk

`func (o *ModuleMessageData) GetMessageCodeOk() (*string, bool)`

GetMessageCodeOk returns a tuple with the MessageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCode

`func (o *ModuleMessageData) SetMessageCode(v string)`

SetMessageCode sets MessageCode field to given value.

### HasMessageCode

`func (o *ModuleMessageData) HasMessageCode() bool`

HasMessageCode returns a boolean if a field has been set.

### GetModule

`func (o *ModuleMessageData) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ModuleMessageData) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ModuleMessageData) SetModule(v Module)`

SetModule sets Module field to given value.


### GetLayerName

`func (o *ModuleMessageData) GetLayerName() string`

GetLayerName returns the LayerName field if non-nil, zero value otherwise.

### GetLayerNameOk

`func (o *ModuleMessageData) GetLayerNameOk() (*string, bool)`

GetLayerNameOk returns a tuple with the LayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerName

`func (o *ModuleMessageData) SetLayerName(v string)`

SetLayerName sets LayerName field to given value.


### GetLayerType

`func (o *ModuleMessageData) GetLayerType() NodeType`

GetLayerType returns the LayerType field if non-nil, zero value otherwise.

### GetLayerTypeOk

`func (o *ModuleMessageData) GetLayerTypeOk() (*NodeType, bool)`

GetLayerTypeOk returns a tuple with the LayerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerType

`func (o *ModuleMessageData) SetLayerType(v NodeType)`

SetLayerType sets LayerType field to given value.


### GetVisualizerName

`func (o *ModuleMessageData) GetVisualizerName() string`

GetVisualizerName returns the VisualizerName field if non-nil, zero value otherwise.

### GetVisualizerNameOk

`func (o *ModuleMessageData) GetVisualizerNameOk() (*string, bool)`

GetVisualizerNameOk returns a tuple with the VisualizerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizerName

`func (o *ModuleMessageData) SetVisualizerName(v string)`

SetVisualizerName sets VisualizerName field to given value.


### GetNodeId

`func (o *ModuleMessageData) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModuleMessageData) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModuleMessageData) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetLineNumber

`func (o *ModuleMessageData) GetLineNumber() float64`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *ModuleMessageData) GetLineNumberOk() (*float64, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *ModuleMessageData) SetLineNumber(v float64)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *ModuleMessageData) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetLine

`func (o *ModuleMessageData) GetLine() float64`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *ModuleMessageData) GetLineOk() (*float64, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *ModuleMessageData) SetLine(v float64)`

SetLine sets Line field to given value.


### GetFuncName

`func (o *ModuleMessageData) GetFuncName() string`

GetFuncName returns the FuncName field if non-nil, zero value otherwise.

### GetFuncNameOk

`func (o *ModuleMessageData) GetFuncNameOk() (*string, bool)`

GetFuncNameOk returns a tuple with the FuncName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncName

`func (o *ModuleMessageData) SetFuncName(v string)`

SetFuncName sets FuncName field to given value.


### GetLossName

`func (o *ModuleMessageData) GetLossName() string`

GetLossName returns the LossName field if non-nil, zero value otherwise.

### GetLossNameOk

`func (o *ModuleMessageData) GetLossNameOk() (*string, bool)`

GetLossNameOk returns a tuple with the LossName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossName

`func (o *ModuleMessageData) SetLossName(v string)`

SetLossName sets LossName field to given value.


### GetMetricName

`func (o *ModuleMessageData) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ModuleMessageData) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ModuleMessageData) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetLossNodeId

`func (o *ModuleMessageData) GetLossNodeId() string`

GetLossNodeId returns the LossNodeId field if non-nil, zero value otherwise.

### GetLossNodeIdOk

`func (o *ModuleMessageData) GetLossNodeIdOk() (*string, bool)`

GetLossNodeIdOk returns a tuple with the LossNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossNodeId

`func (o *ModuleMessageData) SetLossNodeId(v string)`

SetLossNodeId sets LossNodeId field to given value.

### HasLossNodeId

`func (o *ModuleMessageData) HasLossNodeId() bool`

HasLossNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


