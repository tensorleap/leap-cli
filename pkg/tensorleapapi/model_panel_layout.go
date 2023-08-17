/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the PanelLayout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PanelLayout{}

// PanelLayout struct for PanelLayout
type PanelLayout struct {
	VisualizationId string `json:"visualizationId"`
	Layout SizedLayout `json:"layout"`
}

// NewPanelLayout instantiates a new PanelLayout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPanelLayout(visualizationId string, layout SizedLayout) *PanelLayout {
	this := PanelLayout{}
	this.VisualizationId = visualizationId
	this.Layout = layout
	return &this
}

// NewPanelLayoutWithDefaults instantiates a new PanelLayout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPanelLayoutWithDefaults() *PanelLayout {
	this := PanelLayout{}
	return &this
}

// GetVisualizationId returns the VisualizationId field value
func (o *PanelLayout) GetVisualizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VisualizationId
}

// GetVisualizationIdOk returns a tuple with the VisualizationId field value
// and a boolean to check if the value has been set.
func (o *PanelLayout) GetVisualizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisualizationId, true
}

// SetVisualizationId sets field value
func (o *PanelLayout) SetVisualizationId(v string) {
	o.VisualizationId = v
}

// GetLayout returns the Layout field value
func (o *PanelLayout) GetLayout() SizedLayout {
	if o == nil {
		var ret SizedLayout
		return ret
	}

	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value
// and a boolean to check if the value has been set.
func (o *PanelLayout) GetLayoutOk() (*SizedLayout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layout, true
}

// SetLayout sets field value
func (o *PanelLayout) SetLayout(v SizedLayout) {
	o.Layout = v
}

func (o PanelLayout) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PanelLayout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["visualizationId"] = o.VisualizationId
	toSerialize["layout"] = o.Layout
	return toSerialize, nil
}

type NullablePanelLayout struct {
	value *PanelLayout
	isSet bool
}

func (v NullablePanelLayout) Get() *PanelLayout {
	return v.value
}

func (v *NullablePanelLayout) Set(val *PanelLayout) {
	v.value = val
	v.isSet = true
}

func (v NullablePanelLayout) IsSet() bool {
	return v.isSet
}

func (v *NullablePanelLayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePanelLayout(val *PanelLayout) *NullablePanelLayout {
	return &NullablePanelLayout{value: val, isSet: true}
}

func (v NullablePanelLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePanelLayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


