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

// ScatterInsightType the model 'ScatterInsightType'
type ScatterInsightType string

// List of ScatterInsightType
const (
	SCATTERINSIGHTTYPE_REDUNDANT_SAMPLES_CLUSTER    ScatterInsightType = "redundant_samples_cluster"
	SCATTERINSIGHTTYPE_WEAK_CLUSTER                 ScatterInsightType = "weak_cluster"
	SCATTERINSIGHTTYPE_UNDER_REPRESENTATION_CLUSTER ScatterInsightType = "under_representation_cluster"
	SCATTERINSIGHTTYPE_OVERFITTING_CLUSTER          ScatterInsightType = "overfitting_cluster"
	SCATTERINSIGHTTYPE_DUPLICATION_INSIGHT          ScatterInsightType = "duplication_insight"
)

// All allowed values of ScatterInsightType enum
var AllowedScatterInsightTypeEnumValues = []ScatterInsightType{
	"redundant_samples_cluster",
	"weak_cluster",
	"under_representation_cluster",
	"overfitting_cluster",
	"duplication_insight",
}

func (v *ScatterInsightType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScatterInsightType(value)
	for _, existing := range AllowedScatterInsightTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScatterInsightType", value)
}

// NewScatterInsightTypeFromValue returns a pointer to a valid ScatterInsightType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScatterInsightTypeFromValue(v string) (*ScatterInsightType, error) {
	ev := ScatterInsightType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScatterInsightType: valid values are %v", v, AllowedScatterInsightTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScatterInsightType) IsValid() bool {
	for _, existing := range AllowedScatterInsightTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterInsightType value
func (v ScatterInsightType) Ptr() *ScatterInsightType {
	return &v
}

type NullableScatterInsightType struct {
	value *ScatterInsightType
	isSet bool
}

func (v NullableScatterInsightType) Get() *ScatterInsightType {
	return v.value
}

func (v *NullableScatterInsightType) Set(val *ScatterInsightType) {
	v.value = val
	v.isSet = true
}

func (v NullableScatterInsightType) IsSet() bool {
	return v.isSet
}

func (v *NullableScatterInsightType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScatterInsightType(val *ScatterInsightType) *NullableScatterInsightType {
	return &NullableScatterInsightType{value: val, isSet: true}
}

func (v NullableScatterInsightType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScatterInsightType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
