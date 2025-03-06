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

// EvaluateParams struct for EvaluateParams
type EvaluateParams struct {
	EvaluateExistingSessionParams *EvaluateExistingSessionParams
	EvaluateNewSessionParams      *EvaluateNewSessionParams
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EvaluateParams) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EvaluateExistingSessionParams
	err = json.Unmarshal(data, &dst.EvaluateExistingSessionParams)
	if err == nil {
		jsonEvaluateExistingSessionParams, _ := json.Marshal(dst.EvaluateExistingSessionParams)
		if string(jsonEvaluateExistingSessionParams) == "{}" { // empty struct
			dst.EvaluateExistingSessionParams = nil
		} else {
			return nil // data stored in dst.EvaluateExistingSessionParams, return on the first match
		}
	} else {
		dst.EvaluateExistingSessionParams = nil
	}

	// try to unmarshal JSON data into EvaluateNewSessionParams
	err = json.Unmarshal(data, &dst.EvaluateNewSessionParams)
	if err == nil {
		jsonEvaluateNewSessionParams, _ := json.Marshal(dst.EvaluateNewSessionParams)
		if string(jsonEvaluateNewSessionParams) == "{}" { // empty struct
			dst.EvaluateNewSessionParams = nil
		} else {
			return nil // data stored in dst.EvaluateNewSessionParams, return on the first match
		}
	} else {
		dst.EvaluateNewSessionParams = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EvaluateParams)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EvaluateParams) MarshalJSON() ([]byte, error) {
	if src.EvaluateExistingSessionParams != nil {
		return json.Marshal(&src.EvaluateExistingSessionParams)
	}

	if src.EvaluateNewSessionParams != nil {
		return json.Marshal(&src.EvaluateNewSessionParams)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEvaluateParams struct {
	value *EvaluateParams
	isSet bool
}

func (v NullableEvaluateParams) Get() *EvaluateParams {
	return v.value
}

func (v *NullableEvaluateParams) Set(val *EvaluateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEvaluateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEvaluateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvaluateParams(val *EvaluateParams) *NullableEvaluateParams {
	return &NullableEvaluateParams{value: val, isSet: true}
}

func (v NullableEvaluateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvaluateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
