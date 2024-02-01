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

// checks if the StopJobParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopJobParams{}

// StopJobParams struct for StopJobParams
type StopJobParams struct {
	JobId string `json:"jobId"`
}

// NewStopJobParams instantiates a new StopJobParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopJobParams(jobId string) *StopJobParams {
	this := StopJobParams{}
	this.JobId = jobId
	return &this
}

// NewStopJobParamsWithDefaults instantiates a new StopJobParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopJobParamsWithDefaults() *StopJobParams {
	this := StopJobParams{}
	return &this
}

// GetJobId returns the JobId field value
func (o *StopJobParams) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *StopJobParams) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *StopJobParams) SetJobId(v string) {
	o.JobId = v
}

func (o StopJobParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopJobParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobId"] = o.JobId
	return toSerialize, nil
}

type NullableStopJobParams struct {
	value *StopJobParams
	isSet bool
}

func (v NullableStopJobParams) Get() *StopJobParams {
	return v.value
}

func (v *NullableStopJobParams) Set(val *StopJobParams) {
	v.value = val
	v.isSet = true
}

func (v NullableStopJobParams) IsSet() bool {
	return v.isSet
}

func (v *NullableStopJobParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopJobParams(val *StopJobParams) *NullableStopJobParams {
	return &NullableStopJobParams{value: val, isSet: true}
}

func (v NullableStopJobParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopJobParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
