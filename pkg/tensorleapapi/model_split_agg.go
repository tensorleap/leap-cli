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

// SplitAgg struct for SplitAgg
type SplitAgg struct {
	ContinuesAgg *ContinuesAgg
	DistinctAgg  *DistinctAgg
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SplitAgg) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ContinuesAgg
	err = json.Unmarshal(data, &dst.ContinuesAgg)
	if err == nil {
		jsonContinuesAgg, _ := json.Marshal(dst.ContinuesAgg)
		if string(jsonContinuesAgg) == "{}" { // empty struct
			dst.ContinuesAgg = nil
		} else {
			return nil // data stored in dst.ContinuesAgg, return on the first match
		}
	} else {
		dst.ContinuesAgg = nil
	}

	// try to unmarshal JSON data into DistinctAgg
	err = json.Unmarshal(data, &dst.DistinctAgg)
	if err == nil {
		jsonDistinctAgg, _ := json.Marshal(dst.DistinctAgg)
		if string(jsonDistinctAgg) == "{}" { // empty struct
			dst.DistinctAgg = nil
		} else {
			return nil // data stored in dst.DistinctAgg, return on the first match
		}
	} else {
		dst.DistinctAgg = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SplitAgg)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SplitAgg) MarshalJSON() ([]byte, error) {
	if src.ContinuesAgg != nil {
		return json.Marshal(&src.ContinuesAgg)
	}

	if src.DistinctAgg != nil {
		return json.Marshal(&src.DistinctAgg)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSplitAgg struct {
	value *SplitAgg
	isSet bool
}

func (v NullableSplitAgg) Get() *SplitAgg {
	return v.value
}

func (v *NullableSplitAgg) Set(val *SplitAgg) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitAgg) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitAgg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitAgg(val *SplitAgg) *NullableSplitAgg {
	return &NullableSplitAgg{value: val, isSet: true}
}

func (v NullableSplitAgg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitAgg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
