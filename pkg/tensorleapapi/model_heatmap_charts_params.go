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

// checks if the HeatmapChartsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeatmapChartsParams{}

// HeatmapChartsParams struct for HeatmapChartsParams
type HeatmapChartsParams struct {
	ProjectId              string               `json:"projectId"`
	XField                 string               `json:"xField"`
	XAxisSizeInterval      float64              `json:"xAxisSizeInterval"`
	XDataDistributionType  DataDistributionType `json:"xDataDistributionType"`
	XOrderParams           *OrderType           `json:"xOrderParams,omitempty"`
	YField                 string               `json:"yField"`
	YAxisSizeInterval      float64              `json:"yAxisSizeInterval"`
	YDataDistributionType  DataDistributionType `json:"yDataDistributionType"`
	YOrderParams           *OrderType           `json:"yOrderParams,omitempty"`
	ColorField             string               `json:"colorField"`
	ColorAggregationMethod AggregationMethod    `json:"colorAggregationMethod"`
	SessionRunIds          []string             `json:"sessionRunIds"`
	VerticalSplit          *string              `json:"verticalSplit,omitempty"`
	HorizontalSplit        *string              `json:"horizontalSplit,omitempty"`
	Filters                []ESFilter           `json:"filters,omitempty"`
	LastEpochOnly          *bool                `json:"lastEpochOnly,omitempty"`
}

// NewHeatmapChartsParams instantiates a new HeatmapChartsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeatmapChartsParams(projectId string, xField string, xAxisSizeInterval float64, xDataDistributionType DataDistributionType, yField string, yAxisSizeInterval float64, yDataDistributionType DataDistributionType, colorField string, colorAggregationMethod AggregationMethod, sessionRunIds []string) *HeatmapChartsParams {
	this := HeatmapChartsParams{}
	this.ProjectId = projectId
	this.XField = xField
	this.XAxisSizeInterval = xAxisSizeInterval
	this.XDataDistributionType = xDataDistributionType
	this.YField = yField
	this.YAxisSizeInterval = yAxisSizeInterval
	this.YDataDistributionType = yDataDistributionType
	this.ColorField = colorField
	this.ColorAggregationMethod = colorAggregationMethod
	this.SessionRunIds = sessionRunIds
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

// GetXField returns the XField field value
func (o *HeatmapChartsParams) GetXField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.XField
}

// GetXFieldOk returns a tuple with the XField field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetXFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XField, true
}

// SetXField sets field value
func (o *HeatmapChartsParams) SetXField(v string) {
	o.XField = v
}

// GetXAxisSizeInterval returns the XAxisSizeInterval field value
func (o *HeatmapChartsParams) GetXAxisSizeInterval() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.XAxisSizeInterval
}

// GetXAxisSizeIntervalOk returns a tuple with the XAxisSizeInterval field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetXAxisSizeIntervalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XAxisSizeInterval, true
}

// SetXAxisSizeInterval sets field value
func (o *HeatmapChartsParams) SetXAxisSizeInterval(v float64) {
	o.XAxisSizeInterval = v
}

// GetXDataDistributionType returns the XDataDistributionType field value
func (o *HeatmapChartsParams) GetXDataDistributionType() DataDistributionType {
	if o == nil {
		var ret DataDistributionType
		return ret
	}

	return o.XDataDistributionType
}

// GetXDataDistributionTypeOk returns a tuple with the XDataDistributionType field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetXDataDistributionTypeOk() (*DataDistributionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XDataDistributionType, true
}

// SetXDataDistributionType sets field value
func (o *HeatmapChartsParams) SetXDataDistributionType(v DataDistributionType) {
	o.XDataDistributionType = v
}

// GetXOrderParams returns the XOrderParams field value if set, zero value otherwise.
func (o *HeatmapChartsParams) GetXOrderParams() OrderType {
	if o == nil || IsNil(o.XOrderParams) {
		var ret OrderType
		return ret
	}
	return *o.XOrderParams
}

