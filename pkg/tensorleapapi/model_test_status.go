/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.285
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// TestStatus the model 'TestStatus'
type TestStatus string

// List of TestStatus
const (
	TESTSTATUS_BEFORE_TEST TestStatus = "beforeTest"
	TESTSTATUS_DURING_TEST TestStatus = "duringTest"
	TESTSTATUS_TEST_SUCCESS TestStatus = "testSuccess"
	TESTSTATUS_TEST_FAIL TestStatus = "testFail"
)

// All allowed values of TestStatus enum
var AllowedTestStatusEnumValues = []TestStatus{
	"beforeTest",
	"duringTest",
	"testSuccess",
	"testFail",
}

func (v *TestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TestStatus(value)
	for _, existing := range AllowedTestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TestStatus", value)
}

// NewTestStatusFromValue returns a pointer to a valid TestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTestStatusFromValue(v string) (*TestStatus, error) {
	ev := TestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TestStatus: valid values are %v", v, AllowedTestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TestStatus) IsValid() bool {
	for _, existing := range AllowedTestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestStatus value
func (v TestStatus) Ptr() *TestStatus {
	return &v
}

type NullableTestStatus struct {
	value *TestStatus
	isSet bool
}

func (v NullableTestStatus) Get() *TestStatus {
	return v.value
}

func (v *NullableTestStatus) Set(val *TestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestStatus(val *TestStatus) *NullableTestStatus {
	return &NullableTestStatus{value: val, isSet: true}
}

func (v NullableTestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

