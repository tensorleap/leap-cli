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

// checks if the GetDashboardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDashboardResponse{}

// GetDashboardResponse struct for GetDashboardResponse
type GetDashboardResponse struct {
	Dashboard Dashboard `json:"dashboard"`
}

// NewGetDashboardResponse instantiates a new GetDashboardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDashboardResponse(dashboard Dashboard) *GetDashboardResponse {
	this := GetDashboardResponse{}
	this.Dashboard = dashboard
	return &this
}

// NewGetDashboardResponseWithDefaults instantiates a new GetDashboardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDashboardResponseWithDefaults() *GetDashboardResponse {
	this := GetDashboardResponse{}
	return &this
}

// GetDashboard returns the Dashboard field value
func (o *GetDashboardResponse) GetDashboard() Dashboard {
	if o == nil {
		var ret Dashboard
		return ret
	}

	return o.Dashboard
}

// GetDashboardOk returns a tuple with the Dashboard field value
// and a boolean to check if the value has been set.
func (o *GetDashboardResponse) GetDashboardOk() (*Dashboard, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dashboard, true
}

// SetDashboard sets field value
func (o *GetDashboardResponse) SetDashboard(v Dashboard) {
	o.Dashboard = v
}

func (o GetDashboardResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDashboardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dashboard"] = o.Dashboard
	return toSerialize, nil
}

type NullableGetDashboardResponse struct {
	value *GetDashboardResponse
	isSet bool
}

func (v NullableGetDashboardResponse) Get() *GetDashboardResponse {
	return v.value
}

func (v *NullableGetDashboardResponse) Set(val *GetDashboardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDashboardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDashboardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDashboardResponse(val *GetDashboardResponse) *NullableGetDashboardResponse {
	return &NullableGetDashboardResponse{value: val, isSet: true}
}

func (v NullableGetDashboardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDashboardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
