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

// checks if the GetTeamsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTeamsResponse{}

// GetTeamsResponse struct for GetTeamsResponse
type GetTeamsResponse struct {
	Teams []SlimTeam `json:"teams"`
}

// NewGetTeamsResponse instantiates a new GetTeamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTeamsResponse(teams []SlimTeam) *GetTeamsResponse {
	this := GetTeamsResponse{}
	this.Teams = teams
	return &this
}

// NewGetTeamsResponseWithDefaults instantiates a new GetTeamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTeamsResponseWithDefaults() *GetTeamsResponse {
	this := GetTeamsResponse{}
	return &this
}

// GetTeams returns the Teams field value
func (o *GetTeamsResponse) GetTeams() []SlimTeam {
	if o == nil {
		var ret []SlimTeam
		return ret
	}

	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value
// and a boolean to check if the value has been set.
func (o *GetTeamsResponse) GetTeamsOk() ([]SlimTeam, bool) {
	if o == nil {
		return nil, false
	}
	return o.Teams, true
}

// SetTeams sets field value
func (o *GetTeamsResponse) SetTeams(v []SlimTeam) {
	o.Teams = v
}

func (o GetTeamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTeamsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["teams"] = o.Teams
	return toSerialize, nil
}

type NullableGetTeamsResponse struct {
	value *GetTeamsResponse
	isSet bool
}

func (v NullableGetTeamsResponse) Get() *GetTeamsResponse {
	return v.value
}

func (v *NullableGetTeamsResponse) Set(val *GetTeamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTeamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTeamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTeamsResponse(val *GetTeamsResponse) *NullableGetTeamsResponse {
	return &NullableGetTeamsResponse{value: val, isSet: true}
}

func (v NullableGetTeamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTeamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
