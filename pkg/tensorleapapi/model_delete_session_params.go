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

// checks if the DeleteSessionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSessionParams{}

// DeleteSessionParams struct for DeleteSessionParams
type DeleteSessionParams struct {
	SessionId string `json:"sessionId"`
	ProjectId string `json:"projectId"`
}

// NewDeleteSessionParams instantiates a new DeleteSessionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSessionParams(sessionId string, projectId string) *DeleteSessionParams {
	this := DeleteSessionParams{}
	this.SessionId = sessionId
	this.ProjectId = projectId
	return &this
}

// NewDeleteSessionParamsWithDefaults instantiates a new DeleteSessionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSessionParamsWithDefaults() *DeleteSessionParams {
	this := DeleteSessionParams{}
	return &this
}

// GetSessionId returns the SessionId field value
func (o *DeleteSessionParams) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *DeleteSessionParams) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *DeleteSessionParams) SetSessionId(v string) {
	o.SessionId = v
}

// GetProjectId returns the ProjectId field value
func (o *DeleteSessionParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DeleteSessionParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DeleteSessionParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o DeleteSessionParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteSessionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionId"] = o.SessionId
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableDeleteSessionParams struct {
	value *DeleteSessionParams
	isSet bool
}

func (v NullableDeleteSessionParams) Get() *DeleteSessionParams {
	return v.value
}

func (v *NullableDeleteSessionParams) Set(val *DeleteSessionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSessionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSessionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSessionParams(val *DeleteSessionParams) *NullableDeleteSessionParams {
	return &NullableDeleteSessionParams{value: val, isSet: true}
}

func (v NullableDeleteSessionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSessionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
