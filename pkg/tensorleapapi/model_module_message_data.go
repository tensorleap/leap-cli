/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.417
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// ModuleMessageData struct for ModuleMessageData
type ModuleMessageData struct {
	DatasetMessageParams    *DatasetMessageParams
	GeneralMessageParams    *GeneralMessageParams
	LossMessageParams       *LossMessageParams
	MetricMessageParams     *MetricMessageParams
	NodeMessageParams       *NodeMessageParams
	VisualizerMessageParams *VisualizerMessageParams
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ModuleMessageData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DatasetMessageParams
	err = json.Unmarshal(data, &dst.DatasetMessageParams)
	if err == nil {
		jsonDatasetMessageParams, _ := json.Marshal(dst.DatasetMessageParams)
		if string(jsonDatasetMessageParams) == "{}" { // empty struct
			dst.DatasetMessageParams = nil
		} else {
			return nil // data stored in dst.DatasetMessageParams, return on the first match
		}
	} else {
		dst.DatasetMessageParams = nil
	}

	// try to unmarshal JSON data into GeneralMessageParams
	err = json.Unmarshal(data, &dst.GeneralMessageParams)
	if err == nil {
		jsonGeneralMessageParams, _ := json.Marshal(dst.GeneralMessageParams)
		if string(jsonGeneralMessageParams) == "{}" { // empty struct
			dst.GeneralMessageParams = nil
		} else {
			return nil // data stored in dst.GeneralMessageParams, return on the first match
		}
	} else {
		dst.GeneralMessageParams = nil
	}

	// try to unmarshal JSON data into LossMessageParams
	err = json.Unmarshal(data, &dst.LossMessageParams)
	if err == nil {
		jsonLossMessageParams, _ := json.Marshal(dst.LossMessageParams)
		if string(jsonLossMessageParams) == "{}" { // empty struct
			dst.LossMessageParams = nil
		} else {
			return nil // data stored in dst.LossMessageParams, return on the first match
		}
	} else {
		dst.LossMessageParams = nil
	}

	// try to unmarshal JSON data into MetricMessageParams
	err = json.Unmarshal(data, &dst.MetricMessageParams)
	if err == nil {
		jsonMetricMessageParams, _ := json.Marshal(dst.MetricMessageParams)
		if string(jsonMetricMessageParams) == "{}" { // empty struct
			dst.MetricMessageParams = nil
		} else {
			return nil // data stored in dst.MetricMessageParams, return on the first match
		}
	} else {
		dst.MetricMessageParams = nil
	}

	// try to unmarshal JSON data into NodeMessageParams
	err = json.Unmarshal(data, &dst.NodeMessageParams)
	if err == nil {
		jsonNodeMessageParams, _ := json.Marshal(dst.NodeMessageParams)
		if string(jsonNodeMessageParams) == "{}" { // empty struct
			dst.NodeMessageParams = nil
		} else {
			return nil // data stored in dst.NodeMessageParams, return on the first match
		}
	} else {
		dst.NodeMessageParams = nil
	}

	// try to unmarshal JSON data into VisualizerMessageParams
	err = json.Unmarshal(data, &dst.VisualizerMessageParams)
	if err == nil {
		jsonVisualizerMessageParams, _ := json.Marshal(dst.VisualizerMessageParams)
		if string(jsonVisualizerMessageParams) == "{}" { // empty struct
			dst.VisualizerMessageParams = nil
		} else {
			return nil // data stored in dst.VisualizerMessageParams, return on the first match
		}
	} else {
		dst.VisualizerMessageParams = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ModuleMessageData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ModuleMessageData) MarshalJSON() ([]byte, error) {
	if src.DatasetMessageParams != nil {
		return json.Marshal(&src.DatasetMessageParams)
	}

	if src.GeneralMessageParams != nil {
		return json.Marshal(&src.GeneralMessageParams)
	}

	if src.LossMessageParams != nil {
		return json.Marshal(&src.LossMessageParams)
	}

	if src.MetricMessageParams != nil {
		return json.Marshal(&src.MetricMessageParams)
	}

	if src.NodeMessageParams != nil {
		return json.Marshal(&src.NodeMessageParams)
	}

	if src.VisualizerMessageParams != nil {
		return json.Marshal(&src.VisualizerMessageParams)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableModuleMessageData struct {
	value *ModuleMessageData
	isSet bool
}

func (v NullableModuleMessageData) Get() *ModuleMessageData {
	return v.value
}

func (v *NullableModuleMessageData) Set(val *ModuleMessageData) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleMessageData) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleMessageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleMessageData(val *ModuleMessageData) *NullableModuleMessageData {
	return &NullableModuleMessageData{value: val, isSet: true}
}

func (v NullableModuleMessageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleMessageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
