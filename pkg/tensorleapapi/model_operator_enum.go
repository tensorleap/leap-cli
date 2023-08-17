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

// OperatorEnum the model 'OperatorEnum'
type OperatorEnum string

// List of OperatorEnum
const (
	OPERATORENUM_SMALLER_THAN OperatorEnum = "smaller_than"
	OPERATORENUM_EQUAL_TO OperatorEnum = "equal_to"
	OPERATORENUM_LARGER_THAN OperatorEnum = "larger_than"
)

// All allowed values of OperatorEnum enum
var AllowedOperatorEnumEnumValues = []OperatorEnum{
	"smaller_than",
	"equal_to",
	"larger_than",
}

func (v *OperatorEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperatorEnum(value)
	for _, existing := range AllowedOperatorEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperatorEnum", value)
}

// NewOperatorEnumFromValue returns a pointer to a valid OperatorEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperatorEnumFromValue(v string) (*OperatorEnum, error) {
	ev := OperatorEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperatorEnum: valid values are %v", v, AllowedOperatorEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperatorEnum) IsValid() bool {
	for _, existing := range AllowedOperatorEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperatorEnum value
func (v OperatorEnum) Ptr() *OperatorEnum {
	return &v
}

type NullableOperatorEnum struct {
	value *OperatorEnum
	isSet bool
}

func (v NullableOperatorEnum) Get() *OperatorEnum {
	return v.value
}

func (v *NullableOperatorEnum) Set(val *OperatorEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorEnum(val *OperatorEnum) *NullableOperatorEnum {
	return &NullableOperatorEnum{value: val, isSet: true}
}

func (v NullableOperatorEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

