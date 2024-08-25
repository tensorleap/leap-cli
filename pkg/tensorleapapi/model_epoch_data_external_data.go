/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.526
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the EpochDataExternalData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EpochDataExternalData{}

// EpochDataExternalData struct for EpochDataExternalData
type EpochDataExternalData struct {
	Metrics map[string]EpochMetricsValue `json:"metrics"`
}

// NewEpochDataExternalData instantiates a new EpochDataExternalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpochDataExternalData(metrics map[string]EpochMetricsValue) *EpochDataExternalData {
	this := EpochDataExternalData{}
	this.Metrics = metrics
	return &this
}

// NewEpochDataExternalDataWithDefaults instantiates a new EpochDataExternalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpochDataExternalDataWithDefaults() *EpochDataExternalData {
	this := EpochDataExternalData{}
	return &this
}

// GetMetrics returns the Metrics field value
func (o *EpochDataExternalData) GetMetrics() map[string]EpochMetricsValue {
	if o == nil {
		var ret map[string]EpochMetricsValue
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *EpochDataExternalData) GetMetricsOk() (*map[string]EpochMetricsValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *EpochDataExternalData) SetMetrics(v map[string]EpochMetricsValue) {
	o.Metrics = v
}

func (o EpochDataExternalData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EpochDataExternalData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metrics"] = o.Metrics
	return toSerialize, nil
}

type NullableEpochDataExternalData struct {
	value *EpochDataExternalData
	isSet bool
}

func (v NullableEpochDataExternalData) Get() *EpochDataExternalData {
	return v.value
}

func (v *NullableEpochDataExternalData) Set(val *EpochDataExternalData) {
	v.value = val
	v.isSet = true
}

func (v NullableEpochDataExternalData) IsSet() bool {
	return v.isSet
}

func (v *NullableEpochDataExternalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpochDataExternalData(val *EpochDataExternalData) *NullableEpochDataExternalData {
	return &NullableEpochDataExternalData{value: val, isSet: true}
}

func (v NullableEpochDataExternalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpochDataExternalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
