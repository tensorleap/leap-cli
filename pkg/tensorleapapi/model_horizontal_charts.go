/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.422
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the HorizontalCharts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HorizontalCharts{}

// HorizontalCharts struct for HorizontalCharts
type HorizontalCharts struct {
	Label string              `json:"label"`
	Chart GenericDataResponse `json:"chart"`
}

// NewHorizontalCharts instantiates a new HorizontalCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHorizontalCharts(label string, chart GenericDataResponse) *HorizontalCharts {
	this := HorizontalCharts{}
	this.Label = label
	this.Chart = chart
	return &this
}

// NewHorizontalChartsWithDefaults instantiates a new HorizontalCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHorizontalChartsWithDefaults() *HorizontalCharts {
	this := HorizontalCharts{}
	return &this
}

// GetLabel returns the Label field value
func (o *HorizontalCharts) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *HorizontalCharts) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *HorizontalCharts) SetLabel(v string) {
	o.Label = v
}

// GetChart returns the Chart field value
func (o *HorizontalCharts) GetChart() GenericDataResponse {
	if o == nil {
		var ret GenericDataResponse
		return ret
	}

	return o.Chart
}

// GetChartOk returns a tuple with the Chart field value
// and a boolean to check if the value has been set.
func (o *HorizontalCharts) GetChartOk() (*GenericDataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chart, true
}

// SetChart sets field value
func (o *HorizontalCharts) SetChart(v GenericDataResponse) {
	o.Chart = v
}

func (o HorizontalCharts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HorizontalCharts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["chart"] = o.Chart
	return toSerialize, nil
}

type NullableHorizontalCharts struct {
	value *HorizontalCharts
	isSet bool
}

func (v NullableHorizontalCharts) Get() *HorizontalCharts {
	return v.value
}

func (v *NullableHorizontalCharts) Set(val *HorizontalCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableHorizontalCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableHorizontalCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHorizontalCharts(val *HorizontalCharts) *NullableHorizontalCharts {
	return &NullableHorizontalCharts{value: val, isSet: true}
}

func (v NullableHorizontalCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHorizontalCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
