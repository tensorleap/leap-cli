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

// checks if the UpdateSessionRunNameParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSessionRunNameParams{}

// UpdateSessionRunNameParams struct for UpdateSessionRunNameParams
type UpdateSessionRunNameParams struct {
	Cid       string `json:"cid"`
	Name      string `json:"name"`
	ProjectId string `json:"projectId"`
}

// NewUpdateSessionRunNameParams instantiates a new UpdateSessionRunNameParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSessionRunNameParams(cid string, name string, projectId string) *UpdateSessionRunNameParams {
	this := UpdateSessionRunNameParams{}
	this.Cid = cid
	this.Name = name
	this.ProjectId = projectId
	return &this
}

// NewUpdateSessionRunNameParamsWithDefaults instantiates a new UpdateSessionRunNameParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSessionRunNameParamsWithDefaults() *UpdateSessionRunNameParams {
	this := UpdateSessionRunNameParams{}
	return &this
}

// GetCid returns the Cid field value
func (o *UpdateSessionRunNameParams) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *UpdateSessionRunNameParams) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *UpdateSessionRunNameParams) SetCid(v string) {
	o.Cid = v
}

// GetName returns the Name field value
func (o *UpdateSessionRunNameParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateSessionRunNameParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateSessionRunNameParams) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value
func (o *UpdateSessionRunNameParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *UpdateSessionRunNameParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *UpdateSessionRunNameParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o UpdateSessionRunNameParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSessionRunNameParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["name"] = o.Name
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableUpdateSessionRunNameParams struct {
	value *UpdateSessionRunNameParams
	isSet bool
}

func (v NullableUpdateSessionRunNameParams) Get() *UpdateSessionRunNameParams {
	return v.value
}

func (v *NullableUpdateSessionRunNameParams) Set(val *UpdateSessionRunNameParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSessionRunNameParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSessionRunNameParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSessionRunNameParams(val *UpdateSessionRunNameParams) *NullableUpdateSessionRunNameParams {
	return &NullableUpdateSessionRunNameParams{value: val, isSet: true}
}

func (v NullableUpdateSessionRunNameParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSessionRunNameParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
