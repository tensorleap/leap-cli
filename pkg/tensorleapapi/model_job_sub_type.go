/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// JobSubType the model 'JobSubType'
type JobSubType string

// List of JobSubType
const (
	JOBSUBTYPE_EVALUATE                   JobSubType = "Evaluate"
	JOBSUBTYPE_TRAIN_FROM_SCRATCH         JobSubType = "Train From Scratch"
	JOBSUBTYPE_CONTINUE_TRAINING          JobSubType = "Continue Training"
	JOBSUBTYPE_TRAIN_FROM_INITIAL_WEIGHTS JobSubType = "Train From Initial Weights"
	JOBSUBTYPE_POPULATION_EXPLORATION     JobSubType = "Population Exploration"
	JOBSUBTYPE_VISUALIZERS_CALCULATION    JobSubType = "Visualizers Calculation"
	JOBSUBTYPE_SAMPLE_SELECTION           JobSubType = "Sample Selection"
	JOBSUBTYPE_SAMPLE_ANALYSIS            JobSubType = "Sample Analysis"
	JOBSUBTYPE_GRAPH_VALIDATE             JobSubType = "Graph Validate"
	JOBSUBTYPE_FETCH_SIMILAR              JobSubType = "Fetch Similar"
	JOBSUBTYPE_EXPORT_PROJECT             JobSubType = "Export Project"
	JOBSUBTYPE_COPY_PROJECT               JobSubType = "Copy Project"
	JOBSUBTYPE_IMPORT_PROJECT             JobSubType = "Import Project"
)

// All allowed values of JobSubType enum
var AllowedJobSubTypeEnumValues = []JobSubType{
	"Evaluate",
	"Train From Scratch",
	"Continue Training",
	"Train From Initial Weights",
	"Population Exploration",
	"Visualizers Calculation",
	"Sample Selection",
	"Sample Analysis",
	"Graph Validate",
	"Fetch Similar",
	"Export Project",
	"Copy Project",
	"Import Project",
}

func (v *JobSubType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JobSubType(value)
	for _, existing := range AllowedJobSubTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JobSubType", value)
}

// NewJobSubTypeFromValue returns a pointer to a valid JobSubType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJobSubTypeFromValue(v string) (*JobSubType, error) {
	ev := JobSubType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JobSubType: valid values are %v", v, AllowedJobSubTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JobSubType) IsValid() bool {
	for _, existing := range AllowedJobSubTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JobSubType value
func (v JobSubType) Ptr() *JobSubType {
	return &v
}

type NullableJobSubType struct {
	value *JobSubType
	isSet bool
}

func (v NullableJobSubType) Get() *JobSubType {
	return v.value
}

func (v *NullableJobSubType) Set(val *JobSubType) {
	v.value = val
	v.isSet = true
}

func (v NullableJobSubType) IsSet() bool {
	return v.isSet
}

func (v *NullableJobSubType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobSubType(val *JobSubType) *NullableJobSubType {
	return &NullableJobSubType{value: val, isSet: true}
}

func (v NullableJobSubType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobSubType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
