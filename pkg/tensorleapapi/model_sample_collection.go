/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the SampleCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SampleCollection{}

// SampleCollection struct for SampleCollection
type SampleCollection struct {
	Samples     []SampleIdentity `json:"samples"`
	CreatedBy   *string          `json:"createdBy,omitempty"`
	CreatedAt   time.Time        `json:"createdAt"`
	Description *string          `json:"description,omitempty"`
	Name        string           `json:"name"`
	Cid         string           `json:"cid"`
	ProjectId   string           `json:"projectId"`
	TeamId      string           `json:"teamId"`
}

// NewSampleCollection instantiates a new SampleCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSampleCollection(samples []SampleIdentity, createdAt time.Time, name string, cid string, projectId string, teamId string) *SampleCollection {
	this := SampleCollection{}
	this.Samples = samples
	this.CreatedAt = createdAt
	this.Name = name
	this.Cid = cid
	this.ProjectId = projectId
	this.TeamId = teamId
	return &this
}

// NewSampleCollectionWithDefaults instantiates a new SampleCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSampleCollectionWithDefaults() *SampleCollection {
	this := SampleCollection{}
	return &this
}

// GetSamples returns the Samples field value
func (o *SampleCollection) GetSamples() []SampleIdentity {
	if o == nil {
		var ret []SampleIdentity
		return ret
	}

	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value
// and a boolean to check if the value has been set.
func (o *SampleCollection) GetSamplesOk() ([]SampleIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Samples, true
}

// SetSamples sets field value
func (o *SampleCollection) SetSamples(v []SampleIdentity) {
	o.Samples = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SampleCollection) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleCollection) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SampleCollection) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SampleCollection) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SampleCollection) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SampleCollection) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SampleCollection) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SampleCollection) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleCollection) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SampleCollection) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SampleCollection) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *SampleCollection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SampleCollection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SampleCollection) SetName(v string) {
	o.Name = v
}

// GetCid returns the Cid field value
func (o *SampleCollection) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *SampleCollection) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *SampleCollection) SetCid(v string) {
	o.Cid = v
}

// GetProjectId returns the ProjectId field value
func (o *SampleCollection) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *SampleCollection) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *SampleCollection) SetProjectId(v string) {
	o.ProjectId = v
}

// GetTeamId returns the TeamId field value
func (o *SampleCollection) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *SampleCollection) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *SampleCollection) SetTeamId(v string) {
	o.TeamId = v
}

func (o SampleCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SampleCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["samples"] = o.Samples
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	toSerialize["cid"] = o.Cid
	toSerialize["projectId"] = o.ProjectId
	toSerialize["teamId"] = o.TeamId
	return toSerialize, nil
}

type NullableSampleCollection struct {
	value *SampleCollection
	isSet bool
}

func (v NullableSampleCollection) Get() *SampleCollection {
	return v.value
}

func (v *NullableSampleCollection) Set(val *SampleCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableSampleCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableSampleCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSampleCollection(val *SampleCollection) *NullableSampleCollection {
	return &NullableSampleCollection{value: val, isSet: true}
}

func (v NullableSampleCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSampleCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
