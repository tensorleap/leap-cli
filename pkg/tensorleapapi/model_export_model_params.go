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

// checks if the ExportModelParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportModelParams{}

// ExportModelParams struct for ExportModelParams
type ExportModelParams struct {
	Type  ExportModelTypeEnum `json:"type"`
	Title string              `json:"title"`
	Epoch float64             `json:"epoch"`
}

// NewExportModelParams instantiates a new ExportModelParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportModelParams(type_ ExportModelTypeEnum, title string, epoch float64) *ExportModelParams {
	this := ExportModelParams{}
	this.Type = type_
	this.Title = title
	this.Epoch = epoch
	return &this
}

// NewExportModelParamsWithDefaults instantiates a new ExportModelParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportModelParamsWithDefaults() *ExportModelParams {
	this := ExportModelParams{}
	return &this
}

// GetType returns the Type field value
func (o *ExportModelParams) GetType() ExportModelTypeEnum {
	if o == nil {
		var ret ExportModelTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExportModelParams) GetTypeOk() (*ExportModelTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExportModelParams) SetType(v ExportModelTypeEnum) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *ExportModelParams) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ExportModelParams) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ExportModelParams) SetTitle(v string) {
	o.Title = v
}

// GetEpoch returns the Epoch field value
func (o *ExportModelParams) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *ExportModelParams) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *ExportModelParams) SetEpoch(v float64) {
	o.Epoch = v
}

func (o ExportModelParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportModelParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["epoch"] = o.Epoch
	return toSerialize, nil
}

type NullableExportModelParams struct {
	value *ExportModelParams
	isSet bool
}

func (v NullableExportModelParams) Get() *ExportModelParams {
	return v.value
}

func (v *NullableExportModelParams) Set(val *ExportModelParams) {
	v.value = val
	v.isSet = true
}

func (v NullableExportModelParams) IsSet() bool {
	return v.isSet
}

func (v *NullableExportModelParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportModelParams(val *ExportModelParams) *NullableExportModelParams {
	return &NullableExportModelParams{value: val, isSet: true}
}

func (v NullableExportModelParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportModelParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
