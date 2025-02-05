/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.592
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SessionRunToEpoch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionRunToEpoch{}

// SessionRunToEpoch struct for SessionRunToEpoch
type SessionRunToEpoch struct {
	SessionRunId string  `json:"sessionRunId"`
	Epoch        float64 `json:"epoch"`
}

// NewSessionRunToEpoch instantiates a new SessionRunToEpoch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionRunToEpoch(sessionRunId string, epoch float64) *SessionRunToEpoch {
	this := SessionRunToEpoch{}
	this.SessionRunId = sessionRunId
	this.Epoch = epoch
	return &this
}

// NewSessionRunToEpochWithDefaults instantiates a new SessionRunToEpoch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionRunToEpochWithDefaults() *SessionRunToEpoch {
	this := SessionRunToEpoch{}
	return &this
}

// GetSessionRunId returns the SessionRunId field value
func (o *SessionRunToEpoch) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *SessionRunToEpoch) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *SessionRunToEpoch) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetEpoch returns the Epoch field value
func (o *SessionRunToEpoch) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *SessionRunToEpoch) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *SessionRunToEpoch) SetEpoch(v float64) {
	o.Epoch = v
}

func (o SessionRunToEpoch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionRunToEpoch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["epoch"] = o.Epoch
	return toSerialize, nil
}

type NullableSessionRunToEpoch struct {
	value *SessionRunToEpoch
	isSet bool
}

func (v NullableSessionRunToEpoch) Get() *SessionRunToEpoch {
	return v.value
}

func (v *NullableSessionRunToEpoch) Set(val *SessionRunToEpoch) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionRunToEpoch) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionRunToEpoch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionRunToEpoch(val *SessionRunToEpoch) *NullableSessionRunToEpoch {
	return &NullableSessionRunToEpoch{value: val, isSet: true}
}

func (v NullableSessionRunToEpoch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionRunToEpoch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
