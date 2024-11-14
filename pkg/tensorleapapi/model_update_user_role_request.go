/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.557
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UpdateUserRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserRoleRequest{}

// UpdateUserRoleRequest struct for UpdateUserRoleRequest
type UpdateUserRoleRequest struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
}

// NewUpdateUserRoleRequest instantiates a new UpdateUserRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserRoleRequest(userId string, role string) *UpdateUserRoleRequest {
	this := UpdateUserRoleRequest{}
	this.UserId = userId
	this.Role = role
	return &this
}

// NewUpdateUserRoleRequestWithDefaults instantiates a new UpdateUserRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserRoleRequestWithDefaults() *UpdateUserRoleRequest {
	this := UpdateUserRoleRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *UpdateUserRoleRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UpdateUserRoleRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UpdateUserRoleRequest) SetUserId(v string) {
	o.UserId = v
}

// GetRole returns the Role field value
func (o *UpdateUserRoleRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UpdateUserRoleRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UpdateUserRoleRequest) SetRole(v string) {
	o.Role = v
}

func (o UpdateUserRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

type NullableUpdateUserRoleRequest struct {
	value *UpdateUserRoleRequest
	isSet bool
}

func (v NullableUpdateUserRoleRequest) Get() *UpdateUserRoleRequest {
	return v.value
}

func (v *NullableUpdateUserRoleRequest) Set(val *UpdateUserRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserRoleRequest(val *UpdateUserRoleRequest) *NullableUpdateUserRoleRequest {
	return &NullableUpdateUserRoleRequest{value: val, isSet: true}
}

func (v NullableUpdateUserRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
