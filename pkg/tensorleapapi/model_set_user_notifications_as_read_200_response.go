/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SetUserNotificationsAsRead200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetUserNotificationsAsRead200Response{}

// SetUserNotificationsAsRead200Response struct for SetUserNotificationsAsRead200Response
type SetUserNotificationsAsRead200Response struct {
	Success bool `json:"success"`
}

// NewSetUserNotificationsAsRead200Response instantiates a new SetUserNotificationsAsRead200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetUserNotificationsAsRead200Response(success bool) *SetUserNotificationsAsRead200Response {
	this := SetUserNotificationsAsRead200Response{}
	this.Success = success
	return &this
}

// NewSetUserNotificationsAsRead200ResponseWithDefaults instantiates a new SetUserNotificationsAsRead200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetUserNotificationsAsRead200ResponseWithDefaults() *SetUserNotificationsAsRead200Response {
	this := SetUserNotificationsAsRead200Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *SetUserNotificationsAsRead200Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *SetUserNotificationsAsRead200Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *SetUserNotificationsAsRead200Response) SetSuccess(v bool) {
	o.Success = v
}

func (o SetUserNotificationsAsRead200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetUserNotificationsAsRead200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

type NullableSetUserNotificationsAsRead200Response struct {
	value *SetUserNotificationsAsRead200Response
	isSet bool
}

func (v NullableSetUserNotificationsAsRead200Response) Get() *SetUserNotificationsAsRead200Response {
	return v.value
}

func (v *NullableSetUserNotificationsAsRead200Response) Set(val *SetUserNotificationsAsRead200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSetUserNotificationsAsRead200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSetUserNotificationsAsRead200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetUserNotificationsAsRead200Response(val *SetUserNotificationsAsRead200Response) *NullableSetUserNotificationsAsRead200Response {
	return &NullableSetUserNotificationsAsRead200Response{value: val, isSet: true}
}

func (v NullableSetUserNotificationsAsRead200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetUserNotificationsAsRead200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
