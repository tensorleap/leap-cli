/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.422
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the LoadVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadVersionResponse{}

// LoadVersionResponse struct for LoadVersionResponse
type LoadVersionResponse struct {
	Project Project `json:"project"`
	Version Version `json:"version"`
}

// NewLoadVersionResponse instantiates a new LoadVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadVersionResponse(project Project, version Version) *LoadVersionResponse {
	this := LoadVersionResponse{}
	this.Project = project
	this.Version = version
	return &this
}

// NewLoadVersionResponseWithDefaults instantiates a new LoadVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadVersionResponseWithDefaults() *LoadVersionResponse {
	this := LoadVersionResponse{}
	return &this
}

// GetProject returns the Project field value
func (o *LoadVersionResponse) GetProject() Project {
	if o == nil {
		var ret Project
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *LoadVersionResponse) GetProjectOk() (*Project, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *LoadVersionResponse) SetProject(v Project) {
	o.Project = v
}

// GetVersion returns the Version field value
func (o *LoadVersionResponse) GetVersion() Version {
	if o == nil {
		var ret Version
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *LoadVersionResponse) GetVersionOk() (*Version, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *LoadVersionResponse) SetVersion(v Version) {
	o.Version = v
}

func (o LoadVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project"] = o.Project
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableLoadVersionResponse struct {
	value *LoadVersionResponse
	isSet bool
}

func (v NullableLoadVersionResponse) Get() *LoadVersionResponse {
	return v.value
}

func (v *NullableLoadVersionResponse) Set(val *LoadVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadVersionResponse(val *LoadVersionResponse) *NullableLoadVersionResponse {
	return &NullableLoadVersionResponse{value: val, isSet: true}
}

func (v NullableLoadVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
