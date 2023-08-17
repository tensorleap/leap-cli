/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
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
	Type                        ScatterInsightType         `json:"type"`
	Filter                      ScatterFilter              `json:"filter"`
	MutualInfoElements          []MutualInformationElement `json:"mutual_info_elements,omitempty"`
	UnderRepresentationDataset  DataStateType              `json:"under_representation_dataset"`
	UnderRepresentationNSamples float64                    `json:"under_representation_n_samples"`
	OverRepresentationDataset   DataStateType              `json:"over_representation_dataset"`
	OverRepresentationNSamples  float64                    `json:"over_representation_n_samples"`
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
	toSerialize["under_representation_dataset"] = o.UnderRepresentationDataset
	toSerialize["under_representation_n_samples"] = o.UnderRepresentationNSamples
	toSerialize["over_representation_dataset"] = o.OverRepresentationDataset
	toSerialize["over_representation_n_samples"] = o.OverRepresentationNSamples
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
