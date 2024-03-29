/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the RunProcess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunProcess{}

// RunProcess struct for RunProcess
type RunProcess struct {
	JobId          string     `json:"jobId"`
	VersionName    string     `json:"versionName"`
	VersionId      string     `json:"versionId"`
	ProjectName    string     `json:"projectName"`
	JobType        string     `json:"jobType"`
	Status         JobStatus  `json:"status"`
	CreatedAt      string     `json:"createdAt"`
	UpdatedAt      string     `json:"updatedAt"`
	SessionName    string     `json:"sessionName"`
	SessionRunName string     `json:"sessionRunName"`
	SessionRunId   *string    `json:"sessionRunId,omitempty"`
	Events         []JobEvent `json:"events"`
	Params         *JobParams `json:"params,omitempty"`
}

// NewRunProcess instantiates a new RunProcess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunProcess(jobId string, versionName string, versionId string, projectName string, jobType string, status JobStatus, createdAt string, updatedAt string, sessionName string, sessionRunName string, events []JobEvent) *RunProcess {
	this := RunProcess{}
	this.JobId = jobId
	this.VersionName = versionName
	this.VersionId = versionId
	this.ProjectName = projectName
	this.JobType = jobType
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.SessionName = sessionName
	this.SessionRunName = sessionRunName
	this.Events = events
	return &this
}

// NewRunProcessWithDefaults instantiates a new RunProcess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunProcessWithDefaults() *RunProcess {
	this := RunProcess{}
	return &this
}

// GetJobId returns the JobId field value
func (o *RunProcess) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *RunProcess) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *RunProcess) SetJobId(v string) {
	o.JobId = v
}

// GetVersionName returns the VersionName field value
func (o *RunProcess) GetVersionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value
// and a boolean to check if the value has been set.
func (o *RunProcess) GetVersionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionName, true
}

// SetVersionName sets field value
func (o *RunProcess) SetVersionName(v string) {
	o.VersionName = v
}

// GetVersionId returns the VersionId field value
func (o *RunProcess) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *RunProcess) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *RunProcess) SetVersionId(v string) {
	o.VersionId = v
}

// GetProjectName returns the ProjectName field value
func (o *RunProcess) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *RunProcess) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *RunProcess) SetProjectName(v string) {
	o.ProjectName = v
}

// GetJobType returns the JobType field value
func (o *RunProcess) GetJobType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *RunProcess) GetJobTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value
func (o *RunProcess) SetJobType(v string) {
	o.JobType = v
}

// GetStatus returns the Status field value
func (o *RunProcess) GetStatus() JobStatus {
	if o == nil {
		var ret JobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RunProcess) GetStatusOk() (*JobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RunProcess) SetStatus(v JobStatus) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RunProcess) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RunProcess) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RunProcess) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *RunProcess) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RunProcess) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *RunProcess) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetSessionName returns the SessionName field value
func (o *RunProcess) GetSessionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionName
}

// GetSessionNameOk returns a tuple with the SessionName field value
// and a boolean to check if the value has been set.
func (o *RunProcess) GetSessionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionName, true
}

// SetSessionName sets field value
func (o *RunProcess) SetSessionName(v string) {
	o.SessionName = v
}

// GetSessionRunName returns the SessionRunName field value
func (o *RunProcess) GetSessionRunName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunName
}

// GetSessionRunNameOk returns a tuple with the SessionRunName field value
// and a boolean to check if the value has been set.
func (o *RunProcess) GetSessionRunNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunName, true
}

// SetSessionRunName sets field value
func (o *RunProcess) SetSessionRunName(v string) {
	o.SessionRunName = v
}

// GetSessionRunId returns the SessionRunId field value if set, zero value otherwise.
func (o *RunProcess) GetSessionRunId() string {
	if o == nil || IsNil(o.SessionRunId) {
		var ret string
		return ret
	}
	return *o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunProcess) GetSessionRunIdOk() (*string, bool) {
	if o == nil || IsNil(o.SessionRunId) {
		return nil, false
	}
	return o.SessionRunId, true
}

// HasSessionRunId returns a boolean if a field has been set.
func (o *RunProcess) HasSessionRunId() bool {
	if o != nil && !IsNil(o.SessionRunId) {
		return true
	}

	return false
}

// SetSessionRunId gets a reference to the given string and assigns it to the SessionRunId field.
func (o *RunProcess) SetSessionRunId(v string) {
	o.SessionRunId = &v
}

// GetEvents returns the Events field value
func (o *RunProcess) GetEvents() []JobEvent {
	if o == nil {
		var ret []JobEvent
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *RunProcess) GetEventsOk() ([]JobEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *RunProcess) SetEvents(v []JobEvent) {
	o.Events = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *RunProcess) GetParams() JobParams {
	if o == nil || IsNil(o.Params) {
		var ret JobParams
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunProcess) GetParamsOk() (*JobParams, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *RunProcess) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given JobParams and assigns it to the Params field.
func (o *RunProcess) SetParams(v JobParams) {
	o.Params = &v
}

func (o RunProcess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunProcess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobId"] = o.JobId
	toSerialize["versionName"] = o.VersionName
	toSerialize["versionId"] = o.VersionId
	toSerialize["projectName"] = o.ProjectName
	toSerialize["jobType"] = o.JobType
	toSerialize["status"] = o.Status
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["sessionName"] = o.SessionName
	toSerialize["sessionRunName"] = o.SessionRunName
	if !IsNil(o.SessionRunId) {
		toSerialize["sessionRunId"] = o.SessionRunId
	}
	toSerialize["events"] = o.Events
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	return toSerialize, nil
}

type NullableRunProcess struct {
	value *RunProcess
	isSet bool
}

func (v NullableRunProcess) Get() *RunProcess {
	return v.value
}

func (v *NullableRunProcess) Set(val *RunProcess) {
	v.value = val
	v.isSet = true
}

func (v NullableRunProcess) IsSet() bool {
	return v.isSet
}

func (v *NullableRunProcess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunProcess(val *RunProcess) *NullableRunProcess {
	return &NullableRunProcess{value: val, isSet: true}
}

func (v NullableRunProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunProcess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
