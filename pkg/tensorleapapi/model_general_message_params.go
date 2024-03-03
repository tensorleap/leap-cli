/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GeneralMessageParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneralMessageParams{}

// GeneralMessageParams struct for GeneralMessageParams
type GeneralMessageParams struct {
	Message     string  `json:"message"`
	MessageCode *string `json:"message_code,omitempty"`
	Module      Module  `json:"module"`
}

// NewGeneralMessageParams instantiates a new GeneralMessageParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneralMessageParams(message string, module Module) *GeneralMessageParams {
	this := GeneralMessageParams{}
	this.Message = message
	this.Module = module
	return &this
}

// NewGeneralMessageParamsWithDefaults instantiates a new GeneralMessageParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralMessageParamsWithDefaults() *GeneralMessageParams {
	this := GeneralMessageParams{}
	return &this
}

// GetMessage returns the Message field value
func (o *GeneralMessageParams) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GeneralMessageParams) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *GeneralMessageParams) SetMessage(v string) {
	o.Message = v
}

// GetMessageCode returns the MessageCode field value if set, zero value otherwise.
func (o *GeneralMessageParams) GetMessageCode() string {
	if o == nil || IsNil(o.MessageCode) {
		var ret string
		return ret
	}
	return *o.MessageCode
}

// GetMessageCodeOk returns a tuple with the MessageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralMessageParams) GetMessageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MessageCode) {
		return nil, false
	}
	return o.MessageCode, true
}

// HasMessageCode returns a boolean if a field has been set.
func (o *GeneralMessageParams) HasMessageCode() bool {
	if o != nil && !IsNil(o.MessageCode) {
		return true
	}

	return false
}

// SetMessageCode gets a reference to the given string and assigns it to the MessageCode field.
func (o *GeneralMessageParams) SetMessageCode(v string) {
	o.MessageCode = &v
}

// GetModule returns the Module field value
func (o *GeneralMessageParams) GetModule() Module {
	if o == nil {
		var ret Module
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *GeneralMessageParams) GetModuleOk() (*Module, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *GeneralMessageParams) SetModule(v Module) {
	o.Module = v
}

func (o GeneralMessageParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeneralMessageParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.MessageCode) {
		toSerialize["message_code"] = o.MessageCode
	}
	toSerialize["module"] = o.Module
	return toSerialize, nil
}

type NullableGeneralMessageParams struct {
	value *GeneralMessageParams
	isSet bool
}

func (v NullableGeneralMessageParams) Get() *GeneralMessageParams {
	return v.value
}

func (v *NullableGeneralMessageParams) Set(val *GeneralMessageParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralMessageParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralMessageParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralMessageParams(val *GeneralMessageParams) *NullableGeneralMessageParams {
	return &NullableGeneralMessageParams{value: val, isSet: true}
}

func (v NullableGeneralMessageParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralMessageParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
