/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.301
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ClientFilterParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientFilterParams{}

// ClientFilterParams struct for ClientFilterParams
type ClientFilterParams struct {
	Field    string             `json:"field"`
	Operator FilterOperatorType `json:"operator"`
	Value    float64            `json:"value"`
	Disable  *bool              `json:"disable,omitempty"`
}

// NewClientFilterParams instantiates a new ClientFilterParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientFilterParams(field string, operator FilterOperatorType, value float64) *ClientFilterParams {
	this := ClientFilterParams{}
	this.Field = field
	this.Operator = operator
	this.Value = value
	return &this
}

// NewClientFilterParamsWithDefaults instantiates a new ClientFilterParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientFilterParamsWithDefaults() *ClientFilterParams {
	this := ClientFilterParams{}
	return &this
}

// GetField returns the Field field value
func (o *ClientFilterParams) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *ClientFilterParams) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *ClientFilterParams) SetField(v string) {
	o.Field = v
}

// GetOperator returns the Operator field value
func (o *ClientFilterParams) GetOperator() FilterOperatorType {
	if o == nil {
		var ret FilterOperatorType
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ClientFilterParams) GetOperatorOk() (*FilterOperatorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *ClientFilterParams) SetOperator(v FilterOperatorType) {
	o.Operator = v
}

// GetValue returns the Value field value
func (o *ClientFilterParams) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ClientFilterParams) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ClientFilterParams) SetValue(v float64) {
	o.Value = v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *ClientFilterParams) GetDisable() bool {
	if o == nil || IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientFilterParams) GetDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *ClientFilterParams) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *ClientFilterParams) SetDisable(v bool) {
	o.Disable = &v
}

func (o ClientFilterParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientFilterParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	toSerialize["operator"] = o.Operator
	toSerialize["value"] = o.Value
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	return toSerialize, nil
}

type NullableClientFilterParams struct {
	value *ClientFilterParams
	isSet bool
}

func (v NullableClientFilterParams) Get() *ClientFilterParams {
	return v.value
}

func (v *NullableClientFilterParams) Set(val *ClientFilterParams) {
	v.value = val
	v.isSet = true
}

func (v NullableClientFilterParams) IsSet() bool {
	return v.isSet
}

func (v *NullableClientFilterParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientFilterParams(val *ClientFilterParams) *NullableClientFilterParams {
	return &NullableClientFilterParams{value: val, isSet: true}
}

func (v NullableClientFilterParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientFilterParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
