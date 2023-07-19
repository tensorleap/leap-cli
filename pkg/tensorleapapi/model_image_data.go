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

// checks if the ImageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageData{}

// ImageData struct for ImageData
type ImageData struct {
	Blob string `json:"blob"`
	Type DataTypeEnum `json:"type"`
}

// NewImageData instantiates a new ImageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageData(blob string, type_ DataTypeEnum) *ImageData {
	this := ImageData{}
	this.Blob = blob
	this.Type = type_
	return &this
}

// NewImageDataWithDefaults instantiates a new ImageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageDataWithDefaults() *ImageData {
	this := ImageData{}
	return &this
}

// GetBlob returns the Blob field value
func (o *ImageData) GetBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blob
}

// GetBlobOk returns a tuple with the Blob field value
// and a boolean to check if the value has been set.
func (o *ImageData) GetBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blob, true
}

// SetBlob sets field value
func (o *ImageData) SetBlob(v string) {
	o.Blob = v
}

// GetType returns the Type field value
func (o *ImageData) GetType() DataTypeEnum {
	if o == nil {
		var ret DataTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImageData) GetTypeOk() (*DataTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ImageData) SetType(v DataTypeEnum) {
	o.Type = v
}

func (o ImageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blob"] = o.Blob
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableImageData struct {
	value *ImageData
	isSet bool
}

func (v NullableImageData) Get() *ImageData {
	return v.value
}

func (v *NullableImageData) Set(val *ImageData) {
	v.value = val
	v.isSet = true
}

func (v NullableImageData) IsSet() bool {
	return v.isSet
}

func (v *NullableImageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageData(val *ImageData) *NullableImageData {
	return &NullableImageData{value: val, isSet: true}
}

func (v NullableImageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


