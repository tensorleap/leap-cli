/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the MappingError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MappingError{}

// MappingError struct for MappingError
type MappingError struct {
	Type    MappingErrorType `json:"type"`
	Message string           `json:"message"`
	Title   string           `json:"title"`
}

// NewMappingError instantiates a new MappingError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMappingError(type_ MappingErrorType, message string, title string) *MappingError {
	this := MappingError{}
	this.Type = type_
	this.Message = message
	this.Title = title
	return &this
}

// NewMappingErrorWithDefaults instantiates a new MappingError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMappingErrorWithDefaults() *MappingError {
	this := MappingError{}
	return &this
}

// GetType returns the Type field value
func (o *MappingError) GetType() MappingErrorType {
	if o == nil {
		var ret MappingErrorType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MappingError) GetTypeOk() (*MappingErrorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MappingError) SetType(v MappingErrorType) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *MappingError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *MappingError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *MappingError) SetMessage(v string) {
	o.Message = v
}

// GetTitle returns the Title field value
func (o *MappingError) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *MappingError) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *MappingError) SetTitle(v string) {
	o.Title = v
}

func (o MappingError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MappingError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["message"] = o.Message
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

type NullableMappingError struct {
	value *MappingError
	isSet bool
}

func (v NullableMappingError) Get() *MappingError {
	return v.value
}

func (v *NullableMappingError) Set(val *MappingError) {
	v.value = val
	v.isSet = true
}

func (v NullableMappingError) IsSet() bool {
	return v.isSet
}

func (v *NullableMappingError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMappingError(val *MappingError) *NullableMappingError {
	return &NullableMappingError{value: val, isSet: true}
}

func (v NullableMappingError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMappingError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
