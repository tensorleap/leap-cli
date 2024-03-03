/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ImageScatterLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageScatterLabel{}

// ImageScatterLabel struct for ImageScatterLabel
type ImageScatterLabel struct {
	Data      [][]float64 `json:"data"`
	Blob      string      `json:"blob"`
	Type      string      `json:"type"`
	InputName *string     `json:"input_name,omitempty"`
}

// NewImageScatterLabel instantiates a new ImageScatterLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageScatterLabel(data [][]float64, blob string, type_ string) *ImageScatterLabel {
	this := ImageScatterLabel{}
	this.Data = data
	this.Blob = blob
	this.Type = type_
	return &this
}

// NewImageScatterLabelWithDefaults instantiates a new ImageScatterLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageScatterLabelWithDefaults() *ImageScatterLabel {
	this := ImageScatterLabel{}
	return &this
}

// GetData returns the Data field value
func (o *ImageScatterLabel) GetData() [][]float64 {
	if o == nil {
		var ret [][]float64
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ImageScatterLabel) GetDataOk() ([][]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ImageScatterLabel) SetData(v [][]float64) {
	o.Data = v
}

// GetBlob returns the Blob field value
func (o *ImageScatterLabel) GetBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blob
}

// GetBlobOk returns a tuple with the Blob field value
// and a boolean to check if the value has been set.
func (o *ImageScatterLabel) GetBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blob, true
}

// SetBlob sets field value
func (o *ImageScatterLabel) SetBlob(v string) {
	o.Blob = v
}

// GetType returns the Type field value
func (o *ImageScatterLabel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImageScatterLabel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ImageScatterLabel) SetType(v string) {
	o.Type = v
}

// GetInputName returns the InputName field value if set, zero value otherwise.
func (o *ImageScatterLabel) GetInputName() string {
	if o == nil || IsNil(o.InputName) {
		var ret string
		return ret
	}
	return *o.InputName
}

// GetInputNameOk returns a tuple with the InputName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageScatterLabel) GetInputNameOk() (*string, bool) {
	if o == nil || IsNil(o.InputName) {
		return nil, false
	}
	return o.InputName, true
}

// HasInputName returns a boolean if a field has been set.
func (o *ImageScatterLabel) HasInputName() bool {
	if o != nil && !IsNil(o.InputName) {
		return true
	}

	return false
}

// SetInputName gets a reference to the given string and assigns it to the InputName field.
func (o *ImageScatterLabel) SetInputName(v string) {
	o.InputName = &v
}

func (o ImageScatterLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageScatterLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["blob"] = o.Blob
	toSerialize["type"] = o.Type
	if !IsNil(o.InputName) {
		toSerialize["input_name"] = o.InputName
	}
	return toSerialize, nil
}

type NullableImageScatterLabel struct {
	value *ImageScatterLabel
	isSet bool
}

func (v NullableImageScatterLabel) Get() *ImageScatterLabel {
	return v.value
}

func (v *NullableImageScatterLabel) Set(val *ImageScatterLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableImageScatterLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableImageScatterLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageScatterLabel(val *ImageScatterLabel) *NullableImageScatterLabel {
	return &NullableImageScatterLabel{value: val, isSet: true}
}

func (v NullableImageScatterLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageScatterLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
