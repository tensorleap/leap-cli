/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.285
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetDatasetVersionsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDatasetVersionsParams{}

// GetDatasetVersionsParams struct for GetDatasetVersionsParams
type GetDatasetVersionsParams struct {
	DatasetId string `json:"datasetId"`
}

// NewGetDatasetVersionsParams instantiates a new GetDatasetVersionsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDatasetVersionsParams(datasetId string) *GetDatasetVersionsParams {
	this := GetDatasetVersionsParams{}
	this.DatasetId = datasetId
	return &this
}

// NewGetDatasetVersionsParamsWithDefaults instantiates a new GetDatasetVersionsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDatasetVersionsParamsWithDefaults() *GetDatasetVersionsParams {
	this := GetDatasetVersionsParams{}
	return &this
}

// GetDatasetId returns the DatasetId field value
func (o *GetDatasetVersionsParams) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *GetDatasetVersionsParams) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value
func (o *GetDatasetVersionsParams) SetDatasetId(v string) {
	o.DatasetId = v
}

func (o GetDatasetVersionsParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDatasetVersionsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datasetId"] = o.DatasetId
	return toSerialize, nil
}

type NullableGetDatasetVersionsParams struct {
	value *GetDatasetVersionsParams
	isSet bool
}

func (v NullableGetDatasetVersionsParams) Get() *GetDatasetVersionsParams {
	return v.value
}

func (v *NullableGetDatasetVersionsParams) Set(val *GetDatasetVersionsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDatasetVersionsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDatasetVersionsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDatasetVersionsParams(val *GetDatasetVersionsParams) *NullableGetDatasetVersionsParams {
	return &NullableGetDatasetVersionsParams{value: val, isSet: true}
}

func (v NullableGetDatasetVersionsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDatasetVersionsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


