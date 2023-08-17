/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetJobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetJobsResponse{}

// GetJobsResponse struct for GetJobsResponse
type GetJobsResponse struct {
	Processes []RunProcess `json:"processes"`
}

// NewGetJobsResponse instantiates a new GetJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetJobsResponse(processes []RunProcess) *GetJobsResponse {
	this := GetJobsResponse{}
	this.Processes = processes
	return &this
}

// NewGetJobsResponseWithDefaults instantiates a new GetJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetJobsResponseWithDefaults() *GetJobsResponse {
	this := GetJobsResponse{}
	return &this
}

// GetProcesses returns the Processes field value
func (o *GetJobsResponse) GetProcesses() []RunProcess {
	if o == nil {
		var ret []RunProcess
		return ret
	}

	return o.Processes
}

// GetProcessesOk returns a tuple with the Processes field value
// and a boolean to check if the value has been set.
func (o *GetJobsResponse) GetProcessesOk() ([]RunProcess, bool) {
	if o == nil {
		return nil, false
	}
	return o.Processes, true
}

// SetProcesses sets field value
func (o *GetJobsResponse) SetProcesses(v []RunProcess) {
	o.Processes = v
}

func (o GetJobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetJobsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["processes"] = o.Processes
	return toSerialize, nil
}

type NullableGetJobsResponse struct {
	value *GetJobsResponse
	isSet bool
}

func (v NullableGetJobsResponse) Get() *GetJobsResponse {
	return v.value
}

func (v *NullableGetJobsResponse) Set(val *GetJobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetJobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetJobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetJobsResponse(val *GetJobsResponse) *NullableGetJobsResponse {
	return &NullableGetJobsResponse{value: val, isSet: true}
}

func (v NullableGetJobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetJobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


