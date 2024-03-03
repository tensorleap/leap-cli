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

// checks if the MaskImageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaskImageData{}

// MaskImageData struct for MaskImageData
type MaskImageData struct {
	Blob     string       `json:"blob"`
	MaskBlob string       `json:"mask_blob"`
	Labels   []string     `json:"labels"`
	Type     DataTypeEnum `json:"type"`
}

// NewMaskImageData instantiates a new MaskImageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaskImageData(blob string, maskBlob string, labels []string, type_ DataTypeEnum) *MaskImageData {
	this := MaskImageData{}
	this.Blob = blob
	this.MaskBlob = maskBlob
	this.Labels = labels
	this.Type = type_
	return &this
}

// NewMaskImageDataWithDefaults instantiates a new MaskImageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaskImageDataWithDefaults() *MaskImageData {
	this := MaskImageData{}
	return &this
}

// GetBlob returns the Blob field value
func (o *MaskImageData) GetBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blob
}

// GetBlobOk returns a tuple with the Blob field value
// and a boolean to check if the value has been set.
func (o *MaskImageData) GetBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blob, true
}

// SetBlob sets field value
func (o *MaskImageData) SetBlob(v string) {
	o.Blob = v
}

// GetMaskBlob returns the MaskBlob field value
func (o *MaskImageData) GetMaskBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaskBlob
}

// GetMaskBlobOk returns a tuple with the MaskBlob field value
// and a boolean to check if the value has been set.
func (o *MaskImageData) GetMaskBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaskBlob, true
}

// SetMaskBlob sets field value
func (o *MaskImageData) SetMaskBlob(v string) {
	o.MaskBlob = v
}

// GetLabels returns the Labels field value
func (o *MaskImageData) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *MaskImageData) GetLabelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *MaskImageData) SetLabels(v []string) {
	o.Labels = v
}

// GetType returns the Type field value
func (o *MaskImageData) GetType() DataTypeEnum {
	if o == nil {
		var ret DataTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MaskImageData) GetTypeOk() (*DataTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MaskImageData) SetType(v DataTypeEnum) {
	o.Type = v
}

func (o MaskImageData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaskImageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blob"] = o.Blob
	toSerialize["mask_blob"] = o.MaskBlob
	toSerialize["labels"] = o.Labels
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableMaskImageData struct {
	value *MaskImageData
	isSet bool
}

func (v NullableMaskImageData) Get() *MaskImageData {
	return v.value
}

func (v *NullableMaskImageData) Set(val *MaskImageData) {
	v.value = val
	v.isSet = true
}

func (v NullableMaskImageData) IsSet() bool {
	return v.isSet
}

func (v *NullableMaskImageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaskImageData(val *MaskImageData) *NullableMaskImageData {
	return &NullableMaskImageData{value: val, isSet: true}
}

func (v NullableMaskImageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaskImageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
