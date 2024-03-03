/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// FilterDisplayData struct for FilterDisplayData
type FilterDisplayData struct {
	FetchSimilarFilterDisplayData *FetchSimilarFilterDisplayData
	InsightFilterDisplayData      *InsightFilterDisplayData
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FilterDisplayData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FetchSimilarFilterDisplayData
	err = json.Unmarshal(data, &dst.FetchSimilarFilterDisplayData)
	if err == nil {
		jsonFetchSimilarFilterDisplayData, _ := json.Marshal(dst.FetchSimilarFilterDisplayData)
		if string(jsonFetchSimilarFilterDisplayData) == "{}" { // empty struct
			dst.FetchSimilarFilterDisplayData = nil
		} else {
			return nil // data stored in dst.FetchSimilarFilterDisplayData, return on the first match
		}
	} else {
		dst.FetchSimilarFilterDisplayData = nil
	}

	// try to unmarshal JSON data into InsightFilterDisplayData
	err = json.Unmarshal(data, &dst.InsightFilterDisplayData)
	if err == nil {
		jsonInsightFilterDisplayData, _ := json.Marshal(dst.InsightFilterDisplayData)
		if string(jsonInsightFilterDisplayData) == "{}" { // empty struct
			dst.InsightFilterDisplayData = nil
		} else {
			return nil // data stored in dst.InsightFilterDisplayData, return on the first match
		}
	} else {
		dst.InsightFilterDisplayData = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(FilterDisplayData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FilterDisplayData) MarshalJSON() ([]byte, error) {
	if src.FetchSimilarFilterDisplayData != nil {
		return json.Marshal(&src.FetchSimilarFilterDisplayData)
	}

	if src.InsightFilterDisplayData != nil {
		return json.Marshal(&src.InsightFilterDisplayData)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFilterDisplayData struct {
	value *FilterDisplayData
	isSet bool
}

func (v NullableFilterDisplayData) Get() *FilterDisplayData {
	return v.value
}

func (v *NullableFilterDisplayData) Set(val *FilterDisplayData) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterDisplayData) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterDisplayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterDisplayData(val *FilterDisplayData) *NullableFilterDisplayData {
	return &NullableFilterDisplayData{value: val, isSet: true}
}

func (v NullableFilterDisplayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterDisplayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
