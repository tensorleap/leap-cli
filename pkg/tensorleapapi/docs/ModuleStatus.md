# ModuleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Name** | [**MonitoredModuleType**](MonitoredModuleType.md) |  | 

## Methods

### NewModuleStatus

`func NewModuleStatus(status bool, name MonitoredModuleType, ) *ModuleStatus`

NewModuleStatus instantiates a new ModuleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleStatusWithDefaults

`func NewModuleStatusWithDefaults() *ModuleStatus`

NewModuleStatusWithDefaults instantiates a new ModuleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ModuleStatus) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModuleStatus) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModuleStatus) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetError

`func (o *ModuleStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ModuleStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ModuleStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ModuleStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetName

`func (o *ModuleStatus) GetName() MonitoredModuleType`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModuleStatus) GetNameOk() (*MonitoredModuleType, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModuleStatus) SetName(v MonitoredModuleType)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


