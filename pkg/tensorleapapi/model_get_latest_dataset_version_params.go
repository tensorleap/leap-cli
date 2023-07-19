/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetLatestDatasetVersionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLatestDatasetVersionParams{}

// GetLatestDatasetVersionParams struct for GetLatestDatasetVersionParams
type GetLatestDatasetVersionParams struct {
	DatasetId string `json:"datasetId"`
}

// NewGetLatestDatasetVersionParams instantiates a new GetLatestDatasetVersionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLatestDatasetVersionParams(datasetId string) *GetLatestDatasetVersionParams {
	this := GetLatestDatasetVersionParams{}
	this.DatasetId = datasetId
	return &this
}

// NewGetLatestDatasetVersionParamsWithDefaults instantiates a new GetLatestDatasetVersionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLatestDatasetVersionParamsWithDefaults() *GetLatestDatasetVersionParams {
	this := GetLatestDatasetVersionParams{}
	return &this
}

// GetDatasetId returns the DatasetId field value
func (o *GetLatestDatasetVersionParams) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *GetLatestDatasetVersionParams) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value
func (o *GetLatestDatasetVersionParams) SetDatasetId(v string) {
	o.DatasetId = v
}

func (o GetLatestDatasetVersionParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLatestDatasetVersionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datasetId"] = o.DatasetId
	return toSerialize, nil
}

type NullableGetLatestDatasetVersionParams struct {
	value *GetLatestDatasetVersionParams
	isSet bool
}

func (v NullableGetLatestDatasetVersionParams) Get() *GetLatestDatasetVersionParams {
	return v.value
}

func (v *NullableGetLatestDatasetVersionParams) Set(val *GetLatestDatasetVersionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLatestDatasetVersionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLatestDatasetVersionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLatestDatasetVersionParams(val *GetLatestDatasetVersionParams) *NullableGetLatestDatasetVersionParams {
	return &NullableGetLatestDatasetVersionParams{value: val, isSet: true}
}

func (v NullableGetLatestDatasetVersionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLatestDatasetVersionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


