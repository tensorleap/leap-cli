/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.557
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the FetchSimilarFilterDisplayData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchSimilarFilterDisplayData{}

// FetchSimilarFilterDisplayData struct for FetchSimilarFilterDisplayData
type FetchSimilarFilterDisplayData struct {
	Type        string           `json:"type"`
	Limit       float64          `json:"limit"`
	SampleIds   []SampleIdentity `json:"sampleIds"`
	Epoch       float64          `json:"epoch"`
	FiltersUsed []string         `json:"filtersUsed"`
	SessionRun  FilterSessionRun `json:"sessionRun"`
}

// NewFetchSimilarFilterDisplayData instantiates a new FetchSimilarFilterDisplayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchSimilarFilterDisplayData(type_ string, limit float64, sampleIds []SampleIdentity, epoch float64, filtersUsed []string, sessionRun FilterSessionRun) *FetchSimilarFilterDisplayData {
	this := FetchSimilarFilterDisplayData{}
	this.Type = type_
	this.Limit = limit
	this.SampleIds = sampleIds
	this.Epoch = epoch
	this.FiltersUsed = filtersUsed
	this.SessionRun = sessionRun
	return &this
}

// NewFetchSimilarFilterDisplayDataWithDefaults instantiates a new FetchSimilarFilterDisplayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchSimilarFilterDisplayDataWithDefaults() *FetchSimilarFilterDisplayData {
	this := FetchSimilarFilterDisplayData{}
	return &this
}

// GetType returns the Type field value
func (o *FetchSimilarFilterDisplayData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarFilterDisplayData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FetchSimilarFilterDisplayData) SetType(v string) {
	o.Type = v
}

// GetLimit returns the Limit field value
func (o *FetchSimilarFilterDisplayData) GetLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarFilterDisplayData) GetLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *FetchSimilarFilterDisplayData) SetLimit(v float64) {
	o.Limit = v
}

// GetSampleIds returns the SampleIds field value
func (o *FetchSimilarFilterDisplayData) GetSampleIds() []SampleIdentity {
	if o == nil {
		var ret []SampleIdentity
		return ret
	}

	return o.SampleIds
}

// GetSampleIdsOk returns a tuple with the SampleIds field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarFilterDisplayData) GetSampleIdsOk() ([]SampleIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return o.SampleIds, true
}

// SetSampleIds sets field value
func (o *FetchSimilarFilterDisplayData) SetSampleIds(v []SampleIdentity) {
	o.SampleIds = v
}

// GetEpoch returns the Epoch field value
func (o *FetchSimilarFilterDisplayData) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarFilterDisplayData) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *FetchSimilarFilterDisplayData) SetEpoch(v float64) {
	o.Epoch = v
}

// GetFiltersUsed returns the FiltersUsed field value
func (o *FetchSimilarFilterDisplayData) GetFiltersUsed() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FiltersUsed
}

// GetFiltersUsedOk returns a tuple with the FiltersUsed field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarFilterDisplayData) GetFiltersUsedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FiltersUsed, true
}

// SetFiltersUsed sets field value
func (o *FetchSimilarFilterDisplayData) SetFiltersUsed(v []string) {
	o.FiltersUsed = v
}

// GetSessionRun returns the SessionRun field value
func (o *FetchSimilarFilterDisplayData) GetSessionRun() FilterSessionRun {
	if o == nil {
		var ret FilterSessionRun
		return ret
	}

	return o.SessionRun
}

// GetSessionRunOk returns a tuple with the SessionRun field value
// and a boolean to check if the value has been set.
func (o *FetchSimilarFilterDisplayData) GetSessionRunOk() (*FilterSessionRun, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRun, true
}

// SetSessionRun sets field value
func (o *FetchSimilarFilterDisplayData) SetSessionRun(v FilterSessionRun) {
	o.SessionRun = v
}

func (o FetchSimilarFilterDisplayData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchSimilarFilterDisplayData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["limit"] = o.Limit
	toSerialize["sampleIds"] = o.SampleIds
	toSerialize["epoch"] = o.Epoch
	toSerialize["filtersUsed"] = o.FiltersUsed
	toSerialize["sessionRun"] = o.SessionRun
	return toSerialize, nil
}

type NullableFetchSimilarFilterDisplayData struct {
	value *FetchSimilarFilterDisplayData
	isSet bool
}

func (v NullableFetchSimilarFilterDisplayData) Get() *FetchSimilarFilterDisplayData {
	return v.value
}

func (v *NullableFetchSimilarFilterDisplayData) Set(val *FetchSimilarFilterDisplayData) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchSimilarFilterDisplayData) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchSimilarFilterDisplayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchSimilarFilterDisplayData(val *FetchSimilarFilterDisplayData) *NullableFetchSimilarFilterDisplayData {
	return &NullableFetchSimilarFilterDisplayData{value: val, isSet: true}
}

func (v NullableFetchSimilarFilterDisplayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchSimilarFilterDisplayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
