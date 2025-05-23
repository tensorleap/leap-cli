/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ConfusionMatrixLabelsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfusionMatrixLabelsResponse{}

// ConfusionMatrixLabelsResponse struct for ConfusionMatrixLabelsResponse
type ConfusionMatrixLabelsResponse struct {
	LabelsByPrediction []PredictionLabels `json:"labelsByPrediction"`
}

// NewConfusionMatrixLabelsResponse instantiates a new ConfusionMatrixLabelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfusionMatrixLabelsResponse(labelsByPrediction []PredictionLabels) *ConfusionMatrixLabelsResponse {
	this := ConfusionMatrixLabelsResponse{}
	this.LabelsByPrediction = labelsByPrediction
	return &this
}

// NewConfusionMatrixLabelsResponseWithDefaults instantiates a new ConfusionMatrixLabelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfusionMatrixLabelsResponseWithDefaults() *ConfusionMatrixLabelsResponse {
	this := ConfusionMatrixLabelsResponse{}
	return &this
}

// GetLabelsByPrediction returns the LabelsByPrediction field value
func (o *ConfusionMatrixLabelsResponse) GetLabelsByPrediction() []PredictionLabels {
	if o == nil {
		var ret []PredictionLabels
		return ret
	}

	return o.LabelsByPrediction
}

// GetLabelsByPredictionOk returns a tuple with the LabelsByPrediction field value
// and a boolean to check if the value has been set.
func (o *ConfusionMatrixLabelsResponse) GetLabelsByPredictionOk() ([]PredictionLabels, bool) {
	if o == nil {
		return nil, false
	}
	return o.LabelsByPrediction, true
}

// SetLabelsByPrediction sets field value
func (o *ConfusionMatrixLabelsResponse) SetLabelsByPrediction(v []PredictionLabels) {
	o.LabelsByPrediction = v
}

func (o ConfusionMatrixLabelsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfusionMatrixLabelsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labelsByPrediction"] = o.LabelsByPrediction
	return toSerialize, nil
}

type NullableConfusionMatrixLabelsResponse struct {
	value *ConfusionMatrixLabelsResponse
	isSet bool
}

func (v NullableConfusionMatrixLabelsResponse) Get() *ConfusionMatrixLabelsResponse {
	return v.value
}

func (v *NullableConfusionMatrixLabelsResponse) Set(val *ConfusionMatrixLabelsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfusionMatrixLabelsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfusionMatrixLabelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfusionMatrixLabelsResponse(val *ConfusionMatrixLabelsResponse) *NullableConfusionMatrixLabelsResponse {
	return &NullableConfusionMatrixLabelsResponse{value: val, isSet: true}
}

func (v NullableConfusionMatrixLabelsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfusionMatrixLabelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
