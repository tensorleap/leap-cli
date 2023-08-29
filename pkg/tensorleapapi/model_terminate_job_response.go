/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.378-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the TerminateJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminateJobResponse{}

// TerminateJobResponse struct for TerminateJobResponse
type TerminateJobResponse struct {
	Status string `json:"status"`
	JobId string `json:"job_id"`
}

// NewTerminateJobResponse instantiates a new TerminateJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminateJobResponse(status string, jobId string) *TerminateJobResponse {
	this := TerminateJobResponse{}
	this.Status = status
	this.JobId = jobId
	return &this
}

// NewTerminateJobResponseWithDefaults instantiates a new TerminateJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminateJobResponseWithDefaults() *TerminateJobResponse {
	this := TerminateJobResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *TerminateJobResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TerminateJobResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TerminateJobResponse) SetStatus(v string) {
	o.Status = v
}

// GetJobId returns the JobId field value
func (o *TerminateJobResponse) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *TerminateJobResponse) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *TerminateJobResponse) SetJobId(v string) {
	o.JobId = v
}

func (o TerminateJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminateJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["job_id"] = o.JobId
	return toSerialize, nil
}

type NullableTerminateJobResponse struct {
	value *TerminateJobResponse
	isSet bool
}

func (v NullableTerminateJobResponse) Get() *TerminateJobResponse {
	return v.value
}

func (v *NullableTerminateJobResponse) Set(val *TerminateJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminateJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminateJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminateJobResponse(val *TerminateJobResponse) *NullableTerminateJobResponse {
	return &NullableTerminateJobResponse{value: val, isSet: true}
}

func (v NullableTerminateJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminateJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


