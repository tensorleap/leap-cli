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

// checks if the UpdateSecretManagerParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSecretManagerParams{}

// UpdateSecretManagerParams struct for UpdateSecretManagerParams
type UpdateSecretManagerParams struct {
	Cid      string  `json:"cid"`
	Name     string  `json:"name"`
	AuthData *string `json:"authData,omitempty"`
}

// NewUpdateSecretManagerParams instantiates a new UpdateSecretManagerParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSecretManagerParams(cid string, name string) *UpdateSecretManagerParams {
	this := UpdateSecretManagerParams{}
	this.Cid = cid
	this.Name = name
	return &this
}

// NewUpdateSecretManagerParamsWithDefaults instantiates a new UpdateSecretManagerParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSecretManagerParamsWithDefaults() *UpdateSecretManagerParams {
	this := UpdateSecretManagerParams{}
	return &this
}

// GetCid returns the Cid field value
func (o *UpdateSecretManagerParams) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *UpdateSecretManagerParams) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *UpdateSecretManagerParams) SetCid(v string) {
	o.Cid = v
}

// GetName returns the Name field value
func (o *UpdateSecretManagerParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateSecretManagerParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateSecretManagerParams) SetName(v string) {
	o.Name = v
}

// GetAuthData returns the AuthData field value if set, zero value otherwise.
func (o *UpdateSecretManagerParams) GetAuthData() string {
	if o == nil || IsNil(o.AuthData) {
		var ret string
		return ret
	}
	return *o.AuthData
}

// GetAuthDataOk returns a tuple with the AuthData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSecretManagerParams) GetAuthDataOk() (*string, bool) {
	if o == nil || IsNil(o.AuthData) {
		return nil, false
	}
	return o.AuthData, true
}

// HasAuthData returns a boolean if a field has been set.
func (o *UpdateSecretManagerParams) HasAuthData() bool {
	if o != nil && !IsNil(o.AuthData) {
		return true
	}

	return false
}

// SetAuthData gets a reference to the given string and assigns it to the AuthData field.
func (o *UpdateSecretManagerParams) SetAuthData(v string) {
	o.AuthData = &v
}

func (o UpdateSecretManagerParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSecretManagerParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["name"] = o.Name
	if !IsNil(o.AuthData) {
		toSerialize["authData"] = o.AuthData
	}
	return toSerialize, nil
}

type NullableUpdateSecretManagerParams struct {
	value *UpdateSecretManagerParams
	isSet bool
}

func (v NullableUpdateSecretManagerParams) Get() *UpdateSecretManagerParams {
	return v.value
}

func (v *NullableUpdateSecretManagerParams) Set(val *UpdateSecretManagerParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSecretManagerParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSecretManagerParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSecretManagerParams(val *UpdateSecretManagerParams) *NullableUpdateSecretManagerParams {
	return &NullableUpdateSecretManagerParams{value: val, isSet: true}
}

func (v NullableUpdateSecretManagerParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSecretManagerParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
