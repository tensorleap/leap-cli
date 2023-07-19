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

// checks if the MultiChartsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiChartsResponse{}

// MultiChartsResponse struct for MultiChartsResponse
type MultiChartsResponse struct {
	Vertical []VerticalCharts `json:"vertical"`
}

// NewMultiChartsResponse instantiates a new MultiChartsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiChartsResponse(vertical []VerticalCharts) *MultiChartsResponse {
	this := MultiChartsResponse{}
	this.Vertical = vertical
	return &this
}

// NewMultiChartsResponseWithDefaults instantiates a new MultiChartsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiChartsResponseWithDefaults() *MultiChartsResponse {
	this := MultiChartsResponse{}
	return &this
}

// GetVertical returns the Vertical field value
func (o *MultiChartsResponse) GetVertical() []VerticalCharts {
	if o == nil {
		var ret []VerticalCharts
		return ret
	}

	return o.Vertical
}

// GetVerticalOk returns a tuple with the Vertical field value
// and a boolean to check if the value has been set.
func (o *MultiChartsResponse) GetVerticalOk() ([]VerticalCharts, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vertical, true
}

// SetVertical sets field value
func (o *MultiChartsResponse) SetVertical(v []VerticalCharts) {
	o.Vertical = v
}

func (o MultiChartsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiChartsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vertical"] = o.Vertical
	return toSerialize, nil
}

type NullableMultiChartsResponse struct {
	value *MultiChartsResponse
	isSet bool
}

func (v NullableMultiChartsResponse) Get() *MultiChartsResponse {
	return v.value
}

func (v *NullableMultiChartsResponse) Set(val *MultiChartsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiChartsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiChartsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiChartsResponse(val *MultiChartsResponse) *NullableMultiChartsResponse {
	return &NullableMultiChartsResponse{value: val, isSet: true}
}

func (v NullableMultiChartsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiChartsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


