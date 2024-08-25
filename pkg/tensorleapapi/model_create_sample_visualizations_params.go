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

// checks if the CreateSampleVisualizationsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSampleVisualizationsParams{}

// CreateSampleVisualizationsParams struct for CreateSampleVisualizationsParams
type CreateSampleVisualizationsParams struct {
	ProjectId        string           `json:"projectId"`
	SessionRunId     string           `json:"sessionRunId"`
	Epoch            float64          `json:"epoch"`
	SampleIdentities []SampleIdentity `json:"sampleIdentities"`
	Digest           string           `json:"digest"`
}

// NewCreateSampleVisualizationsParams instantiates a new CreateSampleVisualizationsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSampleVisualizationsParams(projectId string, sessionRunId string, epoch float64, sampleIdentities []SampleIdentity, digest string) *CreateSampleVisualizationsParams {
	this := CreateSampleVisualizationsParams{}
	this.ProjectId = projectId
	this.SessionRunId = sessionRunId
	this.Epoch = epoch
	this.SampleIdentities = sampleIdentities
	this.Digest = digest
	return &this
}

// NewCreateSampleVisualizationsParamsWithDefaults instantiates a new CreateSampleVisualizationsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSampleVisualizationsParamsWithDefaults() *CreateSampleVisualizationsParams {
	this := CreateSampleVisualizationsParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *CreateSampleVisualizationsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CreateSampleVisualizationsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *CreateSampleVisualizationsParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSessionRunId returns the SessionRunId field value
func (o *CreateSampleVisualizationsParams) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *CreateSampleVisualizationsParams) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *CreateSampleVisualizationsParams) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetEpoch returns the Epoch field value
func (o *CreateSampleVisualizationsParams) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *CreateSampleVisualizationsParams) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *CreateSampleVisualizationsParams) SetEpoch(v float64) {
	o.Epoch = v
}

// GetSampleIdentities returns the SampleIdentities field value
func (o *CreateSampleVisualizationsParams) GetSampleIdentities() []SampleIdentity {
	if o == nil {
		var ret []SampleIdentity
		return ret
	}

	return o.SampleIdentities
}

// GetSampleIdentitiesOk returns a tuple with the SampleIdentities field value
// and a boolean to check if the value has been set.
func (o *CreateSampleVisualizationsParams) GetSampleIdentitiesOk() ([]SampleIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return o.SampleIdentities, true
}

// SetSampleIdentities sets field value
func (o *CreateSampleVisualizationsParams) SetSampleIdentities(v []SampleIdentity) {
	o.SampleIdentities = v
}

// GetDigest returns the Digest field value
func (o *CreateSampleVisualizationsParams) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *CreateSampleVisualizationsParams) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *CreateSampleVisualizationsParams) SetDigest(v string) {
	o.Digest = v
}

func (o CreateSampleVisualizationsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSampleVisualizationsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["epoch"] = o.Epoch
	toSerialize["sampleIdentities"] = o.SampleIdentities
	toSerialize["digest"] = o.Digest
	return toSerialize, nil
}

type NullableCreateSampleVisualizationsParams struct {
	value *CreateSampleVisualizationsParams
	isSet bool
}

func (v NullableCreateSampleVisualizationsParams) Get() *CreateSampleVisualizationsParams {
	return v.value
}

func (v *NullableCreateSampleVisualizationsParams) Set(val *CreateSampleVisualizationsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSampleVisualizationsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSampleVisualizationsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSampleVisualizationsParams(val *CreateSampleVisualizationsParams) *NullableCreateSampleVisualizationsParams {
	return &NullableCreateSampleVisualizationsParams{value: val, isSet: true}
}

func (v NullableCreateSampleVisualizationsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSampleVisualizationsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
