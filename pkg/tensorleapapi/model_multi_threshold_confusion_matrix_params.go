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

// checks if the MultiThresholdConfusionMatrixParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiThresholdConfusionMatrixParams{}

// MultiThresholdConfusionMatrixParams struct for MultiThresholdConfusionMatrixParams
type MultiThresholdConfusionMatrixParams struct {
	SessionRunsToEpochs []SessionRunToEpoch `json:"sessionRunsToEpochs"`
	ProjectId           string              `json:"projectId"`
	CustomMetricName    string              `json:"customMetricName"`
	VerticalSplit       *SplitAgg           `json:"verticalSplit,omitempty"`
	HorizontalSplit     *SplitAgg           `json:"horizontalSplit,omitempty"`
	Filters             []ESFilter          `json:"filters,omitempty"`
}

// NewMultiThresholdConfusionMatrixParams instantiates a new MultiThresholdConfusionMatrixParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiThresholdConfusionMatrixParams(sessionRunsToEpochs []SessionRunToEpoch, projectId string, customMetricName string) *MultiThresholdConfusionMatrixParams {
	this := MultiThresholdConfusionMatrixParams{}
	this.SessionRunsToEpochs = sessionRunsToEpochs
	this.ProjectId = projectId
	this.CustomMetricName = customMetricName
	return &this
}

// NewMultiThresholdConfusionMatrixParamsWithDefaults instantiates a new MultiThresholdConfusionMatrixParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiThresholdConfusionMatrixParamsWithDefaults() *MultiThresholdConfusionMatrixParams {
	this := MultiThresholdConfusionMatrixParams{}
	return &this
}

// GetSessionRunsToEpochs returns the SessionRunsToEpochs field value
func (o *MultiThresholdConfusionMatrixParams) GetSessionRunsToEpochs() []SessionRunToEpoch {
	if o == nil {
		var ret []SessionRunToEpoch
		return ret
	}

	return o.SessionRunsToEpochs
}

// GetSessionRunsToEpochsOk returns a tuple with the SessionRunsToEpochs field value
// and a boolean to check if the value has been set.
func (o *MultiThresholdConfusionMatrixParams) GetSessionRunsToEpochsOk() ([]SessionRunToEpoch, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionRunsToEpochs, true
}

// SetSessionRunsToEpochs sets field value
func (o *MultiThresholdConfusionMatrixParams) SetSessionRunsToEpochs(v []SessionRunToEpoch) {
	o.SessionRunsToEpochs = v
}

// GetProjectId returns the ProjectId field value
func (o *MultiThresholdConfusionMatrixParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *MultiThresholdConfusionMatrixParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *MultiThresholdConfusionMatrixParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetCustomMetricName returns the CustomMetricName field value
func (o *MultiThresholdConfusionMatrixParams) GetCustomMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomMetricName
}

// GetCustomMetricNameOk returns a tuple with the CustomMetricName field value
// and a boolean to check if the value has been set.
func (o *MultiThresholdConfusionMatrixParams) GetCustomMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomMetricName, true
}

// SetCustomMetricName sets field value
func (o *MultiThresholdConfusionMatrixParams) SetCustomMetricName(v string) {
	o.CustomMetricName = v
}

// GetVerticalSplit returns the VerticalSplit field value if set, zero value otherwise.
func (o *MultiThresholdConfusionMatrixParams) GetVerticalSplit() SplitAgg {
	if o == nil || IsNil(o.VerticalSplit) {
		var ret SplitAgg
		return ret
	}
	return *o.VerticalSplit
}

// GetVerticalSplitOk returns a tuple with the VerticalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiThresholdConfusionMatrixParams) GetVerticalSplitOk() (*SplitAgg, bool) {
	if o == nil || IsNil(o.VerticalSplit) {
		return nil, false
	}
	return o.VerticalSplit, true
}

// HasVerticalSplit returns a boolean if a field has been set.
func (o *MultiThresholdConfusionMatrixParams) HasVerticalSplit() bool {
	if o != nil && !IsNil(o.VerticalSplit) {
		return true
	}

	return false
}

// SetVerticalSplit gets a reference to the given SplitAgg and assigns it to the VerticalSplit field.
func (o *MultiThresholdConfusionMatrixParams) SetVerticalSplit(v SplitAgg) {
	o.VerticalSplit = &v
}

// GetHorizontalSplit returns the HorizontalSplit field value if set, zero value otherwise.
func (o *MultiThresholdConfusionMatrixParams) GetHorizontalSplit() SplitAgg {
	if o == nil || IsNil(o.HorizontalSplit) {
		var ret SplitAgg
		return ret
	}
	return *o.HorizontalSplit
}

// GetHorizontalSplitOk returns a tuple with the HorizontalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiThresholdConfusionMatrixParams) GetHorizontalSplitOk() (*SplitAgg, bool) {
	if o == nil || IsNil(o.HorizontalSplit) {
		return nil, false
	}
	return o.HorizontalSplit, true
}

// HasHorizontalSplit returns a boolean if a field has been set.
func (o *MultiThresholdConfusionMatrixParams) HasHorizontalSplit() bool {
	if o != nil && !IsNil(o.HorizontalSplit) {
		return true
	}

	return false
}

// SetHorizontalSplit gets a reference to the given SplitAgg and assigns it to the HorizontalSplit field.
func (o *MultiThresholdConfusionMatrixParams) SetHorizontalSplit(v SplitAgg) {
	o.HorizontalSplit = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *MultiThresholdConfusionMatrixParams) GetFilters() []ESFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []ESFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiThresholdConfusionMatrixParams) GetFiltersOk() ([]ESFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *MultiThresholdConfusionMatrixParams) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []ESFilter and assigns it to the Filters field.
func (o *MultiThresholdConfusionMatrixParams) SetFilters(v []ESFilter) {
	o.Filters = v
}

func (o MultiThresholdConfusionMatrixParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiThresholdConfusionMatrixParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionRunsToEpochs"] = o.SessionRunsToEpochs
	toSerialize["projectId"] = o.ProjectId
	toSerialize["customMetricName"] = o.CustomMetricName
	if !IsNil(o.VerticalSplit) {
		toSerialize["verticalSplit"] = o.VerticalSplit
	}
	if !IsNil(o.HorizontalSplit) {
		toSerialize["horizontalSplit"] = o.HorizontalSplit
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return toSerialize, nil
}

type NullableMultiThresholdConfusionMatrixParams struct {
	value *MultiThresholdConfusionMatrixParams
	isSet bool
}

func (v NullableMultiThresholdConfusionMatrixParams) Get() *MultiThresholdConfusionMatrixParams {
	return v.value
}

func (v *NullableMultiThresholdConfusionMatrixParams) Set(val *MultiThresholdConfusionMatrixParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiThresholdConfusionMatrixParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiThresholdConfusionMatrixParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiThresholdConfusionMatrixParams(val *MultiThresholdConfusionMatrixParams) *NullableMultiThresholdConfusionMatrixParams {
	return &NullableMultiThresholdConfusionMatrixParams{value: val, isSet: true}
}

func (v NullableMultiThresholdConfusionMatrixParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiThresholdConfusionMatrixParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
