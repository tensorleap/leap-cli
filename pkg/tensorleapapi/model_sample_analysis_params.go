/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SampleAnalysisParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SampleAnalysisParams{}

// SampleAnalysisParams struct for SampleAnalysisParams
type SampleAnalysisParams struct {
	SessionRunId   string             `json:"sessionRunId"`
	ProjectId      string             `json:"projectId"`
	SampleIdentity SampleIdentity     `json:"sampleIdentity"`
	FromEpoch      float64            `json:"fromEpoch"`
	Algo           SampleAnalysisAlgo `json:"algo"`
}

// NewSampleAnalysisParams instantiates a new SampleAnalysisParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSampleAnalysisParams(sessionRunId string, projectId string, sampleIdentity SampleIdentity, fromEpoch float64, algo SampleAnalysisAlgo) *SampleAnalysisParams {
	this := SampleAnalysisParams{}
	this.SessionRunId = sessionRunId
	this.ProjectId = projectId
	this.SampleIdentity = sampleIdentity
	this.FromEpoch = fromEpoch
	this.Algo = algo
	return &this
}

// NewSampleAnalysisParamsWithDefaults instantiates a new SampleAnalysisParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSampleAnalysisParamsWithDefaults() *SampleAnalysisParams {
	this := SampleAnalysisParams{}
	return &this
}

// GetSessionRunId returns the SessionRunId field value
func (o *SampleAnalysisParams) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisParams) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *SampleAnalysisParams) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetProjectId returns the ProjectId field value
func (o *SampleAnalysisParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *SampleAnalysisParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSampleIdentity returns the SampleIdentity field value
func (o *SampleAnalysisParams) GetSampleIdentity() SampleIdentity {
	if o == nil {
		var ret SampleIdentity
		return ret
	}

	return o.SampleIdentity
}

// GetSampleIdentityOk returns a tuple with the SampleIdentity field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisParams) GetSampleIdentityOk() (*SampleIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SampleIdentity, true
}

// SetSampleIdentity sets field value
func (o *SampleAnalysisParams) SetSampleIdentity(v SampleIdentity) {
	o.SampleIdentity = v
}

// GetFromEpoch returns the FromEpoch field value
func (o *SampleAnalysisParams) GetFromEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.FromEpoch
}

// GetFromEpochOk returns a tuple with the FromEpoch field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisParams) GetFromEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromEpoch, true
}

// SetFromEpoch sets field value
func (o *SampleAnalysisParams) SetFromEpoch(v float64) {
	o.FromEpoch = v
}

// GetAlgo returns the Algo field value
func (o *SampleAnalysisParams) GetAlgo() SampleAnalysisAlgo {
	if o == nil {
		var ret SampleAnalysisAlgo
		return ret
	}

	return o.Algo
}

// GetAlgoOk returns a tuple with the Algo field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisParams) GetAlgoOk() (*SampleAnalysisAlgo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algo, true
}

// SetAlgo sets field value
func (o *SampleAnalysisParams) SetAlgo(v SampleAnalysisAlgo) {
	o.Algo = v
}

func (o SampleAnalysisParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SampleAnalysisParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["projectId"] = o.ProjectId
	toSerialize["sampleIdentity"] = o.SampleIdentity
	toSerialize["fromEpoch"] = o.FromEpoch
	toSerialize["algo"] = o.Algo
	return toSerialize, nil
}

type NullableSampleAnalysisParams struct {
	value *SampleAnalysisParams
	isSet bool
}

func (v NullableSampleAnalysisParams) Get() *SampleAnalysisParams {
	return v.value
}

func (v *NullableSampleAnalysisParams) Set(val *SampleAnalysisParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSampleAnalysisParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSampleAnalysisParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSampleAnalysisParams(val *SampleAnalysisParams) *NullableSampleAnalysisParams {
	return &NullableSampleAnalysisParams{value: val, isSet: true}
}

func (v NullableSampleAnalysisParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSampleAnalysisParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
