/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.526
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetConfusionMatrixResultCombinationsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetConfusionMatrixResultCombinationsParams{}

// GetConfusionMatrixResultCombinationsParams struct for GetConfusionMatrixResultCombinationsParams
type GetConfusionMatrixResultCombinationsParams struct {
	ProjectId        string     `json:"projectId"`
	SessionRunIds    []string   `json:"sessionRunIds"`
	X                SplitAgg   `json:"x"`
	VerticalSplit    *SplitAgg  `json:"verticalSplit,omitempty"`
	HorizontalSplit  *SplitAgg  `json:"horizontalSplit,omitempty"`
	InnerSplit       *SplitAgg  `json:"innerSplit,omitempty"`
	Threshold        *float64   `json:"threshold,omitempty"`
	CustomMetricName string     `json:"customMetricName"`
	Filters          []ESFilter `json:"filters,omitempty"`
	LastEpochOnly    *bool      `json:"lastEpochOnly,omitempty"`
	FilterLabels     []string   `json:"filterLabels,omitempty"`
}

// NewGetConfusionMatrixResultCombinationsParams instantiates a new GetConfusionMatrixResultCombinationsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConfusionMatrixResultCombinationsParams(projectId string, sessionRunIds []string, x SplitAgg, customMetricName string) *GetConfusionMatrixResultCombinationsParams {
	this := GetConfusionMatrixResultCombinationsParams{}
	this.ProjectId = projectId
	this.SessionRunIds = sessionRunIds
	this.X = x
	this.CustomMetricName = customMetricName
	return &this
}

// NewGetConfusionMatrixResultCombinationsParamsWithDefaults instantiates a new GetConfusionMatrixResultCombinationsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConfusionMatrixResultCombinationsParamsWithDefaults() *GetConfusionMatrixResultCombinationsParams {
	this := GetConfusionMatrixResultCombinationsParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GetConfusionMatrixResultCombinationsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetConfusionMatrixResultCombinationsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetConfusionMatrixResultCombinationsParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSessionRunIds returns the SessionRunIds field value
func (o *GetConfusionMatrixResultCombinationsParams) GetSessionRunIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SessionRunIds
}

// GetSessionRunIdsOk returns a tuple with the SessionRunIds field value
// and a boolean to check if the value has been set.
func (o *GetConfusionMatrixResultCombinationsParams) GetSessionRunIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionRunIds, true
}

// SetSessionRunIds sets field value
func (o *GetConfusionMatrixResultCombinationsParams) SetSessionRunIds(v []string) {
	o.SessionRunIds = v
}

// GetX returns the X field value
func (o *GetConfusionMatrixResultCombinationsParams) GetX() SplitAgg {
	if o == nil {
		var ret SplitAgg
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *GetConfusionMatrixResultCombinationsParams) GetXOk() (*SplitAgg, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *GetConfusionMatrixResultCombinationsParams) SetX(v SplitAgg) {
	o.X = v
}

// GetVerticalSplit returns the VerticalSplit field value if set, zero value otherwise.
func (o *GetConfusionMatrixResultCombinationsParams) GetVerticalSplit() SplitAgg {
	if o == nil || IsNil(o.VerticalSplit) {
		var ret SplitAgg
		return ret
	}
	return *o.VerticalSplit
}

// GetVerticalSplitOk returns a tuple with the VerticalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfusionMatrixResultCombinationsParams) GetVerticalSplitOk() (*SplitAgg, bool) {
	if o == nil || IsNil(o.VerticalSplit) {
		return nil, false
	}
	return o.VerticalSplit, true
}

// HasVerticalSplit returns a boolean if a field has been set.
func (o *GetConfusionMatrixResultCombinationsParams) HasVerticalSplit() bool {
	if o != nil && !IsNil(o.VerticalSplit) {
		return true
	}

	return false
}

// SetVerticalSplit gets a reference to the given SplitAgg and assigns it to the VerticalSplit field.
func (o *GetConfusionMatrixResultCombinationsParams) SetVerticalSplit(v SplitAgg) {
	o.VerticalSplit = &v
}

// GetHorizontalSplit returns the HorizontalSplit field value if set, zero value otherwise.
func (o *GetConfusionMatrixResultCombinationsParams) GetHorizontalSplit() SplitAgg {
	if o == nil || IsNil(o.HorizontalSplit) {
		var ret SplitAgg
		return ret
	}
	return *o.HorizontalSplit
}

// GetHorizontalSplitOk returns a tuple with the HorizontalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfusionMatrixResultCombinationsParams) GetHorizontalSplitOk() (*SplitAgg, bool) {
	if o == nil || IsNil(o.HorizontalSplit) {
		return nil, false
	}
	return o.HorizontalSplit, true
}

// HasHorizontalSplit returns a boolean if a field has been set.
func (o *GetConfusionMatrixResultCombinationsParams) HasHorizontalSplit() bool {
	if o != nil && !IsNil(o.HorizontalSplit) {
		return true
	}

	return false
}

// SetHorizontalSplit gets a reference to the given SplitAgg and assigns it to the HorizontalSplit field.
func (o *GetConfusionMatrixResultCombinationsParams) SetHorizontalSplit(v SplitAgg) {
	o.HorizontalSplit = &v
}

// GetInnerSplit returns the InnerSplit field value if set, zero value otherwise.
func (o *GetConfusionMatrixResultCombinationsParams) GetInnerSplit() SplitAgg {
	if o == nil || IsNil(o.InnerSplit) {
		var ret SplitAgg
		return ret
	}
	return *o.InnerSplit
}

// GetInnerSplitOk returns a tuple with the InnerSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfusionMatrixResultCombinationsParams) GetInnerSplitOk() (*SplitAgg, bool) {
	if o == nil || IsNil(o.InnerSplit) {
		return nil, false
	}
	return o.InnerSplit, true
}

