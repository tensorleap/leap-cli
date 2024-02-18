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

// checks if the ModelGraph type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGraph{}

// ModelGraph struct for ModelGraph
type ModelGraph struct {
	Id string `json:"id"`
	// Construct a type with a set of properties K of type T
	Nodes map[string]interface{} `json:"nodes"`
}

// NewModelGraph instantiates a new ModelGraph object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGraph(id string, nodes map[string]interface{}) *ModelGraph {
	this := ModelGraph{}
	this.Id = id
	this.Nodes = nodes
	return &this
}

// NewModelGraphWithDefaults instantiates a new ModelGraph object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGraphWithDefaults() *ModelGraph {
	this := ModelGraph{}
	return &this
}

// GetId returns the Id field value
func (o *ModelGraph) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelGraph) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModelGraph) SetId(v string) {
	o.Id = v
}

// GetNodes returns the Nodes field value
func (o *ModelGraph) GetNodes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *ModelGraph) GetNodesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Nodes, true
}

// SetNodes sets field value
func (o *ModelGraph) SetNodes(v map[string]interface{}) {
	o.Nodes = v
}

func (o ModelGraph) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGraph) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["nodes"] = o.Nodes
	return toSerialize, nil
}

type NullableModelGraph struct {
	value *ModelGraph
	isSet bool
}

func (v NullableModelGraph) Get() *ModelGraph {
	return v.value
}

func (v *NullableModelGraph) Set(val *ModelGraph) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGraph) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGraph) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGraph(val *ModelGraph) *NullableModelGraph {
	return &NullableModelGraph{value: val, isSet: true}
}

func (v NullableModelGraph) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGraph) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
