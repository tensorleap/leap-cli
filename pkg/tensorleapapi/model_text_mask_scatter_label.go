/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.365
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the TextMaskScatterLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextMaskScatterLabel{}

// TextMaskScatterLabel struct for TextMaskScatterLabel
type TextMaskScatterLabel struct {
	Type               string      `json:"type"`
	TextVisualizations []TextViz   `json:"text_visualizations"`
	Mask               [][]float64 `json:"mask"`
	Labels             []string    `json:"labels"`
	InputName          *string     `json:"input_name,omitempty"`
}

// NewTextMaskScatterLabel instantiates a new TextMaskScatterLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextMaskScatterLabel(type_ string, textVisualizations []TextViz, mask [][]float64, labels []string) *TextMaskScatterLabel {
	this := TextMaskScatterLabel{}
	this.Type = type_
	this.TextVisualizations = textVisualizations
	this.Mask = mask
	this.Labels = labels
	return &this
}

// NewTextMaskScatterLabelWithDefaults instantiates a new TextMaskScatterLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextMaskScatterLabelWithDefaults() *TextMaskScatterLabel {
	this := TextMaskScatterLabel{}
	return &this
}

// GetType returns the Type field value
func (o *TextMaskScatterLabel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TextMaskScatterLabel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TextMaskScatterLabel) SetType(v string) {
	o.Type = v
}

// GetTextVisualizations returns the TextVisualizations field value
func (o *TextMaskScatterLabel) GetTextVisualizations() []TextViz {
	if o == nil {
		var ret []TextViz
		return ret
	}

	return o.TextVisualizations
}

// GetTextVisualizationsOk returns a tuple with the TextVisualizations field value
// and a boolean to check if the value has been set.
func (o *TextMaskScatterLabel) GetTextVisualizationsOk() ([]TextViz, bool) {
	if o == nil {
		return nil, false
	}
	return o.TextVisualizations, true
}

// SetTextVisualizations sets field value
func (o *TextMaskScatterLabel) SetTextVisualizations(v []TextViz) {
	o.TextVisualizations = v
}

// GetMask returns the Mask field value
func (o *TextMaskScatterLabel) GetMask() [][]float64 {
	if o == nil {
		var ret [][]float64
		return ret
	}

	return o.Mask
}

// GetMaskOk returns a tuple with the Mask field value
// and a boolean to check if the value has been set.
func (o *TextMaskScatterLabel) GetMaskOk() ([][]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mask, true
}

// SetMask sets field value
func (o *TextMaskScatterLabel) SetMask(v [][]float64) {
	o.Mask = v
}

// GetLabels returns the Labels field value
func (o *TextMaskScatterLabel) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *TextMaskScatterLabel) GetLabelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *TextMaskScatterLabel) SetLabels(v []string) {
	o.Labels = v
}

// GetInputName returns the InputName field value if set, zero value otherwise.
func (o *TextMaskScatterLabel) GetInputName() string {
	if o == nil || IsNil(o.InputName) {
		var ret string
		return ret
	}
	return *o.InputName
}

// GetInputNameOk returns a tuple with the InputName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextMaskScatterLabel) GetInputNameOk() (*string, bool) {
	if o == nil || IsNil(o.InputName) {
		return nil, false
	}
	return o.InputName, true
}

// HasInputName returns a boolean if a field has been set.
func (o *TextMaskScatterLabel) HasInputName() bool {
	if o != nil && !IsNil(o.InputName) {
		return true
	}

	return false
}

// SetInputName gets a reference to the given string and assigns it to the InputName field.
func (o *TextMaskScatterLabel) SetInputName(v string) {
	o.InputName = &v
}

func (o TextMaskScatterLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextMaskScatterLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["text_visualizations"] = o.TextVisualizations
	toSerialize["mask"] = o.Mask
	toSerialize["labels"] = o.Labels
	if !IsNil(o.InputName) {
		toSerialize["input_name"] = o.InputName
	}
	return toSerialize, nil
}

type NullableTextMaskScatterLabel struct {
	value *TextMaskScatterLabel
	isSet bool
}

func (v NullableTextMaskScatterLabel) Get() *TextMaskScatterLabel {
	return v.value
}

func (v *NullableTextMaskScatterLabel) Set(val *TextMaskScatterLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableTextMaskScatterLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableTextMaskScatterLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextMaskScatterLabel(val *TextMaskScatterLabel) *NullableTextMaskScatterLabel {
	return &NullableTextMaskScatterLabel{value: val, isSet: true}
}

func (v NullableTextMaskScatterLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextMaskScatterLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
