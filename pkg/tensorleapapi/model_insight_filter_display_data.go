/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the InsightFilterDisplayData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InsightFilterDisplayData{}

// InsightFilterDisplayData struct for InsightFilterDisplayData
type InsightFilterDisplayData struct {
	Type        string           `json:"type"`
	InsightData InsightType      `json:"insightData"`
	SessionRun  FilterSessionRun `json:"sessionRun"`
}

// NewInsightFilterDisplayData instantiates a new InsightFilterDisplayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsightFilterDisplayData(type_ string, insightData InsightType, sessionRun FilterSessionRun) *InsightFilterDisplayData {
	this := InsightFilterDisplayData{}
	this.Type = type_
	this.InsightData = insightData
	this.SessionRun = sessionRun
	return &this
}

// NewInsightFilterDisplayDataWithDefaults instantiates a new InsightFilterDisplayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsightFilterDisplayDataWithDefaults() *InsightFilterDisplayData {
	this := InsightFilterDisplayData{}
	return &this
}

// GetType returns the Type field value
func (o *InsightFilterDisplayData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InsightFilterDisplayData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InsightFilterDisplayData) SetType(v string) {
	o.Type = v
}

// GetInsightData returns the InsightData field value
func (o *InsightFilterDisplayData) GetInsightData() InsightType {
	if o == nil {
		var ret InsightType
		return ret
	}

	return o.InsightData
}

// GetInsightDataOk returns a tuple with the InsightData field value
// and a boolean to check if the value has been set.
func (o *InsightFilterDisplayData) GetInsightDataOk() (*InsightType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InsightData, true
}

// SetInsightData sets field value
func (o *InsightFilterDisplayData) SetInsightData(v InsightType) {
	o.InsightData = v
}

// GetSessionRun returns the SessionRun field value
func (o *InsightFilterDisplayData) GetSessionRun() FilterSessionRun {
	if o == nil {
		var ret FilterSessionRun
		return ret
	}

	return o.SessionRun
}

// GetSessionRunOk returns a tuple with the SessionRun field value
// and a boolean to check if the value has been set.
func (o *InsightFilterDisplayData) GetSessionRunOk() (*FilterSessionRun, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRun, true
}

// SetSessionRun sets field value
func (o *InsightFilterDisplayData) SetSessionRun(v FilterSessionRun) {
	o.SessionRun = v
}

func (o InsightFilterDisplayData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InsightFilterDisplayData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["insightData"] = o.InsightData
	toSerialize["sessionRun"] = o.SessionRun
	return toSerialize, nil
}

type NullableInsightFilterDisplayData struct {
	value *InsightFilterDisplayData
	isSet bool
}

func (v NullableInsightFilterDisplayData) Get() *InsightFilterDisplayData {
	return v.value
}

func (v *NullableInsightFilterDisplayData) Set(val *InsightFilterDisplayData) {
	v.value = val
	v.isSet = true
}

func (v NullableInsightFilterDisplayData) IsSet() bool {
	return v.isSet
}

func (v *NullableInsightFilterDisplayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsightFilterDisplayData(val *InsightFilterDisplayData) *NullableInsightFilterDisplayData {
	return &NullableInsightFilterDisplayData{value: val, isSet: true}
}

func (v NullableInsightFilterDisplayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsightFilterDisplayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
