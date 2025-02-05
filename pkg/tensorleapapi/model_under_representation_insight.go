/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.592
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UnderRepresentationInsight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnderRepresentationInsight{}

// UnderRepresentationInsight struct for UnderRepresentationInsight
type UnderRepresentationInsight struct {
	Type                         ScatterInsightType             `json:"type"`
	Filter                       ScatterFilter                  `json:"filter"`
	MutualInfoElements           []MutualInformationElement     `json:"mutual_info_elements,omitempty"`
	BlobPath                     *string                        `json:"blob_path,omitempty"`
	InsightSubCategoryDsCuration *InsightSubCategoryDSCuration  `json:"insight_sub_category_ds_curation,omitempty"`
	InsightCategoryPerformance   *InsightSubCategoryPerformance `json:"insight_category_performance,omitempty"`
	Index                        *float64                       `json:"index,omitempty"`
	UnderRepresentationDataset   DataStateType                  `json:"under_representation_dataset"`
	UnderRepresentationNSamples  float64                        `json:"under_representation_n_samples"`
	OverRepresentationDataset    DataStateType                  `json:"over_representation_dataset"`
	OverRepresentationNSamples   float64                        `json:"over_representation_n_samples"`
	// Construct a type with a set of properties K of type T
	Configuration map[string]interface{} `json:"configuration,omitempty"`
}

// NewUnderRepresentationInsight instantiates a new UnderRepresentationInsight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnderRepresentationInsight(type_ ScatterInsightType, filter ScatterFilter, underRepresentationDataset DataStateType, underRepresentationNSamples float64, overRepresentationDataset DataStateType, overRepresentationNSamples float64) *UnderRepresentationInsight {
	this := UnderRepresentationInsight{}
	this.Type = type_
	this.Filter = filter
	this.UnderRepresentationDataset = underRepresentationDataset
	this.UnderRepresentationNSamples = underRepresentationNSamples
	this.OverRepresentationDataset = overRepresentationDataset
	this.OverRepresentationNSamples = overRepresentationNSamples
	return &this
}

// NewUnderRepresentationInsightWithDefaults instantiates a new UnderRepresentationInsight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnderRepresentationInsightWithDefaults() *UnderRepresentationInsight {
	this := UnderRepresentationInsight{}
	return &this
}

// GetType returns the Type field value
func (o *UnderRepresentationInsight) GetType() ScatterInsightType {
	if o == nil {
		var ret ScatterInsightType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UnderRepresentationInsight) GetTypeOk() (*ScatterInsightType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UnderRepresentationInsight) SetType(v ScatterInsightType) {
	o.Type = v
}

// GetFilter returns the Filter field value
func (o *UnderRepresentationInsight) GetFilter() ScatterFilter {
	if o == nil {
		var ret ScatterFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *UnderRepresentationInsight) GetFilterOk() (*ScatterFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *UnderRepresentationInsight) SetFilter(v ScatterFilter) {
	o.Filter = v
}

// GetMutualInfoElements returns the MutualInfoElements field value if set, zero value otherwise.
func (o *UnderRepresentationInsight) GetMutualInfoElements() []MutualInformationElement {
	if o == nil || IsNil(o.MutualInfoElements) {
		var ret []MutualInformationElement
		return ret
	}
	return o.MutualInfoElements
}

// GetMutualInfoElementsOk returns a tuple with the MutualInfoElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnderRepresentationInsight) GetMutualInfoElementsOk() ([]MutualInformationElement, bool) {
	if o == nil || IsNil(o.MutualInfoElements) {
		return nil, false
	}
	return o.MutualInfoElements, true
}

// HasMutualInfoElements returns a boolean if a field has been set.
func (o *UnderRepresentationInsight) HasMutualInfoElements() bool {
	if o != nil && !IsNil(o.MutualInfoElements) {
		return true
	}

	return false
}

// SetMutualInfoElements gets a reference to the given []MutualInformationElement and assigns it to the MutualInfoElements field.
func (o *UnderRepresentationInsight) SetMutualInfoElements(v []MutualInformationElement) {
	o.MutualInfoElements = v
}

// GetBlobPath returns the BlobPath field value if set, zero value otherwise.
func (o *UnderRepresentationInsight) GetBlobPath() string {
	if o == nil || IsNil(o.BlobPath) {
		var ret string
		return ret
	}
	return *o.BlobPath
}

// GetBlobPathOk returns a tuple with the BlobPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnderRepresentationInsight) GetBlobPathOk() (*string, bool) {
	if o == nil || IsNil(o.BlobPath) {
		return nil, false
	}
	return o.BlobPath, true
}

