/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.423
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the TextScatterLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextScatterLabel{}

// TextScatterLabel struct for TextScatterLabel
type TextScatterLabel struct {
	Type string `json:"type"`
	TextVisualizations []TextViz `json:"text_visualizations"`
	InputName *string `json:"input_name,omitempty"`
}

// NewTextScatterLabel instantiates a new TextScatterLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextScatterLabel(type_ string, textVisualizations []TextViz) *TextScatterLabel {
	this := TextScatterLabel{}
	this.Type = type_
	this.TextVisualizations = textVisualizations
	return &this
}

// NewTextScatterLabelWithDefaults instantiates a new TextScatterLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextScatterLabelWithDefaults() *TextScatterLabel {
	this := TextScatterLabel{}
	return &this
}

// GetType returns the Type field value
func (o *TextScatterLabel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TextScatterLabel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TextScatterLabel) SetType(v string) {
	o.Type = v
}

// GetTextVisualizations returns the TextVisualizations field value
func (o *TextScatterLabel) GetTextVisualizations() []TextViz {
	if o == nil {
		var ret []TextViz
		return ret
	}

	return o.TextVisualizations
}

// GetTextVisualizationsOk returns a tuple with the TextVisualizations field value
// and a boolean to check if the value has been set.
func (o *TextScatterLabel) GetTextVisualizationsOk() ([]TextViz, bool) {
	if o == nil {
		return nil, false
	}
	return o.TextVisualizations, true
}

// SetTextVisualizations sets field value
func (o *TextScatterLabel) SetTextVisualizations(v []TextViz) {
	o.TextVisualizations = v
}

// GetInputName returns the InputName field value if set, zero value otherwise.
func (o *TextScatterLabel) GetInputName() string {
	if o == nil || IsNil(o.InputName) {
		var ret string
		return ret
	}
	return *o.InputName
}

// GetInputNameOk returns a tuple with the InputName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextScatterLabel) GetInputNameOk() (*string, bool) {
	if o == nil || IsNil(o.InputName) {
		return nil, false
	}
	return o.InputName, true
}

// HasInputName returns a boolean if a field has been set.
func (o *TextScatterLabel) HasInputName() bool {
	if o != nil && !IsNil(o.InputName) {
		return true
	}

	return false
}

// SetInputName gets a reference to the given string and assigns it to the InputName field.
func (o *TextScatterLabel) SetInputName(v string) {
	o.InputName = &v
}

func (o TextScatterLabel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextScatterLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["text_visualizations"] = o.TextVisualizations
	if !IsNil(o.InputName) {
		toSerialize["input_name"] = o.InputName
	}
	return toSerialize, nil
}

type NullableTextScatterLabel struct {
	value *TextScatterLabel
	isSet bool
}

func (v NullableTextScatterLabel) Get() *TextScatterLabel {
	return v.value
}

func (v *NullableTextScatterLabel) Set(val *TextScatterLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableTextScatterLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableTextScatterLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextScatterLabel(val *TextScatterLabel) *NullableTextScatterLabel {
	return &NullableTextScatterLabel{value: val, isSet: true}
}

func (v NullableTextScatterLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextScatterLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


