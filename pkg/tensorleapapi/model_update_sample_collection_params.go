/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UpdateSampleCollectionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSampleCollectionParams{}

// UpdateSampleCollectionParams struct for UpdateSampleCollectionParams
type UpdateSampleCollectionParams struct {
	Samples            []SampleIdentity `json:"samples,omitempty"`
	Description        *string          `json:"description,omitempty"`
	Name               *string          `json:"name,omitempty"`
	SampleCollectionId string           `json:"sampleCollectionId"`
	ProjectId          string           `json:"projectId"`
}

// NewUpdateSampleCollectionParams instantiates a new UpdateSampleCollectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSampleCollectionParams(sampleCollectionId string, projectId string) *UpdateSampleCollectionParams {
	this := UpdateSampleCollectionParams{}
	this.SampleCollectionId = sampleCollectionId
	this.ProjectId = projectId
	return &this
}

// NewUpdateSampleCollectionParamsWithDefaults instantiates a new UpdateSampleCollectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSampleCollectionParamsWithDefaults() *UpdateSampleCollectionParams {
	this := UpdateSampleCollectionParams{}
	return &this
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *UpdateSampleCollectionParams) GetSamples() []SampleIdentity {
	if o == nil || IsNil(o.Samples) {
		var ret []SampleIdentity
		return ret
	}
	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSampleCollectionParams) GetSamplesOk() ([]SampleIdentity, bool) {
	if o == nil || IsNil(o.Samples) {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *UpdateSampleCollectionParams) HasSamples() bool {
	if o != nil && !IsNil(o.Samples) {
		return true
	}

	return false
}

// SetSamples gets a reference to the given []SampleIdentity and assigns it to the Samples field.
func (o *UpdateSampleCollectionParams) SetSamples(v []SampleIdentity) {
	o.Samples = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateSampleCollectionParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSampleCollectionParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateSampleCollectionParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateSampleCollectionParams) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateSampleCollectionParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSampleCollectionParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateSampleCollectionParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateSampleCollectionParams) SetName(v string) {
	o.Name = &v
}

// GetSampleCollectionId returns the SampleCollectionId field value
func (o *UpdateSampleCollectionParams) GetSampleCollectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SampleCollectionId
}

// GetSampleCollectionIdOk returns a tuple with the SampleCollectionId field value
// and a boolean to check if the value has been set.
func (o *UpdateSampleCollectionParams) GetSampleCollectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SampleCollectionId, true
}

// SetSampleCollectionId sets field value
func (o *UpdateSampleCollectionParams) SetSampleCollectionId(v string) {
	o.SampleCollectionId = v
}

// GetProjectId returns the ProjectId field value
func (o *UpdateSampleCollectionParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *UpdateSampleCollectionParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *UpdateSampleCollectionParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o UpdateSampleCollectionParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSampleCollectionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Samples) {
		toSerialize["samples"] = o.Samples
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["sampleCollectionId"] = o.SampleCollectionId
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableUpdateSampleCollectionParams struct {
	value *UpdateSampleCollectionParams
	isSet bool
}

func (v NullableUpdateSampleCollectionParams) Get() *UpdateSampleCollectionParams {
	return v.value
}

func (v *NullableUpdateSampleCollectionParams) Set(val *UpdateSampleCollectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSampleCollectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSampleCollectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSampleCollectionParams(val *UpdateSampleCollectionParams) *NullableUpdateSampleCollectionParams {
	return &NullableUpdateSampleCollectionParams{value: val, isSet: true}
}

func (v NullableUpdateSampleCollectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSampleCollectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
