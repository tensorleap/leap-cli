/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.365
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the PopulationExplorationParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PopulationExplorationParams{}

// PopulationExplorationParams struct for PopulationExplorationParams
type PopulationExplorationParams struct {
	SessionRunId string  `json:"sessionRunId"`
	ProjectId    string  `json:"projectId"`
	BatchSize    float64 `json:"batchSize"`
	NumOfSamples float64 `json:"numOfSamples"`
	FromEpoch    float64 `json:"fromEpoch"`
}

// NewPopulationExplorationParams instantiates a new PopulationExplorationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopulationExplorationParams(sessionRunId string, projectId string, batchSize float64, numOfSamples float64, fromEpoch float64) *PopulationExplorationParams {
	this := PopulationExplorationParams{}
	this.SessionRunId = sessionRunId
	this.ProjectId = projectId
	this.BatchSize = batchSize
	this.NumOfSamples = numOfSamples
	this.FromEpoch = fromEpoch
	return &this
}

// NewPopulationExplorationParamsWithDefaults instantiates a new PopulationExplorationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopulationExplorationParamsWithDefaults() *PopulationExplorationParams {
	this := PopulationExplorationParams{}
	return &this
}

// GetSessionRunId returns the SessionRunId field value
func (o *PopulationExplorationParams) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *PopulationExplorationParams) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *PopulationExplorationParams) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetProjectId returns the ProjectId field value
func (o *PopulationExplorationParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *PopulationExplorationParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *PopulationExplorationParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetBatchSize returns the BatchSize field value
func (o *PopulationExplorationParams) GetBatchSize() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value
// and a boolean to check if the value has been set.
func (o *PopulationExplorationParams) GetBatchSizeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchSize, true
}

// SetBatchSize sets field value
func (o *PopulationExplorationParams) SetBatchSize(v float64) {
	o.BatchSize = v
}

// GetNumOfSamples returns the NumOfSamples field value
func (o *PopulationExplorationParams) GetNumOfSamples() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.NumOfSamples
}

// GetNumOfSamplesOk returns a tuple with the NumOfSamples field value
// and a boolean to check if the value has been set.
func (o *PopulationExplorationParams) GetNumOfSamplesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumOfSamples, true
}

// SetNumOfSamples sets field value
func (o *PopulationExplorationParams) SetNumOfSamples(v float64) {
	o.NumOfSamples = v
}

// GetFromEpoch returns the FromEpoch field value
func (o *PopulationExplorationParams) GetFromEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.FromEpoch
}

// GetFromEpochOk returns a tuple with the FromEpoch field value
// and a boolean to check if the value has been set.
func (o *PopulationExplorationParams) GetFromEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromEpoch, true
}

// SetFromEpoch sets field value
func (o *PopulationExplorationParams) SetFromEpoch(v float64) {
	o.FromEpoch = v
}

func (o PopulationExplorationParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PopulationExplorationParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["projectId"] = o.ProjectId
	toSerialize["batchSize"] = o.BatchSize
	toSerialize["numOfSamples"] = o.NumOfSamples
	toSerialize["fromEpoch"] = o.FromEpoch
	return toSerialize, nil
}

type NullablePopulationExplorationParams struct {
	value *PopulationExplorationParams
	isSet bool
}

func (v NullablePopulationExplorationParams) Get() *PopulationExplorationParams {
	return v.value
}

func (v *NullablePopulationExplorationParams) Set(val *PopulationExplorationParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePopulationExplorationParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePopulationExplorationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopulationExplorationParams(val *PopulationExplorationParams) *NullablePopulationExplorationParams {
	return &NullablePopulationExplorationParams{value: val, isSet: true}
}

func (v NullablePopulationExplorationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopulationExplorationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
