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

// checks if the RemoveSampleCollectionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveSampleCollectionParams{}

// RemoveSampleCollectionParams struct for RemoveSampleCollectionParams
type RemoveSampleCollectionParams struct {
	SampleCollectionIds []string `json:"sampleCollectionIds"`
	ProjectId           string   `json:"projectId"`
}

// NewRemoveSampleCollectionParams instantiates a new RemoveSampleCollectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveSampleCollectionParams(sampleCollectionIds []string, projectId string) *RemoveSampleCollectionParams {
	this := RemoveSampleCollectionParams{}
	this.SampleCollectionIds = sampleCollectionIds
	this.ProjectId = projectId
	return &this
}

// NewRemoveSampleCollectionParamsWithDefaults instantiates a new RemoveSampleCollectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveSampleCollectionParamsWithDefaults() *RemoveSampleCollectionParams {
	this := RemoveSampleCollectionParams{}
	return &this
}

// GetSampleCollectionIds returns the SampleCollectionIds field value
func (o *RemoveSampleCollectionParams) GetSampleCollectionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SampleCollectionIds
}

// GetSampleCollectionIdsOk returns a tuple with the SampleCollectionIds field value
// and a boolean to check if the value has been set.
func (o *RemoveSampleCollectionParams) GetSampleCollectionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SampleCollectionIds, true
}

// SetSampleCollectionIds sets field value
func (o *RemoveSampleCollectionParams) SetSampleCollectionIds(v []string) {
	o.SampleCollectionIds = v
}

// GetProjectId returns the ProjectId field value
func (o *RemoveSampleCollectionParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *RemoveSampleCollectionParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *RemoveSampleCollectionParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o RemoveSampleCollectionParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveSampleCollectionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sampleCollectionIds"] = o.SampleCollectionIds
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableRemoveSampleCollectionParams struct {
	value *RemoveSampleCollectionParams
	isSet bool
}

func (v NullableRemoveSampleCollectionParams) Get() *RemoveSampleCollectionParams {
	return v.value
}

func (v *NullableRemoveSampleCollectionParams) Set(val *RemoveSampleCollectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveSampleCollectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveSampleCollectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveSampleCollectionParams(val *RemoveSampleCollectionParams) *NullableRemoveSampleCollectionParams {
	return &NullableRemoveSampleCollectionParams{value: val, isSet: true}
}

func (v NullableRemoveSampleCollectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveSampleCollectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
