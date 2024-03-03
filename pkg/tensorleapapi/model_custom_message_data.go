/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the CustomMessageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomMessageData{}

// CustomMessageData struct for CustomMessageData
type CustomMessageData struct {
	Level  MessageLevel            `json:"level"`
	Params CustomMessageDataParams `json:"params"`
}

// NewCustomMessageData instantiates a new CustomMessageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomMessageData(level MessageLevel, params CustomMessageDataParams) *CustomMessageData {
	this := CustomMessageData{}
	this.Level = level
	this.Params = params
	return &this
}

// NewCustomMessageDataWithDefaults instantiates a new CustomMessageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomMessageDataWithDefaults() *CustomMessageData {
	this := CustomMessageData{}
	return &this
}

// GetLevel returns the Level field value
func (o *CustomMessageData) GetLevel() MessageLevel {
	if o == nil {
		var ret MessageLevel
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *CustomMessageData) GetLevelOk() (*MessageLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *CustomMessageData) SetLevel(v MessageLevel) {
	o.Level = v
}

// GetParams returns the Params field value
func (o *CustomMessageData) GetParams() CustomMessageDataParams {
	if o == nil {
		var ret CustomMessageDataParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *CustomMessageData) GetParamsOk() (*CustomMessageDataParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *CustomMessageData) SetParams(v CustomMessageDataParams) {
	o.Params = v
}

func (o CustomMessageData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomMessageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

type NullableCustomMessageData struct {
	value *CustomMessageData
	isSet bool
}

func (v NullableCustomMessageData) Get() *CustomMessageData {
	return v.value
}

func (v *NullableCustomMessageData) Set(val *CustomMessageData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomMessageData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomMessageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomMessageData(val *CustomMessageData) *NullableCustomMessageData {
	return &NullableCustomMessageData{value: val, isSet: true}
}

func (v NullableCustomMessageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomMessageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
