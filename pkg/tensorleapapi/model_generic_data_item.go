/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GenericDataItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericDataItem{}

// GenericDataItem struct for GenericDataItem
type GenericDataItem struct {
	// Construct a type with a set of properties K of type T
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Construct a type with a set of properties K of type T
	Data map[string]interface{} `json:"data"`
}

// NewGenericDataItem instantiates a new GenericDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericDataItem(data map[string]interface{}) *GenericDataItem {
	this := GenericDataItem{}
	this.Data = data
	return &this
}

// NewGenericDataItemWithDefaults instantiates a new GenericDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericDataItemWithDefaults() *GenericDataItem {
	this := GenericDataItem{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GenericDataItem) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericDataItem) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GenericDataItem) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GenericDataItem) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *GenericDataItem) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GenericDataItem) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GenericDataItem) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o GenericDataItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericDataItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGenericDataItem struct {
	value *GenericDataItem
	isSet bool
}

func (v NullableGenericDataItem) Get() *GenericDataItem {
	return v.value
}

func (v *NullableGenericDataItem) Set(val *GenericDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericDataItem(val *GenericDataItem) *NullableGenericDataItem {
	return &NullableGenericDataItem{value: val, isSet: true}
}

func (v NullableGenericDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
