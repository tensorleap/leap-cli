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

// checks if the GetStatisticsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatisticsParams{}

// GetStatisticsParams struct for GetStatisticsParams
type GetStatisticsParams struct {
	ProjectId string `json:"projectId"`
}

// NewGetStatisticsParams instantiates a new GetStatisticsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatisticsParams(projectId string) *GetStatisticsParams {
	this := GetStatisticsParams{}
	this.ProjectId = projectId
	return &this
}

// NewGetStatisticsParamsWithDefaults instantiates a new GetStatisticsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatisticsParamsWithDefaults() *GetStatisticsParams {
	this := GetStatisticsParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GetStatisticsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetStatisticsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetStatisticsParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o GetStatisticsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatisticsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableGetStatisticsParams struct {
	value *GetStatisticsParams
	isSet bool
}

func (v NullableGetStatisticsParams) Get() *GetStatisticsParams {
	return v.value
}

func (v *NullableGetStatisticsParams) Set(val *GetStatisticsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatisticsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatisticsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatisticsParams(val *GetStatisticsParams) *NullableGetStatisticsParams {
	return &NullableGetStatisticsParams{value: val, isSet: true}
}

func (v NullableGetStatisticsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatisticsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
