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

// IssueStatus the model 'IssueStatus'
type IssueStatus string

// List of IssueStatus
const (
	ISSUESTATUS_OPEN    IssueStatus = "Open"
	ISSUESTATUS_AT_WORK IssueStatus = "AtWork"
	ISSUESTATUS_CLOSED  IssueStatus = "Closed"
)

// All allowed values of IssueStatus enum
var AllowedIssueStatusEnumValues = []IssueStatus{
	"Open",
	"AtWork",
	"Closed",
}

func (v *IssueStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IssueStatus(value)
	for _, existing := range AllowedIssueStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IssueStatus", value)
}

// NewIssueStatusFromValue returns a pointer to a valid IssueStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIssueStatusFromValue(v string) (*IssueStatus, error) {
	ev := IssueStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IssueStatus: valid values are %v", v, AllowedIssueStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IssueStatus) IsValid() bool {
	for _, existing := range AllowedIssueStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueStatus value
func (v IssueStatus) Ptr() *IssueStatus {
	return &v
}

type NullableIssueStatus struct {
	value *IssueStatus
	isSet bool
}

func (v NullableIssueStatus) Get() *IssueStatus {
	return v.value
}

func (v *NullableIssueStatus) Set(val *IssueStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueStatus(val *IssueStatus) *NullableIssueStatus {
	return &NullableIssueStatus{value: val, isSet: true}
}

func (v NullableIssueStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
