/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.423
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SampleSelectionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SampleSelectionInfo{}

// SampleSelectionInfo struct for SampleSelectionInfo
type SampleSelectionInfo struct {
	JobType          JobTypeEnum     `json:"job_type"`
	AnalyzeType      AnalyzeTypeEnum `json:"analyze_type"`
	Guid             string          `json:"guid"`
	NSelectedSamples float64         `json:"n_selected_samples"`
	SelectedSample   SampleIdentity  `json:"selected_sample"`
}

// NewSampleSelectionInfo instantiates a new SampleSelectionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSampleSelectionInfo(jobType JobTypeEnum, analyzeType AnalyzeTypeEnum, guid string, nSelectedSamples float64, selectedSample SampleIdentity) *SampleSelectionInfo {
	this := SampleSelectionInfo{}
	this.JobType = jobType
	this.AnalyzeType = analyzeType
	this.Guid = guid
	this.NSelectedSamples = nSelectedSamples
	this.SelectedSample = selectedSample
	return &this
}

// NewSampleSelectionInfoWithDefaults instantiates a new SampleSelectionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSampleSelectionInfoWithDefaults() *SampleSelectionInfo {
	this := SampleSelectionInfo{}
	return &this
}

// GetJobType returns the JobType field value
func (o *SampleSelectionInfo) GetJobType() JobTypeEnum {
	if o == nil {
		var ret JobTypeEnum
		return ret
	}

	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *SampleSelectionInfo) GetJobTypeOk() (*JobTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value
func (o *SampleSelectionInfo) SetJobType(v JobTypeEnum) {
	o.JobType = v
}

// GetAnalyzeType returns the AnalyzeType field value
func (o *SampleSelectionInfo) GetAnalyzeType() AnalyzeTypeEnum {
	if o == nil {
		var ret AnalyzeTypeEnum
		return ret
	}

	return o.AnalyzeType
}

// GetAnalyzeTypeOk returns a tuple with the AnalyzeType field value
// and a boolean to check if the value has been set.
func (o *SampleSelectionInfo) GetAnalyzeTypeOk() (*AnalyzeTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalyzeType, true
}

// SetAnalyzeType sets field value
func (o *SampleSelectionInfo) SetAnalyzeType(v AnalyzeTypeEnum) {
	o.AnalyzeType = v
}

// GetGuid returns the Guid field value
func (o *SampleSelectionInfo) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *SampleSelectionInfo) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *SampleSelectionInfo) SetGuid(v string) {
	o.Guid = v
}

// GetNSelectedSamples returns the NSelectedSamples field value
func (o *SampleSelectionInfo) GetNSelectedSamples() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.NSelectedSamples
}

// GetNSelectedSamplesOk returns a tuple with the NSelectedSamples field value
// and a boolean to check if the value has been set.
func (o *SampleSelectionInfo) GetNSelectedSamplesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NSelectedSamples, true
}

// SetNSelectedSamples sets field value
func (o *SampleSelectionInfo) SetNSelectedSamples(v float64) {
	o.NSelectedSamples = v
}

// GetSelectedSample returns the SelectedSample field value
func (o *SampleSelectionInfo) GetSelectedSample() SampleIdentity {
	if o == nil {
		var ret SampleIdentity
		return ret
	}

	return o.SelectedSample
}

// GetSelectedSampleOk returns a tuple with the SelectedSample field value
// and a boolean to check if the value has been set.
func (o *SampleSelectionInfo) GetSelectedSampleOk() (*SampleIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedSample, true
}

// SetSelectedSample sets field value
func (o *SampleSelectionInfo) SetSelectedSample(v SampleIdentity) {
	o.SelectedSample = v
}

func (o SampleSelectionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SampleSelectionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["job_type"] = o.JobType
	toSerialize["analyze_type"] = o.AnalyzeType
	toSerialize["guid"] = o.Guid
	toSerialize["n_selected_samples"] = o.NSelectedSamples
	toSerialize["selected_sample"] = o.SelectedSample
	return toSerialize, nil
}

type NullableSampleSelectionInfo struct {
	value *SampleSelectionInfo
	isSet bool
}

func (v NullableSampleSelectionInfo) Get() *SampleSelectionInfo {
	return v.value
}

func (v *NullableSampleSelectionInfo) Set(val *SampleSelectionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSampleSelectionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSampleSelectionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSampleSelectionInfo(val *SampleSelectionInfo) *NullableSampleSelectionInfo {
	return &NullableSampleSelectionInfo{value: val, isSet: true}
}

func (v NullableSampleSelectionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSampleSelectionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
