/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.464
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the MultiChartsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiChartsParams{}

// MultiChartsParams struct for MultiChartsParams
type MultiChartsParams struct {
	ProjectId            string               `json:"projectId"`
	XField               string               `json:"xField"`
	YField               string               `json:"yField"`
	AggregationMethod    AggregationMethod    `json:"aggregationMethod"`
	SessionRunIds        []string             `json:"sessionRunIds"`
	VerticalSplit        *string              `json:"verticalSplit,omitempty"`
	HorizontalSplit      *string              `json:"horizontalSplit,omitempty"`
	InnerSplit           *string              `json:"innerSplit,omitempty"`
	Filters              []ESFilter           `json:"filters,omitempty"`
	DataDistributionType DataDistributionType `json:"dataDistributionType"`
	OrderByParam         *string              `json:"orderByParam,omitempty"`
	OrderParams          *OrderType           `json:"orderParams,omitempty"`
	XAxisSizeInterval    float64              `json:"xAxisSizeInterval"`
	LastEpochOnly        *bool                `json:"lastEpochOnly,omitempty"`
}

// NewMultiChartsParams instantiates a new MultiChartsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiChartsParams(projectId string, xField string, yField string, aggregationMethod AggregationMethod, sessionRunIds []string, dataDistributionType DataDistributionType, xAxisSizeInterval float64) *MultiChartsParams {
	this := MultiChartsParams{}
	this.ProjectId = projectId
	this.XField = xField
	this.YField = yField
	this.AggregationMethod = aggregationMethod
	this.SessionRunIds = sessionRunIds
	this.DataDistributionType = dataDistributionType
	this.XAxisSizeInterval = xAxisSizeInterval
	return &this
}

// NewMultiChartsParamsWithDefaults instantiates a new MultiChartsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiChartsParamsWithDefaults() *MultiChartsParams {
	this := MultiChartsParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *MultiChartsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *MultiChartsParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetXField returns the XField field value
func (o *MultiChartsParams) GetXField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.XField
}

// GetXFieldOk returns a tuple with the XField field value
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetXFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XField, true
}

// SetXField sets field value
func (o *MultiChartsParams) SetXField(v string) {
	o.XField = v
}

// GetYField returns the YField field value
func (o *MultiChartsParams) GetYField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YField
}

// GetYFieldOk returns a tuple with the YField field value
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetYFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YField, true
}

// SetYField sets field value
func (o *MultiChartsParams) SetYField(v string) {
	o.YField = v
}

// GetAggregationMethod returns the AggregationMethod field value
func (o *MultiChartsParams) GetAggregationMethod() AggregationMethod {
	if o == nil {
		var ret AggregationMethod
		return ret
	}

	return o.AggregationMethod
}

// GetAggregationMethodOk returns a tuple with the AggregationMethod field value
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetAggregationMethodOk() (*AggregationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AggregationMethod, true
}

// SetAggregationMethod sets field value
func (o *MultiChartsParams) SetAggregationMethod(v AggregationMethod) {
	o.AggregationMethod = v
}

// GetSessionRunIds returns the SessionRunIds field value
func (o *MultiChartsParams) GetSessionRunIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SessionRunIds
}

// GetSessionRunIdsOk returns a tuple with the SessionRunIds field value
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetSessionRunIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionRunIds, true
}

// SetSessionRunIds sets field value
func (o *MultiChartsParams) SetSessionRunIds(v []string) {
	o.SessionRunIds = v
}

// GetVerticalSplit returns the VerticalSplit field value if set, zero value otherwise.
func (o *MultiChartsParams) GetVerticalSplit() string {
	if o == nil || IsNil(o.VerticalSplit) {
		var ret string
		return ret
	}
	return *o.VerticalSplit
}

// GetVerticalSplitOk returns a tuple with the VerticalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetVerticalSplitOk() (*string, bool) {
	if o == nil || IsNil(o.VerticalSplit) {
		return nil, false
	}
	return o.VerticalSplit, true
}

