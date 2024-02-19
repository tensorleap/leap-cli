/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.464
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ConfusionMetricNamesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfusionMetricNamesResponse{}

// ConfusionMetricNamesResponse struct for ConfusionMetricNamesResponse
type ConfusionMetricNamesResponse struct {
	ConfusionMetricNames []string `json:"confusionMetricNames"`
}

// NewConfusionMetricNamesResponse instantiates a new ConfusionMetricNamesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfusionMetricNamesResponse(confusionMetricNames []string) *ConfusionMetricNamesResponse {
	this := ConfusionMetricNamesResponse{}
	this.ConfusionMetricNames = confusionMetricNames
	return &this
}

// NewConfusionMetricNamesResponseWithDefaults instantiates a new ConfusionMetricNamesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfusionMetricNamesResponseWithDefaults() *ConfusionMetricNamesResponse {
	this := ConfusionMetricNamesResponse{}
	return &this
}

// GetConfusionMetricNames returns the ConfusionMetricNames field value
func (o *ConfusionMetricNamesResponse) GetConfusionMetricNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConfusionMetricNames
}

// GetConfusionMetricNamesOk returns a tuple with the ConfusionMetricNames field value
// and a boolean to check if the value has been set.
func (o *ConfusionMetricNamesResponse) GetConfusionMetricNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfusionMetricNames, true
}

// SetConfusionMetricNames sets field value
func (o *ConfusionMetricNamesResponse) SetConfusionMetricNames(v []string) {
	o.ConfusionMetricNames = v
}

func (o ConfusionMetricNamesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfusionMetricNamesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["confusionMetricNames"] = o.ConfusionMetricNames
	return toSerialize, nil
}

type NullableConfusionMetricNamesResponse struct {
	value *ConfusionMetricNamesResponse
	isSet bool
}

func (v NullableConfusionMetricNamesResponse) Get() *ConfusionMetricNamesResponse {
	return v.value
}

func (v *NullableConfusionMetricNamesResponse) Set(val *ConfusionMetricNamesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfusionMetricNamesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfusionMetricNamesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfusionMetricNamesResponse(val *ConfusionMetricNamesResponse) *NullableConfusionMetricNamesResponse {
	return &NullableConfusionMetricNamesResponse{value: val, isSet: true}
}

func (v NullableConfusionMetricNamesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfusionMetricNamesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
