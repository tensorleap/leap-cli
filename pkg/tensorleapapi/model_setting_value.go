/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// SettingValue struct for SettingValue
type SettingValue struct {
	bool    *bool
	float64 *float64
	string  *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SettingValue) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into bool
	err = json.Unmarshal(data, &dst.bool)
	if err == nil {
		jsonbool, _ := json.Marshal(dst.bool)
		if string(jsonbool) == "{}" { // empty struct
			dst.bool = nil
		} else {
			return nil // data stored in dst.bool, return on the first match
		}
	} else {
		dst.bool = nil
	}

	// try to unmarshal JSON data into float64
	err = json.Unmarshal(data, &dst.float64)
	if err == nil {
		jsonfloat64, _ := json.Marshal(dst.float64)
		if string(jsonfloat64) == "{}" { // empty struct
			dst.float64 = nil
		} else {
			return nil // data stored in dst.float64, return on the first match
		}
	} else {
		dst.float64 = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SettingValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SettingValue) MarshalJSON() ([]byte, error) {
	if src.bool != nil {
		return json.Marshal(&src.bool)
	}

	if src.float64 != nil {
		return json.Marshal(&src.float64)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSettingValue struct {
	value *SettingValue
	isSet bool
}

func (v NullableSettingValue) Get() *SettingValue {
	return v.value
}

func (v *NullableSettingValue) Set(val *SettingValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingValue(val *SettingValue) *NullableSettingValue {
	return &NullableSettingValue{value: val, isSet: true}
}

func (v NullableSettingValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
