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

// checks if the ActivationParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivationParams{}

// ActivationParams struct for ActivationParams
type ActivationParams struct {
	Token string `json:"token"`
}

// NewActivationParams instantiates a new ActivationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivationParams(token string) *ActivationParams {
	this := ActivationParams{}
	this.Token = token
	return &this
}

// NewActivationParamsWithDefaults instantiates a new ActivationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivationParamsWithDefaults() *ActivationParams {
	this := ActivationParams{}
	return &this
}

// GetToken returns the Token field value
func (o *ActivationParams) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ActivationParams) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ActivationParams) SetToken(v string) {
	o.Token = v
}

func (o ActivationParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivationParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableActivationParams struct {
	value *ActivationParams
	isSet bool
}

func (v NullableActivationParams) Get() *ActivationParams {
	return v.value
}

func (v *NullableActivationParams) Set(val *ActivationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableActivationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableActivationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivationParams(val *ActivationParams) *NullableActivationParams {
	return &NullableActivationParams{value: val, isSet: true}
}

func (v NullableActivationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
