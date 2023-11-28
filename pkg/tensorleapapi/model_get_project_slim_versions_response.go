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

// checks if the GetProjectSlimVersionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectSlimVersionsResponse{}

// GetProjectSlimVersionsResponse struct for GetProjectSlimVersionsResponse
type GetProjectSlimVersionsResponse struct {
	Versions []SlimVersion `json:"versions"`
}

// NewGetProjectSlimVersionsResponse instantiates a new GetProjectSlimVersionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectSlimVersionsResponse(versions []SlimVersion) *GetProjectSlimVersionsResponse {
	this := GetProjectSlimVersionsResponse{}
	this.Versions = versions
	return &this
}

// NewGetProjectSlimVersionsResponseWithDefaults instantiates a new GetProjectSlimVersionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectSlimVersionsResponseWithDefaults() *GetProjectSlimVersionsResponse {
	this := GetProjectSlimVersionsResponse{}
	return &this
}

// GetVersions returns the Versions field value
func (o *GetProjectSlimVersionsResponse) GetVersions() []SlimVersion {
	if o == nil {
		var ret []SlimVersion
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *GetProjectSlimVersionsResponse) GetVersionsOk() ([]SlimVersion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Versions, true
}

// SetVersions sets field value
func (o *GetProjectSlimVersionsResponse) SetVersions(v []SlimVersion) {
	o.Versions = v
}

func (o GetProjectSlimVersionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectSlimVersionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versions"] = o.Versions
	return toSerialize, nil
}

type NullableGetProjectSlimVersionsResponse struct {
	value *GetProjectSlimVersionsResponse
	isSet bool
}

func (v NullableGetProjectSlimVersionsResponse) Get() *GetProjectSlimVersionsResponse {
	return v.value
}

func (v *NullableGetProjectSlimVersionsResponse) Set(val *GetProjectSlimVersionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectSlimVersionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectSlimVersionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectSlimVersionsResponse(val *GetProjectSlimVersionsResponse) *NullableGetProjectSlimVersionsResponse {
	return &NullableGetProjectSlimVersionsResponse{value: val, isSet: true}
}

func (v NullableGetProjectSlimVersionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectSlimVersionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
