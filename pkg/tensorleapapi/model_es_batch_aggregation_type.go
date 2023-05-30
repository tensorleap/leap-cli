/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.321
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// EsBatchAggregationType the model 'EsBatchAggregationType'
type EsBatchAggregationType string

// List of EsBatchAggregationType
const (
	ESBATCHAGGREGATIONTYPE_HISTOGRAM EsBatchAggregationType = "histogram"
	ESBATCHAGGREGATIONTYPE_BATCH     EsBatchAggregationType = "batch"
)

// All allowed values of EsBatchAggregationType enum
var AllowedEsBatchAggregationTypeEnumValues = []EsBatchAggregationType{
	"histogram",
	"batch",
}

func (v *EsBatchAggregationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EsBatchAggregationType(value)
	for _, existing := range AllowedEsBatchAggregationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EsBatchAggregationType", value)
}

// NewEsBatchAggregationTypeFromValue returns a pointer to a valid EsBatchAggregationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEsBatchAggregationTypeFromValue(v string) (*EsBatchAggregationType, error) {
	ev := EsBatchAggregationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EsBatchAggregationType: valid values are %v", v, AllowedEsBatchAggregationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EsBatchAggregationType) IsValid() bool {
	for _, existing := range AllowedEsBatchAggregationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EsBatchAggregationType value
func (v EsBatchAggregationType) Ptr() *EsBatchAggregationType {
	return &v
}

type NullableEsBatchAggregationType struct {
	value *EsBatchAggregationType
	isSet bool
}

func (v NullableEsBatchAggregationType) Get() *EsBatchAggregationType {
	return v.value
}

func (v *NullableEsBatchAggregationType) Set(val *EsBatchAggregationType) {
	v.value = val
	v.isSet = true
}

func (v NullableEsBatchAggregationType) IsSet() bool {
	return v.isSet
}

func (v *NullableEsBatchAggregationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsBatchAggregationType(val *EsBatchAggregationType) *NullableEsBatchAggregationType {
	return &NullableEsBatchAggregationType{value: val, isSet: true}
}

func (v NullableEsBatchAggregationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsBatchAggregationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
