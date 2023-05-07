# GenericDataQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionRunIds** | **[]string** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**VerticalSplit** | Pointer to **string** |  | [optional] 
**HorizontalSplit** | Pointer to **string** |  | [optional] 
**Aggregations** | [**[]Aggregations**](Aggregations.md) |  | 
**Buckets** | [**[]BucketAggregation**](BucketAggregation.md) |  | 
**LastEpochOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewGenericDataQueryParams

`func NewGenericDataQueryParams(sessionRunIds []string, aggregations []Aggregations, buckets []BucketAggregation, ) *GenericDataQueryParams`

NewGenericDataQueryParams instantiates a new GenericDataQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericDataQueryParamsWithDefaults

`func NewGenericDataQueryParamsWithDefaults() *GenericDataQueryParams`

NewGenericDataQueryParamsWithDefaults instantiates a new GenericDataQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionRunIds

`func (o *GenericDataQueryParams) GetSessionRunIds() []string`

GetSessionRunIds returns the SessionRunIds field if non-nil, zero value otherwise.

### GetSessionRunIdsOk

`func (o *GenericDataQueryParams) GetSessionRunIdsOk() (*[]string, bool)`

GetSessionRunIdsOk returns a tuple with the SessionRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunIds

`func (o *GenericDataQueryParams) SetSessionRunIds(v []string)`

SetSessionRunIds sets SessionRunIds field to given value.


### GetFilters

`func (o *GenericDataQueryParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GenericDataQueryParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GenericDataQueryParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GenericDataQueryParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetVerticalSplit

`func (o *GenericDataQueryParams) GetVerticalSplit() string`

GetVerticalSplit returns the VerticalSplit field if non-nil, zero value otherwise.

### GetVerticalSplitOk

`func (o *GenericDataQueryParams) GetVerticalSplitOk() (*string, bool)`

GetVerticalSplitOk returns a tuple with the VerticalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalSplit

`func (o *GenericDataQueryParams) SetVerticalSplit(v string)`

SetVerticalSplit sets VerticalSplit field to given value.

### HasVerticalSplit

`func (o *GenericDataQueryParams) HasVerticalSplit() bool`

HasVerticalSplit returns a boolean if a field has been set.

### GetHorizontalSplit

`func (o *GenericDataQueryParams) GetHorizontalSplit() string`

GetHorizontalSplit returns the HorizontalSplit field if non-nil, zero value otherwise.

### GetHorizontalSplitOk

`func (o *GenericDataQueryParams) GetHorizontalSplitOk() (*string, bool)`

GetHorizontalSplitOk returns a tuple with the HorizontalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalSplit

`func (o *GenericDataQueryParams) SetHorizontalSplit(v string)`

SetHorizontalSplit sets HorizontalSplit field to given value.

### HasHorizontalSplit

`func (o *GenericDataQueryParams) HasHorizontalSplit() bool`

HasHorizontalSplit returns a boolean if a field has been set.

### GetAggregations

`func (o *GenericDataQueryParams) GetAggregations() []Aggregations`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *GenericDataQueryParams) GetAggregationsOk() (*[]Aggregations, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *GenericDataQueryParams) SetAggregations(v []Aggregations)`

SetAggregations sets Aggregations field to given value.


### GetBuckets

`func (o *GenericDataQueryParams) GetBuckets() []BucketAggregation`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *GenericDataQueryParams) GetBucketsOk() (*[]BucketAggregation, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *GenericDataQueryParams) SetBuckets(v []BucketAggregation)`

SetBuckets sets Buckets field to given value.


### GetLastEpochOnly

`func (o *GenericDataQueryParams) GetLastEpochOnly() bool`

GetLastEpochOnly returns the LastEpochOnly field if non-nil, zero value otherwise.

### GetLastEpochOnlyOk

`func (o *GenericDataQueryParams) GetLastEpochOnlyOk() (*bool, bool)`

GetLastEpochOnlyOk returns a tuple with the LastEpochOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEpochOnly

`func (o *GenericDataQueryParams) SetLastEpochOnly(v bool)`

SetLastEpochOnly sets LastEpochOnly field to given value.

### HasLastEpochOnly

`func (o *GenericDataQueryParams) HasLastEpochOnly() bool`

HasLastEpochOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


