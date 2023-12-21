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

// checks if the GetSlimVisualizationParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSlimVisualizationParams{}

// GetSlimVisualizationParams struct for GetSlimVisualizationParams
type GetSlimVisualizationParams struct {
	VisualizationId string `json:"visualizationId"`
	ProjectId       string `json:"projectId"`
}

// NewGetSlimVisualizationParams instantiates a new GetSlimVisualizationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSlimVisualizationParams(visualizationId string, projectId string) *GetSlimVisualizationParams {
	this := GetSlimVisualizationParams{}
	this.VisualizationId = visualizationId
	this.ProjectId = projectId
	return &this
}

// NewGetSlimVisualizationParamsWithDefaults instantiates a new GetSlimVisualizationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSlimVisualizationParamsWithDefaults() *GetSlimVisualizationParams {
	this := GetSlimVisualizationParams{}
	return &this
}

// GetVisualizationId returns the VisualizationId field value
func (o *GetSlimVisualizationParams) GetVisualizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VisualizationId
}

// GetVisualizationIdOk returns a tuple with the VisualizationId field value
// and a boolean to check if the value has been set.
func (o *GetSlimVisualizationParams) GetVisualizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisualizationId, true
}

// SetVisualizationId sets field value
func (o *GetSlimVisualizationParams) SetVisualizationId(v string) {
	o.VisualizationId = v
}

// GetProjectId returns the ProjectId field value
func (o *GetSlimVisualizationParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetSlimVisualizationParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetSlimVisualizationParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o GetSlimVisualizationParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSlimVisualizationParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["visualizationId"] = o.VisualizationId
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableGetSlimVisualizationParams struct {
	value *GetSlimVisualizationParams
	isSet bool
}

func (v NullableGetSlimVisualizationParams) Get() *GetSlimVisualizationParams {
	return v.value
}

func (v *NullableGetSlimVisualizationParams) Set(val *GetSlimVisualizationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSlimVisualizationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSlimVisualizationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSlimVisualizationParams(val *GetSlimVisualizationParams) *NullableGetSlimVisualizationParams {
	return &NullableGetSlimVisualizationParams{value: val, isSet: true}
}

func (v NullableGetSlimVisualizationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSlimVisualizationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
