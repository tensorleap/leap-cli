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

// checks if the VerticalCharts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerticalCharts{}

// VerticalCharts struct for VerticalCharts
type VerticalCharts struct {
	Label      string             `json:"label"`
	Horizontal []HorizontalCharts `json:"horizontal"`
}

// NewVerticalCharts instantiates a new VerticalCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerticalCharts(label string, horizontal []HorizontalCharts) *VerticalCharts {
	this := VerticalCharts{}
	this.Label = label
	this.Horizontal = horizontal
	return &this
}

// NewVerticalChartsWithDefaults instantiates a new VerticalCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerticalChartsWithDefaults() *VerticalCharts {
	this := VerticalCharts{}
	return &this
}

// GetLabel returns the Label field value
func (o *VerticalCharts) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *VerticalCharts) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *VerticalCharts) SetLabel(v string) {
	o.Label = v
}

// GetHorizontal returns the Horizontal field value
func (o *VerticalCharts) GetHorizontal() []HorizontalCharts {
	if o == nil {
		var ret []HorizontalCharts
		return ret
	}

	return o.Horizontal
}

// GetHorizontalOk returns a tuple with the Horizontal field value
// and a boolean to check if the value has been set.
func (o *VerticalCharts) GetHorizontalOk() ([]HorizontalCharts, bool) {
	if o == nil {
		return nil, false
	}
	return o.Horizontal, true
}

// SetHorizontal sets field value
func (o *VerticalCharts) SetHorizontal(v []HorizontalCharts) {
	o.Horizontal = v
}

func (o VerticalCharts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerticalCharts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["horizontal"] = o.Horizontal
	return toSerialize, nil
}

type NullableVerticalCharts struct {
	value *VerticalCharts
	isSet bool
}

func (v NullableVerticalCharts) Get() *VerticalCharts {
	return v.value
}

func (v *NullableVerticalCharts) Set(val *VerticalCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableVerticalCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableVerticalCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerticalCharts(val *VerticalCharts) *NullableVerticalCharts {
	return &NullableVerticalCharts{value: val, isSet: true}
}

func (v NullableVerticalCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerticalCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
