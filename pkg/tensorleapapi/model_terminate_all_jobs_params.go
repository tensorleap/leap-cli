/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the TerminateAllJobsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminateAllJobsParams{}

// TerminateAllJobsParams struct for TerminateAllJobsParams
type TerminateAllJobsParams struct {
	ProjectId *string `json:"projectId,omitempty"`
}

// NewTerminateAllJobsParams instantiates a new TerminateAllJobsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminateAllJobsParams() *TerminateAllJobsParams {
	this := TerminateAllJobsParams{}
	return &this
}

// NewTerminateAllJobsParamsWithDefaults instantiates a new TerminateAllJobsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminateAllJobsParamsWithDefaults() *TerminateAllJobsParams {
	this := TerminateAllJobsParams{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *TerminateAllJobsParams) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminateAllJobsParams) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *TerminateAllJobsParams) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *TerminateAllJobsParams) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o TerminateAllJobsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminateAllJobsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableTerminateAllJobsParams struct {
	value *TerminateAllJobsParams
	isSet bool
}

func (v NullableTerminateAllJobsParams) Get() *TerminateAllJobsParams {
	return v.value
}

func (v *NullableTerminateAllJobsParams) Set(val *TerminateAllJobsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminateAllJobsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminateAllJobsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminateAllJobsParams(val *TerminateAllJobsParams) *NullableTerminateAllJobsParams {
	return &NullableTerminateAllJobsParams{value: val, isSet: true}
}

func (v NullableTerminateAllJobsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminateAllJobsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
