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

// checks if the ContinueTrainParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContinueTrainParams{}

// ContinueTrainParams struct for ContinueTrainParams
type ContinueTrainParams struct {
	VersionId      string         `json:"versionId"`
	SessionId      string         `json:"sessionId"`
	ProjectId      string         `json:"projectId"`
	FromEpoch      float64        `json:"fromEpoch"`
	TrainingParams TrainingParams `json:"trainingParams"`
}

// NewContinueTrainParams instantiates a new ContinueTrainParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinueTrainParams(versionId string, sessionId string, projectId string, fromEpoch float64, trainingParams TrainingParams) *ContinueTrainParams {
	this := ContinueTrainParams{}
	this.VersionId = versionId
	this.SessionId = sessionId
	this.ProjectId = projectId
	this.FromEpoch = fromEpoch
	this.TrainingParams = trainingParams
	return &this
}

// NewContinueTrainParamsWithDefaults instantiates a new ContinueTrainParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinueTrainParamsWithDefaults() *ContinueTrainParams {
	this := ContinueTrainParams{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *ContinueTrainParams) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *ContinueTrainParams) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *ContinueTrainParams) SetVersionId(v string) {
	o.VersionId = v
}

// GetSessionId returns the SessionId field value
func (o *ContinueTrainParams) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *ContinueTrainParams) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *ContinueTrainParams) SetSessionId(v string) {
	o.SessionId = v
}

// GetProjectId returns the ProjectId field value
func (o *ContinueTrainParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ContinueTrainParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ContinueTrainParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetFromEpoch returns the FromEpoch field value
func (o *ContinueTrainParams) GetFromEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.FromEpoch
}

// GetFromEpochOk returns a tuple with the FromEpoch field value
// and a boolean to check if the value has been set.
func (o *ContinueTrainParams) GetFromEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromEpoch, true
}

// SetFromEpoch sets field value
func (o *ContinueTrainParams) SetFromEpoch(v float64) {
	o.FromEpoch = v
}

// GetTrainingParams returns the TrainingParams field value
func (o *ContinueTrainParams) GetTrainingParams() TrainingParams {
	if o == nil {
		var ret TrainingParams
		return ret
	}

	return o.TrainingParams
}

// GetTrainingParamsOk returns a tuple with the TrainingParams field value
// and a boolean to check if the value has been set.
func (o *ContinueTrainParams) GetTrainingParamsOk() (*TrainingParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrainingParams, true
}

// SetTrainingParams sets field value
func (o *ContinueTrainParams) SetTrainingParams(v TrainingParams) {
	o.TrainingParams = v
}

func (o ContinueTrainParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContinueTrainParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	toSerialize["sessionId"] = o.SessionId
	toSerialize["projectId"] = o.ProjectId
	toSerialize["fromEpoch"] = o.FromEpoch
	toSerialize["trainingParams"] = o.TrainingParams
	return toSerialize, nil
}

type NullableContinueTrainParams struct {
	value *ContinueTrainParams
	isSet bool
}

func (v NullableContinueTrainParams) Get() *ContinueTrainParams {
	return v.value
}

func (v *NullableContinueTrainParams) Set(val *ContinueTrainParams) {
	v.value = val
	v.isSet = true
}

func (v NullableContinueTrainParams) IsSet() bool {
	return v.isSet
}

func (v *NullableContinueTrainParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinueTrainParams(val *ContinueTrainParams) *NullableContinueTrainParams {
	return &NullableContinueTrainParams{value: val, isSet: true}
}

func (v NullableContinueTrainParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinueTrainParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
