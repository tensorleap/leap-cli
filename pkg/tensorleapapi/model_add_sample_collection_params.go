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

// checks if the AddSampleCollectionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSampleCollectionParams{}

// AddSampleCollectionParams struct for AddSampleCollectionParams
type AddSampleCollectionParams struct {
	Samples []SampleIdentity `json:"samples"`
	ProjectId string `json:"projectId"`
	Description *string `json:"description,omitempty"`
	Name string `json:"name"`
}

// NewAddSampleCollectionParams instantiates a new AddSampleCollectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSampleCollectionParams(samples []SampleIdentity, projectId string, name string) *AddSampleCollectionParams {
	this := AddSampleCollectionParams{}
	this.Samples = samples
	this.ProjectId = projectId
	this.Name = name
	return &this
}

// NewAddSampleCollectionParamsWithDefaults instantiates a new AddSampleCollectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSampleCollectionParamsWithDefaults() *AddSampleCollectionParams {
	this := AddSampleCollectionParams{}
	return &this
}

// GetSamples returns the Samples field value
func (o *AddSampleCollectionParams) GetSamples() []SampleIdentity {
	if o == nil {
		var ret []SampleIdentity
		return ret
	}

	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value
// and a boolean to check if the value has been set.
func (o *AddSampleCollectionParams) GetSamplesOk() ([]SampleIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Samples, true
}

// SetSamples sets field value
func (o *AddSampleCollectionParams) SetSamples(v []SampleIdentity) {
	o.Samples = v
}

// GetProjectId returns the ProjectId field value
func (o *AddSampleCollectionParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AddSampleCollectionParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AddSampleCollectionParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSampleCollectionParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSampleCollectionParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSampleCollectionParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSampleCollectionParams) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *AddSampleCollectionParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddSampleCollectionParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddSampleCollectionParams) SetName(v string) {
	o.Name = v
}

func (o AddSampleCollectionParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSampleCollectionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["samples"] = o.Samples
	toSerialize["projectId"] = o.ProjectId
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableAddSampleCollectionParams struct {
	value *AddSampleCollectionParams
	isSet bool
}

func (v NullableAddSampleCollectionParams) Get() *AddSampleCollectionParams {
	return v.value
}

func (v *NullableAddSampleCollectionParams) Set(val *AddSampleCollectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSampleCollectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSampleCollectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSampleCollectionParams(val *AddSampleCollectionParams) *NullableAddSampleCollectionParams {
	return &NullableAddSampleCollectionParams{value: val, isSet: true}
}

func (v NullableAddSampleCollectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSampleCollectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


