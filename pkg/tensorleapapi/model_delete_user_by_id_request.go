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

// checks if the DeleteUserByIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteUserByIdRequest{}

// DeleteUserByIdRequest struct for DeleteUserByIdRequest
type DeleteUserByIdRequest struct {
	UserId string `json:"userId"`
}

// NewDeleteUserByIdRequest instantiates a new DeleteUserByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteUserByIdRequest(userId string) *DeleteUserByIdRequest {
	this := DeleteUserByIdRequest{}
	this.UserId = userId
	return &this
}

// NewDeleteUserByIdRequestWithDefaults instantiates a new DeleteUserByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteUserByIdRequestWithDefaults() *DeleteUserByIdRequest {
	this := DeleteUserByIdRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *DeleteUserByIdRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *DeleteUserByIdRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *DeleteUserByIdRequest) SetUserId(v string) {
	o.UserId = v
}

func (o DeleteUserByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteUserByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

type NullableDeleteUserByIdRequest struct {
	value *DeleteUserByIdRequest
	isSet bool
}

func (v NullableDeleteUserByIdRequest) Get() *DeleteUserByIdRequest {
	return v.value
}

func (v *NullableDeleteUserByIdRequest) Set(val *DeleteUserByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteUserByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteUserByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteUserByIdRequest(val *DeleteUserByIdRequest) *NullableDeleteUserByIdRequest {
	return &NullableDeleteUserByIdRequest{value: val, isSet: true}
}

func (v NullableDeleteUserByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteUserByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
