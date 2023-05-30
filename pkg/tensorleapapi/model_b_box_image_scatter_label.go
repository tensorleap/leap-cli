/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.321
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the BBoxImageScatterLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BBoxImageScatterLabel{}

// BBoxImageScatterLabel struct for BBoxImageScatterLabel
type BBoxImageScatterLabel struct {
	Data          [][]float64     `json:"data"`
	Blob          string          `json:"blob"`
	Type          string          `json:"type"`
	BoundingBoxes [][]BoundingBox `json:"bounding_boxes"`
	InputName     *string         `json:"input_name,omitempty"`
}

// NewBBoxImageScatterLabel instantiates a new BBoxImageScatterLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBBoxImageScatterLabel(data [][]float64, blob string, type_ string, boundingBoxes [][]BoundingBox) *BBoxImageScatterLabel {
	this := BBoxImageScatterLabel{}
	this.Data = data
	this.Blob = blob
	this.Type = type_
	this.BoundingBoxes = boundingBoxes
	return &this
}

// NewBBoxImageScatterLabelWithDefaults instantiates a new BBoxImageScatterLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBBoxImageScatterLabelWithDefaults() *BBoxImageScatterLabel {
	this := BBoxImageScatterLabel{}
	return &this
}

// GetData returns the Data field value
func (o *BBoxImageScatterLabel) GetData() [][]float64 {
	if o == nil {
		var ret [][]float64
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BBoxImageScatterLabel) GetDataOk() ([][]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BBoxImageScatterLabel) SetData(v [][]float64) {
	o.Data = v
}

// GetBlob returns the Blob field value
func (o *BBoxImageScatterLabel) GetBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blob
}

// GetBlobOk returns a tuple with the Blob field value
// and a boolean to check if the value has been set.
func (o *BBoxImageScatterLabel) GetBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blob, true
}

// SetBlob sets field value
func (o *BBoxImageScatterLabel) SetBlob(v string) {
	o.Blob = v
}

// GetType returns the Type field value
func (o *BBoxImageScatterLabel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BBoxImageScatterLabel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BBoxImageScatterLabel) SetType(v string) {
	o.Type = v
}

// GetBoundingBoxes returns the BoundingBoxes field value
func (o *BBoxImageScatterLabel) GetBoundingBoxes() [][]BoundingBox {
	if o == nil {
		var ret [][]BoundingBox
		return ret
	}

	return o.BoundingBoxes
}

// GetBoundingBoxesOk returns a tuple with the BoundingBoxes field value
// and a boolean to check if the value has been set.
func (o *BBoxImageScatterLabel) GetBoundingBoxesOk() ([][]BoundingBox, bool) {
	if o == nil {
		return nil, false
	}
	return o.BoundingBoxes, true
}

// SetBoundingBoxes sets field value
func (o *BBoxImageScatterLabel) SetBoundingBoxes(v [][]BoundingBox) {
	o.BoundingBoxes = v
}

// GetInputName returns the InputName field value if set, zero value otherwise.
func (o *BBoxImageScatterLabel) GetInputName() string {
	if o == nil || IsNil(o.InputName) {
		var ret string
		return ret
	}
	return *o.InputName
}

// GetInputNameOk returns a tuple with the InputName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BBoxImageScatterLabel) GetInputNameOk() (*string, bool) {
	if o == nil || IsNil(o.InputName) {
		return nil, false
	}
	return o.InputName, true
}

// HasInputName returns a boolean if a field has been set.
func (o *BBoxImageScatterLabel) HasInputName() bool {
	if o != nil && !IsNil(o.InputName) {
		return true
	}

	return false
}

// SetInputName gets a reference to the given string and assigns it to the InputName field.
func (o *BBoxImageScatterLabel) SetInputName(v string) {
	o.InputName = &v
}

func (o BBoxImageScatterLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BBoxImageScatterLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["blob"] = o.Blob
	toSerialize["type"] = o.Type
	toSerialize["bounding_boxes"] = o.BoundingBoxes
	if !IsNil(o.InputName) {
		toSerialize["input_name"] = o.InputName
	}
	return toSerialize, nil
}

type NullableBBoxImageScatterLabel struct {
	value *BBoxImageScatterLabel
	isSet bool
}

func (v NullableBBoxImageScatterLabel) Get() *BBoxImageScatterLabel {
	return v.value
}

func (v *NullableBBoxImageScatterLabel) Set(val *BBoxImageScatterLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableBBoxImageScatterLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableBBoxImageScatterLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBBoxImageScatterLabel(val *BBoxImageScatterLabel) *NullableBBoxImageScatterLabel {
	return &NullableBBoxImageScatterLabel{value: val, isSet: true}
}

func (v NullableBBoxImageScatterLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBBoxImageScatterLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
