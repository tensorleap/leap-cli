/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// DataTypeEnum the model 'DataTypeEnum'
type DataTypeEnum string

// List of DataTypeEnum
const (
	DATATYPEENUM_IMAGE DataTypeEnum = "image"
	DATATYPEENUM_WORD DataTypeEnum = "word"
	DATATYPEENUM_TEXT DataTypeEnum = "text"
	DATATYPEENUM_MASK_TEXT DataTypeEnum = "mask_text"
	DATATYPEENUM_GRAPH DataTypeEnum = "graph"
	DATATYPEENUM_HBAR DataTypeEnum = "hbar"
	DATATYPEENUM_BBOX_IMAGE DataTypeEnum = "bbox_image"
	DATATYPEENUM_MASK_IMAGE DataTypeEnum = "mask_image"
)

// All allowed values of DataTypeEnum enum
var AllowedDataTypeEnumEnumValues = []DataTypeEnum{
	"image",
	"word",
	"text",
	"mask_text",
	"graph",
	"hbar",
	"bbox_image",
	"mask_image",
}

func (v *DataTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataTypeEnum(value)
	for _, existing := range AllowedDataTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataTypeEnum", value)
}

// NewDataTypeEnumFromValue returns a pointer to a valid DataTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataTypeEnumFromValue(v string) (*DataTypeEnum, error) {
	ev := DataTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataTypeEnum: valid values are %v", v, AllowedDataTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataTypeEnum) IsValid() bool {
	for _, existing := range AllowedDataTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataTypeEnum value
func (v DataTypeEnum) Ptr() *DataTypeEnum {
	return &v
}

type NullableDataTypeEnum struct {
	value *DataTypeEnum
	isSet bool
}

func (v NullableDataTypeEnum) Get() *DataTypeEnum {
	return v.value
}

func (v *NullableDataTypeEnum) Set(val *DataTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableDataTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableDataTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataTypeEnum(val *DataTypeEnum) *NullableDataTypeEnum {
	return &NullableDataTypeEnum{value: val, isSet: true}
}

func (v NullableDataTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

