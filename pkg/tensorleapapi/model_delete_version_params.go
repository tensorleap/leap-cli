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

// checks if the DeleteVersionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteVersionParams{}

// DeleteVersionParams struct for DeleteVersionParams
type DeleteVersionParams struct {
	VersionId string `json:"versionId"`
	ProjectId string `json:"projectId"`
}

// NewDeleteVersionParams instantiates a new DeleteVersionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVersionParams(versionId string, projectId string) *DeleteVersionParams {
	this := DeleteVersionParams{}
	this.VersionId = versionId
	this.ProjectId = projectId
	return &this
}

// NewDeleteVersionParamsWithDefaults instantiates a new DeleteVersionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVersionParamsWithDefaults() *DeleteVersionParams {
	this := DeleteVersionParams{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *DeleteVersionParams) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *DeleteVersionParams) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *DeleteVersionParams) SetVersionId(v string) {
	o.VersionId = v
}

// GetProjectId returns the ProjectId field value
func (o *DeleteVersionParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DeleteVersionParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DeleteVersionParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o DeleteVersionParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteVersionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableDeleteVersionParams struct {
	value *DeleteVersionParams
	isSet bool
}

func (v NullableDeleteVersionParams) Get() *DeleteVersionParams {
	return v.value
}

func (v *NullableDeleteVersionParams) Set(val *DeleteVersionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVersionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVersionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVersionParams(val *DeleteVersionParams) *NullableDeleteVersionParams {
	return &NullableDeleteVersionParams{value: val, isSet: true}
}

func (v NullableDeleteVersionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVersionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
