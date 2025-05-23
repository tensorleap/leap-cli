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

// checks if the SetSettingValueWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetSettingValueWrapper{}

// SetSettingValueWrapper struct for SetSettingValueWrapper
type SetSettingValueWrapper struct {
	Unset []string       `json:"unset,omitempty"`
	Set   []ValueWithKey `json:"set,omitempty"`
}

// NewSetSettingValueWrapper instantiates a new SetSettingValueWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetSettingValueWrapper() *SetSettingValueWrapper {
	this := SetSettingValueWrapper{}
	return &this
}

// NewSetSettingValueWrapperWithDefaults instantiates a new SetSettingValueWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetSettingValueWrapperWithDefaults() *SetSettingValueWrapper {
	this := SetSettingValueWrapper{}
	return &this
}

// GetUnset returns the Unset field value if set, zero value otherwise.
func (o *SetSettingValueWrapper) GetUnset() []string {
	if o == nil || IsNil(o.Unset) {
		var ret []string
		return ret
	}
	return o.Unset
}

// GetUnsetOk returns a tuple with the Unset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetSettingValueWrapper) GetUnsetOk() ([]string, bool) {
	if o == nil || IsNil(o.Unset) {
		return nil, false
	}
	return o.Unset, true
}

// HasUnset returns a boolean if a field has been set.
func (o *SetSettingValueWrapper) HasUnset() bool {
	if o != nil && !IsNil(o.Unset) {
		return true
	}

	return false
}

// SetUnset gets a reference to the given []string and assigns it to the Unset field.
func (o *SetSettingValueWrapper) SetUnset(v []string) {
	o.Unset = v
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *SetSettingValueWrapper) GetSet() []ValueWithKey {
	if o == nil || IsNil(o.Set) {
		var ret []ValueWithKey
		return ret
	}
	return o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetSettingValueWrapper) GetSetOk() ([]ValueWithKey, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *SetSettingValueWrapper) HasSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given []ValueWithKey and assigns it to the Set field.
func (o *SetSettingValueWrapper) SetSet(v []ValueWithKey) {
	o.Set = v
}

func (o SetSettingValueWrapper) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetSettingValueWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Unset) {
		toSerialize["unset"] = o.Unset
	}
	if !IsNil(o.Set) {
		toSerialize["set"] = o.Set
	}
	return toSerialize, nil
}

type NullableSetSettingValueWrapper struct {
	value *SetSettingValueWrapper
	isSet bool
}

func (v NullableSetSettingValueWrapper) Get() *SetSettingValueWrapper {
	return v.value
}

func (v *NullableSetSettingValueWrapper) Set(val *SetSettingValueWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableSetSettingValueWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableSetSettingValueWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetSettingValueWrapper(val *SetSettingValueWrapper) *NullableSetSettingValueWrapper {
	return &NullableSetSettingValueWrapper{value: val, isSet: true}
}

func (v NullableSetSettingValueWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetSettingValueWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
