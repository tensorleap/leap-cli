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

// checks if the HeatmapChartsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeatmapChartsParams{}

// HeatmapChartsParams struct for HeatmapChartsParams
type HeatmapChartsParams struct {
	ProjectId           string              `json:"projectId"`
	X                   SplitAgg            `json:"x"`
	Y                   SplitAgg            `json:"y"`
	Color               Aggregations        `json:"color"`
	SessionRunsToEpochs []SessionRunToEpoch `json:"sessionRunsToEpochs"`
	ShowAllEpochs       bool                `json:"showAllEpochs"`
	VerticalSplit       *SplitAgg           `json:"verticalSplit,omitempty"`
	HorizontalSplit     *SplitAgg           `json:"horizontalSplit,omitempty"`
	Filters             []ESFilter          `json:"filters,omitempty"`
}

// NewHeatmapChartsParams instantiates a new HeatmapChartsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeatmapChartsParams(projectId string, x SplitAgg, y SplitAgg, color Aggregations, sessionRunsToEpochs []SessionRunToEpoch, showAllEpochs bool) *HeatmapChartsParams {
	this := HeatmapChartsParams{}
	this.ProjectId = projectId
	this.X = x
	this.Y = y
	this.Color = color
	this.SessionRunsToEpochs = sessionRunsToEpochs
	this.ShowAllEpochs = showAllEpochs
	return &this
}

// NewHeatmapChartsParamsWithDefaults instantiates a new HeatmapChartsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeatmapChartsParamsWithDefaults() *HeatmapChartsParams {
	this := HeatmapChartsParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *HeatmapChartsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *HeatmapChartsParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetX returns the X field value
func (o *HeatmapChartsParams) GetX() SplitAgg {
	if o == nil {
		var ret SplitAgg
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetXOk() (*SplitAgg, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *HeatmapChartsParams) SetX(v SplitAgg) {
	o.X = v
}

// GetY returns the Y field value
func (o *HeatmapChartsParams) GetY() SplitAgg {
	if o == nil {
		var ret SplitAgg
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetYOk() (*SplitAgg, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *HeatmapChartsParams) SetY(v SplitAgg) {
	o.Y = v
}

// GetColor returns the Color field value
func (o *HeatmapChartsParams) GetColor() Aggregations {
	if o == nil {
		var ret Aggregations
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetColorOk() (*Aggregations, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *HeatmapChartsParams) SetColor(v Aggregations) {
	o.Color = v
}

// GetSessionRunsToEpochs returns the SessionRunsToEpochs field value
func (o *HeatmapChartsParams) GetSessionRunsToEpochs() []SessionRunToEpoch {
	if o == nil {
		var ret []SessionRunToEpoch
		return ret
	}

	return o.SessionRunsToEpochs
}

// GetSessionRunsToEpochsOk returns a tuple with the SessionRunsToEpochs field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetSessionRunsToEpochsOk() ([]SessionRunToEpoch, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionRunsToEpochs, true
}

// SetSessionRunsToEpochs sets field value
func (o *HeatmapChartsParams) SetSessionRunsToEpochs(v []SessionRunToEpoch) {
	o.SessionRunsToEpochs = v
}

// GetShowAllEpochs returns the ShowAllEpochs field value
func (o *HeatmapChartsParams) GetShowAllEpochs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowAllEpochs
}

// GetShowAllEpochsOk returns a tuple with the ShowAllEpochs field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetShowAllEpochsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowAllEpochs, true
}

// SetShowAllEpochs sets field value
func (o *HeatmapChartsParams) SetShowAllEpochs(v bool) {
	o.ShowAllEpochs = v
}

// GetVerticalSplit returns the VerticalSplit field value if set, zero value otherwise.
func (o *HeatmapChartsParams) GetVerticalSplit() SplitAgg {
	if o == nil || IsNil(o.VerticalSplit) {
		var ret SplitAgg
		return ret
	}
	return *o.VerticalSplit
}

// GetVerticalSplitOk returns a tuple with the VerticalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetVerticalSplitOk() (*SplitAgg, bool) {
	if o == nil || IsNil(o.VerticalSplit) {
		return nil, false
	}
	return o.VerticalSplit, true
}

// HasVerticalSplit returns a boolean if a field has been set.
func (o *HeatmapChartsParams) HasVerticalSplit() bool {
	if o != nil && !IsNil(o.VerticalSplit) {
		return true
	}

	return false
}

// SetVerticalSplit gets a reference to the given SplitAgg and assigns it to the VerticalSplit field.
func (o *HeatmapChartsParams) SetVerticalSplit(v SplitAgg) {
	o.VerticalSplit = &v
}

// GetHorizontalSplit returns the HorizontalSplit field value if set, zero value otherwise.
func (o *HeatmapChartsParams) GetHorizontalSplit() SplitAgg {
	if o == nil || IsNil(o.HorizontalSplit) {
		var ret SplitAgg
		return ret
	}
	return *o.HorizontalSplit
}

// GetHorizontalSplitOk returns a tuple with the HorizontalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetHorizontalSplitOk() (*SplitAgg, bool) {
	if o == nil || IsNil(o.HorizontalSplit) {
		return nil, false
	}
	return o.HorizontalSplit, true
}

// HasHorizontalSplit returns a boolean if a field has been set.
func (o *HeatmapChartsParams) HasHorizontalSplit() bool {
	if o != nil && !IsNil(o.HorizontalSplit) {
		return true
	}

	return false
}

// SetHorizontalSplit gets a reference to the given SplitAgg and assigns it to the HorizontalSplit field.
func (o *HeatmapChartsParams) SetHorizontalSplit(v SplitAgg) {
	o.HorizontalSplit = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *HeatmapChartsParams) GetFilters() []ESFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []ESFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetFiltersOk() ([]ESFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *HeatmapChartsParams) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []ESFilter and assigns it to the Filters field.
func (o *HeatmapChartsParams) SetFilters(v []ESFilter) {
	o.Filters = v
}

func (o HeatmapChartsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeatmapChartsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	toSerialize["color"] = o.Color
	toSerialize["sessionRunsToEpochs"] = o.SessionRunsToEpochs
	toSerialize["showAllEpochs"] = o.ShowAllEpochs
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

type NullableHeatmapChartsParams struct {
	value *HeatmapChartsParams
	isSet bool
}

func (v NullableHeatmapChartsParams) Get() *HeatmapChartsParams {
	return v.value
}

func (v *NullableHeatmapChartsParams) Set(val *HeatmapChartsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableHeatmapChartsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableHeatmapChartsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeatmapChartsParams(val *HeatmapChartsParams) *NullableHeatmapChartsParams {
	return &NullableHeatmapChartsParams{value: val, isSet: true}
}

func (v NullableHeatmapChartsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeatmapChartsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
