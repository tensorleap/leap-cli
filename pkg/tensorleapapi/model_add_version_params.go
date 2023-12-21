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

// checks if the AddVersionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddVersionParams{}

// AddVersionParams struct for AddVersionParams
type AddVersionParams struct {
	ProjectId                string       `json:"projectId"`
	ModelGraph               ModelGraph   `json:"modelGraph"`
	BranchName               string       `json:"branchName"`
	Description              string       `json:"description"`
	CodeIntegrationVersionId *string      `json:"codeIntegrationVersionId,omitempty"`
	DatasetSetup             DatasetSetup `json:"datasetSetup"`
	Hash                     *string      `json:"hash,omitempty"`
	CopySessionIds           []string     `json:"copySessionIds,omitempty"`
}

// NewAddVersionParams instantiates a new AddVersionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVersionParams(projectId string, modelGraph ModelGraph, branchName string, description string, datasetSetup DatasetSetup) *AddVersionParams {
	this := AddVersionParams{}
	this.ProjectId = projectId
	this.ModelGraph = modelGraph
	this.BranchName = branchName
	this.Description = description
	this.DatasetSetup = datasetSetup
	return &this
}

// NewAddVersionParamsWithDefaults instantiates a new AddVersionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVersionParamsWithDefaults() *AddVersionParams {
	this := AddVersionParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *AddVersionParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AddVersionParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AddVersionParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetModelGraph returns the ModelGraph field value
func (o *AddVersionParams) GetModelGraph() ModelGraph {
	if o == nil {
		var ret ModelGraph
		return ret
	}

	return o.ModelGraph
}

// GetModelGraphOk returns a tuple with the ModelGraph field value
// and a boolean to check if the value has been set.
func (o *AddVersionParams) GetModelGraphOk() (*ModelGraph, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelGraph, true
}

// SetModelGraph sets field value
func (o *AddVersionParams) SetModelGraph(v ModelGraph) {
	o.ModelGraph = v
}

// GetBranchName returns the BranchName field value
func (o *AddVersionParams) GetBranchName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value
// and a boolean to check if the value has been set.
func (o *AddVersionParams) GetBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchName, true
}

// SetBranchName sets field value
func (o *AddVersionParams) SetBranchName(v string) {
	o.BranchName = v
}

// GetDescription returns the Description field value
func (o *AddVersionParams) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AddVersionParams) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AddVersionParams) SetDescription(v string) {
	o.Description = v
}

// GetCodeIntegrationVersionId returns the CodeIntegrationVersionId field value if set, zero value otherwise.
func (o *AddVersionParams) GetCodeIntegrationVersionId() string {
	if o == nil || IsNil(o.CodeIntegrationVersionId) {
		var ret string
		return ret
	}
	return *o.CodeIntegrationVersionId
}

// GetCodeIntegrationVersionIdOk returns a tuple with the CodeIntegrationVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVersionParams) GetCodeIntegrationVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CodeIntegrationVersionId) {
		return nil, false
	}
	return o.CodeIntegrationVersionId, true
}

// HasCodeIntegrationVersionId returns a boolean if a field has been set.
func (o *AddVersionParams) HasCodeIntegrationVersionId() bool {
	if o != nil && !IsNil(o.CodeIntegrationVersionId) {
		return true
	}

	return false
}

// SetCodeIntegrationVersionId gets a reference to the given string and assigns it to the CodeIntegrationVersionId field.
func (o *AddVersionParams) SetCodeIntegrationVersionId(v string) {
	o.CodeIntegrationVersionId = &v
}

// GetDatasetSetup returns the DatasetSetup field value
func (o *AddVersionParams) GetDatasetSetup() DatasetSetup {
	if o == nil {
		var ret DatasetSetup
		return ret
	}

	return o.DatasetSetup
}

// GetDatasetSetupOk returns a tuple with the DatasetSetup field value
// and a boolean to check if the value has been set.
func (o *AddVersionParams) GetDatasetSetupOk() (*DatasetSetup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetSetup, true
}

// SetDatasetSetup sets field value
func (o *AddVersionParams) SetDatasetSetup(v DatasetSetup) {
	o.DatasetSetup = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *AddVersionParams) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVersionParams) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *AddVersionParams) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *AddVersionParams) SetHash(v string) {
	o.Hash = &v
}

// GetCopySessionIds returns the CopySessionIds field value if set, zero value otherwise.
func (o *AddVersionParams) GetCopySessionIds() []string {
	if o == nil || IsNil(o.CopySessionIds) {
		var ret []string
		return ret
	}
	return o.CopySessionIds
}

// GetCopySessionIdsOk returns a tuple with the CopySessionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVersionParams) GetCopySessionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CopySessionIds) {
		return nil, false
	}
	return o.CopySessionIds, true
}

// HasCopySessionIds returns a boolean if a field has been set.
func (o *AddVersionParams) HasCopySessionIds() bool {
	if o != nil && !IsNil(o.CopySessionIds) {
		return true
	}

	return false
}

// SetCopySessionIds gets a reference to the given []string and assigns it to the CopySessionIds field.
func (o *AddVersionParams) SetCopySessionIds(v []string) {
	o.CopySessionIds = v
}

func (o AddVersionParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddVersionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["modelGraph"] = o.ModelGraph
	toSerialize["branchName"] = o.BranchName
	toSerialize["description"] = o.Description
	if !IsNil(o.CodeIntegrationVersionId) {
		toSerialize["codeIntegrationVersionId"] = o.CodeIntegrationVersionId
	}
	toSerialize["datasetSetup"] = o.DatasetSetup
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.CopySessionIds) {
		toSerialize["copySessionIds"] = o.CopySessionIds
	}
	return toSerialize, nil
}

type NullableAddVersionParams struct {
	value *AddVersionParams
	isSet bool
}

func (v NullableAddVersionParams) Get() *AddVersionParams {
	return v.value
}

func (v *NullableAddVersionParams) Set(val *AddVersionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVersionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVersionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVersionParams(val *AddVersionParams) *NullableAddVersionParams {
	return &NullableAddVersionParams{value: val, isSet: true}
}

func (v NullableAddVersionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVersionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
