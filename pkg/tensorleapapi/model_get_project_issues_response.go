/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.378-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetProjectIssuesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectIssuesResponse{}

// GetProjectIssuesResponse struct for GetProjectIssuesResponse
type GetProjectIssuesResponse struct {
	Issues []Issue `json:"issues"`
}

// NewGetProjectIssuesResponse instantiates a new GetProjectIssuesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectIssuesResponse(issues []Issue) *GetProjectIssuesResponse {
	this := GetProjectIssuesResponse{}
	this.Issues = issues
	return &this
}

// NewGetProjectIssuesResponseWithDefaults instantiates a new GetProjectIssuesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectIssuesResponseWithDefaults() *GetProjectIssuesResponse {
	this := GetProjectIssuesResponse{}
	return &this
}

// GetIssues returns the Issues field value
func (o *GetProjectIssuesResponse) GetIssues() []Issue {
	if o == nil {
		var ret []Issue
		return ret
	}

	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value
// and a boolean to check if the value has been set.
func (o *GetProjectIssuesResponse) GetIssuesOk() ([]Issue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issues, true
}

// SetIssues sets field value
func (o *GetProjectIssuesResponse) SetIssues(v []Issue) {
	o.Issues = v
}

func (o GetProjectIssuesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectIssuesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["issues"] = o.Issues
	return toSerialize, nil
}

type NullableGetProjectIssuesResponse struct {
	value *GetProjectIssuesResponse
	isSet bool
}

func (v NullableGetProjectIssuesResponse) Get() *GetProjectIssuesResponse {
	return v.value
}

func (v *NullableGetProjectIssuesResponse) Set(val *GetProjectIssuesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectIssuesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectIssuesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectIssuesResponse(val *GetProjectIssuesResponse) *NullableGetProjectIssuesResponse {
	return &NullableGetProjectIssuesResponse{value: val, isSet: true}
}

func (v NullableGetProjectIssuesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectIssuesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


