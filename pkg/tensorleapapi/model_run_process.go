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

// checks if the RunProcess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunProcess{}

// RunProcess struct for RunProcess
type RunProcess struct {
	JobId              string              `json:"jobId"`
	ProcessId          string              `json:"processId"`
	VersionName        *string             `json:"versionName,omitempty"`
	VersionId          *string             `json:"versionId,omitempty"`
	ProjectId          *string             `json:"projectId,omitempty"`
	ProjectName        *string             `json:"projectName,omitempty"`
	DatasetName        *string             `json:"datasetName,omitempty"`
	JobType            string              `json:"jobType"`
	Status             JobStatus           `json:"status"`
	CreatedAt          string              `json:"createdAt"`
	UpdatedAt          string              `json:"updatedAt"`
	SessionName        *string             `json:"sessionName,omitempty"`
	SessionRunName     *string             `json:"sessionRunName,omitempty"`
	SessionRunId       *string             `json:"sessionRunId,omitempty"`
	Events             []JobEvent          `json:"events"`
	Params             *JobParams          `json:"params,omitempty"`
	MachineType        *string             `json:"machineType,omitempty"`
	BatchSize          *float64            `json:"batchSize,omitempty"`
	DatasetVersionInfo *DatasetVersionInfo `json:"datasetVersionInfo,omitempty"`
	LogsBlobName       *string             `json:"logsBlobName,omitempty"`
}

// NewRunProcess instantiates a new RunProcess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunProcess(jobId string, processId string, jobType string, status JobStatus, createdAt string, updatedAt string, events []JobEvent) *RunProcess {
	this := RunProcess{}
	this.JobId = jobId
	this.ProcessId = processId
	this.JobType = jobType
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
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

// GetProcessId returns the ProcessId field value
func (o *RunProcess) GetProcessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value
// and a boolean to check if the value has been set.
func (o *RunProcess) GetProcessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessId, true
}

// SetProcessId sets field value
func (o *RunProcess) SetProcessId(v string) {
	o.ProcessId = v
}

// GetVersionName returns the VersionName field value if set, zero value otherwise.
func (o *RunProcess) GetVersionName() string {
	if o == nil || IsNil(o.VersionName) {
		var ret string
		return ret
	}
	return *o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunProcess) GetVersionNameOk() (*string, bool) {
	if o == nil || IsNil(o.VersionName) {
		return nil, false
	}
	return o.VersionName, true
}

// HasVersionName returns a boolean if a field has been set.
func (o *RunProcess) HasVersionName() bool {
	if o != nil && !IsNil(o.VersionName) {
		return true
	}

	return false
}

// SetVersionName gets a reference to the given string and assigns it to the VersionName field.
func (o *RunProcess) SetVersionName(v string) {
	o.VersionName = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *RunProcess) GetVersionId() string {
	if o == nil || IsNil(o.VersionId) {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunProcess) GetVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.VersionId) {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *RunProcess) HasVersionId() bool {
	if o != nil && !IsNil(o.VersionId) {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *RunProcess) SetVersionId(v string) {
	o.VersionId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *RunProcess) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunProcess) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *RunProcess) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *RunProcess) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *RunProcess) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunProcess) GetProjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *RunProcess) HasProjectName() bool {
	if o != nil && !IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *RunProcess) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetDatasetName returns the DatasetName field value if set, zero value otherwise.
func (o *RunProcess) GetDatasetName() string {
	if o == nil || IsNil(o.DatasetName) {
		var ret string
		return ret
	}
	return *o.DatasetName
}

// GetDatasetNameOk returns a tuple with the DatasetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunProcess) GetDatasetNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetName) {
		return nil, false
	}
	return o.DatasetName, true
}

// HasDatasetName returns a boolean if a field has been set.
func (o *RunProcess) HasDatasetName() bool {
	if o != nil && !IsNil(o.DatasetName) {
		return true
	}

	return false
}

// SetDatasetName gets a reference to the given string and assigns it to the DatasetName field.
func (o *RunProcess) SetDatasetName(v string) {
	o.DatasetName = &v
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

// GetSessionName returns the SessionName field value if set, zero value otherwise.
func (o *RunProcess) GetSessionName() string {
	if o == nil || IsNil(o.SessionName) {
		var ret string
		return ret
	}
	return *o.SessionName
}

// GetSessionNameOk returns a tuple with the SessionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunProcess) GetSessionNameOk() (*string, bool) {
	if o == nil || IsNil(o.SessionName) {
		return nil, false
	}
	return o.SessionName, true
}

// HasSessionName returns a boolean if a field has been set.
func (o *RunProcess) HasSessionName() bool {
	if o != nil && !IsNil(o.SessionName) {
		return true
	}

	return false
}

// SetSessionName gets a reference to the given string and assigns it to the SessionName field.
func (o *RunProcess) SetSessionName(v string) {
	o.SessionName = &v
}

// GetSessionRunName returns the SessionRunName field value if set, zero value otherwise.
func (o *RunProcess) GetSessionRunName() string {
	if o == nil || IsNil(o.SessionRunName) {
		var ret string
		return ret
	}
	return *o.SessionRunName
}

