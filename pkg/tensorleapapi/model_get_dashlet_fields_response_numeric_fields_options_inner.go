/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.592
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetDashletFieldsResponseNumericFieldsOptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDashletFieldsResponseNumericFieldsOptionsInner{}

// GetDashletFieldsResponseNumericFieldsOptionsInner struct for GetDashletFieldsResponseNumericFieldsOptionsInner
type GetDashletFieldsResponseNumericFieldsOptionsInner struct {
	Values []NumberOrString `json:"values"`
	Field  string           `json:"field"`
}

// NewGetDashletFieldsResponseNumericFieldsOptionsInner instantiates a new GetDashletFieldsResponseNumericFieldsOptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDashletFieldsResponseNumericFieldsOptionsInner(values []NumberOrString, field string) *GetDashletFieldsResponseNumericFieldsOptionsInner {
	this := GetDashletFieldsResponseNumericFieldsOptionsInner{}
	this.Values = values
	this.Field = field
	return &this
}

// NewGetDashletFieldsResponseNumericFieldsOptionsInnerWithDefaults instantiates a new GetDashletFieldsResponseNumericFieldsOptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDashletFieldsResponseNumericFieldsOptionsInnerWithDefaults() *GetDashletFieldsResponseNumericFieldsOptionsInner {
	this := GetDashletFieldsResponseNumericFieldsOptionsInner{}
	return &this
}

// GetValues returns the Values field value
func (o *GetDashletFieldsResponseNumericFieldsOptionsInner) GetValues() []NumberOrString {
	if o == nil {
		var ret []NumberOrString
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetDashletFieldsResponseNumericFieldsOptionsInner) GetValuesOk() ([]NumberOrString, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetDashletFieldsResponseNumericFieldsOptionsInner) SetValues(v []NumberOrString) {
	o.Values = v
}

// GetField returns the Field field value
func (o *GetDashletFieldsResponseNumericFieldsOptionsInner) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *GetDashletFieldsResponseNumericFieldsOptionsInner) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *GetDashletFieldsResponseNumericFieldsOptionsInner) SetField(v string) {
	o.Field = v
}

func (o GetDashletFieldsResponseNumericFieldsOptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDashletFieldsResponseNumericFieldsOptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	toSerialize["field"] = o.Field
	return toSerialize, nil
}

type NullableGetDashletFieldsResponseNumericFieldsOptionsInner struct {
	value *GetDashletFieldsResponseNumericFieldsOptionsInner
	isSet bool
}

func (v NullableGetDashletFieldsResponseNumericFieldsOptionsInner) Get() *GetDashletFieldsResponseNumericFieldsOptionsInner {
	return v.value
}

func (v *NullableGetDashletFieldsResponseNumericFieldsOptionsInner) Set(val *GetDashletFieldsResponseNumericFieldsOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDashletFieldsResponseNumericFieldsOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDashletFieldsResponseNumericFieldsOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDashletFieldsResponseNumericFieldsOptionsInner(val *GetDashletFieldsResponseNumericFieldsOptionsInner) *NullableGetDashletFieldsResponseNumericFieldsOptionsInner {
	return &NullableGetDashletFieldsResponseNumericFieldsOptionsInner{value: val, isSet: true}
}

func (v NullableGetDashletFieldsResponseNumericFieldsOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDashletFieldsResponseNumericFieldsOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
