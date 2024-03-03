/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the JobNotificationSampleContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobNotificationSampleContext{}

// JobNotificationSampleContext struct for JobNotificationSampleContext
type JobNotificationSampleContext struct {
	JobId       string         `json:"jobId"`
	JobType     JobType        `json:"jobType"`
	ProjectName string         `json:"projectName"`
	ProjectId   string         `json:"projectId"`
	ModelName   string         `json:"modelName"`
	ModelExtId  string         `json:"modelExtId"`
	Sample      SampleIdentity `json:"sample"`
}

// NewJobNotificationSampleContext instantiates a new JobNotificationSampleContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobNotificationSampleContext(jobId string, jobType JobType, projectName string, projectId string, modelName string, modelExtId string, sample SampleIdentity) *JobNotificationSampleContext {
	this := JobNotificationSampleContext{}
	this.JobId = jobId
	this.JobType = jobType
	this.ProjectName = projectName
	this.ProjectId = projectId
	this.ModelName = modelName
	this.ModelExtId = modelExtId
	this.Sample = sample
	return &this
}

// NewJobNotificationSampleContextWithDefaults instantiates a new JobNotificationSampleContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobNotificationSampleContextWithDefaults() *JobNotificationSampleContext {
	this := JobNotificationSampleContext{}
	return &this
}

// GetJobId returns the JobId field value
func (o *JobNotificationSampleContext) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *JobNotificationSampleContext) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *JobNotificationSampleContext) SetJobId(v string) {
	o.JobId = v
}

// GetJobType returns the JobType field value
func (o *JobNotificationSampleContext) GetJobType() JobType {
	if o == nil {
		var ret JobType
		return ret
	}

	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *JobNotificationSampleContext) GetJobTypeOk() (*JobType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value
func (o *JobNotificationSampleContext) SetJobType(v JobType) {
	o.JobType = v
}

// GetProjectName returns the ProjectName field value
func (o *JobNotificationSampleContext) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *JobNotificationSampleContext) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *JobNotificationSampleContext) SetProjectName(v string) {
	o.ProjectName = v
}

// GetProjectId returns the ProjectId field value
func (o *JobNotificationSampleContext) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *JobNotificationSampleContext) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *JobNotificationSampleContext) SetProjectId(v string) {
	o.ProjectId = v
}

// GetModelName returns the ModelName field value
func (o *JobNotificationSampleContext) GetModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value
// and a boolean to check if the value has been set.
func (o *JobNotificationSampleContext) GetModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelName, true
}

// SetModelName sets field value
func (o *JobNotificationSampleContext) SetModelName(v string) {
	o.ModelName = v
}

// GetModelExtId returns the ModelExtId field value
func (o *JobNotificationSampleContext) GetModelExtId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelExtId
}

// GetModelExtIdOk returns a tuple with the ModelExtId field value
// and a boolean to check if the value has been set.
func (o *JobNotificationSampleContext) GetModelExtIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelExtId, true
}

// SetModelExtId sets field value
func (o *JobNotificationSampleContext) SetModelExtId(v string) {
	o.ModelExtId = v
}

// GetSample returns the Sample field value
func (o *JobNotificationSampleContext) GetSample() SampleIdentity {
	if o == nil {
		var ret SampleIdentity
		return ret
	}

	return o.Sample
}

// GetSampleOk returns a tuple with the Sample field value
// and a boolean to check if the value has been set.
func (o *JobNotificationSampleContext) GetSampleOk() (*SampleIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sample, true
}

// SetSample sets field value
func (o *JobNotificationSampleContext) SetSample(v SampleIdentity) {
	o.Sample = v
}

func (o JobNotificationSampleContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobNotificationSampleContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobId"] = o.JobId
	toSerialize["jobType"] = o.JobType
	toSerialize["projectName"] = o.ProjectName
	toSerialize["projectId"] = o.ProjectId
	toSerialize["modelName"] = o.ModelName
	toSerialize["modelExtId"] = o.ModelExtId
	toSerialize["sample"] = o.Sample
	return toSerialize, nil
}

type NullableJobNotificationSampleContext struct {
	value *JobNotificationSampleContext
	isSet bool
}

func (v NullableJobNotificationSampleContext) Get() *JobNotificationSampleContext {
	return v.value
}

func (v *NullableJobNotificationSampleContext) Set(val *JobNotificationSampleContext) {
	v.value = val
	v.isSet = true
}

func (v NullableJobNotificationSampleContext) IsSet() bool {
	return v.isSet
}

func (v *NullableJobNotificationSampleContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobNotificationSampleContext(val *JobNotificationSampleContext) *NullableJobNotificationSampleContext {
	return &NullableJobNotificationSampleContext{value: val, isSet: true}
}

func (v NullableJobNotificationSampleContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobNotificationSampleContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
