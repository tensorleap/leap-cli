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

// checks if the GetDatasetVersionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDatasetVersionParams{}

// GetDatasetVersionParams struct for GetDatasetVersionParams
type GetDatasetVersionParams struct {
	DatasetVersionId string `json:"datasetVersionId"`
}

// NewGetDatasetVersionParams instantiates a new GetDatasetVersionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDatasetVersionParams(datasetVersionId string) *GetDatasetVersionParams {
	this := GetDatasetVersionParams{}
	this.DatasetVersionId = datasetVersionId
	return &this
}

// NewGetDatasetVersionParamsWithDefaults instantiates a new GetDatasetVersionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDatasetVersionParamsWithDefaults() *GetDatasetVersionParams {
	this := GetDatasetVersionParams{}
	return &this
}

// GetDatasetVersionId returns the DatasetVersionId field value
func (o *GetDatasetVersionParams) GetDatasetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasetVersionId
}

// GetDatasetVersionIdOk returns a tuple with the DatasetVersionId field value
// and a boolean to check if the value has been set.
func (o *GetDatasetVersionParams) GetDatasetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetVersionId, true
}

// SetDatasetVersionId sets field value
func (o *GetDatasetVersionParams) SetDatasetVersionId(v string) {
	o.DatasetVersionId = v
}

func (o GetDatasetVersionParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDatasetVersionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datasetVersionId"] = o.DatasetVersionId
	return toSerialize, nil
}

type NullableGetDatasetVersionParams struct {
	value *GetDatasetVersionParams
	isSet bool
}

func (v NullableGetDatasetVersionParams) Get() *GetDatasetVersionParams {
	return v.value
}

func (v *NullableGetDatasetVersionParams) Set(val *GetDatasetVersionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDatasetVersionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDatasetVersionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDatasetVersionParams(val *GetDatasetVersionParams) *NullableGetDatasetVersionParams {
	return &NullableGetDatasetVersionParams{value: val, isSet: true}
}

func (v NullableGetDatasetVersionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDatasetVersionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
