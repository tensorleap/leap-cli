/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ModuleStatusAllModules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleStatusAllModules{}

// ModuleStatusAllModules struct for ModuleStatusAllModules
type ModuleStatusAllModules struct {
	Name MonitoredModuleType `json:"name"`
}

// NewModuleStatusAllModules instantiates a new ModuleStatusAllModules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleStatusAllModules(name MonitoredModuleType) *ModuleStatusAllModules {
	this := ModuleStatusAllModules{}
	this.Name = name
	return &this
}

// NewModuleStatusAllModulesWithDefaults instantiates a new ModuleStatusAllModules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleStatusAllModulesWithDefaults() *ModuleStatusAllModules {
	this := ModuleStatusAllModules{}
	return &this
}

// GetName returns the Name field value
func (o *ModuleStatusAllModules) GetName() MonitoredModuleType {
	if o == nil {
		var ret MonitoredModuleType
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModuleStatusAllModules) GetNameOk() (*MonitoredModuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModuleStatusAllModules) SetName(v MonitoredModuleType) {
	o.Name = v
}

func (o ModuleStatusAllModules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleStatusAllModules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableModuleStatusAllModules struct {
	value *ModuleStatusAllModules
	isSet bool
}

func (v NullableModuleStatusAllModules) Get() *ModuleStatusAllModules {
	return v.value
}

func (v *NullableModuleStatusAllModules) Set(val *ModuleStatusAllModules) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleStatusAllModules) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleStatusAllModules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleStatusAllModules(val *ModuleStatusAllModules) *NullableModuleStatusAllModules {
	return &NullableModuleStatusAllModules{value: val, isSet: true}
}

func (v NullableModuleStatusAllModules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleStatusAllModules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
