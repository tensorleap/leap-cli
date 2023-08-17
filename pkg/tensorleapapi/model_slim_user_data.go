/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the SlimUserData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlimUserData{}

// SlimUserData struct for SlimUserData
type SlimUserData struct {
	Cid           string            `json:"cid"`
	Role          string            `json:"role"`
	RecordSession bool              `json:"recordSession"`
	Local         SlimUserDataLocal `json:"local"`
	CreatedAt     time.Time         `json:"createdAt"`
}

// NewSlimUserData instantiates a new SlimUserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlimUserData(cid string, role string, recordSession bool, local SlimUserDataLocal, createdAt time.Time) *SlimUserData {
	this := SlimUserData{}
	this.Cid = cid
	this.Role = role
	this.RecordSession = recordSession
	this.Local = local
	this.CreatedAt = createdAt
	return &this
}

// NewSlimUserDataWithDefaults instantiates a new SlimUserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlimUserDataWithDefaults() *SlimUserData {
	this := SlimUserData{}
	return &this
}

// GetCid returns the Cid field value
func (o *SlimUserData) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *SlimUserData) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *SlimUserData) SetCid(v string) {
	o.Cid = v
}

// GetRole returns the Role field value
func (o *SlimUserData) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *SlimUserData) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *SlimUserData) SetRole(v string) {
	o.Role = v
}

// GetRecordSession returns the RecordSession field value
func (o *SlimUserData) GetRecordSession() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecordSession
}

// GetRecordSessionOk returns a tuple with the RecordSession field value
// and a boolean to check if the value has been set.
func (o *SlimUserData) GetRecordSessionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordSession, true
}

// SetRecordSession sets field value
func (o *SlimUserData) SetRecordSession(v bool) {
	o.RecordSession = v
}

// GetLocal returns the Local field value
func (o *SlimUserData) GetLocal() SlimUserDataLocal {
	if o == nil {
		var ret SlimUserDataLocal
		return ret
	}

	return o.Local
}

// GetLocalOk returns a tuple with the Local field value
// and a boolean to check if the value has been set.
func (o *SlimUserData) GetLocalOk() (*SlimUserDataLocal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Local, true
}

// SetLocal sets field value
func (o *SlimUserData) SetLocal(v SlimUserDataLocal) {
	o.Local = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SlimUserData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SlimUserData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SlimUserData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o SlimUserData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlimUserData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["role"] = o.Role
	toSerialize["recordSession"] = o.RecordSession
	toSerialize["local"] = o.Local
	toSerialize["createdAt"] = o.CreatedAt
	return toSerialize, nil
}

type NullableSlimUserData struct {
	value *SlimUserData
	isSet bool
}

func (v NullableSlimUserData) Get() *SlimUserData {
	return v.value
}

func (v *NullableSlimUserData) Set(val *SlimUserData) {
	v.value = val
	v.isSet = true
}

func (v NullableSlimUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableSlimUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlimUserData(val *SlimUserData) *NullableSlimUserData {
	return &NullableSlimUserData{value: val, isSet: true}
}

func (v NullableSlimUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlimUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
