/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GradsAnalysis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GradsAnalysis{}

// GradsAnalysis struct for GradsAnalysis
type GradsAnalysis struct {
	GradCam []GradsItem `json:"grad_cam"`
	GradByLoss []GradsItem `json:"grad_by_loss"`
}

// NewGradsAnalysis instantiates a new GradsAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGradsAnalysis(gradCam []GradsItem, gradByLoss []GradsItem) *GradsAnalysis {
	this := GradsAnalysis{}
	this.GradCam = gradCam
	this.GradByLoss = gradByLoss
	return &this
}

// NewGradsAnalysisWithDefaults instantiates a new GradsAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGradsAnalysisWithDefaults() *GradsAnalysis {
	this := GradsAnalysis{}
	return &this
}

// GetGradCam returns the GradCam field value
func (o *GradsAnalysis) GetGradCam() []GradsItem {
	if o == nil {
		var ret []GradsItem
		return ret
	}

	return o.GradCam
}

// GetGradCamOk returns a tuple with the GradCam field value
// and a boolean to check if the value has been set.
func (o *GradsAnalysis) GetGradCamOk() ([]GradsItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.GradCam, true
}

// SetGradCam sets field value
func (o *GradsAnalysis) SetGradCam(v []GradsItem) {
	o.GradCam = v
}

// GetGradByLoss returns the GradByLoss field value
func (o *GradsAnalysis) GetGradByLoss() []GradsItem {
	if o == nil {
		var ret []GradsItem
		return ret
	}

	return o.GradByLoss
}

// GetGradByLossOk returns a tuple with the GradByLoss field value
// and a boolean to check if the value has been set.
func (o *GradsAnalysis) GetGradByLossOk() ([]GradsItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.GradByLoss, true
}

// SetGradByLoss sets field value
func (o *GradsAnalysis) SetGradByLoss(v []GradsItem) {
	o.GradByLoss = v
}

func (o GradsAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GradsAnalysis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grad_cam"] = o.GradCam
	toSerialize["grad_by_loss"] = o.GradByLoss
	return toSerialize, nil
}

type NullableGradsAnalysis struct {
	value *GradsAnalysis
	isSet bool
}

func (v NullableGradsAnalysis) Get() *GradsAnalysis {
	return v.value
}

func (v *NullableGradsAnalysis) Set(val *GradsAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableGradsAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableGradsAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGradsAnalysis(val *GradsAnalysis) *NullableGradsAnalysis {
	return &NullableGradsAnalysis{value: val, isSet: true}
}

func (v NullableGradsAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGradsAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


