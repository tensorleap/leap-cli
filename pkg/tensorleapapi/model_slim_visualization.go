/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the SlimVisualization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlimVisualization{}

// SlimVisualization struct for SlimVisualization
type SlimVisualization struct {
	ProjectId         string          `json:"projectId"`
	Cid               string          `json:"cid"`
	JobId             string          `json:"jobId"`
	SessionRunId      string          `json:"sessionRunId"`
	JobParms          *JobParams      `json:"jobParms,omitempty"`
	Type              AnalyzeTypeEnum `json:"type"`
	CreatedAt         time.Time       `json:"createdAt"`
	Epoch             float64         `json:"epoch"`
	SampleId          *float64        `json:"sampleId,omitempty"`
	Layout            *SizedLayout    `json:"layout,omitempty"`
	VisualizationUuid string          `json:"visualizationUuid"`
	Blob              string          `json:"blob"`
}

// NewSlimVisualization instantiates a new SlimVisualization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlimVisualization(projectId string, cid string, jobId string, sessionRunId string, type_ AnalyzeTypeEnum, createdAt time.Time, epoch float64, visualizationUuid string, blob string) *SlimVisualization {
	this := SlimVisualization{}
	this.ProjectId = projectId
	this.Cid = cid
	this.JobId = jobId
	this.SessionRunId = sessionRunId
	this.Type = type_
	this.CreatedAt = createdAt
	this.Epoch = epoch
	this.VisualizationUuid = visualizationUuid
	this.Blob = blob
	return &this
}

// NewSlimVisualizationWithDefaults instantiates a new SlimVisualization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlimVisualizationWithDefaults() *SlimVisualization {
	this := SlimVisualization{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *SlimVisualization) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *SlimVisualization) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *SlimVisualization) SetProjectId(v string) {
	o.ProjectId = v
}

// GetCid returns the Cid field value
func (o *SlimVisualization) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *SlimVisualization) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *SlimVisualization) SetCid(v string) {
	o.Cid = v
}

// GetJobId returns the JobId field value
func (o *SlimVisualization) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *SlimVisualization) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *SlimVisualization) SetJobId(v string) {
	o.JobId = v
}

// GetSessionRunId returns the SessionRunId field value
func (o *SlimVisualization) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *SlimVisualization) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *SlimVisualization) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetJobParms returns the JobParms field value if set, zero value otherwise.
func (o *SlimVisualization) GetJobParms() JobParams {
	if o == nil || IsNil(o.JobParms) {
		var ret JobParams
		return ret
	}
	return *o.JobParms
}

// GetJobParmsOk returns a tuple with the JobParms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimVisualization) GetJobParmsOk() (*JobParams, bool) {
	if o == nil || IsNil(o.JobParms) {
		return nil, false
	}
	return o.JobParms, true
}

// HasJobParms returns a boolean if a field has been set.
func (o *SlimVisualization) HasJobParms() bool {
	if o != nil && !IsNil(o.JobParms) {
		return true
	}

	return false
}

// SetJobParms gets a reference to the given JobParams and assigns it to the JobParms field.
func (o *SlimVisualization) SetJobParms(v JobParams) {
	o.JobParms = &v
}

// GetType returns the Type field value
func (o *SlimVisualization) GetType() AnalyzeTypeEnum {
	if o == nil {
		var ret AnalyzeTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SlimVisualization) GetTypeOk() (*AnalyzeTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SlimVisualization) SetType(v AnalyzeTypeEnum) {
	o.Type = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SlimVisualization) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SlimVisualization) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SlimVisualization) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEpoch returns the Epoch field value
func (o *SlimVisualization) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *SlimVisualization) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *SlimVisualization) SetEpoch(v float64) {
	o.Epoch = v
}

// GetSampleId returns the SampleId field value if set, zero value otherwise.
func (o *SlimVisualization) GetSampleId() float64 {
	if o == nil || IsNil(o.SampleId) {
		var ret float64
		return ret
	}
	return *o.SampleId
}

// GetSampleIdOk returns a tuple with the SampleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimVisualization) GetSampleIdOk() (*float64, bool) {
	if o == nil || IsNil(o.SampleId) {
		return nil, false
	}
	return o.SampleId, true
}

// HasSampleId returns a boolean if a field has been set.
func (o *SlimVisualization) HasSampleId() bool {
	if o != nil && !IsNil(o.SampleId) {
		return true
	}

	return false
}

// SetSampleId gets a reference to the given float64 and assigns it to the SampleId field.
func (o *SlimVisualization) SetSampleId(v float64) {
	o.SampleId = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *SlimVisualization) GetLayout() SizedLayout {
	if o == nil || IsNil(o.Layout) {
		var ret SizedLayout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimVisualization) GetLayoutOk() (*SizedLayout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *SlimVisualization) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given SizedLayout and assigns it to the Layout field.
func (o *SlimVisualization) SetLayout(v SizedLayout) {
	o.Layout = &v
}

// GetVisualizationUuid returns the VisualizationUuid field value
func (o *SlimVisualization) GetVisualizationUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VisualizationUuid
}

// GetVisualizationUuidOk returns a tuple with the VisualizationUuid field value
// and a boolean to check if the value has been set.
func (o *SlimVisualization) GetVisualizationUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisualizationUuid, true
}

// SetVisualizationUuid sets field value
func (o *SlimVisualization) SetVisualizationUuid(v string) {
	o.VisualizationUuid = v
}

// GetBlob returns the Blob field value
func (o *SlimVisualization) GetBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blob
}

// GetBlobOk returns a tuple with the Blob field value
// and a boolean to check if the value has been set.
func (o *SlimVisualization) GetBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blob, true
}

// SetBlob sets field value
func (o *SlimVisualization) SetBlob(v string) {
	o.Blob = v
}

func (o SlimVisualization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlimVisualization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["cid"] = o.Cid
	toSerialize["jobId"] = o.JobId
	toSerialize["sessionRunId"] = o.SessionRunId
	if !IsNil(o.JobParms) {
		toSerialize["jobParms"] = o.JobParms
	}
	toSerialize["type"] = o.Type
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["epoch"] = o.Epoch
	if !IsNil(o.SampleId) {
		toSerialize["sampleId"] = o.SampleId
	}
	if !IsNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}
	toSerialize["visualizationUuid"] = o.VisualizationUuid
	toSerialize["blob"] = o.Blob
	return toSerialize, nil
}

type NullableSlimVisualization struct {
	value *SlimVisualization
	isSet bool
}

func (v NullableSlimVisualization) Get() *SlimVisualization {
	return v.value
}

func (v *NullableSlimVisualization) Set(val *SlimVisualization) {
	v.value = val
	v.isSet = true
}

func (v NullableSlimVisualization) IsSet() bool {
	return v.isSet
}

func (v *NullableSlimVisualization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlimVisualization(val *SlimVisualization) *NullableSlimVisualization {
	return &NullableSlimVisualization{value: val, isSet: true}
}

func (v NullableSlimVisualization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlimVisualization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