// HasVerticalSplit returns a boolean if a field has been set.
func (o *MultiChartsParams) HasVerticalSplit() bool {
	if o != nil && !IsNil(o.VerticalSplit) {
		return true
	}

	return false
}

// SetVerticalSplit gets a reference to the given string and assigns it to the VerticalSplit field.
func (o *MultiChartsParams) SetVerticalSplit(v string) {
	o.VerticalSplit = &v
}

// GetHorizontalSplit returns the HorizontalSplit field value if set, zero value otherwise.
func (o *MultiChartsParams) GetHorizontalSplit() string {
	if o == nil || IsNil(o.HorizontalSplit) {
		var ret string
		return ret
	}
	return *o.HorizontalSplit
}

// GetHorizontalSplitOk returns a tuple with the HorizontalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetHorizontalSplitOk() (*string, bool) {
	if o == nil || IsNil(o.HorizontalSplit) {
		return nil, false
	}
	return o.HorizontalSplit, true
}

// HasHorizontalSplit returns a boolean if a field has been set.
func (o *MultiChartsParams) HasHorizontalSplit() bool {
	if o != nil && !IsNil(o.HorizontalSplit) {
		return true
	}

	return false
}

// SetHorizontalSplit gets a reference to the given string and assigns it to the HorizontalSplit field.
func (o *MultiChartsParams) SetHorizontalSplit(v string) {
	o.HorizontalSplit = &v
}

// GetInnerSplit returns the InnerSplit field value if set, zero value otherwise.
func (o *MultiChartsParams) GetInnerSplit() string {
	if o == nil || IsNil(o.InnerSplit) {
		var ret string
		return ret
	}
	return *o.InnerSplit
}

// GetInnerSplitOk returns a tuple with the InnerSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetInnerSplitOk() (*string, bool) {
	if o == nil || IsNil(o.InnerSplit) {
		return nil, false
	}
	return o.InnerSplit, true
}

// HasInnerSplit returns a boolean if a field has been set.
func (o *MultiChartsParams) HasInnerSplit() bool {
	if o != nil && !IsNil(o.InnerSplit) {
		return true
	}

	return false
}

// SetInnerSplit gets a reference to the given string and assigns it to the InnerSplit field.
func (o *MultiChartsParams) SetInnerSplit(v string) {
	o.InnerSplit = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *MultiChartsParams) GetFilters() []ESFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []ESFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetFiltersOk() ([]ESFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *MultiChartsParams) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []ESFilter and assigns it to the Filters field.
func (o *MultiChartsParams) SetFilters(v []ESFilter) {
	o.Filters = v
}

// GetDataDistributionType returns the DataDistributionType field value
func (o *MultiChartsParams) GetDataDistributionType() DataDistributionType {
	if o == nil {
		var ret DataDistributionType
		return ret
	}

	return o.DataDistributionType
}

// GetDataDistributionTypeOk returns a tuple with the DataDistributionType field value
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetDataDistributionTypeOk() (*DataDistributionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataDistributionType, true
}

// SetDataDistributionType sets field value
func (o *MultiChartsParams) SetDataDistributionType(v DataDistributionType) {
	o.DataDistributionType = v
}

// GetOrderByParam returns the OrderByParam field value if set, zero value otherwise.
func (o *MultiChartsParams) GetOrderByParam() string {
	if o == nil || IsNil(o.OrderByParam) {
		var ret string
		return ret
	}
	return *o.OrderByParam
}

// GetOrderByParamOk returns a tuple with the OrderByParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetOrderByParamOk() (*string, bool) {
	if o == nil || IsNil(o.OrderByParam) {
		return nil, false
	}
	return o.OrderByParam, true
}

// HasOrderByParam returns a boolean if a field has been set.
func (o *MultiChartsParams) HasOrderByParam() bool {
	if o != nil && !IsNil(o.OrderByParam) {
		return true
	}

	return false
}

// SetOrderByParam gets a reference to the given string and assigns it to the OrderByParam field.
func (o *MultiChartsParams) SetOrderByParam(v string) {
	o.OrderByParam = &v
}

