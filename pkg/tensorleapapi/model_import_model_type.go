/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.301
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// ImportModelType the model 'ImportModelType'
type ImportModelType string

// List of ImportModelType
const (
	IMPORTMODELTYPE_JSON_TF2 ImportModelType = "JSON_TF2"
	IMPORTMODELTYPE_ONNX     ImportModelType = "ONNX"
	IMPORTMODELTYPE_PB_TF2   ImportModelType = "PB_TF2"
	IMPORTMODELTYPE_H5_TF2   ImportModelType = "H5_TF2"
)

// All allowed values of ImportModelType enum
var AllowedImportModelTypeEnumValues = []ImportModelType{
	"JSON_TF2",
	"ONNX",
	"PB_TF2",
	"H5_TF2",
}

func (v *ImportModelType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImportModelType(value)
	for _, existing := range AllowedImportModelTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImportModelType", value)
}

// NewImportModelTypeFromValue returns a pointer to a valid ImportModelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImportModelTypeFromValue(v string) (*ImportModelType, error) {
	ev := ImportModelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImportModelType: valid values are %v", v, AllowedImportModelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImportModelType) IsValid() bool {
	for _, existing := range AllowedImportModelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImportModelType value
func (v ImportModelType) Ptr() *ImportModelType {
	return &v
}

type NullableImportModelType struct {
	value *ImportModelType
	isSet bool
}

func (v NullableImportModelType) Get() *ImportModelType {
	return v.value
}

func (v *NullableImportModelType) Set(val *ImportModelType) {
	v.value = val
	v.isSet = true
}

func (v NullableImportModelType) IsSet() bool {
	return v.isSet
}

func (v *NullableImportModelType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportModelType(val *ImportModelType) *NullableImportModelType {
	return &NullableImportModelType{value: val, isSet: true}
}

func (v NullableImportModelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportModelType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
