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

// checks if the GetExportedSessionRunJobsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetExportedSessionRunJobsParams{}

// GetExportedSessionRunJobsParams struct for GetExportedSessionRunJobsParams
type GetExportedSessionRunJobsParams struct {
	SessionId string `json:"sessionId"`
	ProjectId string `json:"projectId"`
}

// NewGetExportedSessionRunJobsParams instantiates a new GetExportedSessionRunJobsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExportedSessionRunJobsParams(sessionId string, projectId string) *GetExportedSessionRunJobsParams {
	this := GetExportedSessionRunJobsParams{}
	this.SessionId = sessionId
	this.ProjectId = projectId
	return &this
}

// NewGetExportedSessionRunJobsParamsWithDefaults instantiates a new GetExportedSessionRunJobsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExportedSessionRunJobsParamsWithDefaults() *GetExportedSessionRunJobsParams {
	this := GetExportedSessionRunJobsParams{}
	return &this
}

// GetSessionId returns the SessionId field value
func (o *GetExportedSessionRunJobsParams) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *GetExportedSessionRunJobsParams) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *GetExportedSessionRunJobsParams) SetSessionId(v string) {
	o.SessionId = v
}

// GetProjectId returns the ProjectId field value
func (o *GetExportedSessionRunJobsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetExportedSessionRunJobsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetExportedSessionRunJobsParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o GetExportedSessionRunJobsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExportedSessionRunJobsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionId"] = o.SessionId
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableGetExportedSessionRunJobsParams struct {
	value *GetExportedSessionRunJobsParams
	isSet bool
}

func (v NullableGetExportedSessionRunJobsParams) Get() *GetExportedSessionRunJobsParams {
	return v.value
}

func (v *NullableGetExportedSessionRunJobsParams) Set(val *GetExportedSessionRunJobsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExportedSessionRunJobsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExportedSessionRunJobsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExportedSessionRunJobsParams(val *GetExportedSessionRunJobsParams) *NullableGetExportedSessionRunJobsParams {
	return &NullableGetExportedSessionRunJobsParams{value: val, isSet: true}
}

func (v NullableGetExportedSessionRunJobsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExportedSessionRunJobsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
