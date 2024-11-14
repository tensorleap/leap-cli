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

// checks if the DeleteProjectParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteProjectParams{}

// DeleteProjectParams struct for DeleteProjectParams
type DeleteProjectParams struct {
	ProjectId string `json:"projectId"`
}

// NewDeleteProjectParams instantiates a new DeleteProjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteProjectParams(projectId string) *DeleteProjectParams {
	this := DeleteProjectParams{}
	this.ProjectId = projectId
	return &this
}

// NewDeleteProjectParamsWithDefaults instantiates a new DeleteProjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteProjectParamsWithDefaults() *DeleteProjectParams {
	this := DeleteProjectParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *DeleteProjectParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DeleteProjectParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DeleteProjectParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o DeleteProjectParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteProjectParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableDeleteProjectParams struct {
	value *DeleteProjectParams
	isSet bool
}

func (v NullableDeleteProjectParams) Get() *DeleteProjectParams {
	return v.value
}

func (v *NullableDeleteProjectParams) Set(val *DeleteProjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteProjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteProjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteProjectParams(val *DeleteProjectParams) *NullableDeleteProjectParams {
	return &NullableDeleteProjectParams{value: val, isSet: true}
}

func (v NullableDeleteProjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteProjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
