# ValueWithKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyName** | **string** |  | 
**Value** | [**NullableSettingValue**](SettingValue.md) |  | 

## Methods

### NewValueWithKey

`func NewValueWithKey(keyName string, value NullableSettingValue, ) *ValueWithKey`

NewValueWithKey instantiates a new ValueWithKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueWithKeyWithDefaults

`func NewValueWithKeyWithDefaults() *ValueWithKey`

NewValueWithKeyWithDefaults instantiates a new ValueWithKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyName

`func (o *ValueWithKey) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *ValueWithKey) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *ValueWithKey) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetValue

`func (o *ValueWithKey) GetValue() SettingValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ValueWithKey) GetValueOk() (*SettingValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ValueWithKey) SetValue(v SettingValue)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ValueWithKey) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ValueWithKey) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


