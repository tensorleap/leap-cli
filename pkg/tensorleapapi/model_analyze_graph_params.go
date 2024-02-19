/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.464
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the AnalyzeGraphParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyzeGraphParams{}

// AnalyzeGraphParams struct for AnalyzeGraphParams
type AnalyzeGraphParams struct {
	Graph            ModelGraph `json:"graph"`
	DatasetVersionId *string    `json:"datasetVersionId,omitempty"`
	VersionId        string     `json:"versionId"`
	ProjectId        string     `json:"projectId"`
	Digest           string     `json:"digest"`
}

// NewAnalyzeGraphParams instantiates a new AnalyzeGraphParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyzeGraphParams(graph ModelGraph, versionId string, projectId string, digest string) *AnalyzeGraphParams {
	this := AnalyzeGraphParams{}
	this.Graph = graph
	this.VersionId = versionId
	this.ProjectId = projectId
	this.Digest = digest
	return &this
}

// NewAnalyzeGraphParamsWithDefaults instantiates a new AnalyzeGraphParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyzeGraphParamsWithDefaults() *AnalyzeGraphParams {
	this := AnalyzeGraphParams{}
	return &this
}

// GetGraph returns the Graph field value
func (o *AnalyzeGraphParams) GetGraph() ModelGraph {
	if o == nil {
		var ret ModelGraph
		return ret
	}

	return o.Graph
}

// GetGraphOk returns a tuple with the Graph field value
// and a boolean to check if the value has been set.
func (o *AnalyzeGraphParams) GetGraphOk() (*ModelGraph, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Graph, true
}

// SetGraph sets field value
func (o *AnalyzeGraphParams) SetGraph(v ModelGraph) {
	o.Graph = v
}

// GetDatasetVersionId returns the DatasetVersionId field value if set, zero value otherwise.
func (o *AnalyzeGraphParams) GetDatasetVersionId() string {
	if o == nil || IsNil(o.DatasetVersionId) {
		var ret string
		return ret
	}
	return *o.DatasetVersionId
}

// GetDatasetVersionIdOk returns a tuple with the DatasetVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyzeGraphParams) GetDatasetVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetVersionId) {
		return nil, false
	}
	return o.DatasetVersionId, true
}

// HasDatasetVersionId returns a boolean if a field has been set.
func (o *AnalyzeGraphParams) HasDatasetVersionId() bool {
	if o != nil && !IsNil(o.DatasetVersionId) {
		return true
	}

	return false
}

// SetDatasetVersionId gets a reference to the given string and assigns it to the DatasetVersionId field.
func (o *AnalyzeGraphParams) SetDatasetVersionId(v string) {
	o.DatasetVersionId = &v
}

// GetVersionId returns the VersionId field value
func (o *AnalyzeGraphParams) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *AnalyzeGraphParams) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *AnalyzeGraphParams) SetVersionId(v string) {
	o.VersionId = v
}

// GetProjectId returns the ProjectId field value
func (o *AnalyzeGraphParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AnalyzeGraphParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AnalyzeGraphParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetDigest returns the Digest field value
func (o *AnalyzeGraphParams) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *AnalyzeGraphParams) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *AnalyzeGraphParams) SetDigest(v string) {
	o.Digest = v
}

func (o AnalyzeGraphParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyzeGraphParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["graph"] = o.Graph
	if !IsNil(o.DatasetVersionId) {
		toSerialize["datasetVersionId"] = o.DatasetVersionId
	}
	toSerialize["versionId"] = o.VersionId
	toSerialize["projectId"] = o.ProjectId
	toSerialize["digest"] = o.Digest
	return toSerialize, nil
}

type NullableAnalyzeGraphParams struct {
	value *AnalyzeGraphParams
	isSet bool
}

func (v NullableAnalyzeGraphParams) Get() *AnalyzeGraphParams {
	return v.value
}

func (v *NullableAnalyzeGraphParams) Set(val *AnalyzeGraphParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyzeGraphParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyzeGraphParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyzeGraphParams(val *AnalyzeGraphParams) *NullableAnalyzeGraphParams {
	return &NullableAnalyzeGraphParams{value: val, isSet: true}
}

func (v NullableAnalyzeGraphParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyzeGraphParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
