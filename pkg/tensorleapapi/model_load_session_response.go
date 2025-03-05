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

// checks if the LoadSessionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadSessionResponse{}

// LoadSessionResponse struct for LoadSessionResponse
type LoadSessionResponse struct {
	Session SessionPopulatedJob `json:"session"`
}

// NewLoadSessionResponse instantiates a new LoadSessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadSessionResponse(session SessionPopulatedJob) *LoadSessionResponse {
	this := LoadSessionResponse{}
	this.Session = session
	return &this
}

// NewLoadSessionResponseWithDefaults instantiates a new LoadSessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadSessionResponseWithDefaults() *LoadSessionResponse {
	this := LoadSessionResponse{}
	return &this
}

// GetSession returns the Session field value
func (o *LoadSessionResponse) GetSession() SessionPopulatedJob {
	if o == nil {
		var ret SessionPopulatedJob
		return ret
	}

	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *LoadSessionResponse) GetSessionOk() (*SessionPopulatedJob, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value
func (o *LoadSessionResponse) SetSession(v SessionPopulatedJob) {
	o.Session = v
}

func (o LoadSessionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadSessionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["session"] = o.Session
	return toSerialize, nil
}

type NullableLoadSessionResponse struct {
	value *LoadSessionResponse
	isSet bool
}

func (v NullableLoadSessionResponse) Get() *LoadSessionResponse {
	return v.value
}

func (v *NullableLoadSessionResponse) Set(val *LoadSessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadSessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadSessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadSessionResponse(val *LoadSessionResponse) *NullableLoadSessionResponse {
	return &NullableLoadSessionResponse{value: val, isSet: true}
}

func (v NullableLoadSessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadSessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
