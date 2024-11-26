/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.557
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// CodeIntegrationBinder struct for CodeIntegrationBinder
type CodeIntegrationBinder struct {
	CodeIntegrationBinderByBranch *CodeIntegrationBinderByBranch
	CodeIntegrationBinderById     *CodeIntegrationBinderById
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CodeIntegrationBinder) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CodeIntegrationBinderByBranch
	err = json.Unmarshal(data, &dst.CodeIntegrationBinderByBranch)
	if err == nil {
		jsonCodeIntegrationBinderByBranch, _ := json.Marshal(dst.CodeIntegrationBinderByBranch)
		if string(jsonCodeIntegrationBinderByBranch) == "{}" { // empty struct
			dst.CodeIntegrationBinderByBranch = nil
		} else {
			return nil // data stored in dst.CodeIntegrationBinderByBranch, return on the first match
		}
	} else {
		dst.CodeIntegrationBinderByBranch = nil
	}

	// try to unmarshal JSON data into CodeIntegrationBinderById
	err = json.Unmarshal(data, &dst.CodeIntegrationBinderById)
	if err == nil {
		jsonCodeIntegrationBinderById, _ := json.Marshal(dst.CodeIntegrationBinderById)
		if string(jsonCodeIntegrationBinderById) == "{}" { // empty struct
			dst.CodeIntegrationBinderById = nil
		} else {
			return nil // data stored in dst.CodeIntegrationBinderById, return on the first match
		}
	} else {
		dst.CodeIntegrationBinderById = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CodeIntegrationBinder)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CodeIntegrationBinder) MarshalJSON() ([]byte, error) {
	if src.CodeIntegrationBinderByBranch != nil {
		return json.Marshal(&src.CodeIntegrationBinderByBranch)
	}

	if src.CodeIntegrationBinderById != nil {
		return json.Marshal(&src.CodeIntegrationBinderById)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCodeIntegrationBinder struct {
	value *CodeIntegrationBinder
	isSet bool
}

func (v NullableCodeIntegrationBinder) Get() *CodeIntegrationBinder {
	return v.value
}

func (v *NullableCodeIntegrationBinder) Set(val *CodeIntegrationBinder) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeIntegrationBinder) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeIntegrationBinder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeIntegrationBinder(val *CodeIntegrationBinder) *NullableCodeIntegrationBinder {
	return &NullableCodeIntegrationBinder{value: val, isSet: true}
}

func (v NullableCodeIntegrationBinder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeIntegrationBinder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}