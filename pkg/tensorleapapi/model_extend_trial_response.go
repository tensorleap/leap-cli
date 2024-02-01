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

// checks if the ExtendTrialResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtendTrialResponse{}

// ExtendTrialResponse struct for ExtendTrialResponse
type ExtendTrialResponse struct {
	Success bool `json:"success"`
}

// NewExtendTrialResponse instantiates a new ExtendTrialResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendTrialResponse(success bool) *ExtendTrialResponse {
	this := ExtendTrialResponse{}
	this.Success = success
	return &this
}

// NewExtendTrialResponseWithDefaults instantiates a new ExtendTrialResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendTrialResponseWithDefaults() *ExtendTrialResponse {
	this := ExtendTrialResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *ExtendTrialResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ExtendTrialResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ExtendTrialResponse) SetSuccess(v bool) {
	o.Success = v
}

func (o ExtendTrialResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendTrialResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

type NullableExtendTrialResponse struct {
	value *ExtendTrialResponse
	isSet bool
}

func (v NullableExtendTrialResponse) Get() *ExtendTrialResponse {
	return v.value
}

func (v *NullableExtendTrialResponse) Set(val *ExtendTrialResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendTrialResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendTrialResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendTrialResponse(val *ExtendTrialResponse) *NullableExtendTrialResponse {
	return &NullableExtendTrialResponse{value: val, isSet: true}
}

func (v NullableExtendTrialResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendTrialResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
