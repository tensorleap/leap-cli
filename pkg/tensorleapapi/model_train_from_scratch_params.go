/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.557
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the TrainFromScratchParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrainFromScratchParams{}

// TrainFromScratchParams struct for TrainFromScratchParams
type TrainFromScratchParams struct {
	VersionId      string         `json:"versionId"`
	ProjectId      string         `json:"projectId"`
	SessionName    string         `json:"sessionName"`
	TrainingParams TrainingParams `json:"trainingParams"`
}

// NewTrainFromScratchParams instantiates a new TrainFromScratchParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrainFromScratchParams(versionId string, projectId string, sessionName string, trainingParams TrainingParams) *TrainFromScratchParams {
	this := TrainFromScratchParams{}
	this.VersionId = versionId
	this.ProjectId = projectId
	this.SessionName = sessionName
	this.TrainingParams = trainingParams
	return &this
}

// NewTrainFromScratchParamsWithDefaults instantiates a new TrainFromScratchParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrainFromScratchParamsWithDefaults() *TrainFromScratchParams {
	this := TrainFromScratchParams{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *TrainFromScratchParams) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *TrainFromScratchParams) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *TrainFromScratchParams) SetVersionId(v string) {
	o.VersionId = v
}

// GetProjectId returns the ProjectId field value
func (o *TrainFromScratchParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *TrainFromScratchParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *TrainFromScratchParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSessionName returns the SessionName field value
func (o *TrainFromScratchParams) GetSessionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionName
}

// GetSessionNameOk returns a tuple with the SessionName field value
// and a boolean to check if the value has been set.
func (o *TrainFromScratchParams) GetSessionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionName, true
}

// SetSessionName sets field value
func (o *TrainFromScratchParams) SetSessionName(v string) {
	o.SessionName = v
}

// GetTrainingParams returns the TrainingParams field value
func (o *TrainFromScratchParams) GetTrainingParams() TrainingParams {
	if o == nil {
		var ret TrainingParams
		return ret
	}

	return o.TrainingParams
}

// GetTrainingParamsOk returns a tuple with the TrainingParams field value
// and a boolean to check if the value has been set.
func (o *TrainFromScratchParams) GetTrainingParamsOk() (*TrainingParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrainingParams, true
}

// SetTrainingParams sets field value
func (o *TrainFromScratchParams) SetTrainingParams(v TrainingParams) {
	o.TrainingParams = v
}

func (o TrainFromScratchParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrainFromScratchParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	toSerialize["projectId"] = o.ProjectId
	toSerialize["sessionName"] = o.SessionName
	toSerialize["trainingParams"] = o.TrainingParams
	return toSerialize, nil
}

type NullableTrainFromScratchParams struct {
	value *TrainFromScratchParams
	isSet bool
}

func (v NullableTrainFromScratchParams) Get() *TrainFromScratchParams {
	return v.value
}

func (v *NullableTrainFromScratchParams) Set(val *TrainFromScratchParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTrainFromScratchParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTrainFromScratchParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrainFromScratchParams(val *TrainFromScratchParams) *NullableTrainFromScratchParams {
	return &NullableTrainFromScratchParams{value: val, isSet: true}
}

func (v NullableTrainFromScratchParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrainFromScratchParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
