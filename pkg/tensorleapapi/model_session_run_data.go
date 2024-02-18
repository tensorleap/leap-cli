/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.456
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the SessionRunData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionRunData{}

// SessionRunData struct for SessionRunData
type SessionRunData struct {
	ProjectId    string            `json:"projectId"`
	Cid          string            `json:"cid"`
	SessionId    string            `json:"sessionId"`
	Name         string            `json:"name"`
	TeamId       string            `json:"teamId"`
	IsEvaluate   bool              `json:"isEvaluate"`
	CreatedAt    time.Time         `json:"createdAt"`
	CreatedBy    string            `json:"createdBy"`
	WeightAssets []WeightAssetData `json:"weightAssets"`
	Jobs         []Job             `json:"jobs"`
}

// NewSessionRunData instantiates a new SessionRunData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionRunData(projectId string, cid string, sessionId string, name string, teamId string, isEvaluate bool, createdAt time.Time, createdBy string, weightAssets []WeightAssetData, jobs []Job) *SessionRunData {
	this := SessionRunData{}
	this.ProjectId = projectId
	this.Cid = cid
	this.SessionId = sessionId
	this.Name = name
	this.TeamId = teamId
	this.IsEvaluate = isEvaluate
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.WeightAssets = weightAssets
	this.Jobs = jobs
	return &this
}

// NewSessionRunDataWithDefaults instantiates a new SessionRunData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionRunDataWithDefaults() *SessionRunData {
	this := SessionRunData{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *SessionRunData) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *SessionRunData) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *SessionRunData) SetProjectId(v string) {
	o.ProjectId = v
}

// GetCid returns the Cid field value
func (o *SessionRunData) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *SessionRunData) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *SessionRunData) SetCid(v string) {
	o.Cid = v
}

// GetSessionId returns the SessionId field value
func (o *SessionRunData) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *SessionRunData) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *SessionRunData) SetSessionId(v string) {
	o.SessionId = v
}

// GetName returns the Name field value
func (o *SessionRunData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SessionRunData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SessionRunData) SetName(v string) {
	o.Name = v
}

// GetTeamId returns the TeamId field value
func (o *SessionRunData) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *SessionRunData) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *SessionRunData) SetTeamId(v string) {
	o.TeamId = v
}

// GetIsEvaluate returns the IsEvaluate field value
func (o *SessionRunData) GetIsEvaluate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEvaluate
}

// GetIsEvaluateOk returns a tuple with the IsEvaluate field value
// and a boolean to check if the value has been set.
func (o *SessionRunData) GetIsEvaluateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEvaluate, true
}

// SetIsEvaluate sets field value
func (o *SessionRunData) SetIsEvaluate(v bool) {
	o.IsEvaluate = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SessionRunData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SessionRunData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SessionRunData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *SessionRunData) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SessionRunData) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *SessionRunData) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetWeightAssets returns the WeightAssets field value
func (o *SessionRunData) GetWeightAssets() []WeightAssetData {
	if o == nil {
		var ret []WeightAssetData
		return ret
	}

	return o.WeightAssets
}

// GetWeightAssetsOk returns a tuple with the WeightAssets field value
// and a boolean to check if the value has been set.
func (o *SessionRunData) GetWeightAssetsOk() ([]WeightAssetData, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeightAssets, true
}

// SetWeightAssets sets field value
func (o *SessionRunData) SetWeightAssets(v []WeightAssetData) {
	o.WeightAssets = v
}

// GetJobs returns the Jobs field value
func (o *SessionRunData) GetJobs() []Job {
	if o == nil {
		var ret []Job
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *SessionRunData) GetJobsOk() ([]Job, bool) {
	if o == nil {
		return nil, false
	}
	return o.Jobs, true
}

// SetJobs sets field value
func (o *SessionRunData) SetJobs(v []Job) {
	o.Jobs = v
}

func (o SessionRunData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionRunData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["cid"] = o.Cid
	toSerialize["sessionId"] = o.SessionId
	toSerialize["name"] = o.Name
	toSerialize["teamId"] = o.TeamId
	toSerialize["isEvaluate"] = o.IsEvaluate
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["weightAssets"] = o.WeightAssets
	toSerialize["jobs"] = o.Jobs
	return toSerialize, nil
}

type NullableSessionRunData struct {
	value *SessionRunData
	isSet bool
}

func (v NullableSessionRunData) Get() *SessionRunData {
	return v.value
}

func (v *NullableSessionRunData) Set(val *SessionRunData) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionRunData) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionRunData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionRunData(val *SessionRunData) *NullableSessionRunData {
	return &NullableSessionRunData{value: val, isSet: true}
}

func (v NullableSessionRunData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionRunData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
