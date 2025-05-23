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

// checks if the Aggregations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Aggregations{}

// Aggregations struct for Aggregations
type Aggregations struct {
	Field       string            `json:"field"`
	Aggregation AggregationMethod `json:"aggregation"`
}

// NewAggregations instantiates a new Aggregations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregations(field string, aggregation AggregationMethod) *Aggregations {
	this := Aggregations{}
	this.Field = field
	this.Aggregation = aggregation
	return &this
}

// NewAggregationsWithDefaults instantiates a new Aggregations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationsWithDefaults() *Aggregations {
	this := Aggregations{}
	return &this
}

// GetField returns the Field field value
func (o *Aggregations) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *Aggregations) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *Aggregations) SetField(v string) {
	o.Field = v
}

// GetAggregation returns the Aggregation field value
func (o *Aggregations) GetAggregation() AggregationMethod {
	if o == nil {
		var ret AggregationMethod
		return ret
	}

	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *Aggregations) GetAggregationOk() (*AggregationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value
func (o *Aggregations) SetAggregation(v AggregationMethod) {
	o.Aggregation = v
}

func (o Aggregations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Aggregations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	toSerialize["aggregation"] = o.Aggregation
	return toSerialize, nil
}

type NullableAggregations struct {
	value *Aggregations
	isSet bool
}

func (v NullableAggregations) Get() *Aggregations {
	return v.value
}

func (v *NullableAggregations) Set(val *Aggregations) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregations) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregations(val *Aggregations) *NullableAggregations {
	return &NullableAggregations{value: val, isSet: true}
}

func (v NullableAggregations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
