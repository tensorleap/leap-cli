/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.557
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the AddSecretManagerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSecretManagerResponse{}

// AddSecretManagerResponse struct for AddSecretManagerResponse
type AddSecretManagerResponse struct {
	Cid     string `json:"cid"`
	Success bool   `json:"success"`
}

// NewAddSecretManagerResponse instantiates a new AddSecretManagerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSecretManagerResponse(cid string, success bool) *AddSecretManagerResponse {
	this := AddSecretManagerResponse{}
	this.Cid = cid
	this.Success = success
	return &this
}

// NewAddSecretManagerResponseWithDefaults instantiates a new AddSecretManagerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSecretManagerResponseWithDefaults() *AddSecretManagerResponse {
	this := AddSecretManagerResponse{}
	return &this
}

// GetCid returns the Cid field value
func (o *AddSecretManagerResponse) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *AddSecretManagerResponse) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *AddSecretManagerResponse) SetCid(v string) {
	o.Cid = v
}

// GetSuccess returns the Success field value
func (o *AddSecretManagerResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AddSecretManagerResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AddSecretManagerResponse) SetSuccess(v bool) {
	o.Success = v
}

func (o AddSecretManagerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSecretManagerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

type NullableAddSecretManagerResponse struct {
	value *AddSecretManagerResponse
	isSet bool
}

func (v NullableAddSecretManagerResponse) Get() *AddSecretManagerResponse {
	return v.value
}

func (v *NullableAddSecretManagerResponse) Set(val *AddSecretManagerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSecretManagerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSecretManagerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSecretManagerResponse(val *AddSecretManagerResponse) *NullableAddSecretManagerResponse {
	return &NullableAddSecretManagerResponse{value: val, isSet: true}
}

func (v NullableAddSecretManagerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSecretManagerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
