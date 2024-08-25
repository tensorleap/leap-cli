/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.526
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the TagModelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagModelRequest{}

// TagModelRequest struct for TagModelRequest
type TagModelRequest struct {
	ProjectId    string   `json:"projectId"`
	ExperimentId string   `json:"experimentId"`
	Tags         []string `json:"tags"`
	Epoch        float64  `json:"epoch"`
}

// NewTagModelRequest instantiates a new TagModelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagModelRequest(projectId string, experimentId string, tags []string, epoch float64) *TagModelRequest {
	this := TagModelRequest{}
	this.ProjectId = projectId
	this.ExperimentId = experimentId
	this.Tags = tags
	this.Epoch = epoch
	return &this
}

// NewTagModelRequestWithDefaults instantiates a new TagModelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagModelRequestWithDefaults() *TagModelRequest {
	this := TagModelRequest{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *TagModelRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *TagModelRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *TagModelRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetExperimentId returns the ExperimentId field value
func (o *TagModelRequest) GetExperimentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExperimentId
}

// GetExperimentIdOk returns a tuple with the ExperimentId field value
// and a boolean to check if the value has been set.
func (o *TagModelRequest) GetExperimentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExperimentId, true
}

// SetExperimentId sets field value
func (o *TagModelRequest) SetExperimentId(v string) {
	o.ExperimentId = v
}

// GetTags returns the Tags field value
func (o *TagModelRequest) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *TagModelRequest) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *TagModelRequest) SetTags(v []string) {
	o.Tags = v
}

// GetEpoch returns the Epoch field value
func (o *TagModelRequest) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *TagModelRequest) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *TagModelRequest) SetEpoch(v float64) {
	o.Epoch = v
}

func (o TagModelRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagModelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["experimentId"] = o.ExperimentId
	toSerialize["tags"] = o.Tags
	toSerialize["epoch"] = o.Epoch
	return toSerialize, nil
}

type NullableTagModelRequest struct {
	value *TagModelRequest
	isSet bool
}

func (v NullableTagModelRequest) Get() *TagModelRequest {
	return v.value
}

func (v *NullableTagModelRequest) Set(val *TagModelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTagModelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTagModelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagModelRequest(val *TagModelRequest) *NullableTagModelRequest {
	return &NullableTagModelRequest{value: val, isSet: true}
}

func (v NullableTagModelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagModelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
