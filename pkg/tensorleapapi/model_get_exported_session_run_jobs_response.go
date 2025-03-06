/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetExportedSessionRunJobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetExportedSessionRunJobsResponse{}

// GetExportedSessionRunJobsResponse struct for GetExportedSessionRunJobsResponse
type GetExportedSessionRunJobsResponse struct {
	ExportedModelData []ExportedModelData `json:"exportedModelData"`
}

// NewGetExportedSessionRunJobsResponse instantiates a new GetExportedSessionRunJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExportedSessionRunJobsResponse(exportedModelData []ExportedModelData) *GetExportedSessionRunJobsResponse {
	this := GetExportedSessionRunJobsResponse{}
	this.ExportedModelData = exportedModelData
	return &this
}

// NewGetExportedSessionRunJobsResponseWithDefaults instantiates a new GetExportedSessionRunJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExportedSessionRunJobsResponseWithDefaults() *GetExportedSessionRunJobsResponse {
	this := GetExportedSessionRunJobsResponse{}
	return &this
}

// GetExportedModelData returns the ExportedModelData field value
func (o *GetExportedSessionRunJobsResponse) GetExportedModelData() []ExportedModelData {
	if o == nil {
		var ret []ExportedModelData
		return ret
	}

	return o.ExportedModelData
}

// GetExportedModelDataOk returns a tuple with the ExportedModelData field value
// and a boolean to check if the value has been set.
func (o *GetExportedSessionRunJobsResponse) GetExportedModelDataOk() ([]ExportedModelData, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExportedModelData, true
}

// SetExportedModelData sets field value
func (o *GetExportedSessionRunJobsResponse) SetExportedModelData(v []ExportedModelData) {
	o.ExportedModelData = v
}

func (o GetExportedSessionRunJobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExportedSessionRunJobsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exportedModelData"] = o.ExportedModelData
	return toSerialize, nil
}

type NullableGetExportedSessionRunJobsResponse struct {
	value *GetExportedSessionRunJobsResponse
	isSet bool
}

func (v NullableGetExportedSessionRunJobsResponse) Get() *GetExportedSessionRunJobsResponse {
	return v.value
}

func (v *NullableGetExportedSessionRunJobsResponse) Set(val *GetExportedSessionRunJobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExportedSessionRunJobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExportedSessionRunJobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExportedSessionRunJobsResponse(val *GetExportedSessionRunJobsResponse) *NullableGetExportedSessionRunJobsResponse {
	return &NullableGetExportedSessionRunJobsResponse{value: val, isSet: true}
}

func (v NullableGetExportedSessionRunJobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExportedSessionRunJobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
