/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.423
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the MetricInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricInstance{}

// MetricInstance struct for MetricInstance
type MetricInstance struct {
	Name string `json:"name"`
	ArgNames []string `json:"arg_names"`
}

// NewMetricInstance instantiates a new MetricInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricInstance(name string, argNames []string) *MetricInstance {
	this := MetricInstance{}
	this.Name = name
	this.ArgNames = argNames
	return &this
}

// NewMetricInstanceWithDefaults instantiates a new MetricInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricInstanceWithDefaults() *MetricInstance {
	this := MetricInstance{}
	return &this
}

// GetName returns the Name field value
func (o *MetricInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetricInstance) SetName(v string) {
	o.Name = v
}

// GetArgNames returns the ArgNames field value
func (o *MetricInstance) GetArgNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ArgNames
}

// GetArgNamesOk returns a tuple with the ArgNames field value
// and a boolean to check if the value has been set.
func (o *MetricInstance) GetArgNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArgNames, true
}

// SetArgNames sets field value
func (o *MetricInstance) SetArgNames(v []string) {
	o.ArgNames = v
}

func (o MetricInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["arg_names"] = o.ArgNames
	return toSerialize, nil
}

type NullableMetricInstance struct {
	value *MetricInstance
	isSet bool
}

func (v NullableMetricInstance) Get() *MetricInstance {
	return v.value
}

func (v *NullableMetricInstance) Set(val *MetricInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricInstance(val *MetricInstance) *NullableMetricInstance {
	return &NullableMetricInstance{value: val, isSet: true}
}

func (v NullableMetricInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


