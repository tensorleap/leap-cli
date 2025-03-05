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

// checks if the GetSessionRunsVisualizationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSessionRunsVisualizationsResponse{}

// GetSessionRunsVisualizationsResponse struct for GetSessionRunsVisualizationsResponse
type GetSessionRunsVisualizationsResponse struct {
	SlimVisualizations []SlimVisualization `json:"slimVisualizations"`
}

// NewGetSessionRunsVisualizationsResponse instantiates a new GetSessionRunsVisualizationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSessionRunsVisualizationsResponse(slimVisualizations []SlimVisualization) *GetSessionRunsVisualizationsResponse {
	this := GetSessionRunsVisualizationsResponse{}
	this.SlimVisualizations = slimVisualizations
	return &this
}

// NewGetSessionRunsVisualizationsResponseWithDefaults instantiates a new GetSessionRunsVisualizationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSessionRunsVisualizationsResponseWithDefaults() *GetSessionRunsVisualizationsResponse {
	this := GetSessionRunsVisualizationsResponse{}
	return &this
}

// GetSlimVisualizations returns the SlimVisualizations field value
func (o *GetSessionRunsVisualizationsResponse) GetSlimVisualizations() []SlimVisualization {
	if o == nil {
		var ret []SlimVisualization
		return ret
	}

	return o.SlimVisualizations
}

// GetSlimVisualizationsOk returns a tuple with the SlimVisualizations field value
// and a boolean to check if the value has been set.
func (o *GetSessionRunsVisualizationsResponse) GetSlimVisualizationsOk() ([]SlimVisualization, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlimVisualizations, true
}

// SetSlimVisualizations sets field value
func (o *GetSessionRunsVisualizationsResponse) SetSlimVisualizations(v []SlimVisualization) {
	o.SlimVisualizations = v
}

func (o GetSessionRunsVisualizationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSessionRunsVisualizationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["slimVisualizations"] = o.SlimVisualizations
	return toSerialize, nil
}

type NullableGetSessionRunsVisualizationsResponse struct {
	value *GetSessionRunsVisualizationsResponse
	isSet bool
}

func (v NullableGetSessionRunsVisualizationsResponse) Get() *GetSessionRunsVisualizationsResponse {
	return v.value
}

func (v *NullableGetSessionRunsVisualizationsResponse) Set(val *GetSessionRunsVisualizationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSessionRunsVisualizationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSessionRunsVisualizationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSessionRunsVisualizationsResponse(val *GetSessionRunsVisualizationsResponse) *NullableGetSessionRunsVisualizationsResponse {
	return &NullableGetSessionRunsVisualizationsResponse{value: val, isSet: true}
}

func (v NullableGetSessionRunsVisualizationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSessionRunsVisualizationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
