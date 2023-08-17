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

// checks if the GetValidateGraphProcessStateParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetValidateGraphProcessStateParams{}

// GetValidateGraphProcessStateParams struct for GetValidateGraphProcessStateParams
type GetValidateGraphProcessStateParams struct {
	ProjectId string `json:"projectId"`
	Digest    string `json:"digest"`
}

// NewGetValidateGraphProcessStateParams instantiates a new GetValidateGraphProcessStateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetValidateGraphProcessStateParams(projectId string, digest string) *GetValidateGraphProcessStateParams {
	this := GetValidateGraphProcessStateParams{}
	this.ProjectId = projectId
	this.Digest = digest
	return &this
}

// NewGetValidateGraphProcessStateParamsWithDefaults instantiates a new GetValidateGraphProcessStateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetValidateGraphProcessStateParamsWithDefaults() *GetValidateGraphProcessStateParams {
	this := GetValidateGraphProcessStateParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GetValidateGraphProcessStateParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetValidateGraphProcessStateParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetValidateGraphProcessStateParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetDigest returns the Digest field value
func (o *GetValidateGraphProcessStateParams) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *GetValidateGraphProcessStateParams) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *GetValidateGraphProcessStateParams) SetDigest(v string) {
	o.Digest = v
}

func (o GetValidateGraphProcessStateParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetValidateGraphProcessStateParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["digest"] = o.Digest
	return toSerialize, nil
}

type NullableGetValidateGraphProcessStateParams struct {
	value *GetValidateGraphProcessStateParams
	isSet bool
}

func (v NullableGetValidateGraphProcessStateParams) Get() *GetValidateGraphProcessStateParams {
	return v.value
}

func (v *NullableGetValidateGraphProcessStateParams) Set(val *GetValidateGraphProcessStateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetValidateGraphProcessStateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetValidateGraphProcessStateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetValidateGraphProcessStateParams(val *GetValidateGraphProcessStateParams) *NullableGetValidateGraphProcessStateParams {
	return &NullableGetValidateGraphProcessStateParams{value: val, isSet: true}
}

func (v NullableGetValidateGraphProcessStateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetValidateGraphProcessStateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
