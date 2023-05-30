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

// checks if the ESFilterValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ESFilterValue{}

// ESFilterValue struct for ESFilterValue
type ESFilterValue struct {
	Gt  *float64         `json:"gt,omitempty"`
	Lt  *float64         `json:"lt,omitempty"`
	Eq  *NumberOrString  `json:"eq,omitempty"`
	Lst []NumberOrString `json:"lst,omitempty"`
}

// NewESFilterValue instantiates a new ESFilterValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewESFilterValue() *ESFilterValue {
	this := ESFilterValue{}
	return &this
}

// NewESFilterValueWithDefaults instantiates a new ESFilterValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewESFilterValueWithDefaults() *ESFilterValue {
	this := ESFilterValue{}
	return &this
}

// GetGt returns the Gt field value if set, zero value otherwise.
func (o *ESFilterValue) GetGt() float64 {
	if o == nil || IsNil(o.Gt) {
		var ret float64
		return ret
	}
	return *o.Gt
}

// GetGtOk returns a tuple with the Gt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESFilterValue) GetGtOk() (*float64, bool) {
	if o == nil || IsNil(o.Gt) {
		return nil, false
	}
	return o.Gt, true
}

// HasGt returns a boolean if a field has been set.
func (o *ESFilterValue) HasGt() bool {
	if o != nil && !IsNil(o.Gt) {
		return true
	}

	return false
}

// SetGt gets a reference to the given float64 and assigns it to the Gt field.
func (o *ESFilterValue) SetGt(v float64) {
	o.Gt = &v
}

// GetLt returns the Lt field value if set, zero value otherwise.
func (o *ESFilterValue) GetLt() float64 {
	if o == nil || IsNil(o.Lt) {
		var ret float64
		return ret
	}
	return *o.Lt
}

// GetLtOk returns a tuple with the Lt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESFilterValue) GetLtOk() (*float64, bool) {
	if o == nil || IsNil(o.Lt) {
		return nil, false
	}
	return o.Lt, true
}

// HasLt returns a boolean if a field has been set.
func (o *ESFilterValue) HasLt() bool {
	if o != nil && !IsNil(o.Lt) {
		return true
	}

	return false
}

// SetLt gets a reference to the given float64 and assigns it to the Lt field.
func (o *ESFilterValue) SetLt(v float64) {
	o.Lt = &v
}

// GetEq returns the Eq field value if set, zero value otherwise.
func (o *ESFilterValue) GetEq() NumberOrString {
	if o == nil || IsNil(o.Eq) {
		var ret NumberOrString
		return ret
	}
	return *o.Eq
}

// GetEqOk returns a tuple with the Eq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESFilterValue) GetEqOk() (*NumberOrString, bool) {
	if o == nil || IsNil(o.Eq) {
		return nil, false
	}
	return o.Eq, true
}

// HasEq returns a boolean if a field has been set.
func (o *ESFilterValue) HasEq() bool {
	if o != nil && !IsNil(o.Eq) {
		return true
	}

	return false
}

// SetEq gets a reference to the given NumberOrString and assigns it to the Eq field.
func (o *ESFilterValue) SetEq(v NumberOrString) {
	o.Eq = &v
}

// GetLst returns the Lst field value if set, zero value otherwise.
func (o *ESFilterValue) GetLst() []NumberOrString {
	if o == nil || IsNil(o.Lst) {
		var ret []NumberOrString
		return ret
	}
	return o.Lst
}

// GetLstOk returns a tuple with the Lst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESFilterValue) GetLstOk() ([]NumberOrString, bool) {
	if o == nil || IsNil(o.Lst) {
		return nil, false
	}
	return o.Lst, true
}

// HasLst returns a boolean if a field has been set.
func (o *ESFilterValue) HasLst() bool {
	if o != nil && !IsNil(o.Lst) {
		return true
	}

	return false
}

// SetLst gets a reference to the given []NumberOrString and assigns it to the Lst field.
func (o *ESFilterValue) SetLst(v []NumberOrString) {
	o.Lst = v
}

func (o ESFilterValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ESFilterValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gt) {
		toSerialize["gt"] = o.Gt
	}
	if !IsNil(o.Lt) {
		toSerialize["lt"] = o.Lt
	}
	if !IsNil(o.Eq) {
		toSerialize["eq"] = o.Eq
	}
	if !IsNil(o.Lst) {
		toSerialize["lst"] = o.Lst
	}
	return toSerialize, nil
}

type NullableESFilterValue struct {
	value *ESFilterValue
	isSet bool
}

func (v NullableESFilterValue) Get() *ESFilterValue {
	return v.value
}

func (v *NullableESFilterValue) Set(val *ESFilterValue) {
	v.value = val
	v.isSet = true
}

func (v NullableESFilterValue) IsSet() bool {
	return v.isSet
}

func (v *NullableESFilterValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableESFilterValue(val *ESFilterValue) *NullableESFilterValue {
	return &NullableESFilterValue{value: val, isSet: true}
}

func (v NullableESFilterValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableESFilterValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
