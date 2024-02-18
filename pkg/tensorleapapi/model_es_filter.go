/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.456
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ESFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ESFilter{}

// ESFilter struct for ESFilter
type ESFilter struct {
	Operator    FilterOperatorType `json:"operator"`
	Field       string             `json:"field"`
	Disable     *bool              `json:"disable,omitempty"`
	Value       ESFilterValue      `json:"value"`
	DisplayData *FilterDisplayData `json:"displayData,omitempty"`
}

// NewESFilter instantiates a new ESFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewESFilter(operator FilterOperatorType, field string, value ESFilterValue) *ESFilter {
	this := ESFilter{}
	this.Operator = operator
	this.Field = field
	this.Value = value
	return &this
}

// NewESFilterWithDefaults instantiates a new ESFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewESFilterWithDefaults() *ESFilter {
	this := ESFilter{}
	return &this
}

// GetOperator returns the Operator field value
func (o *ESFilter) GetOperator() FilterOperatorType {
	if o == nil {
		var ret FilterOperatorType
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ESFilter) GetOperatorOk() (*FilterOperatorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *ESFilter) SetOperator(v FilterOperatorType) {
	o.Operator = v
}

// GetField returns the Field field value
func (o *ESFilter) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *ESFilter) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *ESFilter) SetField(v string) {
	o.Field = v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *ESFilter) GetDisable() bool {
	if o == nil || IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESFilter) GetDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *ESFilter) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *ESFilter) SetDisable(v bool) {
	o.Disable = &v
}

// GetValue returns the Value field value
func (o *ESFilter) GetValue() ESFilterValue {
	if o == nil {
		var ret ESFilterValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ESFilter) GetValueOk() (*ESFilterValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ESFilter) SetValue(v ESFilterValue) {
	o.Value = v
}

// GetDisplayData returns the DisplayData field value if set, zero value otherwise.
func (o *ESFilter) GetDisplayData() FilterDisplayData {
	if o == nil || IsNil(o.DisplayData) {
		var ret FilterDisplayData
		return ret
	}
	return *o.DisplayData
}

// GetDisplayDataOk returns a tuple with the DisplayData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESFilter) GetDisplayDataOk() (*FilterDisplayData, bool) {
	if o == nil || IsNil(o.DisplayData) {
		return nil, false
	}
	return o.DisplayData, true
}

// HasDisplayData returns a boolean if a field has been set.
func (o *ESFilter) HasDisplayData() bool {
	if o != nil && !IsNil(o.DisplayData) {
		return true
	}

	return false
}

// SetDisplayData gets a reference to the given FilterDisplayData and assigns it to the DisplayData field.
func (o *ESFilter) SetDisplayData(v FilterDisplayData) {
	o.DisplayData = &v
}

func (o ESFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ESFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operator"] = o.Operator
	toSerialize["field"] = o.Field
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	toSerialize["value"] = o.Value
	if !IsNil(o.DisplayData) {
		toSerialize["displayData"] = o.DisplayData
	}
	return toSerialize, nil
}

type NullableESFilter struct {
	value *ESFilter
	isSet bool
}

func (v NullableESFilter) Get() *ESFilter {
	return v.value
}

func (v *NullableESFilter) Set(val *ESFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableESFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableESFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableESFilter(val *ESFilter) *NullableESFilter {
	return &NullableESFilter{value: val, isSet: true}
}

func (v NullableESFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableESFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
