/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.592
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the TrainingParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrainingParams{}

// TrainingParams struct for TrainingParams
type TrainingParams struct {
	Epochs          float64          `json:"epochs"`
	BatchSize       float64          `json:"batch_size"`
	EarlyStopParams *EarlyStopParams `json:"early_stop_params,omitempty"`
}

// NewTrainingParams instantiates a new TrainingParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrainingParams(epochs float64, batchSize float64) *TrainingParams {
	this := TrainingParams{}
	this.Epochs = epochs
	this.BatchSize = batchSize
	return &this
}

// NewTrainingParamsWithDefaults instantiates a new TrainingParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrainingParamsWithDefaults() *TrainingParams {
	this := TrainingParams{}
	return &this
}

// GetEpochs returns the Epochs field value
func (o *TrainingParams) GetEpochs() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epochs
}

// GetEpochsOk returns a tuple with the Epochs field value
// and a boolean to check if the value has been set.
func (o *TrainingParams) GetEpochsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epochs, true
}

// SetEpochs sets field value
func (o *TrainingParams) SetEpochs(v float64) {
	o.Epochs = v
}

// GetBatchSize returns the BatchSize field value
func (o *TrainingParams) GetBatchSize() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value
// and a boolean to check if the value has been set.
func (o *TrainingParams) GetBatchSizeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchSize, true
}

// SetBatchSize sets field value
func (o *TrainingParams) SetBatchSize(v float64) {
	o.BatchSize = v
}

// GetEarlyStopParams returns the EarlyStopParams field value if set, zero value otherwise.
func (o *TrainingParams) GetEarlyStopParams() EarlyStopParams {
	if o == nil || IsNil(o.EarlyStopParams) {
		var ret EarlyStopParams
		return ret
	}
	return *o.EarlyStopParams
}

// GetEarlyStopParamsOk returns a tuple with the EarlyStopParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrainingParams) GetEarlyStopParamsOk() (*EarlyStopParams, bool) {
	if o == nil || IsNil(o.EarlyStopParams) {
		return nil, false
	}
	return o.EarlyStopParams, true
}

// HasEarlyStopParams returns a boolean if a field has been set.
func (o *TrainingParams) HasEarlyStopParams() bool {
	if o != nil && !IsNil(o.EarlyStopParams) {
		return true
	}

	return false
}

// SetEarlyStopParams gets a reference to the given EarlyStopParams and assigns it to the EarlyStopParams field.
func (o *TrainingParams) SetEarlyStopParams(v EarlyStopParams) {
	o.EarlyStopParams = &v
}

func (o TrainingParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrainingParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["epochs"] = o.Epochs
	toSerialize["batch_size"] = o.BatchSize
	if !IsNil(o.EarlyStopParams) {
		toSerialize["early_stop_params"] = o.EarlyStopParams
	}
	return toSerialize, nil
}

type NullableTrainingParams struct {
	value *TrainingParams
	isSet bool
}

func (v NullableTrainingParams) Get() *TrainingParams {
	return v.value
}

func (v *NullableTrainingParams) Set(val *TrainingParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTrainingParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTrainingParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrainingParams(val *TrainingParams) *NullableTrainingParams {
	return &NullableTrainingParams{value: val, isSet: true}
}

func (v NullableTrainingParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrainingParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
