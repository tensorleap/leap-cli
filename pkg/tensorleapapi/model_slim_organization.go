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

// checks if the SlimOrganization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlimOrganization{}

// SlimOrganization struct for SlimOrganization
type SlimOrganization struct {
	Id string `json:"_id"`
	Name string `json:"name"`
	PublicName string `json:"publicName"`
	Type string `json:"type"`
	IsDefault bool `json:"isDefault"`
	MachineTypeId *string `json:"machineTypeId,omitempty"`
}

// NewSlimOrganization instantiates a new SlimOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlimOrganization(id string, name string, publicName string, type_ string, isDefault bool) *SlimOrganization {
	this := SlimOrganization{}
	this.Id = id
	this.Name = name
	this.PublicName = publicName
	this.Type = type_
	this.IsDefault = isDefault
	return &this
}

// NewSlimOrganizationWithDefaults instantiates a new SlimOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlimOrganizationWithDefaults() *SlimOrganization {
	this := SlimOrganization{}
	return &this
}

// GetId returns the Id field value
func (o *SlimOrganization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SlimOrganization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SlimOrganization) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SlimOrganization) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SlimOrganization) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SlimOrganization) SetName(v string) {
	o.Name = v
}

// GetPublicName returns the PublicName field value
func (o *SlimOrganization) GetPublicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicName
}

// GetPublicNameOk returns a tuple with the PublicName field value
// and a boolean to check if the value has been set.
func (o *SlimOrganization) GetPublicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicName, true
}

// SetPublicName sets field value
func (o *SlimOrganization) SetPublicName(v string) {
	o.PublicName = v
}

// GetType returns the Type field value
func (o *SlimOrganization) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SlimOrganization) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SlimOrganization) SetType(v string) {
	o.Type = v
}

// GetIsDefault returns the IsDefault field value
func (o *SlimOrganization) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *SlimOrganization) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *SlimOrganization) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetMachineTypeId returns the MachineTypeId field value if set, zero value otherwise.
func (o *SlimOrganization) GetMachineTypeId() string {
	if o == nil || IsNil(o.MachineTypeId) {
		var ret string
		return ret
	}
	return *o.MachineTypeId
}

// GetMachineTypeIdOk returns a tuple with the MachineTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimOrganization) GetMachineTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.MachineTypeId) {
		return nil, false
	}
	return o.MachineTypeId, true
}

// HasMachineTypeId returns a boolean if a field has been set.
func (o *SlimOrganization) HasMachineTypeId() bool {
	if o != nil && !IsNil(o.MachineTypeId) {
		return true
	}

	return false
}

// SetMachineTypeId gets a reference to the given string and assigns it to the MachineTypeId field.
func (o *SlimOrganization) SetMachineTypeId(v string) {
	o.MachineTypeId = &v
}

func (o SlimOrganization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlimOrganization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["publicName"] = o.PublicName
	toSerialize["type"] = o.Type
	toSerialize["isDefault"] = o.IsDefault
	if !IsNil(o.MachineTypeId) {
		toSerialize["machineTypeId"] = o.MachineTypeId
	}
	return toSerialize, nil
}

type NullableSlimOrganization struct {
	value *SlimOrganization
	isSet bool
}

func (v NullableSlimOrganization) Get() *SlimOrganization {
	return v.value
}

func (v *NullableSlimOrganization) Set(val *SlimOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableSlimOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableSlimOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlimOrganization(val *SlimOrganization) *NullableSlimOrganization {
	return &NullableSlimOrganization{value: val, isSet: true}
}

func (v NullableSlimOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlimOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


