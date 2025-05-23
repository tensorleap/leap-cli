/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the AddProjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddProjectResponse{}

// AddProjectResponse struct for AddProjectResponse
type AddProjectResponse struct {
	Project Project `json:"project"`
	Version Version `json:"version"`
}

// NewAddProjectResponse instantiates a new AddProjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddProjectResponse(project Project, version Version) *AddProjectResponse {
	this := AddProjectResponse{}
	this.Project = project
	this.Version = version
	return &this
}

// NewAddProjectResponseWithDefaults instantiates a new AddProjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddProjectResponseWithDefaults() *AddProjectResponse {
	this := AddProjectResponse{}
	return &this
}

// GetProject returns the Project field value
func (o *AddProjectResponse) GetProject() Project {
	if o == nil {
		var ret Project
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *AddProjectResponse) GetProjectOk() (*Project, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *AddProjectResponse) SetProject(v Project) {
	o.Project = v
}

// GetVersion returns the Version field value
func (o *AddProjectResponse) GetVersion() Version {
	if o == nil {
		var ret Version
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AddProjectResponse) GetVersionOk() (*Version, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AddProjectResponse) SetVersion(v Version) {
	o.Version = v
}

func (o AddProjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddProjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project"] = o.Project
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableAddProjectResponse struct {
	value *AddProjectResponse
	isSet bool
}

func (v NullableAddProjectResponse) Get() *AddProjectResponse {
	return v.value
}

func (v *NullableAddProjectResponse) Set(val *AddProjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddProjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddProjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddProjectResponse(val *AddProjectResponse) *NullableAddProjectResponse {
	return &NullableAddProjectResponse{value: val, isSet: true}
}

func (v NullableAddProjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddProjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
