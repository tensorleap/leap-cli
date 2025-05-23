/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ValueWithKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueWithKey{}

// ValueWithKey struct for ValueWithKey
type ValueWithKey struct {
	KeyName string               `json:"keyName"`
	Value   NullableSettingValue `json:"value"`
}

// NewValueWithKey instantiates a new ValueWithKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueWithKey(keyName string, value NullableSettingValue) *ValueWithKey {
	this := ValueWithKey{}
	this.KeyName = keyName
	this.Value = value
	return &this
}

// NewValueWithKeyWithDefaults instantiates a new ValueWithKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueWithKeyWithDefaults() *ValueWithKey {
	this := ValueWithKey{}
	return &this
}

// GetKeyName returns the KeyName field value
func (o *ValueWithKey) GetKeyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *ValueWithKey) GetKeyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *ValueWithKey) SetKeyName(v string) {
	o.KeyName = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for SettingValue will be returned
func (o *ValueWithKey) GetValue() SettingValue {
	if o == nil || o.Value.Get() == nil {
		var ret SettingValue
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValueWithKey) GetValueOk() (*SettingValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *ValueWithKey) SetValue(v SettingValue) {
	o.Value.Set(&v)
}

func (o ValueWithKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueWithKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keyName"] = o.KeyName
	toSerialize["value"] = o.Value.Get()
	return toSerialize, nil
}

type NullableValueWithKey struct {
	value *ValueWithKey
	isSet bool
}

func (v NullableValueWithKey) Get() *ValueWithKey {
	return v.value
}

func (v *NullableValueWithKey) Set(val *ValueWithKey) {
	v.value = val
	v.isSet = true
}

func (v NullableValueWithKey) IsSet() bool {
	return v.isSet
}

func (v *NullableValueWithKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueWithKey(val *ValueWithKey) *NullableValueWithKey {
	return &NullableValueWithKey{value: val, isSet: true}
}

func (v NullableValueWithKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueWithKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
