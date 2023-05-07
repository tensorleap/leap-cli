# MetricMessageParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**MessageCode** | Pointer to **string** |  | [optional] 
**MetricName** | **string** |  | 
**LossNodeId** | Pointer to **string** |  | [optional] 
**LossName** | Pointer to **string** |  | [optional] 
**Module** | [**Module**](Module.md) |  | 
**NodeId** | Pointer to **string** |  | [optional] 
**LineNumber** | Pointer to **float64** |  | [optional] 

## Methods

### NewMetricMessageParams

`func NewMetricMessageParams(message string, metricName string, module Module, ) *MetricMessageParams`

NewMetricMessageParams instantiates a new MetricMessageParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricMessageParamsWithDefaults

`func NewMetricMessageParamsWithDefaults() *MetricMessageParams`

NewMetricMessageParamsWithDefaults instantiates a new MetricMessageParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MetricMessageParams) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MetricMessageParams) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MetricMessageParams) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessageCode

`func (o *MetricMessageParams) GetMessageCode() string`

GetMessageCode returns the MessageCode field if non-nil, zero value otherwise.

### GetMessageCodeOk

`func (o *MetricMessageParams) GetMessageCodeOk() (*string, bool)`

GetMessageCodeOk returns a tuple with the MessageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCode

`func (o *MetricMessageParams) SetMessageCode(v string)`

SetMessageCode sets MessageCode field to given value.

### HasMessageCode

`func (o *MetricMessageParams) HasMessageCode() bool`

HasMessageCode returns a boolean if a field has been set.

### GetMetricName

`func (o *MetricMessageParams) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricMessageParams) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricMessageParams) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetLossNodeId

`func (o *MetricMessageParams) GetLossNodeId() string`

GetLossNodeId returns the LossNodeId field if non-nil, zero value otherwise.

### GetLossNodeIdOk

`func (o *MetricMessageParams) GetLossNodeIdOk() (*string, bool)`

GetLossNodeIdOk returns a tuple with the LossNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossNodeId

`func (o *MetricMessageParams) SetLossNodeId(v string)`

SetLossNodeId sets LossNodeId field to given value.

### HasLossNodeId

`func (o *MetricMessageParams) HasLossNodeId() bool`

HasLossNodeId returns a boolean if a field has been set.

### GetLossName

`func (o *MetricMessageParams) GetLossName() string`

GetLossName returns the LossName field if non-nil, zero value otherwise.

### GetLossNameOk

`func (o *MetricMessageParams) GetLossNameOk() (*string, bool)`

GetLossNameOk returns a tuple with the LossName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossName

`func (o *MetricMessageParams) SetLossName(v string)`

SetLossName sets LossName field to given value.

### HasLossName

`func (o *MetricMessageParams) HasLossName() bool`

HasLossName returns a boolean if a field has been set.

### GetModule

`func (o *MetricMessageParams) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *MetricMessageParams) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *MetricMessageParams) SetModule(v Module)`

SetModule sets Module field to given value.


### GetNodeId

`func (o *MetricMessageParams) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *MetricMessageParams) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *MetricMessageParams) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *MetricMessageParams) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetLineNumber

`func (o *MetricMessageParams) GetLineNumber() float64`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *MetricMessageParams) GetLineNumberOk() (*float64, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *MetricMessageParams) SetLineNumber(v float64)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *MetricMessageParams) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


