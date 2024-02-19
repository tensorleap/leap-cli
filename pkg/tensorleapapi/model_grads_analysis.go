/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.464
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
	Data   []GradsItem  `json:"data"`
	IsLoss bool         `json:"is_loss"`
	Type   DataTypeEnum `json:"type"`
}

// NewGradsAnalysis instantiates a new GradsAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGradsAnalysis(data []GradsItem, isLoss bool, type_ DataTypeEnum) *GradsAnalysis {
	this := GradsAnalysis{}
	this.Data = data
	this.IsLoss = isLoss
	this.Type = type_
	return &this
}

// NewGradsAnalysisWithDefaults instantiates a new GradsAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGradsAnalysisWithDefaults() *GradsAnalysis {
	this := GradsAnalysis{}
	return &this
}

// GetData returns the Data field value
func (o *GradsAnalysis) GetData() []GradsItem {
	if o == nil {
		var ret []GradsItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GradsAnalysis) GetDataOk() ([]GradsItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GradsAnalysis) SetData(v []GradsItem) {
	o.Data = v
}

// GetIsLoss returns the IsLoss field value
func (o *GradsAnalysis) GetIsLoss() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLoss
}

// GetIsLossOk returns a tuple with the IsLoss field value
// and a boolean to check if the value has been set.
func (o *GradsAnalysis) GetIsLossOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLoss, true
}

// SetIsLoss sets field value
func (o *GradsAnalysis) SetIsLoss(v bool) {
	o.IsLoss = v
}

// GetType returns the Type field value
func (o *GradsAnalysis) GetType() DataTypeEnum {
	if o == nil {
		var ret DataTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GradsAnalysis) GetTypeOk() (*DataTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GradsAnalysis) SetType(v DataTypeEnum) {
	o.Type = v
}

func (o GradsAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GradsAnalysis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["is_loss"] = o.IsLoss
	toSerialize["type"] = o.Type
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
