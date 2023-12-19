/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.422
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the BoolSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoolSchema{}

// BoolSchema struct for BoolSchema
type BoolSchema struct {
	Type        string  `json:"type"`
	Title       *string `json:"title,omitempty"`
	Def         *bool   `json:"def,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewBoolSchema instantiates a new BoolSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoolSchema(type_ string) *BoolSchema {
	this := BoolSchema{}
	this.Type = type_
	return &this
}

// NewBoolSchemaWithDefaults instantiates a new BoolSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoolSchemaWithDefaults() *BoolSchema {
	this := BoolSchema{}
	return &this
}

// GetType returns the Type field value
func (o *BoolSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BoolSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BoolSchema) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BoolSchema) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoolSchema) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BoolSchema) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BoolSchema) SetTitle(v string) {
	o.Title = &v
}

// GetDef returns the Def field value if set, zero value otherwise.
func (o *BoolSchema) GetDef() bool {
	if o == nil || IsNil(o.Def) {
		var ret bool
		return ret
	}
	return *o.Def
}

// GetDefOk returns a tuple with the Def field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoolSchema) GetDefOk() (*bool, bool) {
	if o == nil || IsNil(o.Def) {
		return nil, false
	}
	return o.Def, true
}

// HasDef returns a boolean if a field has been set.
func (o *BoolSchema) HasDef() bool {
	if o != nil && !IsNil(o.Def) {
		return true
	}

	return false
}

// SetDef gets a reference to the given bool and assigns it to the Def field.
func (o *BoolSchema) SetDef(v bool) {
	o.Def = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BoolSchema) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoolSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BoolSchema) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BoolSchema) SetDescription(v string) {
	o.Description = &v
}

func (o BoolSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoolSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Def) {
		toSerialize["def"] = o.Def
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableBoolSchema struct {
	value *BoolSchema
	isSet bool
}

func (v NullableBoolSchema) Get() *BoolSchema {
	return v.value
}

func (v *NullableBoolSchema) Set(val *BoolSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableBoolSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableBoolSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoolSchema(val *BoolSchema) *NullableBoolSchema {
	return &NullableBoolSchema{value: val, isSet: true}
}

func (v NullableBoolSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoolSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
