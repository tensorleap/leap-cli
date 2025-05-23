/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the TextData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextData{}

// TextData struct for TextData
type TextData struct {
	Body    []string     `json:"body"`
	Heatmap *Heatmap     `json:"heatmap,omitempty"`
	Type    DataTypeEnum `json:"type"`
}

// NewTextData instantiates a new TextData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextData(body []string, type_ DataTypeEnum) *TextData {
	this := TextData{}
	this.Body = body
	this.Type = type_
	return &this
}

// NewTextDataWithDefaults instantiates a new TextData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextDataWithDefaults() *TextData {
	this := TextData{}
	return &this
}

// GetBody returns the Body field value
func (o *TextData) GetBody() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *TextData) GetBodyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body, true
}

// SetBody sets field value
func (o *TextData) SetBody(v []string) {
	o.Body = v
}

// GetHeatmap returns the Heatmap field value if set, zero value otherwise.
func (o *TextData) GetHeatmap() Heatmap {
	if o == nil || IsNil(o.Heatmap) {
		var ret Heatmap
		return ret
	}
	return *o.Heatmap
}

// GetHeatmapOk returns a tuple with the Heatmap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextData) GetHeatmapOk() (*Heatmap, bool) {
	if o == nil || IsNil(o.Heatmap) {
		return nil, false
	}
	return o.Heatmap, true
}

// HasHeatmap returns a boolean if a field has been set.
func (o *TextData) HasHeatmap() bool {
	if o != nil && !IsNil(o.Heatmap) {
		return true
	}

	return false
}

// SetHeatmap gets a reference to the given Heatmap and assigns it to the Heatmap field.
func (o *TextData) SetHeatmap(v Heatmap) {
	o.Heatmap = &v
}

// GetType returns the Type field value
func (o *TextData) GetType() DataTypeEnum {
	if o == nil {
		var ret DataTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TextData) GetTypeOk() (*DataTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TextData) SetType(v DataTypeEnum) {
	o.Type = v
}

func (o TextData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	if !IsNil(o.Heatmap) {
		toSerialize["heatmap"] = o.Heatmap
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTextData struct {
	value *TextData
	isSet bool
}

func (v NullableTextData) Get() *TextData {
	return v.value
}

func (v *NullableTextData) Set(val *TextData) {
	v.value = val
	v.isSet = true
}

func (v NullableTextData) IsSet() bool {
	return v.isSet
}

func (v *NullableTextData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextData(val *TextData) *NullableTextData {
	return &NullableTextData{value: val, isSet: true}
}

func (v NullableTextData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
