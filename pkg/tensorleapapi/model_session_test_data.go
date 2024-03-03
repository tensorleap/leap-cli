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

// checks if the SessionTestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionTestData{}

// SessionTestData struct for SessionTestData
type SessionTestData struct {
	SessionRunId string  `json:"sessionRunId"`
	ProjectId    string  `json:"projectId"`
	Epoch        float64 `json:"epoch"`
}

// NewSessionTestData instantiates a new SessionTestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionTestData(sessionRunId string, projectId string, epoch float64) *SessionTestData {
	this := SessionTestData{}
	this.SessionRunId = sessionRunId
	this.ProjectId = projectId
	this.Epoch = epoch
	return &this
}

// NewSessionTestDataWithDefaults instantiates a new SessionTestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionTestDataWithDefaults() *SessionTestData {
	this := SessionTestData{}
	return &this
}

// GetSessionRunId returns the SessionRunId field value
func (o *SessionTestData) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *SessionTestData) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *SessionTestData) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetProjectId returns the ProjectId field value
func (o *SessionTestData) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *SessionTestData) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *SessionTestData) SetProjectId(v string) {
	o.ProjectId = v
}

// GetEpoch returns the Epoch field value
func (o *SessionTestData) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *SessionTestData) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *SessionTestData) SetEpoch(v float64) {
	o.Epoch = v
}

func (o SessionTestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionTestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["projectId"] = o.ProjectId
	toSerialize["epoch"] = o.Epoch
	return toSerialize, nil
}

type NullableSessionTestData struct {
	value *SessionTestData
	isSet bool
}

func (v NullableSessionTestData) Get() *SessionTestData {
	return v.value
}

func (v *NullableSessionTestData) Set(val *SessionTestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionTestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionTestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionTestData(val *SessionTestData) *NullableSessionTestData {
	return &NullableSessionTestData{value: val, isSet: true}
}

func (v NullableSessionTestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionTestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
