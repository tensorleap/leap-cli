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

// checks if the Clusters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Clusters{}

// Clusters struct for Clusters
type Clusters struct {
	Data   []float64 `json:"data"`
	Labels []string  `json:"labels"`
}

// NewClusters instantiates a new Clusters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusters(data []float64, labels []string) *Clusters {
	this := Clusters{}
	this.Data = data
	this.Labels = labels
	return &this
}

// NewClustersWithDefaults instantiates a new Clusters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClustersWithDefaults() *Clusters {
	this := Clusters{}
	return &this
}

// GetData returns the Data field value
func (o *Clusters) GetData() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Clusters) GetDataOk() ([]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *Clusters) SetData(v []float64) {
	o.Data = v
}

// GetLabels returns the Labels field value
func (o *Clusters) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *Clusters) GetLabelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *Clusters) SetLabels(v []string) {
	o.Labels = v
}

func (o Clusters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Clusters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["labels"] = o.Labels
	return toSerialize, nil
}

type NullableClusters struct {
	value *Clusters
	isSet bool
}

func (v NullableClusters) Get() *Clusters {
	return v.value
}

func (v *NullableClusters) Set(val *Clusters) {
	v.value = val
	v.isSet = true
}

func (v NullableClusters) IsSet() bool {
	return v.isSet
}

func (v *NullableClusters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusters(val *Clusters) *NullableClusters {
	return &NullableClusters{value: val, isSet: true}
}

func (v NullableClusters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
