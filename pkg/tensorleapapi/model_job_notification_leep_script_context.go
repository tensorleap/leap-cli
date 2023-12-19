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

// checks if the JobNotificationLeepScriptContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobNotificationLeepScriptContext{}

// JobNotificationLeepScriptContext struct for JobNotificationLeepScriptContext
type JobNotificationLeepScriptContext struct {
	JobId               string      `json:"jobId"`
	JobType             JobTypeEnum `json:"jobType"`
	LeepScriptName      string      `json:"leepScriptName"`
	LeepScriptId        string      `json:"leepScriptId"`
	LeepScriptVersion   *string     `json:"leepScriptVersion,omitempty"`
	LeepScriptVersionId *string     `json:"leepScriptVersionId,omitempty"`
}

// NewJobNotificationLeepScriptContext instantiates a new JobNotificationLeepScriptContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobNotificationLeepScriptContext(jobId string, jobType JobTypeEnum, leepScriptName string, leepScriptId string) *JobNotificationLeepScriptContext {
	this := JobNotificationLeepScriptContext{}
	this.JobId = jobId
	this.JobType = jobType
	this.LeepScriptName = leepScriptName
	this.LeepScriptId = leepScriptId
	return &this
}

// NewJobNotificationLeepScriptContextWithDefaults instantiates a new JobNotificationLeepScriptContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobNotificationLeepScriptContextWithDefaults() *JobNotificationLeepScriptContext {
	this := JobNotificationLeepScriptContext{}
	return &this
}

// GetJobId returns the JobId field value
func (o *JobNotificationLeepScriptContext) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *JobNotificationLeepScriptContext) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *JobNotificationLeepScriptContext) SetJobId(v string) {
	o.JobId = v
}

// GetJobType returns the JobType field value
func (o *JobNotificationLeepScriptContext) GetJobType() JobTypeEnum {
	if o == nil {
		var ret JobTypeEnum
		return ret
	}

	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *JobNotificationLeepScriptContext) GetJobTypeOk() (*JobTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value
func (o *JobNotificationLeepScriptContext) SetJobType(v JobTypeEnum) {
	o.JobType = v
}

// GetLeepScriptName returns the LeepScriptName field value
func (o *JobNotificationLeepScriptContext) GetLeepScriptName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LeepScriptName
}

// GetLeepScriptNameOk returns a tuple with the LeepScriptName field value
// and a boolean to check if the value has been set.
func (o *JobNotificationLeepScriptContext) GetLeepScriptNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeepScriptName, true
}

// SetLeepScriptName sets field value
func (o *JobNotificationLeepScriptContext) SetLeepScriptName(v string) {
	o.LeepScriptName = v
}

// GetLeepScriptId returns the LeepScriptId field value
func (o *JobNotificationLeepScriptContext) GetLeepScriptId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LeepScriptId
}

// GetLeepScriptIdOk returns a tuple with the LeepScriptId field value
// and a boolean to check if the value has been set.
func (o *JobNotificationLeepScriptContext) GetLeepScriptIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeepScriptId, true
}

// SetLeepScriptId sets field value
func (o *JobNotificationLeepScriptContext) SetLeepScriptId(v string) {
	o.LeepScriptId = v
}

// GetLeepScriptVersion returns the LeepScriptVersion field value if set, zero value otherwise.
func (o *JobNotificationLeepScriptContext) GetLeepScriptVersion() string {
	if o == nil || IsNil(o.LeepScriptVersion) {
		var ret string
		return ret
	}
	return *o.LeepScriptVersion
}

// GetLeepScriptVersionOk returns a tuple with the LeepScriptVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobNotificationLeepScriptContext) GetLeepScriptVersionOk() (*string, bool) {
	if o == nil || IsNil(o.LeepScriptVersion) {
		return nil, false
	}
	return o.LeepScriptVersion, true
}

// HasLeepScriptVersion returns a boolean if a field has been set.
func (o *JobNotificationLeepScriptContext) HasLeepScriptVersion() bool {
	if o != nil && !IsNil(o.LeepScriptVersion) {
		return true
	}

	return false
}

// SetLeepScriptVersion gets a reference to the given string and assigns it to the LeepScriptVersion field.
func (o *JobNotificationLeepScriptContext) SetLeepScriptVersion(v string) {
	o.LeepScriptVersion = &v
}

// GetLeepScriptVersionId returns the LeepScriptVersionId field value if set, zero value otherwise.
func (o *JobNotificationLeepScriptContext) GetLeepScriptVersionId() string {
	if o == nil || IsNil(o.LeepScriptVersionId) {
		var ret string
		return ret
	}
	return *o.LeepScriptVersionId
}

// GetLeepScriptVersionIdOk returns a tuple with the LeepScriptVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobNotificationLeepScriptContext) GetLeepScriptVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.LeepScriptVersionId) {
		return nil, false
	}
	return o.LeepScriptVersionId, true
}

// HasLeepScriptVersionId returns a boolean if a field has been set.
func (o *JobNotificationLeepScriptContext) HasLeepScriptVersionId() bool {
	if o != nil && !IsNil(o.LeepScriptVersionId) {
		return true
	}

	return false
}

// SetLeepScriptVersionId gets a reference to the given string and assigns it to the LeepScriptVersionId field.
func (o *JobNotificationLeepScriptContext) SetLeepScriptVersionId(v string) {
	o.LeepScriptVersionId = &v
}

func (o JobNotificationLeepScriptContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobNotificationLeepScriptContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobId"] = o.JobId
	toSerialize["jobType"] = o.JobType
	toSerialize["leepScriptName"] = o.LeepScriptName
	toSerialize["leepScriptId"] = o.LeepScriptId
	if !IsNil(o.LeepScriptVersion) {
		toSerialize["leepScriptVersion"] = o.LeepScriptVersion
	}
	if !IsNil(o.LeepScriptVersionId) {
		toSerialize["leepScriptVersionId"] = o.LeepScriptVersionId
	}
	return toSerialize, nil
}

type NullableJobNotificationLeepScriptContext struct {
	value *JobNotificationLeepScriptContext
	isSet bool
}

func (v NullableJobNotificationLeepScriptContext) Get() *JobNotificationLeepScriptContext {
	return v.value
}

func (v *NullableJobNotificationLeepScriptContext) Set(val *JobNotificationLeepScriptContext) {
	v.value = val
	v.isSet = true
}

func (v NullableJobNotificationLeepScriptContext) IsSet() bool {
	return v.isSet
}

func (v *NullableJobNotificationLeepScriptContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobNotificationLeepScriptContext(val *JobNotificationLeepScriptContext) *NullableJobNotificationLeepScriptContext {
	return &NullableJobNotificationLeepScriptContext{value: val, isSet: true}
}

func (v NullableJobNotificationLeepScriptContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobNotificationLeepScriptContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
