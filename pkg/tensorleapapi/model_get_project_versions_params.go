/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.301
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetProjectVersionsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectVersionsParams{}

// GetProjectVersionsParams struct for GetProjectVersionsParams
type GetProjectVersionsParams struct {
	ProjectId string `json:"projectId"`
}

// NewGetProjectVersionsParams instantiates a new GetProjectVersionsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectVersionsParams(projectId string) *GetProjectVersionsParams {
	this := GetProjectVersionsParams{}
	this.ProjectId = projectId
	return &this
}

// NewGetProjectVersionsParamsWithDefaults instantiates a new GetProjectVersionsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectVersionsParamsWithDefaults() *GetProjectVersionsParams {
	this := GetProjectVersionsParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GetProjectVersionsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetProjectVersionsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetProjectVersionsParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o GetProjectVersionsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectVersionsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableGetProjectVersionsParams struct {
	value *GetProjectVersionsParams
	isSet bool
}

func (v NullableGetProjectVersionsParams) Get() *GetProjectVersionsParams {
	return v.value
}

func (v *NullableGetProjectVersionsParams) Set(val *GetProjectVersionsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectVersionsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectVersionsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectVersionsParams(val *GetProjectVersionsParams) *NullableGetProjectVersionsParams {
	return &NullableGetProjectVersionsParams{value: val, isSet: true}
}

func (v NullableGetProjectVersionsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectVersionsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
