/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.422
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the AddVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddVersionResponse{}

// AddVersionResponse struct for AddVersionResponse
type AddVersionResponse struct {
	Version Version `json:"version"`
}

// NewAddVersionResponse instantiates a new AddVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVersionResponse(version Version) *AddVersionResponse {
	this := AddVersionResponse{}
	this.Version = version
	return &this
}

// NewAddVersionResponseWithDefaults instantiates a new AddVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVersionResponseWithDefaults() *AddVersionResponse {
	this := AddVersionResponse{}
	return &this
}

// GetVersion returns the Version field value
func (o *AddVersionResponse) GetVersion() Version {
	if o == nil {
		var ret Version
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AddVersionResponse) GetVersionOk() (*Version, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AddVersionResponse) SetVersion(v Version) {
	o.Version = v
}

func (o AddVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableAddVersionResponse struct {
	value *AddVersionResponse
	isSet bool
}

func (v NullableAddVersionResponse) Get() *AddVersionResponse {
	return v.value
}

func (v *NullableAddVersionResponse) Set(val *AddVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVersionResponse(val *AddVersionResponse) *NullableAddVersionResponse {
	return &NullableAddVersionResponse{value: val, isSet: true}
}

func (v NullableAddVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
