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

// checks if the AnalyticsDashletData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsDashletData{}

// AnalyticsDashletData struct for AnalyticsDashletData
type AnalyticsDashletData struct {
	// Construct a type with a set of properties K of type T
	Data map[string]interface{} `json:"data"`
	Name string                 `json:"name"`
	Type AnalyticsDashletType   `json:"type"`
}

// NewAnalyticsDashletData instantiates a new AnalyticsDashletData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsDashletData(data map[string]interface{}, name string, type_ AnalyticsDashletType) *AnalyticsDashletData {
	this := AnalyticsDashletData{}
	this.Data = data
	this.Name = name
	this.Type = type_
	return &this
}

// NewAnalyticsDashletDataWithDefaults instantiates a new AnalyticsDashletData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsDashletDataWithDefaults() *AnalyticsDashletData {
	this := AnalyticsDashletData{}
	return &this
}

// GetData returns the Data field value
func (o *AnalyticsDashletData) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AnalyticsDashletData) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AnalyticsDashletData) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetName returns the Name field value
func (o *AnalyticsDashletData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnalyticsDashletData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnalyticsDashletData) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *AnalyticsDashletData) GetType() AnalyticsDashletType {
	if o == nil {
		var ret AnalyticsDashletType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AnalyticsDashletData) GetTypeOk() (*AnalyticsDashletType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AnalyticsDashletData) SetType(v AnalyticsDashletType) {
	o.Type = v
}

func (o AnalyticsDashletData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsDashletData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAnalyticsDashletData struct {
	value *AnalyticsDashletData
	isSet bool
}

func (v NullableAnalyticsDashletData) Get() *AnalyticsDashletData {
	return v.value
}

func (v *NullableAnalyticsDashletData) Set(val *AnalyticsDashletData) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsDashletData) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsDashletData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsDashletData(val *AnalyticsDashletData) *NullableAnalyticsDashletData {
	return &NullableAnalyticsDashletData{value: val, isSet: true}
}

func (v NullableAnalyticsDashletData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsDashletData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
