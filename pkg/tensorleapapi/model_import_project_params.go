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

// checks if the ImportProjectParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportProjectParams{}

// ImportProjectParams struct for ImportProjectParams
type ImportProjectParams struct {
	ImportUrl   string      `json:"importUrl"`
	ProjectMeta ProjectMeta `json:"projectMeta"`
}

// NewImportProjectParams instantiates a new ImportProjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportProjectParams(importUrl string, projectMeta ProjectMeta) *ImportProjectParams {
	this := ImportProjectParams{}
	this.ImportUrl = importUrl
	this.ProjectMeta = projectMeta
	return &this
}

// NewImportProjectParamsWithDefaults instantiates a new ImportProjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportProjectParamsWithDefaults() *ImportProjectParams {
	this := ImportProjectParams{}
	return &this
}

// GetImportUrl returns the ImportUrl field value
func (o *ImportProjectParams) GetImportUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImportUrl
}

// GetImportUrlOk returns a tuple with the ImportUrl field value
// and a boolean to check if the value has been set.
func (o *ImportProjectParams) GetImportUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportUrl, true
}

// SetImportUrl sets field value
func (o *ImportProjectParams) SetImportUrl(v string) {
	o.ImportUrl = v
}

// GetProjectMeta returns the ProjectMeta field value
func (o *ImportProjectParams) GetProjectMeta() ProjectMeta {
	if o == nil {
		var ret ProjectMeta
		return ret
	}

	return o.ProjectMeta
}

// GetProjectMetaOk returns a tuple with the ProjectMeta field value
// and a boolean to check if the value has been set.
func (o *ImportProjectParams) GetProjectMetaOk() (*ProjectMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectMeta, true
}

// SetProjectMeta sets field value
func (o *ImportProjectParams) SetProjectMeta(v ProjectMeta) {
	o.ProjectMeta = v
}

func (o ImportProjectParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportProjectParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["importUrl"] = o.ImportUrl
	toSerialize["projectMeta"] = o.ProjectMeta
	return toSerialize, nil
}

type NullableImportProjectParams struct {
	value *ImportProjectParams
	isSet bool
}

func (v NullableImportProjectParams) Get() *ImportProjectParams {
	return v.value
}

func (v *NullableImportProjectParams) Set(val *ImportProjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableImportProjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableImportProjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportProjectParams(val *ImportProjectParams) *NullableImportProjectParams {
	return &NullableImportProjectParams{value: val, isSet: true}
}

func (v NullableImportProjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportProjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}