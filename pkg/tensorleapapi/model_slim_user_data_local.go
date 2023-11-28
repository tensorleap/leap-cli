/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.417
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SlimUserDataLocal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlimUserDataLocal{}

// SlimUserDataLocal struct for SlimUserDataLocal
type SlimUserDataLocal struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	TeamId    string `json:"teamId"`
	Activated bool   `json:"activated"`
}

// NewSlimUserDataLocal instantiates a new SlimUserDataLocal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlimUserDataLocal(name string, email string, teamId string, activated bool) *SlimUserDataLocal {
	this := SlimUserDataLocal{}
	this.Name = name
	this.Email = email
	this.TeamId = teamId
	this.Activated = activated
	return &this
}

// NewSlimUserDataLocalWithDefaults instantiates a new SlimUserDataLocal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlimUserDataLocalWithDefaults() *SlimUserDataLocal {
	this := SlimUserDataLocal{}
	return &this
}

// GetName returns the Name field value
func (o *SlimUserDataLocal) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SlimUserDataLocal) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SlimUserDataLocal) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *SlimUserDataLocal) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SlimUserDataLocal) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SlimUserDataLocal) SetEmail(v string) {
	o.Email = v
}

// GetTeamId returns the TeamId field value
func (o *SlimUserDataLocal) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *SlimUserDataLocal) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *SlimUserDataLocal) SetTeamId(v string) {
	o.TeamId = v
}

// GetActivated returns the Activated field value
func (o *SlimUserDataLocal) GetActivated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Activated
}

// GetActivatedOk returns a tuple with the Activated field value
// and a boolean to check if the value has been set.
func (o *SlimUserDataLocal) GetActivatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Activated, true
}

// SetActivated sets field value
func (o *SlimUserDataLocal) SetActivated(v bool) {
	o.Activated = v
}

func (o SlimUserDataLocal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlimUserDataLocal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	toSerialize["teamId"] = o.TeamId
	toSerialize["activated"] = o.Activated
	return toSerialize, nil
}

type NullableSlimUserDataLocal struct {
	value *SlimUserDataLocal
	isSet bool
}

func (v NullableSlimUserDataLocal) Get() *SlimUserDataLocal {
	return v.value
}

func (v *NullableSlimUserDataLocal) Set(val *SlimUserDataLocal) {
	v.value = val
	v.isSet = true
}

func (v NullableSlimUserDataLocal) IsSet() bool {
	return v.isSet
}

func (v *NullableSlimUserDataLocal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlimUserDataLocal(val *SlimUserDataLocal) *NullableSlimUserDataLocal {
	return &NullableSlimUserDataLocal{value: val, isSet: true}
}

func (v NullableSlimUserDataLocal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlimUserDataLocal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
