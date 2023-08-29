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

// checks if the ExtendTrialParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtendTrialParams{}

// ExtendTrialParams struct for ExtendTrialParams
type ExtendTrialParams struct {
	Token string `json:"token"`
}

// NewExtendTrialParams instantiates a new ExtendTrialParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendTrialParams(token string) *ExtendTrialParams {
	this := ExtendTrialParams{}
	this.Token = token
	return &this
}

// NewExtendTrialParamsWithDefaults instantiates a new ExtendTrialParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendTrialParamsWithDefaults() *ExtendTrialParams {
	this := ExtendTrialParams{}
	return &this
}

// GetToken returns the Token field value
func (o *ExtendTrialParams) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ExtendTrialParams) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ExtendTrialParams) SetToken(v string) {
	o.Token = v
}

func (o ExtendTrialParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendTrialParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableExtendTrialParams struct {
	value *ExtendTrialParams
	isSet bool
}

func (v NullableExtendTrialParams) Get() *ExtendTrialParams {
	return v.value
}

func (v *NullableExtendTrialParams) Set(val *ExtendTrialParams) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendTrialParams) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendTrialParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendTrialParams(val *ExtendTrialParams) *NullableExtendTrialParams {
	return &NullableExtendTrialParams{value: val, isSet: true}
}

func (v NullableExtendTrialParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendTrialParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


