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

// checks if the DatasetSetup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetSetup{}

// DatasetSetup struct for DatasetSetup
type DatasetSetup struct {
	Preprocess      DatasetPreprocess         `json:"preprocess"`
	Inputs          []DatasetInputInstance    `json:"inputs"`
	Metadata        []DatasetMetadataInstance `json:"metadata"`
	Outputs         []DatasetOutputInstance   `json:"outputs"`
	Visualizers     []VisualizerInstance      `json:"visualizers"`
	PredictionTypes []PredictionTypeInstance  `json:"prediction_types"`
	CustomLosses    []CustomLossInstance      `json:"custom_losses"`
	Metrics         []MetricInstance          `json:"metrics"`
}

// NewDatasetSetup instantiates a new DatasetSetup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetSetup(preprocess DatasetPreprocess, inputs []DatasetInputInstance, metadata []DatasetMetadataInstance, outputs []DatasetOutputInstance, visualizers []VisualizerInstance, predictionTypes []PredictionTypeInstance, customLosses []CustomLossInstance, metrics []MetricInstance) *DatasetSetup {
	this := DatasetSetup{}
	this.Preprocess = preprocess
	this.Inputs = inputs
	this.Metadata = metadata
	this.Outputs = outputs
	this.Visualizers = visualizers
	this.PredictionTypes = predictionTypes
	this.CustomLosses = customLosses
	this.Metrics = metrics
	return &this
}

// NewDatasetSetupWithDefaults instantiates a new DatasetSetup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetSetupWithDefaults() *DatasetSetup {
	this := DatasetSetup{}
	return &this
}

// GetPreprocess returns the Preprocess field value
func (o *DatasetSetup) GetPreprocess() DatasetPreprocess {
	if o == nil {
		var ret DatasetPreprocess
		return ret
	}

	return o.Preprocess
}

// GetPreprocessOk returns a tuple with the Preprocess field value
// and a boolean to check if the value has been set.
func (o *DatasetSetup) GetPreprocessOk() (*DatasetPreprocess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Preprocess, true
}

// SetPreprocess sets field value
func (o *DatasetSetup) SetPreprocess(v DatasetPreprocess) {
	o.Preprocess = v
}

// GetInputs returns the Inputs field value
func (o *DatasetSetup) GetInputs() []DatasetInputInstance {
	if o == nil {
		var ret []DatasetInputInstance
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *DatasetSetup) GetInputsOk() ([]DatasetInputInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *DatasetSetup) SetInputs(v []DatasetInputInstance) {
	o.Inputs = v
}

// GetMetadata returns the Metadata field value
func (o *DatasetSetup) GetMetadata() []DatasetMetadataInstance {
	if o == nil {
		var ret []DatasetMetadataInstance
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *DatasetSetup) GetMetadataOk() ([]DatasetMetadataInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *DatasetSetup) SetMetadata(v []DatasetMetadataInstance) {
	o.Metadata = v
}

// GetOutputs returns the Outputs field value
func (o *DatasetSetup) GetOutputs() []DatasetOutputInstance {
	if o == nil {
		var ret []DatasetOutputInstance
		return ret
	}

	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value
// and a boolean to check if the value has been set.
func (o *DatasetSetup) GetOutputsOk() ([]DatasetOutputInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outputs, true
}

// SetOutputs sets field value
func (o *DatasetSetup) SetOutputs(v []DatasetOutputInstance) {
	o.Outputs = v
}

// GetVisualizers returns the Visualizers field value
func (o *DatasetSetup) GetVisualizers() []VisualizerInstance {
	if o == nil {
		var ret []VisualizerInstance
		return ret
	}

	return o.Visualizers
}

// GetVisualizersOk returns a tuple with the Visualizers field value
// and a boolean to check if the value has been set.
func (o *DatasetSetup) GetVisualizersOk() ([]VisualizerInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.Visualizers, true
}

// SetVisualizers sets field value
func (o *DatasetSetup) SetVisualizers(v []VisualizerInstance) {
	o.Visualizers = v
}

// GetPredictionTypes returns the PredictionTypes field value
func (o *DatasetSetup) GetPredictionTypes() []PredictionTypeInstance {
	if o == nil {
		var ret []PredictionTypeInstance
		return ret
	}

	return o.PredictionTypes
}

// GetPredictionTypesOk returns a tuple with the PredictionTypes field value
// and a boolean to check if the value has been set.
func (o *DatasetSetup) GetPredictionTypesOk() ([]PredictionTypeInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.PredictionTypes, true
}

// SetPredictionTypes sets field value
func (o *DatasetSetup) SetPredictionTypes(v []PredictionTypeInstance) {
	o.PredictionTypes = v
}

// GetCustomLosses returns the CustomLosses field value
func (o *DatasetSetup) GetCustomLosses() []CustomLossInstance {
	if o == nil {
		var ret []CustomLossInstance
		return ret
	}

	return o.CustomLosses
}

// GetCustomLossesOk returns a tuple with the CustomLosses field value
// and a boolean to check if the value has been set.
func (o *DatasetSetup) GetCustomLossesOk() ([]CustomLossInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomLosses, true
}

// SetCustomLosses sets field value
func (o *DatasetSetup) SetCustomLosses(v []CustomLossInstance) {
	o.CustomLosses = v
}

// GetMetrics returns the Metrics field value
func (o *DatasetSetup) GetMetrics() []MetricInstance {
	if o == nil {
		var ret []MetricInstance
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *DatasetSetup) GetMetricsOk() ([]MetricInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics, true
}

// SetMetrics sets field value
func (o *DatasetSetup) SetMetrics(v []MetricInstance) {
	o.Metrics = v
}

func (o DatasetSetup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetSetup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["preprocess"] = o.Preprocess
	toSerialize["inputs"] = o.Inputs
	toSerialize["metadata"] = o.Metadata
	toSerialize["outputs"] = o.Outputs
	toSerialize["visualizers"] = o.Visualizers
	toSerialize["prediction_types"] = o.PredictionTypes
	toSerialize["custom_losses"] = o.CustomLosses
	toSerialize["metrics"] = o.Metrics
	return toSerialize, nil
}

type NullableDatasetSetup struct {
	value *DatasetSetup
	isSet bool
}

func (v NullableDatasetSetup) Get() *DatasetSetup {
	return v.value
}

func (v *NullableDatasetSetup) Set(val *DatasetSetup) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetSetup) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetSetup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetSetup(val *DatasetSetup) *NullableDatasetSetup {
	return &NullableDatasetSetup{value: val, isSet: true}
}

func (v NullableDatasetSetup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetSetup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
