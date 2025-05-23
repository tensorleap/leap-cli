/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetProjectDashboardsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectDashboardsParams{}

// GetProjectDashboardsParams struct for GetProjectDashboardsParams
type GetProjectDashboardsParams struct {
	ProjectId string `json:"projectId"`
}

// NewGetProjectDashboardsParams instantiates a new GetProjectDashboardsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectDashboardsParams(projectId string) *GetProjectDashboardsParams {
	this := GetProjectDashboardsParams{}
	this.ProjectId = projectId
	return &this
}

// NewGetProjectDashboardsParamsWithDefaults instantiates a new GetProjectDashboardsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectDashboardsParamsWithDefaults() *GetProjectDashboardsParams {
	this := GetProjectDashboardsParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GetProjectDashboardsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetProjectDashboardsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetProjectDashboardsParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o GetProjectDashboardsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectDashboardsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableGetProjectDashboardsParams struct {
	value *GetProjectDashboardsParams
	isSet bool
}

func (v NullableGetProjectDashboardsParams) Get() *GetProjectDashboardsParams {
	return v.value
}

func (v *NullableGetProjectDashboardsParams) Set(val *GetProjectDashboardsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectDashboardsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectDashboardsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectDashboardsParams(val *GetProjectDashboardsParams) *NullableGetProjectDashboardsParams {
	return &NullableGetProjectDashboardsParams{value: val, isSet: true}
}

func (v NullableGetProjectDashboardsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectDashboardsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
