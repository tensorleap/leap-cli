/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.526
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the CustomLayerInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomLayerInstance{}

// CustomLayerInstance struct for CustomLayerInstance
type CustomLayerInstance struct {
	Name                 string   `json:"name"`
	InitArgNames         []string `json:"init_arg_names"`
	CallArgNames         []string `json:"call_arg_names"`
	UseCustomLatentSpace *bool    `json:"use_custom_latent_space,omitempty"`
}

// NewCustomLayerInstance instantiates a new CustomLayerInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomLayerInstance(name string, initArgNames []string, callArgNames []string) *CustomLayerInstance {
	this := CustomLayerInstance{}
	this.Name = name
	this.InitArgNames = initArgNames
	this.CallArgNames = callArgNames
	return &this
}

// NewCustomLayerInstanceWithDefaults instantiates a new CustomLayerInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomLayerInstanceWithDefaults() *CustomLayerInstance {
	this := CustomLayerInstance{}
	return &this
}

// GetName returns the Name field value
func (o *CustomLayerInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomLayerInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomLayerInstance) SetName(v string) {
	o.Name = v
}

// GetInitArgNames returns the InitArgNames field value
func (o *CustomLayerInstance) GetInitArgNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.InitArgNames
}

// GetInitArgNamesOk returns a tuple with the InitArgNames field value
// and a boolean to check if the value has been set.
func (o *CustomLayerInstance) GetInitArgNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitArgNames, true
}

// SetInitArgNames sets field value
func (o *CustomLayerInstance) SetInitArgNames(v []string) {
	o.InitArgNames = v
}

// GetCallArgNames returns the CallArgNames field value
func (o *CustomLayerInstance) GetCallArgNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CallArgNames
}

// GetCallArgNamesOk returns a tuple with the CallArgNames field value
// and a boolean to check if the value has been set.
func (o *CustomLayerInstance) GetCallArgNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CallArgNames, true
}

// SetCallArgNames sets field value
func (o *CustomLayerInstance) SetCallArgNames(v []string) {
	o.CallArgNames = v
}

// GetUseCustomLatentSpace returns the UseCustomLatentSpace field value if set, zero value otherwise.
func (o *CustomLayerInstance) GetUseCustomLatentSpace() bool {
	if o == nil || IsNil(o.UseCustomLatentSpace) {
		var ret bool
		return ret
	}
	return *o.UseCustomLatentSpace
}

// GetUseCustomLatentSpaceOk returns a tuple with the UseCustomLatentSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLayerInstance) GetUseCustomLatentSpaceOk() (*bool, bool) {
	if o == nil || IsNil(o.UseCustomLatentSpace) {
		return nil, false
	}
	return o.UseCustomLatentSpace, true
}

// HasUseCustomLatentSpace returns a boolean if a field has been set.
func (o *CustomLayerInstance) HasUseCustomLatentSpace() bool {
	if o != nil && !IsNil(o.UseCustomLatentSpace) {
		return true
	}

	return false
}

// SetUseCustomLatentSpace gets a reference to the given bool and assigns it to the UseCustomLatentSpace field.
func (o *CustomLayerInstance) SetUseCustomLatentSpace(v bool) {
	o.UseCustomLatentSpace = &v
}

func (o CustomLayerInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomLayerInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["init_arg_names"] = o.InitArgNames
	toSerialize["call_arg_names"] = o.CallArgNames
	if !IsNil(o.UseCustomLatentSpace) {
		toSerialize["use_custom_latent_space"] = o.UseCustomLatentSpace
	}
	return toSerialize, nil
}

type NullableCustomLayerInstance struct {
	value *CustomLayerInstance
	isSet bool
}

func (v NullableCustomLayerInstance) Get() *CustomLayerInstance {
	return v.value
}

func (v *NullableCustomLayerInstance) Set(val *CustomLayerInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomLayerInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomLayerInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomLayerInstance(val *CustomLayerInstance) *NullableCustomLayerInstance {
	return &NullableCustomLayerInstance{value: val, isSet: true}
}

func (v NullableCustomLayerInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomLayerInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
