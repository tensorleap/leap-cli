/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.423
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the Job type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Job{}

// Job struct for Job
type Job struct {
	ProjectId    string      `json:"projectId"`
	Cid          string      `json:"cid"`
	CreatedBy    string      `json:"createdBy"`
	Version      *string     `json:"version,omitempty"`
	Type         JobTypeEnum `json:"type"`
	SubType      *JobSubType `json:"subType,omitempty"`
	Status       JobStatus   `json:"status"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
	Params       *JobParams  `json:"params,omitempty"`
	SessionRunId *string     `json:"sessionRunId,omitempty"`
	TeamId       string      `json:"teamId"`
}

// NewJob instantiates a new Job object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJob(projectId string, cid string, createdBy string, type_ JobTypeEnum, status JobStatus, createdAt time.Time, updatedAt time.Time, teamId string) *Job {
	this := Job{}
	this.ProjectId = projectId
	this.Cid = cid
	this.CreatedBy = createdBy
	this.Type = type_
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.TeamId = teamId
	return &this
}

// NewJobWithDefaults instantiates a new Job object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobWithDefaults() *Job {
	this := Job{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *Job) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Job) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Job) SetProjectId(v string) {
	o.ProjectId = v
}

// GetCid returns the Cid field value
func (o *Job) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *Job) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *Job) SetCid(v string) {
	o.Cid = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Job) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Job) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Job) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Job) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Job) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Job) SetVersion(v string) {
	o.Version = &v
}

// GetType returns the Type field value
func (o *Job) GetType() JobTypeEnum {
	if o == nil {
		var ret JobTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Job) GetTypeOk() (*JobTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Job) SetType(v JobTypeEnum) {
	o.Type = v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *Job) GetSubType() JobSubType {
	if o == nil || IsNil(o.SubType) {
		var ret JobSubType
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetSubTypeOk() (*JobSubType, bool) {
	if o == nil || IsNil(o.SubType) {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *Job) HasSubType() bool {
	if o != nil && !IsNil(o.SubType) {
		return true
	}

	return false
}

// SetSubType gets a reference to the given JobSubType and assigns it to the SubType field.
func (o *Job) SetSubType(v JobSubType) {
	o.SubType = &v
}

// GetStatus returns the Status field value
func (o *Job) GetStatus() JobStatus {
	if o == nil {
		var ret JobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Job) GetStatusOk() (*JobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Job) SetStatus(v JobStatus) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Job) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Job) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Job) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Job) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Job) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Job) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *Job) GetParams() JobParams {
	if o == nil || IsNil(o.Params) {
		var ret JobParams
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetParamsOk() (*JobParams, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *Job) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given JobParams and assigns it to the Params field.
func (o *Job) SetParams(v JobParams) {
	o.Params = &v
}

// GetSessionRunId returns the SessionRunId field value if set, zero value otherwise.
func (o *Job) GetSessionRunId() string {
	if o == nil || IsNil(o.SessionRunId) {
		var ret string
		return ret
	}
	return *o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetSessionRunIdOk() (*string, bool) {
	if o == nil || IsNil(o.SessionRunId) {
		return nil, false
	}
	return o.SessionRunId, true
}

// HasSessionRunId returns a boolean if a field has been set.
func (o *Job) HasSessionRunId() bool {
	if o != nil && !IsNil(o.SessionRunId) {
		return true
	}

	return false
}

// SetSessionRunId gets a reference to the given string and assigns it to the SessionRunId field.
func (o *Job) SetSessionRunId(v string) {
	o.SessionRunId = &v
}

// GetTeamId returns the TeamId field value
func (o *Job) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *Job) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *Job) SetTeamId(v string) {
	o.TeamId = v
}

func (o Job) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Job) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["cid"] = o.Cid
	toSerialize["createdBy"] = o.CreatedBy
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.SubType) {
		toSerialize["subType"] = o.SubType
	}
	toSerialize["status"] = o.Status
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.SessionRunId) {
		toSerialize["sessionRunId"] = o.SessionRunId
	}
	toSerialize["teamId"] = o.TeamId
	return toSerialize, nil
}

type NullableJob struct {
	value *Job
	isSet bool
}

func (v NullableJob) Get() *Job {
	return v.value
}

func (v *NullableJob) Set(val *Job) {
	v.value = val
	v.isSet = true
}

func (v NullableJob) IsSet() bool {
	return v.isSet
}

func (v *NullableJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJob(val *Job) *NullableJob {
	return &NullableJob{value: val, isSet: true}
}

func (v NullableJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
