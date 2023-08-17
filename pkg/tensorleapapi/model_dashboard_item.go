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

// DashboardItem struct for DashboardItem
type DashboardItem struct {
	CustomVisualizationItem *CustomVisualizationItem
	VisualizationItem *VisualizationItem
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DashboardItem) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CustomVisualizationItem
	err = json.Unmarshal(data, &dst.CustomVisualizationItem);
	if err == nil {
		jsonCustomVisualizationItem, _ := json.Marshal(dst.CustomVisualizationItem)
		if string(jsonCustomVisualizationItem) == "{}" { // empty struct
			dst.CustomVisualizationItem = nil
		} else {
			return nil // data stored in dst.CustomVisualizationItem, return on the first match
		}
	} else {
		dst.CustomVisualizationItem = nil
	}

	// try to unmarshal JSON data into VisualizationItem
	err = json.Unmarshal(data, &dst.VisualizationItem);
	if err == nil {
		jsonVisualizationItem, _ := json.Marshal(dst.VisualizationItem)
		if string(jsonVisualizationItem) == "{}" { // empty struct
			dst.VisualizationItem = nil
		} else {
			return nil // data stored in dst.VisualizationItem, return on the first match
		}
	} else {
		dst.VisualizationItem = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(DashboardItem)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DashboardItem) MarshalJSON() ([]byte, error) {
	if src.CustomVisualizationItem != nil {
		return json.Marshal(&src.CustomVisualizationItem)
	}

	if src.VisualizationItem != nil {
		return json.Marshal(&src.VisualizationItem)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDashboardItem struct {
	value *DashboardItem
	isSet bool
}

func (v NullableDashboardItem) Get() *DashboardItem {
	return v.value
}

func (v *NullableDashboardItem) Set(val *DashboardItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardItem(val *DashboardItem) *NullableDashboardItem {
	return &NullableDashboardItem{value: val, isSet: true}
}

func (v NullableDashboardItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


