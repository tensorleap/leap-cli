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

// LeapDataType the model 'LeapDataType'
type LeapDataType string

// List of LeapDataType
const (
	LEAPDATATYPE_IMAGE              LeapDataType = "Image"
	LEAPDATATYPE_VIDEO              LeapDataType = "Video"
	LEAPDATATYPE_TEXT               LeapDataType = "Text"
	LEAPDATATYPE_GRAPH              LeapDataType = "Graph"
	LEAPDATATYPE_HORIZONTAL_BAR     LeapDataType = "HorizontalBar"
	LEAPDATATYPE_IMAGE_MASK         LeapDataType = "ImageMask"
	LEAPDATATYPE_TEXT_MASK          LeapDataType = "TextMask"
	LEAPDATATYPE_IMAGE_WITH_B_BOX   LeapDataType = "ImageWithBBox"
	LEAPDATATYPE_IMAGE_WITH_HEATMAP LeapDataType = "ImageWithHeatmap"
)

// All allowed values of LeapDataType enum
var AllowedLeapDataTypeEnumValues = []LeapDataType{
	"Image",
	"Video",
	"Text",
	"Graph",
	"HorizontalBar",
	"ImageMask",
	"TextMask",
	"ImageWithBBox",
	"ImageWithHeatmap",
}

func (v *LeapDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LeapDataType(value)
	for _, existing := range AllowedLeapDataTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LeapDataType", value)
}

// NewLeapDataTypeFromValue returns a pointer to a valid LeapDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLeapDataTypeFromValue(v string) (*LeapDataType, error) {
	ev := LeapDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LeapDataType: valid values are %v", v, AllowedLeapDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LeapDataType) IsValid() bool {
	for _, existing := range AllowedLeapDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LeapDataType value
func (v LeapDataType) Ptr() *LeapDataType {
	return &v
}

type NullableLeapDataType struct {
	value *LeapDataType
	isSet bool
}

func (v NullableLeapDataType) Get() *LeapDataType {
	return v.value
}

func (v *NullableLeapDataType) Set(val *LeapDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableLeapDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableLeapDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeapDataType(val *LeapDataType) *NullableLeapDataType {
	return &NullableLeapDataType{value: val, isSet: true}
}

func (v NullableLeapDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLeapDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
