/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the SlimVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlimVersion{}

// SlimVersion struct for SlimVersion
type SlimVersion struct {
	Cid                      string         `json:"cid"`
	CreatedBy                string         `json:"createdBy"`
	ProjectId                string         `json:"projectId"`
	BranchName               string         `json:"branchName"`
	Tags                     string         `json:"tags"`
	CreatedAt                time.Time      `json:"createdAt"`
	Notes                    string         `json:"notes"`
	CodeIntegrationVersionId *string        `json:"codeIntegrationVersionId,omitempty"`
	DatasetSetup             DatasetSetup   `json:"datasetSetup"`
	IsFavourite              bool           `json:"isFavourite"`
	Hash                     NullableString `json:"hash,omitempty"`
	Sessions                 []Session      `json:"sessions"`
}

// NewSlimVersion instantiates a new SlimVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlimVersion(cid string, createdBy string, projectId string, branchName string, tags string, createdAt time.Time, notes string, datasetSetup DatasetSetup, isFavourite bool, sessions []Session) *SlimVersion {
	this := SlimVersion{}
	this.Cid = cid
	this.CreatedBy = createdBy
	this.ProjectId = projectId
	this.BranchName = branchName
	this.Tags = tags
	this.CreatedAt = createdAt
	this.Notes = notes
	this.DatasetSetup = datasetSetup
	this.IsFavourite = isFavourite
	this.Sessions = sessions
	return &this
}

// NewSlimVersionWithDefaults instantiates a new SlimVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlimVersionWithDefaults() *SlimVersion {
	this := SlimVersion{}
	return &this
}

// GetCid returns the Cid field value
func (o *SlimVersion) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *SlimVersion) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *SlimVersion) SetCid(v string) {
	o.Cid = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *SlimVersion) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SlimVersion) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *SlimVersion) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetProjectId returns the ProjectId field value
func (o *SlimVersion) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *SlimVersion) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *SlimVersion) SetProjectId(v string) {
	o.ProjectId = v
}

// GetBranchName returns the BranchName field value
func (o *SlimVersion) GetBranchName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value
// and a boolean to check if the value has been set.
func (o *SlimVersion) GetBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchName, true
}

// SetBranchName sets field value
func (o *SlimVersion) SetBranchName(v string) {
	o.BranchName = v
}

// GetTags returns the Tags field value
func (o *SlimVersion) GetTags() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *SlimVersion) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *SlimVersion) SetTags(v string) {
	o.Tags = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SlimVersion) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SlimVersion) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SlimVersion) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetNotes returns the Notes field value
func (o *SlimVersion) GetNotes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *SlimVersion) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notes, true
}

// SetNotes sets field value
func (o *SlimVersion) SetNotes(v string) {
	o.Notes = v
}

// GetCodeIntegrationVersionId returns the CodeIntegrationVersionId field value if set, zero value otherwise.
func (o *SlimVersion) GetCodeIntegrationVersionId() string {
	if o == nil || IsNil(o.CodeIntegrationVersionId) {
		var ret string
		return ret
	}
	return *o.CodeIntegrationVersionId
}

// GetCodeIntegrationVersionIdOk returns a tuple with the CodeIntegrationVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlimVersion) GetCodeIntegrationVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CodeIntegrationVersionId) {
		return nil, false
	}
	return o.CodeIntegrationVersionId, true
}

// HasCodeIntegrationVersionId returns a boolean if a field has been set.
func (o *SlimVersion) HasCodeIntegrationVersionId() bool {
	if o != nil && !IsNil(o.CodeIntegrationVersionId) {
		return true
	}

	return false
}

// SetCodeIntegrationVersionId gets a reference to the given string and assigns it to the CodeIntegrationVersionId field.
func (o *SlimVersion) SetCodeIntegrationVersionId(v string) {
	o.CodeIntegrationVersionId = &v
}

// GetDatasetSetup returns the DatasetSetup field value
func (o *SlimVersion) GetDatasetSetup() DatasetSetup {
	if o == nil {
		var ret DatasetSetup
		return ret
	}

	return o.DatasetSetup
}

// GetDatasetSetupOk returns a tuple with the DatasetSetup field value
// and a boolean to check if the value has been set.
func (o *SlimVersion) GetDatasetSetupOk() (*DatasetSetup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetSetup, true
}

// SetDatasetSetup sets field value
func (o *SlimVersion) SetDatasetSetup(v DatasetSetup) {
	o.DatasetSetup = v
}

// GetIsFavourite returns the IsFavourite field value
func (o *SlimVersion) GetIsFavourite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFavourite
}

// GetIsFavouriteOk returns a tuple with the IsFavourite field value
// and a boolean to check if the value has been set.
func (o *SlimVersion) GetIsFavouriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFavourite, true
}

// SetIsFavourite sets field value
func (o *SlimVersion) SetIsFavourite(v bool) {
	o.IsFavourite = v
}

// GetHash returns the Hash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SlimVersion) GetHash() string {
	if o == nil || IsNil(o.Hash.Get()) {
		var ret string
		return ret
	}
	return *o.Hash.Get()
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SlimVersion) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hash.Get(), o.Hash.IsSet()
}

// HasHash returns a boolean if a field has been set.
func (o *SlimVersion) HasHash() bool {
	if o != nil && o.Hash.IsSet() {
		return true
	}

	return false
}

// SetHash gets a reference to the given NullableString and assigns it to the Hash field.
func (o *SlimVersion) SetHash(v string) {
	o.Hash.Set(&v)
}

// SetHashNil sets the value for Hash to be an explicit nil
func (o *SlimVersion) SetHashNil() {
	o.Hash.Set(nil)
}

// UnsetHash ensures that no value is present for Hash, not even an explicit nil
func (o *SlimVersion) UnsetHash() {
	o.Hash.Unset()
}

// GetSessions returns the Sessions field value
func (o *SlimVersion) GetSessions() []Session {
	if o == nil {
		var ret []Session
		return ret
	}

	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *SlimVersion) GetSessionsOk() ([]Session, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sessions, true
}

// SetSessions sets field value
func (o *SlimVersion) SetSessions(v []Session) {
	o.Sessions = v
}

func (o SlimVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlimVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["projectId"] = o.ProjectId
	toSerialize["branchName"] = o.BranchName
	toSerialize["tags"] = o.Tags
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["notes"] = o.Notes
	if !IsNil(o.CodeIntegrationVersionId) {
		toSerialize["codeIntegrationVersionId"] = o.CodeIntegrationVersionId
	}
	toSerialize["datasetSetup"] = o.DatasetSetup
	toSerialize["isFavourite"] = o.IsFavourite
	if o.Hash.IsSet() {
		toSerialize["hash"] = o.Hash.Get()
	}
	toSerialize["sessions"] = o.Sessions
	return toSerialize, nil
}

type NullableSlimVersion struct {
	value *SlimVersion
	isSet bool
}

func (v NullableSlimVersion) Get() *SlimVersion {
	return v.value
}

func (v *NullableSlimVersion) Set(val *SlimVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableSlimVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableSlimVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlimVersion(val *SlimVersion) *NullableSlimVersion {
	return &NullableSlimVersion{value: val, isSet: true}
}

func (v NullableSlimVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlimVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
