# LossMessageParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**MessageCode** | Pointer to **string** |  | [optional] 
**LossName** | **string** |  | 
**NodeId** | **string** |  | 
**Module** | [**Module**](Module.md) |  | 

## Methods

### NewLossMessageParams

`func NewLossMessageParams(message string, lossName string, nodeId string, module Module, ) *LossMessageParams`

NewLossMessageParams instantiates a new LossMessageParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLossMessageParamsWithDefaults

`func NewLossMessageParamsWithDefaults() *LossMessageParams`

NewLossMessageParamsWithDefaults instantiates a new LossMessageParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *LossMessageParams) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LossMessageParams) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LossMessageParams) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessageCode

`func (o *LossMessageParams) GetMessageCode() string`

GetMessageCode returns the MessageCode field if non-nil, zero value otherwise.

### GetMessageCodeOk

`func (o *LossMessageParams) GetMessageCodeOk() (*string, bool)`

GetMessageCodeOk returns a tuple with the MessageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCode

`func (o *LossMessageParams) SetMessageCode(v string)`

SetMessageCode sets MessageCode field to given value.

### HasMessageCode

`func (o *LossMessageParams) HasMessageCode() bool`

HasMessageCode returns a boolean if a field has been set.

### GetLossName

`func (o *LossMessageParams) GetLossName() string`

GetLossName returns the LossName field if non-nil, zero value otherwise.

### GetLossNameOk

`func (o *LossMessageParams) GetLossNameOk() (*string, bool)`

GetLossNameOk returns a tuple with the LossName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossName

`func (o *LossMessageParams) SetLossName(v string)`

SetLossName sets LossName field to given value.


### GetNodeId

`func (o *LossMessageParams) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *LossMessageParams) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *LossMessageParams) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetModule

`func (o *LossMessageParams) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *LossMessageParams) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *LossMessageParams) SetModule(v Module)`

SetModule sets Module field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


