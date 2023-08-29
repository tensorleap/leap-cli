/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.378-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetSessionVisualizationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSessionVisualizationsResponse{}

// GetSessionVisualizationsResponse struct for GetSessionVisualizationsResponse
type GetSessionVisualizationsResponse struct {
	SlimVisualizations []SlimVisualization `json:"slimVisualizations"`
}

// NewGetSessionVisualizationsResponse instantiates a new GetSessionVisualizationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSessionVisualizationsResponse(slimVisualizations []SlimVisualization) *GetSessionVisualizationsResponse {
	this := GetSessionVisualizationsResponse{}
	this.SlimVisualizations = slimVisualizations
	return &this
}

// NewGetSessionVisualizationsResponseWithDefaults instantiates a new GetSessionVisualizationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSessionVisualizationsResponseWithDefaults() *GetSessionVisualizationsResponse {
	this := GetSessionVisualizationsResponse{}
	return &this
}

// GetSlimVisualizations returns the SlimVisualizations field value
func (o *GetSessionVisualizationsResponse) GetSlimVisualizations() []SlimVisualization {
	if o == nil {
		var ret []SlimVisualization
		return ret
	}

	return o.SlimVisualizations
}

// GetSlimVisualizationsOk returns a tuple with the SlimVisualizations field value
// and a boolean to check if the value has been set.
func (o *GetSessionVisualizationsResponse) GetSlimVisualizationsOk() ([]SlimVisualization, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlimVisualizations, true
}

// SetSlimVisualizations sets field value
func (o *GetSessionVisualizationsResponse) SetSlimVisualizations(v []SlimVisualization) {
	o.SlimVisualizations = v
}

func (o GetSessionVisualizationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSessionVisualizationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["slimVisualizations"] = o.SlimVisualizations
	return toSerialize, nil
}

type NullableGetSessionVisualizationsResponse struct {
	value *GetSessionVisualizationsResponse
	isSet bool
}

func (v NullableGetSessionVisualizationsResponse) Get() *GetSessionVisualizationsResponse {
	return v.value
}

func (v *NullableGetSessionVisualizationsResponse) Set(val *GetSessionVisualizationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSessionVisualizationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSessionVisualizationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSessionVisualizationsResponse(val *GetSessionVisualizationsResponse) *NullableGetSessionVisualizationsResponse {
	return &NullableGetSessionVisualizationsResponse{value: val, isSet: true}
}

func (v NullableGetSessionVisualizationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSessionVisualizationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


