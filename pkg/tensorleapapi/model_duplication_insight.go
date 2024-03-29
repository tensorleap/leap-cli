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

// checks if the DuplicationInsight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DuplicationInsight{}

// DuplicationInsight struct for DuplicationInsight
type DuplicationInsight struct {
	Type                     ScatterInsightType         `json:"type"`
	Filter                   ScatterFilter              `json:"filter"`
	MutualInfoElements       []MutualInformationElement `json:"mutual_info_elements,omitempty"`
	BlobPath                 *string                    `json:"blob_path,omitempty"`
	NDuplicateSamples        float64                    `json:"n_duplicate_samples"`
	NCrossSubsetDuplications float64                    `json:"n_cross_subset_duplications"`
}

// NewDuplicationInsight instantiates a new DuplicationInsight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuplicationInsight(type_ ScatterInsightType, filter ScatterFilter, nDuplicateSamples float64, nCrossSubsetDuplications float64) *DuplicationInsight {
	this := DuplicationInsight{}
	this.Type = type_
	this.Filter = filter
	this.NDuplicateSamples = nDuplicateSamples
	this.NCrossSubsetDuplications = nCrossSubsetDuplications
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

// GetNCrossSubsetDuplications returns the NCrossSubsetDuplications field value
func (o *DuplicationInsight) GetNCrossSubsetDuplications() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.NCrossSubsetDuplications
}

// GetNCrossSubsetDuplicationsOk returns a tuple with the NCrossSubsetDuplications field value
// and a boolean to check if the value has been set.
func (o *DuplicationInsight) GetNCrossSubsetDuplicationsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NCrossSubsetDuplications, true
}

// SetNCrossSubsetDuplications sets field value
func (o *DuplicationInsight) SetNCrossSubsetDuplications(v float64) {
	o.NCrossSubsetDuplications = v
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
	toSerialize["n_duplicate_samples"] = o.NDuplicateSamples
	toSerialize["n_cross_subset_duplications"] = o.NCrossSubsetDuplications
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
