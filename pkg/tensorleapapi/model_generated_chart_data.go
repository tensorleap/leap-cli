/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// GeneratedChartData struct for GeneratedChartData
type GeneratedChartData struct {
	HeatmapChart *HeatmapChart
	XYChart      *XYChart
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GeneratedChartData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into HeatmapChart
	err = json.Unmarshal(data, &dst.HeatmapChart)
	if err == nil {
		jsonHeatmapChart, _ := json.Marshal(dst.HeatmapChart)
		if string(jsonHeatmapChart) == "{}" { // empty struct
			dst.HeatmapChart = nil
		} else {
			return nil // data stored in dst.HeatmapChart, return on the first match
		}
	} else {
		dst.HeatmapChart = nil
	}

	// try to unmarshal JSON data into XYChart
	err = json.Unmarshal(data, &dst.XYChart)
	if err == nil {
		jsonXYChart, _ := json.Marshal(dst.XYChart)
		if string(jsonXYChart) == "{}" { // empty struct
			dst.XYChart = nil
		} else {
			return nil // data stored in dst.XYChart, return on the first match
		}
	} else {
		dst.XYChart = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GeneratedChartData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GeneratedChartData) MarshalJSON() ([]byte, error) {
	if src.HeatmapChart != nil {
		return json.Marshal(&src.HeatmapChart)
	}

	if src.XYChart != nil {
		return json.Marshal(&src.XYChart)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGeneratedChartData struct {
	value *GeneratedChartData
	isSet bool
}

func (v NullableGeneratedChartData) Get() *GeneratedChartData {
	return v.value
}

func (v *NullableGeneratedChartData) Set(val *GeneratedChartData) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneratedChartData) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneratedChartData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneratedChartData(val *GeneratedChartData) *NullableGeneratedChartData {
	return &NullableGeneratedChartData{value: val, isSet: true}
}

func (v NullableGeneratedChartData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneratedChartData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
