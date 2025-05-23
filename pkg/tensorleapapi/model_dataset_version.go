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

// checks if the DatasetVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetVersion{}

// DatasetVersion struct for DatasetVersion
type DatasetVersion struct {
	Cid                  string          `json:"cid"`
	TeamId               string          `json:"teamId"`
	DatasetId            string          `json:"datasetId"`
	Note                 string          `json:"note"`
	CreatedAt            string          `json:"createdAt"`
	CreatedBy            string          `json:"createdBy"`
	TestStatus           TestStatus      `json:"testStatus"`
	Metadata             DatasetMetadata `json:"metadata"`
	BlobPath             string          `json:"blobPath"`
	Branch               string          `json:"branch"`
	CodeEntryFile        string          `json:"codeEntryFile"`
	GenericBaseImageType *string         `json:"genericBaseImageType,omitempty"`
}

// NewDatasetVersion instantiates a new DatasetVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetVersion(cid string, teamId string, datasetId string, note string, createdAt string, createdBy string, testStatus TestStatus, metadata DatasetMetadata, blobPath string, branch string, codeEntryFile string) *DatasetVersion {
	this := DatasetVersion{}
	this.Cid = cid
	this.TeamId = teamId
	this.DatasetId = datasetId
	this.Note = note
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.TestStatus = testStatus
	this.Metadata = metadata
	this.BlobPath = blobPath
	this.Branch = branch
	this.CodeEntryFile = codeEntryFile
	return &this
}

// NewDatasetVersionWithDefaults instantiates a new DatasetVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetVersionWithDefaults() *DatasetVersion {
	this := DatasetVersion{}
	return &this
}

// GetCid returns the Cid field value
func (o *DatasetVersion) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *DatasetVersion) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *DatasetVersion) SetCid(v string) {
	o.Cid = v
}

// GetTeamId returns the TeamId field value
func (o *DatasetVersion) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *DatasetVersion) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *DatasetVersion) SetTeamId(v string) {
	o.TeamId = v
}

// GetDatasetId returns the DatasetId field value
func (o *DatasetVersion) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *DatasetVersion) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value
func (o *DatasetVersion) SetDatasetId(v string) {
	o.DatasetId = v
}

// GetNote returns the Note field value
func (o *DatasetVersion) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *DatasetVersion) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *DatasetVersion) SetNote(v string) {
	o.Note = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DatasetVersion) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DatasetVersion) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DatasetVersion) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *DatasetVersion) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *DatasetVersion) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *DatasetVersion) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetTestStatus returns the TestStatus field value
func (o *DatasetVersion) GetTestStatus() TestStatus {
	if o == nil {
		var ret TestStatus
		return ret
	}

	return o.TestStatus
}

// GetTestStatusOk returns a tuple with the TestStatus field value
// and a boolean to check if the value has been set.
func (o *DatasetVersion) GetTestStatusOk() (*TestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestStatus, true
}

// SetTestStatus sets field value
func (o *DatasetVersion) SetTestStatus(v TestStatus) {
	o.TestStatus = v
}

// GetMetadata returns the Metadata field value
func (o *DatasetVersion) GetMetadata() DatasetMetadata {
	if o == nil {
		var ret DatasetMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *DatasetVersion) GetMetadataOk() (*DatasetMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *DatasetVersion) SetMetadata(v DatasetMetadata) {
	o.Metadata = v
}

// GetBlobPath returns the BlobPath field value
func (o *DatasetVersion) GetBlobPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlobPath
}

// GetBlobPathOk returns a tuple with the BlobPath field value
// and a boolean to check if the value has been set.
func (o *DatasetVersion) GetBlobPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlobPath, true
}

// SetBlobPath sets field value
func (o *DatasetVersion) SetBlobPath(v string) {
	o.BlobPath = v
}

// GetBranch returns the Branch field value
func (o *DatasetVersion) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *DatasetVersion) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *DatasetVersion) SetBranch(v string) {
	o.Branch = v
}

// GetCodeEntryFile returns the CodeEntryFile field value
func (o *DatasetVersion) GetCodeEntryFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CodeEntryFile
}

// GetCodeEntryFileOk returns a tuple with the CodeEntryFile field value
// and a boolean to check if the value has been set.
func (o *DatasetVersion) GetCodeEntryFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeEntryFile, true
}

// SetCodeEntryFile sets field value
func (o *DatasetVersion) SetCodeEntryFile(v string) {
	o.CodeEntryFile = v
}

// GetGenericBaseImageType returns the GenericBaseImageType field value if set, zero value otherwise.
func (o *DatasetVersion) GetGenericBaseImageType() string {
	if o == nil || IsNil(o.GenericBaseImageType) {
		var ret string
		return ret
	}
	return *o.GenericBaseImageType
}

// GetGenericBaseImageTypeOk returns a tuple with the GenericBaseImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetVersion) GetGenericBaseImageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GenericBaseImageType) {
		return nil, false
	}
	return o.GenericBaseImageType, true
}

// HasGenericBaseImageType returns a boolean if a field has been set.
func (o *DatasetVersion) HasGenericBaseImageType() bool {
	if o != nil && !IsNil(o.GenericBaseImageType) {
		return true
	}

	return false
}

// SetGenericBaseImageType gets a reference to the given string and assigns it to the GenericBaseImageType field.
func (o *DatasetVersion) SetGenericBaseImageType(v string) {
	o.GenericBaseImageType = &v
}

func (o DatasetVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["teamId"] = o.TeamId
	toSerialize["datasetId"] = o.DatasetId
	toSerialize["note"] = o.Note
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["testStatus"] = o.TestStatus
	toSerialize["metadata"] = o.Metadata
	toSerialize["blobPath"] = o.BlobPath
	toSerialize["branch"] = o.Branch
	toSerialize["codeEntryFile"] = o.CodeEntryFile
	if !IsNil(o.GenericBaseImageType) {
		toSerialize["genericBaseImageType"] = o.GenericBaseImageType
	}
	return toSerialize, nil
}

type NullableDatasetVersion struct {
	value *DatasetVersion
	isSet bool
}

func (v NullableDatasetVersion) Get() *DatasetVersion {
	return v.value
}

func (v *NullableDatasetVersion) Set(val *DatasetVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetVersion(val *DatasetVersion) *NullableDatasetVersion {
	return &NullableDatasetVersion{value: val, isSet: true}
}

func (v NullableDatasetVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
