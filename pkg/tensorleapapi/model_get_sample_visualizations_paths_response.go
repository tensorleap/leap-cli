/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.423
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetSampleVisualizationsPathsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSampleVisualizationsPathsResponse{}

// GetSampleVisualizationsPathsResponse struct for GetSampleVisualizationsPathsResponse
type GetSampleVisualizationsPathsResponse struct {
	Paths []string `json:"paths"`
}

// NewGetSampleVisualizationsPathsResponse instantiates a new GetSampleVisualizationsPathsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSampleVisualizationsPathsResponse(paths []string) *GetSampleVisualizationsPathsResponse {
	this := GetSampleVisualizationsPathsResponse{}
	this.Paths = paths
	return &this
}

// NewGetSampleVisualizationsPathsResponseWithDefaults instantiates a new GetSampleVisualizationsPathsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSampleVisualizationsPathsResponseWithDefaults() *GetSampleVisualizationsPathsResponse {
	this := GetSampleVisualizationsPathsResponse{}
	return &this
}

// GetPaths returns the Paths field value
func (o *GetSampleVisualizationsPathsResponse) GetPaths() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value
// and a boolean to check if the value has been set.
func (o *GetSampleVisualizationsPathsResponse) GetPathsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Paths, true
}

// SetPaths sets field value
func (o *GetSampleVisualizationsPathsResponse) SetPaths(v []string) {
	o.Paths = v
}

func (o GetSampleVisualizationsPathsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSampleVisualizationsPathsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paths"] = o.Paths
	return toSerialize, nil
}

type NullableGetSampleVisualizationsPathsResponse struct {
	value *GetSampleVisualizationsPathsResponse
	isSet bool
}

func (v NullableGetSampleVisualizationsPathsResponse) Get() *GetSampleVisualizationsPathsResponse {
	return v.value
}

func (v *NullableGetSampleVisualizationsPathsResponse) Set(val *GetSampleVisualizationsPathsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSampleVisualizationsPathsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSampleVisualizationsPathsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSampleVisualizationsPathsResponse(val *GetSampleVisualizationsPathsResponse) *NullableGetSampleVisualizationsPathsResponse {
	return &NullableGetSampleVisualizationsPathsResponse{value: val, isSet: true}
}

func (v NullableGetSampleVisualizationsPathsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSampleVisualizationsPathsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
