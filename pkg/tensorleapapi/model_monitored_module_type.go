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

// MonitoredModuleType the model 'MonitoredModuleType'
type MonitoredModuleType string

// List of MonitoredModuleType
const (
	MONITOREDMODULETYPE_MONGO    MonitoredModuleType = "mongo"
	MONITOREDMODULETYPE_RABBITMQ MonitoredModuleType = "rabbitmq"
)

// All allowed values of MonitoredModuleType enum
var AllowedMonitoredModuleTypeEnumValues = []MonitoredModuleType{
	"mongo",
	"rabbitmq",
}

func (v *MonitoredModuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MonitoredModuleType(value)
	for _, existing := range AllowedMonitoredModuleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MonitoredModuleType", value)
}

// NewMonitoredModuleTypeFromValue returns a pointer to a valid MonitoredModuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMonitoredModuleTypeFromValue(v string) (*MonitoredModuleType, error) {
	ev := MonitoredModuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MonitoredModuleType: valid values are %v", v, AllowedMonitoredModuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MonitoredModuleType) IsValid() bool {
	for _, existing := range AllowedMonitoredModuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitoredModuleType value
func (v MonitoredModuleType) Ptr() *MonitoredModuleType {
	return &v
}

type NullableMonitoredModuleType struct {
	value *MonitoredModuleType
	isSet bool
}

func (v NullableMonitoredModuleType) Get() *MonitoredModuleType {
	return v.value
}

func (v *NullableMonitoredModuleType) Set(val *MonitoredModuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoredModuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoredModuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoredModuleType(val *MonitoredModuleType) *NullableMonitoredModuleType {
	return &NullableMonitoredModuleType{value: val, isSet: true}
}

func (v NullableMonitoredModuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoredModuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
