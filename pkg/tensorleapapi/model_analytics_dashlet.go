/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.456
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the AnalyticsDashlet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsDashlet{}

// AnalyticsDashlet struct for AnalyticsDashlet
type AnalyticsDashlet struct {
	Cid    string               `json:"cid"`
	Layout SizedLayout          `json:"layout"`
	Type   string               `json:"type"`
	Data   AnalyticsDashletData `json:"data"`
}

// NewAnalyticsDashlet instantiates a new AnalyticsDashlet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsDashlet(cid string, layout SizedLayout, type_ string, data AnalyticsDashletData) *AnalyticsDashlet {
	this := AnalyticsDashlet{}
	this.Cid = cid
	this.Layout = layout
	this.Type = type_
	this.Data = data
	return &this
}

// NewAnalyticsDashletWithDefaults instantiates a new AnalyticsDashlet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsDashletWithDefaults() *AnalyticsDashlet {
	this := AnalyticsDashlet{}
	return &this
}

// GetCid returns the Cid field value
func (o *AnalyticsDashlet) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *AnalyticsDashlet) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *AnalyticsDashlet) SetCid(v string) {
	o.Cid = v
}

// GetLayout returns the Layout field value
func (o *AnalyticsDashlet) GetLayout() SizedLayout {
	if o == nil {
		var ret SizedLayout
		return ret
	}

	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value
// and a boolean to check if the value has been set.
func (o *AnalyticsDashlet) GetLayoutOk() (*SizedLayout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layout, true
}

// SetLayout sets field value
func (o *AnalyticsDashlet) SetLayout(v SizedLayout) {
	o.Layout = v
}

// GetType returns the Type field value
func (o *AnalyticsDashlet) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AnalyticsDashlet) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AnalyticsDashlet) SetType(v string) {
	o.Type = v
}

// GetData returns the Data field value
func (o *AnalyticsDashlet) GetData() AnalyticsDashletData {
	if o == nil {
		var ret AnalyticsDashletData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AnalyticsDashlet) GetDataOk() (*AnalyticsDashletData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AnalyticsDashlet) SetData(v AnalyticsDashletData) {
	o.Data = v
}

func (o AnalyticsDashlet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsDashlet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["layout"] = o.Layout
	toSerialize["type"] = o.Type
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAnalyticsDashlet struct {
	value *AnalyticsDashlet
	isSet bool
}

func (v NullableAnalyticsDashlet) Get() *AnalyticsDashlet {
	return v.value
}

func (v *NullableAnalyticsDashlet) Set(val *AnalyticsDashlet) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsDashlet) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsDashlet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsDashlet(val *AnalyticsDashlet) *NullableAnalyticsDashlet {
	return &NullableAnalyticsDashlet{value: val, isSet: true}
}

func (v NullableAnalyticsDashlet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsDashlet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