// GetXOrderParamsOk returns a tuple with the XOrderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetXOrderParamsOk() (*OrderType, bool) {
	if o == nil || IsNil(o.XOrderParams) {
		return nil, false
	}
	return o.XOrderParams, true
}

// HasXOrderParams returns a boolean if a field has been set.
func (o *HeatmapChartsParams) HasXOrderParams() bool {
	if o != nil && !IsNil(o.XOrderParams) {
		return true
	}

	return false
}

// SetXOrderParams gets a reference to the given OrderType and assigns it to the XOrderParams field.
func (o *HeatmapChartsParams) SetXOrderParams(v OrderType) {
	o.XOrderParams = &v
}

// GetYField returns the YField field value
func (o *HeatmapChartsParams) GetYField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YField
}

// GetYFieldOk returns a tuple with the YField field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetYFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YField, true
}

// SetYField sets field value
func (o *HeatmapChartsParams) SetYField(v string) {
	o.YField = v
}

// GetYAxisSizeInterval returns the YAxisSizeInterval field value
func (o *HeatmapChartsParams) GetYAxisSizeInterval() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.YAxisSizeInterval
}

// GetYAxisSizeIntervalOk returns a tuple with the YAxisSizeInterval field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetYAxisSizeIntervalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YAxisSizeInterval, true
}

// SetYAxisSizeInterval sets field value
func (o *HeatmapChartsParams) SetYAxisSizeInterval(v float64) {
	o.YAxisSizeInterval = v
}

// GetYDataDistributionType returns the YDataDistributionType field value
func (o *HeatmapChartsParams) GetYDataDistributionType() DataDistributionType {
	if o == nil {
		var ret DataDistributionType
		return ret
	}

	return o.YDataDistributionType
}

// GetYDataDistributionTypeOk returns a tuple with the YDataDistributionType field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetYDataDistributionTypeOk() (*DataDistributionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YDataDistributionType, true
}

// SetYDataDistributionType sets field value
func (o *HeatmapChartsParams) SetYDataDistributionType(v DataDistributionType) {
	o.YDataDistributionType = v
}

// GetYOrderParams returns the YOrderParams field value if set, zero value otherwise.
func (o *HeatmapChartsParams) GetYOrderParams() OrderType {
	if o == nil || IsNil(o.YOrderParams) {
		var ret OrderType
		return ret
	}
	return *o.YOrderParams
}

// GetYOrderParamsOk returns a tuple with the YOrderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetYOrderParamsOk() (*OrderType, bool) {
	if o == nil || IsNil(o.YOrderParams) {
		return nil, false
	}
	return o.YOrderParams, true
}

// HasYOrderParams returns a boolean if a field has been set.
func (o *HeatmapChartsParams) HasYOrderParams() bool {
	if o != nil && !IsNil(o.YOrderParams) {
		return true
	}

	return false
}

// SetYOrderParams gets a reference to the given OrderType and assigns it to the YOrderParams field.
func (o *HeatmapChartsParams) SetYOrderParams(v OrderType) {
	o.YOrderParams = &v
}

// GetColorField returns the ColorField field value
func (o *HeatmapChartsParams) GetColorField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ColorField
}

// GetColorFieldOk returns a tuple with the ColorField field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetColorFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColorField, true
}

// SetColorField sets field value
func (o *HeatmapChartsParams) SetColorField(v string) {
	o.ColorField = v
}

// GetColorAggregationMethod returns the ColorAggregationMethod field value
func (o *HeatmapChartsParams) GetColorAggregationMethod() AggregationMethod {
	if o == nil {
		var ret AggregationMethod
		return ret
	}

	return o.ColorAggregationMethod
}

// GetColorAggregationMethodOk returns a tuple with the ColorAggregationMethod field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetColorAggregationMethodOk() (*AggregationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColorAggregationMethod, true
}

