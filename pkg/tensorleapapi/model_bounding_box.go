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

// checks if the BoundingBox type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoundingBox{}

// BoundingBox struct for BoundingBox
type BoundingBox struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Width float64 `json:"width"`
	Height float64 `json:"height"`
	Confidence float64 `json:"confidence"`
	Label string `json:"label"`
}

// NewBoundingBox instantiates a new BoundingBox object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoundingBox(x float64, y float64, width float64, height float64, confidence float64, label string) *BoundingBox {
	this := BoundingBox{}
	this.X = x
	this.Y = y
	this.Width = width
	this.Height = height
	this.Confidence = confidence
	this.Label = label
	return &this
}

// NewBoundingBoxWithDefaults instantiates a new BoundingBox object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoundingBoxWithDefaults() *BoundingBox {
	this := BoundingBox{}
	return &this
}

// GetX returns the X field value
func (o *BoundingBox) GetX() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *BoundingBox) GetXOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *BoundingBox) SetX(v float64) {
	o.X = v
}

// GetY returns the Y field value
func (o *BoundingBox) GetY() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *BoundingBox) GetYOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *BoundingBox) SetY(v float64) {
	o.Y = v
}

// GetWidth returns the Width field value
func (o *BoundingBox) GetWidth() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *BoundingBox) GetWidthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *BoundingBox) SetWidth(v float64) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *BoundingBox) GetHeight() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *BoundingBox) GetHeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *BoundingBox) SetHeight(v float64) {
	o.Height = v
}

// GetConfidence returns the Confidence field value
func (o *BoundingBox) GetConfidence() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *BoundingBox) GetConfidenceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *BoundingBox) SetConfidence(v float64) {
	o.Confidence = v
}

// GetLabel returns the Label field value
func (o *BoundingBox) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *BoundingBox) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *BoundingBox) SetLabel(v string) {
	o.Label = v
}

func (o BoundingBox) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoundingBox) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	toSerialize["confidence"] = o.Confidence
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

type NullableBoundingBox struct {
	value *BoundingBox
	isSet bool
}

func (v NullableBoundingBox) Get() *BoundingBox {
	return v.value
}

func (v *NullableBoundingBox) Set(val *BoundingBox) {
	v.value = val
	v.isSet = true
}

func (v NullableBoundingBox) IsSet() bool {
	return v.isSet
}

func (v *NullableBoundingBox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoundingBox(val *BoundingBox) *NullableBoundingBox {
	return &NullableBoundingBox{value: val, isSet: true}
}

func (v NullableBoundingBox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoundingBox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


