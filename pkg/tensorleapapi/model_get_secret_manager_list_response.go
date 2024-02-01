/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.445
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetSecretManagerListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSecretManagerListResponse{}

// GetSecretManagerListResponse struct for GetSecretManagerListResponse
type GetSecretManagerListResponse struct {
	Results []SecretManager `json:"results"`
}

// NewGetSecretManagerListResponse instantiates a new GetSecretManagerListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSecretManagerListResponse(results []SecretManager) *GetSecretManagerListResponse {
	this := GetSecretManagerListResponse{}
	this.Results = results
	return &this
}

// NewGetSecretManagerListResponseWithDefaults instantiates a new GetSecretManagerListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSecretManagerListResponseWithDefaults() *GetSecretManagerListResponse {
	this := GetSecretManagerListResponse{}
	return &this
}

// GetResults returns the Results field value
func (o *GetSecretManagerListResponse) GetResults() []SecretManager {
	if o == nil {
		var ret []SecretManager
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *GetSecretManagerListResponse) GetResultsOk() ([]SecretManager, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *GetSecretManagerListResponse) SetResults(v []SecretManager) {
	o.Results = v
}

func (o GetSecretManagerListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSecretManagerListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableGetSecretManagerListResponse struct {
	value *GetSecretManagerListResponse
	isSet bool
}

func (v NullableGetSecretManagerListResponse) Get() *GetSecretManagerListResponse {
	return v.value
}

func (v *NullableGetSecretManagerListResponse) Set(val *GetSecretManagerListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSecretManagerListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSecretManagerListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSecretManagerListResponse(val *GetSecretManagerListResponse) *NullableGetSecretManagerListResponse {
	return &NullableGetSecretManagerListResponse{value: val, isSet: true}
}

func (v NullableGetSecretManagerListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSecretManagerListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
