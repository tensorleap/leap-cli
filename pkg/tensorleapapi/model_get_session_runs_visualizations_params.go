/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.557
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetSessionRunsVisualizationsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSessionRunsVisualizationsParams{}

// GetSessionRunsVisualizationsParams struct for GetSessionRunsVisualizationsParams
type GetSessionRunsVisualizationsParams struct {
	SessionRunIds []string `json:"sessionRunIds"`
	ProjectId     string   `json:"projectId"`
}

// NewGetSessionRunsVisualizationsParams instantiates a new GetSessionRunsVisualizationsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSessionRunsVisualizationsParams(sessionRunIds []string, projectId string) *GetSessionRunsVisualizationsParams {
	this := GetSessionRunsVisualizationsParams{}
	this.SessionRunIds = sessionRunIds
	this.ProjectId = projectId
	return &this
}

// NewGetSessionRunsVisualizationsParamsWithDefaults instantiates a new GetSessionRunsVisualizationsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSessionRunsVisualizationsParamsWithDefaults() *GetSessionRunsVisualizationsParams {
	this := GetSessionRunsVisualizationsParams{}
	return &this
}

// GetSessionRunIds returns the SessionRunIds field value
func (o *GetSessionRunsVisualizationsParams) GetSessionRunIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SessionRunIds
}

// GetSessionRunIdsOk returns a tuple with the SessionRunIds field value
// and a boolean to check if the value has been set.
func (o *GetSessionRunsVisualizationsParams) GetSessionRunIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionRunIds, true
}

// SetSessionRunIds sets field value
func (o *GetSessionRunsVisualizationsParams) SetSessionRunIds(v []string) {
	o.SessionRunIds = v
}

// GetProjectId returns the ProjectId field value
func (o *GetSessionRunsVisualizationsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetSessionRunsVisualizationsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetSessionRunsVisualizationsParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o GetSessionRunsVisualizationsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSessionRunsVisualizationsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionRunIds"] = o.SessionRunIds
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableGetSessionRunsVisualizationsParams struct {
	value *GetSessionRunsVisualizationsParams
	isSet bool
}

func (v NullableGetSessionRunsVisualizationsParams) Get() *GetSessionRunsVisualizationsParams {
	return v.value
}

func (v *NullableGetSessionRunsVisualizationsParams) Set(val *GetSessionRunsVisualizationsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSessionRunsVisualizationsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSessionRunsVisualizationsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSessionRunsVisualizationsParams(val *GetSessionRunsVisualizationsParams) *NullableGetSessionRunsVisualizationsParams {
	return &NullableGetSessionRunsVisualizationsParams{value: val, isSet: true}
}

func (v NullableGetSessionRunsVisualizationsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSessionRunsVisualizationsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
