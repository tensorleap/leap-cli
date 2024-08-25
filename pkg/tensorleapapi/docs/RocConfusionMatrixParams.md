# RocConfusionMatrixParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionRunIds** | **[]string** |  | 
**ProjectId** | **string** |  | 
**CustomMetricName** | **string** |  | 
**VerticalSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**HorizontalSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**AbsAxis** | Pointer to **bool** |  | [optional] 

## Methods

### NewRocConfusionMatrixParams

`func NewRocConfusionMatrixParams(sessionRunIds []string, projectId string, customMetricName string, ) *RocConfusionMatrixParams`

NewRocConfusionMatrixParams instantiates a new RocConfusionMatrixParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRocConfusionMatrixParamsWithDefaults

`func NewRocConfusionMatrixParamsWithDefaults() *RocConfusionMatrixParams`

NewRocConfusionMatrixParamsWithDefaults instantiates a new RocConfusionMatrixParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionRunIds

`func (o *RocConfusionMatrixParams) GetSessionRunIds() []string`

GetSessionRunIds returns the SessionRunIds field if non-nil, zero value otherwise.

### GetSessionRunIdsOk

`func (o *RocConfusionMatrixParams) GetSessionRunIdsOk() (*[]string, bool)`

GetSessionRunIdsOk returns a tuple with the SessionRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunIds

`func (o *RocConfusionMatrixParams) SetSessionRunIds(v []string)`

SetSessionRunIds sets SessionRunIds field to given value.


### GetProjectId

`func (o *RocConfusionMatrixParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RocConfusionMatrixParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RocConfusionMatrixParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCustomMetricName

`func (o *RocConfusionMatrixParams) GetCustomMetricName() string`

GetCustomMetricName returns the CustomMetricName field if non-nil, zero value otherwise.

### GetCustomMetricNameOk

`func (o *RocConfusionMatrixParams) GetCustomMetricNameOk() (*string, bool)`

GetCustomMetricNameOk returns a tuple with the CustomMetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetricName

`func (o *RocConfusionMatrixParams) SetCustomMetricName(v string)`

SetCustomMetricName sets CustomMetricName field to given value.


### GetVerticalSplit

`func (o *RocConfusionMatrixParams) GetVerticalSplit() SplitAgg`

GetVerticalSplit returns the VerticalSplit field if non-nil, zero value otherwise.

### GetVerticalSplitOk

`func (o *RocConfusionMatrixParams) GetVerticalSplitOk() (*SplitAgg, bool)`

GetVerticalSplitOk returns a tuple with the VerticalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalSplit

`func (o *RocConfusionMatrixParams) SetVerticalSplit(v SplitAgg)`

SetVerticalSplit sets VerticalSplit field to given value.

### HasVerticalSplit

`func (o *RocConfusionMatrixParams) HasVerticalSplit() bool`

HasVerticalSplit returns a boolean if a field has been set.

### GetHorizontalSplit

`func (o *RocConfusionMatrixParams) GetHorizontalSplit() SplitAgg`

GetHorizontalSplit returns the HorizontalSplit field if non-nil, zero value otherwise.

### GetHorizontalSplitOk

`func (o *RocConfusionMatrixParams) GetHorizontalSplitOk() (*SplitAgg, bool)`

GetHorizontalSplitOk returns a tuple with the HorizontalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalSplit

`func (o *RocConfusionMatrixParams) SetHorizontalSplit(v SplitAgg)`

SetHorizontalSplit sets HorizontalSplit field to given value.

### HasHorizontalSplit

`func (o *RocConfusionMatrixParams) HasHorizontalSplit() bool`

HasHorizontalSplit returns a boolean if a field has been set.

### GetFilters

`func (o *RocConfusionMatrixParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *RocConfusionMatrixParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *RocConfusionMatrixParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *RocConfusionMatrixParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetAbsAxis

`func (o *RocConfusionMatrixParams) GetAbsAxis() bool`

GetAbsAxis returns the AbsAxis field if non-nil, zero value otherwise.

### GetAbsAxisOk

`func (o *RocConfusionMatrixParams) GetAbsAxisOk() (*bool, bool)`

GetAbsAxisOk returns a tuple with the AbsAxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsAxis

`func (o *RocConfusionMatrixParams) SetAbsAxis(v bool)`

SetAbsAxis sets AbsAxis field to given value.

### HasAbsAxis

`func (o *RocConfusionMatrixParams) HasAbsAxis() bool`

HasAbsAxis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


