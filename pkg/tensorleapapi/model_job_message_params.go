/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.445
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the JobMessageParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobMessageParams{}

// JobMessageParams struct for JobMessageParams
type JobMessageParams struct {
	Module    string    `json:"module"`
	JobStatus JobStatus `json:"jobStatus"`
}

// NewJobMessageParams instantiates a new JobMessageParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobMessageParams(module string, jobStatus JobStatus) *JobMessageParams {
	this := JobMessageParams{}
	this.Module = module
	this.JobStatus = jobStatus
	return &this
}

// NewJobMessageParamsWithDefaults instantiates a new JobMessageParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobMessageParamsWithDefaults() *JobMessageParams {
	this := JobMessageParams{}
	return &this
}

// GetModule returns the Module field value
func (o *JobMessageParams) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *JobMessageParams) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *JobMessageParams) SetModule(v string) {
	o.Module = v
}

// GetJobStatus returns the JobStatus field value
func (o *JobMessageParams) GetJobStatus() JobStatus {
	if o == nil {
		var ret JobStatus
		return ret
	}

	return o.JobStatus
}

// GetJobStatusOk returns a tuple with the JobStatus field value
// and a boolean to check if the value has been set.
func (o *JobMessageParams) GetJobStatusOk() (*JobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobStatus, true
}

// SetJobStatus sets field value
func (o *JobMessageParams) SetJobStatus(v JobStatus) {
	o.JobStatus = v
}

func (o JobMessageParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobMessageParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["module"] = o.Module
	toSerialize["jobStatus"] = o.JobStatus
	return toSerialize, nil
}

type NullableJobMessageParams struct {
	value *JobMessageParams
	isSet bool
}

func (v NullableJobMessageParams) Get() *JobMessageParams {
	return v.value
}

func (v *NullableJobMessageParams) Set(val *JobMessageParams) {
	v.value = val
	v.isSet = true
}

func (v NullableJobMessageParams) IsSet() bool {
	return v.isSet
}

func (v *NullableJobMessageParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobMessageParams(val *JobMessageParams) *NullableJobMessageParams {
	return &NullableJobMessageParams{value: val, isSet: true}
}

func (v NullableJobMessageParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobMessageParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
