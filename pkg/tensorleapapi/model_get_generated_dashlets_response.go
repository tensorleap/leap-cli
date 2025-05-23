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

// checks if the GetGeneratedDashletsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGeneratedDashletsResponse{}

// GetGeneratedDashletsResponse struct for GetGeneratedDashletsResponse
type GetGeneratedDashletsResponse struct {
	Data []GetGeneratedDashletsBySessionRun `json:"data"`
}

// NewGetGeneratedDashletsResponse instantiates a new GetGeneratedDashletsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGeneratedDashletsResponse(data []GetGeneratedDashletsBySessionRun) *GetGeneratedDashletsResponse {
	this := GetGeneratedDashletsResponse{}
	this.Data = data
	return &this
}

// NewGetGeneratedDashletsResponseWithDefaults instantiates a new GetGeneratedDashletsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGeneratedDashletsResponseWithDefaults() *GetGeneratedDashletsResponse {
	this := GetGeneratedDashletsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetGeneratedDashletsResponse) GetData() []GetGeneratedDashletsBySessionRun {
	if o == nil {
		var ret []GetGeneratedDashletsBySessionRun
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetGeneratedDashletsResponse) GetDataOk() ([]GetGeneratedDashletsBySessionRun, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetGeneratedDashletsResponse) SetData(v []GetGeneratedDashletsBySessionRun) {
	o.Data = v
}

func (o GetGeneratedDashletsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGeneratedDashletsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetGeneratedDashletsResponse struct {
	value *GetGeneratedDashletsResponse
	isSet bool
}

func (v NullableGetGeneratedDashletsResponse) Get() *GetGeneratedDashletsResponse {
	return v.value
}

func (v *NullableGetGeneratedDashletsResponse) Set(val *GetGeneratedDashletsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGeneratedDashletsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGeneratedDashletsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGeneratedDashletsResponse(val *GetGeneratedDashletsResponse) *NullableGetGeneratedDashletsResponse {
	return &NullableGetGeneratedDashletsResponse{value: val, isSet: true}
}

func (v NullableGetGeneratedDashletsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGeneratedDashletsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
