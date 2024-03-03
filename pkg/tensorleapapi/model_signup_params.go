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

// checks if the SignupParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignupParams{}

// SignupParams struct for SignupParams
type SignupParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewSignupParams instantiates a new SignupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignupParams(name string, email string, password string) *SignupParams {
	this := SignupParams{}
	this.Name = name
	this.Email = email
	this.Password = password
	return &this
}

// NewSignupParamsWithDefaults instantiates a new SignupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignupParamsWithDefaults() *SignupParams {
	this := SignupParams{}
	return &this
}

// GetName returns the Name field value
func (o *SignupParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SignupParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SignupParams) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *SignupParams) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SignupParams) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SignupParams) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *SignupParams) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SignupParams) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SignupParams) SetPassword(v string) {
	o.Password = v
}

func (o SignupParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignupParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableSignupParams struct {
	value *SignupParams
	isSet bool
}

func (v NullableSignupParams) Get() *SignupParams {
	return v.value
}

func (v *NullableSignupParams) Set(val *SignupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSignupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSignupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignupParams(val *SignupParams) *NullableSignupParams {
	return &NullableSignupParams{value: val, isSet: true}
}

func (v NullableSignupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
