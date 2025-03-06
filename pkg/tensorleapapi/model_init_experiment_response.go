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

// checks if the InitExperimentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitExperimentResponse{}

// InitExperimentResponse struct for InitExperimentResponse
type InitExperimentResponse struct {
	ExperimentId     string `json:"experimentId"`
	ProjectId        string `json:"projectId"`
	VersionId        string `json:"versionId"`
	IsCreatedProject bool   `json:"isCreatedProject"`
}

// NewInitExperimentResponse instantiates a new InitExperimentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitExperimentResponse(experimentId string, projectId string, versionId string, isCreatedProject bool) *InitExperimentResponse {
	this := InitExperimentResponse{}
	this.ExperimentId = experimentId
	this.ProjectId = projectId
	this.VersionId = versionId
	this.IsCreatedProject = isCreatedProject
	return &this
}

// NewInitExperimentResponseWithDefaults instantiates a new InitExperimentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitExperimentResponseWithDefaults() *InitExperimentResponse {
	this := InitExperimentResponse{}
	return &this
}

// GetExperimentId returns the ExperimentId field value
func (o *InitExperimentResponse) GetExperimentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExperimentId
}

// GetExperimentIdOk returns a tuple with the ExperimentId field value
// and a boolean to check if the value has been set.
func (o *InitExperimentResponse) GetExperimentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExperimentId, true
}

// SetExperimentId sets field value
func (o *InitExperimentResponse) SetExperimentId(v string) {
	o.ExperimentId = v
}

// GetProjectId returns the ProjectId field value
func (o *InitExperimentResponse) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *InitExperimentResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *InitExperimentResponse) SetProjectId(v string) {
	o.ProjectId = v
}

// GetVersionId returns the VersionId field value
func (o *InitExperimentResponse) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *InitExperimentResponse) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *InitExperimentResponse) SetVersionId(v string) {
	o.VersionId = v
}

// GetIsCreatedProject returns the IsCreatedProject field value
func (o *InitExperimentResponse) GetIsCreatedProject() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCreatedProject
}

// GetIsCreatedProjectOk returns a tuple with the IsCreatedProject field value
// and a boolean to check if the value has been set.
func (o *InitExperimentResponse) GetIsCreatedProjectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCreatedProject, true
}

// SetIsCreatedProject sets field value
func (o *InitExperimentResponse) SetIsCreatedProject(v bool) {
	o.IsCreatedProject = v
}

func (o InitExperimentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitExperimentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["experimentId"] = o.ExperimentId
	toSerialize["projectId"] = o.ProjectId
	toSerialize["versionId"] = o.VersionId
	toSerialize["isCreatedProject"] = o.IsCreatedProject
	return toSerialize, nil
}

type NullableInitExperimentResponse struct {
	value *InitExperimentResponse
	isSet bool
}

func (v NullableInitExperimentResponse) Get() *InitExperimentResponse {
	return v.value
}

func (v *NullableInitExperimentResponse) Set(val *InitExperimentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInitExperimentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInitExperimentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitExperimentResponse(val *InitExperimentResponse) *NullableInitExperimentResponse {
	return &NullableInitExperimentResponse{value: val, isSet: true}
}

func (v NullableInitExperimentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitExperimentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
