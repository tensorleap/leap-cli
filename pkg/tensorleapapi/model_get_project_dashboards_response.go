/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.445
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetProjectDashboardsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectDashboardsResponse{}

// GetProjectDashboardsResponse struct for GetProjectDashboardsResponse
type GetProjectDashboardsResponse struct {
	Dashboards []Dashboard `json:"dashboards"`
}

// NewGetProjectDashboardsResponse instantiates a new GetProjectDashboardsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectDashboardsResponse(dashboards []Dashboard) *GetProjectDashboardsResponse {
	this := GetProjectDashboardsResponse{}
	this.Dashboards = dashboards
	return &this
}

// NewGetProjectDashboardsResponseWithDefaults instantiates a new GetProjectDashboardsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectDashboardsResponseWithDefaults() *GetProjectDashboardsResponse {
	this := GetProjectDashboardsResponse{}
	return &this
}

// GetDashboards returns the Dashboards field value
func (o *GetProjectDashboardsResponse) GetDashboards() []Dashboard {
	if o == nil {
		var ret []Dashboard
		return ret
	}

	return o.Dashboards
}

// GetDashboardsOk returns a tuple with the Dashboards field value
// and a boolean to check if the value has been set.
func (o *GetProjectDashboardsResponse) GetDashboardsOk() ([]Dashboard, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dashboards, true
}

// SetDashboards sets field value
func (o *GetProjectDashboardsResponse) SetDashboards(v []Dashboard) {
	o.Dashboards = v
}

func (o GetProjectDashboardsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectDashboardsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dashboards"] = o.Dashboards
	return toSerialize, nil
}

type NullableGetProjectDashboardsResponse struct {
	value *GetProjectDashboardsResponse
	isSet bool
}

func (v NullableGetProjectDashboardsResponse) Get() *GetProjectDashboardsResponse {
	return v.value
}

func (v *NullableGetProjectDashboardsResponse) Set(val *GetProjectDashboardsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectDashboardsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectDashboardsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectDashboardsResponse(val *GetProjectDashboardsResponse) *NullableGetProjectDashboardsResponse {
	return &NullableGetProjectDashboardsResponse{value: val, isSet: true}
}

func (v NullableGetProjectDashboardsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectDashboardsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
