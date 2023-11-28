/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.417
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// AnalyzeTypeEnum the model 'AnalyzeTypeEnum'
type AnalyzeTypeEnum string

// List of AnalyzeTypeEnum
const (
	ANALYZETYPEENUM_POPULATION_EXPLORATION  AnalyzeTypeEnum = "population_exploration"
	ANALYZETYPEENUM_SAMPLE_SELECTION        AnalyzeTypeEnum = "sample_selection"
	ANALYZETYPEENUM_SAMPLE_ANALYSIS         AnalyzeTypeEnum = "sample_analysis"
	ANALYZETYPEENUM_VISUALIZERS_CALCULATION AnalyzeTypeEnum = "visualizers_calculation"
)

// All allowed values of AnalyzeTypeEnum enum
var AllowedAnalyzeTypeEnumEnumValues = []AnalyzeTypeEnum{
	"population_exploration",
	"sample_selection",
	"sample_analysis",
	"visualizers_calculation",
}

func (v *AnalyzeTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AnalyzeTypeEnum(value)
	for _, existing := range AllowedAnalyzeTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AnalyzeTypeEnum", value)
}

// NewAnalyzeTypeEnumFromValue returns a pointer to a valid AnalyzeTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAnalyzeTypeEnumFromValue(v string) (*AnalyzeTypeEnum, error) {
	ev := AnalyzeTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AnalyzeTypeEnum: valid values are %v", v, AllowedAnalyzeTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AnalyzeTypeEnum) IsValid() bool {
	for _, existing := range AllowedAnalyzeTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnalyzeTypeEnum value
func (v AnalyzeTypeEnum) Ptr() *AnalyzeTypeEnum {
	return &v
}

type NullableAnalyzeTypeEnum struct {
	value *AnalyzeTypeEnum
	isSet bool
}

func (v NullableAnalyzeTypeEnum) Get() *AnalyzeTypeEnum {
	return v.value
}

func (v *NullableAnalyzeTypeEnum) Set(val *AnalyzeTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyzeTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyzeTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyzeTypeEnum(val *AnalyzeTypeEnum) *NullableAnalyzeTypeEnum {
	return &NullableAnalyzeTypeEnum{value: val, isSet: true}
}

func (v NullableAnalyzeTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyzeTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
