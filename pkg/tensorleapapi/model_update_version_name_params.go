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

// checks if the UpdateVersionNameParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVersionNameParams{}

// UpdateVersionNameParams struct for UpdateVersionNameParams
type UpdateVersionNameParams struct {
	Cid       string `json:"cid"`
	ProjectId string `json:"projectId"`
	Name      string `json:"name"`
}

// NewUpdateVersionNameParams instantiates a new UpdateVersionNameParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVersionNameParams(cid string, projectId string, name string) *UpdateVersionNameParams {
	this := UpdateVersionNameParams{}
	this.Cid = cid
	this.ProjectId = projectId
	this.Name = name
	return &this
}

// NewUpdateVersionNameParamsWithDefaults instantiates a new UpdateVersionNameParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVersionNameParamsWithDefaults() *UpdateVersionNameParams {
	this := UpdateVersionNameParams{}
	return &this
}

// GetCid returns the Cid field value
func (o *UpdateVersionNameParams) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *UpdateVersionNameParams) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *UpdateVersionNameParams) SetCid(v string) {
	o.Cid = v
}

// GetProjectId returns the ProjectId field value
func (o *UpdateVersionNameParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *UpdateVersionNameParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *UpdateVersionNameParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *UpdateVersionNameParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateVersionNameParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateVersionNameParams) SetName(v string) {
	o.Name = v
}

func (o UpdateVersionNameParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVersionNameParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableUpdateVersionNameParams struct {
	value *UpdateVersionNameParams
	isSet bool
}

func (v NullableUpdateVersionNameParams) Get() *UpdateVersionNameParams {
	return v.value
}

func (v *NullableUpdateVersionNameParams) Set(val *UpdateVersionNameParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVersionNameParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVersionNameParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVersionNameParams(val *UpdateVersionNameParams) *NullableUpdateVersionNameParams {
	return &NullableUpdateVersionNameParams{value: val, isSet: true}
}

func (v NullableUpdateVersionNameParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVersionNameParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
