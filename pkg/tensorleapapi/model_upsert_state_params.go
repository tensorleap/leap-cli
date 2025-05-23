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

// checks if the UpsertStateParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsertStateParams{}

// UpsertStateParams struct for UpsertStateParams
type UpsertStateParams struct {
	ProjectId string `json:"projectId"`
	State     string `json:"state"`
}

// NewUpsertStateParams instantiates a new UpsertStateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertStateParams(projectId string, state string) *UpsertStateParams {
	this := UpsertStateParams{}
	this.ProjectId = projectId
	this.State = state
	return &this
}

// NewUpsertStateParamsWithDefaults instantiates a new UpsertStateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertStateParamsWithDefaults() *UpsertStateParams {
	this := UpsertStateParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *UpsertStateParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *UpsertStateParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *UpsertStateParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetState returns the State field value
func (o *UpsertStateParams) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *UpsertStateParams) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *UpsertStateParams) SetState(v string) {
	o.State = v
}

func (o UpsertStateParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsertStateParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["state"] = o.State
	return toSerialize, nil
}

type NullableUpsertStateParams struct {
	value *UpsertStateParams
	isSet bool
}

func (v NullableUpsertStateParams) Get() *UpsertStateParams {
	return v.value
}

func (v *NullableUpsertStateParams) Set(val *UpsertStateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertStateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertStateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertStateParams(val *UpsertStateParams) *NullableUpsertStateParams {
	return &NullableUpsertStateParams{value: val, isSet: true}
}

func (v NullableUpsertStateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertStateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
