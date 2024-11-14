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

// GeneratedDashletType the model 'GeneratedDashletType'
type GeneratedDashletType string

// List of GeneratedDashletType
const (
	GENERATEDDASHLETTYPE_HEATMAP GeneratedDashletType = "heatmap"
	GENERATEDDASHLETTYPE_BAR     GeneratedDashletType = "bar"
	GENERATEDDASHLETTYPE_LINE    GeneratedDashletType = "line"
)

// All allowed values of GeneratedDashletType enum
var AllowedGeneratedDashletTypeEnumValues = []GeneratedDashletType{
	"heatmap",
	"bar",
	"line",
}

func (v *GeneratedDashletType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GeneratedDashletType(value)
	for _, existing := range AllowedGeneratedDashletTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GeneratedDashletType", value)
}

// NewGeneratedDashletTypeFromValue returns a pointer to a valid GeneratedDashletType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGeneratedDashletTypeFromValue(v string) (*GeneratedDashletType, error) {
	ev := GeneratedDashletType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GeneratedDashletType: valid values are %v", v, AllowedGeneratedDashletTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GeneratedDashletType) IsValid() bool {
	for _, existing := range AllowedGeneratedDashletTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeneratedDashletType value
func (v GeneratedDashletType) Ptr() *GeneratedDashletType {
	return &v
}

type NullableGeneratedDashletType struct {
	value *GeneratedDashletType
	isSet bool
}

func (v NullableGeneratedDashletType) Get() *GeneratedDashletType {
	return v.value
}

func (v *NullableGeneratedDashletType) Set(val *GeneratedDashletType) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneratedDashletType) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneratedDashletType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneratedDashletType(val *GeneratedDashletType) *NullableGeneratedDashletType {
	return &NullableGeneratedDashletType{value: val, isSet: true}
}

func (v NullableGeneratedDashletType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneratedDashletType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
