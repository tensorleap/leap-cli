/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.456
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UpdateTeamPublicNameRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTeamPublicNameRequest{}

// UpdateTeamPublicNameRequest struct for UpdateTeamPublicNameRequest
type UpdateTeamPublicNameRequest struct {
	Cid        string `json:"cid"`
	PublicName string `json:"publicName"`
}

// NewUpdateTeamPublicNameRequest instantiates a new UpdateTeamPublicNameRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTeamPublicNameRequest(cid string, publicName string) *UpdateTeamPublicNameRequest {
	this := UpdateTeamPublicNameRequest{}
	this.Cid = cid
	this.PublicName = publicName
	return &this
}

// NewUpdateTeamPublicNameRequestWithDefaults instantiates a new UpdateTeamPublicNameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTeamPublicNameRequestWithDefaults() *UpdateTeamPublicNameRequest {
	this := UpdateTeamPublicNameRequest{}
	return &this
}

// GetCid returns the Cid field value
func (o *UpdateTeamPublicNameRequest) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *UpdateTeamPublicNameRequest) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *UpdateTeamPublicNameRequest) SetCid(v string) {
	o.Cid = v
}

// GetPublicName returns the PublicName field value
func (o *UpdateTeamPublicNameRequest) GetPublicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicName
}

// GetPublicNameOk returns a tuple with the PublicName field value
// and a boolean to check if the value has been set.
func (o *UpdateTeamPublicNameRequest) GetPublicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicName, true
}

// SetPublicName sets field value
func (o *UpdateTeamPublicNameRequest) SetPublicName(v string) {
	o.PublicName = v
}

func (o UpdateTeamPublicNameRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTeamPublicNameRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["publicName"] = o.PublicName
	return toSerialize, nil
}

type NullableUpdateTeamPublicNameRequest struct {
	value *UpdateTeamPublicNameRequest
	isSet bool
}

func (v NullableUpdateTeamPublicNameRequest) Get() *UpdateTeamPublicNameRequest {
	return v.value
}

func (v *NullableUpdateTeamPublicNameRequest) Set(val *UpdateTeamPublicNameRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTeamPublicNameRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTeamPublicNameRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTeamPublicNameRequest(val *UpdateTeamPublicNameRequest) *NullableUpdateTeamPublicNameRequest {
	return &NullableUpdateTeamPublicNameRequest{value: val, isSet: true}
}

func (v NullableUpdateTeamPublicNameRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTeamPublicNameRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
