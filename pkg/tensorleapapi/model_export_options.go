/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ExportOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportOptions{}

// ExportOptions struct for ExportOptions
type ExportOptions struct {
	ExcludeCalculatedFiles *bool `json:"excludeCalculatedFiles,omitempty"`
}

// NewExportOptions instantiates a new ExportOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportOptions() *ExportOptions {
	this := ExportOptions{}
	return &this
}

// NewExportOptionsWithDefaults instantiates a new ExportOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportOptionsWithDefaults() *ExportOptions {
	this := ExportOptions{}
	return &this
}

// GetExcludeCalculatedFiles returns the ExcludeCalculatedFiles field value if set, zero value otherwise.
func (o *ExportOptions) GetExcludeCalculatedFiles() bool {
	if o == nil || IsNil(o.ExcludeCalculatedFiles) {
		var ret bool
		return ret
	}
	return *o.ExcludeCalculatedFiles
}

// GetExcludeCalculatedFilesOk returns a tuple with the ExcludeCalculatedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportOptions) GetExcludeCalculatedFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeCalculatedFiles) {
		return nil, false
	}
	return o.ExcludeCalculatedFiles, true
}

// HasExcludeCalculatedFiles returns a boolean if a field has been set.
func (o *ExportOptions) HasExcludeCalculatedFiles() bool {
	if o != nil && !IsNil(o.ExcludeCalculatedFiles) {
		return true
	}

	return false
}

// SetExcludeCalculatedFiles gets a reference to the given bool and assigns it to the ExcludeCalculatedFiles field.
func (o *ExportOptions) SetExcludeCalculatedFiles(v bool) {
	o.ExcludeCalculatedFiles = &v
}

func (o ExportOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExcludeCalculatedFiles) {
		toSerialize["excludeCalculatedFiles"] = o.ExcludeCalculatedFiles
	}
	return toSerialize, nil
}

type NullableExportOptions struct {
	value *ExportOptions
	isSet bool
}

func (v NullableExportOptions) Get() *ExportOptions {
	return v.value
}

func (v *NullableExportOptions) Set(val *ExportOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableExportOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableExportOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportOptions(val *ExportOptions) *NullableExportOptions {
	return &NullableExportOptions{value: val, isSet: true}
}

func (v NullableExportOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}