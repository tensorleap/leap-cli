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

// checks if the JobNotificationModelContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobNotificationModelContext{}

// JobNotificationModelContext struct for JobNotificationModelContext
type JobNotificationModelContext struct {
	JobId       string      `json:"jobId"`
	JobType     JobTypeEnum `json:"jobType"`
	ProjectName string      `json:"projectName"`
	ProjectId   string      `json:"projectId"`
	ModelName   string      `json:"modelName"`
	ModelExtId  string      `json:"modelExtId"`
}

// NewJobNotificationModelContext instantiates a new JobNotificationModelContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobNotificationModelContext(jobId string, jobType JobTypeEnum, projectName string, projectId string, modelName string, modelExtId string) *JobNotificationModelContext {
	this := JobNotificationModelContext{}
	this.JobId = jobId
	this.JobType = jobType
	this.ProjectName = projectName
	this.ProjectId = projectId
	this.ModelName = modelName
	this.ModelExtId = modelExtId
	return &this
}

// NewJobNotificationModelContextWithDefaults instantiates a new JobNotificationModelContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobNotificationModelContextWithDefaults() *JobNotificationModelContext {
	this := JobNotificationModelContext{}
	return &this
}

// GetJobId returns the JobId field value
func (o *JobNotificationModelContext) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *JobNotificationModelContext) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *JobNotificationModelContext) SetJobId(v string) {
	o.JobId = v
}

// GetJobType returns the JobType field value
func (o *JobNotificationModelContext) GetJobType() JobTypeEnum {
	if o == nil {
		var ret JobTypeEnum
		return ret
	}

	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *JobNotificationModelContext) GetJobTypeOk() (*JobTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value
func (o *JobNotificationModelContext) SetJobType(v JobTypeEnum) {
	o.JobType = v
}

// GetProjectName returns the ProjectName field value
func (o *JobNotificationModelContext) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *JobNotificationModelContext) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *JobNotificationModelContext) SetProjectName(v string) {
	o.ProjectName = v
}

// GetProjectId returns the ProjectId field value
func (o *JobNotificationModelContext) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *JobNotificationModelContext) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *JobNotificationModelContext) SetProjectId(v string) {
	o.ProjectId = v
}

// GetModelName returns the ModelName field value
func (o *JobNotificationModelContext) GetModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value
// and a boolean to check if the value has been set.
func (o *JobNotificationModelContext) GetModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelName, true
}

// SetModelName sets field value
func (o *JobNotificationModelContext) SetModelName(v string) {
	o.ModelName = v
}

// GetModelExtId returns the ModelExtId field value
func (o *JobNotificationModelContext) GetModelExtId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelExtId
}

// GetModelExtIdOk returns a tuple with the ModelExtId field value
// and a boolean to check if the value has been set.
func (o *JobNotificationModelContext) GetModelExtIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelExtId, true
}

// SetModelExtId sets field value
func (o *JobNotificationModelContext) SetModelExtId(v string) {
	o.ModelExtId = v
}

func (o JobNotificationModelContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobNotificationModelContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobId"] = o.JobId
	toSerialize["jobType"] = o.JobType
	toSerialize["projectName"] = o.ProjectName
	toSerialize["projectId"] = o.ProjectId
	toSerialize["modelName"] = o.ModelName
	toSerialize["modelExtId"] = o.ModelExtId
	return toSerialize, nil
}

type NullableJobNotificationModelContext struct {
	value *JobNotificationModelContext
	isSet bool
}

func (v NullableJobNotificationModelContext) Get() *JobNotificationModelContext {
	return v.value
}

func (v *NullableJobNotificationModelContext) Set(val *JobNotificationModelContext) {
	v.value = val
	v.isSet = true
}

func (v NullableJobNotificationModelContext) IsSet() bool {
	return v.isSet
}

func (v *NullableJobNotificationModelContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobNotificationModelContext(val *JobNotificationModelContext) *NullableJobNotificationModelContext {
	return &NullableJobNotificationModelContext{value: val, isSet: true}
}

func (v NullableJobNotificationModelContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobNotificationModelContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
