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

// checks if the BBoxImageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BBoxImageData{}

// BBoxImageData struct for BBoxImageData
type BBoxImageData struct {
	Blob string `json:"blob"`
	BoundingBox []BoundingBox `json:"bounding_box"`
	Type DataTypeEnum `json:"type"`
}

// NewBBoxImageData instantiates a new BBoxImageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBBoxImageData(blob string, boundingBox []BoundingBox, type_ DataTypeEnum) *BBoxImageData {
	this := BBoxImageData{}
	this.Blob = blob
	this.BoundingBox = boundingBox
	this.Type = type_
	return &this
}

// NewBBoxImageDataWithDefaults instantiates a new BBoxImageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBBoxImageDataWithDefaults() *BBoxImageData {
	this := BBoxImageData{}
	return &this
}

// GetBlob returns the Blob field value
func (o *BBoxImageData) GetBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blob
}

// GetBlobOk returns a tuple with the Blob field value
// and a boolean to check if the value has been set.
func (o *BBoxImageData) GetBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blob, true
}

// SetBlob sets field value
func (o *BBoxImageData) SetBlob(v string) {
	o.Blob = v
}

// GetBoundingBox returns the BoundingBox field value
func (o *BBoxImageData) GetBoundingBox() []BoundingBox {
	if o == nil {
		var ret []BoundingBox
		return ret
	}

	return o.BoundingBox
}

// GetBoundingBoxOk returns a tuple with the BoundingBox field value
// and a boolean to check if the value has been set.
func (o *BBoxImageData) GetBoundingBoxOk() ([]BoundingBox, bool) {
	if o == nil {
		return nil, false
	}
	return o.BoundingBox, true
}

// SetBoundingBox sets field value
func (o *BBoxImageData) SetBoundingBox(v []BoundingBox) {
	o.BoundingBox = v
}

// GetType returns the Type field value
func (o *BBoxImageData) GetType() DataTypeEnum {
	if o == nil {
		var ret DataTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BBoxImageData) GetTypeOk() (*DataTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BBoxImageData) SetType(v DataTypeEnum) {
	o.Type = v
}

func (o BBoxImageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BBoxImageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blob"] = o.Blob
	toSerialize["bounding_box"] = o.BoundingBox
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableBBoxImageData struct {
	value *BBoxImageData
	isSet bool
}

func (v NullableBBoxImageData) Get() *BBoxImageData {
	return v.value
}

func (v *NullableBBoxImageData) Set(val *BBoxImageData) {
	v.value = val
	v.isSet = true
}

func (v NullableBBoxImageData) IsSet() bool {
	return v.isSet
}

func (v *NullableBBoxImageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBBoxImageData(val *BBoxImageData) *NullableBBoxImageData {
	return &NullableBBoxImageData{value: val, isSet: true}
}

func (v NullableBBoxImageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBBoxImageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


