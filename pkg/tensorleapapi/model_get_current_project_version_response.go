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

// checks if the GetCurrentProjectVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCurrentProjectVersionResponse{}

// GetCurrentProjectVersionResponse struct for GetCurrentProjectVersionResponse
type GetCurrentProjectVersionResponse struct {
	ProjectId string `json:"projectId"`
	VersionId string `json:"versionId"`
}

// NewGetCurrentProjectVersionResponse instantiates a new GetCurrentProjectVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCurrentProjectVersionResponse(projectId string, versionId string) *GetCurrentProjectVersionResponse {
	this := GetCurrentProjectVersionResponse{}
	this.ProjectId = projectId
	this.VersionId = versionId
	return &this
}

// NewGetCurrentProjectVersionResponseWithDefaults instantiates a new GetCurrentProjectVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCurrentProjectVersionResponseWithDefaults() *GetCurrentProjectVersionResponse {
	this := GetCurrentProjectVersionResponse{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GetCurrentProjectVersionResponse) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetCurrentProjectVersionResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetCurrentProjectVersionResponse) SetProjectId(v string) {
	o.ProjectId = v
}

// GetVersionId returns the VersionId field value
func (o *GetCurrentProjectVersionResponse) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *GetCurrentProjectVersionResponse) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *GetCurrentProjectVersionResponse) SetVersionId(v string) {
	o.VersionId = v
}

func (o GetCurrentProjectVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCurrentProjectVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["versionId"] = o.VersionId
	return toSerialize, nil
}

type NullableGetCurrentProjectVersionResponse struct {
	value *GetCurrentProjectVersionResponse
	isSet bool
}

func (v NullableGetCurrentProjectVersionResponse) Get() *GetCurrentProjectVersionResponse {
	return v.value
}

func (v *NullableGetCurrentProjectVersionResponse) Set(val *GetCurrentProjectVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCurrentProjectVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCurrentProjectVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCurrentProjectVersionResponse(val *GetCurrentProjectVersionResponse) *NullableGetCurrentProjectVersionResponse {
	return &NullableGetCurrentProjectVersionResponse{value: val, isSet: true}
}

func (v NullableGetCurrentProjectVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCurrentProjectVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
