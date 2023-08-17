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

// checks if the DeleteDashboardParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteDashboardParams{}

// DeleteDashboardParams struct for DeleteDashboardParams
type DeleteDashboardParams struct {
	DashboardId string `json:"dashboardId"`
	ProjectId string `json:"projectId"`
}

// NewDeleteDashboardParams instantiates a new DeleteDashboardParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDashboardParams(dashboardId string, projectId string) *DeleteDashboardParams {
	this := DeleteDashboardParams{}
	this.DashboardId = dashboardId
	this.ProjectId = projectId
	return &this
}

// NewDeleteDashboardParamsWithDefaults instantiates a new DeleteDashboardParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDashboardParamsWithDefaults() *DeleteDashboardParams {
	this := DeleteDashboardParams{}
	return &this
}

// GetDashboardId returns the DashboardId field value
func (o *DeleteDashboardParams) GetDashboardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DashboardId
}

// GetDashboardIdOk returns a tuple with the DashboardId field value
// and a boolean to check if the value has been set.
func (o *DeleteDashboardParams) GetDashboardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DashboardId, true
}

// SetDashboardId sets field value
func (o *DeleteDashboardParams) SetDashboardId(v string) {
	o.DashboardId = v
}

// GetProjectId returns the ProjectId field value
func (o *DeleteDashboardParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DeleteDashboardParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DeleteDashboardParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o DeleteDashboardParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteDashboardParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dashboardId"] = o.DashboardId
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableDeleteDashboardParams struct {
	value *DeleteDashboardParams
	isSet bool
}

func (v NullableDeleteDashboardParams) Get() *DeleteDashboardParams {
	return v.value
}

func (v *NullableDeleteDashboardParams) Set(val *DeleteDashboardParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDashboardParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDashboardParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDashboardParams(val *DeleteDashboardParams) *NullableDeleteDashboardParams {
	return &NullableDeleteDashboardParams{value: val, isSet: true}
}

func (v NullableDeleteDashboardParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDashboardParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


