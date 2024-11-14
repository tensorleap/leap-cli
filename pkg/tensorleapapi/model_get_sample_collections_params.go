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

// checks if the GetSampleCollectionsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSampleCollectionsParams{}

// GetSampleCollectionsParams struct for GetSampleCollectionsParams
type GetSampleCollectionsParams struct {
	ProjectId string `json:"projectId"`
}

// NewGetSampleCollectionsParams instantiates a new GetSampleCollectionsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSampleCollectionsParams(projectId string) *GetSampleCollectionsParams {
	this := GetSampleCollectionsParams{}
	this.ProjectId = projectId
	return &this
}

// NewGetSampleCollectionsParamsWithDefaults instantiates a new GetSampleCollectionsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSampleCollectionsParamsWithDefaults() *GetSampleCollectionsParams {
	this := GetSampleCollectionsParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GetSampleCollectionsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetSampleCollectionsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetSampleCollectionsParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o GetSampleCollectionsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSampleCollectionsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableGetSampleCollectionsParams struct {
	value *GetSampleCollectionsParams
	isSet bool
}

func (v NullableGetSampleCollectionsParams) Get() *GetSampleCollectionsParams {
	return v.value
}

func (v *NullableGetSampleCollectionsParams) Set(val *GetSampleCollectionsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSampleCollectionsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSampleCollectionsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSampleCollectionsParams(val *GetSampleCollectionsParams) *NullableGetSampleCollectionsParams {
	return &NullableGetSampleCollectionsParams{value: val, isSet: true}
}

func (v NullableGetSampleCollectionsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSampleCollectionsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
