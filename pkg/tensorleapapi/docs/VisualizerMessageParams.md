# VisualizerMessageParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**MessageCode** | Pointer to **string** |  | [optional] 
**VisualizerName** | **string** |  | 
**NodeId** | **string** |  | 
**LineNumber** | Pointer to **float64** |  | [optional] 
**Module** | [**Module**](Module.md) |  | 

## Methods

### NewVisualizerMessageParams

`func NewVisualizerMessageParams(message string, visualizerName string, nodeId string, module Module, ) *VisualizerMessageParams`

NewVisualizerMessageParams instantiates a new VisualizerMessageParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualizerMessageParamsWithDefaults

`func NewVisualizerMessageParamsWithDefaults() *VisualizerMessageParams`

NewVisualizerMessageParamsWithDefaults instantiates a new VisualizerMessageParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *VisualizerMessageParams) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VisualizerMessageParams) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VisualizerMessageParams) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessageCode

`func (o *VisualizerMessageParams) GetMessageCode() string`

GetMessageCode returns the MessageCode field if non-nil, zero value otherwise.

### GetMessageCodeOk

`func (o *VisualizerMessageParams) GetMessageCodeOk() (*string, bool)`

GetMessageCodeOk returns a tuple with the MessageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCode

`func (o *VisualizerMessageParams) SetMessageCode(v string)`

SetMessageCode sets MessageCode field to given value.

### HasMessageCode

`func (o *VisualizerMessageParams) HasMessageCode() bool`

HasMessageCode returns a boolean if a field has been set.

### GetVisualizerName

`func (o *VisualizerMessageParams) GetVisualizerName() string`

GetVisualizerName returns the VisualizerName field if non-nil, zero value otherwise.

### GetVisualizerNameOk

`func (o *VisualizerMessageParams) GetVisualizerNameOk() (*string, bool)`

GetVisualizerNameOk returns a tuple with the VisualizerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizerName

`func (o *VisualizerMessageParams) SetVisualizerName(v string)`

SetVisualizerName sets VisualizerName field to given value.


### GetNodeId

`func (o *VisualizerMessageParams) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *VisualizerMessageParams) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *VisualizerMessageParams) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetLineNumber

`func (o *VisualizerMessageParams) GetLineNumber() float64`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *VisualizerMessageParams) GetLineNumberOk() (*float64, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *VisualizerMessageParams) SetLineNumber(v float64)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *VisualizerMessageParams) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetModule

`func (o *VisualizerMessageParams) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *VisualizerMessageParams) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *VisualizerMessageParams) SetModule(v Module)`

SetModule sets Module field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


