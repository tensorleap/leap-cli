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

// checks if the ExternalImportModelStorageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalImportModelStorageResponse{}

// ExternalImportModelStorageResponse struct for ExternalImportModelStorageResponse
type ExternalImportModelStorageResponse struct {
	ImportModelJobId string `json:"importModelJobId"`
}

// NewExternalImportModelStorageResponse instantiates a new ExternalImportModelStorageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalImportModelStorageResponse(importModelJobId string) *ExternalImportModelStorageResponse {
	this := ExternalImportModelStorageResponse{}
	this.ImportModelJobId = importModelJobId
	return &this
}

// NewExternalImportModelStorageResponseWithDefaults instantiates a new ExternalImportModelStorageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalImportModelStorageResponseWithDefaults() *ExternalImportModelStorageResponse {
	this := ExternalImportModelStorageResponse{}
	return &this
}

// GetImportModelJobId returns the ImportModelJobId field value
func (o *ExternalImportModelStorageResponse) GetImportModelJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImportModelJobId
}

// GetImportModelJobIdOk returns a tuple with the ImportModelJobId field value
// and a boolean to check if the value has been set.
func (o *ExternalImportModelStorageResponse) GetImportModelJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportModelJobId, true
}

// SetImportModelJobId sets field value
func (o *ExternalImportModelStorageResponse) SetImportModelJobId(v string) {
	o.ImportModelJobId = v
}

func (o ExternalImportModelStorageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalImportModelStorageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["importModelJobId"] = o.ImportModelJobId
	return toSerialize, nil
}

type NullableExternalImportModelStorageResponse struct {
	value *ExternalImportModelStorageResponse
	isSet bool
}

func (v NullableExternalImportModelStorageResponse) Get() *ExternalImportModelStorageResponse {
	return v.value
}

func (v *NullableExternalImportModelStorageResponse) Set(val *ExternalImportModelStorageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalImportModelStorageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalImportModelStorageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalImportModelStorageResponse(val *ExternalImportModelStorageResponse) *NullableExternalImportModelStorageResponse {
	return &NullableExternalImportModelStorageResponse{value: val, isSet: true}
}

func (v NullableExternalImportModelStorageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalImportModelStorageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
