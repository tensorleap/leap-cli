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

// checks if the TrashSecretManagerParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrashSecretManagerParams{}

// TrashSecretManagerParams struct for TrashSecretManagerParams
type TrashSecretManagerParams struct {
	SecretManagerId string `json:"secretManagerId"`
}

// NewTrashSecretManagerParams instantiates a new TrashSecretManagerParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrashSecretManagerParams(secretManagerId string) *TrashSecretManagerParams {
	this := TrashSecretManagerParams{}
	this.SecretManagerId = secretManagerId
	return &this
}

// NewTrashSecretManagerParamsWithDefaults instantiates a new TrashSecretManagerParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrashSecretManagerParamsWithDefaults() *TrashSecretManagerParams {
	this := TrashSecretManagerParams{}
	return &this
}

// GetSecretManagerId returns the SecretManagerId field value
func (o *TrashSecretManagerParams) GetSecretManagerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretManagerId
}

// GetSecretManagerIdOk returns a tuple with the SecretManagerId field value
// and a boolean to check if the value has been set.
func (o *TrashSecretManagerParams) GetSecretManagerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretManagerId, true
}

// SetSecretManagerId sets field value
func (o *TrashSecretManagerParams) SetSecretManagerId(v string) {
	o.SecretManagerId = v
}

func (o TrashSecretManagerParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrashSecretManagerParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["secretManagerId"] = o.SecretManagerId
	return toSerialize, nil
}

type NullableTrashSecretManagerParams struct {
	value *TrashSecretManagerParams
	isSet bool
}

func (v NullableTrashSecretManagerParams) Get() *TrashSecretManagerParams {
	return v.value
}

func (v *NullableTrashSecretManagerParams) Set(val *TrashSecretManagerParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTrashSecretManagerParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTrashSecretManagerParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrashSecretManagerParams(val *TrashSecretManagerParams) *NullableTrashSecretManagerParams {
	return &NullableTrashSecretManagerParams{value: val, isSet: true}
}

func (v NullableTrashSecretManagerParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrashSecretManagerParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
