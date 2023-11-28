/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.417
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GenericDataQueryParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericDataQueryParams{}

// GenericDataQueryParams struct for GenericDataQueryParams
type GenericDataQueryParams struct {
	ProjectId       string              `json:"projectId"`
	SessionRunIds   []string            `json:"sessionRunIds"`
	Filters         []ESFilter          `json:"filters,omitempty"`
	VerticalSplit   *string             `json:"verticalSplit,omitempty"`
	HorizontalSplit *string             `json:"horizontalSplit,omitempty"`
	Aggregations    []Aggregations      `json:"aggregations"`
	Buckets         []BucketAggregation `json:"buckets"`
	LastEpochOnly   *bool               `json:"lastEpochOnly,omitempty"`
}

// NewGenericDataQueryParams instantiates a new GenericDataQueryParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericDataQueryParams(projectId string, sessionRunIds []string, aggregations []Aggregations, buckets []BucketAggregation) *GenericDataQueryParams {
	this := GenericDataQueryParams{}
	this.ProjectId = projectId
	this.SessionRunIds = sessionRunIds
	this.Aggregations = aggregations
	this.Buckets = buckets
	return &this
}

// NewGenericDataQueryParamsWithDefaults instantiates a new GenericDataQueryParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericDataQueryParamsWithDefaults() *GenericDataQueryParams {
	this := GenericDataQueryParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GenericDataQueryParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GenericDataQueryParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GenericDataQueryParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSessionRunIds returns the SessionRunIds field value
func (o *GenericDataQueryParams) GetSessionRunIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SessionRunIds
}

// GetSessionRunIdsOk returns a tuple with the SessionRunIds field value
// and a boolean to check if the value has been set.
func (o *GenericDataQueryParams) GetSessionRunIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionRunIds, true
}

// SetSessionRunIds sets field value
func (o *GenericDataQueryParams) SetSessionRunIds(v []string) {
	o.SessionRunIds = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *GenericDataQueryParams) GetFilters() []ESFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []ESFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericDataQueryParams) GetFiltersOk() ([]ESFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *GenericDataQueryParams) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []ESFilter and assigns it to the Filters field.
func (o *GenericDataQueryParams) SetFilters(v []ESFilter) {
	o.Filters = v
}

// GetVerticalSplit returns the VerticalSplit field value if set, zero value otherwise.
func (o *GenericDataQueryParams) GetVerticalSplit() string {
	if o == nil || IsNil(o.VerticalSplit) {
		var ret string
		return ret
	}
	return *o.VerticalSplit
}

// GetVerticalSplitOk returns a tuple with the VerticalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericDataQueryParams) GetVerticalSplitOk() (*string, bool) {
	if o == nil || IsNil(o.VerticalSplit) {
		return nil, false
	}
	return o.VerticalSplit, true
}

// HasVerticalSplit returns a boolean if a field has been set.
func (o *GenericDataQueryParams) HasVerticalSplit() bool {
	if o != nil && !IsNil(o.VerticalSplit) {
		return true
	}

	return false
}

// SetVerticalSplit gets a reference to the given string and assigns it to the VerticalSplit field.
func (o *GenericDataQueryParams) SetVerticalSplit(v string) {
	o.VerticalSplit = &v
}

// GetHorizontalSplit returns the HorizontalSplit field value if set, zero value otherwise.
func (o *GenericDataQueryParams) GetHorizontalSplit() string {
	if o == nil || IsNil(o.HorizontalSplit) {
		var ret string
		return ret
	}
	return *o.HorizontalSplit
}

// GetHorizontalSplitOk returns a tuple with the HorizontalSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericDataQueryParams) GetHorizontalSplitOk() (*string, bool) {
	if o == nil || IsNil(o.HorizontalSplit) {
		return nil, false
	}
	return o.HorizontalSplit, true
}

// HasHorizontalSplit returns a boolean if a field has been set.
func (o *GenericDataQueryParams) HasHorizontalSplit() bool {
	if o != nil && !IsNil(o.HorizontalSplit) {
		return true
	}

	return false
}

// SetHorizontalSplit gets a reference to the given string and assigns it to the HorizontalSplit field.
func (o *GenericDataQueryParams) SetHorizontalSplit(v string) {
	o.HorizontalSplit = &v
}

// GetAggregations returns the Aggregations field value
func (o *GenericDataQueryParams) GetAggregations() []Aggregations {
	if o == nil {
		var ret []Aggregations
		return ret
	}

	return o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value
// and a boolean to check if the value has been set.
func (o *GenericDataQueryParams) GetAggregationsOk() ([]Aggregations, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aggregations, true
}

// SetAggregations sets field value
func (o *GenericDataQueryParams) SetAggregations(v []Aggregations) {
	o.Aggregations = v
}

// GetBuckets returns the Buckets field value
func (o *GenericDataQueryParams) GetBuckets() []BucketAggregation {
	if o == nil {
		var ret []BucketAggregation
		return ret
	}

	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value
// and a boolean to check if the value has been set.
func (o *GenericDataQueryParams) GetBucketsOk() ([]BucketAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Buckets, true
}

// SetBuckets sets field value
func (o *GenericDataQueryParams) SetBuckets(v []BucketAggregation) {
	o.Buckets = v
}

// GetLastEpochOnly returns the LastEpochOnly field value if set, zero value otherwise.
func (o *GenericDataQueryParams) GetLastEpochOnly() bool {
	if o == nil || IsNil(o.LastEpochOnly) {
		var ret bool
		return ret
	}
	return *o.LastEpochOnly
}

// GetLastEpochOnlyOk returns a tuple with the LastEpochOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericDataQueryParams) GetLastEpochOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.LastEpochOnly) {
		return nil, false
	}
	return o.LastEpochOnly, true
}

// HasLastEpochOnly returns a boolean if a field has been set.
func (o *GenericDataQueryParams) HasLastEpochOnly() bool {
	if o != nil && !IsNil(o.LastEpochOnly) {
		return true
	}

	return false
}

// SetLastEpochOnly gets a reference to the given bool and assigns it to the LastEpochOnly field.
func (o *GenericDataQueryParams) SetLastEpochOnly(v bool) {
	o.LastEpochOnly = &v
}

func (o GenericDataQueryParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericDataQueryParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["sessionRunIds"] = o.SessionRunIds
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.VerticalSplit) {
		toSerialize["verticalSplit"] = o.VerticalSplit
	}
	if !IsNil(o.HorizontalSplit) {
		toSerialize["horizontalSplit"] = o.HorizontalSplit
	}
	toSerialize["aggregations"] = o.Aggregations
	toSerialize["buckets"] = o.Buckets
	if !IsNil(o.LastEpochOnly) {
		toSerialize["lastEpochOnly"] = o.LastEpochOnly
	}
	return toSerialize, nil
}

type NullableGenericDataQueryParams struct {
	value *GenericDataQueryParams
	isSet bool
}

func (v NullableGenericDataQueryParams) Get() *GenericDataQueryParams {
	return v.value
}

func (v *NullableGenericDataQueryParams) Set(val *GenericDataQueryParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericDataQueryParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericDataQueryParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericDataQueryParams(val *GenericDataQueryParams) *NullableGenericDataQueryParams {
	return &NullableGenericDataQueryParams{value: val, isSet: true}
}

func (v NullableGenericDataQueryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericDataQueryParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
