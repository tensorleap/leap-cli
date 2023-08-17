/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.365
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the RotateApiKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RotateApiKeyResponse{}

// RotateApiKeyResponse struct for RotateApiKeyResponse
type RotateApiKeyResponse struct {
	ApiKey string `json:"apiKey"`
}

// NewRotateApiKeyResponse instantiates a new RotateApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotateApiKeyResponse(apiKey string) *RotateApiKeyResponse {
	this := RotateApiKeyResponse{}
	this.ApiKey = apiKey
	return &this
}

// NewRotateApiKeyResponseWithDefaults instantiates a new RotateApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotateApiKeyResponseWithDefaults() *RotateApiKeyResponse {
	this := RotateApiKeyResponse{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *RotateApiKeyResponse) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *RotateApiKeyResponse) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *RotateApiKeyResponse) SetApiKey(v string) {
	o.ApiKey = v
}

func (o RotateApiKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RotateApiKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiKey"] = o.ApiKey
	return toSerialize, nil
}

type NullableRotateApiKeyResponse struct {
	value *RotateApiKeyResponse
	isSet bool
}

func (v NullableRotateApiKeyResponse) Get() *RotateApiKeyResponse {
	return v.value
}

func (v *NullableRotateApiKeyResponse) Set(val *RotateApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRotateApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRotateApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotateApiKeyResponse(val *RotateApiKeyResponse) *NullableRotateApiKeyResponse {
	return &NullableRotateApiKeyResponse{value: val, isSet: true}
}

func (v NullableRotateApiKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotateApiKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
