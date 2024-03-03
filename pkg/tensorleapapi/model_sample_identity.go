/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SampleIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SampleIdentity{}

// SampleIdentity struct for SampleIdentity
type SampleIdentity struct {
	State DataStateType `json:"state"`
	Index float64       `json:"index"`
}

// NewSampleIdentity instantiates a new SampleIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSampleIdentity(state DataStateType, index float64) *SampleIdentity {
	this := SampleIdentity{}
	this.State = state
	this.Index = index
	return &this
}

// NewSampleIdentityWithDefaults instantiates a new SampleIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSampleIdentityWithDefaults() *SampleIdentity {
	this := SampleIdentity{}
	return &this
}

// GetState returns the State field value
func (o *SampleIdentity) GetState() DataStateType {
	if o == nil {
		var ret DataStateType
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SampleIdentity) GetStateOk() (*DataStateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SampleIdentity) SetState(v DataStateType) {
	o.State = v
}

// GetIndex returns the Index field value
func (o *SampleIdentity) GetIndex() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SampleIdentity) GetIndexOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SampleIdentity) SetIndex(v float64) {
	o.Index = v
}

func (o SampleIdentity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SampleIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["index"] = o.Index
	return toSerialize, nil
}

type NullableSampleIdentity struct {
	value *SampleIdentity
	isSet bool
}

func (v NullableSampleIdentity) Get() *SampleIdentity {
	return v.value
}

func (v *NullableSampleIdentity) Set(val *SampleIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableSampleIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableSampleIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSampleIdentity(val *SampleIdentity) *NullableSampleIdentity {
	return &NullableSampleIdentity{value: val, isSet: true}
}

func (v NullableSampleIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSampleIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
