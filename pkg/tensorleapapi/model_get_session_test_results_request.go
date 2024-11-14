/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.557
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetSessionTestResultsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSessionTestResultsRequest{}

// GetSessionTestResultsRequest struct for GetSessionTestResultsRequest
type GetSessionTestResultsRequest struct {
	ProjectId        string            `json:"projectId"`
	SessionsTestData []SessionTestData `json:"sessionsTestData"`
}

// NewGetSessionTestResultsRequest instantiates a new GetSessionTestResultsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSessionTestResultsRequest(projectId string, sessionsTestData []SessionTestData) *GetSessionTestResultsRequest {
	this := GetSessionTestResultsRequest{}
	this.ProjectId = projectId
	this.SessionsTestData = sessionsTestData
	return &this
}

// NewGetSessionTestResultsRequestWithDefaults instantiates a new GetSessionTestResultsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSessionTestResultsRequestWithDefaults() *GetSessionTestResultsRequest {
	this := GetSessionTestResultsRequest{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GetSessionTestResultsRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetSessionTestResultsRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetSessionTestResultsRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSessionsTestData returns the SessionsTestData field value
func (o *GetSessionTestResultsRequest) GetSessionsTestData() []SessionTestData {
	if o == nil {
		var ret []SessionTestData
		return ret
	}

	return o.SessionsTestData
}

// GetSessionsTestDataOk returns a tuple with the SessionsTestData field value
// and a boolean to check if the value has been set.
func (o *GetSessionTestResultsRequest) GetSessionsTestDataOk() ([]SessionTestData, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionsTestData, true
}

// SetSessionsTestData sets field value
func (o *GetSessionTestResultsRequest) SetSessionsTestData(v []SessionTestData) {
	o.SessionsTestData = v
}

func (o GetSessionTestResultsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSessionTestResultsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["sessionsTestData"] = o.SessionsTestData
	return toSerialize, nil
}

type NullableGetSessionTestResultsRequest struct {
	value *GetSessionTestResultsRequest
	isSet bool
}

func (v NullableGetSessionTestResultsRequest) Get() *GetSessionTestResultsRequest {
	return v.value
}

func (v *NullableGetSessionTestResultsRequest) Set(val *GetSessionTestResultsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSessionTestResultsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSessionTestResultsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSessionTestResultsRequest(val *GetSessionTestResultsRequest) *NullableGetSessionTestResultsRequest {
	return &NullableGetSessionTestResultsRequest{value: val, isSet: true}
}

func (v NullableGetSessionTestResultsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSessionTestResultsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
