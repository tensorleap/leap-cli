/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the MaskTextData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaskTextData{}

// MaskTextData struct for MaskTextData
type MaskTextData struct {
	Text   []string     `json:"text"`
	Mask   []float64    `json:"mask"`
	Labels []string     `json:"labels"`
	Type   DataTypeEnum `json:"type"`
}

// NewMaskTextData instantiates a new MaskTextData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaskTextData(text []string, mask []float64, labels []string, type_ DataTypeEnum) *MaskTextData {
	this := MaskTextData{}
	this.Text = text
	this.Mask = mask
	this.Labels = labels
	this.Type = type_
	return &this
}

// NewMaskTextDataWithDefaults instantiates a new MaskTextData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaskTextDataWithDefaults() *MaskTextData {
	this := MaskTextData{}
	return &this
}

// GetText returns the Text field value
func (o *MaskTextData) GetText() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *MaskTextData) GetTextOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text, true
}

// SetText sets field value
func (o *MaskTextData) SetText(v []string) {
	o.Text = v
}

// GetMask returns the Mask field value
func (o *MaskTextData) GetMask() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}

	return o.Mask
}

// GetMaskOk returns a tuple with the Mask field value
// and a boolean to check if the value has been set.
func (o *MaskTextData) GetMaskOk() ([]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mask, true
}

// SetMask sets field value
func (o *MaskTextData) SetMask(v []float64) {
	o.Mask = v
}

// GetLabels returns the Labels field value
func (o *MaskTextData) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *MaskTextData) GetLabelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *MaskTextData) SetLabels(v []string) {
	o.Labels = v
}

// GetType returns the Type field value
func (o *MaskTextData) GetType() DataTypeEnum {
	if o == nil {
		var ret DataTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MaskTextData) GetTypeOk() (*DataTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MaskTextData) SetType(v DataTypeEnum) {
	o.Type = v
}

func (o MaskTextData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaskTextData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	toSerialize["mask"] = o.Mask
	toSerialize["labels"] = o.Labels
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableMaskTextData struct {
	value *MaskTextData
	isSet bool
}

func (v NullableMaskTextData) Get() *MaskTextData {
	return v.value
}

func (v *NullableMaskTextData) Set(val *MaskTextData) {
	v.value = val
	v.isSet = true
}

func (v NullableMaskTextData) IsSet() bool {
	return v.isSet
}

func (v *NullableMaskTextData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaskTextData(val *MaskTextData) *NullableMaskTextData {
	return &NullableMaskTextData{value: val, isSet: true}
}

func (v NullableMaskTextData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaskTextData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
