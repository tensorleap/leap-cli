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

// checks if the PopulationExplorationDisplayParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PopulationExplorationDisplayParams{}

// PopulationExplorationDisplayParams struct for PopulationExplorationDisplayParams
type PopulationExplorationDisplayParams struct {
	PopulationExplorationNSamples    float64            `json:"population_exploration_n_samples"`
	BalanceBy                        []string           `json:"balance_by"`
	ShouldFillRemainingWithUnbalance bool               `json:"should_fill_remaining_with_unbalance"`
	ReductionAlgorithm               ReductionAlgorithm `json:"reduction_algorithm"`
	OptionalAnalysis                 []OptionalAnalysis `json:"optional_analysis,omitempty"`
}

// NewPopulationExplorationDisplayParams instantiates a new PopulationExplorationDisplayParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopulationExplorationDisplayParams(populationExplorationNSamples float64, balanceBy []string, shouldFillRemainingWithUnbalance bool, reductionAlgorithm ReductionAlgorithm) *PopulationExplorationDisplayParams {
	this := PopulationExplorationDisplayParams{}
	this.PopulationExplorationNSamples = populationExplorationNSamples
	this.BalanceBy = balanceBy
	this.ShouldFillRemainingWithUnbalance = shouldFillRemainingWithUnbalance
	this.ReductionAlgorithm = reductionAlgorithm
	return &this
}

// NewPopulationExplorationDisplayParamsWithDefaults instantiates a new PopulationExplorationDisplayParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopulationExplorationDisplayParamsWithDefaults() *PopulationExplorationDisplayParams {
	this := PopulationExplorationDisplayParams{}
	return &this
}

// GetPopulationExplorationNSamples returns the PopulationExplorationNSamples field value
func (o *PopulationExplorationDisplayParams) GetPopulationExplorationNSamples() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.PopulationExplorationNSamples
}

// GetPopulationExplorationNSamplesOk returns a tuple with the PopulationExplorationNSamples field value
// and a boolean to check if the value has been set.
func (o *PopulationExplorationDisplayParams) GetPopulationExplorationNSamplesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PopulationExplorationNSamples, true
}

// SetPopulationExplorationNSamples sets field value
func (o *PopulationExplorationDisplayParams) SetPopulationExplorationNSamples(v float64) {
	o.PopulationExplorationNSamples = v
}

// GetBalanceBy returns the BalanceBy field value
func (o *PopulationExplorationDisplayParams) GetBalanceBy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BalanceBy
}

// GetBalanceByOk returns a tuple with the BalanceBy field value
// and a boolean to check if the value has been set.
func (o *PopulationExplorationDisplayParams) GetBalanceByOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BalanceBy, true
}

// SetBalanceBy sets field value
func (o *PopulationExplorationDisplayParams) SetBalanceBy(v []string) {
	o.BalanceBy = v
}

// GetShouldFillRemainingWithUnbalance returns the ShouldFillRemainingWithUnbalance field value
func (o *PopulationExplorationDisplayParams) GetShouldFillRemainingWithUnbalance() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShouldFillRemainingWithUnbalance
}

// GetShouldFillRemainingWithUnbalanceOk returns a tuple with the ShouldFillRemainingWithUnbalance field value
// and a boolean to check if the value has been set.
func (o *PopulationExplorationDisplayParams) GetShouldFillRemainingWithUnbalanceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShouldFillRemainingWithUnbalance, true
}

// SetShouldFillRemainingWithUnbalance sets field value
func (o *PopulationExplorationDisplayParams) SetShouldFillRemainingWithUnbalance(v bool) {
	o.ShouldFillRemainingWithUnbalance = v
}

// GetReductionAlgorithm returns the ReductionAlgorithm field value
func (o *PopulationExplorationDisplayParams) GetReductionAlgorithm() ReductionAlgorithm {
	if o == nil {
		var ret ReductionAlgorithm
		return ret
	}

	return o.ReductionAlgorithm
}

// GetReductionAlgorithmOk returns a tuple with the ReductionAlgorithm field value
// and a boolean to check if the value has been set.
func (o *PopulationExplorationDisplayParams) GetReductionAlgorithmOk() (*ReductionAlgorithm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReductionAlgorithm, true
}

// SetReductionAlgorithm sets field value
func (o *PopulationExplorationDisplayParams) SetReductionAlgorithm(v ReductionAlgorithm) {
	o.ReductionAlgorithm = v
}

// GetOptionalAnalysis returns the OptionalAnalysis field value if set, zero value otherwise.
func (o *PopulationExplorationDisplayParams) GetOptionalAnalysis() []OptionalAnalysis {
	if o == nil || IsNil(o.OptionalAnalysis) {
		var ret []OptionalAnalysis
		return ret
	}
	return o.OptionalAnalysis
}

// GetOptionalAnalysisOk returns a tuple with the OptionalAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PopulationExplorationDisplayParams) GetOptionalAnalysisOk() ([]OptionalAnalysis, bool) {
	if o == nil || IsNil(o.OptionalAnalysis) {
		return nil, false
	}
	return o.OptionalAnalysis, true
}

// HasOptionalAnalysis returns a boolean if a field has been set.
func (o *PopulationExplorationDisplayParams) HasOptionalAnalysis() bool {
	if o != nil && !IsNil(o.OptionalAnalysis) {
		return true
	}

	return false
}

// SetOptionalAnalysis gets a reference to the given []OptionalAnalysis and assigns it to the OptionalAnalysis field.
func (o *PopulationExplorationDisplayParams) SetOptionalAnalysis(v []OptionalAnalysis) {
	o.OptionalAnalysis = v
}

func (o PopulationExplorationDisplayParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PopulationExplorationDisplayParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["population_exploration_n_samples"] = o.PopulationExplorationNSamples
	toSerialize["balance_by"] = o.BalanceBy
	toSerialize["should_fill_remaining_with_unbalance"] = o.ShouldFillRemainingWithUnbalance
	toSerialize["reduction_algorithm"] = o.ReductionAlgorithm
	if !IsNil(o.OptionalAnalysis) {
		toSerialize["optional_analysis"] = o.OptionalAnalysis
	}
	return toSerialize, nil
}

type NullablePopulationExplorationDisplayParams struct {
	value *PopulationExplorationDisplayParams
	isSet bool
}

func (v NullablePopulationExplorationDisplayParams) Get() *PopulationExplorationDisplayParams {
	return v.value
}

func (v *NullablePopulationExplorationDisplayParams) Set(val *PopulationExplorationDisplayParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePopulationExplorationDisplayParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePopulationExplorationDisplayParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopulationExplorationDisplayParams(val *PopulationExplorationDisplayParams) *NullablePopulationExplorationDisplayParams {
	return &NullablePopulationExplorationDisplayParams{value: val, isSet: true}
}

func (v NullablePopulationExplorationDisplayParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopulationExplorationDisplayParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
