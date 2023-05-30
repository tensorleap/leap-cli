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

// checks if the CreateOrganizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationResponse{}

// CreateOrganizationResponse struct for CreateOrganizationResponse
type CreateOrganizationResponse struct {
	OrganizationId string `json:"organizationId"`
}

// NewCreateOrganizationResponse instantiates a new CreateOrganizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationResponse(organizationId string) *CreateOrganizationResponse {
	this := CreateOrganizationResponse{}
	this.OrganizationId = organizationId
	return &this
}

// NewCreateOrganizationResponseWithDefaults instantiates a new CreateOrganizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationResponseWithDefaults() *CreateOrganizationResponse {
	this := CreateOrganizationResponse{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value
func (o *CreateOrganizationResponse) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *CreateOrganizationResponse) SetOrganizationId(v string) {
	o.OrganizationId = v
}

func (o CreateOrganizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["organizationId"] = o.OrganizationId
	return toSerialize, nil
}

type NullableCreateOrganizationResponse struct {
	value *CreateOrganizationResponse
	isSet bool
}

func (v NullableCreateOrganizationResponse) Get() *CreateOrganizationResponse {
	return v.value
}

func (v *NullableCreateOrganizationResponse) Set(val *CreateOrganizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationResponse(val *CreateOrganizationResponse) *NullableCreateOrganizationResponse {
	return &NullableCreateOrganizationResponse{value: val, isSet: true}
}

func (v NullableCreateOrganizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
