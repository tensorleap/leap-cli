/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.417
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the VisualizationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisualizationData{}

// VisualizationData struct for VisualizationData
type VisualizationData struct {
	Payload []VizType `json:"payload"`
	Title   string    `json:"title"`
	Epoch   float64   `json:"epoch"`
}

// NewVisualizationData instantiates a new VisualizationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisualizationData(payload []VizType, title string, epoch float64) *VisualizationData {
	this := VisualizationData{}
	this.Payload = payload
	this.Title = title
	this.Epoch = epoch
	return &this
}

// NewVisualizationDataWithDefaults instantiates a new VisualizationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisualizationDataWithDefaults() *VisualizationData {
	this := VisualizationData{}
	return &this
}

// GetPayload returns the Payload field value
func (o *VisualizationData) GetPayload() []VizType {
	if o == nil {
		var ret []VizType
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *VisualizationData) GetPayloadOk() ([]VizType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *VisualizationData) SetPayload(v []VizType) {
	o.Payload = v
}

// GetTitle returns the Title field value
func (o *VisualizationData) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *VisualizationData) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *VisualizationData) SetTitle(v string) {
	o.Title = v
}

// GetEpoch returns the Epoch field value
func (o *VisualizationData) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *VisualizationData) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *VisualizationData) SetEpoch(v float64) {
	o.Epoch = v
}

func (o VisualizationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisualizationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payload"] = o.Payload
	toSerialize["title"] = o.Title
	toSerialize["epoch"] = o.Epoch
	return toSerialize, nil
}

type NullableVisualizationData struct {
	value *VisualizationData
	isSet bool
}

func (v NullableVisualizationData) Get() *VisualizationData {
	return v.value
}

func (v *NullableVisualizationData) Set(val *VisualizationData) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualizationData) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualizationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualizationData(val *VisualizationData) *NullableVisualizationData {
	return &NullableVisualizationData{value: val, isSet: true}
}

func (v NullableVisualizationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualizationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
