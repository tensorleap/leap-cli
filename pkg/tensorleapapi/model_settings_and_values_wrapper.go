/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.423
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SettingsAndValuesWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsAndValuesWrapper{}

// SettingsAndValuesWrapper struct for SettingsAndValuesWrapper
type SettingsAndValuesWrapper struct {
	Schema []SchemaWithKey `json:"schema"`
	Values []ValueWithKey `json:"values"`
}

// NewSettingsAndValuesWrapper instantiates a new SettingsAndValuesWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsAndValuesWrapper(schema []SchemaWithKey, values []ValueWithKey) *SettingsAndValuesWrapper {
	this := SettingsAndValuesWrapper{}
	this.Schema = schema
	this.Values = values
	return &this
}

// NewSettingsAndValuesWrapperWithDefaults instantiates a new SettingsAndValuesWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsAndValuesWrapperWithDefaults() *SettingsAndValuesWrapper {
	this := SettingsAndValuesWrapper{}
	return &this
}

// GetSchema returns the Schema field value
func (o *SettingsAndValuesWrapper) GetSchema() []SchemaWithKey {
	if o == nil {
		var ret []SchemaWithKey
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *SettingsAndValuesWrapper) GetSchemaOk() ([]SchemaWithKey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schema, true
}

// SetSchema sets field value
func (o *SettingsAndValuesWrapper) SetSchema(v []SchemaWithKey) {
	o.Schema = v
}

// GetValues returns the Values field value
func (o *SettingsAndValuesWrapper) GetValues() []ValueWithKey {
	if o == nil {
		var ret []ValueWithKey
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *SettingsAndValuesWrapper) GetValuesOk() ([]ValueWithKey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *SettingsAndValuesWrapper) SetValues(v []ValueWithKey) {
	o.Values = v
}

func (o SettingsAndValuesWrapper) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsAndValuesWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schema"] = o.Schema
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

type NullableSettingsAndValuesWrapper struct {
	value *SettingsAndValuesWrapper
	isSet bool
}

func (v NullableSettingsAndValuesWrapper) Get() *SettingsAndValuesWrapper {
	return v.value
}

func (v *NullableSettingsAndValuesWrapper) Set(val *SettingsAndValuesWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsAndValuesWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsAndValuesWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsAndValuesWrapper(val *SettingsAndValuesWrapper) *NullableSettingsAndValuesWrapper {
	return &NullableSettingsAndValuesWrapper{value: val, isSet: true}
}

func (v NullableSettingsAndValuesWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsAndValuesWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


