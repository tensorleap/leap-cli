/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// HubPublishPolicy the model 'HubPublishPolicy'
type HubPublishPolicy string

// List of HubPublishPolicy
const (
	HUBPUBLISHPOLICY_PUBLIC     HubPublishPolicy = "public"
	HUBPUBLISHPOLICY_ALPHA      HubPublishPolicy = "alpha"
	HUBPUBLISHPOLICY_NO_PUBLISH HubPublishPolicy = "no_publish"
)

// All allowed values of HubPublishPolicy enum
var AllowedHubPublishPolicyEnumValues = []HubPublishPolicy{
	"public",
	"alpha",
	"no_publish",
}

func (v *HubPublishPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HubPublishPolicy(value)
	for _, existing := range AllowedHubPublishPolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HubPublishPolicy", value)
}

// NewHubPublishPolicyFromValue returns a pointer to a valid HubPublishPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHubPublishPolicyFromValue(v string) (*HubPublishPolicy, error) {
	ev := HubPublishPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HubPublishPolicy: valid values are %v", v, AllowedHubPublishPolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HubPublishPolicy) IsValid() bool {
	for _, existing := range AllowedHubPublishPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HubPublishPolicy value
func (v HubPublishPolicy) Ptr() *HubPublishPolicy {
	return &v
}

type NullableHubPublishPolicy struct {
	value *HubPublishPolicy
	isSet bool
}

func (v NullableHubPublishPolicy) Get() *HubPublishPolicy {
	return v.value
}

func (v *NullableHubPublishPolicy) Set(val *HubPublishPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHubPublishPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHubPublishPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHubPublishPolicy(val *HubPublishPolicy) *NullableHubPublishPolicy {
	return &NullableHubPublishPolicy{value: val, isSet: true}
}

func (v NullableHubPublishPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHubPublishPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
