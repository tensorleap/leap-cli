/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.417
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetSlimJobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSlimJobsResponse{}

// GetSlimJobsResponse struct for GetSlimJobsResponse
type GetSlimJobsResponse struct {
	Jobs []Job `json:"jobs"`
}

// NewGetSlimJobsResponse instantiates a new GetSlimJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSlimJobsResponse(jobs []Job) *GetSlimJobsResponse {
	this := GetSlimJobsResponse{}
	this.Jobs = jobs
	return &this
}

// NewGetSlimJobsResponseWithDefaults instantiates a new GetSlimJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSlimJobsResponseWithDefaults() *GetSlimJobsResponse {
	this := GetSlimJobsResponse{}
	return &this
}

// GetJobs returns the Jobs field value
func (o *GetSlimJobsResponse) GetJobs() []Job {
	if o == nil {
		var ret []Job
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *GetSlimJobsResponse) GetJobsOk() ([]Job, bool) {
	if o == nil {
		return nil, false
	}
	return o.Jobs, true
}

// SetJobs sets field value
func (o *GetSlimJobsResponse) SetJobs(v []Job) {
	o.Jobs = v
}

func (o GetSlimJobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSlimJobsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobs"] = o.Jobs
	return toSerialize, nil
}

type NullableGetSlimJobsResponse struct {
	value *GetSlimJobsResponse
	isSet bool
}

func (v NullableGetSlimJobsResponse) Get() *GetSlimJobsResponse {
	return v.value
}

func (v *NullableGetSlimJobsResponse) Set(val *GetSlimJobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSlimJobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSlimJobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSlimJobsResponse(val *GetSlimJobsResponse) *NullableGetSlimJobsResponse {
	return &NullableGetSlimJobsResponse{value: val, isSet: true}
}

func (v NullableGetSlimJobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSlimJobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
