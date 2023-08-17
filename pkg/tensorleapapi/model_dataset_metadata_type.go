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

// DatasetMetadataType the model 'DatasetMetadataType'
type DatasetMetadataType string

// List of DatasetMetadataType
const (
	DATASETMETADATATYPE_FLOAT DatasetMetadataType = "float"
	DATASETMETADATATYPE_STRING DatasetMetadataType = "string"
	DATASETMETADATATYPE_INT DatasetMetadataType = "int"
	DATASETMETADATATYPE_BOOLEAN DatasetMetadataType = "boolean"
)

// All allowed values of DatasetMetadataType enum
var AllowedDatasetMetadataTypeEnumValues = []DatasetMetadataType{
	"float",
	"string",
	"int",
	"boolean",
}

func (v *DatasetMetadataType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DatasetMetadataType(value)
	for _, existing := range AllowedDatasetMetadataTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DatasetMetadataType", value)
}

// NewDatasetMetadataTypeFromValue returns a pointer to a valid DatasetMetadataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDatasetMetadataTypeFromValue(v string) (*DatasetMetadataType, error) {
	ev := DatasetMetadataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DatasetMetadataType: valid values are %v", v, AllowedDatasetMetadataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DatasetMetadataType) IsValid() bool {
	for _, existing := range AllowedDatasetMetadataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatasetMetadataType value
func (v DatasetMetadataType) Ptr() *DatasetMetadataType {
	return &v
}

type NullableDatasetMetadataType struct {
	value *DatasetMetadataType
	isSet bool
}

func (v NullableDatasetMetadataType) Get() *DatasetMetadataType {
	return v.value
}

func (v *NullableDatasetMetadataType) Set(val *DatasetMetadataType) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetMetadataType) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetMetadataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetMetadataType(val *DatasetMetadataType) *NullableDatasetMetadataType {
	return &NullableDatasetMetadataType{value: val, isSet: true}
}

func (v NullableDatasetMetadataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetMetadataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

