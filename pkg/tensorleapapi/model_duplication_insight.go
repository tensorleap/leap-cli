/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the DuplicationInsight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DuplicationInsight{}

// DuplicationInsight struct for DuplicationInsight
type DuplicationInsight struct {
	Type                         ScatterInsightType             `json:"type"`
	Filter                       ScatterFilter                  `json:"filter"`
	MutualInfoElements           []MutualInformationElement     `json:"mutual_info_elements,omitempty"`
	BlobPath                     *string                        `json:"blob_path,omitempty"`
	InsightSubCategoryDsCuration *InsightSubCategoryDSCuration  `json:"insight_sub_category_ds_curation,omitempty"`
	InsightCategoryPerformance   *InsightSubCategoryPerformance `json:"insight_category_performance,omitempty"`
	Index                        *float64                       `json:"index,omitempty"`
	Subset                       DataStateType                  `json:"subset"`
	NDuplicateSamples            float64                        `json:"n_duplicate_samples"`
}

// NewDuplicationInsight instantiates a new DuplicationInsight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuplicationInsight(type_ ScatterInsightType, filter ScatterFilter, subset DataStateType, nDuplicateSamples float64) *DuplicationInsight {
	this := DuplicationInsight{}
	this.Type = type_
	this.Filter = filter
	this.Subset = subset
	this.NDuplicateSamples = nDuplicateSamples
	return &this
}

// NewDuplicationInsightWithDefaults instantiates a new DuplicationInsight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDuplicationInsightWithDefaults() *DuplicationInsight {
	this := DuplicationInsight{}
	return &this
}

// GetType returns the Type field value
func (o *DuplicationInsight) GetType() ScatterInsightType {
	if o == nil {
		var ret ScatterInsightType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DuplicationInsight) GetTypeOk() (*ScatterInsightType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DuplicationInsight) SetType(v ScatterInsightType) {
	o.Type = v
}

// GetFilter returns the Filter field value
func (o *DuplicationInsight) GetFilter() ScatterFilter {
	if o == nil {
		var ret ScatterFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *DuplicationInsight) GetFilterOk() (*ScatterFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *DuplicationInsight) SetFilter(v ScatterFilter) {
	o.Filter = v
}

// GetMutualInfoElements returns the MutualInfoElements field value if set, zero value otherwise.
func (o *DuplicationInsight) GetMutualInfoElements() []MutualInformationElement {
	if o == nil || IsNil(o.MutualInfoElements) {
		var ret []MutualInformationElement
		return ret
	}
	return o.MutualInfoElements
}

// GetMutualInfoElementsOk returns a tuple with the MutualInfoElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DuplicationInsight) GetMutualInfoElementsOk() ([]MutualInformationElement, bool) {
	if o == nil || IsNil(o.MutualInfoElements) {
		return nil, false
	}
	return o.MutualInfoElements, true
}

// HasMutualInfoElements returns a boolean if a field has been set.
func (o *DuplicationInsight) HasMutualInfoElements() bool {
	if o != nil && !IsNil(o.MutualInfoElements) {
		return true
	}

	return false
}

// SetMutualInfoElements gets a reference to the given []MutualInformationElement and assigns it to the MutualInfoElements field.
func (o *DuplicationInsight) SetMutualInfoElements(v []MutualInformationElement) {
	o.MutualInfoElements = v
}

// GetBlobPath returns the BlobPath field value if set, zero value otherwise.
func (o *DuplicationInsight) GetBlobPath() string {
	if o == nil || IsNil(o.BlobPath) {
		var ret string
		return ret
	}
	return *o.BlobPath
}

// GetBlobPathOk returns a tuple with the BlobPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DuplicationInsight) GetBlobPathOk() (*string, bool) {
	if o == nil || IsNil(o.BlobPath) {
		return nil, false
	}
	return o.BlobPath, true
}

// HasBlobPath returns a boolean if a field has been set.
func (o *DuplicationInsight) HasBlobPath() bool {
	if o != nil && !IsNil(o.BlobPath) {
		return true
	}

	return false
}

// SetBlobPath gets a reference to the given string and assigns it to the BlobPath field.
func (o *DuplicationInsight) SetBlobPath(v string) {
	o.BlobPath = &v
}

// GetInsightSubCategoryDsCuration returns the InsightSubCategoryDsCuration field value if set, zero value otherwise.
func (o *DuplicationInsight) GetInsightSubCategoryDsCuration() InsightSubCategoryDSCuration {
	if o == nil || IsNil(o.InsightSubCategoryDsCuration) {
		var ret InsightSubCategoryDSCuration
		return ret
	}
	return *o.InsightSubCategoryDsCuration
}

