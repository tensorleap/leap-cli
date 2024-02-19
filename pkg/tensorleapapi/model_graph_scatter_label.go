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

// checks if the GraphScatterLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphScatterLabel{}

// GraphScatterLabel struct for GraphScatterLabel
type GraphScatterLabel struct {
	Type                string     `json:"type"`
	GraphVisualizations []GraphViz `json:"graph_visualizations"`
	InputName           *string    `json:"input_name,omitempty"`
}

// NewGraphScatterLabel instantiates a new GraphScatterLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphScatterLabel(type_ string, graphVisualizations []GraphViz) *GraphScatterLabel {
	this := GraphScatterLabel{}
	this.Type = type_
	this.GraphVisualizations = graphVisualizations
	return &this
}

// NewGraphScatterLabelWithDefaults instantiates a new GraphScatterLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphScatterLabelWithDefaults() *GraphScatterLabel {
	this := GraphScatterLabel{}
	return &this
}

// GetType returns the Type field value
func (o *GraphScatterLabel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GraphScatterLabel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GraphScatterLabel) SetType(v string) {
	o.Type = v
}

// GetGraphVisualizations returns the GraphVisualizations field value
func (o *GraphScatterLabel) GetGraphVisualizations() []GraphViz {
	if o == nil {
		var ret []GraphViz
		return ret
	}

	return o.GraphVisualizations
}

// GetGraphVisualizationsOk returns a tuple with the GraphVisualizations field value
// and a boolean to check if the value has been set.
func (o *GraphScatterLabel) GetGraphVisualizationsOk() ([]GraphViz, bool) {
	if o == nil {
		return nil, false
	}
	return o.GraphVisualizations, true
}

// SetGraphVisualizations sets field value
func (o *GraphScatterLabel) SetGraphVisualizations(v []GraphViz) {
	o.GraphVisualizations = v
}

// GetInputName returns the InputName field value if set, zero value otherwise.
func (o *GraphScatterLabel) GetInputName() string {
	if o == nil || IsNil(o.InputName) {
		var ret string
		return ret
	}
	return *o.InputName
}

// GetInputNameOk returns a tuple with the InputName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphScatterLabel) GetInputNameOk() (*string, bool) {
	if o == nil || IsNil(o.InputName) {
		return nil, false
	}
	return o.InputName, true
}

// HasInputName returns a boolean if a field has been set.
func (o *GraphScatterLabel) HasInputName() bool {
	if o != nil && !IsNil(o.InputName) {
		return true
	}

	return false
}

// SetInputName gets a reference to the given string and assigns it to the InputName field.
func (o *GraphScatterLabel) SetInputName(v string) {
	o.InputName = &v
}

func (o GraphScatterLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphScatterLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["graph_visualizations"] = o.GraphVisualizations
	if !IsNil(o.InputName) {
		toSerialize["input_name"] = o.InputName
	}
	return toSerialize, nil
}

type NullableGraphScatterLabel struct {
	value *GraphScatterLabel
	isSet bool
}

func (v NullableGraphScatterLabel) Get() *GraphScatterLabel {
	return v.value
}

func (v *NullableGraphScatterLabel) Set(val *GraphScatterLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphScatterLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphScatterLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphScatterLabel(val *GraphScatterLabel) *NullableGraphScatterLabel {
	return &NullableGraphScatterLabel{value: val, isSet: true}
}

func (v NullableGraphScatterLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphScatterLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
