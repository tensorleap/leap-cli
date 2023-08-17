/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// FilterOperatorType the model 'FilterOperatorType'
type FilterOperatorType string

// List of FilterOperatorType
const (
	FILTEROPERATORTYPE_BETWEEN FilterOperatorType = "between"
	FILTEROPERATORTYPE_NOT_BETWEEN FilterOperatorType = "not-between"
	FILTEROPERATORTYPE_EQUAL FilterOperatorType = "equal"
	FILTEROPERATORTYPE_NOT_EQUAL FilterOperatorType = "not-equal"
	FILTEROPERATORTYPE_IN FilterOperatorType = "in"
	FILTEROPERATORTYPE_GREATER_THAN FilterOperatorType = "greater-than"
	FILTEROPERATORTYPE_LESS_THAN FilterOperatorType = "less-than"
)

// All allowed values of FilterOperatorType enum
var AllowedFilterOperatorTypeEnumValues = []FilterOperatorType{
	"between",
	"not-between",
	"equal",
	"not-equal",
	"in",
	"greater-than",
	"less-than",
}

func (v *FilterOperatorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FilterOperatorType(value)
	for _, existing := range AllowedFilterOperatorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FilterOperatorType", value)
}

// NewFilterOperatorTypeFromValue returns a pointer to a valid FilterOperatorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilterOperatorTypeFromValue(v string) (*FilterOperatorType, error) {
	ev := FilterOperatorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilterOperatorType: valid values are %v", v, AllowedFilterOperatorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilterOperatorType) IsValid() bool {
	for _, existing := range AllowedFilterOperatorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FilterOperatorType value
func (v FilterOperatorType) Ptr() *FilterOperatorType {
	return &v
}

type NullableFilterOperatorType struct {
	value *FilterOperatorType
	isSet bool
}

func (v NullableFilterOperatorType) Get() *FilterOperatorType {
	return v.value
}

func (v *NullableFilterOperatorType) Set(val *FilterOperatorType) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterOperatorType) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterOperatorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterOperatorType(val *FilterOperatorType) *NullableFilterOperatorType {
	return &NullableFilterOperatorType{value: val, isSet: true}
}

func (v NullableFilterOperatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterOperatorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

