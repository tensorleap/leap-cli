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

// checks if the RocConfusionMatrixParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RocConfusionMatrixParams{}

// RocConfusionMatrixParams struct for RocConfusionMatrixParams
type RocConfusionMatrixParams struct {
	SessionRunsToEpochs []SessionRunToEpoch `json:"sessionRunsToEpochs"`
	ProjectId           string              `json:"projectId"`
	CustomMetricName    string              `json:"customMetricName"`
	VerticalSplit       *SplitAgg           `json:"verticalSplit,omitempty"`
	HorizontalSplit     *SplitAgg           `json:"horizontalSplit,omitempty"`
	Filters             []ESFilter          `json:"filters,omitempty"`
	AbsAxis             *bool               `json:"absAxis,omitempty"`
}

// NewRocConfusionMatrixParams instantiates a new RocConfusionMatrixParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRocConfusionMatrixParams(sessionRunsToEpochs []SessionRunToEpoch, projectId string, customMetricName string) *RocConfusionMatrixParams {
	this := RocConfusionMatrixParams{}
	this.SessionRunsToEpochs = sessionRunsToEpochs
	this.ProjectId = projectId
	this.CustomMetricName = customMetricName
	return &this
}

// NewRocConfusionMatrixParamsWithDefaults instantiates a new RocConfusionMatrixParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRocConfusionMatrixParamsWithDefaults() *RocConfusionMatrixParams {
	this := RocConfusionMatrixParams{}
	return &this
}

// GetSessionRunsToEpochs returns the SessionRunsToEpochs field value
func (o *RocConfusionMatrixParams) GetSessionRunsToEpochs() []SessionRunToEpoch {
	if o == nil {
		var ret []SessionRunToEpoch
		return ret
	}

	return o.SessionRunsToEpochs
}

// GetSessionRunsToEpochsOk returns a tuple with the SessionRunsToEpochs field value
// and a boolean to check if the value has been set.
func (o *RocConfusionMatrixParams) GetSessionRunsToEpochsOk() ([]SessionRunToEpoch, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionRunsToEpochs, true
}

// SetSessionRunsToEpochs sets field value
func (o *RocConfusionMatrixParams) SetSessionRunsToEpochs(v []SessionRunToEpoch) {
	o.SessionRunsToEpochs = v
}

// GetProjectId returns the ProjectId field value
func (o *RocConfusionMatrixParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *RocConfusionMatrixParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *RocConfusionMatrixParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetCustomMetricName returns the CustomMetricName field value
func (o *RocConfusionMatrixParams) GetCustomMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomMetricName
}

// GetCustomMetricNameOk returns a tuple with the CustomMetricName field value
// and a boolean to check if the value has been set.
func (o *RocConfusionMatrixParams) GetCustomMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomMetricName, true
}

// SetCustomMetricName sets field value
func (o *RocConfusionMatrixParams) SetCustomMetricName(v string) {
	o.CustomMetricName = v
}

// GetVerticalSplit returns the VerticalSplit field value if set, zero value otherwise.
func (o *RocConfusionMatrixParams) GetVerticalSplit() SplitAgg {
	if o == nil || IsNil(o.VerticalSplit) {
		var ret SplitAgg
		return ret
	}
	return *o.VerticalSplit
}

// GetVerticalSplitOk returns a tuple with the VerticalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RocConfusionMatrixParams) GetVerticalSplitOk() (*SplitAgg, bool) {
	if o == nil || IsNil(o.VerticalSplit) {
		return nil, false
	}
	return o.VerticalSplit, true
}

// HasVerticalSplit returns a boolean if a field has been set.
func (o *RocConfusionMatrixParams) HasVerticalSplit() bool {
	if o != nil && !IsNil(o.VerticalSplit) {
		return true
	}

	return false
}

// SetVerticalSplit gets a reference to the given SplitAgg and assigns it to the VerticalSplit field.
func (o *RocConfusionMatrixParams) SetVerticalSplit(v SplitAgg) {
	o.VerticalSplit = &v
}

// GetHorizontalSplit returns the HorizontalSplit field value if set, zero value otherwise.
func (o *RocConfusionMatrixParams) GetHorizontalSplit() SplitAgg {
	if o == nil || IsNil(o.HorizontalSplit) {
		var ret SplitAgg
		return ret
	}
	return *o.HorizontalSplit
}

// GetHorizontalSplitOk returns a tuple with the HorizontalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RocConfusionMatrixParams) GetHorizontalSplitOk() (*SplitAgg, bool) {
	if o == nil || IsNil(o.HorizontalSplit) {
		return nil, false
	}
	return o.HorizontalSplit, true
}

// HasHorizontalSplit returns a boolean if a field has been set.
func (o *RocConfusionMatrixParams) HasHorizontalSplit() bool {
	if o != nil && !IsNil(o.HorizontalSplit) {
		return true
	}

	return false
}

// SetHorizontalSplit gets a reference to the given SplitAgg and assigns it to the HorizontalSplit field.
func (o *RocConfusionMatrixParams) SetHorizontalSplit(v SplitAgg) {
	o.HorizontalSplit = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *RocConfusionMatrixParams) GetFilters() []ESFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []ESFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RocConfusionMatrixParams) GetFiltersOk() ([]ESFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *RocConfusionMatrixParams) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []ESFilter and assigns it to the Filters field.
func (o *RocConfusionMatrixParams) SetFilters(v []ESFilter) {
	o.Filters = v
}

// GetAbsAxis returns the AbsAxis field value if set, zero value otherwise.
func (o *RocConfusionMatrixParams) GetAbsAxis() bool {
	if o == nil || IsNil(o.AbsAxis) {
		var ret bool
		return ret
	}
	return *o.AbsAxis
}

// GetAbsAxisOk returns a tuple with the AbsAxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RocConfusionMatrixParams) GetAbsAxisOk() (*bool, bool) {
	if o == nil || IsNil(o.AbsAxis) {
		return nil, false
	}
	return o.AbsAxis, true
}

// HasAbsAxis returns a boolean if a field has been set.
func (o *RocConfusionMatrixParams) HasAbsAxis() bool {
	if o != nil && !IsNil(o.AbsAxis) {
		return true
	}

	return false
}

// SetAbsAxis gets a reference to the given bool and assigns it to the AbsAxis field.
func (o *RocConfusionMatrixParams) SetAbsAxis(v bool) {
	o.AbsAxis = &v
}

func (o RocConfusionMatrixParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RocConfusionMatrixParams) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AbsAxis) {
		toSerialize["absAxis"] = o.AbsAxis
	}
	return toSerialize, nil
}

type NullableRocConfusionMatrixParams struct {
	value *RocConfusionMatrixParams
	isSet bool
}

func (v NullableRocConfusionMatrixParams) Get() *RocConfusionMatrixParams {
	return v.value
}

func (v *NullableRocConfusionMatrixParams) Set(val *RocConfusionMatrixParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRocConfusionMatrixParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRocConfusionMatrixParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRocConfusionMatrixParams(val *RocConfusionMatrixParams) *NullableRocConfusionMatrixParams {
	return &NullableRocConfusionMatrixParams{value: val, isSet: true}
}

func (v NullableRocConfusionMatrixParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRocConfusionMatrixParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
