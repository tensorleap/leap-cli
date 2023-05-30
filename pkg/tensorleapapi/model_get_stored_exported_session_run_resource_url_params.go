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

// checks if the GetStoredExportedSessionRunResourceUrlParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStoredExportedSessionRunResourceUrlParams{}

// GetStoredExportedSessionRunResourceUrlParams struct for GetStoredExportedSessionRunResourceUrlParams
type GetStoredExportedSessionRunResourceUrlParams struct {
	ExportedSessionRunId string `json:"exportedSessionRunId"`
}

// NewGetStoredExportedSessionRunResourceUrlParams instantiates a new GetStoredExportedSessionRunResourceUrlParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStoredExportedSessionRunResourceUrlParams(exportedSessionRunId string) *GetStoredExportedSessionRunResourceUrlParams {
	this := GetStoredExportedSessionRunResourceUrlParams{}
	this.ExportedSessionRunId = exportedSessionRunId
	return &this
}

// NewGetStoredExportedSessionRunResourceUrlParamsWithDefaults instantiates a new GetStoredExportedSessionRunResourceUrlParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStoredExportedSessionRunResourceUrlParamsWithDefaults() *GetStoredExportedSessionRunResourceUrlParams {
	this := GetStoredExportedSessionRunResourceUrlParams{}
	return &this
}

// GetExportedSessionRunId returns the ExportedSessionRunId field value
func (o *GetStoredExportedSessionRunResourceUrlParams) GetExportedSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExportedSessionRunId
}

// GetExportedSessionRunIdOk returns a tuple with the ExportedSessionRunId field value
// and a boolean to check if the value has been set.
func (o *GetStoredExportedSessionRunResourceUrlParams) GetExportedSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportedSessionRunId, true
}

// SetExportedSessionRunId sets field value
func (o *GetStoredExportedSessionRunResourceUrlParams) SetExportedSessionRunId(v string) {
	o.ExportedSessionRunId = v
}

func (o GetStoredExportedSessionRunResourceUrlParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStoredExportedSessionRunResourceUrlParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exportedSessionRunId"] = o.ExportedSessionRunId
	return toSerialize, nil
}

type NullableGetStoredExportedSessionRunResourceUrlParams struct {
	value *GetStoredExportedSessionRunResourceUrlParams
	isSet bool
}

func (v NullableGetStoredExportedSessionRunResourceUrlParams) Get() *GetStoredExportedSessionRunResourceUrlParams {
	return v.value
}

func (v *NullableGetStoredExportedSessionRunResourceUrlParams) Set(val *GetStoredExportedSessionRunResourceUrlParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStoredExportedSessionRunResourceUrlParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStoredExportedSessionRunResourceUrlParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStoredExportedSessionRunResourceUrlParams(val *GetStoredExportedSessionRunResourceUrlParams) *NullableGetStoredExportedSessionRunResourceUrlParams {
	return &NullableGetStoredExportedSessionRunResourceUrlParams{value: val, isSet: true}
}

func (v NullableGetStoredExportedSessionRunResourceUrlParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStoredExportedSessionRunResourceUrlParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
