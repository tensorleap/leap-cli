/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.417
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UpdateUserStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserStatusRequest{}

// UpdateUserStatusRequest struct for UpdateUserStatusRequest
type UpdateUserStatusRequest struct {
	UserId string `json:"userId"`
	Active bool   `json:"active"`
}

// NewUpdateUserStatusRequest instantiates a new UpdateUserStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserStatusRequest(userId string, active bool) *UpdateUserStatusRequest {
	this := UpdateUserStatusRequest{}
	this.UserId = userId
	this.Active = active
	return &this
}

// NewUpdateUserStatusRequestWithDefaults instantiates a new UpdateUserStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserStatusRequestWithDefaults() *UpdateUserStatusRequest {
	this := UpdateUserStatusRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *UpdateUserStatusRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UpdateUserStatusRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UpdateUserStatusRequest) SetUserId(v string) {
	o.UserId = v
}

// GetActive returns the Active field value
func (o *UpdateUserStatusRequest) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *UpdateUserStatusRequest) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *UpdateUserStatusRequest) SetActive(v bool) {
	o.Active = v
}

func (o UpdateUserStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["active"] = o.Active
	return toSerialize, nil
}

type NullableUpdateUserStatusRequest struct {
	value *UpdateUserStatusRequest
	isSet bool
}

func (v NullableUpdateUserStatusRequest) Get() *UpdateUserStatusRequest {
	return v.value
}

func (v *NullableUpdateUserStatusRequest) Set(val *UpdateUserStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserStatusRequest(val *UpdateUserStatusRequest) *NullableUpdateUserStatusRequest {
	return &NullableUpdateUserStatusRequest{value: val, isSet: true}
}

func (v NullableUpdateUserStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
