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

// checks if the DeleteSessionTestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSessionTestRequest{}

// DeleteSessionTestRequest struct for DeleteSessionTestRequest
type DeleteSessionTestRequest struct {
	Cid       string `json:"cid"`
	ProjectId string `json:"projectId"`
}

// NewDeleteSessionTestRequest instantiates a new DeleteSessionTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSessionTestRequest(cid string, projectId string) *DeleteSessionTestRequest {
	this := DeleteSessionTestRequest{}
	this.Cid = cid
	this.ProjectId = projectId
	return &this
}

// NewDeleteSessionTestRequestWithDefaults instantiates a new DeleteSessionTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSessionTestRequestWithDefaults() *DeleteSessionTestRequest {
	this := DeleteSessionTestRequest{}
	return &this
}

// GetCid returns the Cid field value
func (o *DeleteSessionTestRequest) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *DeleteSessionTestRequest) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *DeleteSessionTestRequest) SetCid(v string) {
	o.Cid = v
}

// GetProjectId returns the ProjectId field value
func (o *DeleteSessionTestRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DeleteSessionTestRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DeleteSessionTestRequest) SetProjectId(v string) {
	o.ProjectId = v
}

func (o DeleteSessionTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteSessionTestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableDeleteSessionTestRequest struct {
	value *DeleteSessionTestRequest
	isSet bool
}

func (v NullableDeleteSessionTestRequest) Get() *DeleteSessionTestRequest {
	return v.value
}

func (v *NullableDeleteSessionTestRequest) Set(val *DeleteSessionTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSessionTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSessionTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSessionTestRequest(val *DeleteSessionTestRequest) *NullableDeleteSessionTestRequest {
	return &NullableDeleteSessionTestRequest{value: val, isSet: true}
}

func (v NullableDeleteSessionTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSessionTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
