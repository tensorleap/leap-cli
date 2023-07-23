/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetDashletFieldsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDashletFieldsParams{}

// GetDashletFieldsParams struct for GetDashletFieldsParams
type GetDashletFieldsParams struct {
	ProjectId string `json:"projectId"`
}

// NewGetDashletFieldsParams instantiates a new GetDashletFieldsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDashletFieldsParams(projectId string) *GetDashletFieldsParams {
	this := GetDashletFieldsParams{}
	this.ProjectId = projectId
	return &this
}

// NewGetDashletFieldsParamsWithDefaults instantiates a new GetDashletFieldsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDashletFieldsParamsWithDefaults() *GetDashletFieldsParams {
	this := GetDashletFieldsParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GetDashletFieldsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetDashletFieldsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetDashletFieldsParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o GetDashletFieldsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDashletFieldsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableGetDashletFieldsParams struct {
	value *GetDashletFieldsParams
	isSet bool
}

func (v NullableGetDashletFieldsParams) Get() *GetDashletFieldsParams {
	return v.value
}

func (v *NullableGetDashletFieldsParams) Set(val *GetDashletFieldsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDashletFieldsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDashletFieldsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDashletFieldsParams(val *GetDashletFieldsParams) *NullableGetDashletFieldsParams {
	return &NullableGetDashletFieldsParams{value: val, isSet: true}
}

func (v NullableGetDashletFieldsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDashletFieldsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