// SetColorAggregationMethod sets field value
func (o *HeatmapChartsParams) SetColorAggregationMethod(v AggregationMethod) {
	o.ColorAggregationMethod = v
}

// GetSessionRunIds returns the SessionRunIds field value
func (o *HeatmapChartsParams) GetSessionRunIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SessionRunIds
}

// GetSessionRunIdsOk returns a tuple with the SessionRunIds field value
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetSessionRunIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionRunIds, true
}

// SetSessionRunIds sets field value
func (o *HeatmapChartsParams) SetSessionRunIds(v []string) {
	o.SessionRunIds = v
}

// GetVerticalSplit returns the VerticalSplit field value if set, zero value otherwise.
func (o *HeatmapChartsParams) GetVerticalSplit() string {
	if o == nil || IsNil(o.VerticalSplit) {
		var ret string
		return ret
	}
	return *o.VerticalSplit
}

// GetVerticalSplitOk returns a tuple with the VerticalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetVerticalSplitOk() (*string, bool) {
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

// SetVerticalSplit gets a reference to the given string and assigns it to the VerticalSplit field.
func (o *HeatmapChartsParams) SetVerticalSplit(v string) {
	o.VerticalSplit = &v
}

// GetHorizontalSplit returns the HorizontalSplit field value if set, zero value otherwise.
func (o *HeatmapChartsParams) GetHorizontalSplit() string {
	if o == nil || IsNil(o.HorizontalSplit) {
		var ret string
		return ret
	}
	return *o.HorizontalSplit
}

// GetHorizontalSplitOk returns a tuple with the HorizontalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetHorizontalSplitOk() (*string, bool) {
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

// SetHorizontalSplit gets a reference to the given string and assigns it to the HorizontalSplit field.
func (o *HeatmapChartsParams) SetHorizontalSplit(v string) {
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

// GetLastEpochOnly returns the LastEpochOnly field value if set, zero value otherwise.
func (o *HeatmapChartsParams) GetLastEpochOnly() bool {
	if o == nil || IsNil(o.LastEpochOnly) {
		var ret bool
		return ret
	}
	return *o.LastEpochOnly
}

// GetLastEpochOnlyOk returns a tuple with the LastEpochOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapChartsParams) GetLastEpochOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.LastEpochOnly) {
		return nil, false
	}
	return o.LastEpochOnly, true
}

// HasLastEpochOnly returns a boolean if a field has been set.
func (o *HeatmapChartsParams) HasLastEpochOnly() bool {
	if o != nil && !IsNil(o.LastEpochOnly) {
		return true
	}

	return false
}

// SetLastEpochOnly gets a reference to the given bool and assigns it to the LastEpochOnly field.
func (o *HeatmapChartsParams) SetLastEpochOnly(v bool) {
	o.LastEpochOnly = &v
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
	toSerialize["xField"] = o.XField
	toSerialize["xAxisSizeInterval"] = o.XAxisSizeInterval
	toSerialize["xDataDistributionType"] = o.XDataDistributionType
	if !IsNil(o.XOrderParams) {
		toSerialize["xOrderParams"] = o.XOrderParams
	}
	toSerialize["yField"] = o.YField
	toSerialize["yAxisSizeInterval"] = o.YAxisSizeInterval
	toSerialize["yDataDistributionType"] = o.YDataDistributionType
	if !IsNil(o.YOrderParams) {
		toSerialize["yOrderParams"] = o.YOrderParams
	}
	toSerialize["colorField"] = o.ColorField
	toSerialize["colorAggregationMethod"] = o.ColorAggregationMethod
	toSerialize["sessionRunIds"] = o.SessionRunIds
	if !IsNil(o.VerticalSplit) {
		toSerialize["verticalSplit"] = o.VerticalSplit
	}
	if !IsNil(o.HorizontalSplit) {
		toSerialize["horizontalSplit"] = o.HorizontalSplit
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.LastEpochOnly) {
		toSerialize["lastEpochOnly"] = o.LastEpochOnly
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
