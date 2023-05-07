/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.285
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SessionHashRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionHashRequestParams{}

// SessionHashRequestParams struct for SessionHashRequestParams
type SessionHashRequestParams struct {
	Hash string `json:"hash"`
}

// NewSessionHashRequestParams instantiates a new SessionHashRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionHashRequestParams(hash string) *SessionHashRequestParams {
	this := SessionHashRequestParams{}
	this.Hash = hash
	return &this
}

// NewSessionHashRequestParamsWithDefaults instantiates a new SessionHashRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionHashRequestParamsWithDefaults() *SessionHashRequestParams {
	this := SessionHashRequestParams{}
	return &this
}

// GetHash returns the Hash field value
func (o *SessionHashRequestParams) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *SessionHashRequestParams) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *SessionHashRequestParams) SetHash(v string) {
	o.Hash = v
}

func (o SessionHashRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionHashRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hash"] = o.Hash
	return toSerialize, nil
}

type NullableSessionHashRequestParams struct {
	value *SessionHashRequestParams
	isSet bool
}

func (v NullableSessionHashRequestParams) Get() *SessionHashRequestParams {
	return v.value
}

func (v *NullableSessionHashRequestParams) Set(val *SessionHashRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionHashRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionHashRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionHashRequestParams(val *SessionHashRequestParams) *NullableSessionHashRequestParams {
	return &NullableSessionHashRequestParams{value: val, isSet: true}
}

func (v NullableSessionHashRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionHashRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


