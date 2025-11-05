# MultiThresholdConfusionMatrixParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionRunsToEpochs** | [**[]SessionRunToEpoch**](SessionRunToEpoch.md) |  | 
**ProjectId** | **string** |  | 
**CustomMetricName** | **string** |  | 
**VerticalSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**HorizontalSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**ElementInstance** | Pointer to **bool** |  | [optional] 

## Methods

### NewMultiThresholdConfusionMatrixParams

`func NewMultiThresholdConfusionMatrixParams(sessionRunsToEpochs []SessionRunToEpoch, projectId string, customMetricName string, ) *MultiThresholdConfusionMatrixParams`

NewMultiThresholdConfusionMatrixParams instantiates a new MultiThresholdConfusionMatrixParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiThresholdConfusionMatrixParamsWithDefaults

`func NewMultiThresholdConfusionMatrixParamsWithDefaults() *MultiThresholdConfusionMatrixParams`

NewMultiThresholdConfusionMatrixParamsWithDefaults instantiates a new MultiThresholdConfusionMatrixParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionRunsToEpochs

`func (o *MultiThresholdConfusionMatrixParams) GetSessionRunsToEpochs() []SessionRunToEpoch`

GetSessionRunsToEpochs returns the SessionRunsToEpochs field if non-nil, zero value otherwise.

### GetSessionRunsToEpochsOk

`func (o *MultiThresholdConfusionMatrixParams) GetSessionRunsToEpochsOk() (*[]SessionRunToEpoch, bool)`

GetSessionRunsToEpochsOk returns a tuple with the SessionRunsToEpochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunsToEpochs

`func (o *MultiThresholdConfusionMatrixParams) SetSessionRunsToEpochs(v []SessionRunToEpoch)`

SetSessionRunsToEpochs sets SessionRunsToEpochs field to given value.


### GetProjectId

`func (o *MultiThresholdConfusionMatrixParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MultiThresholdConfusionMatrixParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MultiThresholdConfusionMatrixParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCustomMetricName

`func (o *MultiThresholdConfusionMatrixParams) GetCustomMetricName() string`

GetCustomMetricName returns the CustomMetricName field if non-nil, zero value otherwise.

### GetCustomMetricNameOk

`func (o *MultiThresholdConfusionMatrixParams) GetCustomMetricNameOk() (*string, bool)`

GetCustomMetricNameOk returns a tuple with the CustomMetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetricName

`func (o *MultiThresholdConfusionMatrixParams) SetCustomMetricName(v string)`

SetCustomMetricName sets CustomMetricName field to given value.


### GetVerticalSplit

`func (o *MultiThresholdConfusionMatrixParams) GetVerticalSplit() SplitAgg`

GetVerticalSplit returns the VerticalSplit field if non-nil, zero value otherwise.

### GetVerticalSplitOk

`func (o *MultiThresholdConfusionMatrixParams) GetVerticalSplitOk() (*SplitAgg, bool)`

GetVerticalSplitOk returns a tuple with the VerticalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalSplit

`func (o *MultiThresholdConfusionMatrixParams) SetVerticalSplit(v SplitAgg)`

SetVerticalSplit sets VerticalSplit field to given value.

### HasVerticalSplit

`func (o *MultiThresholdConfusionMatrixParams) HasVerticalSplit() bool`

HasVerticalSplit returns a boolean if a field has been set.

### GetHorizontalSplit

`func (o *MultiThresholdConfusionMatrixParams) GetHorizontalSplit() SplitAgg`

GetHorizontalSplit returns the HorizontalSplit field if non-nil, zero value otherwise.

### GetHorizontalSplitOk

`func (o *MultiThresholdConfusionMatrixParams) GetHorizontalSplitOk() (*SplitAgg, bool)`

GetHorizontalSplitOk returns a tuple with the HorizontalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalSplit

`func (o *MultiThresholdConfusionMatrixParams) SetHorizontalSplit(v SplitAgg)`

SetHorizontalSplit sets HorizontalSplit field to given value.

### HasHorizontalSplit

`func (o *MultiThresholdConfusionMatrixParams) HasHorizontalSplit() bool`

HasHorizontalSplit returns a boolean if a field has been set.

### GetFilters

`func (o *MultiThresholdConfusionMatrixParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MultiThresholdConfusionMatrixParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MultiThresholdConfusionMatrixParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MultiThresholdConfusionMatrixParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetElementInstance

`func (o *MultiThresholdConfusionMatrixParams) GetElementInstance() bool`

GetElementInstance returns the ElementInstance field if non-nil, zero value otherwise.

### GetElementInstanceOk

`func (o *MultiThresholdConfusionMatrixParams) GetElementInstanceOk() (*bool, bool)`

GetElementInstanceOk returns a tuple with the ElementInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementInstance

`func (o *MultiThresholdConfusionMatrixParams) SetElementInstance(v bool)`

SetElementInstance sets ElementInstance field to given value.

### HasElementInstance

`func (o *MultiThresholdConfusionMatrixParams) HasElementInstance() bool`

HasElementInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


