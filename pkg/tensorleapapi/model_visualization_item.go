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

// checks if the VisualizationItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisualizationItem{}

// VisualizationItem struct for VisualizationItem
type VisualizationItem struct {
	Cid    string      `json:"cid"`
	Layout SizedLayout `json:"layout"`
	Type   string      `json:"type"`
}

// NewVisualizationItem instantiates a new VisualizationItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisualizationItem(cid string, layout SizedLayout, type_ string) *VisualizationItem {
	this := VisualizationItem{}
	this.Cid = cid
	this.Layout = layout
	this.Type = type_
	return &this
}

// NewVisualizationItemWithDefaults instantiates a new VisualizationItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisualizationItemWithDefaults() *VisualizationItem {
	this := VisualizationItem{}
	return &this
}

// GetCid returns the Cid field value
func (o *VisualizationItem) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *VisualizationItem) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *VisualizationItem) SetCid(v string) {
	o.Cid = v
}

// GetLayout returns the Layout field value
func (o *VisualizationItem) GetLayout() SizedLayout {
	if o == nil {
		var ret SizedLayout
		return ret
	}

	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value
// and a boolean to check if the value has been set.
func (o *VisualizationItem) GetLayoutOk() (*SizedLayout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layout, true
}

// SetLayout sets field value
func (o *VisualizationItem) SetLayout(v SizedLayout) {
	o.Layout = v
}

// GetType returns the Type field value
func (o *VisualizationItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VisualizationItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VisualizationItem) SetType(v string) {
	o.Type = v
}

func (o VisualizationItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisualizationItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["layout"] = o.Layout
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableVisualizationItem struct {
	value *VisualizationItem
	isSet bool
}

func (v NullableVisualizationItem) Get() *VisualizationItem {
	return v.value
}

func (v *NullableVisualizationItem) Set(val *VisualizationItem) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualizationItem) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualizationItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualizationItem(val *VisualizationItem) *NullableVisualizationItem {
	return &NullableVisualizationItem{value: val, isSet: true}
}

func (v NullableVisualizationItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualizationItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
