/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.557
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the StringSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringSchema{}

// StringSchema struct for StringSchema
type StringSchema struct {
	Type        string         `json:"type"`
	Title       *string        `json:"title,omitempty"`
	Def         NullableString `json:"def,omitempty"`
	Description *string        `json:"description,omitempty"`
	Options     []string       `json:"options,omitempty"`
	MaxLength   *float64       `json:"maxLength,omitempty"`
	MinLength   *float64       `json:"minLength,omitempty"`
	Pattern     *string        `json:"pattern,omitempty"`
	Required    *bool          `json:"required,omitempty"`
	Placeholder *string        `json:"placeholder,omitempty"`
}

// NewStringSchema instantiates a new StringSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringSchema(type_ string) *StringSchema {
	this := StringSchema{}
	this.Type = type_
	return &this
}

// NewStringSchemaWithDefaults instantiates a new StringSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringSchemaWithDefaults() *StringSchema {
	this := StringSchema{}
	return &this
}

// GetType returns the Type field value
func (o *StringSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StringSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StringSchema) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *StringSchema) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringSchema) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *StringSchema) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *StringSchema) SetTitle(v string) {
	o.Title = &v
}

// GetDef returns the Def field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringSchema) GetDef() string {
	if o == nil || IsNil(o.Def.Get()) {
		var ret string
		return ret
	}
	return *o.Def.Get()
}

// GetDefOk returns a tuple with the Def field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringSchema) GetDefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Def.Get(), o.Def.IsSet()
}

// HasDef returns a boolean if a field has been set.
func (o *StringSchema) HasDef() bool {
	if o != nil && o.Def.IsSet() {
		return true
	}

	return false
}

// SetDef gets a reference to the given NullableString and assigns it to the Def field.
func (o *StringSchema) SetDef(v string) {
	o.Def.Set(&v)
}

// SetDefNil sets the value for Def to be an explicit nil
func (o *StringSchema) SetDefNil() {
	o.Def.Set(nil)
}

// UnsetDef ensures that no value is present for Def, not even an explicit nil
func (o *StringSchema) UnsetDef() {
	o.Def.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StringSchema) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StringSchema) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StringSchema) SetDescription(v string) {
	o.Description = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *StringSchema) GetOptions() []string {
	if o == nil || IsNil(o.Options) {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringSchema) GetOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *StringSchema) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *StringSchema) SetOptions(v []string) {
	o.Options = v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *StringSchema) GetMaxLength() float64 {
	if o == nil || IsNil(o.MaxLength) {
		var ret float64
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringSchema) GetMaxLengthOk() (*float64, bool) {
	if o == nil || IsNil(o.MaxLength) {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *StringSchema) HasMaxLength() bool {
	if o != nil && !IsNil(o.MaxLength) {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given float64 and assigns it to the MaxLength field.
func (o *StringSchema) SetMaxLength(v float64) {
	o.MaxLength = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *StringSchema) GetMinLength() float64 {
	if o == nil || IsNil(o.MinLength) {
		var ret float64
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringSchema) GetMinLengthOk() (*float64, bool) {
	if o == nil || IsNil(o.MinLength) {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *StringSchema) HasMinLength() bool {
	if o != nil && !IsNil(o.MinLength) {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given float64 and assigns it to the MinLength field.
func (o *StringSchema) SetMinLength(v float64) {
	o.MinLength = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *StringSchema) GetPattern() string {
	if o == nil || IsNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringSchema) GetPatternOk() (*string, bool) {
	if o == nil || IsNil(o.Pattern) {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *StringSchema) HasPattern() bool {
	if o != nil && !IsNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *StringSchema) SetPattern(v string) {
	o.Pattern = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *StringSchema) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringSchema) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *StringSchema) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *StringSchema) SetRequired(v bool) {
	o.Required = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *StringSchema) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder) {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringSchema) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *StringSchema) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *StringSchema) SetPlaceholder(v string) {
	o.Placeholder = &v
}

func (o StringSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringSchema) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.MaxLength) {
		toSerialize["maxLength"] = o.MaxLength
	}
	if !IsNil(o.MinLength) {
		toSerialize["minLength"] = o.MinLength
	}
	if !IsNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	return toSerialize, nil
}

type NullableStringSchema struct {
	value *StringSchema
	isSet bool
}

func (v NullableStringSchema) Get() *StringSchema {
	return v.value
}

func (v *NullableStringSchema) Set(val *StringSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableStringSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableStringSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringSchema(val *StringSchema) *NullableStringSchema {
	return &NullableStringSchema{value: val, isSet: true}
}

func (v NullableStringSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
