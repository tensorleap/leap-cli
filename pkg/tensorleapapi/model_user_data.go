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

// checks if the UserData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserData{}

// UserData struct for UserData
type UserData struct {
	Id               string                  `json:"_id"`
	RecordSession    bool                    `json:"recordSession"`
	Role             UserRole                `json:"role"`
	Local            UserDataLocal           `json:"local"`
	Metadata         *UserDataMetadata       `json:"metadata,omitempty"`
	OrganizationName string                  `json:"organizationName"`
	SingleUserMode   *SingleUserModeSettings `json:"singleUserMode,omitempty"`
	Activated        bool                    `json:"activated"`
}

// NewUserData instantiates a new UserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserData(id string, recordSession bool, role UserRole, local UserDataLocal, organizationName string, activated bool) *UserData {
	this := UserData{}
	this.Id = id
	this.RecordSession = recordSession
	this.Role = role
	this.Local = local
	this.OrganizationName = organizationName
	this.Activated = activated
	return &this
}

// NewUserDataWithDefaults instantiates a new UserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataWithDefaults() *UserData {
	this := UserData{}
	return &this
}

// GetId returns the Id field value
func (o *UserData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserData) SetId(v string) {
	o.Id = v
}

// GetRecordSession returns the RecordSession field value
func (o *UserData) GetRecordSession() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecordSession
}

// GetRecordSessionOk returns a tuple with the RecordSession field value
// and a boolean to check if the value has been set.
func (o *UserData) GetRecordSessionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordSession, true
}

// SetRecordSession sets field value
func (o *UserData) SetRecordSession(v bool) {
	o.RecordSession = v
}

// GetRole returns the Role field value
func (o *UserData) GetRole() UserRole {
	if o == nil {
		var ret UserRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UserData) GetRoleOk() (*UserRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UserData) SetRole(v UserRole) {
	o.Role = v
}

// GetLocal returns the Local field value
func (o *UserData) GetLocal() UserDataLocal {
	if o == nil {
		var ret UserDataLocal
		return ret
	}

	return o.Local
}

// GetLocalOk returns a tuple with the Local field value
// and a boolean to check if the value has been set.
func (o *UserData) GetLocalOk() (*UserDataLocal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Local, true
}

// SetLocal sets field value
func (o *UserData) SetLocal(v UserDataLocal) {
	o.Local = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UserData) GetMetadata() UserDataMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret UserDataMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetMetadataOk() (*UserDataMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UserData) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given UserDataMetadata and assigns it to the Metadata field.
func (o *UserData) SetMetadata(v UserDataMetadata) {
	o.Metadata = &v
}

// GetOrganizationName returns the OrganizationName field value
func (o *UserData) GetOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
func (o *UserData) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationName, true
}

// SetOrganizationName sets field value
func (o *UserData) SetOrganizationName(v string) {
	o.OrganizationName = v
}

// GetSingleUserMode returns the SingleUserMode field value if set, zero value otherwise.
func (o *UserData) GetSingleUserMode() SingleUserModeSettings {
	if o == nil || IsNil(o.SingleUserMode) {
		var ret SingleUserModeSettings
		return ret
	}
	return *o.SingleUserMode
}

// GetSingleUserModeOk returns a tuple with the SingleUserMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetSingleUserModeOk() (*SingleUserModeSettings, bool) {
	if o == nil || IsNil(o.SingleUserMode) {
		return nil, false
	}
	return o.SingleUserMode, true
}

// HasSingleUserMode returns a boolean if a field has been set.
func (o *UserData) HasSingleUserMode() bool {
	if o != nil && !IsNil(o.SingleUserMode) {
		return true
	}

	return false
}

// SetSingleUserMode gets a reference to the given SingleUserModeSettings and assigns it to the SingleUserMode field.
func (o *UserData) SetSingleUserMode(v SingleUserModeSettings) {
	o.SingleUserMode = &v
}

// GetActivated returns the Activated field value
func (o *UserData) GetActivated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Activated
}

// GetActivatedOk returns a tuple with the Activated field value
// and a boolean to check if the value has been set.
func (o *UserData) GetActivatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Activated, true
}

// SetActivated sets field value
func (o *UserData) SetActivated(v bool) {
	o.Activated = v
}

func (o UserData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["recordSession"] = o.RecordSession
	toSerialize["role"] = o.Role
	toSerialize["local"] = o.Local
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["organizationName"] = o.OrganizationName
	if !IsNil(o.SingleUserMode) {
		toSerialize["singleUserMode"] = o.SingleUserMode
	}
	toSerialize["activated"] = o.Activated
	return toSerialize, nil
}

type NullableUserData struct {
	value *UserData
	isSet bool
}

func (v NullableUserData) Get() *UserData {
	return v.value
}

func (v *NullableUserData) Set(val *UserData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserData(val *UserData) *NullableUserData {
	return &NullableUserData{value: val, isSet: true}
}

func (v NullableUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
