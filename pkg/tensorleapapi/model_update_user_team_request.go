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

// checks if the UpdateUserTeamRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserTeamRequest{}

// UpdateUserTeamRequest struct for UpdateUserTeamRequest
type UpdateUserTeamRequest struct {
	UserId string `json:"userId"`
	TeamId string `json:"teamId"`
}

// NewUpdateUserTeamRequest instantiates a new UpdateUserTeamRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserTeamRequest(userId string, teamId string) *UpdateUserTeamRequest {
	this := UpdateUserTeamRequest{}
	this.UserId = userId
	this.TeamId = teamId
	return &this
}

// NewUpdateUserTeamRequestWithDefaults instantiates a new UpdateUserTeamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserTeamRequestWithDefaults() *UpdateUserTeamRequest {
	this := UpdateUserTeamRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *UpdateUserTeamRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UpdateUserTeamRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UpdateUserTeamRequest) SetUserId(v string) {
	o.UserId = v
}

// GetTeamId returns the TeamId field value
func (o *UpdateUserTeamRequest) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *UpdateUserTeamRequest) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *UpdateUserTeamRequest) SetTeamId(v string) {
	o.TeamId = v
}

func (o UpdateUserTeamRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserTeamRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["teamId"] = o.TeamId
	return toSerialize, nil
}

type NullableUpdateUserTeamRequest struct {
	value *UpdateUserTeamRequest
	isSet bool
}

func (v NullableUpdateUserTeamRequest) Get() *UpdateUserTeamRequest {
	return v.value
}

func (v *NullableUpdateUserTeamRequest) Set(val *UpdateUserTeamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserTeamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserTeamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserTeamRequest(val *UpdateUserTeamRequest) *NullableUpdateUserTeamRequest {
	return &NullableUpdateUserTeamRequest{value: val, isSet: true}
}

func (v NullableUpdateUserTeamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserTeamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


