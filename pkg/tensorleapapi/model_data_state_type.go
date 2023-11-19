/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.404
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// DataStateType the model 'DataStateType'
type DataStateType string

// List of DataStateType
const (
	DATASTATETYPE_TRAINING   DataStateType = "training"
	DATASTATETYPE_VALIDATION DataStateType = "validation"
	DATASTATETYPE_TEST       DataStateType = "test"
	DATASTATETYPE_UNLABELED  DataStateType = "unlabeled"
)

// All allowed values of DataStateType enum
var AllowedDataStateTypeEnumValues = []DataStateType{
	"training",
	"validation",
	"test",
	"unlabeled",
}

func (v *DataStateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataStateType(value)
	for _, existing := range AllowedDataStateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataStateType", value)
}

// NewDataStateTypeFromValue returns a pointer to a valid DataStateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataStateTypeFromValue(v string) (*DataStateType, error) {
	ev := DataStateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataStateType: valid values are %v", v, AllowedDataStateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataStateType) IsValid() bool {
	for _, existing := range AllowedDataStateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataStateType value
func (v DataStateType) Ptr() *DataStateType {
	return &v
}

type NullableDataStateType struct {
	value *DataStateType
	isSet bool
}

func (v NullableDataStateType) Get() *DataStateType {
	return v.value
}

func (v *NullableDataStateType) Set(val *DataStateType) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStateType) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStateType(val *DataStateType) *NullableDataStateType {
	return &NullableDataStateType{value: val, isSet: true}
}

func (v NullableDataStateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