// GetSessionRunNameOk returns a tuple with the SessionRunName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunProcess) GetSessionRunNameOk() (*string, bool) {
	if o == nil || IsNil(o.SessionRunName) {
		return nil, false
	}
	return o.SessionRunName, true
}

// HasSessionRunName returns a boolean if a field has been set.
func (o *RunProcess) HasSessionRunName() bool {
	if o != nil && !IsNil(o.SessionRunName) {
		return true
	}

	return false
}

// SetSessionRunName gets a reference to the given string and assigns it to the SessionRunName field.
func (o *RunProcess) SetSessionRunName(v string) {
	o.SessionRunName = &v
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

// GetMachineType returns the MachineType field value if set, zero value otherwise.
func (o *RunProcess) GetMachineType() string {
	if o == nil || IsNil(o.MachineType) {
		var ret string
		return ret
	}
	return *o.MachineType
}

// GetMachineTypeOk returns a tuple with the MachineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunProcess) GetMachineTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MachineType) {
		return nil, false
	}
	return o.MachineType, true
}

// HasMachineType returns a boolean if a field has been set.
func (o *RunProcess) HasMachineType() bool {
	if o != nil && !IsNil(o.MachineType) {
		return true
	}

	return false
}

// SetMachineType gets a reference to the given string and assigns it to the MachineType field.
func (o *RunProcess) SetMachineType(v string) {
	o.MachineType = &v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *RunProcess) GetBatchSize() float64 {
	if o == nil || IsNil(o.BatchSize) {
		var ret float64
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunProcess) GetBatchSizeOk() (*float64, bool) {
	if o == nil || IsNil(o.BatchSize) {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *RunProcess) HasBatchSize() bool {
	if o != nil && !IsNil(o.BatchSize) {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given float64 and assigns it to the BatchSize field.
func (o *RunProcess) SetBatchSize(v float64) {
	o.BatchSize = &v
}

// GetDatasetVersionInfo returns the DatasetVersionInfo field value if set, zero value otherwise.
func (o *RunProcess) GetDatasetVersionInfo() DatasetVersionInfo {
	if o == nil || IsNil(o.DatasetVersionInfo) {
		var ret DatasetVersionInfo
		return ret
	}
	return *o.DatasetVersionInfo
}

// GetDatasetVersionInfoOk returns a tuple with the DatasetVersionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunProcess) GetDatasetVersionInfoOk() (*DatasetVersionInfo, bool) {
	if o == nil || IsNil(o.DatasetVersionInfo) {
		return nil, false
	}
	return o.DatasetVersionInfo, true
}

// HasDatasetVersionInfo returns a boolean if a field has been set.
func (o *RunProcess) HasDatasetVersionInfo() bool {
	if o != nil && !IsNil(o.DatasetVersionInfo) {
		return true
	}

	return false
}

// SetDatasetVersionInfo gets a reference to the given DatasetVersionInfo and assigns it to the DatasetVersionInfo field.
func (o *RunProcess) SetDatasetVersionInfo(v DatasetVersionInfo) {
	o.DatasetVersionInfo = &v
}

// GetLogsBlobName returns the LogsBlobName field value if set, zero value otherwise.
func (o *RunProcess) GetLogsBlobName() string {
	if o == nil || IsNil(o.LogsBlobName) {
		var ret string
		return ret
	}
	return *o.LogsBlobName
}

// GetLogsBlobNameOk returns a tuple with the LogsBlobName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunProcess) GetLogsBlobNameOk() (*string, bool) {
	if o == nil || IsNil(o.LogsBlobName) {
		return nil, false
	}
	return o.LogsBlobName, true
}

// HasLogsBlobName returns a boolean if a field has been set.
func (o *RunProcess) HasLogsBlobName() bool {
	if o != nil && !IsNil(o.LogsBlobName) {
		return true
	}

	return false
}

// SetLogsBlobName gets a reference to the given string and assigns it to the LogsBlobName field.
func (o *RunProcess) SetLogsBlobName(v string) {
	o.LogsBlobName = &v
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
	toSerialize["processId"] = o.ProcessId
	if !IsNil(o.VersionName) {
		toSerialize["versionName"] = o.VersionName
	}
	if !IsNil(o.VersionId) {
		toSerialize["versionId"] = o.VersionId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.ProjectName) {
		toSerialize["projectName"] = o.ProjectName
	}
	if !IsNil(o.DatasetName) {
		toSerialize["datasetName"] = o.DatasetName
	}
	toSerialize["jobType"] = o.JobType
	toSerialize["status"] = o.Status
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	if !IsNil(o.SessionName) {
		toSerialize["sessionName"] = o.SessionName
	}
	if !IsNil(o.SessionRunName) {
		toSerialize["sessionRunName"] = o.SessionRunName
	}
	if !IsNil(o.SessionRunId) {
		toSerialize["sessionRunId"] = o.SessionRunId
	}
	toSerialize["events"] = o.Events
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.MachineType) {
		toSerialize["machineType"] = o.MachineType
	}
	if !IsNil(o.BatchSize) {
		toSerialize["batchSize"] = o.BatchSize
	}
	if !IsNil(o.DatasetVersionInfo) {
		toSerialize["datasetVersionInfo"] = o.DatasetVersionInfo
	}
	if !IsNil(o.LogsBlobName) {
		toSerialize["logsBlobName"] = o.LogsBlobName
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
