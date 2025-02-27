# GenericDataQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SessionRunsToEpochs** | [**[]SessionRunToEpoch**](SessionRunToEpoch.md) |  | 
**ShowAllEpochs** | **bool** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**VerticalSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**HorizontalSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**InnerSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**Aggregations** | [**[]Aggregations**](Aggregations.md) |  | 
**Buckets** | [**[]SplitAgg**](SplitAgg.md) |  | 

## Methods

### NewGenericDataQueryParams

`func NewGenericDataQueryParams(projectId string, sessionRunsToEpochs []SessionRunToEpoch, showAllEpochs bool, aggregations []Aggregations, buckets []SplitAgg, ) *GenericDataQueryParams`

NewGenericDataQueryParams instantiates a new GenericDataQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericDataQueryParamsWithDefaults

`func NewGenericDataQueryParamsWithDefaults() *GenericDataQueryParams`

NewGenericDataQueryParamsWithDefaults instantiates a new GenericDataQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GenericDataQueryParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GenericDataQueryParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GenericDataQueryParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSessionRunsToEpochs

`func (o *GenericDataQueryParams) GetSessionRunsToEpochs() []SessionRunToEpoch`

GetSessionRunsToEpochs returns the SessionRunsToEpochs field if non-nil, zero value otherwise.

### GetSessionRunsToEpochsOk

`func (o *GenericDataQueryParams) GetSessionRunsToEpochsOk() (*[]SessionRunToEpoch, bool)`

GetSessionRunsToEpochsOk returns a tuple with the SessionRunsToEpochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunsToEpochs

`func (o *GenericDataQueryParams) SetSessionRunsToEpochs(v []SessionRunToEpoch)`

SetSessionRunsToEpochs sets SessionRunsToEpochs field to given value.


### GetShowAllEpochs

`func (o *GenericDataQueryParams) GetShowAllEpochs() bool`

GetShowAllEpochs returns the ShowAllEpochs field if non-nil, zero value otherwise.

### GetShowAllEpochsOk

`func (o *GenericDataQueryParams) GetShowAllEpochsOk() (*bool, bool)`

GetShowAllEpochsOk returns a tuple with the ShowAllEpochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAllEpochs

`func (o *GenericDataQueryParams) SetShowAllEpochs(v bool)`

SetShowAllEpochs sets ShowAllEpochs field to given value.


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

`func (o *GenericDataQueryParams) GetVerticalSplit() SplitAgg`

GetVerticalSplit returns the VerticalSplit field if non-nil, zero value otherwise.

### GetVerticalSplitOk

`func (o *GenericDataQueryParams) GetVerticalSplitOk() (*SplitAgg, bool)`

GetVerticalSplitOk returns a tuple with the VerticalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalSplit

`func (o *GenericDataQueryParams) SetVerticalSplit(v SplitAgg)`

SetVerticalSplit sets VerticalSplit field to given value.

### HasVerticalSplit

`func (o *GenericDataQueryParams) HasVerticalSplit() bool`

HasVerticalSplit returns a boolean if a field has been set.

### GetHorizontalSplit

`func (o *GenericDataQueryParams) GetHorizontalSplit() SplitAgg`

GetHorizontalSplit returns the HorizontalSplit field if non-nil, zero value otherwise.

### GetHorizontalSplitOk

`func (o *GenericDataQueryParams) GetHorizontalSplitOk() (*SplitAgg, bool)`

GetHorizontalSplitOk returns a tuple with the HorizontalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalSplit

`func (o *GenericDataQueryParams) SetHorizontalSplit(v SplitAgg)`

SetHorizontalSplit sets HorizontalSplit field to given value.

### HasHorizontalSplit

`func (o *GenericDataQueryParams) HasHorizontalSplit() bool`

HasHorizontalSplit returns a boolean if a field has been set.

### GetInnerSplit

`func (o *GenericDataQueryParams) GetInnerSplit() SplitAgg`

GetInnerSplit returns the InnerSplit field if non-nil, zero value otherwise.

### GetInnerSplitOk

`func (o *GenericDataQueryParams) GetInnerSplitOk() (*SplitAgg, bool)`

GetInnerSplitOk returns a tuple with the InnerSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerSplit

`func (o *GenericDataQueryParams) SetInnerSplit(v SplitAgg)`

SetInnerSplit sets InnerSplit field to given value.

### HasInnerSplit

`func (o *GenericDataQueryParams) HasInnerSplit() bool`

HasInnerSplit returns a boolean if a field has been set.

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

`func (o *GenericDataQueryParams) GetBuckets() []SplitAgg`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *GenericDataQueryParams) GetBucketsOk() (*[]SplitAgg, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *GenericDataQueryParams) SetBuckets(v []SplitAgg)`

SetBuckets sets Buckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


