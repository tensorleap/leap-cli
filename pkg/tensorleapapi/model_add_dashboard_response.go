/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.417
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the AddDashboardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddDashboardResponse{}

// AddDashboardResponse struct for AddDashboardResponse
type AddDashboardResponse struct {
	DashboardId string `json:"dashboardId"`
}

// NewAddDashboardResponse instantiates a new AddDashboardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDashboardResponse(dashboardId string) *AddDashboardResponse {
	this := AddDashboardResponse{}
	this.DashboardId = dashboardId
	return &this
}

// NewAddDashboardResponseWithDefaults instantiates a new AddDashboardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDashboardResponseWithDefaults() *AddDashboardResponse {
	this := AddDashboardResponse{}
	return &this
}

// GetDashboardId returns the DashboardId field value
func (o *AddDashboardResponse) GetDashboardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DashboardId
}

// GetDashboardIdOk returns a tuple with the DashboardId field value
// and a boolean to check if the value has been set.
func (o *AddDashboardResponse) GetDashboardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DashboardId, true
}

// SetDashboardId sets field value
func (o *AddDashboardResponse) SetDashboardId(v string) {
	o.DashboardId = v
}

func (o AddDashboardResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddDashboardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dashboardId"] = o.DashboardId
	return toSerialize, nil
}

type NullableAddDashboardResponse struct {
	value *AddDashboardResponse
	isSet bool
}

func (v NullableAddDashboardResponse) Get() *AddDashboardResponse {
	return v.value
}

func (v *NullableAddDashboardResponse) Set(val *AddDashboardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDashboardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDashboardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDashboardResponse(val *AddDashboardResponse) *NullableAddDashboardResponse {
	return &NullableAddDashboardResponse{value: val, isSet: true}
}

func (v NullableAddDashboardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDashboardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
