/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.301
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetDatasetsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDatasetsResponse{}

// GetDatasetsResponse struct for GetDatasetsResponse
type GetDatasetsResponse struct {
	Datasets []Dataset `json:"datasets"`
}

// NewGetDatasetsResponse instantiates a new GetDatasetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDatasetsResponse(datasets []Dataset) *GetDatasetsResponse {
	this := GetDatasetsResponse{}
	this.Datasets = datasets
	return &this
}

// NewGetDatasetsResponseWithDefaults instantiates a new GetDatasetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDatasetsResponseWithDefaults() *GetDatasetsResponse {
	this := GetDatasetsResponse{}
	return &this
}

// GetDatasets returns the Datasets field value
func (o *GetDatasetsResponse) GetDatasets() []Dataset {
	if o == nil {
		var ret []Dataset
		return ret
	}

	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value
// and a boolean to check if the value has been set.
func (o *GetDatasetsResponse) GetDatasetsOk() ([]Dataset, bool) {
	if o == nil {
		return nil, false
	}
	return o.Datasets, true
}

// SetDatasets sets field value
func (o *GetDatasetsResponse) SetDatasets(v []Dataset) {
	o.Datasets = v
}

func (o GetDatasetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDatasetsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datasets"] = o.Datasets
	return toSerialize, nil
}

type NullableGetDatasetsResponse struct {
	value *GetDatasetsResponse
	isSet bool
}

func (v NullableGetDatasetsResponse) Get() *GetDatasetsResponse {
	return v.value
}

func (v *NullableGetDatasetsResponse) Set(val *GetDatasetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDatasetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDatasetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDatasetsResponse(val *GetDatasetsResponse) *NullableGetDatasetsResponse {
	return &NullableGetDatasetsResponse{value: val, isSet: true}
}

func (v NullableGetDatasetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDatasetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
