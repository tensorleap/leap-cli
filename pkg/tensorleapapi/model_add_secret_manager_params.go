/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the AddSecretManagerParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSecretManagerParams{}

// AddSecretManagerParams struct for AddSecretManagerParams
type AddSecretManagerParams struct {
	Name     string `json:"name"`
	AuthData string `json:"authData"`
}

// NewAddSecretManagerParams instantiates a new AddSecretManagerParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSecretManagerParams(name string, authData string) *AddSecretManagerParams {
	this := AddSecretManagerParams{}
	this.Name = name
	this.AuthData = authData
	return &this
}

// NewAddSecretManagerParamsWithDefaults instantiates a new AddSecretManagerParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSecretManagerParamsWithDefaults() *AddSecretManagerParams {
	this := AddSecretManagerParams{}
	return &this
}

// GetName returns the Name field value
func (o *AddSecretManagerParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddSecretManagerParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddSecretManagerParams) SetName(v string) {
	o.Name = v
}

// GetAuthData returns the AuthData field value
func (o *AddSecretManagerParams) GetAuthData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthData
}

// GetAuthDataOk returns a tuple with the AuthData field value
// and a boolean to check if the value has been set.
func (o *AddSecretManagerParams) GetAuthDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthData, true
}

// SetAuthData sets field value
func (o *AddSecretManagerParams) SetAuthData(v string) {
	o.AuthData = v
}

func (o AddSecretManagerParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSecretManagerParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["authData"] = o.AuthData
	return toSerialize, nil
}

type NullableAddSecretManagerParams struct {
	value *AddSecretManagerParams
	isSet bool
}

func (v NullableAddSecretManagerParams) Get() *AddSecretManagerParams {
	return v.value
}

func (v *NullableAddSecretManagerParams) Set(val *AddSecretManagerParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSecretManagerParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSecretManagerParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSecretManagerParams(val *AddSecretManagerParams) *NullableAddSecretManagerParams {
	return &NullableAddSecretManagerParams{value: val, isSet: true}
}

func (v NullableAddSecretManagerParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSecretManagerParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
