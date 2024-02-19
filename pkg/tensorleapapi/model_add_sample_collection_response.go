/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.464
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the AddSampleCollectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSampleCollectionResponse{}

// AddSampleCollectionResponse struct for AddSampleCollectionResponse
type AddSampleCollectionResponse struct {
	Cid string `json:"cid"`
}

// NewAddSampleCollectionResponse instantiates a new AddSampleCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSampleCollectionResponse(cid string) *AddSampleCollectionResponse {
	this := AddSampleCollectionResponse{}
	this.Cid = cid
	return &this
}

// NewAddSampleCollectionResponseWithDefaults instantiates a new AddSampleCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSampleCollectionResponseWithDefaults() *AddSampleCollectionResponse {
	this := AddSampleCollectionResponse{}
	return &this
}

// GetCid returns the Cid field value
func (o *AddSampleCollectionResponse) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *AddSampleCollectionResponse) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *AddSampleCollectionResponse) SetCid(v string) {
	o.Cid = v
}

func (o AddSampleCollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSampleCollectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	return toSerialize, nil
}

type NullableAddSampleCollectionResponse struct {
	value *AddSampleCollectionResponse
	isSet bool
}

func (v NullableAddSampleCollectionResponse) Get() *AddSampleCollectionResponse {
	return v.value
}

func (v *NullableAddSampleCollectionResponse) Set(val *AddSampleCollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSampleCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSampleCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSampleCollectionResponse(val *AddSampleCollectionResponse) *NullableAddSampleCollectionResponse {
	return &NullableAddSampleCollectionResponse{value: val, isSet: true}
}

func (v NullableAddSampleCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSampleCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
