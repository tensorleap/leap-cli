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

// checks if the ModelSetup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSetup{}

// ModelSetup struct for ModelSetup
type ModelSetup struct {
	CustomLayers []CustomLayerInstance `json:"custom_layers"`
}

// NewModelSetup instantiates a new ModelSetup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSetup(customLayers []CustomLayerInstance) *ModelSetup {
	this := ModelSetup{}
	this.CustomLayers = customLayers
	return &this
}

// NewModelSetupWithDefaults instantiates a new ModelSetup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSetupWithDefaults() *ModelSetup {
	this := ModelSetup{}
	return &this
}

// GetCustomLayers returns the CustomLayers field value
func (o *ModelSetup) GetCustomLayers() []CustomLayerInstance {
	if o == nil {
		var ret []CustomLayerInstance
		return ret
	}

	return o.CustomLayers
}

// GetCustomLayersOk returns a tuple with the CustomLayers field value
// and a boolean to check if the value has been set.
func (o *ModelSetup) GetCustomLayersOk() ([]CustomLayerInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomLayers, true
}

// SetCustomLayers sets field value
func (o *ModelSetup) SetCustomLayers(v []CustomLayerInstance) {
	o.CustomLayers = v
}

func (o ModelSetup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSetup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["custom_layers"] = o.CustomLayers
	return toSerialize, nil
}

type NullableModelSetup struct {
	value *ModelSetup
	isSet bool
}

func (v NullableModelSetup) Get() *ModelSetup {
	return v.value
}

func (v *NullableModelSetup) Set(val *ModelSetup) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSetup) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSetup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSetup(val *ModelSetup) *NullableModelSetup {
	return &NullableModelSetup{value: val, isSet: true}
}

func (v NullableModelSetup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSetup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
