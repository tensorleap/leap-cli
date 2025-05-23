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

// checks if the FetchSimilarRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchSimilarRequestParams{}

// FetchSimilarRequestParams struct for FetchSimilarRequestParams
type FetchSimilarRequestParams struct {
	SessionRunId string           `json:"sessionRunId"`
	ProjectId    string           `json:"projectId"`
	Limit        float64          `json:"limit"`
	SampleIds    []SampleIdentity `json:"sampleIds"`
	Epoch        float64          `json:"epoch"`
	Filters      []ESFilter       `json:"filters,omitempty"`
	Digest       string           `json:"digest"`
}

// NewFetchSimilarRequestParams instantiates a new FetchSimilarRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchSimilarRequestParams(sessionRunId string, projectId string, limit float64, sampleIds []SampleIdentity, epoch float64, digest string) *FetchSimilarRequestParams {
	this := FetchSimilarRequestParams{}
	this.SessionRunId = sessionRunId
	this.ProjectId = projectId
	this.Limit = limit
	this.SampleIds = sampleIds
	this.Epoch = epoch
	this.Digest = digest
	return &this
}

// NewFetchSimilarRequestParamsWithDefaults instantiates a new FetchSimilarRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchSimilarRequestParamsWithDefaults() *FetchSimilarRequestParams {
	this := FetchSimilarRequestParams{}
	return &this
}

// GetSessionRunId returns the SessionRunId field value
func (o *FetchSimilarRequestParams) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarRequestParams) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *FetchSimilarRequestParams) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetProjectId returns the ProjectId field value
func (o *FetchSimilarRequestParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarRequestParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *FetchSimilarRequestParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetLimit returns the Limit field value
func (o *FetchSimilarRequestParams) GetLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarRequestParams) GetLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *FetchSimilarRequestParams) SetLimit(v float64) {
	o.Limit = v
}

// GetSampleIds returns the SampleIds field value
func (o *FetchSimilarRequestParams) GetSampleIds() []SampleIdentity {
	if o == nil {
		var ret []SampleIdentity
		return ret
	}

	return o.SampleIds
}

// GetSampleIdsOk returns a tuple with the SampleIds field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarRequestParams) GetSampleIdsOk() ([]SampleIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return o.SampleIds, true
}

// SetSampleIds sets field value
func (o *FetchSimilarRequestParams) SetSampleIds(v []SampleIdentity) {
	o.SampleIds = v
}

// GetEpoch returns the Epoch field value
func (o *FetchSimilarRequestParams) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarRequestParams) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *FetchSimilarRequestParams) SetEpoch(v float64) {
	o.Epoch = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *FetchSimilarRequestParams) GetFilters() []ESFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []ESFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchSimilarRequestParams) GetFiltersOk() ([]ESFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *FetchSimilarRequestParams) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []ESFilter and assigns it to the Filters field.
func (o *FetchSimilarRequestParams) SetFilters(v []ESFilter) {
	o.Filters = v
}

// GetDigest returns the Digest field value
func (o *FetchSimilarRequestParams) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarRequestParams) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *FetchSimilarRequestParams) SetDigest(v string) {
	o.Digest = v
}

func (o FetchSimilarRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchSimilarRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["projectId"] = o.ProjectId
	toSerialize["limit"] = o.Limit
	toSerialize["sampleIds"] = o.SampleIds
	toSerialize["epoch"] = o.Epoch
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	toSerialize["digest"] = o.Digest
	return toSerialize, nil
}

type NullableFetchSimilarRequestParams struct {
	value *FetchSimilarRequestParams
	isSet bool
}

func (v NullableFetchSimilarRequestParams) Get() *FetchSimilarRequestParams {
	return v.value
}

func (v *NullableFetchSimilarRequestParams) Set(val *FetchSimilarRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchSimilarRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchSimilarRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchSimilarRequestParams(val *FetchSimilarRequestParams) *NullableFetchSimilarRequestParams {
	return &NullableFetchSimilarRequestParams{value: val, isSet: true}
}

func (v NullableFetchSimilarRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchSimilarRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
