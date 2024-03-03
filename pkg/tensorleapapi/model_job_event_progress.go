/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the JobEventProgress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobEventProgress{}

// JobEventProgress struct for JobEventProgress
type JobEventProgress struct {
	Total   float64 `json:"total"`
	Current float64 `json:"current"`
}

// NewJobEventProgress instantiates a new JobEventProgress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobEventProgress(total float64, current float64) *JobEventProgress {
	this := JobEventProgress{}
	this.Total = total
	this.Current = current
	return &this
}

// NewJobEventProgressWithDefaults instantiates a new JobEventProgress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobEventProgressWithDefaults() *JobEventProgress {
	this := JobEventProgress{}
	return &this
}

// GetTotal returns the Total field value
func (o *JobEventProgress) GetTotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *JobEventProgress) GetTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *JobEventProgress) SetTotal(v float64) {
	o.Total = v
}

// GetCurrent returns the Current field value
func (o *JobEventProgress) GetCurrent() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *JobEventProgress) GetCurrentOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *JobEventProgress) SetCurrent(v float64) {
	o.Current = v
}

func (o JobEventProgress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobEventProgress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["current"] = o.Current
	return toSerialize, nil
}

type NullableJobEventProgress struct {
	value *JobEventProgress
	isSet bool
}

func (v NullableJobEventProgress) Get() *JobEventProgress {
	return v.value
}

func (v *NullableJobEventProgress) Set(val *JobEventProgress) {
	v.value = val
	v.isSet = true
}

func (v NullableJobEventProgress) IsSet() bool {
	return v.isSet
}

func (v *NullableJobEventProgress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobEventProgress(val *JobEventProgress) *NullableJobEventProgress {
	return &NullableJobEventProgress{value: val, isSet: true}
}

func (v NullableJobEventProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobEventProgress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
