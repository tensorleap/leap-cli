# DatasetMessageParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**MessageCode** | Pointer to **string** |  | [optional] 
**Line** | **float64** |  | 
**FuncName** | **string** |  | 
**Module** | [**Module**](Module.md) |  | 

## Methods

### NewDatasetMessageParams

`func NewDatasetMessageParams(message string, line float64, funcName string, module Module, ) *DatasetMessageParams`

NewDatasetMessageParams instantiates a new DatasetMessageParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetMessageParamsWithDefaults

`func NewDatasetMessageParamsWithDefaults() *DatasetMessageParams`

NewDatasetMessageParamsWithDefaults instantiates a new DatasetMessageParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DatasetMessageParams) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DatasetMessageParams) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DatasetMessageParams) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessageCode

`func (o *DatasetMessageParams) GetMessageCode() string`

GetMessageCode returns the MessageCode field if non-nil, zero value otherwise.

### GetMessageCodeOk

`func (o *DatasetMessageParams) GetMessageCodeOk() (*string, bool)`

GetMessageCodeOk returns a tuple with the MessageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCode

`func (o *DatasetMessageParams) SetMessageCode(v string)`

SetMessageCode sets MessageCode field to given value.

### HasMessageCode

`func (o *DatasetMessageParams) HasMessageCode() bool`

HasMessageCode returns a boolean if a field has been set.

### GetLine

`func (o *DatasetMessageParams) GetLine() float64`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *DatasetMessageParams) GetLineOk() (*float64, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *DatasetMessageParams) SetLine(v float64)`

SetLine sets Line field to given value.


### GetFuncName

`func (o *DatasetMessageParams) GetFuncName() string`

GetFuncName returns the FuncName field if non-nil, zero value otherwise.

### GetFuncNameOk

`func (o *DatasetMessageParams) GetFuncNameOk() (*string, bool)`

GetFuncNameOk returns a tuple with the FuncName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncName

`func (o *DatasetMessageParams) SetFuncName(v string)`

SetFuncName sets FuncName field to given value.


### GetModule

`func (o *DatasetMessageParams) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *DatasetMessageParams) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *DatasetMessageParams) SetModule(v Module)`

SetModule sets Module field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


