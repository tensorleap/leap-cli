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

// checks if the ModuleStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleStatus{}

// ModuleStatus struct for ModuleStatus
type ModuleStatus struct {
	Status bool                `json:"status"`
	Error  *string             `json:"error,omitempty"`
	Name   MonitoredModuleType `json:"name"`
}

// NewModuleStatus instantiates a new ModuleStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleStatus(status bool, name MonitoredModuleType) *ModuleStatus {
	this := ModuleStatus{}
	this.Status = status
	this.Name = name
	return &this
}

// NewModuleStatusWithDefaults instantiates a new ModuleStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleStatusWithDefaults() *ModuleStatus {
	this := ModuleStatus{}
	return &this
}

// GetStatus returns the Status field value
func (o *ModuleStatus) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ModuleStatus) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ModuleStatus) SetStatus(v bool) {
	o.Status = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ModuleStatus) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatus) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ModuleStatus) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ModuleStatus) SetError(v string) {
	o.Error = &v
}

// GetName returns the Name field value
func (o *ModuleStatus) GetName() MonitoredModuleType {
	if o == nil {
		var ret MonitoredModuleType
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModuleStatus) GetNameOk() (*MonitoredModuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModuleStatus) SetName(v MonitoredModuleType) {
	o.Name = v
}

func (o ModuleStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableModuleStatus struct {
	value *ModuleStatus
	isSet bool
}

func (v NullableModuleStatus) Get() *ModuleStatus {
	return v.value
}

func (v *NullableModuleStatus) Set(val *ModuleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleStatus(val *ModuleStatus) *NullableModuleStatus {
	return &NullableModuleStatus{value: val, isSet: true}
}

func (v NullableModuleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
