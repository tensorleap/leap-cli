/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.285
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetVisualizationParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVisualizationParams{}

// GetVisualizationParams struct for GetVisualizationParams
type GetVisualizationParams struct {
	VisualizationId string `json:"visualizationId"`
}

// NewGetVisualizationParams instantiates a new GetVisualizationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVisualizationParams(visualizationId string) *GetVisualizationParams {
	this := GetVisualizationParams{}
	this.VisualizationId = visualizationId
	return &this
}

// NewGetVisualizationParamsWithDefaults instantiates a new GetVisualizationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVisualizationParamsWithDefaults() *GetVisualizationParams {
	this := GetVisualizationParams{}
	return &this
}

// GetVisualizationId returns the VisualizationId field value
func (o *GetVisualizationParams) GetVisualizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VisualizationId
}

// GetVisualizationIdOk returns a tuple with the VisualizationId field value
// and a boolean to check if the value has been set.
func (o *GetVisualizationParams) GetVisualizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisualizationId, true
}

// SetVisualizationId sets field value
func (o *GetVisualizationParams) SetVisualizationId(v string) {
	o.VisualizationId = v
}

func (o GetVisualizationParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVisualizationParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["visualizationId"] = o.VisualizationId
	return toSerialize, nil
}

type NullableGetVisualizationParams struct {
	value *GetVisualizationParams
	isSet bool
}

func (v NullableGetVisualizationParams) Get() *GetVisualizationParams {
	return v.value
}

func (v *NullableGetVisualizationParams) Set(val *GetVisualizationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVisualizationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVisualizationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVisualizationParams(val *GetVisualizationParams) *NullableGetVisualizationParams {
	return &NullableGetVisualizationParams{value: val, isSet: true}
}

func (v NullableGetVisualizationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVisualizationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


