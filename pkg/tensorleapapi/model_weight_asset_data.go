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

// checks if the WeightAssetData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WeightAssetData{}

// WeightAssetData struct for WeightAssetData
type WeightAssetData struct {
	SessionWeightId string `json:"sessionWeightId"`
	EsMetricIndex   string `json:"esMetricIndex"`
}

// NewWeightAssetData instantiates a new WeightAssetData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeightAssetData(sessionWeightId string, esMetricIndex string) *WeightAssetData {
	this := WeightAssetData{}
	this.SessionWeightId = sessionWeightId
	this.EsMetricIndex = esMetricIndex
	return &this
}

// NewWeightAssetDataWithDefaults instantiates a new WeightAssetData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeightAssetDataWithDefaults() *WeightAssetData {
	this := WeightAssetData{}
	return &this
}

// GetSessionWeightId returns the SessionWeightId field value
func (o *WeightAssetData) GetSessionWeightId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionWeightId
}

// GetSessionWeightIdOk returns a tuple with the SessionWeightId field value
// and a boolean to check if the value has been set.
func (o *WeightAssetData) GetSessionWeightIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionWeightId, true
}

// SetSessionWeightId sets field value
func (o *WeightAssetData) SetSessionWeightId(v string) {
	o.SessionWeightId = v
}

// GetEsMetricIndex returns the EsMetricIndex field value
func (o *WeightAssetData) GetEsMetricIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EsMetricIndex
}

// GetEsMetricIndexOk returns a tuple with the EsMetricIndex field value
// and a boolean to check if the value has been set.
func (o *WeightAssetData) GetEsMetricIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EsMetricIndex, true
}

// SetEsMetricIndex sets field value
func (o *WeightAssetData) SetEsMetricIndex(v string) {
	o.EsMetricIndex = v
}

func (o WeightAssetData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WeightAssetData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionWeightId"] = o.SessionWeightId
	toSerialize["esMetricIndex"] = o.EsMetricIndex
	return toSerialize, nil
}

type NullableWeightAssetData struct {
	value *WeightAssetData
	isSet bool
}

func (v NullableWeightAssetData) Get() *WeightAssetData {
	return v.value
}

func (v *NullableWeightAssetData) Set(val *WeightAssetData) {
	v.value = val
	v.isSet = true
}

func (v NullableWeightAssetData) IsSet() bool {
	return v.isSet
}

func (v *NullableWeightAssetData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeightAssetData(val *WeightAssetData) *NullableWeightAssetData {
	return &NullableWeightAssetData{value: val, isSet: true}
}

func (v NullableWeightAssetData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeightAssetData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
