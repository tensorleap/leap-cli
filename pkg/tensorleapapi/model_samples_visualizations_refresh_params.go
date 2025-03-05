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

// checks if the SamplesVisualizationsRefreshParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamplesVisualizationsRefreshParams{}

// SamplesVisualizationsRefreshParams struct for SamplesVisualizationsRefreshParams
type SamplesVisualizationsRefreshParams struct {
	ProjectId    string `json:"projectId"`
	SessionRunId string `json:"sessionRunId"`
}

// NewSamplesVisualizationsRefreshParams instantiates a new SamplesVisualizationsRefreshParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamplesVisualizationsRefreshParams(projectId string, sessionRunId string) *SamplesVisualizationsRefreshParams {
	this := SamplesVisualizationsRefreshParams{}
	this.ProjectId = projectId
	this.SessionRunId = sessionRunId
	return &this
}

// NewSamplesVisualizationsRefreshParamsWithDefaults instantiates a new SamplesVisualizationsRefreshParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamplesVisualizationsRefreshParamsWithDefaults() *SamplesVisualizationsRefreshParams {
	this := SamplesVisualizationsRefreshParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *SamplesVisualizationsRefreshParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *SamplesVisualizationsRefreshParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *SamplesVisualizationsRefreshParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSessionRunId returns the SessionRunId field value
func (o *SamplesVisualizationsRefreshParams) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *SamplesVisualizationsRefreshParams) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *SamplesVisualizationsRefreshParams) SetSessionRunId(v string) {
	o.SessionRunId = v
}

func (o SamplesVisualizationsRefreshParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamplesVisualizationsRefreshParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["sessionRunId"] = o.SessionRunId
	return toSerialize, nil
}

type NullableSamplesVisualizationsRefreshParams struct {
	value *SamplesVisualizationsRefreshParams
	isSet bool
}

func (v NullableSamplesVisualizationsRefreshParams) Get() *SamplesVisualizationsRefreshParams {
	return v.value
}

func (v *NullableSamplesVisualizationsRefreshParams) Set(val *SamplesVisualizationsRefreshParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSamplesVisualizationsRefreshParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSamplesVisualizationsRefreshParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamplesVisualizationsRefreshParams(val *SamplesVisualizationsRefreshParams) *NullableSamplesVisualizationsRefreshParams {
	return &NullableSamplesVisualizationsRefreshParams{value: val, isSet: true}
}

func (v NullableSamplesVisualizationsRefreshParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamplesVisualizationsRefreshParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
