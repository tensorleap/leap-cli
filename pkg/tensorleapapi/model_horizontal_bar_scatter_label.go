/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the HorizontalBarScatterLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HorizontalBarScatterLabel{}

// HorizontalBarScatterLabel struct for HorizontalBarScatterLabel
type HorizontalBarScatterLabel struct {
	Type                        string             `json:"type"`
	HorizontalBarVisualizations []HorizontalBarViz `json:"horizontal_bar_visualizations"`
	InputName                   *string            `json:"input_name,omitempty"`
}

// NewHorizontalBarScatterLabel instantiates a new HorizontalBarScatterLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHorizontalBarScatterLabel(type_ string, horizontalBarVisualizations []HorizontalBarViz) *HorizontalBarScatterLabel {
	this := HorizontalBarScatterLabel{}
	this.Type = type_
	this.HorizontalBarVisualizations = horizontalBarVisualizations
	return &this
}

// NewHorizontalBarScatterLabelWithDefaults instantiates a new HorizontalBarScatterLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHorizontalBarScatterLabelWithDefaults() *HorizontalBarScatterLabel {
	this := HorizontalBarScatterLabel{}
	return &this
}

// GetType returns the Type field value
func (o *HorizontalBarScatterLabel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HorizontalBarScatterLabel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HorizontalBarScatterLabel) SetType(v string) {
	o.Type = v
}

// GetHorizontalBarVisualizations returns the HorizontalBarVisualizations field value
func (o *HorizontalBarScatterLabel) GetHorizontalBarVisualizations() []HorizontalBarViz {
	if o == nil {
		var ret []HorizontalBarViz
		return ret
	}

	return o.HorizontalBarVisualizations
}

// GetHorizontalBarVisualizationsOk returns a tuple with the HorizontalBarVisualizations field value
// and a boolean to check if the value has been set.
func (o *HorizontalBarScatterLabel) GetHorizontalBarVisualizationsOk() ([]HorizontalBarViz, bool) {
	if o == nil {
		return nil, false
	}
	return o.HorizontalBarVisualizations, true
}

// SetHorizontalBarVisualizations sets field value
func (o *HorizontalBarScatterLabel) SetHorizontalBarVisualizations(v []HorizontalBarViz) {
	o.HorizontalBarVisualizations = v
}

// GetInputName returns the InputName field value if set, zero value otherwise.
func (o *HorizontalBarScatterLabel) GetInputName() string {
	if o == nil || IsNil(o.InputName) {
		var ret string
		return ret
	}
	return *o.InputName
}

// GetInputNameOk returns a tuple with the InputName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizontalBarScatterLabel) GetInputNameOk() (*string, bool) {
	if o == nil || IsNil(o.InputName) {
		return nil, false
	}
	return o.InputName, true
}

// HasInputName returns a boolean if a field has been set.
func (o *HorizontalBarScatterLabel) HasInputName() bool {
	if o != nil && !IsNil(o.InputName) {
		return true
	}

	return false
}

// SetInputName gets a reference to the given string and assigns it to the InputName field.
func (o *HorizontalBarScatterLabel) SetInputName(v string) {
	o.InputName = &v
}

func (o HorizontalBarScatterLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HorizontalBarScatterLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["horizontal_bar_visualizations"] = o.HorizontalBarVisualizations
	if !IsNil(o.InputName) {
		toSerialize["input_name"] = o.InputName
	}
	return toSerialize, nil
}

type NullableHorizontalBarScatterLabel struct {
	value *HorizontalBarScatterLabel
	isSet bool
}

func (v NullableHorizontalBarScatterLabel) Get() *HorizontalBarScatterLabel {
	return v.value
}

func (v *NullableHorizontalBarScatterLabel) Set(val *HorizontalBarScatterLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableHorizontalBarScatterLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableHorizontalBarScatterLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHorizontalBarScatterLabel(val *HorizontalBarScatterLabel) *NullableHorizontalBarScatterLabel {
	return &NullableHorizontalBarScatterLabel{value: val, isSet: true}
}

func (v NullableHorizontalBarScatterLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHorizontalBarScatterLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
