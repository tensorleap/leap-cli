/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.445
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetSlimVisualizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSlimVisualizationResponse{}

// GetSlimVisualizationResponse struct for GetSlimVisualizationResponse
type GetSlimVisualizationResponse struct {
	SlimVisualization SlimVisualization `json:"slimVisualization"`
}

// NewGetSlimVisualizationResponse instantiates a new GetSlimVisualizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSlimVisualizationResponse(slimVisualization SlimVisualization) *GetSlimVisualizationResponse {
	this := GetSlimVisualizationResponse{}
	this.SlimVisualization = slimVisualization
	return &this
}

// NewGetSlimVisualizationResponseWithDefaults instantiates a new GetSlimVisualizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSlimVisualizationResponseWithDefaults() *GetSlimVisualizationResponse {
	this := GetSlimVisualizationResponse{}
	return &this
}

// GetSlimVisualization returns the SlimVisualization field value
func (o *GetSlimVisualizationResponse) GetSlimVisualization() SlimVisualization {
	if o == nil {
		var ret SlimVisualization
		return ret
	}

	return o.SlimVisualization
}

// GetSlimVisualizationOk returns a tuple with the SlimVisualization field value
// and a boolean to check if the value has been set.
func (o *GetSlimVisualizationResponse) GetSlimVisualizationOk() (*SlimVisualization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlimVisualization, true
}

// SetSlimVisualization sets field value
func (o *GetSlimVisualizationResponse) SetSlimVisualization(v SlimVisualization) {
	o.SlimVisualization = v
}

func (o GetSlimVisualizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSlimVisualizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["slimVisualization"] = o.SlimVisualization
	return toSerialize, nil
}

type NullableGetSlimVisualizationResponse struct {
	value *GetSlimVisualizationResponse
	isSet bool
}

func (v NullableGetSlimVisualizationResponse) Get() *GetSlimVisualizationResponse {
	return v.value
}

func (v *NullableGetSlimVisualizationResponse) Set(val *GetSlimVisualizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSlimVisualizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSlimVisualizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSlimVisualizationResponse(val *GetSlimVisualizationResponse) *NullableGetSlimVisualizationResponse {
	return &NullableGetSlimVisualizationResponse{value: val, isSet: true}
}

func (v NullableGetSlimVisualizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSlimVisualizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
