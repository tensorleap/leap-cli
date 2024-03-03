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

// checks if the ExportProjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportProjectResponse{}

// ExportProjectResponse struct for ExportProjectResponse
type ExportProjectResponse struct {
	JobId string `json:"jobId"`
}

// NewExportProjectResponse instantiates a new ExportProjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportProjectResponse(jobId string) *ExportProjectResponse {
	this := ExportProjectResponse{}
	this.JobId = jobId
	return &this
}

// NewExportProjectResponseWithDefaults instantiates a new ExportProjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportProjectResponseWithDefaults() *ExportProjectResponse {
	this := ExportProjectResponse{}
	return &this
}

// GetJobId returns the JobId field value
func (o *ExportProjectResponse) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *ExportProjectResponse) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *ExportProjectResponse) SetJobId(v string) {
	o.JobId = v
}

func (o ExportProjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportProjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobId"] = o.JobId
	return toSerialize, nil
}

type NullableExportProjectResponse struct {
	value *ExportProjectResponse
	isSet bool
}

func (v NullableExportProjectResponse) Get() *ExportProjectResponse {
	return v.value
}

func (v *NullableExportProjectResponse) Set(val *ExportProjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExportProjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExportProjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportProjectResponse(val *ExportProjectResponse) *NullableExportProjectResponse {
	return &NullableExportProjectResponse{value: val, isSet: true}
}

func (v NullableExportProjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportProjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
