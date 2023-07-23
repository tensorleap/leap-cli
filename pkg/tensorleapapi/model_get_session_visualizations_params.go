/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetSessionVisualizationsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSessionVisualizationsParams{}

// GetSessionVisualizationsParams struct for GetSessionVisualizationsParams
type GetSessionVisualizationsParams struct {
	SessionRunId string `json:"sessionRunId"`
	ProjectId    string `json:"projectId"`
}

// NewGetSessionVisualizationsParams instantiates a new GetSessionVisualizationsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSessionVisualizationsParams(sessionRunId string, projectId string) *GetSessionVisualizationsParams {
	this := GetSessionVisualizationsParams{}
	this.SessionRunId = sessionRunId
	this.ProjectId = projectId
	return &this
}

// NewGetSessionVisualizationsParamsWithDefaults instantiates a new GetSessionVisualizationsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSessionVisualizationsParamsWithDefaults() *GetSessionVisualizationsParams {
	this := GetSessionVisualizationsParams{}
	return &this
}

// GetSessionRunId returns the SessionRunId field value
func (o *GetSessionVisualizationsParams) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *GetSessionVisualizationsParams) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *GetSessionVisualizationsParams) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetProjectId returns the ProjectId field value
func (o *GetSessionVisualizationsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetSessionVisualizationsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetSessionVisualizationsParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o GetSessionVisualizationsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSessionVisualizationsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableGetSessionVisualizationsParams struct {
	value *GetSessionVisualizationsParams
	isSet bool
}

func (v NullableGetSessionVisualizationsParams) Get() *GetSessionVisualizationsParams {
	return v.value
}

func (v *NullableGetSessionVisualizationsParams) Set(val *GetSessionVisualizationsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSessionVisualizationsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSessionVisualizationsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSessionVisualizationsParams(val *GetSessionVisualizationsParams) *NullableGetSessionVisualizationsParams {
	return &NullableGetSessionVisualizationsParams{value: val, isSet: true}
}

func (v NullableGetSessionVisualizationsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSessionVisualizationsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
