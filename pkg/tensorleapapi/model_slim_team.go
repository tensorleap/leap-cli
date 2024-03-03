/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SlimTeam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlimTeam{}

// SlimTeam struct for SlimTeam
type SlimTeam struct {
	Cid           string  `json:"cid"`
	Name          string  `json:"name"`
	PublicName    string  `json:"publicName"`
	Type          string  `json:"type"`
	IsDefault     bool    `json:"isDefault"`
	MachineTypeId *string `json:"machineTypeId,omitempty"`
}

// NewSlimTeam instantiates a new SlimTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlimTeam(cid string, name string, publicName string, type_ string, isDefault bool) *SlimTeam {
	this := SlimTeam{}
	this.Cid = cid
	this.Name = name
	this.PublicName = publicName
	this.Type = type_
	this.IsDefault = isDefault
	return &this
}

// NewSlimTeamWithDefaults instantiates a new SlimTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlimTeamWithDefaults() *SlimTeam {
	this := SlimTeam{}
	return &this
}

// GetCid returns the Cid field value
func (o *SlimTeam) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *SlimTeam) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *SlimTeam) SetCid(v string) {
	o.Cid = v
}

// GetName returns the Name field value
func (o *SlimTeam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SlimTeam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SlimTeam) SetName(v string) {
	o.Name = v
}

// GetPublicName returns the PublicName field value
func (o *SlimTeam) GetPublicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicName
}

// GetPublicNameOk returns a tuple with the PublicName field value
// and a boolean to check if the value has been set.
func (o *SlimTeam) GetPublicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicName, true
}

// SetPublicName sets field value
func (o *SlimTeam) SetPublicName(v string) {
	o.PublicName = v
}

// GetType returns the Type field value
func (o *SlimTeam) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SlimTeam) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SlimTeam) SetType(v string) {
	o.Type = v
}

// GetIsDefault returns the IsDefault field value
func (o *SlimTeam) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *SlimTeam) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *SlimTeam) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetMachineTypeId returns the MachineTypeId field value if set, zero value otherwise.
func (o *SlimTeam) GetMachineTypeId() string {
	if o == nil || IsNil(o.MachineTypeId) {
		var ret string
		return ret
	}
	return *o.MachineTypeId
}

// GetMachineTypeIdOk returns a tuple with the MachineTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimTeam) GetMachineTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.MachineTypeId) {
		return nil, false
	}
	return o.MachineTypeId, true
}

// HasMachineTypeId returns a boolean if a field has been set.
func (o *SlimTeam) HasMachineTypeId() bool {
	if o != nil && !IsNil(o.MachineTypeId) {
		return true
	}

	return false
}

// SetMachineTypeId gets a reference to the given string and assigns it to the MachineTypeId field.
func (o *SlimTeam) SetMachineTypeId(v string) {
	o.MachineTypeId = &v
}

func (o SlimTeam) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlimTeam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["name"] = o.Name
	toSerialize["publicName"] = o.PublicName
	toSerialize["type"] = o.Type
	toSerialize["isDefault"] = o.IsDefault
	if !IsNil(o.MachineTypeId) {
		toSerialize["machineTypeId"] = o.MachineTypeId
	}
	return toSerialize, nil
}

type NullableSlimTeam struct {
	value *SlimTeam
	isSet bool
}

func (v NullableSlimTeam) Get() *SlimTeam {
	return v.value
}

func (v *NullableSlimTeam) Set(val *SlimTeam) {
	v.value = val
	v.isSet = true
}

func (v NullableSlimTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableSlimTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlimTeam(val *SlimTeam) *NullableSlimTeam {
	return &NullableSlimTeam{value: val, isSet: true}
}

func (v NullableSlimTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlimTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
