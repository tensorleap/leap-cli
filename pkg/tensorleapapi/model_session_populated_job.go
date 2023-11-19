/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.404
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the SessionPopulatedJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionPopulatedJob{}

// SessionPopulatedJob struct for SessionPopulatedJob
type SessionPopulatedJob struct {
	Cid            string              `json:"cid"`
	ExtId          *string             `json:"extId,omitempty"`
	ModelName      string              `json:"modelName"`
	CreatedAt      time.Time           `json:"createdAt"`
	CreatedBy      *string             `json:"createdBy,omitempty"`
	TeamId         string              `json:"teamId"`
	Hash           NullableString      `json:"hash,omitempty"`
	TrainingParams *TrainingParams     `json:"trainingParams,omitempty"`
	SessionRuns    []SessionRunData    `json:"sessionRuns,omitempty"`
	SessionWeights []SessionWeightData `json:"sessionWeights,omitempty"`
}

// NewSessionPopulatedJob instantiates a new SessionPopulatedJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionPopulatedJob(cid string, modelName string, createdAt time.Time, teamId string) *SessionPopulatedJob {
	this := SessionPopulatedJob{}
	this.Cid = cid
	this.ModelName = modelName
	this.CreatedAt = createdAt
	this.TeamId = teamId
	return &this
}

// NewSessionPopulatedJobWithDefaults instantiates a new SessionPopulatedJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionPopulatedJobWithDefaults() *SessionPopulatedJob {
	this := SessionPopulatedJob{}
	return &this
}

// GetCid returns the Cid field value
func (o *SessionPopulatedJob) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *SessionPopulatedJob) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *SessionPopulatedJob) SetCid(v string) {
	o.Cid = v
}

// GetExtId returns the ExtId field value if set, zero value otherwise.
func (o *SessionPopulatedJob) GetExtId() string {
	if o == nil || IsNil(o.ExtId) {
		var ret string
		return ret
	}
	return *o.ExtId
}

// GetExtIdOk returns a tuple with the ExtId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionPopulatedJob) GetExtIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExtId) {
		return nil, false
	}
	return o.ExtId, true
}

// HasExtId returns a boolean if a field has been set.
func (o *SessionPopulatedJob) HasExtId() bool {
	if o != nil && !IsNil(o.ExtId) {
		return true
	}

	return false
}

// SetExtId gets a reference to the given string and assigns it to the ExtId field.
func (o *SessionPopulatedJob) SetExtId(v string) {
	o.ExtId = &v
}

// GetModelName returns the ModelName field value
func (o *SessionPopulatedJob) GetModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value
// and a boolean to check if the value has been set.
func (o *SessionPopulatedJob) GetModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelName, true
}

// SetModelName sets field value
func (o *SessionPopulatedJob) SetModelName(v string) {
	o.ModelName = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SessionPopulatedJob) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SessionPopulatedJob) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SessionPopulatedJob) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SessionPopulatedJob) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionPopulatedJob) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SessionPopulatedJob) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SessionPopulatedJob) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetTeamId returns the TeamId field value
func (o *SessionPopulatedJob) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *SessionPopulatedJob) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *SessionPopulatedJob) SetTeamId(v string) {
	o.TeamId = v
}

// GetHash returns the Hash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionPopulatedJob) GetHash() string {
	if o == nil || IsNil(o.Hash.Get()) {
		var ret string
		return ret
	}
	return *o.Hash.Get()
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionPopulatedJob) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hash.Get(), o.Hash.IsSet()
}

// HasHash returns a boolean if a field has been set.
func (o *SessionPopulatedJob) HasHash() bool {
	if o != nil && o.Hash.IsSet() {
		return true
	}

	return false
}

// SetHash gets a reference to the given NullableString and assigns it to the Hash field.
func (o *SessionPopulatedJob) SetHash(v string) {
	o.Hash.Set(&v)
}

// SetHashNil sets the value for Hash to be an explicit nil
func (o *SessionPopulatedJob) SetHashNil() {
	o.Hash.Set(nil)
}

// UnsetHash ensures that no value is present for Hash, not even an explicit nil
func (o *SessionPopulatedJob) UnsetHash() {
	o.Hash.Unset()
}

