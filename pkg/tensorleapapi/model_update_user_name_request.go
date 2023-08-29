/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.378-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UpdateUserNameRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserNameRequest{}

// UpdateUserNameRequest struct for UpdateUserNameRequest
type UpdateUserNameRequest struct {
	UserName string `json:"userName"`
}

// NewUpdateUserNameRequest instantiates a new UpdateUserNameRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserNameRequest(userName string) *UpdateUserNameRequest {
	this := UpdateUserNameRequest{}
	this.UserName = userName
	return &this
}

// NewUpdateUserNameRequestWithDefaults instantiates a new UpdateUserNameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserNameRequestWithDefaults() *UpdateUserNameRequest {
	this := UpdateUserNameRequest{}
	return &this
}

// GetUserName returns the UserName field value
func (o *UpdateUserNameRequest) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *UpdateUserNameRequest) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *UpdateUserNameRequest) SetUserName(v string) {
	o.UserName = v
}

func (o UpdateUserNameRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserNameRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userName"] = o.UserName
	return toSerialize, nil
}

type NullableUpdateUserNameRequest struct {
	value *UpdateUserNameRequest
	isSet bool
}

func (v NullableUpdateUserNameRequest) Get() *UpdateUserNameRequest {
	return v.value
}

func (v *NullableUpdateUserNameRequest) Set(val *UpdateUserNameRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserNameRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserNameRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserNameRequest(val *UpdateUserNameRequest) *NullableUpdateUserNameRequest {
	return &NullableUpdateUserNameRequest{value: val, isSet: true}
}

func (v NullableUpdateUserNameRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserNameRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


