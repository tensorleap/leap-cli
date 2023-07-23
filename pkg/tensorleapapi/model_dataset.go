/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the Dataset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dataset{}

// Dataset struct for Dataset
type Dataset struct {
	Cid                string          `json:"cid"`
	Name               NullableString  `json:"name"`
	TeamId             string          `json:"teamId"`
	SourceDatasetId    *string         `json:"sourceDatasetId,omitempty"`
	Access             DatasetAccess   `json:"access"`
	CreatedAt          string          `json:"createdAt"`
	LatestValidVersion *DatasetVersion `json:"latestValidVersion,omitempty"`
	LatestVersion      *DatasetVersion `json:"latestVersion,omitempty"`
}

// NewDataset instantiates a new Dataset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataset(cid string, name NullableString, teamId string, access DatasetAccess, createdAt string) *Dataset {
	this := Dataset{}
	this.Cid = cid
	this.Name = name
	this.TeamId = teamId
	this.Access = access
	this.CreatedAt = createdAt
	return &this
}

// NewDatasetWithDefaults instantiates a new Dataset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetWithDefaults() *Dataset {
	this := Dataset{}
	return &this
}

// GetCid returns the Cid field value
func (o *Dataset) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *Dataset) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *Dataset) SetCid(v string) {
	o.Cid = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Dataset) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Dataset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *Dataset) SetName(v string) {
	o.Name.Set(&v)
}

// GetTeamId returns the TeamId field value
func (o *Dataset) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *Dataset) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *Dataset) SetTeamId(v string) {
	o.TeamId = v
}

// GetSourceDatasetId returns the SourceDatasetId field value if set, zero value otherwise.
func (o *Dataset) GetSourceDatasetId() string {
	if o == nil || IsNil(o.SourceDatasetId) {
		var ret string
		return ret
	}
	return *o.SourceDatasetId
}

// GetSourceDatasetIdOk returns a tuple with the SourceDatasetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dataset) GetSourceDatasetIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceDatasetId) {
		return nil, false
	}
	return o.SourceDatasetId, true
}

// HasSourceDatasetId returns a boolean if a field has been set.
func (o *Dataset) HasSourceDatasetId() bool {
	if o != nil && !IsNil(o.SourceDatasetId) {
		return true
	}

	return false
}

// SetSourceDatasetId gets a reference to the given string and assigns it to the SourceDatasetId field.
func (o *Dataset) SetSourceDatasetId(v string) {
	o.SourceDatasetId = &v
}

// GetAccess returns the Access field value
func (o *Dataset) GetAccess() DatasetAccess {
	if o == nil {
		var ret DatasetAccess
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *Dataset) GetAccessOk() (*DatasetAccess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *Dataset) SetAccess(v DatasetAccess) {
	o.Access = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Dataset) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Dataset) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Dataset) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetLatestValidVersion returns the LatestValidVersion field value if set, zero value otherwise.
func (o *Dataset) GetLatestValidVersion() DatasetVersion {
	if o == nil || IsNil(o.LatestValidVersion) {
		var ret DatasetVersion
		return ret
	}
	return *o.LatestValidVersion
}

// GetLatestValidVersionOk returns a tuple with the LatestValidVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dataset) GetLatestValidVersionOk() (*DatasetVersion, bool) {
	if o == nil || IsNil(o.LatestValidVersion) {
		return nil, false
	}
	return o.LatestValidVersion, true
}

// HasLatestValidVersion returns a boolean if a field has been set.
func (o *Dataset) HasLatestValidVersion() bool {
	if o != nil && !IsNil(o.LatestValidVersion) {
		return true
	}

	return false
}

// SetLatestValidVersion gets a reference to the given DatasetVersion and assigns it to the LatestValidVersion field.
func (o *Dataset) SetLatestValidVersion(v DatasetVersion) {
	o.LatestValidVersion = &v
}

// GetLatestVersion returns the LatestVersion field value if set, zero value otherwise.
func (o *Dataset) GetLatestVersion() DatasetVersion {
	if o == nil || IsNil(o.LatestVersion) {
		var ret DatasetVersion
		return ret
	}
	return *o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dataset) GetLatestVersionOk() (*DatasetVersion, bool) {
	if o == nil || IsNil(o.LatestVersion) {
		return nil, false
	}
	return o.LatestVersion, true
}

// HasLatestVersion returns a boolean if a field has been set.
func (o *Dataset) HasLatestVersion() bool {
	if o != nil && !IsNil(o.LatestVersion) {
		return true
	}

	return false
}

// SetLatestVersion gets a reference to the given DatasetVersion and assigns it to the LatestVersion field.
func (o *Dataset) SetLatestVersion(v DatasetVersion) {
	o.LatestVersion = &v
}

func (o Dataset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dataset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["name"] = o.Name.Get()
	toSerialize["teamId"] = o.TeamId
	if !IsNil(o.SourceDatasetId) {
		toSerialize["sourceDatasetId"] = o.SourceDatasetId
	}
	toSerialize["access"] = o.Access
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.LatestValidVersion) {
		toSerialize["latestValidVersion"] = o.LatestValidVersion
	}
	if !IsNil(o.LatestVersion) {
		toSerialize["latestVersion"] = o.LatestVersion
	}
	return toSerialize, nil
}

type NullableDataset struct {
	value *Dataset
	isSet bool
}

func (v NullableDataset) Get() *Dataset {
	return v.value
}

func (v *NullableDataset) Set(val *Dataset) {
	v.value = val
	v.isSet = true
}

func (v NullableDataset) IsSet() bool {
	return v.isSet
}

func (v *NullableDataset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataset(val *Dataset) *NullableDataset {
	return &NullableDataset{value: val, isSet: true}
}

func (v NullableDataset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
