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

// checks if the VisualizerInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisualizerInstance{}

// VisualizerInstance struct for VisualizerInstance
type VisualizerInstance struct {
	Name     string       `json:"name"`
	Type     LeapDataType `json:"type"`
	ArgNames []string     `json:"arg_names"`
}

// NewVisualizerInstance instantiates a new VisualizerInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisualizerInstance(name string, type_ LeapDataType, argNames []string) *VisualizerInstance {
	this := VisualizerInstance{}
	this.Name = name
	this.Type = type_
	this.ArgNames = argNames
	return &this
}

// NewVisualizerInstanceWithDefaults instantiates a new VisualizerInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisualizerInstanceWithDefaults() *VisualizerInstance {
	this := VisualizerInstance{}
	return &this
}

// GetName returns the Name field value
func (o *VisualizerInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VisualizerInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VisualizerInstance) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *VisualizerInstance) GetType() LeapDataType {
	if o == nil {
		var ret LeapDataType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VisualizerInstance) GetTypeOk() (*LeapDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VisualizerInstance) SetType(v LeapDataType) {
	o.Type = v
}

// GetArgNames returns the ArgNames field value
func (o *VisualizerInstance) GetArgNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ArgNames
}

// GetArgNamesOk returns a tuple with the ArgNames field value
// and a boolean to check if the value has been set.
func (o *VisualizerInstance) GetArgNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArgNames, true
}

// SetArgNames sets field value
func (o *VisualizerInstance) SetArgNames(v []string) {
	o.ArgNames = v
}

func (o VisualizerInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisualizerInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["arg_names"] = o.ArgNames
	return toSerialize, nil
}

type NullableVisualizerInstance struct {
	value *VisualizerInstance
	isSet bool
}

func (v NullableVisualizerInstance) Get() *VisualizerInstance {
	return v.value
}

func (v *NullableVisualizerInstance) Set(val *VisualizerInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualizerInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualizerInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualizerInstance(val *VisualizerInstance) *NullableVisualizerInstance {
	return &NullableVisualizerInstance{value: val, isSet: true}
}

func (v NullableVisualizerInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualizerInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
