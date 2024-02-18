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

// checks if the SessionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionsResponse{}

// SessionsResponse struct for SessionsResponse
type SessionsResponse struct {
	Sessions []Session `json:"sessions"`
}

// NewSessionsResponse instantiates a new SessionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionsResponse(sessions []Session) *SessionsResponse {
	this := SessionsResponse{}
	this.Sessions = sessions
	return &this
}

// NewSessionsResponseWithDefaults instantiates a new SessionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionsResponseWithDefaults() *SessionsResponse {
	this := SessionsResponse{}
	return &this
}

// GetSessions returns the Sessions field value
func (o *SessionsResponse) GetSessions() []Session {
	if o == nil {
		var ret []Session
		return ret
	}

	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *SessionsResponse) GetSessionsOk() ([]Session, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sessions, true
}

// SetSessions sets field value
func (o *SessionsResponse) SetSessions(v []Session) {
	o.Sessions = v
}

func (o SessionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessions"] = o.Sessions
	return toSerialize, nil
}

type NullableSessionsResponse struct {
	value *SessionsResponse
	isSet bool
}

func (v NullableSessionsResponse) Get() *SessionsResponse {
	return v.value
}

func (v *NullableSessionsResponse) Set(val *SessionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionsResponse(val *SessionsResponse) *NullableSessionsResponse {
	return &NullableSessionsResponse{value: val, isSet: true}
}

func (v NullableSessionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
