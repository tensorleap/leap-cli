/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.301
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SessionTestResultError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionTestResultError{}

// SessionTestResultError struct for SessionTestResultError
type SessionTestResultError struct {
	SessionRunId string `json:"sessionRunId"`
	QueryStatus  string `json:"queryStatus"`
}

// NewSessionTestResultError instantiates a new SessionTestResultError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionTestResultError(sessionRunId string, queryStatus string) *SessionTestResultError {
	this := SessionTestResultError{}
	this.SessionRunId = sessionRunId
	this.QueryStatus = queryStatus
	return &this
}

// NewSessionTestResultErrorWithDefaults instantiates a new SessionTestResultError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionTestResultErrorWithDefaults() *SessionTestResultError {
	this := SessionTestResultError{}
	return &this
}

// GetSessionRunId returns the SessionRunId field value
func (o *SessionTestResultError) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *SessionTestResultError) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *SessionTestResultError) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetQueryStatus returns the QueryStatus field value
func (o *SessionTestResultError) GetQueryStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryStatus
}

// GetQueryStatusOk returns a tuple with the QueryStatus field value
// and a boolean to check if the value has been set.
func (o *SessionTestResultError) GetQueryStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryStatus, true
}

// SetQueryStatus sets field value
func (o *SessionTestResultError) SetQueryStatus(v string) {
	o.QueryStatus = v
}

func (o SessionTestResultError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionTestResultError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["queryStatus"] = o.QueryStatus
	return toSerialize, nil
}

type NullableSessionTestResultError struct {
	value *SessionTestResultError
	isSet bool
}

func (v NullableSessionTestResultError) Get() *SessionTestResultError {
	return v.value
}

func (v *NullableSessionTestResultError) Set(val *SessionTestResultError) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionTestResultError) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionTestResultError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionTestResultError(val *SessionTestResultError) *NullableSessionTestResultError {
	return &NullableSessionTestResultError{value: val, isSet: true}
}

func (v NullableSessionTestResultError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionTestResultError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
