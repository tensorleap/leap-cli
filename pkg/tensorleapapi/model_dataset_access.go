/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.342
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// DatasetAccess the model 'DatasetAccess'
type DatasetAccess string

// List of DatasetAccess
const (
	DATASETACCESS_PUBLIC DatasetAccess = "public"
	DATASETACCESS_LOCAL DatasetAccess = "local"
)

// All allowed values of DatasetAccess enum
var AllowedDatasetAccessEnumValues = []DatasetAccess{
	"public",
	"local",
}

func (v *DatasetAccess) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DatasetAccess(value)
	for _, existing := range AllowedDatasetAccessEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DatasetAccess", value)
}

// NewDatasetAccessFromValue returns a pointer to a valid DatasetAccess
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDatasetAccessFromValue(v string) (*DatasetAccess, error) {
	ev := DatasetAccess(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DatasetAccess: valid values are %v", v, AllowedDatasetAccessEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DatasetAccess) IsValid() bool {
	for _, existing := range AllowedDatasetAccessEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatasetAccess value
func (v DatasetAccess) Ptr() *DatasetAccess {
	return &v
}

type NullableDatasetAccess struct {
	value *DatasetAccess
	isSet bool
}

func (v NullableDatasetAccess) Get() *DatasetAccess {
	return v.value
}

func (v *NullableDatasetAccess) Set(val *DatasetAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetAccess(val *DatasetAccess) *NullableDatasetAccess {
	return &NullableDatasetAccess{value: val, isSet: true}
}

func (v NullableDatasetAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

