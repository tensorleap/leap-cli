/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.445
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// JobParams struct for JobParams
type JobParams struct {
	AnalyzeParams                  *AnalyzeParams
	ExportModelParams              *ExportModelParams
	FetchSimilarJobParams          *FetchSimilarJobParams
	PopulationExplorationJobParams *PopulationExplorationJobParams
	TrainingParams                 *TrainingParams
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *JobParams) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AnalyzeParams
	err = json.Unmarshal(data, &dst.AnalyzeParams)
	if err == nil {
		jsonAnalyzeParams, _ := json.Marshal(dst.AnalyzeParams)
		if string(jsonAnalyzeParams) == "{}" { // empty struct
			dst.AnalyzeParams = nil
		} else {
			return nil // data stored in dst.AnalyzeParams, return on the first match
		}
	} else {
		dst.AnalyzeParams = nil
	}

	// try to unmarshal JSON data into ExportModelParams
	err = json.Unmarshal(data, &dst.ExportModelParams)
	if err == nil {
		jsonExportModelParams, _ := json.Marshal(dst.ExportModelParams)
		if string(jsonExportModelParams) == "{}" { // empty struct
			dst.ExportModelParams = nil
		} else {
			return nil // data stored in dst.ExportModelParams, return on the first match
		}
	} else {
		dst.ExportModelParams = nil
	}

	// try to unmarshal JSON data into FetchSimilarJobParams
	err = json.Unmarshal(data, &dst.FetchSimilarJobParams)
	if err == nil {
		jsonFetchSimilarJobParams, _ := json.Marshal(dst.FetchSimilarJobParams)
		if string(jsonFetchSimilarJobParams) == "{}" { // empty struct
			dst.FetchSimilarJobParams = nil
		} else {
			return nil // data stored in dst.FetchSimilarJobParams, return on the first match
		}
	} else {
		dst.FetchSimilarJobParams = nil
	}

	// try to unmarshal JSON data into PopulationExplorationJobParams
	err = json.Unmarshal(data, &dst.PopulationExplorationJobParams)
	if err == nil {
		jsonPopulationExplorationJobParams, _ := json.Marshal(dst.PopulationExplorationJobParams)
		if string(jsonPopulationExplorationJobParams) == "{}" { // empty struct
			dst.PopulationExplorationJobParams = nil
		} else {
			return nil // data stored in dst.PopulationExplorationJobParams, return on the first match
		}
	} else {
		dst.PopulationExplorationJobParams = nil
	}

	// try to unmarshal JSON data into TrainingParams
	err = json.Unmarshal(data, &dst.TrainingParams)
	if err == nil {
		jsonTrainingParams, _ := json.Marshal(dst.TrainingParams)
		if string(jsonTrainingParams) == "{}" { // empty struct
			dst.TrainingParams = nil
		} else {
			return nil // data stored in dst.TrainingParams, return on the first match
		}
	} else {
		dst.TrainingParams = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(JobParams)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *JobParams) MarshalJSON() ([]byte, error) {
	if src.AnalyzeParams != nil {
		return json.Marshal(&src.AnalyzeParams)
	}

	if src.ExportModelParams != nil {
		return json.Marshal(&src.ExportModelParams)
	}

	if src.FetchSimilarJobParams != nil {
		return json.Marshal(&src.FetchSimilarJobParams)
	}

	if src.PopulationExplorationJobParams != nil {
		return json.Marshal(&src.PopulationExplorationJobParams)
	}

	if src.TrainingParams != nil {
		return json.Marshal(&src.TrainingParams)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableJobParams struct {
	value *JobParams
	isSet bool
}

func (v NullableJobParams) Get() *JobParams {
	return v.value
}

func (v *NullableJobParams) Set(val *JobParams) {
	v.value = val
	v.isSet = true
}

func (v NullableJobParams) IsSet() bool {
	return v.isSet
}

func (v *NullableJobParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobParams(val *JobParams) *NullableJobParams {
	return &NullableJobParams{value: val, isSet: true}
}

func (v NullableJobParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
