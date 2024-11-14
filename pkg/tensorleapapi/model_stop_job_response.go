/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.557
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the StopJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopJobResponse{}

// StopJobResponse struct for StopJobResponse
type StopJobResponse struct {
	Status string `json:"status"`
	JobId  string `json:"job_id"`
}

// NewStopJobResponse instantiates a new StopJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopJobResponse(status string, jobId string) *StopJobResponse {
	this := StopJobResponse{}
	this.Status = status
	this.JobId = jobId
	return &this
}

// NewStopJobResponseWithDefaults instantiates a new StopJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopJobResponseWithDefaults() *StopJobResponse {
	this := StopJobResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *StopJobResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StopJobResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StopJobResponse) SetStatus(v string) {
	o.Status = v
}

// GetJobId returns the JobId field value
func (o *StopJobResponse) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *StopJobResponse) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *StopJobResponse) SetJobId(v string) {
	o.JobId = v
}

func (o StopJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["job_id"] = o.JobId
	return toSerialize, nil
}

type NullableStopJobResponse struct {
	value *StopJobResponse
	isSet bool
}

func (v NullableStopJobResponse) Get() *StopJobResponse {
	return v.value
}

func (v *NullableStopJobResponse) Set(val *StopJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStopJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStopJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopJobResponse(val *StopJobResponse) *NullableStopJobResponse {
	return &NullableStopJobResponse{value: val, isSet: true}
}

func (v NullableStopJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
