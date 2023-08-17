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

// Direction the model 'Direction'
type Direction string

// List of Direction
const (
	DIRECTION_UP   Direction = "up"
	DIRECTION_DOWN Direction = "down"
)

// All allowed values of Direction enum
var AllowedDirectionEnumValues = []Direction{
	"up",
	"down",
}

func (v *Direction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Direction(value)
	for _, existing := range AllowedDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Direction", value)
}

// NewDirectionFromValue returns a pointer to a valid Direction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDirectionFromValue(v string) (*Direction, error) {
	ev := Direction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Direction: valid values are %v", v, AllowedDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Direction) IsValid() bool {
	for _, existing := range AllowedDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Direction value
func (v Direction) Ptr() *Direction {
	return &v
}

type NullableDirection struct {
	value *Direction
	isSet bool
}

func (v NullableDirection) Get() *Direction {
	return v.value
}

func (v *NullableDirection) Set(val *Direction) {
	v.value = val
	v.isSet = true
}

func (v NullableDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirection(val *Direction) *NullableDirection {
	return &NullableDirection{value: val, isSet: true}
}

func (v NullableDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
