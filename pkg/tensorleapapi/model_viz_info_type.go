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

// VizInfoType struct for VizInfoType
type VizInfoType struct {
	BasicVisualizationInfo  *BasicVisualizationInfo
	PopExploreInfo          *PopExploreInfo
	SampleSelectionInfo     *SampleSelectionInfo
	SampleVisualizationInfo *SampleVisualizationInfo
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *VizInfoType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BasicVisualizationInfo
	err = json.Unmarshal(data, &dst.BasicVisualizationInfo)
	if err == nil {
		jsonBasicVisualizationInfo, _ := json.Marshal(dst.BasicVisualizationInfo)
		if string(jsonBasicVisualizationInfo) == "{}" { // empty struct
			dst.BasicVisualizationInfo = nil
		} else {
			return nil // data stored in dst.BasicVisualizationInfo, return on the first match
		}
	} else {
		dst.BasicVisualizationInfo = nil
	}

	// try to unmarshal JSON data into PopExploreInfo
	err = json.Unmarshal(data, &dst.PopExploreInfo)
	if err == nil {
		jsonPopExploreInfo, _ := json.Marshal(dst.PopExploreInfo)
		if string(jsonPopExploreInfo) == "{}" { // empty struct
			dst.PopExploreInfo = nil
		} else {
			return nil // data stored in dst.PopExploreInfo, return on the first match
		}
	} else {
		dst.PopExploreInfo = nil
	}

	// try to unmarshal JSON data into SampleSelectionInfo
	err = json.Unmarshal(data, &dst.SampleSelectionInfo)
	if err == nil {
		jsonSampleSelectionInfo, _ := json.Marshal(dst.SampleSelectionInfo)
		if string(jsonSampleSelectionInfo) == "{}" { // empty struct
			dst.SampleSelectionInfo = nil
		} else {
			return nil // data stored in dst.SampleSelectionInfo, return on the first match
		}
	} else {
		dst.SampleSelectionInfo = nil
	}

	// try to unmarshal JSON data into SampleVisualizationInfo
	err = json.Unmarshal(data, &dst.SampleVisualizationInfo)
	if err == nil {
		jsonSampleVisualizationInfo, _ := json.Marshal(dst.SampleVisualizationInfo)
		if string(jsonSampleVisualizationInfo) == "{}" { // empty struct
			dst.SampleVisualizationInfo = nil
		} else {
			return nil // data stored in dst.SampleVisualizationInfo, return on the first match
		}
	} else {
		dst.SampleVisualizationInfo = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(VizInfoType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *VizInfoType) MarshalJSON() ([]byte, error) {
	if src.BasicVisualizationInfo != nil {
		return json.Marshal(&src.BasicVisualizationInfo)
	}

	if src.PopExploreInfo != nil {
		return json.Marshal(&src.PopExploreInfo)
	}

	if src.SampleSelectionInfo != nil {
		return json.Marshal(&src.SampleSelectionInfo)
	}

	if src.SampleVisualizationInfo != nil {
		return json.Marshal(&src.SampleVisualizationInfo)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableVizInfoType struct {
	value *VizInfoType
	isSet bool
}

func (v NullableVizInfoType) Get() *VizInfoType {
	return v.value
}

func (v *NullableVizInfoType) Set(val *VizInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableVizInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableVizInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVizInfoType(val *VizInfoType) *NullableVizInfoType {
	return &NullableVizInfoType{value: val, isSet: true}
}

func (v NullableVizInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVizInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
