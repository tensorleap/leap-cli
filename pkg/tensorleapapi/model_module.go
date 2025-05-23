/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// Module the model 'Module'
type Module string

// List of Module
const (
	MODULE_GENERAL      Module = "General"
	MODULE_NODE         Module = "Node"
	MODULE_DATASET      Module = "Dataset"
	MODULE_VISUALIZERS  Module = "Visualizers"
	MODULE_LOSS         Module = "Loss"
	MODULE_METRIC       Module = "Metric"
	MODULE_IMPORT_MODEL Module = "ImportModel"
)

// All allowed values of Module enum
var AllowedModuleEnumValues = []Module{
	"General",
	"Node",
	"Dataset",
	"Visualizers",
	"Loss",
	"Metric",
	"ImportModel",
}

func (v *Module) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Module(value)
	for _, existing := range AllowedModuleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Module", value)
}

// NewModuleFromValue returns a pointer to a valid Module
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModuleFromValue(v string) (*Module, error) {
	ev := Module(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Module: valid values are %v", v, AllowedModuleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Module) IsValid() bool {
	for _, existing := range AllowedModuleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Module value
func (v Module) Ptr() *Module {
	return &v
}

type NullableModule struct {
	value *Module
	isSet bool
}

func (v NullableModule) Get() *Module {
	return v.value
}

func (v *NullableModule) Set(val *Module) {
	v.value = val
	v.isSet = true
}

func (v NullableModule) IsSet() bool {
	return v.isSet
}

func (v *NullableModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModule(val *Module) *NullableModule {
	return &NullableModule{value: val, isSet: true}
}

func (v NullableModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
