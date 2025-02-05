/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.592
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the DatasetMetadataInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetMetadataInstance{}

// DatasetMetadataInstance struct for DatasetMetadataInstance
type DatasetMetadataInstance struct {
	Name string              `json:"name"`
	Type DatasetMetadataType `json:"type"`
}

// NewDatasetMetadataInstance instantiates a new DatasetMetadataInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetMetadataInstance(name string, type_ DatasetMetadataType) *DatasetMetadataInstance {
	this := DatasetMetadataInstance{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewDatasetMetadataInstanceWithDefaults instantiates a new DatasetMetadataInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetMetadataInstanceWithDefaults() *DatasetMetadataInstance {
	this := DatasetMetadataInstance{}
	return &this
}

// GetName returns the Name field value
func (o *DatasetMetadataInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatasetMetadataInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatasetMetadataInstance) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *DatasetMetadataInstance) GetType() DatasetMetadataType {
	if o == nil {
		var ret DatasetMetadataType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatasetMetadataInstance) GetTypeOk() (*DatasetMetadataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DatasetMetadataInstance) SetType(v DatasetMetadataType) {
	o.Type = v
}

func (o DatasetMetadataInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetMetadataInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableDatasetMetadataInstance struct {
	value *DatasetMetadataInstance
	isSet bool
}

func (v NullableDatasetMetadataInstance) Get() *DatasetMetadataInstance {
	return v.value
}

func (v *NullableDatasetMetadataInstance) Set(val *DatasetMetadataInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetMetadataInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetMetadataInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetMetadataInstance(val *DatasetMetadataInstance) *NullableDatasetMetadataInstance {
	return &NullableDatasetMetadataInstance{value: val, isSet: true}
}

func (v NullableDatasetMetadataInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetMetadataInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
