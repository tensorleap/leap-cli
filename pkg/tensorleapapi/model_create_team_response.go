/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the CreateTeamResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTeamResponse{}

// CreateTeamResponse struct for CreateTeamResponse
type CreateTeamResponse struct {
	TeamId string `json:"teamId"`
}

// NewCreateTeamResponse instantiates a new CreateTeamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTeamResponse(teamId string) *CreateTeamResponse {
	this := CreateTeamResponse{}
	this.TeamId = teamId
	return &this
}

// NewCreateTeamResponseWithDefaults instantiates a new CreateTeamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTeamResponseWithDefaults() *CreateTeamResponse {
	this := CreateTeamResponse{}
	return &this
}

// GetTeamId returns the TeamId field value
func (o *CreateTeamResponse) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *CreateTeamResponse) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *CreateTeamResponse) SetTeamId(v string) {
	o.TeamId = v
}

func (o CreateTeamResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTeamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["teamId"] = o.TeamId
	return toSerialize, nil
}

type NullableCreateTeamResponse struct {
	value *CreateTeamResponse
	isSet bool
}

func (v NullableCreateTeamResponse) Get() *CreateTeamResponse {
	return v.value
}

func (v *NullableCreateTeamResponse) Set(val *CreateTeamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTeamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTeamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTeamResponse(val *CreateTeamResponse) *NullableCreateTeamResponse {
	return &NullableCreateTeamResponse{value: val, isSet: true}
}

func (v NullableCreateTeamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTeamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


