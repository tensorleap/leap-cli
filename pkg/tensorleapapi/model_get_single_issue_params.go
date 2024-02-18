/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.456
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetSingleIssueParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSingleIssueParams{}

// GetSingleIssueParams struct for GetSingleIssueParams
type GetSingleIssueParams struct {
	Cid       string `json:"cid"`
	ProjectId string `json:"projectId"`
}

// NewGetSingleIssueParams instantiates a new GetSingleIssueParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSingleIssueParams(cid string, projectId string) *GetSingleIssueParams {
	this := GetSingleIssueParams{}
	this.Cid = cid
	this.ProjectId = projectId
	return &this
}

// NewGetSingleIssueParamsWithDefaults instantiates a new GetSingleIssueParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSingleIssueParamsWithDefaults() *GetSingleIssueParams {
	this := GetSingleIssueParams{}
	return &this
}

// GetCid returns the Cid field value
func (o *GetSingleIssueParams) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *GetSingleIssueParams) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *GetSingleIssueParams) SetCid(v string) {
	o.Cid = v
}

// GetProjectId returns the ProjectId field value
func (o *GetSingleIssueParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetSingleIssueParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetSingleIssueParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o GetSingleIssueParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSingleIssueParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableGetSingleIssueParams struct {
	value *GetSingleIssueParams
	isSet bool
}

func (v NullableGetSingleIssueParams) Get() *GetSingleIssueParams {
	return v.value
}

func (v *NullableGetSingleIssueParams) Set(val *GetSingleIssueParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSingleIssueParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSingleIssueParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSingleIssueParams(val *GetSingleIssueParams) *NullableGetSingleIssueParams {
	return &NullableGetSingleIssueParams{value: val, isSet: true}
}

func (v NullableGetSingleIssueParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSingleIssueParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
