# CodeMessageParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**MessageCode** | Pointer to **string** |  | [optional] 
**Line** | **float64** |  | 
**FuncName** | **string** |  | 
**Module** | [**Module**](Module.md) |  | 

## Methods

### NewCodeMessageParams

`func NewCodeMessageParams(message string, line float64, funcName string, module Module, ) *CodeMessageParams`

NewCodeMessageParams instantiates a new CodeMessageParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeMessageParamsWithDefaults

`func NewCodeMessageParamsWithDefaults() *CodeMessageParams`

NewCodeMessageParamsWithDefaults instantiates a new CodeMessageParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CodeMessageParams) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CodeMessageParams) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CodeMessageParams) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessageCode

`func (o *CodeMessageParams) GetMessageCode() string`

GetMessageCode returns the MessageCode field if non-nil, zero value otherwise.

### GetMessageCodeOk

`func (o *CodeMessageParams) GetMessageCodeOk() (*string, bool)`

GetMessageCodeOk returns a tuple with the MessageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCode

`func (o *CodeMessageParams) SetMessageCode(v string)`

SetMessageCode sets MessageCode field to given value.

### HasMessageCode

`func (o *CodeMessageParams) HasMessageCode() bool`

HasMessageCode returns a boolean if a field has been set.

### GetLine

`func (o *CodeMessageParams) GetLine() float64`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *CodeMessageParams) GetLineOk() (*float64, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *CodeMessageParams) SetLine(v float64)`

SetLine sets Line field to given value.


### GetFuncName

`func (o *CodeMessageParams) GetFuncName() string`

GetFuncName returns the FuncName field if non-nil, zero value otherwise.

### GetFuncNameOk

`func (o *CodeMessageParams) GetFuncNameOk() (*string, bool)`

GetFuncNameOk returns a tuple with the FuncName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncName

`func (o *CodeMessageParams) SetFuncName(v string)`

SetFuncName sets FuncName field to given value.


### GetModule

`func (o *CodeMessageParams) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CodeMessageParams) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CodeMessageParams) SetModule(v Module)`

SetModule sets Module field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


