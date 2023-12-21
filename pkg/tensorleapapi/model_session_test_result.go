/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.423
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// SessionTestResult struct for SessionTestResult
type SessionTestResult struct {
	SessionTestResultError    *SessionTestResultError
	SessionTestResultNotFound *SessionTestResultNotFound
	SessionTestResultSuccess  *SessionTestResultSuccess
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SessionTestResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SessionTestResultError
	err = json.Unmarshal(data, &dst.SessionTestResultError)
	if err == nil {
		jsonSessionTestResultError, _ := json.Marshal(dst.SessionTestResultError)
		if string(jsonSessionTestResultError) == "{}" { // empty struct
			dst.SessionTestResultError = nil
		} else {
			return nil // data stored in dst.SessionTestResultError, return on the first match
		}
	} else {
		dst.SessionTestResultError = nil
	}

	// try to unmarshal JSON data into SessionTestResultNotFound
	err = json.Unmarshal(data, &dst.SessionTestResultNotFound)
	if err == nil {
		jsonSessionTestResultNotFound, _ := json.Marshal(dst.SessionTestResultNotFound)
		if string(jsonSessionTestResultNotFound) == "{}" { // empty struct
			dst.SessionTestResultNotFound = nil
		} else {
			return nil // data stored in dst.SessionTestResultNotFound, return on the first match
		}
	} else {
		dst.SessionTestResultNotFound = nil
	}

	// try to unmarshal JSON data into SessionTestResultSuccess
	err = json.Unmarshal(data, &dst.SessionTestResultSuccess)
	if err == nil {
		jsonSessionTestResultSuccess, _ := json.Marshal(dst.SessionTestResultSuccess)
		if string(jsonSessionTestResultSuccess) == "{}" { // empty struct
			dst.SessionTestResultSuccess = nil
		} else {
			return nil // data stored in dst.SessionTestResultSuccess, return on the first match
		}
	} else {
		dst.SessionTestResultSuccess = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SessionTestResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SessionTestResult) MarshalJSON() ([]byte, error) {
	if src.SessionTestResultError != nil {
		return json.Marshal(&src.SessionTestResultError)
	}

	if src.SessionTestResultNotFound != nil {
		return json.Marshal(&src.SessionTestResultNotFound)
	}

	if src.SessionTestResultSuccess != nil {
		return json.Marshal(&src.SessionTestResultSuccess)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSessionTestResult struct {
	value *SessionTestResult
	isSet bool
}

func (v NullableSessionTestResult) Get() *SessionTestResult {
	return v.value
}

func (v *NullableSessionTestResult) Set(val *SessionTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionTestResult(val *SessionTestResult) *NullableSessionTestResult {
	return &NullableSessionTestResult{value: val, isSet: true}
}

func (v NullableSessionTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
