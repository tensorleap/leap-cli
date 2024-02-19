/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.464
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UpdateSessionNameParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSessionNameParams{}

// UpdateSessionNameParams struct for UpdateSessionNameParams
type UpdateSessionNameParams struct {
	Cid       string `json:"cid"`
	ProjectId string `json:"projectId"`
	Name      string `json:"name"`
}

// NewUpdateSessionNameParams instantiates a new UpdateSessionNameParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSessionNameParams(cid string, projectId string, name string) *UpdateSessionNameParams {
	this := UpdateSessionNameParams{}
	this.Cid = cid
	this.ProjectId = projectId
	this.Name = name
	return &this
}

// NewUpdateSessionNameParamsWithDefaults instantiates a new UpdateSessionNameParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSessionNameParamsWithDefaults() *UpdateSessionNameParams {
	this := UpdateSessionNameParams{}
	return &this
}

// GetCid returns the Cid field value
func (o *UpdateSessionNameParams) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *UpdateSessionNameParams) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *UpdateSessionNameParams) SetCid(v string) {
	o.Cid = v
}

// GetProjectId returns the ProjectId field value
func (o *UpdateSessionNameParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *UpdateSessionNameParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *UpdateSessionNameParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *UpdateSessionNameParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateSessionNameParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateSessionNameParams) SetName(v string) {
	o.Name = v
}

func (o UpdateSessionNameParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSessionNameParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableUpdateSessionNameParams struct {
	value *UpdateSessionNameParams
	isSet bool
}

func (v NullableUpdateSessionNameParams) Get() *UpdateSessionNameParams {
	return v.value
}

func (v *NullableUpdateSessionNameParams) Set(val *UpdateSessionNameParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSessionNameParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSessionNameParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSessionNameParams(val *UpdateSessionNameParams) *NullableUpdateSessionNameParams {
	return &NullableUpdateSessionNameParams{value: val, isSet: true}
}

func (v NullableUpdateSessionNameParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSessionNameParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
