/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.445
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SessionTestResultNotFound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionTestResultNotFound{}

// SessionTestResultNotFound struct for SessionTestResultNotFound
type SessionTestResultNotFound struct {
	SessionRunId string `json:"sessionRunId"`
	QueryStatus  string `json:"queryStatus"`
}

// NewSessionTestResultNotFound instantiates a new SessionTestResultNotFound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionTestResultNotFound(sessionRunId string, queryStatus string) *SessionTestResultNotFound {
	this := SessionTestResultNotFound{}
	this.SessionRunId = sessionRunId
	this.QueryStatus = queryStatus
	return &this
}

// NewSessionTestResultNotFoundWithDefaults instantiates a new SessionTestResultNotFound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionTestResultNotFoundWithDefaults() *SessionTestResultNotFound {
	this := SessionTestResultNotFound{}
	return &this
}

// GetSessionRunId returns the SessionRunId field value
func (o *SessionTestResultNotFound) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *SessionTestResultNotFound) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *SessionTestResultNotFound) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetQueryStatus returns the QueryStatus field value
func (o *SessionTestResultNotFound) GetQueryStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryStatus
}

// GetQueryStatusOk returns a tuple with the QueryStatus field value
// and a boolean to check if the value has been set.
func (o *SessionTestResultNotFound) GetQueryStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryStatus, true
}

// SetQueryStatus sets field value
func (o *SessionTestResultNotFound) SetQueryStatus(v string) {
	o.QueryStatus = v
}

func (o SessionTestResultNotFound) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionTestResultNotFound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["queryStatus"] = o.QueryStatus
	return toSerialize, nil
}

type NullableSessionTestResultNotFound struct {
	value *SessionTestResultNotFound
	isSet bool
}

func (v NullableSessionTestResultNotFound) Get() *SessionTestResultNotFound {
	return v.value
}

func (v *NullableSessionTestResultNotFound) Set(val *SessionTestResultNotFound) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionTestResultNotFound) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionTestResultNotFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionTestResultNotFound(val *SessionTestResultNotFound) *NullableSessionTestResultNotFound {
	return &NullableSessionTestResultNotFound{value: val, isSet: true}
}

func (v NullableSessionTestResultNotFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionTestResultNotFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
