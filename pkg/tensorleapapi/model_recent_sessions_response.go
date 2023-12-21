/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.423
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the RecentSessionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecentSessionsResponse{}

// RecentSessionsResponse struct for RecentSessionsResponse
type RecentSessionsResponse struct {
	Sessions []SessionPopulatedJob `json:"sessions"`
}

// NewRecentSessionsResponse instantiates a new RecentSessionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecentSessionsResponse(sessions []SessionPopulatedJob) *RecentSessionsResponse {
	this := RecentSessionsResponse{}
	this.Sessions = sessions
	return &this
}

// NewRecentSessionsResponseWithDefaults instantiates a new RecentSessionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecentSessionsResponseWithDefaults() *RecentSessionsResponse {
	this := RecentSessionsResponse{}
	return &this
}

// GetSessions returns the Sessions field value
func (o *RecentSessionsResponse) GetSessions() []SessionPopulatedJob {
	if o == nil {
		var ret []SessionPopulatedJob
		return ret
	}

	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *RecentSessionsResponse) GetSessionsOk() ([]SessionPopulatedJob, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sessions, true
}

// SetSessions sets field value
func (o *RecentSessionsResponse) SetSessions(v []SessionPopulatedJob) {
	o.Sessions = v
}

func (o RecentSessionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecentSessionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessions"] = o.Sessions
	return toSerialize, nil
}

type NullableRecentSessionsResponse struct {
	value *RecentSessionsResponse
	isSet bool
}

func (v NullableRecentSessionsResponse) Get() *RecentSessionsResponse {
	return v.value
}

func (v *NullableRecentSessionsResponse) Set(val *RecentSessionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRecentSessionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRecentSessionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecentSessionsResponse(val *RecentSessionsResponse) *NullableRecentSessionsResponse {
	return &NullableRecentSessionsResponse{value: val, isSet: true}
}

func (v NullableRecentSessionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecentSessionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
