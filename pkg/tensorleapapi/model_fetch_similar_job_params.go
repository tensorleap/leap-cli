/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the FetchSimilarJobParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchSimilarJobParams{}

// FetchSimilarJobParams struct for FetchSimilarJobParams
type FetchSimilarJobParams struct {
	Digest       string           `json:"digest"`
	Filters      []ESFilter       `json:"filters,omitempty"`
	Epoch        float64          `json:"epoch"`
	SampleIds    []SampleIdentity `json:"sampleIds"`
	Limit        float64          `json:"limit"`
	SessionRunId string           `json:"sessionRunId"`
	Type         string           `json:"type"`
}

// NewFetchSimilarJobParams instantiates a new FetchSimilarJobParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchSimilarJobParams(digest string, epoch float64, sampleIds []SampleIdentity, limit float64, sessionRunId string, type_ string) *FetchSimilarJobParams {
	this := FetchSimilarJobParams{}
	this.Digest = digest
	this.Epoch = epoch
	this.SampleIds = sampleIds
	this.Limit = limit
	this.SessionRunId = sessionRunId
	this.Type = type_
	return &this
}

// NewFetchSimilarJobParamsWithDefaults instantiates a new FetchSimilarJobParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchSimilarJobParamsWithDefaults() *FetchSimilarJobParams {
	this := FetchSimilarJobParams{}
	return &this
}

// GetDigest returns the Digest field value
func (o *FetchSimilarJobParams) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarJobParams) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *FetchSimilarJobParams) SetDigest(v string) {
	o.Digest = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *FetchSimilarJobParams) GetFilters() []ESFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []ESFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchSimilarJobParams) GetFiltersOk() ([]ESFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *FetchSimilarJobParams) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []ESFilter and assigns it to the Filters field.
func (o *FetchSimilarJobParams) SetFilters(v []ESFilter) {
	o.Filters = v
}

// GetEpoch returns the Epoch field value
func (o *FetchSimilarJobParams) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarJobParams) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *FetchSimilarJobParams) SetEpoch(v float64) {
	o.Epoch = v
}

// GetSampleIds returns the SampleIds field value
func (o *FetchSimilarJobParams) GetSampleIds() []SampleIdentity {
	if o == nil {
		var ret []SampleIdentity
		return ret
	}

	return o.SampleIds
}

// GetSampleIdsOk returns a tuple with the SampleIds field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarJobParams) GetSampleIdsOk() ([]SampleIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return o.SampleIds, true
}

// SetSampleIds sets field value
func (o *FetchSimilarJobParams) SetSampleIds(v []SampleIdentity) {
	o.SampleIds = v
}

// GetLimit returns the Limit field value
func (o *FetchSimilarJobParams) GetLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarJobParams) GetLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *FetchSimilarJobParams) SetLimit(v float64) {
	o.Limit = v
}

// GetSessionRunId returns the SessionRunId field value
func (o *FetchSimilarJobParams) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarJobParams) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *FetchSimilarJobParams) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetType returns the Type field value
func (o *FetchSimilarJobParams) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarJobParams) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FetchSimilarJobParams) SetType(v string) {
	o.Type = v
}

func (o FetchSimilarJobParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchSimilarJobParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["digest"] = o.Digest
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	toSerialize["epoch"] = o.Epoch
	toSerialize["sampleIds"] = o.SampleIds
	toSerialize["limit"] = o.Limit
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableFetchSimilarJobParams struct {
	value *FetchSimilarJobParams
	isSet bool
}

func (v NullableFetchSimilarJobParams) Get() *FetchSimilarJobParams {
	return v.value
}

func (v *NullableFetchSimilarJobParams) Set(val *FetchSimilarJobParams) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchSimilarJobParams) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchSimilarJobParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchSimilarJobParams(val *FetchSimilarJobParams) *NullableFetchSimilarJobParams {
	return &NullableFetchSimilarJobParams{value: val, isSet: true}
}

func (v NullableFetchSimilarJobParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchSimilarJobParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
