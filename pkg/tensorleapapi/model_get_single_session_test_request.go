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

// checks if the GetSingleSessionTestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSingleSessionTestRequest{}

// GetSingleSessionTestRequest struct for GetSingleSessionTestRequest
type GetSingleSessionTestRequest struct {
	Cid       string `json:"cid"`
	ProjectId string `json:"projectId"`
}

// NewGetSingleSessionTestRequest instantiates a new GetSingleSessionTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSingleSessionTestRequest(cid string, projectId string) *GetSingleSessionTestRequest {
	this := GetSingleSessionTestRequest{}
	this.Cid = cid
	this.ProjectId = projectId
	return &this
}

// NewGetSingleSessionTestRequestWithDefaults instantiates a new GetSingleSessionTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSingleSessionTestRequestWithDefaults() *GetSingleSessionTestRequest {
	this := GetSingleSessionTestRequest{}
	return &this
}

// GetCid returns the Cid field value
func (o *GetSingleSessionTestRequest) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *GetSingleSessionTestRequest) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *GetSingleSessionTestRequest) SetCid(v string) {
	o.Cid = v
}

// GetProjectId returns the ProjectId field value
func (o *GetSingleSessionTestRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetSingleSessionTestRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetSingleSessionTestRequest) SetProjectId(v string) {
	o.ProjectId = v
}

func (o GetSingleSessionTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSingleSessionTestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableGetSingleSessionTestRequest struct {
	value *GetSingleSessionTestRequest
	isSet bool
}

func (v NullableGetSingleSessionTestRequest) Get() *GetSingleSessionTestRequest {
	return v.value
}

func (v *NullableGetSingleSessionTestRequest) Set(val *GetSingleSessionTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSingleSessionTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSingleSessionTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSingleSessionTestRequest(val *GetSingleSessionTestRequest) *NullableGetSingleSessionTestRequest {
	return &NullableGetSingleSessionTestRequest{value: val, isSet: true}
}

func (v NullableGetSingleSessionTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSingleSessionTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