// GetInsightSubCategoryDsCurationOk returns a tuple with the InsightSubCategoryDsCuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DuplicationInsight) GetInsightSubCategoryDsCurationOk() (*InsightSubCategoryDSCuration, bool) {
	if o == nil || IsNil(o.InsightSubCategoryDsCuration) {
		return nil, false
	}
	return o.InsightSubCategoryDsCuration, true
}

// HasInsightSubCategoryDsCuration returns a boolean if a field has been set.
func (o *DuplicationInsight) HasInsightSubCategoryDsCuration() bool {
	if o != nil && !IsNil(o.InsightSubCategoryDsCuration) {
		return true
	}

	return false
}

// SetInsightSubCategoryDsCuration gets a reference to the given InsightSubCategoryDSCuration and assigns it to the InsightSubCategoryDsCuration field.
func (o *DuplicationInsight) SetInsightSubCategoryDsCuration(v InsightSubCategoryDSCuration) {
	o.InsightSubCategoryDsCuration = &v
}

// GetInsightCategoryPerformance returns the InsightCategoryPerformance field value if set, zero value otherwise.
func (o *DuplicationInsight) GetInsightCategoryPerformance() InsightSubCategoryPerformance {
	if o == nil || IsNil(o.InsightCategoryPerformance) {
		var ret InsightSubCategoryPerformance
		return ret
	}
	return *o.InsightCategoryPerformance
}

// GetInsightCategoryPerformanceOk returns a tuple with the InsightCategoryPerformance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DuplicationInsight) GetInsightCategoryPerformanceOk() (*InsightSubCategoryPerformance, bool) {
	if o == nil || IsNil(o.InsightCategoryPerformance) {
		return nil, false
	}
	return o.InsightCategoryPerformance, true
}

// HasInsightCategoryPerformance returns a boolean if a field has been set.
func (o *DuplicationInsight) HasInsightCategoryPerformance() bool {
	if o != nil && !IsNil(o.InsightCategoryPerformance) {
		return true
	}

	return false
}

// SetInsightCategoryPerformance gets a reference to the given InsightSubCategoryPerformance and assigns it to the InsightCategoryPerformance field.
func (o *DuplicationInsight) SetInsightCategoryPerformance(v InsightSubCategoryPerformance) {
	o.InsightCategoryPerformance = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *DuplicationInsight) GetIndex() float64 {
	if o == nil || IsNil(o.Index) {
		var ret float64
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DuplicationInsight) GetIndexOk() (*float64, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *DuplicationInsight) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given float64 and assigns it to the Index field.
func (o *DuplicationInsight) SetIndex(v float64) {
	o.Index = &v
}

// GetSubset returns the Subset field value
func (o *DuplicationInsight) GetSubset() DataStateType {
	if o == nil {
		var ret DataStateType
		return ret
	}

	return o.Subset
}

// GetSubsetOk returns a tuple with the Subset field value
// and a boolean to check if the value has been set.
func (o *DuplicationInsight) GetSubsetOk() (*DataStateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subset, true
}

// SetSubset sets field value
func (o *DuplicationInsight) SetSubset(v DataStateType) {
	o.Subset = v
}

// GetNDuplicateSamples returns the NDuplicateSamples field value
func (o *DuplicationInsight) GetNDuplicateSamples() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.NDuplicateSamples
}

// GetNDuplicateSamplesOk returns a tuple with the NDuplicateSamples field value
// and a boolean to check if the value has been set.
func (o *DuplicationInsight) GetNDuplicateSamplesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NDuplicateSamples, true
}

// SetNDuplicateSamples sets field value
func (o *DuplicationInsight) SetNDuplicateSamples(v float64) {
	o.NDuplicateSamples = v
}

func (o DuplicationInsight) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DuplicationInsight) ToMap() (map[string]interface{}, error) {
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
	toSerialize["subset"] = o.Subset
	toSerialize["n_duplicate_samples"] = o.NDuplicateSamples
	return toSerialize, nil
}

type NullableDuplicationInsight struct {
	value *DuplicationInsight
	isSet bool
}

func (v NullableDuplicationInsight) Get() *DuplicationInsight {
	return v.value
}

func (v *NullableDuplicationInsight) Set(val *DuplicationInsight) {
	v.value = val
	v.isSet = true
}

func (v NullableDuplicationInsight) IsSet() bool {
	return v.isSet
}

func (v *NullableDuplicationInsight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuplicationInsight(val *DuplicationInsight) *NullableDuplicationInsight {
	return &NullableDuplicationInsight{value: val, isSet: true}
}

func (v NullableDuplicationInsight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuplicationInsight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