// HasInnerSplit returns a boolean if a field has been set.
func (o *GetConfusionMatrixResultCombinationsParams) HasInnerSplit() bool {
	if o != nil && !IsNil(o.InnerSplit) {
		return true
	}

	return false
}

// SetInnerSplit gets a reference to the given SplitAgg and assigns it to the InnerSplit field.
func (o *GetConfusionMatrixResultCombinationsParams) SetInnerSplit(v SplitAgg) {
	o.InnerSplit = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *GetConfusionMatrixResultCombinationsParams) GetThreshold() float64 {
	if o == nil || IsNil(o.Threshold) {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfusionMatrixResultCombinationsParams) GetThresholdOk() (*float64, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *GetConfusionMatrixResultCombinationsParams) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *GetConfusionMatrixResultCombinationsParams) SetThreshold(v float64) {
	o.Threshold = &v
}

// GetCustomMetricName returns the CustomMetricName field value
func (o *GetConfusionMatrixResultCombinationsParams) GetCustomMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomMetricName
}

// GetCustomMetricNameOk returns a tuple with the CustomMetricName field value
// and a boolean to check if the value has been set.
func (o *GetConfusionMatrixResultCombinationsParams) GetCustomMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomMetricName, true
}

// SetCustomMetricName sets field value
func (o *GetConfusionMatrixResultCombinationsParams) SetCustomMetricName(v string) {
	o.CustomMetricName = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *GetConfusionMatrixResultCombinationsParams) GetFilters() []ESFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []ESFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfusionMatrixResultCombinationsParams) GetFiltersOk() ([]ESFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *GetConfusionMatrixResultCombinationsParams) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []ESFilter and assigns it to the Filters field.
func (o *GetConfusionMatrixResultCombinationsParams) SetFilters(v []ESFilter) {
	o.Filters = v
}

// GetLastEpochOnly returns the LastEpochOnly field value if set, zero value otherwise.
func (o *GetConfusionMatrixResultCombinationsParams) GetLastEpochOnly() bool {
	if o == nil || IsNil(o.LastEpochOnly) {
		var ret bool
		return ret
	}
	return *o.LastEpochOnly
}

// GetLastEpochOnlyOk returns a tuple with the LastEpochOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfusionMatrixResultCombinationsParams) GetLastEpochOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.LastEpochOnly) {
		return nil, false
	}
	return o.LastEpochOnly, true
}

// HasLastEpochOnly returns a boolean if a field has been set.
func (o *GetConfusionMatrixResultCombinationsParams) HasLastEpochOnly() bool {
	if o != nil && !IsNil(o.LastEpochOnly) {
		return true
	}

	return false
}

// SetLastEpochOnly gets a reference to the given bool and assigns it to the LastEpochOnly field.
func (o *GetConfusionMatrixResultCombinationsParams) SetLastEpochOnly(v bool) {
	o.LastEpochOnly = &v
}

// GetFilterLabels returns the FilterLabels field value if set, zero value otherwise.
func (o *GetConfusionMatrixResultCombinationsParams) GetFilterLabels() []string {
	if o == nil || IsNil(o.FilterLabels) {
		var ret []string
		return ret
	}
	return o.FilterLabels
}

// GetFilterLabelsOk returns a tuple with the FilterLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfusionMatrixResultCombinationsParams) GetFilterLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.FilterLabels) {
		return nil, false
	}
	return o.FilterLabels, true
}

// HasFilterLabels returns a boolean if a field has been set.
func (o *GetConfusionMatrixResultCombinationsParams) HasFilterLabels() bool {
	if o != nil && !IsNil(o.FilterLabels) {
		return true
	}

	return false
}

// SetFilterLabels gets a reference to the given []string and assigns it to the FilterLabels field.
func (o *GetConfusionMatrixResultCombinationsParams) SetFilterLabels(v []string) {
	o.FilterLabels = v
}

func (o GetConfusionMatrixResultCombinationsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetConfusionMatrixResultCombinationsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["sessionRunIds"] = o.SessionRunIds
	toSerialize["x"] = o.X
	if !IsNil(o.VerticalSplit) {
		toSerialize["verticalSplit"] = o.VerticalSplit
	}
	if !IsNil(o.HorizontalSplit) {
		toSerialize["horizontalSplit"] = o.HorizontalSplit
	}
	if !IsNil(o.InnerSplit) {
		toSerialize["innerSplit"] = o.InnerSplit
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	toSerialize["customMetricName"] = o.CustomMetricName
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.LastEpochOnly) {
		toSerialize["lastEpochOnly"] = o.LastEpochOnly
	}
	if !IsNil(o.FilterLabels) {
		toSerialize["filterLabels"] = o.FilterLabels
	}
	return toSerialize, nil
}

type NullableGetConfusionMatrixResultCombinationsParams struct {
	value *GetConfusionMatrixResultCombinationsParams
	isSet bool
}

func (v NullableGetConfusionMatrixResultCombinationsParams) Get() *GetConfusionMatrixResultCombinationsParams {
	return v.value
}

func (v *NullableGetConfusionMatrixResultCombinationsParams) Set(val *GetConfusionMatrixResultCombinationsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConfusionMatrixResultCombinationsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConfusionMatrixResultCombinationsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConfusionMatrixResultCombinationsParams(val *GetConfusionMatrixResultCombinationsParams) *NullableGetConfusionMatrixResultCombinationsParams {
	return &NullableGetConfusionMatrixResultCombinationsParams{value: val, isSet: true}
}

func (v NullableGetConfusionMatrixResultCombinationsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConfusionMatrixResultCombinationsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
