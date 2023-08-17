/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the Version type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Version{}

// Version struct for Version
type Version struct {
	Cid                      string       `json:"cid"`
	ExtId                    string       `json:"extId"`
	CreatedBy                string       `json:"createdBy"`
	TeamId                   string       `json:"teamId"`
	ProjectId                string       `json:"projectId"`
	Branch                   string       `json:"branch"`
	Tag                      string       `json:"tag"`
	Data                     ModelGraph   `json:"data"`
	CreatedAt                time.Time    `json:"createdAt"`
	Notes                    string       `json:"notes"`
	Status                   string       `json:"status"`
	IsFavourite              bool         `json:"isFavourite"`
	CodeIntegrationVersionId *string      `json:"codeIntegrationVersionId,omitempty"`
	DatasetSetup             DatasetSetup `json:"datasetSetup"`
	Hash                     *string      `json:"hash,omitempty"`
}

// NewVersion instantiates a new Version object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersion(cid string, extId string, createdBy string, teamId string, projectId string, branch string, tag string, data ModelGraph, createdAt time.Time, notes string, status string, isFavourite bool, datasetSetup DatasetSetup) *Version {
	this := Version{}
	this.Cid = cid
	this.ExtId = extId
	this.CreatedBy = createdBy
	this.TeamId = teamId
	this.ProjectId = projectId
	this.Branch = branch
	this.Tag = tag
	this.Data = data
	this.CreatedAt = createdAt
	this.Notes = notes
	this.Status = status
	this.IsFavourite = isFavourite
	this.DatasetSetup = datasetSetup
	return &this
}

// NewVersionWithDefaults instantiates a new Version object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionWithDefaults() *Version {
	this := Version{}
	return &this
}

// GetCid returns the Cid field value
func (o *Version) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *Version) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *Version) SetCid(v string) {
	o.Cid = v
}

// GetExtId returns the ExtId field value
func (o *Version) GetExtId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtId
}

// GetExtIdOk returns a tuple with the ExtId field value
// and a boolean to check if the value has been set.
func (o *Version) GetExtIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtId, true
}

// SetExtId sets field value
func (o *Version) SetExtId(v string) {
	o.ExtId = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Version) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Version) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Version) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetTeamId returns the TeamId field value
func (o *Version) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *Version) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *Version) SetTeamId(v string) {
	o.TeamId = v
}

// GetProjectId returns the ProjectId field value
func (o *Version) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Version) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Version) SetProjectId(v string) {
	o.ProjectId = v
}

// GetBranch returns the Branch field value
func (o *Version) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *Version) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *Version) SetBranch(v string) {
	o.Branch = v
}

// GetTag returns the Tag field value
func (o *Version) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *Version) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *Version) SetTag(v string) {
	o.Tag = v
}

// GetData returns the Data field value
func (o *Version) GetData() ModelGraph {
	if o == nil {
		var ret ModelGraph
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Version) GetDataOk() (*ModelGraph, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Version) SetData(v ModelGraph) {
	o.Data = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Version) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Version) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Version) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetNotes returns the Notes field value
func (o *Version) GetNotes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *Version) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notes, true
}

// SetNotes sets field value
func (o *Version) SetNotes(v string) {
	o.Notes = v
}

// GetStatus returns the Status field value
func (o *Version) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Version) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Version) SetStatus(v string) {
	o.Status = v
}

// GetIsFavourite returns the IsFavourite field value
func (o *Version) GetIsFavourite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFavourite
}

// GetIsFavouriteOk returns a tuple with the IsFavourite field value
// and a boolean to check if the value has been set.
func (o *Version) GetIsFavouriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFavourite, true
}

// SetIsFavourite sets field value
func (o *Version) SetIsFavourite(v bool) {
	o.IsFavourite = v
}

// GetCodeIntegrationVersionId returns the CodeIntegrationVersionId field value if set, zero value otherwise.
func (o *Version) GetCodeIntegrationVersionId() string {
	if o == nil || IsNil(o.CodeIntegrationVersionId) {
		var ret string
		return ret
	}
	return *o.CodeIntegrationVersionId
}

// GetCodeIntegrationVersionIdOk returns a tuple with the CodeIntegrationVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetCodeIntegrationVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CodeIntegrationVersionId) {
		return nil, false
	}
	return o.CodeIntegrationVersionId, true
}

// HasCodeIntegrationVersionId returns a boolean if a field has been set.
func (o *Version) HasCodeIntegrationVersionId() bool {
	if o != nil && !IsNil(o.CodeIntegrationVersionId) {
		return true
	}

	return false
}

// SetCodeIntegrationVersionId gets a reference to the given string and assigns it to the CodeIntegrationVersionId field.
func (o *Version) SetCodeIntegrationVersionId(v string) {
	o.CodeIntegrationVersionId = &v
}

// GetDatasetSetup returns the DatasetSetup field value
func (o *Version) GetDatasetSetup() DatasetSetup {
	if o == nil {
		var ret DatasetSetup
		return ret
	}

	return o.DatasetSetup
}

// GetDatasetSetupOk returns a tuple with the DatasetSetup field value
// and a boolean to check if the value has been set.
func (o *Version) GetDatasetSetupOk() (*DatasetSetup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetSetup, true
}

// SetDatasetSetup sets field value
func (o *Version) SetDatasetSetup(v DatasetSetup) {
	o.DatasetSetup = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *Version) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *Version) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *Version) SetHash(v string) {
	o.Hash = &v
}

func (o Version) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Version) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["extId"] = o.ExtId
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["teamId"] = o.TeamId
	toSerialize["projectId"] = o.ProjectId
	toSerialize["branch"] = o.Branch
	toSerialize["tag"] = o.Tag
	toSerialize["data"] = o.Data
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["notes"] = o.Notes
	toSerialize["status"] = o.Status
	toSerialize["isFavourite"] = o.IsFavourite
	if !IsNil(o.CodeIntegrationVersionId) {
		toSerialize["codeIntegrationVersionId"] = o.CodeIntegrationVersionId
	}
	toSerialize["datasetSetup"] = o.DatasetSetup
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	return toSerialize, nil
}

type NullableVersion struct {
	value *Version
	isSet bool
}

func (v NullableVersion) Get() *Version {
	return v.value
}

func (v *NullableVersion) Set(val *Version) {
	v.value = val
	v.isSet = true
}

func (v NullableVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersion(val *Version) *NullableVersion {
	return &NullableVersion{value: val, isSet: true}
}

func (v NullableVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
