/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.445
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the HealthCheckResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckResponse{}

// HealthCheckResponse struct for HealthCheckResponse
type HealthCheckResponse struct {
	AllModules []ModuleStatus `json:"allModules"`
}

// NewHealthCheckResponse instantiates a new HealthCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckResponse(allModules []ModuleStatus) *HealthCheckResponse {
	this := HealthCheckResponse{}
	this.AllModules = allModules
	return &this
}

// NewHealthCheckResponseWithDefaults instantiates a new HealthCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckResponseWithDefaults() *HealthCheckResponse {
	this := HealthCheckResponse{}
	return &this
}

// GetAllModules returns the AllModules field value
func (o *HealthCheckResponse) GetAllModules() []ModuleStatus {
	if o == nil {
		var ret []ModuleStatus
		return ret
	}

	return o.AllModules
}

// GetAllModulesOk returns a tuple with the AllModules field value
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetAllModulesOk() ([]ModuleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllModules, true
}

// SetAllModules sets field value
func (o *HealthCheckResponse) SetAllModules(v []ModuleStatus) {
	o.AllModules = v
}

func (o HealthCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allModules"] = o.AllModules
	return toSerialize, nil
}

type NullableHealthCheckResponse struct {
	value *HealthCheckResponse
	isSet bool
}

func (v NullableHealthCheckResponse) Get() *HealthCheckResponse {
	return v.value
}

func (v *NullableHealthCheckResponse) Set(val *HealthCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckResponse(val *HealthCheckResponse) *NullableHealthCheckResponse {
	return &NullableHealthCheckResponse{value: val, isSet: true}
}

func (v NullableHealthCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
