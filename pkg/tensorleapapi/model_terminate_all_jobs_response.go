/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the TerminateAllJobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminateAllJobsResponse{}

// TerminateAllJobsResponse struct for TerminateAllJobsResponse
type TerminateAllJobsResponse struct {
	Status string   `json:"status"`
	JobIds []string `json:"job_ids"`
}

// NewTerminateAllJobsResponse instantiates a new TerminateAllJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminateAllJobsResponse(status string, jobIds []string) *TerminateAllJobsResponse {
	this := TerminateAllJobsResponse{}
	this.Status = status
	this.JobIds = jobIds
	return &this
}

// NewTerminateAllJobsResponseWithDefaults instantiates a new TerminateAllJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminateAllJobsResponseWithDefaults() *TerminateAllJobsResponse {
	this := TerminateAllJobsResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *TerminateAllJobsResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TerminateAllJobsResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TerminateAllJobsResponse) SetStatus(v string) {
	o.Status = v
}

// GetJobIds returns the JobIds field value
func (o *TerminateAllJobsResponse) GetJobIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.JobIds
}

// GetJobIdsOk returns a tuple with the JobIds field value
// and a boolean to check if the value has been set.
func (o *TerminateAllJobsResponse) GetJobIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobIds, true
}

// SetJobIds sets field value
func (o *TerminateAllJobsResponse) SetJobIds(v []string) {
	o.JobIds = v
}

func (o TerminateAllJobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminateAllJobsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["job_ids"] = o.JobIds
	return toSerialize, nil
}

type NullableTerminateAllJobsResponse struct {
	value *TerminateAllJobsResponse
	isSet bool
}

func (v NullableTerminateAllJobsResponse) Get() *TerminateAllJobsResponse {
	return v.value
}

func (v *NullableTerminateAllJobsResponse) Set(val *TerminateAllJobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminateAllJobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminateAllJobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminateAllJobsResponse(val *TerminateAllJobsResponse) *NullableTerminateAllJobsResponse {
	return &NullableTerminateAllJobsResponse{value: val, isSet: true}
}

func (v NullableTerminateAllJobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminateAllJobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
