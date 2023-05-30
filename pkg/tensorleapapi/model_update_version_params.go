/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.321
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UpdateVersionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVersionParams{}

// UpdateVersionParams struct for UpdateVersionParams
type UpdateVersionParams struct {
	VersionId                string       `json:"versionId"`
	Data                     ModelGraph   `json:"data"`
	CodeIntegrationVersionId *string      `json:"codeIntegrationVersionId,omitempty"`
	DatasetSetup             DatasetSetup `json:"datasetSetup"`
	OverrideHash             bool         `json:"overrideHash"`
	Hash                     *string      `json:"hash,omitempty"`
}

// NewUpdateVersionParams instantiates a new UpdateVersionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVersionParams(versionId string, data ModelGraph, datasetSetup DatasetSetup, overrideHash bool) *UpdateVersionParams {
	this := UpdateVersionParams{}
	this.VersionId = versionId
	this.Data = data
	this.DatasetSetup = datasetSetup
	this.OverrideHash = overrideHash
	return &this
}

// NewUpdateVersionParamsWithDefaults instantiates a new UpdateVersionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVersionParamsWithDefaults() *UpdateVersionParams {
	this := UpdateVersionParams{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *UpdateVersionParams) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *UpdateVersionParams) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *UpdateVersionParams) SetVersionId(v string) {
	o.VersionId = v
}

// GetData returns the Data field value
func (o *UpdateVersionParams) GetData() ModelGraph {
	if o == nil {
		var ret ModelGraph
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateVersionParams) GetDataOk() (*ModelGraph, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateVersionParams) SetData(v ModelGraph) {
	o.Data = v
}

// GetCodeIntegrationVersionId returns the CodeIntegrationVersionId field value if set, zero value otherwise.
func (o *UpdateVersionParams) GetCodeIntegrationVersionId() string {
	if o == nil || IsNil(o.CodeIntegrationVersionId) {
		var ret string
		return ret
	}
	return *o.CodeIntegrationVersionId
}

// GetCodeIntegrationVersionIdOk returns a tuple with the CodeIntegrationVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVersionParams) GetCodeIntegrationVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CodeIntegrationVersionId) {
		return nil, false
	}
	return o.CodeIntegrationVersionId, true
}

// HasCodeIntegrationVersionId returns a boolean if a field has been set.
func (o *UpdateVersionParams) HasCodeIntegrationVersionId() bool {
	if o != nil && !IsNil(o.CodeIntegrationVersionId) {
		return true
	}

	return false
}

// SetCodeIntegrationVersionId gets a reference to the given string and assigns it to the CodeIntegrationVersionId field.
func (o *UpdateVersionParams) SetCodeIntegrationVersionId(v string) {
	o.CodeIntegrationVersionId = &v
}

// GetDatasetSetup returns the DatasetSetup field value
func (o *UpdateVersionParams) GetDatasetSetup() DatasetSetup {
	if o == nil {
		var ret DatasetSetup
		return ret
	}

	return o.DatasetSetup
}

// GetDatasetSetupOk returns a tuple with the DatasetSetup field value
// and a boolean to check if the value has been set.
func (o *UpdateVersionParams) GetDatasetSetupOk() (*DatasetSetup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetSetup, true
}

// SetDatasetSetup sets field value
func (o *UpdateVersionParams) SetDatasetSetup(v DatasetSetup) {
	o.DatasetSetup = v
}

// GetOverrideHash returns the OverrideHash field value
func (o *UpdateVersionParams) GetOverrideHash() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OverrideHash
}

// GetOverrideHashOk returns a tuple with the OverrideHash field value
// and a boolean to check if the value has been set.
func (o *UpdateVersionParams) GetOverrideHashOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverrideHash, true
}

// SetOverrideHash sets field value
func (o *UpdateVersionParams) SetOverrideHash(v bool) {
	o.OverrideHash = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *UpdateVersionParams) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVersionParams) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *UpdateVersionParams) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *UpdateVersionParams) SetHash(v string) {
	o.Hash = &v
}

func (o UpdateVersionParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVersionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	toSerialize["data"] = o.Data
	if !IsNil(o.CodeIntegrationVersionId) {
		toSerialize["codeIntegrationVersionId"] = o.CodeIntegrationVersionId
	}
	toSerialize["datasetSetup"] = o.DatasetSetup
	toSerialize["overrideHash"] = o.OverrideHash
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	return toSerialize, nil
}

type NullableUpdateVersionParams struct {
	value *UpdateVersionParams
	isSet bool
}

func (v NullableUpdateVersionParams) Get() *UpdateVersionParams {
	return v.value
}

func (v *NullableUpdateVersionParams) Set(val *UpdateVersionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVersionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVersionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVersionParams(val *UpdateVersionParams) *NullableUpdateVersionParams {
	return &NullableUpdateVersionParams{value: val, isSet: true}
}

func (v NullableUpdateVersionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVersionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
