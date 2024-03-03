/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the CreateTeamRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTeamRequest{}

// CreateTeamRequest struct for CreateTeamRequest
type CreateTeamRequest struct {
	Name       string `json:"name"`
	PublicName string `json:"publicName"`
}

// NewCreateTeamRequest instantiates a new CreateTeamRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTeamRequest(name string, publicName string) *CreateTeamRequest {
	this := CreateTeamRequest{}
	this.Name = name
	this.PublicName = publicName
	return &this
}

// NewCreateTeamRequestWithDefaults instantiates a new CreateTeamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTeamRequestWithDefaults() *CreateTeamRequest {
	this := CreateTeamRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateTeamRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTeamRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateTeamRequest) SetName(v string) {
	o.Name = v
}

// GetPublicName returns the PublicName field value
func (o *CreateTeamRequest) GetPublicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicName
}

// GetPublicNameOk returns a tuple with the PublicName field value
// and a boolean to check if the value has been set.
func (o *CreateTeamRequest) GetPublicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicName, true
}

// SetPublicName sets field value
func (o *CreateTeamRequest) SetPublicName(v string) {
	o.PublicName = v
}

func (o CreateTeamRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTeamRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["publicName"] = o.PublicName
	return toSerialize, nil
}

type NullableCreateTeamRequest struct {
	value *CreateTeamRequest
	isSet bool
}

func (v NullableCreateTeamRequest) Get() *CreateTeamRequest {
	return v.value
}

func (v *NullableCreateTeamRequest) Set(val *CreateTeamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTeamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTeamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTeamRequest(val *CreateTeamRequest) *NullableCreateTeamRequest {
	return &NullableCreateTeamRequest{value: val, isSet: true}
}

func (v NullableCreateTeamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTeamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
