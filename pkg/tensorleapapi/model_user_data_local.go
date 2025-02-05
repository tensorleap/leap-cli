/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.592
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UserDataLocal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDataLocal{}

// UserDataLocal struct for UserDataLocal
type UserDataLocal struct {
	Email          string `json:"email"`
	Name           string `json:"name"`
	TeamId         string `json:"teamId"`
	TrialRequested *bool  `json:"trialRequested,omitempty"`
}

// NewUserDataLocal instantiates a new UserDataLocal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDataLocal(email string, name string, teamId string) *UserDataLocal {
	this := UserDataLocal{}
	this.Email = email
	this.Name = name
	this.TeamId = teamId
	return &this
}

// NewUserDataLocalWithDefaults instantiates a new UserDataLocal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataLocalWithDefaults() *UserDataLocal {
	this := UserDataLocal{}
	return &this
}

// GetEmail returns the Email field value
func (o *UserDataLocal) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserDataLocal) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserDataLocal) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *UserDataLocal) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserDataLocal) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserDataLocal) SetName(v string) {
	o.Name = v
}

// GetTeamId returns the TeamId field value
func (o *UserDataLocal) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *UserDataLocal) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *UserDataLocal) SetTeamId(v string) {
	o.TeamId = v
}

// GetTrialRequested returns the TrialRequested field value if set, zero value otherwise.
func (o *UserDataLocal) GetTrialRequested() bool {
	if o == nil || IsNil(o.TrialRequested) {
		var ret bool
		return ret
	}
	return *o.TrialRequested
}

// GetTrialRequestedOk returns a tuple with the TrialRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataLocal) GetTrialRequestedOk() (*bool, bool) {
	if o == nil || IsNil(o.TrialRequested) {
		return nil, false
	}
	return o.TrialRequested, true
}

// HasTrialRequested returns a boolean if a field has been set.
func (o *UserDataLocal) HasTrialRequested() bool {
	if o != nil && !IsNil(o.TrialRequested) {
		return true
	}

	return false
}

// SetTrialRequested gets a reference to the given bool and assigns it to the TrialRequested field.
func (o *UserDataLocal) SetTrialRequested(v bool) {
	o.TrialRequested = &v
}

func (o UserDataLocal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDataLocal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["name"] = o.Name
	toSerialize["teamId"] = o.TeamId
	if !IsNil(o.TrialRequested) {
		toSerialize["trialRequested"] = o.TrialRequested
	}
	return toSerialize, nil
}

type NullableUserDataLocal struct {
	value *UserDataLocal
	isSet bool
}

func (v NullableUserDataLocal) Get() *UserDataLocal {
	return v.value
}

func (v *NullableUserDataLocal) Set(val *UserDataLocal) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataLocal) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataLocal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataLocal(val *UserDataLocal) *NullableUserDataLocal {
	return &NullableUserDataLocal{value: val, isSet: true}
}

func (v NullableUserDataLocal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataLocal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
