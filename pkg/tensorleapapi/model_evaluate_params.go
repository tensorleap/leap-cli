/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.423
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the EvaluateParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvaluateParams{}

// EvaluateParams struct for EvaluateParams
type EvaluateParams struct {
	VersionId      string          `json:"versionId"`
	ProjectId      string          `json:"projectId"`
	SessionId      string          `json:"sessionId"`
	BatchSize      float64         `json:"batchSize"`
	DataStates     []DataStateType `json:"dataStates"`
	EvaluatedEpoch float64         `json:"evaluatedEpoch"`
	Name           string          `json:"name"`
}

// NewEvaluateParams instantiates a new EvaluateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvaluateParams(versionId string, projectId string, sessionId string, batchSize float64, dataStates []DataStateType, evaluatedEpoch float64, name string) *EvaluateParams {
	this := EvaluateParams{}
	this.VersionId = versionId
	this.ProjectId = projectId
	this.SessionId = sessionId
	this.BatchSize = batchSize
	this.DataStates = dataStates
	this.EvaluatedEpoch = evaluatedEpoch
	this.Name = name
	return &this
}

// NewEvaluateParamsWithDefaults instantiates a new EvaluateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvaluateParamsWithDefaults() *EvaluateParams {
	this := EvaluateParams{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *EvaluateParams) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *EvaluateParams) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *EvaluateParams) SetVersionId(v string) {
	o.VersionId = v
}

// GetProjectId returns the ProjectId field value
func (o *EvaluateParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *EvaluateParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *EvaluateParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSessionId returns the SessionId field value
func (o *EvaluateParams) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *EvaluateParams) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *EvaluateParams) SetSessionId(v string) {
	o.SessionId = v
}

// GetBatchSize returns the BatchSize field value
func (o *EvaluateParams) GetBatchSize() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value
// and a boolean to check if the value has been set.
func (o *EvaluateParams) GetBatchSizeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchSize, true
}

// SetBatchSize sets field value
func (o *EvaluateParams) SetBatchSize(v float64) {
	o.BatchSize = v
}

// GetDataStates returns the DataStates field value
func (o *EvaluateParams) GetDataStates() []DataStateType {
	if o == nil {
		var ret []DataStateType
		return ret
	}

	return o.DataStates
}

// GetDataStatesOk returns a tuple with the DataStates field value
// and a boolean to check if the value has been set.
func (o *EvaluateParams) GetDataStatesOk() ([]DataStateType, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataStates, true
}

// SetDataStates sets field value
func (o *EvaluateParams) SetDataStates(v []DataStateType) {
	o.DataStates = v
}

// GetEvaluatedEpoch returns the EvaluatedEpoch field value
func (o *EvaluateParams) GetEvaluatedEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.EvaluatedEpoch
}

// GetEvaluatedEpochOk returns a tuple with the EvaluatedEpoch field value
// and a boolean to check if the value has been set.
func (o *EvaluateParams) GetEvaluatedEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluatedEpoch, true
}

// SetEvaluatedEpoch sets field value
func (o *EvaluateParams) SetEvaluatedEpoch(v float64) {
	o.EvaluatedEpoch = v
}

// GetName returns the Name field value
func (o *EvaluateParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EvaluateParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EvaluateParams) SetName(v string) {
	o.Name = v
}

func (o EvaluateParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvaluateParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	toSerialize["projectId"] = o.ProjectId
	toSerialize["sessionId"] = o.SessionId
	toSerialize["batchSize"] = o.BatchSize
	toSerialize["dataStates"] = o.DataStates
	toSerialize["evaluatedEpoch"] = o.EvaluatedEpoch
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableEvaluateParams struct {
	value *EvaluateParams
	isSet bool
}

func (v NullableEvaluateParams) Get() *EvaluateParams {
	return v.value
}

func (v *NullableEvaluateParams) Set(val *EvaluateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEvaluateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEvaluateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvaluateParams(val *EvaluateParams) *NullableEvaluateParams {
	return &NullableEvaluateParams{value: val, isSet: true}
}

func (v NullableEvaluateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvaluateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
