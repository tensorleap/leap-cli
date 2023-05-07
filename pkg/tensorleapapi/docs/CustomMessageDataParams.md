# CustomMessageDataParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**MessageCode** | Pointer to **string** |  | [optional] 
**Module** | **string** |  | 
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
**JobStatus** | [**JobStatus**](JobStatus.md) |  | 

## Methods

### NewCustomMessageDataParams

`func NewCustomMessageDataParams(message string, module string, layerName string, layerType NodeType, visualizerName string, nodeId string, line float64, funcName string, lossName string, metricName string, jobStatus JobStatus, ) *CustomMessageDataParams`

NewCustomMessageDataParams instantiates a new CustomMessageDataParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomMessageDataParamsWithDefaults

`func NewCustomMessageDataParamsWithDefaults() *CustomMessageDataParams`

NewCustomMessageDataParamsWithDefaults instantiates a new CustomMessageDataParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CustomMessageDataParams) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CustomMessageDataParams) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CustomMessageDataParams) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessageCode

`func (o *CustomMessageDataParams) GetMessageCode() string`

GetMessageCode returns the MessageCode field if non-nil, zero value otherwise.

### GetMessageCodeOk

`func (o *CustomMessageDataParams) GetMessageCodeOk() (*string, bool)`

GetMessageCodeOk returns a tuple with the MessageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCode

`func (o *CustomMessageDataParams) SetMessageCode(v string)`

SetMessageCode sets MessageCode field to given value.

### HasMessageCode

`func (o *CustomMessageDataParams) HasMessageCode() bool`

HasMessageCode returns a boolean if a field has been set.

### GetModule

`func (o *CustomMessageDataParams) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CustomMessageDataParams) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CustomMessageDataParams) SetModule(v string)`

SetModule sets Module field to given value.


### GetLayerName

`func (o *CustomMessageDataParams) GetLayerName() string`

GetLayerName returns the LayerName field if non-nil, zero value otherwise.

### GetLayerNameOk

`func (o *CustomMessageDataParams) GetLayerNameOk() (*string, bool)`

GetLayerNameOk returns a tuple with the LayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerName

`func (o *CustomMessageDataParams) SetLayerName(v string)`

SetLayerName sets LayerName field to given value.


### GetLayerType

`func (o *CustomMessageDataParams) GetLayerType() NodeType`

GetLayerType returns the LayerType field if non-nil, zero value otherwise.

### GetLayerTypeOk

`func (o *CustomMessageDataParams) GetLayerTypeOk() (*NodeType, bool)`

GetLayerTypeOk returns a tuple with the LayerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerType

`func (o *CustomMessageDataParams) SetLayerType(v NodeType)`

SetLayerType sets LayerType field to given value.


### GetVisualizerName

`func (o *CustomMessageDataParams) GetVisualizerName() string`

GetVisualizerName returns the VisualizerName field if non-nil, zero value otherwise.

### GetVisualizerNameOk

`func (o *CustomMessageDataParams) GetVisualizerNameOk() (*string, bool)`

GetVisualizerNameOk returns a tuple with the VisualizerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizerName

`func (o *CustomMessageDataParams) SetVisualizerName(v string)`

SetVisualizerName sets VisualizerName field to given value.


### GetNodeId

`func (o *CustomMessageDataParams) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *CustomMessageDataParams) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *CustomMessageDataParams) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetLineNumber

`func (o *CustomMessageDataParams) GetLineNumber() float64`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *CustomMessageDataParams) GetLineNumberOk() (*float64, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *CustomMessageDataParams) SetLineNumber(v float64)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *CustomMessageDataParams) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetLine

`func (o *CustomMessageDataParams) GetLine() float64`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *CustomMessageDataParams) GetLineOk() (*float64, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *CustomMessageDataParams) SetLine(v float64)`

SetLine sets Line field to given value.


### GetFuncName

`func (o *CustomMessageDataParams) GetFuncName() string`

GetFuncName returns the FuncName field if non-nil, zero value otherwise.

### GetFuncNameOk

`func (o *CustomMessageDataParams) GetFuncNameOk() (*string, bool)`

GetFuncNameOk returns a tuple with the FuncName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncName

`func (o *CustomMessageDataParams) SetFuncName(v string)`

SetFuncName sets FuncName field to given value.


### GetLossName

`func (o *CustomMessageDataParams) GetLossName() string`

GetLossName returns the LossName field if non-nil, zero value otherwise.

### GetLossNameOk

`func (o *CustomMessageDataParams) GetLossNameOk() (*string, bool)`

GetLossNameOk returns a tuple with the LossName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossName

`func (o *CustomMessageDataParams) SetLossName(v string)`

SetLossName sets LossName field to given value.


### GetMetricName

`func (o *CustomMessageDataParams) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *CustomMessageDataParams) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *CustomMessageDataParams) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetLossNodeId

`func (o *CustomMessageDataParams) GetLossNodeId() string`

GetLossNodeId returns the LossNodeId field if non-nil, zero value otherwise.

### GetLossNodeIdOk

`func (o *CustomMessageDataParams) GetLossNodeIdOk() (*string, bool)`

GetLossNodeIdOk returns a tuple with the LossNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossNodeId

`func (o *CustomMessageDataParams) SetLossNodeId(v string)`

SetLossNodeId sets LossNodeId field to given value.

### HasLossNodeId

`func (o *CustomMessageDataParams) HasLossNodeId() bool`

HasLossNodeId returns a boolean if a field has been set.

### GetJobStatus

`func (o *CustomMessageDataParams) GetJobStatus() JobStatus`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *CustomMessageDataParams) GetJobStatusOk() (*JobStatus, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *CustomMessageDataParams) SetJobStatus(v JobStatus)`

SetJobStatus sets JobStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


