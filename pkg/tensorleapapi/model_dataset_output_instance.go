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

// checks if the DatasetOutputInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetOutputInstance{}

// DatasetOutputInstance struct for DatasetOutputInstance
type DatasetOutputInstance struct {
	Name  string    `json:"name"`
	Shape []float64 `json:"shape"`
}

// NewDatasetOutputInstance instantiates a new DatasetOutputInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetOutputInstance(name string, shape []float64) *DatasetOutputInstance {
	this := DatasetOutputInstance{}
	this.Name = name
	this.Shape = shape
	return &this
}

// NewDatasetOutputInstanceWithDefaults instantiates a new DatasetOutputInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetOutputInstanceWithDefaults() *DatasetOutputInstance {
	this := DatasetOutputInstance{}
	return &this
}

// GetName returns the Name field value
func (o *DatasetOutputInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatasetOutputInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatasetOutputInstance) SetName(v string) {
	o.Name = v
}

// GetShape returns the Shape field value
func (o *DatasetOutputInstance) GetShape() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}

	return o.Shape
}

// GetShapeOk returns a tuple with the Shape field value
// and a boolean to check if the value has been set.
func (o *DatasetOutputInstance) GetShapeOk() ([]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shape, true
}

// SetShape sets field value
func (o *DatasetOutputInstance) SetShape(v []float64) {
	o.Shape = v
}

func (o DatasetOutputInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetOutputInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["shape"] = o.Shape
	return toSerialize, nil
}

type NullableDatasetOutputInstance struct {
	value *DatasetOutputInstance
	isSet bool
}

func (v NullableDatasetOutputInstance) Get() *DatasetOutputInstance {
	return v.value
}

func (v *NullableDatasetOutputInstance) Set(val *DatasetOutputInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetOutputInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetOutputInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetOutputInstance(val *DatasetOutputInstance) *NullableDatasetOutputInstance {
	return &NullableDatasetOutputInstance{value: val, isSet: true}
}

func (v NullableDatasetOutputInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetOutputInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
