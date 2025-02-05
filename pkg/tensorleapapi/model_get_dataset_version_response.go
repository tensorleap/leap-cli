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

// checks if the GetDatasetVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDatasetVersionResponse{}

// GetDatasetVersionResponse struct for GetDatasetVersionResponse
type GetDatasetVersionResponse struct {
	DatasetVersion DatasetVersion `json:"datasetVersion"`
}

// NewGetDatasetVersionResponse instantiates a new GetDatasetVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDatasetVersionResponse(datasetVersion DatasetVersion) *GetDatasetVersionResponse {
	this := GetDatasetVersionResponse{}
	this.DatasetVersion = datasetVersion
	return &this
}

// NewGetDatasetVersionResponseWithDefaults instantiates a new GetDatasetVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDatasetVersionResponseWithDefaults() *GetDatasetVersionResponse {
	this := GetDatasetVersionResponse{}
	return &this
}

// GetDatasetVersion returns the DatasetVersion field value
func (o *GetDatasetVersionResponse) GetDatasetVersion() DatasetVersion {
	if o == nil {
		var ret DatasetVersion
		return ret
	}

	return o.DatasetVersion
}

// GetDatasetVersionOk returns a tuple with the DatasetVersion field value
// and a boolean to check if the value has been set.
func (o *GetDatasetVersionResponse) GetDatasetVersionOk() (*DatasetVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetVersion, true
}

// SetDatasetVersion sets field value
func (o *GetDatasetVersionResponse) SetDatasetVersion(v DatasetVersion) {
	o.DatasetVersion = v
}

func (o GetDatasetVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDatasetVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datasetVersion"] = o.DatasetVersion
	return toSerialize, nil
}

type NullableGetDatasetVersionResponse struct {
	value *GetDatasetVersionResponse
	isSet bool
}

func (v NullableGetDatasetVersionResponse) Get() *GetDatasetVersionResponse {
	return v.value
}

func (v *NullableGetDatasetVersionResponse) Set(val *GetDatasetVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDatasetVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDatasetVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDatasetVersionResponse(val *GetDatasetVersionResponse) *NullableGetDatasetVersionResponse {
	return &NullableGetDatasetVersionResponse{value: val, isSet: true}
}

func (v NullableGetDatasetVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDatasetVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
