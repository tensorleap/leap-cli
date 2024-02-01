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

// MessageLevel the model 'MessageLevel'
type MessageLevel string

// List of MessageLevel
const (
	MESSAGELEVEL_ERROR   MessageLevel = "Error"
	MESSAGELEVEL_WARNING MessageLevel = "Warning"
	MESSAGELEVEL_INFO    MessageLevel = "Info"
	MESSAGELEVEL_VERBOSE MessageLevel = "Verbose"
)

// All allowed values of MessageLevel enum
var AllowedMessageLevelEnumValues = []MessageLevel{
	"Error",
	"Warning",
	"Info",
	"Verbose",
}

func (v *MessageLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageLevel(value)
	for _, existing := range AllowedMessageLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MessageLevel", value)
}

// NewMessageLevelFromValue returns a pointer to a valid MessageLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageLevelFromValue(v string) (*MessageLevel, error) {
	ev := MessageLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageLevel: valid values are %v", v, AllowedMessageLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageLevel) IsValid() bool {
	for _, existing := range AllowedMessageLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageLevel value
func (v MessageLevel) Ptr() *MessageLevel {
	return &v
}

type NullableMessageLevel struct {
	value *MessageLevel
	isSet bool
}

func (v NullableMessageLevel) Get() *MessageLevel {
	return v.value
}

func (v *NullableMessageLevel) Set(val *MessageLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageLevel(val *MessageLevel) *NullableMessageLevel {
	return &NullableMessageLevel{value: val, isSet: true}
}

func (v NullableMessageLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
