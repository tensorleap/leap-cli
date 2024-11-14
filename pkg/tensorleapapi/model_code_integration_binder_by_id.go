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

// checks if the CodeIntegrationBinderById type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CodeIntegrationBinderById{}

// CodeIntegrationBinderById struct for CodeIntegrationBinderById
type CodeIntegrationBinderById struct {
	CodeIntegrationVersionId string `json:"codeIntegrationVersionId"`
	Type                     string `json:"type"`
}

// NewCodeIntegrationBinderById instantiates a new CodeIntegrationBinderById object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeIntegrationBinderById(codeIntegrationVersionId string, type_ string) *CodeIntegrationBinderById {
	this := CodeIntegrationBinderById{}
	this.CodeIntegrationVersionId = codeIntegrationVersionId
	this.Type = type_
	return &this
}

// NewCodeIntegrationBinderByIdWithDefaults instantiates a new CodeIntegrationBinderById object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeIntegrationBinderByIdWithDefaults() *CodeIntegrationBinderById {
	this := CodeIntegrationBinderById{}
	return &this
}

// GetCodeIntegrationVersionId returns the CodeIntegrationVersionId field value
func (o *CodeIntegrationBinderById) GetCodeIntegrationVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CodeIntegrationVersionId
}

// GetCodeIntegrationVersionIdOk returns a tuple with the CodeIntegrationVersionId field value
// and a boolean to check if the value has been set.
func (o *CodeIntegrationBinderById) GetCodeIntegrationVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeIntegrationVersionId, true
}

// SetCodeIntegrationVersionId sets field value
func (o *CodeIntegrationBinderById) SetCodeIntegrationVersionId(v string) {
	o.CodeIntegrationVersionId = v
}

// GetType returns the Type field value
func (o *CodeIntegrationBinderById) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CodeIntegrationBinderById) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CodeIntegrationBinderById) SetType(v string) {
	o.Type = v
}

func (o CodeIntegrationBinderById) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodeIntegrationBinderById) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["codeIntegrationVersionId"] = o.CodeIntegrationVersionId
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCodeIntegrationBinderById struct {
	value *CodeIntegrationBinderById
	isSet bool
}

func (v NullableCodeIntegrationBinderById) Get() *CodeIntegrationBinderById {
	return v.value
}

func (v *NullableCodeIntegrationBinderById) Set(val *CodeIntegrationBinderById) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeIntegrationBinderById) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeIntegrationBinderById) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeIntegrationBinderById(val *CodeIntegrationBinderById) *NullableCodeIntegrationBinderById {
	return &NullableCodeIntegrationBinderById{value: val, isSet: true}
}

func (v NullableCodeIntegrationBinderById) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeIntegrationBinderById) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
