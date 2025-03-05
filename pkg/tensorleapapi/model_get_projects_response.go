/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetProjectsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectsResponse{}

// GetProjectsResponse struct for GetProjectsResponse
type GetProjectsResponse struct {
	Data []Project `json:"data"`
}

// NewGetProjectsResponse instantiates a new GetProjectsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectsResponse(data []Project) *GetProjectsResponse {
	this := GetProjectsResponse{}
	this.Data = data
	return &this
}

// NewGetProjectsResponseWithDefaults instantiates a new GetProjectsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectsResponseWithDefaults() *GetProjectsResponse {
	this := GetProjectsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetProjectsResponse) GetData() []Project {
	if o == nil {
		var ret []Project
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetProjectsResponse) GetDataOk() ([]Project, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetProjectsResponse) SetData(v []Project) {
	o.Data = v
}

func (o GetProjectsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetProjectsResponse struct {
	value *GetProjectsResponse
	isSet bool
}

func (v NullableGetProjectsResponse) Get() *GetProjectsResponse {
	return v.value
}

func (v *NullableGetProjectsResponse) Set(val *GetProjectsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectsResponse(val *GetProjectsResponse) *NullableGetProjectsResponse {
	return &NullableGetProjectsResponse{value: val, isSet: true}
}

func (v NullableGetProjectsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