// HasBlobPath returns a boolean if a field has been set.
func (o *UnderRepresentationInsight) HasBlobPath() bool {
	if o != nil && !IsNil(o.BlobPath) {
		return true
	}

	return false
}

// SetBlobPath gets a reference to the given string and assigns it to the BlobPath field.
func (o *UnderRepresentationInsight) SetBlobPath(v string) {
	o.BlobPath = &v
}

// GetInsightSubCategoryDsCuration returns the InsightSubCategoryDsCuration field value if set, zero value otherwise.
func (o *UnderRepresentationInsight) GetInsightSubCategoryDsCuration() InsightSubCategoryDSCuration {
	if o == nil || IsNil(o.InsightSubCategoryDsCuration) {
		var ret InsightSubCategoryDSCuration
		return ret
	}
	return *o.InsightSubCategoryDsCuration
}

// GetInsightSubCategoryDsCurationOk returns a tuple with the InsightSubCategoryDsCuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnderRepresentationInsight) GetInsightSubCategoryDsCurationOk() (*InsightSubCategoryDSCuration, bool) {
	if o == nil || IsNil(o.InsightSubCategoryDsCuration) {
		return nil, false
	}
	return o.InsightSubCategoryDsCuration, true
}

// HasInsightSubCategoryDsCuration returns a boolean if a field has been set.
func (o *UnderRepresentationInsight) HasInsightSubCategoryDsCuration() bool {
	if o != nil && !IsNil(o.InsightSubCategoryDsCuration) {
		return true
	}

	return false
}

// SetInsightSubCategoryDsCuration gets a reference to the given InsightSubCategoryDSCuration and assigns it to the InsightSubCategoryDsCuration field.
func (o *UnderRepresentationInsight) SetInsightSubCategoryDsCuration(v InsightSubCategoryDSCuration) {
	o.InsightSubCategoryDsCuration = &v
}

// GetInsightCategoryPerformance returns the InsightCategoryPerformance field value if set, zero value otherwise.
func (o *UnderRepresentationInsight) GetInsightCategoryPerformance() InsightSubCategoryPerformance {
	if o == nil || IsNil(o.InsightCategoryPerformance) {
		var ret InsightSubCategoryPerformance
		return ret
	}
	return *o.InsightCategoryPerformance
}

// GetInsightCategoryPerformanceOk returns a tuple with the InsightCategoryPerformance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnderRepresentationInsight) GetInsightCategoryPerformanceOk() (*InsightSubCategoryPerformance, bool) {
	if o == nil || IsNil(o.InsightCategoryPerformance) {
		return nil, false
	}
	return o.InsightCategoryPerformance, true
}

// HasInsightCategoryPerformance returns a boolean if a field has been set.
func (o *UnderRepresentationInsight) HasInsightCategoryPerformance() bool {
	if o != nil && !IsNil(o.InsightCategoryPerformance) {
		return true
	}

	return false
}

// SetInsightCategoryPerformance gets a reference to the given InsightSubCategoryPerformance and assigns it to the InsightCategoryPerformance field.
func (o *UnderRepresentationInsight) SetInsightCategoryPerformance(v InsightSubCategoryPerformance) {
	o.InsightCategoryPerformance = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *UnderRepresentationInsight) GetIndex() float64 {
	if o == nil || IsNil(o.Index) {
		var ret float64
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnderRepresentationInsight) GetIndexOk() (*float64, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *UnderRepresentationInsight) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given float64 and assigns it to the Index field.
func (o *UnderRepresentationInsight) SetIndex(v float64) {
	o.Index = &v
}

// GetUnderRepresentationDataset returns the UnderRepresentationDataset field value
func (o *UnderRepresentationInsight) GetUnderRepresentationDataset() DataStateType {
	if o == nil {
		var ret DataStateType
		return ret
	}

	return o.UnderRepresentationDataset
}

// GetUnderRepresentationDatasetOk returns a tuple with the UnderRepresentationDataset field value
// and a boolean to check if the value has been set.
func (o *UnderRepresentationInsight) GetUnderRepresentationDatasetOk() (*DataStateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderRepresentationDataset, true
}

// SetUnderRepresentationDataset sets field value
func (o *UnderRepresentationInsight) SetUnderRepresentationDataset(v DataStateType) {
	o.UnderRepresentationDataset = v
}

// GetUnderRepresentationNSamples returns the UnderRepresentationNSamples field value
func (o *UnderRepresentationInsight) GetUnderRepresentationNSamples() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.UnderRepresentationNSamples
}

