# NodeMessageParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**MessageCode** | Pointer to **string** |  | [optional] 
**LayerName** | **string** |  | 
**LayerType** | [**NodeType**](NodeType.md) |  | 
**Module** | [**Module**](Module.md) |  | 

## Methods

### NewNodeMessageParams

`func NewNodeMessageParams(message string, layerName string, layerType NodeType, module Module, ) *NodeMessageParams`

NewNodeMessageParams instantiates a new NodeMessageParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeMessageParamsWithDefaults

`func NewNodeMessageParamsWithDefaults() *NodeMessageParams`

NewNodeMessageParamsWithDefaults instantiates a new NodeMessageParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *NodeMessageParams) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NodeMessageParams) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NodeMessageParams) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessageCode

`func (o *NodeMessageParams) GetMessageCode() string`

GetMessageCode returns the MessageCode field if non-nil, zero value otherwise.

### GetMessageCodeOk

`func (o *NodeMessageParams) GetMessageCodeOk() (*string, bool)`

GetMessageCodeOk returns a tuple with the MessageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCode

`func (o *NodeMessageParams) SetMessageCode(v string)`

SetMessageCode sets MessageCode field to given value.

### HasMessageCode

`func (o *NodeMessageParams) HasMessageCode() bool`

HasMessageCode returns a boolean if a field has been set.

### GetLayerName

`func (o *NodeMessageParams) GetLayerName() string`

GetLayerName returns the LayerName field if non-nil, zero value otherwise.

### GetLayerNameOk

`func (o *NodeMessageParams) GetLayerNameOk() (*string, bool)`

GetLayerNameOk returns a tuple with the LayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerName

`func (o *NodeMessageParams) SetLayerName(v string)`

SetLayerName sets LayerName field to given value.


### GetLayerType

`func (o *NodeMessageParams) GetLayerType() NodeType`

GetLayerType returns the LayerType field if non-nil, zero value otherwise.

### GetLayerTypeOk

`func (o *NodeMessageParams) GetLayerTypeOk() (*NodeType, bool)`

GetLayerTypeOk returns a tuple with the LayerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerType

`func (o *NodeMessageParams) SetLayerType(v NodeType)`

SetLayerType sets LayerType field to given value.


### GetModule

`func (o *NodeMessageParams) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *NodeMessageParams) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *NodeMessageParams) SetModule(v Module)`

SetModule sets Module field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


