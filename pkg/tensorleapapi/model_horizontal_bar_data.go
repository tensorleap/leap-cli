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

// checks if the HorizontalBarData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HorizontalBarData{}

// HorizontalBarData struct for HorizontalBarData
type HorizontalBarData struct {
	Body   []float64    `json:"body"`
	Labels []string     `json:"labels"`
	Type   DataTypeEnum `json:"type"`
}

// NewHorizontalBarData instantiates a new HorizontalBarData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHorizontalBarData(body []float64, labels []string, type_ DataTypeEnum) *HorizontalBarData {
	this := HorizontalBarData{}
	this.Body = body
	this.Labels = labels
	this.Type = type_
	return &this
}

// NewHorizontalBarDataWithDefaults instantiates a new HorizontalBarData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHorizontalBarDataWithDefaults() *HorizontalBarData {
	this := HorizontalBarData{}
	return &this
}

// GetBody returns the Body field value
func (o *HorizontalBarData) GetBody() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *HorizontalBarData) GetBodyOk() ([]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body, true
}

// SetBody sets field value
func (o *HorizontalBarData) SetBody(v []float64) {
	o.Body = v
}

// GetLabels returns the Labels field value
func (o *HorizontalBarData) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *HorizontalBarData) GetLabelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *HorizontalBarData) SetLabels(v []string) {
	o.Labels = v
}

// GetType returns the Type field value
func (o *HorizontalBarData) GetType() DataTypeEnum {
	if o == nil {
		var ret DataTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HorizontalBarData) GetTypeOk() (*DataTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HorizontalBarData) SetType(v DataTypeEnum) {
	o.Type = v
}

func (o HorizontalBarData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HorizontalBarData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	toSerialize["labels"] = o.Labels
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableHorizontalBarData struct {
	value *HorizontalBarData
	isSet bool
}

func (v NullableHorizontalBarData) Get() *HorizontalBarData {
	return v.value
}

func (v *NullableHorizontalBarData) Set(val *HorizontalBarData) {
	v.value = val
	v.isSet = true
}

func (v NullableHorizontalBarData) IsSet() bool {
	return v.isSet
}

func (v *NullableHorizontalBarData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHorizontalBarData(val *HorizontalBarData) *NullableHorizontalBarData {
	return &NullableHorizontalBarData{value: val, isSet: true}
}

func (v NullableHorizontalBarData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHorizontalBarData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