// GetTrainingParams returns the TrainingParams field value if set, zero value otherwise.
func (o *SessionPopulatedJob) GetTrainingParams() TrainingParams {
	if o == nil || IsNil(o.TrainingParams) {
		var ret TrainingParams
		return ret
	}
	return *o.TrainingParams
}

// GetTrainingParamsOk returns a tuple with the TrainingParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionPopulatedJob) GetTrainingParamsOk() (*TrainingParams, bool) {
	if o == nil || IsNil(o.TrainingParams) {
		return nil, false
	}
	return o.TrainingParams, true
}

// HasTrainingParams returns a boolean if a field has been set.
func (o *SessionPopulatedJob) HasTrainingParams() bool {
	if o != nil && !IsNil(o.TrainingParams) {
		return true
	}

	return false
}

// SetTrainingParams gets a reference to the given TrainingParams and assigns it to the TrainingParams field.
func (o *SessionPopulatedJob) SetTrainingParams(v TrainingParams) {
	o.TrainingParams = &v
}

// GetSessionRuns returns the SessionRuns field value if set, zero value otherwise.
func (o *SessionPopulatedJob) GetSessionRuns() []SessionRunData {
	if o == nil || IsNil(o.SessionRuns) {
		var ret []SessionRunData
		return ret
	}
	return o.SessionRuns
}

// GetSessionRunsOk returns a tuple with the SessionRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionPopulatedJob) GetSessionRunsOk() ([]SessionRunData, bool) {
	if o == nil || IsNil(o.SessionRuns) {
		return nil, false
	}
	return o.SessionRuns, true
}

// HasSessionRuns returns a boolean if a field has been set.
func (o *SessionPopulatedJob) HasSessionRuns() bool {
	if o != nil && !IsNil(o.SessionRuns) {
		return true
	}

	return false
}

// SetSessionRuns gets a reference to the given []SessionRunData and assigns it to the SessionRuns field.
func (o *SessionPopulatedJob) SetSessionRuns(v []SessionRunData) {
	o.SessionRuns = v
}

// GetSessionWeights returns the SessionWeights field value if set, zero value otherwise.
func (o *SessionPopulatedJob) GetSessionWeights() []SessionWeightData {
	if o == nil || IsNil(o.SessionWeights) {
		var ret []SessionWeightData
		return ret
	}
	return o.SessionWeights
}

// GetSessionWeightsOk returns a tuple with the SessionWeights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionPopulatedJob) GetSessionWeightsOk() ([]SessionWeightData, bool) {
	if o == nil || IsNil(o.SessionWeights) {
		return nil, false
	}
	return o.SessionWeights, true
}

// HasSessionWeights returns a boolean if a field has been set.
func (o *SessionPopulatedJob) HasSessionWeights() bool {
	if o != nil && !IsNil(o.SessionWeights) {
		return true
	}

	return false
}

// SetSessionWeights gets a reference to the given []SessionWeightData and assigns it to the SessionWeights field.
func (o *SessionPopulatedJob) SetSessionWeights(v []SessionWeightData) {
	o.SessionWeights = v
}

func (o SessionPopulatedJob) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionPopulatedJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	if !IsNil(o.ExtId) {
		toSerialize["extId"] = o.ExtId
	}
	toSerialize["modelName"] = o.ModelName
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	toSerialize["teamId"] = o.TeamId
	if o.Hash.IsSet() {
		toSerialize["hash"] = o.Hash.Get()
	}
	if !IsNil(o.TrainingParams) {
		toSerialize["trainingParams"] = o.TrainingParams
	}
	if !IsNil(o.SessionRuns) {
		toSerialize["sessionRuns"] = o.SessionRuns
	}
	if !IsNil(o.SessionWeights) {
		toSerialize["sessionWeights"] = o.SessionWeights
	}
	return toSerialize, nil
}

type NullableSessionPopulatedJob struct {
	value *SessionPopulatedJob
	isSet bool
}

func (v NullableSessionPopulatedJob) Get() *SessionPopulatedJob {
	return v.value
}

func (v *NullableSessionPopulatedJob) Set(val *SessionPopulatedJob) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionPopulatedJob) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionPopulatedJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionPopulatedJob(val *SessionPopulatedJob) *NullableSessionPopulatedJob {
	return &NullableSessionPopulatedJob{value: val, isSet: true}
}

func (v NullableSessionPopulatedJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionPopulatedJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
