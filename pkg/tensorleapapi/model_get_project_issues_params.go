/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetProjectIssuesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectIssuesParams{}

// GetProjectIssuesParams struct for GetProjectIssuesParams
type GetProjectIssuesParams struct {
	ProjectId string `json:"projectId"`
}

// NewGetProjectIssuesParams instantiates a new GetProjectIssuesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectIssuesParams(projectId string) *GetProjectIssuesParams {
	this := GetProjectIssuesParams{}
	this.ProjectId = projectId
	return &this
}

// NewGetProjectIssuesParamsWithDefaults instantiates a new GetProjectIssuesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectIssuesParamsWithDefaults() *GetProjectIssuesParams {
	this := GetProjectIssuesParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GetProjectIssuesParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetProjectIssuesParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetProjectIssuesParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o GetProjectIssuesParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectIssuesParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableGetProjectIssuesParams struct {
	value *GetProjectIssuesParams
	isSet bool
}

func (v NullableGetProjectIssuesParams) Get() *GetProjectIssuesParams {
	return v.value
}

func (v *NullableGetProjectIssuesParams) Set(val *GetProjectIssuesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectIssuesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectIssuesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectIssuesParams(val *GetProjectIssuesParams) *NullableGetProjectIssuesParams {
	return &NullableGetProjectIssuesParams{value: val, isSet: true}
}

func (v NullableGetProjectIssuesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectIssuesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
