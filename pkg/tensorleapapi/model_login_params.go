/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the LoginParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginParams{}

// LoginParams struct for LoginParams
type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewLoginParams instantiates a new LoginParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginParams(email string, password string) *LoginParams {
	this := LoginParams{}
	this.Email = email
	this.Password = password
	return &this
}

// NewLoginParamsWithDefaults instantiates a new LoginParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginParamsWithDefaults() *LoginParams {
	this := LoginParams{}
	return &this
}

// GetEmail returns the Email field value
func (o *LoginParams) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *LoginParams) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *LoginParams) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *LoginParams) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *LoginParams) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *LoginParams) SetPassword(v string) {
	o.Password = v
}

func (o LoginParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableLoginParams struct {
	value *LoginParams
	isSet bool
}

func (v NullableLoginParams) Get() *LoginParams {
	return v.value
}

func (v *NullableLoginParams) Set(val *LoginParams) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginParams) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginParams(val *LoginParams) *NullableLoginParams {
	return &NullableLoginParams{value: val, isSet: true}
}

func (v NullableLoginParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
