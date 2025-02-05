/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.592
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetStateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStateResponse{}

// GetStateResponse struct for GetStateResponse
type GetStateResponse struct {
	State string `json:"state"`
}

// NewGetStateResponse instantiates a new GetStateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStateResponse(state string) *GetStateResponse {
	this := GetStateResponse{}
	this.State = state
	return &this
}

// NewGetStateResponseWithDefaults instantiates a new GetStateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStateResponseWithDefaults() *GetStateResponse {
	this := GetStateResponse{}
	return &this
}

// GetState returns the State field value
func (o *GetStateResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *GetStateResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *GetStateResponse) SetState(v string) {
	o.State = v
}

func (o GetStateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	return toSerialize, nil
}

type NullableGetStateResponse struct {
	value *GetStateResponse
	isSet bool
}

func (v NullableGetStateResponse) Get() *GetStateResponse {
	return v.value
}

func (v *NullableGetStateResponse) Set(val *GetStateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStateResponse(val *GetStateResponse) *NullableGetStateResponse {
	return &NullableGetStateResponse{value: val, isSet: true}
}

func (v NullableGetStateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
