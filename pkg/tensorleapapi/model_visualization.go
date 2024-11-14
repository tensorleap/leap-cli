/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.557
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the Visualization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Visualization{}

// Visualization struct for Visualization
type Visualization struct {
	ProjectId         string                `json:"projectId"`
	Cid               string                `json:"cid"`
	JobId             string                `json:"jobId"`
	SessionRunId      string                `json:"sessionRunId"`
	JobParms          *JobParams            `json:"jobParms,omitempty"`
	Type              AnalyzeTypeEnum       `json:"type"`
	CreatedAt         time.Time             `json:"createdAt"`
	Epoch             float64               `json:"epoch"`
	SampleId          *float64              `json:"sampleId,omitempty"`
	Layout            *SizedLayout          `json:"layout,omitempty"`
	VisualizationUuid string                `json:"visualizationUuid"`
	Blob              string                `json:"blob"`
	CsvBlob           string                `json:"csvBlob"`
	Data              VisualizationResponse `json:"data"`
}

// NewVisualization instantiates a new Visualization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisualization(projectId string, cid string, jobId string, sessionRunId string, type_ AnalyzeTypeEnum, createdAt time.Time, epoch float64, visualizationUuid string, blob string, csvBlob string, data VisualizationResponse) *Visualization {
	this := Visualization{}
	this.ProjectId = projectId
	this.Cid = cid
	this.JobId = jobId
	this.SessionRunId = sessionRunId
	this.Type = type_
	this.CreatedAt = createdAt
	this.Epoch = epoch
	this.VisualizationUuid = visualizationUuid
	this.Blob = blob
	this.CsvBlob = csvBlob
	this.Data = data
	return &this
}

// NewVisualizationWithDefaults instantiates a new Visualization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisualizationWithDefaults() *Visualization {
	this := Visualization{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *Visualization) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Visualization) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Visualization) SetProjectId(v string) {
	o.ProjectId = v
}

// GetCid returns the Cid field value
func (o *Visualization) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *Visualization) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *Visualization) SetCid(v string) {
	o.Cid = v
}

// GetJobId returns the JobId field value
func (o *Visualization) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *Visualization) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *Visualization) SetJobId(v string) {
	o.JobId = v
}

// GetSessionRunId returns the SessionRunId field value
func (o *Visualization) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *Visualization) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *Visualization) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetJobParms returns the JobParms field value if set, zero value otherwise.
func (o *Visualization) GetJobParms() JobParams {
	if o == nil || IsNil(o.JobParms) {
		var ret JobParams
		return ret
	}
	return *o.JobParms
}

// GetJobParmsOk returns a tuple with the JobParms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Visualization) GetJobParmsOk() (*JobParams, bool) {
	if o == nil || IsNil(o.JobParms) {
		return nil, false
	}
	return o.JobParms, true
}

// HasJobParms returns a boolean if a field has been set.
func (o *Visualization) HasJobParms() bool {
	if o != nil && !IsNil(o.JobParms) {
		return true
	}

	return false
}

// SetJobParms gets a reference to the given JobParams and assigns it to the JobParms field.
func (o *Visualization) SetJobParms(v JobParams) {
	o.JobParms = &v
}

// GetType returns the Type field value
func (o *Visualization) GetType() AnalyzeTypeEnum {
	if o == nil {
		var ret AnalyzeTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Visualization) GetTypeOk() (*AnalyzeTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Visualization) SetType(v AnalyzeTypeEnum) {
	o.Type = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Visualization) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Visualization) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Visualization) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEpoch returns the Epoch field value
func (o *Visualization) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *Visualization) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *Visualization) SetEpoch(v float64) {
	o.Epoch = v
}

// GetSampleId returns the SampleId field value if set, zero value otherwise.
func (o *Visualization) GetSampleId() float64 {
	if o == nil || IsNil(o.SampleId) {
		var ret float64
		return ret
	}
	return *o.SampleId
}

// GetSampleIdOk returns a tuple with the SampleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Visualization) GetSampleIdOk() (*float64, bool) {
	if o == nil || IsNil(o.SampleId) {
		return nil, false
	}
	return o.SampleId, true
}

// HasSampleId returns a boolean if a field has been set.
func (o *Visualization) HasSampleId() bool {
	if o != nil && !IsNil(o.SampleId) {
		return true
	}

	return false
}

// SetSampleId gets a reference to the given float64 and assigns it to the SampleId field.
func (o *Visualization) SetSampleId(v float64) {
	o.SampleId = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *Visualization) GetLayout() SizedLayout {
	if o == nil || IsNil(o.Layout) {
		var ret SizedLayout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Visualization) GetLayoutOk() (*SizedLayout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *Visualization) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given SizedLayout and assigns it to the Layout field.
func (o *Visualization) SetLayout(v SizedLayout) {
	o.Layout = &v
}

// GetVisualizationUuid returns the VisualizationUuid field value
func (o *Visualization) GetVisualizationUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VisualizationUuid
}

// GetVisualizationUuidOk returns a tuple with the VisualizationUuid field value
// and a boolean to check if the value has been set.
func (o *Visualization) GetVisualizationUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisualizationUuid, true
}

// SetVisualizationUuid sets field value
func (o *Visualization) SetVisualizationUuid(v string) {
	o.VisualizationUuid = v
}

// GetBlob returns the Blob field value
func (o *Visualization) GetBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blob
}

// GetBlobOk returns a tuple with the Blob field value
// and a boolean to check if the value has been set.
func (o *Visualization) GetBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blob, true
}

// SetBlob sets field value
func (o *Visualization) SetBlob(v string) {
	o.Blob = v
}

// GetCsvBlob returns the CsvBlob field value
func (o *Visualization) GetCsvBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CsvBlob
}

// GetCsvBlobOk returns a tuple with the CsvBlob field value
// and a boolean to check if the value has been set.
func (o *Visualization) GetCsvBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CsvBlob, true
}

// SetCsvBlob sets field value
func (o *Visualization) SetCsvBlob(v string) {
	o.CsvBlob = v
}

// GetData returns the Data field value
func (o *Visualization) GetData() VisualizationResponse {
	if o == nil {
		var ret VisualizationResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Visualization) GetDataOk() (*VisualizationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Visualization) SetData(v VisualizationResponse) {
	o.Data = v
}

func (o Visualization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Visualization) ToMap() (map[string]interface{}, error) {
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
	toSerialize["csvBlob"] = o.CsvBlob
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableVisualization struct {
	value *Visualization
	isSet bool
}

func (v NullableVisualization) Get() *Visualization {
	return v.value
}

func (v *NullableVisualization) Set(val *Visualization) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualization) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualization(val *Visualization) *NullableVisualization {
	return &NullableVisualization{value: val, isSet: true}
}

func (v NullableVisualization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
