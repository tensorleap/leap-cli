/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.417
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the CodeIntegrationMappingErrorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CodeIntegrationMappingErrorsResponse{}

// CodeIntegrationMappingErrorsResponse struct for CodeIntegrationMappingErrorsResponse
type CodeIntegrationMappingErrorsResponse struct {
	MappingErrors []MappingError `json:"mappingErrors,omitempty"`
}

// NewCodeIntegrationMappingErrorsResponse instantiates a new CodeIntegrationMappingErrorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeIntegrationMappingErrorsResponse() *CodeIntegrationMappingErrorsResponse {
	this := CodeIntegrationMappingErrorsResponse{}
	return &this
}

// NewCodeIntegrationMappingErrorsResponseWithDefaults instantiates a new CodeIntegrationMappingErrorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeIntegrationMappingErrorsResponseWithDefaults() *CodeIntegrationMappingErrorsResponse {
	this := CodeIntegrationMappingErrorsResponse{}
	return &this
}

// GetMappingErrors returns the MappingErrors field value if set, zero value otherwise.
func (o *CodeIntegrationMappingErrorsResponse) GetMappingErrors() []MappingError {
	if o == nil || IsNil(o.MappingErrors) {
		var ret []MappingError
		return ret
	}
	return o.MappingErrors
}

// GetMappingErrorsOk returns a tuple with the MappingErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeIntegrationMappingErrorsResponse) GetMappingErrorsOk() ([]MappingError, bool) {
	if o == nil || IsNil(o.MappingErrors) {
		return nil, false
	}
	return o.MappingErrors, true
}

// HasMappingErrors returns a boolean if a field has been set.
func (o *CodeIntegrationMappingErrorsResponse) HasMappingErrors() bool {
	if o != nil && !IsNil(o.MappingErrors) {
		return true
	}

	return false
}

// SetMappingErrors gets a reference to the given []MappingError and assigns it to the MappingErrors field.
func (o *CodeIntegrationMappingErrorsResponse) SetMappingErrors(v []MappingError) {
	o.MappingErrors = v
}

func (o CodeIntegrationMappingErrorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodeIntegrationMappingErrorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MappingErrors) {
		toSerialize["mappingErrors"] = o.MappingErrors
	}
	return toSerialize, nil
}

type NullableCodeIntegrationMappingErrorsResponse struct {
	value *CodeIntegrationMappingErrorsResponse
	isSet bool
}

func (v NullableCodeIntegrationMappingErrorsResponse) Get() *CodeIntegrationMappingErrorsResponse {
	return v.value
}

func (v *NullableCodeIntegrationMappingErrorsResponse) Set(val *CodeIntegrationMappingErrorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeIntegrationMappingErrorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeIntegrationMappingErrorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeIntegrationMappingErrorsResponse(val *CodeIntegrationMappingErrorsResponse) *NullableCodeIntegrationMappingErrorsResponse {
	return &NullableCodeIntegrationMappingErrorsResponse{value: val, isSet: true}
}

func (v NullableCodeIntegrationMappingErrorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeIntegrationMappingErrorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
