/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UpdateVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVersionResponse{}

// UpdateVersionResponse struct for UpdateVersionResponse
type UpdateVersionResponse struct {
	VersionId string `json:"versionId"`
}

// NewUpdateVersionResponse instantiates a new UpdateVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVersionResponse(versionId string) *UpdateVersionResponse {
	this := UpdateVersionResponse{}
	this.VersionId = versionId
	return &this
}

// NewUpdateVersionResponseWithDefaults instantiates a new UpdateVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVersionResponseWithDefaults() *UpdateVersionResponse {
	this := UpdateVersionResponse{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *UpdateVersionResponse) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *UpdateVersionResponse) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *UpdateVersionResponse) SetVersionId(v string) {
	o.VersionId = v
}

func (o UpdateVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	return toSerialize, nil
}

type NullableUpdateVersionResponse struct {
	value *UpdateVersionResponse
	isSet bool
}

func (v NullableUpdateVersionResponse) Get() *UpdateVersionResponse {
	return v.value
}

func (v *NullableUpdateVersionResponse) Set(val *UpdateVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVersionResponse(val *UpdateVersionResponse) *NullableUpdateVersionResponse {
	return &NullableUpdateVersionResponse{value: val, isSet: true}
}

func (v NullableUpdateVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


