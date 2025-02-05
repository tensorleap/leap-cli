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

// checks if the GetDatasetVersionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDatasetVersionsResponse{}

// GetDatasetVersionsResponse struct for GetDatasetVersionsResponse
type GetDatasetVersionsResponse struct {
	DatasetVersions []DatasetVersion `json:"datasetVersions"`
}

// NewGetDatasetVersionsResponse instantiates a new GetDatasetVersionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDatasetVersionsResponse(datasetVersions []DatasetVersion) *GetDatasetVersionsResponse {
	this := GetDatasetVersionsResponse{}
	this.DatasetVersions = datasetVersions
	return &this
}

// NewGetDatasetVersionsResponseWithDefaults instantiates a new GetDatasetVersionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDatasetVersionsResponseWithDefaults() *GetDatasetVersionsResponse {
	this := GetDatasetVersionsResponse{}
	return &this
}

// GetDatasetVersions returns the DatasetVersions field value
func (o *GetDatasetVersionsResponse) GetDatasetVersions() []DatasetVersion {
	if o == nil {
		var ret []DatasetVersion
		return ret
	}

	return o.DatasetVersions
}

// GetDatasetVersionsOk returns a tuple with the DatasetVersions field value
// and a boolean to check if the value has been set.
func (o *GetDatasetVersionsResponse) GetDatasetVersionsOk() ([]DatasetVersion, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatasetVersions, true
}

// SetDatasetVersions sets field value
func (o *GetDatasetVersionsResponse) SetDatasetVersions(v []DatasetVersion) {
	o.DatasetVersions = v
}

func (o GetDatasetVersionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDatasetVersionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datasetVersions"] = o.DatasetVersions
	return toSerialize, nil
}

type NullableGetDatasetVersionsResponse struct {
	value *GetDatasetVersionsResponse
	isSet bool
}

func (v NullableGetDatasetVersionsResponse) Get() *GetDatasetVersionsResponse {
	return v.value
}

func (v *NullableGetDatasetVersionsResponse) Set(val *GetDatasetVersionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDatasetVersionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDatasetVersionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDatasetVersionsResponse(val *GetDatasetVersionsResponse) *NullableGetDatasetVersionsResponse {
	return &NullableGetDatasetVersionsResponse{value: val, isSet: true}
}

func (v NullableGetDatasetVersionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDatasetVersionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
