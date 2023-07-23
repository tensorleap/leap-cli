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

// checks if the LoadVersionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadVersionParams{}

// LoadVersionParams struct for LoadVersionParams
type LoadVersionParams struct {
	VersionId string `json:"versionId"`
	ProjectId string `json:"projectId"`
}

// NewLoadVersionParams instantiates a new LoadVersionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadVersionParams(versionId string, projectId string) *LoadVersionParams {
	this := LoadVersionParams{}
	this.VersionId = versionId
	this.ProjectId = projectId
	return &this
}

// NewLoadVersionParamsWithDefaults instantiates a new LoadVersionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadVersionParamsWithDefaults() *LoadVersionParams {
	this := LoadVersionParams{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *LoadVersionParams) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *LoadVersionParams) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *LoadVersionParams) SetVersionId(v string) {
	o.VersionId = v
}

// GetProjectId returns the ProjectId field value
func (o *LoadVersionParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *LoadVersionParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *LoadVersionParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o LoadVersionParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadVersionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableLoadVersionParams struct {
	value *LoadVersionParams
	isSet bool
}

func (v NullableLoadVersionParams) Get() *LoadVersionParams {
	return v.value
}

func (v *NullableLoadVersionParams) Set(val *LoadVersionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadVersionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadVersionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadVersionParams(val *LoadVersionParams) *NullableLoadVersionParams {
	return &NullableLoadVersionParams{value: val, isSet: true}
}

func (v NullableLoadVersionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadVersionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
