/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the Layout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Layout{}

// Layout struct for Layout
type Layout struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`
}

// NewLayout instantiates a new Layout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLayout(x float64, y float64, w float64, h float64) *Layout {
	this := Layout{}
	this.X = x
	this.Y = y
	this.W = w
	this.H = h
	return &this
}

// NewLayoutWithDefaults instantiates a new Layout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLayoutWithDefaults() *Layout {
	this := Layout{}
	return &this
}

// GetX returns the X field value
func (o *Layout) GetX() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *Layout) GetXOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *Layout) SetX(v float64) {
	o.X = v
}

// GetY returns the Y field value
func (o *Layout) GetY() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *Layout) GetYOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *Layout) SetY(v float64) {
	o.Y = v
}

// GetW returns the W field value
func (o *Layout) GetW() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.W
}

// GetWOk returns a tuple with the W field value
// and a boolean to check if the value has been set.
func (o *Layout) GetWOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.W, true
}

// SetW sets field value
func (o *Layout) SetW(v float64) {
	o.W = v
}

// GetH returns the H field value
func (o *Layout) GetH() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.H
}

// GetHOk returns a tuple with the H field value
// and a boolean to check if the value has been set.
func (o *Layout) GetHOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.H, true
}

// SetH sets field value
func (o *Layout) SetH(v float64) {
	o.H = v
}

func (o Layout) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Layout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	toSerialize["w"] = o.W
	toSerialize["h"] = o.H
	return toSerialize, nil
}

type NullableLayout struct {
	value *Layout
	isSet bool
}

func (v NullableLayout) Get() *Layout {
	return v.value
}

func (v *NullableLayout) Set(val *Layout) {
	v.value = val
	v.isSet = true
}

func (v NullableLayout) IsSet() bool {
	return v.isSet
}

func (v *NullableLayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayout(val *Layout) *NullableLayout {
	return &NullableLayout{value: val, isSet: true}
}

func (v NullableLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


