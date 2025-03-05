/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetSessionsEpochsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSessionsEpochsRequest{}

// GetSessionsEpochsRequest struct for GetSessionsEpochsRequest
type GetSessionsEpochsRequest struct {
	ProjectId  string   `json:"projectId"`
	SessionIds []string `json:"sessionIds"`
}

// NewGetSessionsEpochsRequest instantiates a new GetSessionsEpochsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSessionsEpochsRequest(projectId string, sessionIds []string) *GetSessionsEpochsRequest {
	this := GetSessionsEpochsRequest{}
	this.ProjectId = projectId
	this.SessionIds = sessionIds
	return &this
}

// NewGetSessionsEpochsRequestWithDefaults instantiates a new GetSessionsEpochsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSessionsEpochsRequestWithDefaults() *GetSessionsEpochsRequest {
	this := GetSessionsEpochsRequest{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GetSessionsEpochsRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetSessionsEpochsRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetSessionsEpochsRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSessionIds returns the SessionIds field value
func (o *GetSessionsEpochsRequest) GetSessionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SessionIds
}

// GetSessionIdsOk returns a tuple with the SessionIds field value
// and a boolean to check if the value has been set.
func (o *GetSessionsEpochsRequest) GetSessionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionIds, true
}

// SetSessionIds sets field value
func (o *GetSessionsEpochsRequest) SetSessionIds(v []string) {
	o.SessionIds = v
}

func (o GetSessionsEpochsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSessionsEpochsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["sessionIds"] = o.SessionIds
	return toSerialize, nil
}

type NullableGetSessionsEpochsRequest struct {
	value *GetSessionsEpochsRequest
	isSet bool
}

func (v NullableGetSessionsEpochsRequest) Get() *GetSessionsEpochsRequest {
	return v.value
}

func (v *NullableGetSessionsEpochsRequest) Set(val *GetSessionsEpochsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSessionsEpochsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSessionsEpochsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSessionsEpochsRequest(val *GetSessionsEpochsRequest) *NullableGetSessionsEpochsRequest {
	return &NullableGetSessionsEpochsRequest{value: val, isSet: true}
}

func (v NullableGetSessionsEpochsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSessionsEpochsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
