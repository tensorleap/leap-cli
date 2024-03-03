/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// VisualizedItemData struct for VisualizedItemData
type VisualizedItemData struct {
	GradsAnalysis *GradsAnalysis
	VisData       *VisData
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *VisualizedItemData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into GradsAnalysis
	err = json.Unmarshal(data, &dst.GradsAnalysis)
	if err == nil {
		jsonGradsAnalysis, _ := json.Marshal(dst.GradsAnalysis)
		if string(jsonGradsAnalysis) == "{}" { // empty struct
			dst.GradsAnalysis = nil
		} else {
			return nil // data stored in dst.GradsAnalysis, return on the first match
		}
	} else {
		dst.GradsAnalysis = nil
	}

	// try to unmarshal JSON data into VisData
	err = json.Unmarshal(data, &dst.VisData)
	if err == nil {
		jsonVisData, _ := json.Marshal(dst.VisData)
		if string(jsonVisData) == "{}" { // empty struct
			dst.VisData = nil
		} else {
			return nil // data stored in dst.VisData, return on the first match
		}
	} else {
		dst.VisData = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(VisualizedItemData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *VisualizedItemData) MarshalJSON() ([]byte, error) {
	if src.GradsAnalysis != nil {
		return json.Marshal(&src.GradsAnalysis)
	}

	if src.VisData != nil {
		return json.Marshal(&src.VisData)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableVisualizedItemData struct {
	value *VisualizedItemData
	isSet bool
}

func (v NullableVisualizedItemData) Get() *VisualizedItemData {
	return v.value
}

func (v *NullableVisualizedItemData) Set(val *VisualizedItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualizedItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualizedItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualizedItemData(val *VisualizedItemData) *NullableVisualizedItemData {
	return &NullableVisualizedItemData{value: val, isSet: true}
}

func (v NullableVisualizedItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualizedItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
