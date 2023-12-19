/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.422
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the AddDatasetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddDatasetResponse{}

// AddDatasetResponse struct for AddDatasetResponse
type AddDatasetResponse struct {
	Dataset Dataset `json:"dataset"`
}

// NewAddDatasetResponse instantiates a new AddDatasetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDatasetResponse(dataset Dataset) *AddDatasetResponse {
	this := AddDatasetResponse{}
	this.Dataset = dataset
	return &this
}

// NewAddDatasetResponseWithDefaults instantiates a new AddDatasetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDatasetResponseWithDefaults() *AddDatasetResponse {
	this := AddDatasetResponse{}
	return &this
}

// GetDataset returns the Dataset field value
func (o *AddDatasetResponse) GetDataset() Dataset {
	if o == nil {
		var ret Dataset
		return ret
	}

	return o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value
// and a boolean to check if the value has been set.
func (o *AddDatasetResponse) GetDatasetOk() (*Dataset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dataset, true
}

// SetDataset sets field value
func (o *AddDatasetResponse) SetDataset(v Dataset) {
	o.Dataset = v
}

func (o AddDatasetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddDatasetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataset"] = o.Dataset
	return toSerialize, nil
}

type NullableAddDatasetResponse struct {
	value *AddDatasetResponse
	isSet bool
}

func (v NullableAddDatasetResponse) Get() *AddDatasetResponse {
	return v.value
}

func (v *NullableAddDatasetResponse) Set(val *AddDatasetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDatasetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDatasetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDatasetResponse(val *AddDatasetResponse) *NullableAddDatasetResponse {
	return &NullableAddDatasetResponse{value: val, isSet: true}
}

func (v NullableAddDatasetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDatasetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
