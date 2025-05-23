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

// checks if the SizedLayout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SizedLayout{}

// SizedLayout struct for SizedLayout
type SizedLayout struct {
	Xs Layout `json:"xs"`
	Sm Layout `json:"sm"`
	Md Layout `json:"md"`
	Lg Layout `json:"lg"`
}

// NewSizedLayout instantiates a new SizedLayout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSizedLayout(xs Layout, sm Layout, md Layout, lg Layout) *SizedLayout {
	this := SizedLayout{}
	this.Xs = xs
	this.Sm = sm
	this.Md = md
	this.Lg = lg
	return &this
}

// NewSizedLayoutWithDefaults instantiates a new SizedLayout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSizedLayoutWithDefaults() *SizedLayout {
	this := SizedLayout{}
	return &this
}

// GetXs returns the Xs field value
func (o *SizedLayout) GetXs() Layout {
	if o == nil {
		var ret Layout
		return ret
	}

	return o.Xs
}

// GetXsOk returns a tuple with the Xs field value
// and a boolean to check if the value has been set.
func (o *SizedLayout) GetXsOk() (*Layout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Xs, true
}

// SetXs sets field value
func (o *SizedLayout) SetXs(v Layout) {
	o.Xs = v
}

// GetSm returns the Sm field value
func (o *SizedLayout) GetSm() Layout {
	if o == nil {
		var ret Layout
		return ret
	}

	return o.Sm
}

// GetSmOk returns a tuple with the Sm field value
// and a boolean to check if the value has been set.
func (o *SizedLayout) GetSmOk() (*Layout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sm, true
}

// SetSm sets field value
func (o *SizedLayout) SetSm(v Layout) {
	o.Sm = v
}

// GetMd returns the Md field value
func (o *SizedLayout) GetMd() Layout {
	if o == nil {
		var ret Layout
		return ret
	}

	return o.Md
}

// GetMdOk returns a tuple with the Md field value
// and a boolean to check if the value has been set.
func (o *SizedLayout) GetMdOk() (*Layout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Md, true
}

// SetMd sets field value
func (o *SizedLayout) SetMd(v Layout) {
	o.Md = v
}

// GetLg returns the Lg field value
func (o *SizedLayout) GetLg() Layout {
	if o == nil {
		var ret Layout
		return ret
	}

	return o.Lg
}

// GetLgOk returns a tuple with the Lg field value
// and a boolean to check if the value has been set.
func (o *SizedLayout) GetLgOk() (*Layout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lg, true
}

// SetLg sets field value
func (o *SizedLayout) SetLg(v Layout) {
	o.Lg = v
}

func (o SizedLayout) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SizedLayout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["xs"] = o.Xs
	toSerialize["sm"] = o.Sm
	toSerialize["md"] = o.Md
	toSerialize["lg"] = o.Lg
	return toSerialize, nil
}

type NullableSizedLayout struct {
	value *SizedLayout
	isSet bool
}

func (v NullableSizedLayout) Get() *SizedLayout {
	return v.value
}

func (v *NullableSizedLayout) Set(val *SizedLayout) {
	v.value = val
	v.isSet = true
}

func (v NullableSizedLayout) IsSet() bool {
	return v.isSet
}

func (v *NullableSizedLayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSizedLayout(val *SizedLayout) *NullableSizedLayout {
	return &NullableSizedLayout{value: val, isSet: true}
}

func (v NullableSizedLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSizedLayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
