/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.404
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the SecretManager type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretManager{}

// SecretManager struct for SecretManager
type SecretManager struct {
	Cid       string    `json:"cid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// NewSecretManager instantiates a new SecretManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretManager(cid string, name string, createdAt time.Time) *SecretManager {
	this := SecretManager{}
	this.Cid = cid
	this.Name = name
	this.CreatedAt = createdAt
	return &this
}

// NewSecretManagerWithDefaults instantiates a new SecretManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretManagerWithDefaults() *SecretManager {
	this := SecretManager{}
	return &this
}

// GetCid returns the Cid field value
func (o *SecretManager) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *SecretManager) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *SecretManager) SetCid(v string) {
	o.Cid = v
}

// GetName returns the Name field value
func (o *SecretManager) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecretManager) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecretManager) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SecretManager) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SecretManager) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SecretManager) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o SecretManager) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretManager) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["name"] = o.Name
	toSerialize["createdAt"] = o.CreatedAt
	return toSerialize, nil
}

type NullableSecretManager struct {
	value *SecretManager
	isSet bool
}

func (v NullableSecretManager) Get() *SecretManager {
	return v.value
}

func (v *NullableSecretManager) Set(val *SecretManager) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretManager) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretManager(val *SecretManager) *NullableSecretManager {
	return &NullableSecretManager{value: val, isSet: true}
}

func (v NullableSecretManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
