# SetSettingValueWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unset** | Pointer to **[]string** |  | [optional] 
**Set** | Pointer to [**[]ValueWithKey**](ValueWithKey.md) |  | [optional] 
**ProjectId** | **string** |  | 

## Methods

### NewSetSettingValueWrapper

`func NewSetSettingValueWrapper(projectId string, ) *SetSettingValueWrapper`

NewSetSettingValueWrapper instantiates a new SetSettingValueWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetSettingValueWrapperWithDefaults

`func NewSetSettingValueWrapperWithDefaults() *SetSettingValueWrapper`

NewSetSettingValueWrapperWithDefaults instantiates a new SetSettingValueWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnset

`func (o *SetSettingValueWrapper) GetUnset() []string`

GetUnset returns the Unset field if non-nil, zero value otherwise.

### GetUnsetOk

`func (o *SetSettingValueWrapper) GetUnsetOk() (*[]string, bool)`

GetUnsetOk returns a tuple with the Unset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnset

`func (o *SetSettingValueWrapper) SetUnset(v []string)`

SetUnset sets Unset field to given value.

### HasUnset

`func (o *SetSettingValueWrapper) HasUnset() bool`

HasUnset returns a boolean if a field has been set.

### GetSet

`func (o *SetSettingValueWrapper) GetSet() []ValueWithKey`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *SetSettingValueWrapper) GetSetOk() (*[]ValueWithKey, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *SetSettingValueWrapper) SetSet(v []ValueWithKey)`

SetSet sets Set field to given value.

### HasSet

`func (o *SetSettingValueWrapper) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetProjectId

`func (o *SetSettingValueWrapper) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SetSettingValueWrapper) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SetSettingValueWrapper) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


