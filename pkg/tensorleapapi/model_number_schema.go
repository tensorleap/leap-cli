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

// checks if the NumberSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NumberSchema{}

// NumberSchema struct for NumberSchema
type NumberSchema struct {
	Type string `json:"type"`
	Title *string `json:"title,omitempty"`
	Def NullableFloat64 `json:"def,omitempty"`
	Description *string `json:"description,omitempty"`
	Options []float64 `json:"options,omitempty"`
	Step *float64 `json:"step,omitempty"`
	Max *float64 `json:"max,omitempty"`
	Min *float64 `json:"min,omitempty"`
	Required *bool `json:"required,omitempty"`
	Placeholder *string `json:"placeholder,omitempty"`
}

// NewNumberSchema instantiates a new NumberSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumberSchema(type_ string) *NumberSchema {
	this := NumberSchema{}
	this.Type = type_
	return &this
}

// NewNumberSchemaWithDefaults instantiates a new NumberSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumberSchemaWithDefaults() *NumberSchema {
	this := NumberSchema{}
	return &this
}

// GetType returns the Type field value
func (o *NumberSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NumberSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NumberSchema) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *NumberSchema) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberSchema) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *NumberSchema) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *NumberSchema) SetTitle(v string) {
	o.Title = &v
}

// GetDef returns the Def field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NumberSchema) GetDef() float64 {
	if o == nil || IsNil(o.Def.Get()) {
		var ret float64
		return ret
	}
	return *o.Def.Get()
}

// GetDefOk returns a tuple with the Def field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NumberSchema) GetDefOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Def.Get(), o.Def.IsSet()
}

// HasDef returns a boolean if a field has been set.
func (o *NumberSchema) HasDef() bool {
	if o != nil && o.Def.IsSet() {
		return true
	}

	return false
}

// SetDef gets a reference to the given NullableFloat64 and assigns it to the Def field.
func (o *NumberSchema) SetDef(v float64) {
	o.Def.Set(&v)
}
// SetDefNil sets the value for Def to be an explicit nil
func (o *NumberSchema) SetDefNil() {
	o.Def.Set(nil)
}

// UnsetDef ensures that no value is present for Def, not even an explicit nil
func (o *NumberSchema) UnsetDef() {
	o.Def.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NumberSchema) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NumberSchema) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NumberSchema) SetDescription(v string) {
	o.Description = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *NumberSchema) GetOptions() []float64 {
	if o == nil || IsNil(o.Options) {
		var ret []float64
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberSchema) GetOptionsOk() ([]float64, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *NumberSchema) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []float64 and assigns it to the Options field.
func (o *NumberSchema) SetOptions(v []float64) {
	o.Options = v
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *NumberSchema) GetStep() float64 {
	if o == nil || IsNil(o.Step) {
		var ret float64
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberSchema) GetStepOk() (*float64, bool) {
	if o == nil || IsNil(o.Step) {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *NumberSchema) HasStep() bool {
	if o != nil && !IsNil(o.Step) {
		return true
	}

	return false
}

// SetStep gets a reference to the given float64 and assigns it to the Step field.
func (o *NumberSchema) SetStep(v float64) {
	o.Step = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *NumberSchema) GetMax() float64 {
	if o == nil || IsNil(o.Max) {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberSchema) GetMaxOk() (*float64, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *NumberSchema) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *NumberSchema) SetMax(v float64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *NumberSchema) GetMin() float64 {
	if o == nil || IsNil(o.Min) {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberSchema) GetMinOk() (*float64, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *NumberSchema) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *NumberSchema) SetMin(v float64) {
	o.Min = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *NumberSchema) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberSchema) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *NumberSchema) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *NumberSchema) SetRequired(v bool) {
	o.Required = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *NumberSchema) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder) {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberSchema) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *NumberSchema) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *NumberSchema) SetPlaceholder(v string) {
	o.Placeholder = &v
}

func (o NumberSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NumberSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if o.Def.IsSet() {
		toSerialize["def"] = o.Def.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Step) {
		toSerialize["step"] = o.Step
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	return toSerialize, nil
}

type NullableNumberSchema struct {
	value *NumberSchema
	isSet bool
}

func (v NullableNumberSchema) Get() *NumberSchema {
	return v.value
}

func (v *NullableNumberSchema) Set(val *NumberSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableNumberSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableNumberSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumberSchema(val *NumberSchema) *NullableNumberSchema {
	return &NullableNumberSchema{value: val, isSet: true}
}

func (v NullableNumberSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumberSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


