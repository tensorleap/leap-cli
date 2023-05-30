/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.321
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the TrainFromInitialWeightsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrainFromInitialWeightsParams{}

// TrainFromInitialWeightsParams struct for TrainFromInitialWeightsParams
type TrainFromInitialWeightsParams struct {
	VersionId                      string         `json:"versionId"`
	FromSessionId                  string         `json:"fromSessionId"`
	FromEpoch                      float64        `json:"fromEpoch"`
	ModelName                      string         `json:"modelName"`
	TrainingParams                 TrainingParams `json:"trainingParams"`
	ShouldRunPopulationExploration bool           `json:"shouldRunPopulationExploration"`
}

// NewTrainFromInitialWeightsParams instantiates a new TrainFromInitialWeightsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrainFromInitialWeightsParams(versionId string, fromSessionId string, fromEpoch float64, modelName string, trainingParams TrainingParams, shouldRunPopulationExploration bool) *TrainFromInitialWeightsParams {
	this := TrainFromInitialWeightsParams{}
	this.VersionId = versionId
	this.FromSessionId = fromSessionId
	this.FromEpoch = fromEpoch
	this.ModelName = modelName
	this.TrainingParams = trainingParams
	this.ShouldRunPopulationExploration = shouldRunPopulationExploration
	return &this
}

// NewTrainFromInitialWeightsParamsWithDefaults instantiates a new TrainFromInitialWeightsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrainFromInitialWeightsParamsWithDefaults() *TrainFromInitialWeightsParams {
	this := TrainFromInitialWeightsParams{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *TrainFromInitialWeightsParams) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *TrainFromInitialWeightsParams) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *TrainFromInitialWeightsParams) SetVersionId(v string) {
	o.VersionId = v
}

// GetFromSessionId returns the FromSessionId field value
func (o *TrainFromInitialWeightsParams) GetFromSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromSessionId
}

// GetFromSessionIdOk returns a tuple with the FromSessionId field value
// and a boolean to check if the value has been set.
func (o *TrainFromInitialWeightsParams) GetFromSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromSessionId, true
}

// SetFromSessionId sets field value
func (o *TrainFromInitialWeightsParams) SetFromSessionId(v string) {
	o.FromSessionId = v
}

// GetFromEpoch returns the FromEpoch field value
func (o *TrainFromInitialWeightsParams) GetFromEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.FromEpoch
}

// GetFromEpochOk returns a tuple with the FromEpoch field value
// and a boolean to check if the value has been set.
func (o *TrainFromInitialWeightsParams) GetFromEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromEpoch, true
}

// SetFromEpoch sets field value
func (o *TrainFromInitialWeightsParams) SetFromEpoch(v float64) {
	o.FromEpoch = v
}

// GetModelName returns the ModelName field value
func (o *TrainFromInitialWeightsParams) GetModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value
// and a boolean to check if the value has been set.
func (o *TrainFromInitialWeightsParams) GetModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelName, true
}

// SetModelName sets field value
func (o *TrainFromInitialWeightsParams) SetModelName(v string) {
	o.ModelName = v
}

// GetTrainingParams returns the TrainingParams field value
func (o *TrainFromInitialWeightsParams) GetTrainingParams() TrainingParams {
	if o == nil {
		var ret TrainingParams
		return ret
	}

	return o.TrainingParams
}

// GetTrainingParamsOk returns a tuple with the TrainingParams field value
// and a boolean to check if the value has been set.
func (o *TrainFromInitialWeightsParams) GetTrainingParamsOk() (*TrainingParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrainingParams, true
}

// SetTrainingParams sets field value
func (o *TrainFromInitialWeightsParams) SetTrainingParams(v TrainingParams) {
	o.TrainingParams = v
}

// GetShouldRunPopulationExploration returns the ShouldRunPopulationExploration field value
func (o *TrainFromInitialWeightsParams) GetShouldRunPopulationExploration() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShouldRunPopulationExploration
}

// GetShouldRunPopulationExplorationOk returns a tuple with the ShouldRunPopulationExploration field value
// and a boolean to check if the value has been set.
func (o *TrainFromInitialWeightsParams) GetShouldRunPopulationExplorationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShouldRunPopulationExploration, true
}

// SetShouldRunPopulationExploration sets field value
func (o *TrainFromInitialWeightsParams) SetShouldRunPopulationExploration(v bool) {
	o.ShouldRunPopulationExploration = v
}

func (o TrainFromInitialWeightsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrainFromInitialWeightsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	toSerialize["fromSessionId"] = o.FromSessionId
	toSerialize["fromEpoch"] = o.FromEpoch
	toSerialize["modelName"] = o.ModelName
	toSerialize["trainingParams"] = o.TrainingParams
	toSerialize["shouldRunPopulationExploration"] = o.ShouldRunPopulationExploration
	return toSerialize, nil
}

type NullableTrainFromInitialWeightsParams struct {
	value *TrainFromInitialWeightsParams
	isSet bool
}

func (v NullableTrainFromInitialWeightsParams) Get() *TrainFromInitialWeightsParams {
	return v.value
}

func (v *NullableTrainFromInitialWeightsParams) Set(val *TrainFromInitialWeightsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTrainFromInitialWeightsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTrainFromInitialWeightsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrainFromInitialWeightsParams(val *TrainFromInitialWeightsParams) *NullableTrainFromInitialWeightsParams {
	return &NullableTrainFromInitialWeightsParams{value: val, isSet: true}
}

func (v NullableTrainFromInitialWeightsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrainFromInitialWeightsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
