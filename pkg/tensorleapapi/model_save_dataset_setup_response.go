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

// checks if the SaveDatasetSetupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveDatasetSetupResponse{}

// SaveDatasetSetupResponse struct for SaveDatasetSetupResponse
type SaveDatasetSetupResponse struct {
	Dataset Dataset `json:"dataset"`
}

// NewSaveDatasetSetupResponse instantiates a new SaveDatasetSetupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveDatasetSetupResponse(dataset Dataset) *SaveDatasetSetupResponse {
	this := SaveDatasetSetupResponse{}
	this.Dataset = dataset
	return &this
}

// NewSaveDatasetSetupResponseWithDefaults instantiates a new SaveDatasetSetupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveDatasetSetupResponseWithDefaults() *SaveDatasetSetupResponse {
	this := SaveDatasetSetupResponse{}
	return &this
}

// GetDataset returns the Dataset field value
func (o *SaveDatasetSetupResponse) GetDataset() Dataset {
	if o == nil {
		var ret Dataset
		return ret
	}

	return o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value
// and a boolean to check if the value has been set.
func (o *SaveDatasetSetupResponse) GetDatasetOk() (*Dataset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dataset, true
}

// SetDataset sets field value
func (o *SaveDatasetSetupResponse) SetDataset(v Dataset) {
	o.Dataset = v
}

func (o SaveDatasetSetupResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveDatasetSetupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataset"] = o.Dataset
	return toSerialize, nil
}

type NullableSaveDatasetSetupResponse struct {
	value *SaveDatasetSetupResponse
	isSet bool
}

func (v NullableSaveDatasetSetupResponse) Get() *SaveDatasetSetupResponse {
	return v.value
}

func (v *NullableSaveDatasetSetupResponse) Set(val *SaveDatasetSetupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveDatasetSetupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveDatasetSetupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveDatasetSetupResponse(val *SaveDatasetSetupResponse) *NullableSaveDatasetSetupResponse {
	return &NullableSaveDatasetSetupResponse{value: val, isSet: true}
}

func (v NullableSaveDatasetSetupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveDatasetSetupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
