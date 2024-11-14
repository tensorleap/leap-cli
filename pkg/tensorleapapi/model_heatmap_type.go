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

// HeatmapType the model 'HeatmapType'
type HeatmapType string

// List of HeatmapType
const (
	HEATMAPTYPE_SINGLE   HeatmapType = "single"
	HEATMAPTYPE_MULTIPLE HeatmapType = "multiple"
)

// All allowed values of HeatmapType enum
var AllowedHeatmapTypeEnumValues = []HeatmapType{
	"single",
	"multiple",
}

func (v *HeatmapType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HeatmapType(value)
	for _, existing := range AllowedHeatmapTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HeatmapType", value)
}

// NewHeatmapTypeFromValue returns a pointer to a valid HeatmapType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHeatmapTypeFromValue(v string) (*HeatmapType, error) {
	ev := HeatmapType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HeatmapType: valid values are %v", v, AllowedHeatmapTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HeatmapType) IsValid() bool {
	for _, existing := range AllowedHeatmapTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatmapType value
func (v HeatmapType) Ptr() *HeatmapType {
	return &v
}

type NullableHeatmapType struct {
	value *HeatmapType
	isSet bool
}

func (v NullableHeatmapType) Get() *HeatmapType {
	return v.value
}

func (v *NullableHeatmapType) Set(val *HeatmapType) {
	v.value = val
	v.isSet = true
}

func (v NullableHeatmapType) IsSet() bool {
	return v.isSet
}

func (v *NullableHeatmapType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeatmapType(val *HeatmapType) *NullableHeatmapType {
	return &NullableHeatmapType{value: val, isSet: true}
}

func (v NullableHeatmapType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeatmapType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
