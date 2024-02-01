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

// checks if the PopExploreInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PopExploreInfo{}

// PopExploreInfo struct for PopExploreInfo
type PopExploreInfo struct {
	JobType                       JobTypeEnum     `json:"job_type"`
	AnalyzeType                   AnalyzeTypeEnum `json:"analyze_type"`
	Guid                          string          `json:"guid"`
	PopulationExplorationNSamples float64         `json:"population_exploration_n_samples"`
}

// NewPopExploreInfo instantiates a new PopExploreInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopExploreInfo(jobType JobTypeEnum, analyzeType AnalyzeTypeEnum, guid string, populationExplorationNSamples float64) *PopExploreInfo {
	this := PopExploreInfo{}
	this.JobType = jobType
	this.AnalyzeType = analyzeType
	this.Guid = guid
	this.PopulationExplorationNSamples = populationExplorationNSamples
	return &this
}

// NewPopExploreInfoWithDefaults instantiates a new PopExploreInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopExploreInfoWithDefaults() *PopExploreInfo {
	this := PopExploreInfo{}
	return &this
}

// GetJobType returns the JobType field value
func (o *PopExploreInfo) GetJobType() JobTypeEnum {
	if o == nil {
		var ret JobTypeEnum
		return ret
	}

	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *PopExploreInfo) GetJobTypeOk() (*JobTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value
func (o *PopExploreInfo) SetJobType(v JobTypeEnum) {
	o.JobType = v
}

// GetAnalyzeType returns the AnalyzeType field value
func (o *PopExploreInfo) GetAnalyzeType() AnalyzeTypeEnum {
	if o == nil {
		var ret AnalyzeTypeEnum
		return ret
	}

	return o.AnalyzeType
}

// GetAnalyzeTypeOk returns a tuple with the AnalyzeType field value
// and a boolean to check if the value has been set.
func (o *PopExploreInfo) GetAnalyzeTypeOk() (*AnalyzeTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalyzeType, true
}

// SetAnalyzeType sets field value
func (o *PopExploreInfo) SetAnalyzeType(v AnalyzeTypeEnum) {
	o.AnalyzeType = v
}

// GetGuid returns the Guid field value
func (o *PopExploreInfo) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *PopExploreInfo) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *PopExploreInfo) SetGuid(v string) {
	o.Guid = v
}

// GetPopulationExplorationNSamples returns the PopulationExplorationNSamples field value
func (o *PopExploreInfo) GetPopulationExplorationNSamples() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.PopulationExplorationNSamples
}

// GetPopulationExplorationNSamplesOk returns a tuple with the PopulationExplorationNSamples field value
// and a boolean to check if the value has been set.
func (o *PopExploreInfo) GetPopulationExplorationNSamplesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PopulationExplorationNSamples, true
}

// SetPopulationExplorationNSamples sets field value
func (o *PopExploreInfo) SetPopulationExplorationNSamples(v float64) {
	o.PopulationExplorationNSamples = v
}

func (o PopExploreInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PopExploreInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["job_type"] = o.JobType
	toSerialize["analyze_type"] = o.AnalyzeType
	toSerialize["guid"] = o.Guid
	toSerialize["population_exploration_n_samples"] = o.PopulationExplorationNSamples
	return toSerialize, nil
}

type NullablePopExploreInfo struct {
	value *PopExploreInfo
	isSet bool
}

func (v NullablePopExploreInfo) Get() *PopExploreInfo {
	return v.value
}

func (v *NullablePopExploreInfo) Set(val *PopExploreInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePopExploreInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePopExploreInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopExploreInfo(val *PopExploreInfo) *NullablePopExploreInfo {
	return &NullablePopExploreInfo{value: val, isSet: true}
}

func (v NullablePopExploreInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopExploreInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
