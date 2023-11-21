/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.404
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SampleAssetId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SampleAssetId{}

// SampleAssetId struct for SampleAssetId
type SampleAssetId struct {
	Name string           `json:"name"`
	Type SampleAssetNames `json:"type"`
}

// NewSampleAssetId instantiates a new SampleAssetId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSampleAssetId(name string, type_ SampleAssetNames) *SampleAssetId {
	this := SampleAssetId{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewSampleAssetIdWithDefaults instantiates a new SampleAssetId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSampleAssetIdWithDefaults() *SampleAssetId {
	this := SampleAssetId{}
	return &this
}

// GetName returns the Name field value
func (o *SampleAssetId) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SampleAssetId) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SampleAssetId) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *SampleAssetId) GetType() SampleAssetNames {
	if o == nil {
		var ret SampleAssetNames
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SampleAssetId) GetTypeOk() (*SampleAssetNames, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SampleAssetId) SetType(v SampleAssetNames) {
	o.Type = v
}

func (o SampleAssetId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SampleAssetId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableSampleAssetId struct {
	value *SampleAssetId
	isSet bool
}

func (v NullableSampleAssetId) Get() *SampleAssetId {
	return v.value
}

func (v *NullableSampleAssetId) Set(val *SampleAssetId) {
	v.value = val
	v.isSet = true
}

func (v NullableSampleAssetId) IsSet() bool {
	return v.isSet
}

func (v *NullableSampleAssetId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSampleAssetId(val *SampleAssetId) *NullableSampleAssetId {
	return &NullableSampleAssetId{value: val, isSet: true}
}

func (v NullableSampleAssetId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSampleAssetId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}