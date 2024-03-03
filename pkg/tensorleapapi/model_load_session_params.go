/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the LoadSessionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadSessionParams{}

// LoadSessionParams struct for LoadSessionParams
type LoadSessionParams struct {
	SessionId string `json:"sessionId"`
	ProjectId string `json:"projectId"`
}

// NewLoadSessionParams instantiates a new LoadSessionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadSessionParams(sessionId string, projectId string) *LoadSessionParams {
	this := LoadSessionParams{}
	this.SessionId = sessionId
	this.ProjectId = projectId
	return &this
}

// NewLoadSessionParamsWithDefaults instantiates a new LoadSessionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadSessionParamsWithDefaults() *LoadSessionParams {
	this := LoadSessionParams{}
	return &this
}

// GetSessionId returns the SessionId field value
func (o *LoadSessionParams) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *LoadSessionParams) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *LoadSessionParams) SetSessionId(v string) {
	o.SessionId = v
}

// GetProjectId returns the ProjectId field value
func (o *LoadSessionParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *LoadSessionParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *LoadSessionParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o LoadSessionParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadSessionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionId"] = o.SessionId
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableLoadSessionParams struct {
	value *LoadSessionParams
	isSet bool
}

func (v NullableLoadSessionParams) Get() *LoadSessionParams {
	return v.value
}

func (v *NullableLoadSessionParams) Set(val *LoadSessionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadSessionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadSessionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadSessionParams(val *LoadSessionParams) *NullableLoadSessionParams {
	return &NullableLoadSessionParams{value: val, isSet: true}
}

func (v NullableLoadSessionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadSessionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
