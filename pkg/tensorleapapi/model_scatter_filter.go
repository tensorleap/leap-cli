/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ScatterFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScatterFilter{}

// ScatterFilter struct for ScatterFilter
type ScatterFilter struct {
	Metric    string         `json:"metric"`
	Value     string         `json:"value"`
	Operator  OperatorEnum   `json:"operator"`
	DataState *DataStateType `json:"data_state,omitempty"`
}

// NewScatterFilter instantiates a new ScatterFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScatterFilter(metric string, value string, operator OperatorEnum) *ScatterFilter {
	this := ScatterFilter{}
	this.Metric = metric
	this.Value = value
	this.Operator = operator
	return &this
}

// NewScatterFilterWithDefaults instantiates a new ScatterFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScatterFilterWithDefaults() *ScatterFilter {
	this := ScatterFilter{}
	return &this
}

// GetMetric returns the Metric field value
func (o *ScatterFilter) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *ScatterFilter) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *ScatterFilter) SetMetric(v string) {
	o.Metric = v
}

// GetValue returns the Value field value
func (o *ScatterFilter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ScatterFilter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ScatterFilter) SetValue(v string) {
	o.Value = v
}

// GetOperator returns the Operator field value
func (o *ScatterFilter) GetOperator() OperatorEnum {
	if o == nil {
		var ret OperatorEnum
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ScatterFilter) GetOperatorOk() (*OperatorEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *ScatterFilter) SetOperator(v OperatorEnum) {
	o.Operator = v
}

// GetDataState returns the DataState field value if set, zero value otherwise.
func (o *ScatterFilter) GetDataState() DataStateType {
	if o == nil || IsNil(o.DataState) {
		var ret DataStateType
		return ret
	}
	return *o.DataState
}

// GetDataStateOk returns a tuple with the DataState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterFilter) GetDataStateOk() (*DataStateType, bool) {
	if o == nil || IsNil(o.DataState) {
		return nil, false
	}
	return o.DataState, true
}

// HasDataState returns a boolean if a field has been set.
func (o *ScatterFilter) HasDataState() bool {
	if o != nil && !IsNil(o.DataState) {
		return true
	}

	return false
}

// SetDataState gets a reference to the given DataStateType and assigns it to the DataState field.
func (o *ScatterFilter) SetDataState(v DataStateType) {
	o.DataState = &v
}

func (o ScatterFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScatterFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metric"] = o.Metric
	toSerialize["value"] = o.Value
	toSerialize["operator"] = o.Operator
	if !IsNil(o.DataState) {
		toSerialize["data_state"] = o.DataState
	}
	return toSerialize, nil
}

type NullableScatterFilter struct {
	value *ScatterFilter
	isSet bool
}

func (v NullableScatterFilter) Get() *ScatterFilter {
	return v.value
}

func (v *NullableScatterFilter) Set(val *ScatterFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableScatterFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableScatterFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScatterFilter(val *ScatterFilter) *NullableScatterFilter {
	return &NullableScatterFilter{value: val, isSet: true}
}

func (v NullableScatterFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScatterFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