// GetOrderParams returns the OrderParams field value if set, zero value otherwise.
func (o *MultiChartsParams) GetOrderParams() OrderType {
	if o == nil || IsNil(o.OrderParams) {
		var ret OrderType
		return ret
	}
	return *o.OrderParams
}

// GetOrderParamsOk returns a tuple with the OrderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetOrderParamsOk() (*OrderType, bool) {
	if o == nil || IsNil(o.OrderParams) {
		return nil, false
	}
	return o.OrderParams, true
}

// HasOrderParams returns a boolean if a field has been set.
func (o *MultiChartsParams) HasOrderParams() bool {
	if o != nil && !IsNil(o.OrderParams) {
		return true
	}

	return false
}

// SetOrderParams gets a reference to the given OrderType and assigns it to the OrderParams field.
func (o *MultiChartsParams) SetOrderParams(v OrderType) {
	o.OrderParams = &v
}

// GetXAxisSizeInterval returns the XAxisSizeInterval field value
func (o *MultiChartsParams) GetXAxisSizeInterval() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.XAxisSizeInterval
}

// GetXAxisSizeIntervalOk returns a tuple with the XAxisSizeInterval field value
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetXAxisSizeIntervalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XAxisSizeInterval, true
}

// SetXAxisSizeInterval sets field value
func (o *MultiChartsParams) SetXAxisSizeInterval(v float64) {
	o.XAxisSizeInterval = v
}

// GetLastEpochOnly returns the LastEpochOnly field value if set, zero value otherwise.
func (o *MultiChartsParams) GetLastEpochOnly() bool {
	if o == nil || IsNil(o.LastEpochOnly) {
		var ret bool
		return ret
	}
	return *o.LastEpochOnly
}

// GetLastEpochOnlyOk returns a tuple with the LastEpochOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiChartsParams) GetLastEpochOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.LastEpochOnly) {
		return nil, false
	}
	return o.LastEpochOnly, true
}

// HasLastEpochOnly returns a boolean if a field has been set.
func (o *MultiChartsParams) HasLastEpochOnly() bool {
	if o != nil && !IsNil(o.LastEpochOnly) {
		return true
	}

	return false
}

// SetLastEpochOnly gets a reference to the given bool and assigns it to the LastEpochOnly field.
func (o *MultiChartsParams) SetLastEpochOnly(v bool) {
	o.LastEpochOnly = &v
}

func (o MultiChartsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiChartsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["xField"] = o.XField
	toSerialize["yField"] = o.YField
	toSerialize["aggregationMethod"] = o.AggregationMethod
	toSerialize["sessionRunIds"] = o.SessionRunIds
	if !IsNil(o.VerticalSplit) {
		toSerialize["verticalSplit"] = o.VerticalSplit
	}
	if !IsNil(o.HorizontalSplit) {
		toSerialize["horizontalSplit"] = o.HorizontalSplit
	}
	if !IsNil(o.InnerSplit) {
		toSerialize["innerSplit"] = o.InnerSplit
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	toSerialize["dataDistributionType"] = o.DataDistributionType
	if !IsNil(o.OrderByParam) {
		toSerialize["orderByParam"] = o.OrderByParam
	}
	if !IsNil(o.OrderParams) {
		toSerialize["orderParams"] = o.OrderParams
	}
	toSerialize["xAxisSizeInterval"] = o.XAxisSizeInterval
	if !IsNil(o.LastEpochOnly) {
		toSerialize["lastEpochOnly"] = o.LastEpochOnly
	}
	return toSerialize, nil
}

type NullableMultiChartsParams struct {
	value *MultiChartsParams
	isSet bool
}

func (v NullableMultiChartsParams) Get() *MultiChartsParams {
	return v.value
}

func (v *NullableMultiChartsParams) Set(val *MultiChartsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiChartsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiChartsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiChartsParams(val *MultiChartsParams) *NullableMultiChartsParams {
	return &NullableMultiChartsParams{value: val, isSet: true}
}

func (v NullableMultiChartsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiChartsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
