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

// checks if the SetDefaultOrganizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetDefaultOrganizationRequest{}

// SetDefaultOrganizationRequest struct for SetDefaultOrganizationRequest
type SetDefaultOrganizationRequest struct {
	Id string `json:"_id"`
}

// NewSetDefaultOrganizationRequest instantiates a new SetDefaultOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetDefaultOrganizationRequest(id string) *SetDefaultOrganizationRequest {
	this := SetDefaultOrganizationRequest{}
	this.Id = id
	return &this
}

// NewSetDefaultOrganizationRequestWithDefaults instantiates a new SetDefaultOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetDefaultOrganizationRequestWithDefaults() *SetDefaultOrganizationRequest {
	this := SetDefaultOrganizationRequest{}
	return &this
}

// GetId returns the Id field value
func (o *SetDefaultOrganizationRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SetDefaultOrganizationRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SetDefaultOrganizationRequest) SetId(v string) {
	o.Id = v
}

func (o SetDefaultOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetDefaultOrganizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	return toSerialize, nil
}

type NullableSetDefaultOrganizationRequest struct {
	value *SetDefaultOrganizationRequest
	isSet bool
}

func (v NullableSetDefaultOrganizationRequest) Get() *SetDefaultOrganizationRequest {
	return v.value
}

func (v *NullableSetDefaultOrganizationRequest) Set(val *SetDefaultOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetDefaultOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetDefaultOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetDefaultOrganizationRequest(val *SetDefaultOrganizationRequest) *NullableSetDefaultOrganizationRequest {
	return &NullableSetDefaultOrganizationRequest{value: val, isSet: true}
}

func (v NullableSetDefaultOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetDefaultOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


