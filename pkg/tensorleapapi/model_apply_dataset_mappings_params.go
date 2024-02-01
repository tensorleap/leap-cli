/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.445
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ApplyDatasetMappingsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplyDatasetMappingsParams{}

// ApplyDatasetMappingsParams struct for ApplyDatasetMappingsParams
type ApplyDatasetMappingsParams struct {
	ModelGraph       ModelGraph `json:"modelGraph"`
	Yaml             string     `json:"yaml"`
	DatasetVersionId *string    `json:"datasetVersionId,omitempty"`
}

// NewApplyDatasetMappingsParams instantiates a new ApplyDatasetMappingsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyDatasetMappingsParams(modelGraph ModelGraph, yaml string) *ApplyDatasetMappingsParams {
	this := ApplyDatasetMappingsParams{}
	this.ModelGraph = modelGraph
	this.Yaml = yaml
	return &this
}

// NewApplyDatasetMappingsParamsWithDefaults instantiates a new ApplyDatasetMappingsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyDatasetMappingsParamsWithDefaults() *ApplyDatasetMappingsParams {
	this := ApplyDatasetMappingsParams{}
	return &this
}

// GetModelGraph returns the ModelGraph field value
func (o *ApplyDatasetMappingsParams) GetModelGraph() ModelGraph {
	if o == nil {
		var ret ModelGraph
		return ret
	}

	return o.ModelGraph
}

// GetModelGraphOk returns a tuple with the ModelGraph field value
// and a boolean to check if the value has been set.
func (o *ApplyDatasetMappingsParams) GetModelGraphOk() (*ModelGraph, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelGraph, true
}

// SetModelGraph sets field value
func (o *ApplyDatasetMappingsParams) SetModelGraph(v ModelGraph) {
	o.ModelGraph = v
}

// GetYaml returns the Yaml field value
func (o *ApplyDatasetMappingsParams) GetYaml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Yaml
}

// GetYamlOk returns a tuple with the Yaml field value
// and a boolean to check if the value has been set.
func (o *ApplyDatasetMappingsParams) GetYamlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Yaml, true
}

// SetYaml sets field value
func (o *ApplyDatasetMappingsParams) SetYaml(v string) {
	o.Yaml = v
}

// GetDatasetVersionId returns the DatasetVersionId field value if set, zero value otherwise.
func (o *ApplyDatasetMappingsParams) GetDatasetVersionId() string {
	if o == nil || IsNil(o.DatasetVersionId) {
		var ret string
		return ret
	}
	return *o.DatasetVersionId
}

// GetDatasetVersionIdOk returns a tuple with the DatasetVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyDatasetMappingsParams) GetDatasetVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetVersionId) {
		return nil, false
	}
	return o.DatasetVersionId, true
}

// HasDatasetVersionId returns a boolean if a field has been set.
func (o *ApplyDatasetMappingsParams) HasDatasetVersionId() bool {
	if o != nil && !IsNil(o.DatasetVersionId) {
		return true
	}

	return false
}

// SetDatasetVersionId gets a reference to the given string and assigns it to the DatasetVersionId field.
func (o *ApplyDatasetMappingsParams) SetDatasetVersionId(v string) {
	o.DatasetVersionId = &v
}

func (o ApplyDatasetMappingsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyDatasetMappingsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["modelGraph"] = o.ModelGraph
	toSerialize["yaml"] = o.Yaml
	if !IsNil(o.DatasetVersionId) {
		toSerialize["datasetVersionId"] = o.DatasetVersionId
	}
	return toSerialize, nil
}

type NullableApplyDatasetMappingsParams struct {
	value *ApplyDatasetMappingsParams
	isSet bool
}

func (v NullableApplyDatasetMappingsParams) Get() *ApplyDatasetMappingsParams {
	return v.value
}

func (v *NullableApplyDatasetMappingsParams) Set(val *ApplyDatasetMappingsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyDatasetMappingsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyDatasetMappingsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyDatasetMappingsParams(val *ApplyDatasetMappingsParams) *NullableApplyDatasetMappingsParams {
	return &NullableApplyDatasetMappingsParams{value: val, isSet: true}
}

func (v NullableApplyDatasetMappingsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyDatasetMappingsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
