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

// checks if the Heatmap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Heatmap{}

// Heatmap struct for Heatmap
type Heatmap struct {
	Body []float64 `json:"body"`
	Type HeatmapType `json:"type"`
}

// NewHeatmap instantiates a new Heatmap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeatmap(body []float64, type_ HeatmapType) *Heatmap {
	this := Heatmap{}
	this.Body = body
	this.Type = type_
	return &this
}

// NewHeatmapWithDefaults instantiates a new Heatmap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeatmapWithDefaults() *Heatmap {
	this := Heatmap{}
	return &this
}

// GetBody returns the Body field value
func (o *Heatmap) GetBody() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *Heatmap) GetBodyOk() ([]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body, true
}

// SetBody sets field value
func (o *Heatmap) SetBody(v []float64) {
	o.Body = v
}

// GetType returns the Type field value
func (o *Heatmap) GetType() HeatmapType {
	if o == nil {
		var ret HeatmapType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Heatmap) GetTypeOk() (*HeatmapType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Heatmap) SetType(v HeatmapType) {
	o.Type = v
}

func (o Heatmap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Heatmap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableHeatmap struct {
	value *Heatmap
	isSet bool
}

func (v NullableHeatmap) Get() *Heatmap {
	return v.value
}

func (v *NullableHeatmap) Set(val *Heatmap) {
	v.value = val
	v.isSet = true
}

func (v NullableHeatmap) IsSet() bool {
	return v.isSet
}

func (v *NullableHeatmap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeatmap(val *Heatmap) *NullableHeatmap {
	return &NullableHeatmap{value: val, isSet: true}
}

func (v NullableHeatmap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeatmap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


