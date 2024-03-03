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

// checks if the SetDefaultTeamRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetDefaultTeamRequest{}

// SetDefaultTeamRequest struct for SetDefaultTeamRequest
type SetDefaultTeamRequest struct {
	Cid string `json:"cid"`
}

// NewSetDefaultTeamRequest instantiates a new SetDefaultTeamRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetDefaultTeamRequest(cid string) *SetDefaultTeamRequest {
	this := SetDefaultTeamRequest{}
	this.Cid = cid
	return &this
}

// NewSetDefaultTeamRequestWithDefaults instantiates a new SetDefaultTeamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetDefaultTeamRequestWithDefaults() *SetDefaultTeamRequest {
	this := SetDefaultTeamRequest{}
	return &this
}

// GetCid returns the Cid field value
func (o *SetDefaultTeamRequest) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *SetDefaultTeamRequest) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *SetDefaultTeamRequest) SetCid(v string) {
	o.Cid = v
}

func (o SetDefaultTeamRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetDefaultTeamRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	return toSerialize, nil
}

type NullableSetDefaultTeamRequest struct {
	value *SetDefaultTeamRequest
	isSet bool
}

func (v NullableSetDefaultTeamRequest) Get() *SetDefaultTeamRequest {
	return v.value
}

func (v *NullableSetDefaultTeamRequest) Set(val *SetDefaultTeamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetDefaultTeamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetDefaultTeamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetDefaultTeamRequest(val *SetDefaultTeamRequest) *NullableSetDefaultTeamRequest {
	return &NullableSetDefaultTeamRequest{value: val, isSet: true}
}

func (v NullableSetDefaultTeamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetDefaultTeamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
