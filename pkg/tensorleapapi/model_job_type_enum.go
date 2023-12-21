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

// JobTypeEnum the model 'JobTypeEnum'
type JobTypeEnum string

// List of JobTypeEnum
const (
	JOBTYPEENUM_TRAINING           JobTypeEnum = "TRAINING"
	JOBTYPEENUM_IMPORT_MODEL       JobTypeEnum = "IMPORT_MODEL"
	JOBTYPEENUM_ANALYZE            JobTypeEnum = "ANALYZE"
	JOBTYPEENUM_WARMUP             JobTypeEnum = "WARMUP"
	JOBTYPEENUM_TEST_STUB_FUNCTION JobTypeEnum = "TEST_STUB_FUNCTION"
	JOBTYPEENUM_TEST_CUSTOM_LOSS   JobTypeEnum = "TEST_CUSTOM_LOSS"
	JOBTYPEENUM_EXPORT_MODEL       JobTypeEnum = "EXPORT_MODEL"
	JOBTYPEENUM_DATASET_PARSE      JobTypeEnum = "DATASET_PARSE"
	JOBTYPEENUM_ANALYZE_GRAPH      JobTypeEnum = "ANALYZE_GRAPH"
	JOBTYPEENUM_DRY_RUN_GRAPH      JobTypeEnum = "DRY_RUN_GRAPH"
	JOBTYPEENUM_SLIM_LS            JobTypeEnum = "SLIM_LS"
)

// All allowed values of JobTypeEnum enum
var AllowedJobTypeEnumEnumValues = []JobTypeEnum{
	"TRAINING",
	"IMPORT_MODEL",
	"ANALYZE",
	"WARMUP",
	"TEST_STUB_FUNCTION",
	"TEST_CUSTOM_LOSS",
	"EXPORT_MODEL",
	"DATASET_PARSE",
	"ANALYZE_GRAPH",
	"DRY_RUN_GRAPH",
	"SLIM_LS",
}

func (v *JobTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JobTypeEnum(value)
	for _, existing := range AllowedJobTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JobTypeEnum", value)
}

// NewJobTypeEnumFromValue returns a pointer to a valid JobTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJobTypeEnumFromValue(v string) (*JobTypeEnum, error) {
	ev := JobTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JobTypeEnum: valid values are %v", v, AllowedJobTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JobTypeEnum) IsValid() bool {
	for _, existing := range AllowedJobTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JobTypeEnum value
func (v JobTypeEnum) Ptr() *JobTypeEnum {
	return &v
}

type NullableJobTypeEnum struct {
	value *JobTypeEnum
	isSet bool
}

func (v NullableJobTypeEnum) Get() *JobTypeEnum {
	return v.value
}

func (v *NullableJobTypeEnum) Set(val *JobTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableJobTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableJobTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobTypeEnum(val *JobTypeEnum) *NullableJobTypeEnum {
	return &NullableJobTypeEnum{value: val, isSet: true}
}

func (v NullableJobTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
