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

// checks if the GetScatterSampleVisualizationsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetScatterSampleVisualizationsParams{}

// GetScatterSampleVisualizationsParams struct for GetScatterSampleVisualizationsParams
type GetScatterSampleVisualizationsParams struct {
	SessionRunId string  `json:"sessionRunId"`
	ProjectId    string  `json:"projectId"`
	Epoch        float64 `json:"epoch"`
}

// NewGetScatterSampleVisualizationsParams instantiates a new GetScatterSampleVisualizationsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetScatterSampleVisualizationsParams(sessionRunId string, projectId string, epoch float64) *GetScatterSampleVisualizationsParams {
	this := GetScatterSampleVisualizationsParams{}
	this.SessionRunId = sessionRunId
	this.ProjectId = projectId
	this.Epoch = epoch
	return &this
}

// NewGetScatterSampleVisualizationsParamsWithDefaults instantiates a new GetScatterSampleVisualizationsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetScatterSampleVisualizationsParamsWithDefaults() *GetScatterSampleVisualizationsParams {
	this := GetScatterSampleVisualizationsParams{}
	return &this
}

// GetSessionRunId returns the SessionRunId field value
func (o *GetScatterSampleVisualizationsParams) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *GetScatterSampleVisualizationsParams) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *GetScatterSampleVisualizationsParams) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetProjectId returns the ProjectId field value
func (o *GetScatterSampleVisualizationsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetScatterSampleVisualizationsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetScatterSampleVisualizationsParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetEpoch returns the Epoch field value
func (o *GetScatterSampleVisualizationsParams) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *GetScatterSampleVisualizationsParams) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *GetScatterSampleVisualizationsParams) SetEpoch(v float64) {
	o.Epoch = v
}

func (o GetScatterSampleVisualizationsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetScatterSampleVisualizationsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["projectId"] = o.ProjectId
	toSerialize["epoch"] = o.Epoch
	return toSerialize, nil
}

type NullableGetScatterSampleVisualizationsParams struct {
	value *GetScatterSampleVisualizationsParams
	isSet bool
}

func (v NullableGetScatterSampleVisualizationsParams) Get() *GetScatterSampleVisualizationsParams {
	return v.value
}

func (v *NullableGetScatterSampleVisualizationsParams) Set(val *GetScatterSampleVisualizationsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetScatterSampleVisualizationsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetScatterSampleVisualizationsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetScatterSampleVisualizationsParams(val *GetScatterSampleVisualizationsParams) *NullableGetScatterSampleVisualizationsParams {
	return &NullableGetScatterSampleVisualizationsParams{value: val, isSet: true}
}

func (v NullableGetScatterSampleVisualizationsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetScatterSampleVisualizationsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
