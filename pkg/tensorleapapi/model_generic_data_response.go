/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GenericDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericDataResponse{}

// GenericDataResponse struct for GenericDataResponse
type GenericDataResponse struct {
	Data []GenericDataItem `json:"data"`
}

// NewGenericDataResponse instantiates a new GenericDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericDataResponse(data []GenericDataItem) *GenericDataResponse {
	this := GenericDataResponse{}
	this.Data = data
	return &this
}

// NewGenericDataResponseWithDefaults instantiates a new GenericDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericDataResponseWithDefaults() *GenericDataResponse {
	this := GenericDataResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GenericDataResponse) GetData() []GenericDataItem {
	if o == nil {
		var ret []GenericDataItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GenericDataResponse) GetDataOk() ([]GenericDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GenericDataResponse) SetData(v []GenericDataItem) {
	o.Data = v
}

func (o GenericDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGenericDataResponse struct {
	value *GenericDataResponse
	isSet bool
}

func (v NullableGenericDataResponse) Get() *GenericDataResponse {
	return v.value
}

func (v *NullableGenericDataResponse) Set(val *GenericDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericDataResponse(val *GenericDataResponse) *NullableGenericDataResponse {
	return &NullableGenericDataResponse{value: val, isSet: true}
}

func (v NullableGenericDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
