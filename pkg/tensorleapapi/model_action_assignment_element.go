/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.378-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ActionAssignmentElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionAssignmentElement{}

// ActionAssignmentElement struct for ActionAssignmentElement
type ActionAssignmentElement struct {
	AssignedUser string `json:"assignedUser"`
}

// NewActionAssignmentElement instantiates a new ActionAssignmentElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionAssignmentElement(assignedUser string) *ActionAssignmentElement {
	this := ActionAssignmentElement{}
	this.AssignedUser = assignedUser
	return &this
}

// NewActionAssignmentElementWithDefaults instantiates a new ActionAssignmentElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionAssignmentElementWithDefaults() *ActionAssignmentElement {
	this := ActionAssignmentElement{}
	return &this
}

// GetAssignedUser returns the AssignedUser field value
func (o *ActionAssignmentElement) GetAssignedUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedUser
}

// GetAssignedUserOk returns a tuple with the AssignedUser field value
// and a boolean to check if the value has been set.
func (o *ActionAssignmentElement) GetAssignedUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedUser, true
}

// SetAssignedUser sets field value
func (o *ActionAssignmentElement) SetAssignedUser(v string) {
	o.AssignedUser = v
}

func (o ActionAssignmentElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionAssignmentElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assignedUser"] = o.AssignedUser
	return toSerialize, nil
}

type NullableActionAssignmentElement struct {
	value *ActionAssignmentElement
	isSet bool
}

func (v NullableActionAssignmentElement) Get() *ActionAssignmentElement {
	return v.value
}

func (v *NullableActionAssignmentElement) Set(val *ActionAssignmentElement) {
	v.value = val
	v.isSet = true
}

func (v NullableActionAssignmentElement) IsSet() bool {
	return v.isSet
}

func (v *NullableActionAssignmentElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionAssignmentElement(val *ActionAssignmentElement) *NullableActionAssignmentElement {
	return &NullableActionAssignmentElement{value: val, isSet: true}
}

func (v NullableActionAssignmentElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionAssignmentElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


