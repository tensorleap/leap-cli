/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// NumberOrString struct for NumberOrString
type NumberOrString struct {
	float64 *float64
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NumberOrString) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into float64
	err = json.Unmarshal(data, &dst.float64);
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
	err = json.Unmarshal(data, &dst.string);
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

	return fmt.Errorf("data failed to match schemas in anyOf(NumberOrString)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NumberOrString) MarshalJSON() ([]byte, error) {
	if src.float64 != nil {
		return json.Marshal(&src.float64)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNumberOrString struct {
	value *NumberOrString
	isSet bool
}

func (v NullableNumberOrString) Get() *NumberOrString {
	return v.value
}

func (v *NullableNumberOrString) Set(val *NumberOrString) {
	v.value = val
	v.isSet = true
}

func (v NullableNumberOrString) IsSet() bool {
	return v.isSet
}

func (v *NullableNumberOrString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumberOrString(val *NumberOrString) *NullableNumberOrString {
	return &NullableNumberOrString{value: val, isSet: true}
}

func (v NullableNumberOrString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumberOrString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


