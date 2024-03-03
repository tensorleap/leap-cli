/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ModifyDatasetVersionNoteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyDatasetVersionNoteResponse{}

// ModifyDatasetVersionNoteResponse struct for ModifyDatasetVersionNoteResponse
type ModifyDatasetVersionNoteResponse struct {
	DatasetVersion DatasetVersion `json:"datasetVersion"`
}

// NewModifyDatasetVersionNoteResponse instantiates a new ModifyDatasetVersionNoteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyDatasetVersionNoteResponse(datasetVersion DatasetVersion) *ModifyDatasetVersionNoteResponse {
	this := ModifyDatasetVersionNoteResponse{}
	this.DatasetVersion = datasetVersion
	return &this
}

// NewModifyDatasetVersionNoteResponseWithDefaults instantiates a new ModifyDatasetVersionNoteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyDatasetVersionNoteResponseWithDefaults() *ModifyDatasetVersionNoteResponse {
	this := ModifyDatasetVersionNoteResponse{}
	return &this
}

// GetDatasetVersion returns the DatasetVersion field value
func (o *ModifyDatasetVersionNoteResponse) GetDatasetVersion() DatasetVersion {
	if o == nil {
		var ret DatasetVersion
		return ret
	}

	return o.DatasetVersion
}

// GetDatasetVersionOk returns a tuple with the DatasetVersion field value
// and a boolean to check if the value has been set.
func (o *ModifyDatasetVersionNoteResponse) GetDatasetVersionOk() (*DatasetVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetVersion, true
}

// SetDatasetVersion sets field value
func (o *ModifyDatasetVersionNoteResponse) SetDatasetVersion(v DatasetVersion) {
	o.DatasetVersion = v
}

func (o ModifyDatasetVersionNoteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyDatasetVersionNoteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datasetVersion"] = o.DatasetVersion
	return toSerialize, nil
}

type NullableModifyDatasetVersionNoteResponse struct {
	value *ModifyDatasetVersionNoteResponse
	isSet bool
}

func (v NullableModifyDatasetVersionNoteResponse) Get() *ModifyDatasetVersionNoteResponse {
	return v.value
}

func (v *NullableModifyDatasetVersionNoteResponse) Set(val *ModifyDatasetVersionNoteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyDatasetVersionNoteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyDatasetVersionNoteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyDatasetVersionNoteResponse(val *ModifyDatasetVersionNoteResponse) *NullableModifyDatasetVersionNoteResponse {
	return &NullableModifyDatasetVersionNoteResponse{value: val, isSet: true}
}

func (v NullableModifyDatasetVersionNoteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyDatasetVersionNoteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
