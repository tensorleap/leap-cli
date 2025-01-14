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

// checks if the GraphData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphData{}

// GraphData struct for GraphData
type GraphData struct {
	Body    [][]float64  `json:"body"`
	Heatmap *Heatmap     `json:"heatmap,omitempty"`
	Type    DataTypeEnum `json:"type"`
}

// NewGraphData instantiates a new GraphData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphData(body [][]float64, type_ DataTypeEnum) *GraphData {
	this := GraphData{}
	this.Body = body
	this.Type = type_
	return &this
}

// NewGraphDataWithDefaults instantiates a new GraphData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphDataWithDefaults() *GraphData {
	this := GraphData{}
	return &this
}

// GetBody returns the Body field value
func (o *GraphData) GetBody() [][]float64 {
	if o == nil {
		var ret [][]float64
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *GraphData) GetBodyOk() ([][]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body, true
}

// SetBody sets field value
func (o *GraphData) SetBody(v [][]float64) {
	o.Body = v
}

// GetHeatmap returns the Heatmap field value if set, zero value otherwise.
func (o *GraphData) GetHeatmap() Heatmap {
	if o == nil || IsNil(o.Heatmap) {
		var ret Heatmap
		return ret
	}
	return *o.Heatmap
}

// GetHeatmapOk returns a tuple with the Heatmap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphData) GetHeatmapOk() (*Heatmap, bool) {
	if o == nil || IsNil(o.Heatmap) {
		return nil, false
	}
	return o.Heatmap, true
}

// HasHeatmap returns a boolean if a field has been set.
func (o *GraphData) HasHeatmap() bool {
	if o != nil && !IsNil(o.Heatmap) {
		return true
	}

	return false
}

// SetHeatmap gets a reference to the given Heatmap and assigns it to the Heatmap field.
func (o *GraphData) SetHeatmap(v Heatmap) {
	o.Heatmap = &v
}

// GetType returns the Type field value
func (o *GraphData) GetType() DataTypeEnum {
	if o == nil {
		var ret DataTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GraphData) GetTypeOk() (*DataTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GraphData) SetType(v DataTypeEnum) {
	o.Type = v
}

func (o GraphData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	if !IsNil(o.Heatmap) {
		toSerialize["heatmap"] = o.Heatmap
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableGraphData struct {
	value *GraphData
	isSet bool
}

func (v NullableGraphData) Get() *GraphData {
	return v.value
}

func (v *NullableGraphData) Set(val *GraphData) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphData) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphData(val *GraphData) *NullableGraphData {
	return &NullableGraphData{value: val, isSet: true}
}

func (v NullableGraphData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
