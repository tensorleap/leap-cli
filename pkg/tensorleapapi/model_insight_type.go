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

// InsightType struct for InsightType
type InsightType struct {
	ClusterInsight             *ClusterInsight
	DataLeakageInsight         *DataLeakageInsight
	DuplicationInsight         *DuplicationInsight
	OverfittingInsight         *OverfittingInsight
	UnderRepresentationInsight *UnderRepresentationInsight
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *InsightType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ClusterInsight
	err = json.Unmarshal(data, &dst.ClusterInsight)
	if err == nil {
		jsonClusterInsight, _ := json.Marshal(dst.ClusterInsight)
		if string(jsonClusterInsight) == "{}" { // empty struct
			dst.ClusterInsight = nil
		} else {
			return nil // data stored in dst.ClusterInsight, return on the first match
		}
	} else {
		dst.ClusterInsight = nil
	}

	// try to unmarshal JSON data into DataLeakageInsight
	err = json.Unmarshal(data, &dst.DataLeakageInsight)
	if err == nil {
		jsonDataLeakageInsight, _ := json.Marshal(dst.DataLeakageInsight)
		if string(jsonDataLeakageInsight) == "{}" { // empty struct
			dst.DataLeakageInsight = nil
		} else {
			return nil // data stored in dst.DataLeakageInsight, return on the first match
		}
	} else {
		dst.DataLeakageInsight = nil
	}

	// try to unmarshal JSON data into DuplicationInsight
	err = json.Unmarshal(data, &dst.DuplicationInsight)
	if err == nil {
		jsonDuplicationInsight, _ := json.Marshal(dst.DuplicationInsight)
		if string(jsonDuplicationInsight) == "{}" { // empty struct
			dst.DuplicationInsight = nil
		} else {
			return nil // data stored in dst.DuplicationInsight, return on the first match
		}
	} else {
		dst.DuplicationInsight = nil
	}

	// try to unmarshal JSON data into OverfittingInsight
	err = json.Unmarshal(data, &dst.OverfittingInsight)
	if err == nil {
		jsonOverfittingInsight, _ := json.Marshal(dst.OverfittingInsight)
		if string(jsonOverfittingInsight) == "{}" { // empty struct
			dst.OverfittingInsight = nil
		} else {
			return nil // data stored in dst.OverfittingInsight, return on the first match
		}
	} else {
		dst.OverfittingInsight = nil
	}

	// try to unmarshal JSON data into UnderRepresentationInsight
	err = json.Unmarshal(data, &dst.UnderRepresentationInsight)
	if err == nil {
		jsonUnderRepresentationInsight, _ := json.Marshal(dst.UnderRepresentationInsight)
		if string(jsonUnderRepresentationInsight) == "{}" { // empty struct
			dst.UnderRepresentationInsight = nil
		} else {
			return nil // data stored in dst.UnderRepresentationInsight, return on the first match
		}
	} else {
		dst.UnderRepresentationInsight = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(InsightType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *InsightType) MarshalJSON() ([]byte, error) {
	if src.ClusterInsight != nil {
		return json.Marshal(&src.ClusterInsight)
	}

	if src.DataLeakageInsight != nil {
		return json.Marshal(&src.DataLeakageInsight)
	}

	if src.DuplicationInsight != nil {
		return json.Marshal(&src.DuplicationInsight)
	}

	if src.OverfittingInsight != nil {
		return json.Marshal(&src.OverfittingInsight)
	}

	if src.UnderRepresentationInsight != nil {
		return json.Marshal(&src.UnderRepresentationInsight)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableInsightType struct {
	value *InsightType
	isSet bool
}

func (v NullableInsightType) Get() *InsightType {
	return v.value
}

func (v *NullableInsightType) Set(val *InsightType) {
	v.value = val
	v.isSet = true
}

func (v NullableInsightType) IsSet() bool {
	return v.isSet
}

func (v *NullableInsightType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsightType(val *InsightType) *NullableInsightType {
	return &NullableInsightType{value: val, isSet: true}
}

func (v NullableInsightType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsightType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
