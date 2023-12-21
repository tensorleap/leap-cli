/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.423
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// SettingSchema struct for SettingSchema
type SettingSchema struct {
	BoolSchema   *BoolSchema
	NumberSchema *NumberSchema
	StringSchema *StringSchema
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SettingSchema) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BoolSchema
	err = json.Unmarshal(data, &dst.BoolSchema)
	if err == nil {
		jsonBoolSchema, _ := json.Marshal(dst.BoolSchema)
		if string(jsonBoolSchema) == "{}" { // empty struct
			dst.BoolSchema = nil
		} else {
			return nil // data stored in dst.BoolSchema, return on the first match
		}
	} else {
		dst.BoolSchema = nil
	}

	// try to unmarshal JSON data into NumberSchema
	err = json.Unmarshal(data, &dst.NumberSchema)
	if err == nil {
		jsonNumberSchema, _ := json.Marshal(dst.NumberSchema)
		if string(jsonNumberSchema) == "{}" { // empty struct
			dst.NumberSchema = nil
		} else {
			return nil // data stored in dst.NumberSchema, return on the first match
		}
	} else {
		dst.NumberSchema = nil
	}

	// try to unmarshal JSON data into StringSchema
	err = json.Unmarshal(data, &dst.StringSchema)
	if err == nil {
		jsonStringSchema, _ := json.Marshal(dst.StringSchema)
		if string(jsonStringSchema) == "{}" { // empty struct
			dst.StringSchema = nil
		} else {
			return nil // data stored in dst.StringSchema, return on the first match
		}
	} else {
		dst.StringSchema = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SettingSchema)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SettingSchema) MarshalJSON() ([]byte, error) {
	if src.BoolSchema != nil {
		return json.Marshal(&src.BoolSchema)
	}

	if src.NumberSchema != nil {
		return json.Marshal(&src.NumberSchema)
	}

	if src.StringSchema != nil {
		return json.Marshal(&src.StringSchema)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSettingSchema struct {
	value *SettingSchema
	isSet bool
}

func (v NullableSettingSchema) Get() *SettingSchema {
	return v.value
}

func (v *NullableSettingSchema) Set(val *SettingSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingSchema(val *SettingSchema) *NullableSettingSchema {
	return &NullableSettingSchema{value: val, isSet: true}
}

func (v NullableSettingSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
