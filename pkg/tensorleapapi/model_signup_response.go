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

// checks if the SignupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignupResponse{}

// SignupResponse struct for SignupResponse
type SignupResponse struct {
	UserCid string `json:"userCid"`
	Success bool   `json:"success"`
}

// NewSignupResponse instantiates a new SignupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignupResponse(userCid string, success bool) *SignupResponse {
	this := SignupResponse{}
	this.UserCid = userCid
	this.Success = success
	return &this
}

// NewSignupResponseWithDefaults instantiates a new SignupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignupResponseWithDefaults() *SignupResponse {
	this := SignupResponse{}
	return &this
}

// GetUserCid returns the UserCid field value
func (o *SignupResponse) GetUserCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserCid
}

// GetUserCidOk returns a tuple with the UserCid field value
// and a boolean to check if the value has been set.
func (o *SignupResponse) GetUserCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserCid, true
}

// SetUserCid sets field value
func (o *SignupResponse) SetUserCid(v string) {
	o.UserCid = v
}

// GetSuccess returns the Success field value
func (o *SignupResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *SignupResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *SignupResponse) SetSuccess(v bool) {
	o.Success = v
}

func (o SignupResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userCid"] = o.UserCid
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

type NullableSignupResponse struct {
	value *SignupResponse
	isSet bool
}

func (v NullableSignupResponse) Get() *SignupResponse {
	return v.value
}

func (v *NullableSignupResponse) Set(val *SignupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSignupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSignupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignupResponse(val *SignupResponse) *NullableSignupResponse {
	return &NullableSignupResponse{value: val, isSet: true}
}

func (v NullableSignupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
