/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.445
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// AggregationMethod the model 'AggregationMethod'
type AggregationMethod string

// List of AggregationMethod
const (
	AGGREGATIONMETHOD_AVERAGE AggregationMethod = "Average"
	AGGREGATIONMETHOD_COUNT   AggregationMethod = "Count"
	AGGREGATIONMETHOD_MIN     AggregationMethod = "Min"
	AGGREGATIONMETHOD_MAX     AggregationMethod = "Max"
)

// All allowed values of AggregationMethod enum
var AllowedAggregationMethodEnumValues = []AggregationMethod{
	"Average",
	"Count",
	"Min",
	"Max",
}

func (v *AggregationMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AggregationMethod(value)
	for _, existing := range AllowedAggregationMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AggregationMethod", value)
}

// NewAggregationMethodFromValue returns a pointer to a valid AggregationMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAggregationMethodFromValue(v string) (*AggregationMethod, error) {
	ev := AggregationMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AggregationMethod: valid values are %v", v, AllowedAggregationMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AggregationMethod) IsValid() bool {
	for _, existing := range AllowedAggregationMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregationMethod value
func (v AggregationMethod) Ptr() *AggregationMethod {
	return &v
}

type NullableAggregationMethod struct {
	value *AggregationMethod
	isSet bool
}

func (v NullableAggregationMethod) Get() *AggregationMethod {
	return v.value
}

func (v *NullableAggregationMethod) Set(val *AggregationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregationMethod(val *AggregationMethod) *NullableAggregationMethod {
	return &NullableAggregationMethod{value: val, isSet: true}
}

func (v NullableAggregationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
