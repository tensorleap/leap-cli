/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.321
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ModifyDatasetVersionNoteParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyDatasetVersionNoteParams{}

// ModifyDatasetVersionNoteParams struct for ModifyDatasetVersionNoteParams
type ModifyDatasetVersionNoteParams struct {
	DatasetVersionId string `json:"datasetVersionId"`
	Note             string `json:"note"`
}

// NewModifyDatasetVersionNoteParams instantiates a new ModifyDatasetVersionNoteParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyDatasetVersionNoteParams(datasetVersionId string, note string) *ModifyDatasetVersionNoteParams {
	this := ModifyDatasetVersionNoteParams{}
	this.DatasetVersionId = datasetVersionId
	this.Note = note
	return &this
}

// NewModifyDatasetVersionNoteParamsWithDefaults instantiates a new ModifyDatasetVersionNoteParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyDatasetVersionNoteParamsWithDefaults() *ModifyDatasetVersionNoteParams {
	this := ModifyDatasetVersionNoteParams{}
	return &this
}

// GetDatasetVersionId returns the DatasetVersionId field value
func (o *ModifyDatasetVersionNoteParams) GetDatasetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasetVersionId
}

// GetDatasetVersionIdOk returns a tuple with the DatasetVersionId field value
// and a boolean to check if the value has been set.
func (o *ModifyDatasetVersionNoteParams) GetDatasetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetVersionId, true
}

// SetDatasetVersionId sets field value
func (o *ModifyDatasetVersionNoteParams) SetDatasetVersionId(v string) {
	o.DatasetVersionId = v
}

// GetNote returns the Note field value
func (o *ModifyDatasetVersionNoteParams) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *ModifyDatasetVersionNoteParams) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *ModifyDatasetVersionNoteParams) SetNote(v string) {
	o.Note = v
}

func (o ModifyDatasetVersionNoteParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyDatasetVersionNoteParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datasetVersionId"] = o.DatasetVersionId
	toSerialize["note"] = o.Note
	return toSerialize, nil
}

type NullableModifyDatasetVersionNoteParams struct {
	value *ModifyDatasetVersionNoteParams
	isSet bool
}

func (v NullableModifyDatasetVersionNoteParams) Get() *ModifyDatasetVersionNoteParams {
	return v.value
}

func (v *NullableModifyDatasetVersionNoteParams) Set(val *ModifyDatasetVersionNoteParams) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyDatasetVersionNoteParams) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyDatasetVersionNoteParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyDatasetVersionNoteParams(val *ModifyDatasetVersionNoteParams) *NullableModifyDatasetVersionNoteParams {
	return &NullableModifyDatasetVersionNoteParams{value: val, isSet: true}
}

func (v NullableModifyDatasetVersionNoteParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyDatasetVersionNoteParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
