/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ConfusionMetricNamesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfusionMetricNamesParams{}

// ConfusionMetricNamesParams struct for ConfusionMetricNamesParams
type ConfusionMetricNamesParams struct {
	ProjectId string `json:"projectId"`
}

// NewConfusionMetricNamesParams instantiates a new ConfusionMetricNamesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfusionMetricNamesParams(projectId string) *ConfusionMetricNamesParams {
	this := ConfusionMetricNamesParams{}
	this.ProjectId = projectId
	return &this
}

// NewConfusionMetricNamesParamsWithDefaults instantiates a new ConfusionMetricNamesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfusionMetricNamesParamsWithDefaults() *ConfusionMetricNamesParams {
	this := ConfusionMetricNamesParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *ConfusionMetricNamesParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ConfusionMetricNamesParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ConfusionMetricNamesParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o ConfusionMetricNamesParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfusionMetricNamesParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableConfusionMetricNamesParams struct {
	value *ConfusionMetricNamesParams
	isSet bool
}

func (v NullableConfusionMetricNamesParams) Get() *ConfusionMetricNamesParams {
	return v.value
}

func (v *NullableConfusionMetricNamesParams) Set(val *ConfusionMetricNamesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableConfusionMetricNamesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableConfusionMetricNamesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfusionMetricNamesParams(val *ConfusionMetricNamesParams) *NullableConfusionMetricNamesParams {
	return &NullableConfusionMetricNamesParams{value: val, isSet: true}
}

func (v NullableConfusionMetricNamesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfusionMetricNamesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


