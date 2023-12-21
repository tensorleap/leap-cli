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

// checks if the TerminateJobParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminateJobParams{}

// TerminateJobParams struct for TerminateJobParams
type TerminateJobParams struct {
	JobId string `json:"jobId"`
}

// NewTerminateJobParams instantiates a new TerminateJobParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminateJobParams(jobId string) *TerminateJobParams {
	this := TerminateJobParams{}
	this.JobId = jobId
	return &this
}

// NewTerminateJobParamsWithDefaults instantiates a new TerminateJobParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminateJobParamsWithDefaults() *TerminateJobParams {
	this := TerminateJobParams{}
	return &this
}

// GetJobId returns the JobId field value
func (o *TerminateJobParams) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *TerminateJobParams) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *TerminateJobParams) SetJobId(v string) {
	o.JobId = v
}

func (o TerminateJobParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminateJobParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobId"] = o.JobId
	return toSerialize, nil
}

type NullableTerminateJobParams struct {
	value *TerminateJobParams
	isSet bool
}

func (v NullableTerminateJobParams) Get() *TerminateJobParams {
	return v.value
}

func (v *NullableTerminateJobParams) Set(val *TerminateJobParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminateJobParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminateJobParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminateJobParams(val *TerminateJobParams) *NullableTerminateJobParams {
	return &NullableTerminateJobParams{value: val, isSet: true}
}

func (v NullableTerminateJobParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminateJobParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


