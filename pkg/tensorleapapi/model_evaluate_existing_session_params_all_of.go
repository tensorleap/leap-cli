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

// checks if the EvaluateExistingSessionParamsAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvaluateExistingSessionParamsAllOf{}

// EvaluateExistingSessionParamsAllOf struct for EvaluateExistingSessionParamsAllOf
type EvaluateExistingSessionParamsAllOf struct {
	EvaluatedEpoch float64 `json:"evaluatedEpoch"`
	SessionId      string  `json:"sessionId"`
}

// NewEvaluateExistingSessionParamsAllOf instantiates a new EvaluateExistingSessionParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvaluateExistingSessionParamsAllOf(evaluatedEpoch float64, sessionId string) *EvaluateExistingSessionParamsAllOf {
	this := EvaluateExistingSessionParamsAllOf{}
	this.EvaluatedEpoch = evaluatedEpoch
	this.SessionId = sessionId
	return &this
}

// NewEvaluateExistingSessionParamsAllOfWithDefaults instantiates a new EvaluateExistingSessionParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvaluateExistingSessionParamsAllOfWithDefaults() *EvaluateExistingSessionParamsAllOf {
	this := EvaluateExistingSessionParamsAllOf{}
	return &this
}

// GetEvaluatedEpoch returns the EvaluatedEpoch field value
func (o *EvaluateExistingSessionParamsAllOf) GetEvaluatedEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.EvaluatedEpoch
}

// GetEvaluatedEpochOk returns a tuple with the EvaluatedEpoch field value
// and a boolean to check if the value has been set.
func (o *EvaluateExistingSessionParamsAllOf) GetEvaluatedEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluatedEpoch, true
}

// SetEvaluatedEpoch sets field value
func (o *EvaluateExistingSessionParamsAllOf) SetEvaluatedEpoch(v float64) {
	o.EvaluatedEpoch = v
}

// GetSessionId returns the SessionId field value
func (o *EvaluateExistingSessionParamsAllOf) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *EvaluateExistingSessionParamsAllOf) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *EvaluateExistingSessionParamsAllOf) SetSessionId(v string) {
	o.SessionId = v
}

func (o EvaluateExistingSessionParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvaluateExistingSessionParamsAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["evaluatedEpoch"] = o.EvaluatedEpoch
	toSerialize["sessionId"] = o.SessionId
	return toSerialize, nil
}

type NullableEvaluateExistingSessionParamsAllOf struct {
	value *EvaluateExistingSessionParamsAllOf
	isSet bool
}

func (v NullableEvaluateExistingSessionParamsAllOf) Get() *EvaluateExistingSessionParamsAllOf {
	return v.value
}

func (v *NullableEvaluateExistingSessionParamsAllOf) Set(val *EvaluateExistingSessionParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEvaluateExistingSessionParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEvaluateExistingSessionParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvaluateExistingSessionParamsAllOf(val *EvaluateExistingSessionParamsAllOf) *NullableEvaluateExistingSessionParamsAllOf {
	return &NullableEvaluateExistingSessionParamsAllOf{value: val, isSet: true}
}

func (v NullableEvaluateExistingSessionParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvaluateExistingSessionParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
