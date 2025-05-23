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

// checks if the SampleVisualizationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SampleVisualizationInfo{}

// SampleVisualizationInfo struct for SampleVisualizationInfo
type SampleVisualizationInfo struct {
	JobType        JobTypeEnum     `json:"job_type"`
	AnalyzeType    AnalyzeTypeEnum `json:"analyze_type"`
	Guid           string          `json:"guid"`
	SelectedSample SampleIdentity  `json:"selected_sample"`
}

// NewSampleVisualizationInfo instantiates a new SampleVisualizationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSampleVisualizationInfo(jobType JobTypeEnum, analyzeType AnalyzeTypeEnum, guid string, selectedSample SampleIdentity) *SampleVisualizationInfo {
	this := SampleVisualizationInfo{}
	this.JobType = jobType
	this.AnalyzeType = analyzeType
	this.Guid = guid
	this.SelectedSample = selectedSample
	return &this
}

// NewSampleVisualizationInfoWithDefaults instantiates a new SampleVisualizationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSampleVisualizationInfoWithDefaults() *SampleVisualizationInfo {
	this := SampleVisualizationInfo{}
	return &this
}

// GetJobType returns the JobType field value
func (o *SampleVisualizationInfo) GetJobType() JobTypeEnum {
	if o == nil {
		var ret JobTypeEnum
		return ret
	}

	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *SampleVisualizationInfo) GetJobTypeOk() (*JobTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value
func (o *SampleVisualizationInfo) SetJobType(v JobTypeEnum) {
	o.JobType = v
}

// GetAnalyzeType returns the AnalyzeType field value
func (o *SampleVisualizationInfo) GetAnalyzeType() AnalyzeTypeEnum {
	if o == nil {
		var ret AnalyzeTypeEnum
		return ret
	}

	return o.AnalyzeType
}

// GetAnalyzeTypeOk returns a tuple with the AnalyzeType field value
// and a boolean to check if the value has been set.
func (o *SampleVisualizationInfo) GetAnalyzeTypeOk() (*AnalyzeTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalyzeType, true
}

// SetAnalyzeType sets field value
func (o *SampleVisualizationInfo) SetAnalyzeType(v AnalyzeTypeEnum) {
	o.AnalyzeType = v
}

// GetGuid returns the Guid field value
func (o *SampleVisualizationInfo) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *SampleVisualizationInfo) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *SampleVisualizationInfo) SetGuid(v string) {
	o.Guid = v
}

// GetSelectedSample returns the SelectedSample field value
func (o *SampleVisualizationInfo) GetSelectedSample() SampleIdentity {
	if o == nil {
		var ret SampleIdentity
		return ret
	}

	return o.SelectedSample
}

// GetSelectedSampleOk returns a tuple with the SelectedSample field value
// and a boolean to check if the value has been set.
func (o *SampleVisualizationInfo) GetSelectedSampleOk() (*SampleIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedSample, true
}

// SetSelectedSample sets field value
func (o *SampleVisualizationInfo) SetSelectedSample(v SampleIdentity) {
	o.SelectedSample = v
}

func (o SampleVisualizationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SampleVisualizationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["job_type"] = o.JobType
	toSerialize["analyze_type"] = o.AnalyzeType
	toSerialize["guid"] = o.Guid
	toSerialize["selected_sample"] = o.SelectedSample
	return toSerialize, nil
}

type NullableSampleVisualizationInfo struct {
	value *SampleVisualizationInfo
	isSet bool
}

func (v NullableSampleVisualizationInfo) Get() *SampleVisualizationInfo {
	return v.value
}

func (v *NullableSampleVisualizationInfo) Set(val *SampleVisualizationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSampleVisualizationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSampleVisualizationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSampleVisualizationInfo(val *SampleVisualizationInfo) *NullableSampleVisualizationInfo {
	return &NullableSampleVisualizationInfo{value: val, isSet: true}
}

func (v NullableSampleVisualizationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSampleVisualizationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
