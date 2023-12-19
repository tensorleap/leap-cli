/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.422
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// JobTrigger the model 'JobTrigger'
type JobTrigger string

// List of JobTrigger
const (
	JOBTRIGGER_MANUAL JobTrigger = "Manual"
	JOBTRIGGER_AUTO   JobTrigger = "Auto"
)

// All allowed values of JobTrigger enum
var AllowedJobTriggerEnumValues = []JobTrigger{
	"Manual",
	"Auto",
}

func (v *JobTrigger) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JobTrigger(value)
	for _, existing := range AllowedJobTriggerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JobTrigger", value)
}

// NewJobTriggerFromValue returns a pointer to a valid JobTrigger
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJobTriggerFromValue(v string) (*JobTrigger, error) {
	ev := JobTrigger(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JobTrigger: valid values are %v", v, AllowedJobTriggerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JobTrigger) IsValid() bool {
	for _, existing := range AllowedJobTriggerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JobTrigger value
func (v JobTrigger) Ptr() *JobTrigger {
	return &v
}

type NullableJobTrigger struct {
	value *JobTrigger
	isSet bool
}

func (v NullableJobTrigger) Get() *JobTrigger {
	return v.value
}

func (v *NullableJobTrigger) Set(val *JobTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableJobTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableJobTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobTrigger(val *JobTrigger) *NullableJobTrigger {
	return &NullableJobTrigger{value: val, isSet: true}
}

func (v NullableJobTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
