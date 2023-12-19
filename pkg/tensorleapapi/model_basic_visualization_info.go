/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.422
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the BasicVisualizationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasicVisualizationInfo{}

// BasicVisualizationInfo struct for BasicVisualizationInfo
type BasicVisualizationInfo struct {
	JobType     JobTypeEnum     `json:"job_type"`
	AnalyzeType AnalyzeTypeEnum `json:"analyze_type"`
	Guid        string          `json:"guid"`
}

// NewBasicVisualizationInfo instantiates a new BasicVisualizationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicVisualizationInfo(jobType JobTypeEnum, analyzeType AnalyzeTypeEnum, guid string) *BasicVisualizationInfo {
	this := BasicVisualizationInfo{}
	this.JobType = jobType
	this.AnalyzeType = analyzeType
	this.Guid = guid
	return &this
}

// NewBasicVisualizationInfoWithDefaults instantiates a new BasicVisualizationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicVisualizationInfoWithDefaults() *BasicVisualizationInfo {
	this := BasicVisualizationInfo{}
	return &this
}

// GetJobType returns the JobType field value
func (o *BasicVisualizationInfo) GetJobType() JobTypeEnum {
	if o == nil {
		var ret JobTypeEnum
		return ret
	}

	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *BasicVisualizationInfo) GetJobTypeOk() (*JobTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value
func (o *BasicVisualizationInfo) SetJobType(v JobTypeEnum) {
	o.JobType = v
}

// GetAnalyzeType returns the AnalyzeType field value
func (o *BasicVisualizationInfo) GetAnalyzeType() AnalyzeTypeEnum {
	if o == nil {
		var ret AnalyzeTypeEnum
		return ret
	}

	return o.AnalyzeType
}

// GetAnalyzeTypeOk returns a tuple with the AnalyzeType field value
// and a boolean to check if the value has been set.
func (o *BasicVisualizationInfo) GetAnalyzeTypeOk() (*AnalyzeTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalyzeType, true
}

// SetAnalyzeType sets field value
func (o *BasicVisualizationInfo) SetAnalyzeType(v AnalyzeTypeEnum) {
	o.AnalyzeType = v
}

// GetGuid returns the Guid field value
func (o *BasicVisualizationInfo) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *BasicVisualizationInfo) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *BasicVisualizationInfo) SetGuid(v string) {
	o.Guid = v
}

func (o BasicVisualizationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasicVisualizationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["job_type"] = o.JobType
	toSerialize["analyze_type"] = o.AnalyzeType
	toSerialize["guid"] = o.Guid
	return toSerialize, nil
}

type NullableBasicVisualizationInfo struct {
	value *BasicVisualizationInfo
	isSet bool
}

func (v NullableBasicVisualizationInfo) Get() *BasicVisualizationInfo {
	return v.value
}

func (v *NullableBasicVisualizationInfo) Set(val *BasicVisualizationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicVisualizationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicVisualizationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicVisualizationInfo(val *BasicVisualizationInfo) *NullableBasicVisualizationInfo {
	return &NullableBasicVisualizationInfo{value: val, isSet: true}
}

func (v NullableBasicVisualizationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicVisualizationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
