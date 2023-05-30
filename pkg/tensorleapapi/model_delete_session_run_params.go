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

// checks if the DeleteSessionRunParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSessionRunParams{}

// DeleteSessionRunParams struct for DeleteSessionRunParams
type DeleteSessionRunParams struct {
	SessionRunId string `json:"sessionRunId"`
}

// NewDeleteSessionRunParams instantiates a new DeleteSessionRunParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSessionRunParams(sessionRunId string) *DeleteSessionRunParams {
	this := DeleteSessionRunParams{}
	this.SessionRunId = sessionRunId
	return &this
}

// NewDeleteSessionRunParamsWithDefaults instantiates a new DeleteSessionRunParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSessionRunParamsWithDefaults() *DeleteSessionRunParams {
	this := DeleteSessionRunParams{}
	return &this
}

// GetSessionRunId returns the SessionRunId field value
func (o *DeleteSessionRunParams) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *DeleteSessionRunParams) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *DeleteSessionRunParams) SetSessionRunId(v string) {
	o.SessionRunId = v
}

func (o DeleteSessionRunParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteSessionRunParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionRunId"] = o.SessionRunId
	return toSerialize, nil
}

type NullableDeleteSessionRunParams struct {
	value *DeleteSessionRunParams
	isSet bool
}

func (v NullableDeleteSessionRunParams) Get() *DeleteSessionRunParams {
	return v.value
}

func (v *NullableDeleteSessionRunParams) Set(val *DeleteSessionRunParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSessionRunParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSessionRunParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSessionRunParams(val *DeleteSessionRunParams) *NullableDeleteSessionRunParams {
	return &NullableDeleteSessionRunParams{value: val, isSet: true}
}

func (v NullableDeleteSessionRunParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSessionRunParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
