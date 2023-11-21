/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.404
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the JobNotificationAnalyzeContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobNotificationAnalyzeContext{}

// JobNotificationAnalyzeContext struct for JobNotificationAnalyzeContext
type JobNotificationAnalyzeContext struct {
	JobId        string      `json:"jobId"`
	JobType      JobTypeEnum `json:"jobType"`
	ProjectName  string      `json:"projectName"`
	ProjectId    string      `json:"projectId"`
	ModelName    string      `json:"modelName"`
	ModelExtId   string      `json:"modelExtId"`
	SessionRunId string      `json:"sessionRunId"`
}

// NewJobNotificationAnalyzeContext instantiates a new JobNotificationAnalyzeContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobNotificationAnalyzeContext(jobId string, jobType JobTypeEnum, projectName string, projectId string, modelName string, modelExtId string, sessionRunId string) *JobNotificationAnalyzeContext {
	this := JobNotificationAnalyzeContext{}
	this.JobId = jobId
	this.JobType = jobType
	this.ProjectName = projectName
	this.ProjectId = projectId
	this.ModelName = modelName
	this.ModelExtId = modelExtId
	this.SessionRunId = sessionRunId
	return &this
}

// NewJobNotificationAnalyzeContextWithDefaults instantiates a new JobNotificationAnalyzeContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobNotificationAnalyzeContextWithDefaults() *JobNotificationAnalyzeContext {
	this := JobNotificationAnalyzeContext{}
	return &this
}

// GetJobId returns the JobId field value
func (o *JobNotificationAnalyzeContext) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *JobNotificationAnalyzeContext) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *JobNotificationAnalyzeContext) SetJobId(v string) {
	o.JobId = v
}

// GetJobType returns the JobType field value
func (o *JobNotificationAnalyzeContext) GetJobType() JobTypeEnum {
	if o == nil {
		var ret JobTypeEnum
		return ret
	}

	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *JobNotificationAnalyzeContext) GetJobTypeOk() (*JobTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value
func (o *JobNotificationAnalyzeContext) SetJobType(v JobTypeEnum) {
	o.JobType = v
}

// GetProjectName returns the ProjectName field value
func (o *JobNotificationAnalyzeContext) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *JobNotificationAnalyzeContext) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *JobNotificationAnalyzeContext) SetProjectName(v string) {
	o.ProjectName = v
}

// GetProjectId returns the ProjectId field value
func (o *JobNotificationAnalyzeContext) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *JobNotificationAnalyzeContext) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *JobNotificationAnalyzeContext) SetProjectId(v string) {
	o.ProjectId = v
}

// GetModelName returns the ModelName field value
func (o *JobNotificationAnalyzeContext) GetModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value
// and a boolean to check if the value has been set.
func (o *JobNotificationAnalyzeContext) GetModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelName, true
}

// SetModelName sets field value
func (o *JobNotificationAnalyzeContext) SetModelName(v string) {
	o.ModelName = v
}

// GetModelExtId returns the ModelExtId field value
func (o *JobNotificationAnalyzeContext) GetModelExtId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelExtId
}

// GetModelExtIdOk returns a tuple with the ModelExtId field value
// and a boolean to check if the value has been set.
func (o *JobNotificationAnalyzeContext) GetModelExtIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelExtId, true
}

// SetModelExtId sets field value
func (o *JobNotificationAnalyzeContext) SetModelExtId(v string) {
	o.ModelExtId = v
}

// GetSessionRunId returns the SessionRunId field value
func (o *JobNotificationAnalyzeContext) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *JobNotificationAnalyzeContext) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *JobNotificationAnalyzeContext) SetSessionRunId(v string) {
	o.SessionRunId = v
}

func (o JobNotificationAnalyzeContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobNotificationAnalyzeContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobId"] = o.JobId
	toSerialize["jobType"] = o.JobType
	toSerialize["projectName"] = o.ProjectName
	toSerialize["projectId"] = o.ProjectId
	toSerialize["modelName"] = o.ModelName
	toSerialize["modelExtId"] = o.ModelExtId
	toSerialize["sessionRunId"] = o.SessionRunId
	return toSerialize, nil
}

type NullableJobNotificationAnalyzeContext struct {
	value *JobNotificationAnalyzeContext
	isSet bool
}

func (v NullableJobNotificationAnalyzeContext) Get() *JobNotificationAnalyzeContext {
	return v.value
}

func (v *NullableJobNotificationAnalyzeContext) Set(val *JobNotificationAnalyzeContext) {
	v.value = val
	v.isSet = true
}

func (v NullableJobNotificationAnalyzeContext) IsSet() bool {
	return v.isSet
}

func (v *NullableJobNotificationAnalyzeContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobNotificationAnalyzeContext(val *JobNotificationAnalyzeContext) *NullableJobNotificationAnalyzeContext {
	return &NullableJobNotificationAnalyzeContext{value: val, isSet: true}
}

func (v NullableJobNotificationAnalyzeContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobNotificationAnalyzeContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}