// GetUnderRepresentationNSamplesOk returns a tuple with the UnderRepresentationNSamples field value
// and a boolean to check if the value has been set.
func (o *UnderRepresentationInsight) GetUnderRepresentationNSamplesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnderRepresentationNSamples, true
}

// SetUnderRepresentationNSamples sets field value
func (o *UnderRepresentationInsight) SetUnderRepresentationNSamples(v float64) {
	o.UnderRepresentationNSamples = v
}

// GetOverRepresentationDataset returns the OverRepresentationDataset field value
func (o *UnderRepresentationInsight) GetOverRepresentationDataset() DataStateType {
	if o == nil {
		var ret DataStateType
		return ret
	}

	return o.OverRepresentationDataset
}

// GetOverRepresentationDatasetOk returns a tuple with the OverRepresentationDataset field value
// and a boolean to check if the value has been set.
func (o *UnderRepresentationInsight) GetOverRepresentationDatasetOk() (*DataStateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverRepresentationDataset, true
}

// SetOverRepresentationDataset sets field value
func (o *UnderRepresentationInsight) SetOverRepresentationDataset(v DataStateType) {
	o.OverRepresentationDataset = v
}

// GetOverRepresentationNSamples returns the OverRepresentationNSamples field value
func (o *UnderRepresentationInsight) GetOverRepresentationNSamples() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.OverRepresentationNSamples
}

// GetOverRepresentationNSamplesOk returns a tuple with the OverRepresentationNSamples field value
// and a boolean to check if the value has been set.
func (o *UnderRepresentationInsight) GetOverRepresentationNSamplesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverRepresentationNSamples, true
}

// SetOverRepresentationNSamples sets field value
func (o *UnderRepresentationInsight) SetOverRepresentationNSamples(v float64) {
	o.OverRepresentationNSamples = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *UnderRepresentationInsight) GetConfiguration() map[string]interface{} {
	if o == nil || IsNil(o.Configuration) {
		var ret map[string]interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnderRepresentationInsight) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Configuration) {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *UnderRepresentationInsight) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]interface{} and assigns it to the Configuration field.
func (o *UnderRepresentationInsight) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

func (o UnderRepresentationInsight) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnderRepresentationInsight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["filter"] = o.Filter
	if !IsNil(o.MutualInfoElements) {
		toSerialize["mutual_info_elements"] = o.MutualInfoElements
	}
	if !IsNil(o.BlobPath) {
		toSerialize["blob_path"] = o.BlobPath
	}
	if !IsNil(o.InsightSubCategoryDsCuration) {
		toSerialize["insight_sub_category_ds_curation"] = o.InsightSubCategoryDsCuration
	}
	if !IsNil(o.InsightCategoryPerformance) {
		toSerialize["insight_category_performance"] = o.InsightCategoryPerformance
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	toSerialize["under_representation_dataset"] = o.UnderRepresentationDataset
	toSerialize["under_representation_n_samples"] = o.UnderRepresentationNSamples
	toSerialize["over_representation_dataset"] = o.OverRepresentationDataset
	toSerialize["over_representation_n_samples"] = o.OverRepresentationNSamples
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	return toSerialize, nil
}

type NullableUnderRepresentationInsight struct {
	value *UnderRepresentationInsight
	isSet bool
}

func (v NullableUnderRepresentationInsight) Get() *UnderRepresentationInsight {
	return v.value
}

func (v *NullableUnderRepresentationInsight) Set(val *UnderRepresentationInsight) {
	v.value = val
	v.isSet = true
}

func (v NullableUnderRepresentationInsight) IsSet() bool {
	return v.isSet
}

func (v *NullableUnderRepresentationInsight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnderRepresentationInsight(val *UnderRepresentationInsight) *NullableUnderRepresentationInsight {
	return &NullableUnderRepresentationInsight{value: val, isSet: true}
}

func (v NullableUnderRepresentationInsight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnderRepresentationInsight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
