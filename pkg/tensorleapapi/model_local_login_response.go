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

// checks if the LocalLoginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalLoginResponse{}

// LocalLoginResponse struct for LocalLoginResponse
type LocalLoginResponse struct {
	Token *string   `json:"token,omitempty"`
	User  *UserData `json:"user,omitempty"`
}

// NewLocalLoginResponse instantiates a new LocalLoginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalLoginResponse() *LocalLoginResponse {
	this := LocalLoginResponse{}
	return &this
}

// NewLocalLoginResponseWithDefaults instantiates a new LocalLoginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalLoginResponseWithDefaults() *LocalLoginResponse {
	this := LocalLoginResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LocalLoginResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalLoginResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LocalLoginResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LocalLoginResponse) SetToken(v string) {
	o.Token = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LocalLoginResponse) GetUser() UserData {
	if o == nil || IsNil(o.User) {
		var ret UserData
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalLoginResponse) GetUserOk() (*UserData, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LocalLoginResponse) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserData and assigns it to the User field.
func (o *LocalLoginResponse) SetUser(v UserData) {
	o.User = &v
}

func (o LocalLoginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalLoginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableLocalLoginResponse struct {
	value *LocalLoginResponse
	isSet bool
}

func (v NullableLocalLoginResponse) Get() *LocalLoginResponse {
	return v.value
}

func (v *NullableLocalLoginResponse) Set(val *LocalLoginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalLoginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalLoginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalLoginResponse(val *LocalLoginResponse) *NullableLocalLoginResponse {
	return &NullableLocalLoginResponse{value: val, isSet: true}
}

func (v NullableLocalLoginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalLoginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}