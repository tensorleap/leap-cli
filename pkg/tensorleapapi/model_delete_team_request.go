/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the DeleteTeamRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteTeamRequest{}

// DeleteTeamRequest struct for DeleteTeamRequest
type DeleteTeamRequest struct {
	Cid string `json:"cid"`
}

// NewDeleteTeamRequest instantiates a new DeleteTeamRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTeamRequest(cid string) *DeleteTeamRequest {
	this := DeleteTeamRequest{}
	this.Cid = cid
	return &this
}

// NewDeleteTeamRequestWithDefaults instantiates a new DeleteTeamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTeamRequestWithDefaults() *DeleteTeamRequest {
	this := DeleteTeamRequest{}
	return &this
}

// GetCid returns the Cid field value
func (o *DeleteTeamRequest) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *DeleteTeamRequest) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *DeleteTeamRequest) SetCid(v string) {
	o.Cid = v
}

func (o DeleteTeamRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTeamRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	return toSerialize, nil
}

type NullableDeleteTeamRequest struct {
	value *DeleteTeamRequest
	isSet bool
}

func (v NullableDeleteTeamRequest) Get() *DeleteTeamRequest {
	return v.value
}

func (v *NullableDeleteTeamRequest) Set(val *DeleteTeamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTeamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTeamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTeamRequest(val *DeleteTeamRequest) *NullableDeleteTeamRequest {
	return &NullableDeleteTeamRequest{value: val, isSet: true}
}

func (v NullableDeleteTeamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTeamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


