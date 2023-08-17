/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the Insight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Insight{}

// Insight struct for Insight
type Insight struct {
	ProjectId string `json:"projectId"`
	Cid string `json:"cid"`
	SessionRunId string `json:"sessionRunId"`
	ModelEpoch float64 `json:"modelEpoch"`
	ModelExtId string `json:"modelExtId"`
	TeamId string `json:"teamId"`
	CreatedBy string `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	VisualizationUuid string `json:"visualizationUuid"`
	Data InsightType `json:"data"`
}

// NewInsight instantiates a new Insight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsight(projectId string, cid string, sessionRunId string, modelEpoch float64, modelExtId string, teamId string, createdBy string, createdAt time.Time, updatedAt time.Time, visualizationUuid string, data InsightType) *Insight {
	this := Insight{}
	this.ProjectId = projectId
	this.Cid = cid
	this.SessionRunId = sessionRunId
	this.ModelEpoch = modelEpoch
	this.ModelExtId = modelExtId
	this.TeamId = teamId
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.VisualizationUuid = visualizationUuid
	this.Data = data
	return &this
}

// NewInsightWithDefaults instantiates a new Insight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsightWithDefaults() *Insight {
	this := Insight{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *Insight) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Insight) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Insight) SetProjectId(v string) {
	o.ProjectId = v
}

// GetCid returns the Cid field value
func (o *Insight) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *Insight) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *Insight) SetCid(v string) {
	o.Cid = v
}

// GetSessionRunId returns the SessionRunId field value
func (o *Insight) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *Insight) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *Insight) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetModelEpoch returns the ModelEpoch field value
func (o *Insight) GetModelEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ModelEpoch
}

// GetModelEpochOk returns a tuple with the ModelEpoch field value
// and a boolean to check if the value has been set.
func (o *Insight) GetModelEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelEpoch, true
}

// SetModelEpoch sets field value
func (o *Insight) SetModelEpoch(v float64) {
	o.ModelEpoch = v
}

// GetModelExtId returns the ModelExtId field value
func (o *Insight) GetModelExtId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelExtId
}

// GetModelExtIdOk returns a tuple with the ModelExtId field value
// and a boolean to check if the value has been set.
func (o *Insight) GetModelExtIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelExtId, true
}

// SetModelExtId sets field value
func (o *Insight) SetModelExtId(v string) {
	o.ModelExtId = v
}

// GetTeamId returns the TeamId field value
func (o *Insight) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *Insight) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *Insight) SetTeamId(v string) {
	o.TeamId = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Insight) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Insight) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Insight) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Insight) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Insight) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Insight) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Insight) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Insight) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Insight) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetVisualizationUuid returns the VisualizationUuid field value
func (o *Insight) GetVisualizationUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VisualizationUuid
}

// GetVisualizationUuidOk returns a tuple with the VisualizationUuid field value
// and a boolean to check if the value has been set.
func (o *Insight) GetVisualizationUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisualizationUuid, true
}

// SetVisualizationUuid sets field value
func (o *Insight) SetVisualizationUuid(v string) {
	o.VisualizationUuid = v
}

// GetData returns the Data field value
func (o *Insight) GetData() InsightType {
	if o == nil {
		var ret InsightType
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Insight) GetDataOk() (*InsightType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Insight) SetData(v InsightType) {
	o.Data = v
}

func (o Insight) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Insight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["cid"] = o.Cid
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["modelEpoch"] = o.ModelEpoch
	toSerialize["modelExtId"] = o.ModelExtId
	toSerialize["teamId"] = o.TeamId
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["visualizationUuid"] = o.VisualizationUuid
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableInsight struct {
	value *Insight
	isSet bool
}

func (v NullableInsight) Get() *Insight {
	return v.value
}

func (v *NullableInsight) Set(val *Insight) {
	v.value = val
	v.isSet = true
}

func (v NullableInsight) IsSet() bool {
	return v.isSet
}

func (v *NullableInsight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsight(val *Insight) *NullableInsight {
	return &NullableInsight{value: val, isSet: true}
}

func (v NullableInsight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


