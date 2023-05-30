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

// checks if the GetMaxActiveUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMaxActiveUsersResponse{}

// GetMaxActiveUsersResponse struct for GetMaxActiveUsersResponse
type GetMaxActiveUsersResponse struct {
	MaxActiveUsers float64 `json:"maxActiveUsers"`
}

// NewGetMaxActiveUsersResponse instantiates a new GetMaxActiveUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMaxActiveUsersResponse(maxActiveUsers float64) *GetMaxActiveUsersResponse {
	this := GetMaxActiveUsersResponse{}
	this.MaxActiveUsers = maxActiveUsers
	return &this
}

// NewGetMaxActiveUsersResponseWithDefaults instantiates a new GetMaxActiveUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMaxActiveUsersResponseWithDefaults() *GetMaxActiveUsersResponse {
	this := GetMaxActiveUsersResponse{}
	return &this
}

// GetMaxActiveUsers returns the MaxActiveUsers field value
func (o *GetMaxActiveUsersResponse) GetMaxActiveUsers() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.MaxActiveUsers
}

// GetMaxActiveUsersOk returns a tuple with the MaxActiveUsers field value
// and a boolean to check if the value has been set.
func (o *GetMaxActiveUsersResponse) GetMaxActiveUsersOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxActiveUsers, true
}

// SetMaxActiveUsers sets field value
func (o *GetMaxActiveUsersResponse) SetMaxActiveUsers(v float64) {
	o.MaxActiveUsers = v
}

func (o GetMaxActiveUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMaxActiveUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxActiveUsers"] = o.MaxActiveUsers
	return toSerialize, nil
}

type NullableGetMaxActiveUsersResponse struct {
	value *GetMaxActiveUsersResponse
	isSet bool
}

func (v NullableGetMaxActiveUsersResponse) Get() *GetMaxActiveUsersResponse {
	return v.value
}

func (v *NullableGetMaxActiveUsersResponse) Set(val *GetMaxActiveUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMaxActiveUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMaxActiveUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMaxActiveUsersResponse(val *GetMaxActiveUsersResponse) *NullableGetMaxActiveUsersResponse {
	return &NullableGetMaxActiveUsersResponse{value: val, isSet: true}
}

func (v NullableGetMaxActiveUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMaxActiveUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
