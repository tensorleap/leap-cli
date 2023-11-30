/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.417
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SetSettingValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetSettingValues{}

// SetSettingValues struct for SetSettingValues
type SetSettingValues struct {
	Unset []string               `json:"unset,omitempty"`
	Set   map[string]interface{} `json:"set,omitempty"`
}

// NewSetSettingValues instantiates a new SetSettingValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetSettingValues() *SetSettingValues {
	this := SetSettingValues{}
	return &this
}

// NewSetSettingValuesWithDefaults instantiates a new SetSettingValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetSettingValuesWithDefaults() *SetSettingValues {
	this := SetSettingValues{}
	return &this
}

// GetUnset returns the Unset field value if set, zero value otherwise.
func (o *SetSettingValues) GetUnset() []string {
	if o == nil || IsNil(o.Unset) {
		var ret []string
		return ret
	}
	return o.Unset
}

// GetUnsetOk returns a tuple with the Unset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetSettingValues) GetUnsetOk() ([]string, bool) {
	if o == nil || IsNil(o.Unset) {
		return nil, false
	}
	return o.Unset, true
}

// HasUnset returns a boolean if a field has been set.
func (o *SetSettingValues) HasUnset() bool {
	if o != nil && !IsNil(o.Unset) {
		return true
	}

	return false
}

// SetUnset gets a reference to the given []string and assigns it to the Unset field.
func (o *SetSettingValues) SetUnset(v []string) {
	o.Unset = v
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *SetSettingValues) GetSet() map[string]interface{} {
	if o == nil || IsNil(o.Set) {
		var ret map[string]interface{}
		return ret
	}
	return o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetSettingValues) GetSetOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Set) {
		return map[string]interface{}{}, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *SetSettingValues) HasSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given map[string]interface{} and assigns it to the Set field.
func (o *SetSettingValues) SetSet(v map[string]interface{}) {
	o.Set = v
}

func (o SetSettingValues) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetSettingValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Unset) {
		toSerialize["unset"] = o.Unset
	}
	if !IsNil(o.Set) {
		toSerialize["set"] = o.Set
	}
	return toSerialize, nil
}

type NullableSetSettingValues struct {
	value *SetSettingValues
	isSet bool
}

func (v NullableSetSettingValues) Get() *SetSettingValues {
	return v.value
}

func (v *NullableSetSettingValues) Set(val *SetSettingValues) {
	v.value = val
	v.isSet = true
}

func (v NullableSetSettingValues) IsSet() bool {
	return v.isSet
}

func (v *NullableSetSettingValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetSettingValues(val *SetSettingValues) *NullableSetSettingValues {
	return &NullableSetSettingValues{value: val, isSet: true}
}

func (v NullableSetSettingValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetSettingValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}