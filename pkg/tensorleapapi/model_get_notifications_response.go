/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.464
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetNotificationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNotificationsResponse{}

// GetNotificationsResponse struct for GetNotificationsResponse
type GetNotificationsResponse struct {
	Notifications []Notification `json:"notifications"`
}

// NewGetNotificationsResponse instantiates a new GetNotificationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNotificationsResponse(notifications []Notification) *GetNotificationsResponse {
	this := GetNotificationsResponse{}
	this.Notifications = notifications
	return &this
}

// NewGetNotificationsResponseWithDefaults instantiates a new GetNotificationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNotificationsResponseWithDefaults() *GetNotificationsResponse {
	this := GetNotificationsResponse{}
	return &this
}

// GetNotifications returns the Notifications field value
func (o *GetNotificationsResponse) GetNotifications() []Notification {
	if o == nil {
		var ret []Notification
		return ret
	}

	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value
// and a boolean to check if the value has been set.
func (o *GetNotificationsResponse) GetNotificationsOk() ([]Notification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notifications, true
}

// SetNotifications sets field value
func (o *GetNotificationsResponse) SetNotifications(v []Notification) {
	o.Notifications = v
}

func (o GetNotificationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNotificationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifications"] = o.Notifications
	return toSerialize, nil
}

type NullableGetNotificationsResponse struct {
	value *GetNotificationsResponse
	isSet bool
}

func (v NullableGetNotificationsResponse) Get() *GetNotificationsResponse {
	return v.value
}

func (v *NullableGetNotificationsResponse) Set(val *GetNotificationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNotificationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNotificationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNotificationsResponse(val *GetNotificationsResponse) *NullableGetNotificationsResponse {
	return &NullableGetNotificationsResponse{value: val, isSet: true}
}

func (v NullableGetNotificationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNotificationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
