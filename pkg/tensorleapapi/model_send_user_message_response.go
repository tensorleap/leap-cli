/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.321
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SendUserMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendUserMessageResponse{}

// SendUserMessageResponse struct for SendUserMessageResponse
type SendUserMessageResponse struct {
	Success bool `json:"success"`
}

// NewSendUserMessageResponse instantiates a new SendUserMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendUserMessageResponse(success bool) *SendUserMessageResponse {
	this := SendUserMessageResponse{}
	this.Success = success
	return &this
}

// NewSendUserMessageResponseWithDefaults instantiates a new SendUserMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendUserMessageResponseWithDefaults() *SendUserMessageResponse {
	this := SendUserMessageResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *SendUserMessageResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *SendUserMessageResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *SendUserMessageResponse) SetSuccess(v bool) {
	o.Success = v
}

func (o SendUserMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendUserMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

type NullableSendUserMessageResponse struct {
	value *SendUserMessageResponse
	isSet bool
}

func (v NullableSendUserMessageResponse) Get() *SendUserMessageResponse {
	return v.value
}

func (v *NullableSendUserMessageResponse) Set(val *SendUserMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSendUserMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSendUserMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendUserMessageResponse(val *SendUserMessageResponse) *NullableSendUserMessageResponse {
	return &NullableSendUserMessageResponse{value: val, isSet: true}
}

func (v NullableSendUserMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendUserMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
