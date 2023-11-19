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

// checks if the PopulationExplorationDashlet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PopulationExplorationDashlet{}

// PopulationExplorationDashlet struct for PopulationExplorationDashlet
type PopulationExplorationDashlet struct {
	Cid    string      `json:"cid"`
	Layout SizedLayout `json:"layout"`
	Type   string      `json:"type"`
	Name   string      `json:"name"`
}

// NewPopulationExplorationDashlet instantiates a new PopulationExplorationDashlet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopulationExplorationDashlet(cid string, layout SizedLayout, type_ string, name string) *PopulationExplorationDashlet {
	this := PopulationExplorationDashlet{}
	this.Cid = cid
	this.Layout = layout
	this.Type = type_
	this.Name = name
	return &this
}

// NewPopulationExplorationDashletWithDefaults instantiates a new PopulationExplorationDashlet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopulationExplorationDashletWithDefaults() *PopulationExplorationDashlet {
	this := PopulationExplorationDashlet{}
	return &this
}

// GetCid returns the Cid field value
func (o *PopulationExplorationDashlet) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *PopulationExplorationDashlet) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *PopulationExplorationDashlet) SetCid(v string) {
	o.Cid = v
}

// GetLayout returns the Layout field value
func (o *PopulationExplorationDashlet) GetLayout() SizedLayout {
	if o == nil {
		var ret SizedLayout
		return ret
	}

	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value
// and a boolean to check if the value has been set.
func (o *PopulationExplorationDashlet) GetLayoutOk() (*SizedLayout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layout, true
}

// SetLayout sets field value
func (o *PopulationExplorationDashlet) SetLayout(v SizedLayout) {
	o.Layout = v
}

// GetType returns the Type field value
func (o *PopulationExplorationDashlet) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PopulationExplorationDashlet) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PopulationExplorationDashlet) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *PopulationExplorationDashlet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PopulationExplorationDashlet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PopulationExplorationDashlet) SetName(v string) {
	o.Name = v
}

func (o PopulationExplorationDashlet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PopulationExplorationDashlet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["layout"] = o.Layout
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullablePopulationExplorationDashlet struct {
	value *PopulationExplorationDashlet
	isSet bool
}

func (v NullablePopulationExplorationDashlet) Get() *PopulationExplorationDashlet {
	return v.value
}

func (v *NullablePopulationExplorationDashlet) Set(val *PopulationExplorationDashlet) {
	v.value = val
	v.isSet = true
}

func (v NullablePopulationExplorationDashlet) IsSet() bool {
	return v.isSet
}

func (v *NullablePopulationExplorationDashlet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopulationExplorationDashlet(val *PopulationExplorationDashlet) *NullablePopulationExplorationDashlet {
	return &NullablePopulationExplorationDashlet{value: val, isSet: true}
}

func (v NullablePopulationExplorationDashlet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopulationExplorationDashlet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
