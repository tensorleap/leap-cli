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

// checks if the DeleteIssueParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteIssueParams{}

// DeleteIssueParams struct for DeleteIssueParams
type DeleteIssueParams struct {
	Cid       string `json:"cid"`
	ProjectId string `json:"projectId"`
}

// NewDeleteIssueParams instantiates a new DeleteIssueParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteIssueParams(cid string, projectId string) *DeleteIssueParams {
	this := DeleteIssueParams{}
	this.Cid = cid
	this.ProjectId = projectId
	return &this
}

// NewDeleteIssueParamsWithDefaults instantiates a new DeleteIssueParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteIssueParamsWithDefaults() *DeleteIssueParams {
	this := DeleteIssueParams{}
	return &this
}

// GetCid returns the Cid field value
func (o *DeleteIssueParams) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *DeleteIssueParams) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *DeleteIssueParams) SetCid(v string) {
	o.Cid = v
}

// GetProjectId returns the ProjectId field value
func (o *DeleteIssueParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DeleteIssueParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DeleteIssueParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o DeleteIssueParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteIssueParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableDeleteIssueParams struct {
	value *DeleteIssueParams
	isSet bool
}

func (v NullableDeleteIssueParams) Get() *DeleteIssueParams {
	return v.value
}

func (v *NullableDeleteIssueParams) Set(val *DeleteIssueParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteIssueParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteIssueParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteIssueParams(val *DeleteIssueParams) *NullableDeleteIssueParams {
	return &NullableDeleteIssueParams{value: val, isSet: true}
}

func (v NullableDeleteIssueParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteIssueParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
