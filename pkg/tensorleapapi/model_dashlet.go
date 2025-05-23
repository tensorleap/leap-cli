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

// Dashlet struct for Dashlet
type Dashlet struct {
	AnalyticsDashlet             *AnalyticsDashlet
	PopulationExplorationDashlet *PopulationExplorationDashlet
	SampleAnalysisDashlet        *SampleAnalysisDashlet
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Dashlet) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AnalyticsDashlet
	err = json.Unmarshal(data, &dst.AnalyticsDashlet)
	if err == nil {
		jsonAnalyticsDashlet, _ := json.Marshal(dst.AnalyticsDashlet)
		if string(jsonAnalyticsDashlet) == "{}" { // empty struct
			dst.AnalyticsDashlet = nil
		} else {
			return nil // data stored in dst.AnalyticsDashlet, return on the first match
		}
	} else {
		dst.AnalyticsDashlet = nil
	}

	// try to unmarshal JSON data into PopulationExplorationDashlet
	err = json.Unmarshal(data, &dst.PopulationExplorationDashlet)
	if err == nil {
		jsonPopulationExplorationDashlet, _ := json.Marshal(dst.PopulationExplorationDashlet)
		if string(jsonPopulationExplorationDashlet) == "{}" { // empty struct
			dst.PopulationExplorationDashlet = nil
		} else {
			return nil // data stored in dst.PopulationExplorationDashlet, return on the first match
		}
	} else {
		dst.PopulationExplorationDashlet = nil
	}

	// try to unmarshal JSON data into SampleAnalysisDashlet
	err = json.Unmarshal(data, &dst.SampleAnalysisDashlet)
	if err == nil {
		jsonSampleAnalysisDashlet, _ := json.Marshal(dst.SampleAnalysisDashlet)
		if string(jsonSampleAnalysisDashlet) == "{}" { // empty struct
			dst.SampleAnalysisDashlet = nil
		} else {
			return nil // data stored in dst.SampleAnalysisDashlet, return on the first match
		}
	} else {
		dst.SampleAnalysisDashlet = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Dashlet)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Dashlet) MarshalJSON() ([]byte, error) {
	if src.AnalyticsDashlet != nil {
		return json.Marshal(&src.AnalyticsDashlet)
	}

	if src.PopulationExplorationDashlet != nil {
		return json.Marshal(&src.PopulationExplorationDashlet)
	}

	if src.SampleAnalysisDashlet != nil {
		return json.Marshal(&src.SampleAnalysisDashlet)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDashlet struct {
	value *Dashlet
	isSet bool
}

func (v NullableDashlet) Get() *Dashlet {
	return v.value
}

func (v *NullableDashlet) Set(val *Dashlet) {
	v.value = val
	v.isSet = true
}

func (v NullableDashlet) IsSet() bool {
	return v.isSet
}

func (v *NullableDashlet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashlet(val *Dashlet) *NullableDashlet {
	return &NullableDashlet{value: val, isSet: true}
}

func (v NullableDashlet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashlet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
