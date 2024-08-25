/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.526
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetTeamUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTeamUsersResponse{}

// GetTeamUsersResponse struct for GetTeamUsersResponse
type GetTeamUsersResponse struct {
	Users []SlimUserData `json:"users"`
}

// NewGetTeamUsersResponse instantiates a new GetTeamUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTeamUsersResponse(users []SlimUserData) *GetTeamUsersResponse {
	this := GetTeamUsersResponse{}
	this.Users = users
	return &this
}

// NewGetTeamUsersResponseWithDefaults instantiates a new GetTeamUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTeamUsersResponseWithDefaults() *GetTeamUsersResponse {
	this := GetTeamUsersResponse{}
	return &this
}

// GetUsers returns the Users field value
func (o *GetTeamUsersResponse) GetUsers() []SlimUserData {
	if o == nil {
		var ret []SlimUserData
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *GetTeamUsersResponse) GetUsersOk() ([]SlimUserData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *GetTeamUsersResponse) SetUsers(v []SlimUserData) {
	o.Users = v
}

func (o GetTeamUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTeamUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

type NullableGetTeamUsersResponse struct {
	value *GetTeamUsersResponse
	isSet bool
}

func (v NullableGetTeamUsersResponse) Get() *GetTeamUsersResponse {
	return v.value
}

func (v *NullableGetTeamUsersResponse) Set(val *GetTeamUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTeamUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTeamUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTeamUsersResponse(val *GetTeamUsersResponse) *NullableGetTeamUsersResponse {
	return &NullableGetTeamUsersResponse{value: val, isSet: true}
}

func (v NullableGetTeamUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTeamUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
