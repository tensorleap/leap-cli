/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.365
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the DatasetInputInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetInputInstance{}

// DatasetInputInstance struct for DatasetInputInstance
type DatasetInputInstance struct {
	Name  string    `json:"name"`
	Shape []float64 `json:"shape"`
}

// NewDatasetInputInstance instantiates a new DatasetInputInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetInputInstance(name string, shape []float64) *DatasetInputInstance {
	this := DatasetInputInstance{}
	this.Name = name
	this.Shape = shape
	return &this
}

// NewDatasetInputInstanceWithDefaults instantiates a new DatasetInputInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetInputInstanceWithDefaults() *DatasetInputInstance {
	this := DatasetInputInstance{}
	return &this
}

// GetName returns the Name field value
func (o *DatasetInputInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatasetInputInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatasetInputInstance) SetName(v string) {
	o.Name = v
}

// GetShape returns the Shape field value
func (o *DatasetInputInstance) GetShape() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}

	return o.Shape
}

// GetShapeOk returns a tuple with the Shape field value
// and a boolean to check if the value has been set.
func (o *DatasetInputInstance) GetShapeOk() ([]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shape, true
}

// SetShape sets field value
func (o *DatasetInputInstance) SetShape(v []float64) {
	o.Shape = v
}

func (o DatasetInputInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetInputInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["shape"] = o.Shape
	return toSerialize, nil
}

type NullableDatasetInputInstance struct {
	value *DatasetInputInstance
	isSet bool
}

func (v NullableDatasetInputInstance) Get() *DatasetInputInstance {
	return v.value
}

func (v *NullableDatasetInputInstance) Set(val *DatasetInputInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetInputInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetInputInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetInputInstance(val *DatasetInputInstance) *NullableDatasetInputInstance {
	return &NullableDatasetInputInstance{value: val, isSet: true}
}

func (v NullableDatasetInputInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetInputInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
