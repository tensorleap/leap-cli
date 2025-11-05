# ConfusionMatrixTableParams

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
**SplitByLabelOrder** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**SplitByLabel** | **bool** |  | 
**FilterLabels** | Pointer to **[]string** |  | [optional] 
**Threshold** | Pointer to **float64** |  | [optional] 

## Methods

### NewConfusionMatrixTableParams

`func NewConfusionMatrixTableParams(sessionRunsToEpochs []SessionRunToEpoch, projectId string, customMetricName string, splitByLabel bool, ) *ConfusionMatrixTableParams`

NewConfusionMatrixTableParams instantiates a new ConfusionMatrixTableParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfusionMatrixTableParamsWithDefaults

`func NewConfusionMatrixTableParamsWithDefaults() *ConfusionMatrixTableParams`

NewConfusionMatrixTableParamsWithDefaults instantiates a new ConfusionMatrixTableParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionRunsToEpochs

`func (o *ConfusionMatrixTableParams) GetSessionRunsToEpochs() []SessionRunToEpoch`

GetSessionRunsToEpochs returns the SessionRunsToEpochs field if non-nil, zero value otherwise.

### GetSessionRunsToEpochsOk

`func (o *ConfusionMatrixTableParams) GetSessionRunsToEpochsOk() (*[]SessionRunToEpoch, bool)`

GetSessionRunsToEpochsOk returns a tuple with the SessionRunsToEpochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunsToEpochs

`func (o *ConfusionMatrixTableParams) SetSessionRunsToEpochs(v []SessionRunToEpoch)`

SetSessionRunsToEpochs sets SessionRunsToEpochs field to given value.


### GetProjectId

`func (o *ConfusionMatrixTableParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ConfusionMatrixTableParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ConfusionMatrixTableParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCustomMetricName

`func (o *ConfusionMatrixTableParams) GetCustomMetricName() string`

GetCustomMetricName returns the CustomMetricName field if non-nil, zero value otherwise.

### GetCustomMetricNameOk

`func (o *ConfusionMatrixTableParams) GetCustomMetricNameOk() (*string, bool)`

GetCustomMetricNameOk returns a tuple with the CustomMetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetricName

`func (o *ConfusionMatrixTableParams) SetCustomMetricName(v string)`

SetCustomMetricName sets CustomMetricName field to given value.


### GetVerticalSplit

`func (o *ConfusionMatrixTableParams) GetVerticalSplit() SplitAgg`

GetVerticalSplit returns the VerticalSplit field if non-nil, zero value otherwise.

### GetVerticalSplitOk

`func (o *ConfusionMatrixTableParams) GetVerticalSplitOk() (*SplitAgg, bool)`

GetVerticalSplitOk returns a tuple with the VerticalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalSplit

`func (o *ConfusionMatrixTableParams) SetVerticalSplit(v SplitAgg)`

SetVerticalSplit sets VerticalSplit field to given value.

### HasVerticalSplit

`func (o *ConfusionMatrixTableParams) HasVerticalSplit() bool`

HasVerticalSplit returns a boolean if a field has been set.

### GetHorizontalSplit

`func (o *ConfusionMatrixTableParams) GetHorizontalSplit() SplitAgg`

GetHorizontalSplit returns the HorizontalSplit field if non-nil, zero value otherwise.

### GetHorizontalSplitOk

`func (o *ConfusionMatrixTableParams) GetHorizontalSplitOk() (*SplitAgg, bool)`

GetHorizontalSplitOk returns a tuple with the HorizontalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalSplit

`func (o *ConfusionMatrixTableParams) SetHorizontalSplit(v SplitAgg)`

SetHorizontalSplit sets HorizontalSplit field to given value.

### HasHorizontalSplit

`func (o *ConfusionMatrixTableParams) HasHorizontalSplit() bool`

HasHorizontalSplit returns a boolean if a field has been set.

### GetFilters

`func (o *ConfusionMatrixTableParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ConfusionMatrixTableParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ConfusionMatrixTableParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ConfusionMatrixTableParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetElementInstance

`func (o *ConfusionMatrixTableParams) GetElementInstance() bool`

GetElementInstance returns the ElementInstance field if non-nil, zero value otherwise.

### GetElementInstanceOk

`func (o *ConfusionMatrixTableParams) GetElementInstanceOk() (*bool, bool)`

GetElementInstanceOk returns a tuple with the ElementInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementInstance

`func (o *ConfusionMatrixTableParams) SetElementInstance(v bool)`

SetElementInstance sets ElementInstance field to given value.

### HasElementInstance

`func (o *ConfusionMatrixTableParams) HasElementInstance() bool`

HasElementInstance returns a boolean if a field has been set.

### GetSplitByLabelOrder

`func (o *ConfusionMatrixTableParams) GetSplitByLabelOrder() OrderType`

GetSplitByLabelOrder returns the SplitByLabelOrder field if non-nil, zero value otherwise.

### GetSplitByLabelOrderOk

`func (o *ConfusionMatrixTableParams) GetSplitByLabelOrderOk() (*OrderType, bool)`

GetSplitByLabelOrderOk returns a tuple with the SplitByLabelOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitByLabelOrder

`func (o *ConfusionMatrixTableParams) SetSplitByLabelOrder(v OrderType)`

SetSplitByLabelOrder sets SplitByLabelOrder field to given value.

### HasSplitByLabelOrder

`func (o *ConfusionMatrixTableParams) HasSplitByLabelOrder() bool`

HasSplitByLabelOrder returns a boolean if a field has been set.

### GetSplitByLabel

`func (o *ConfusionMatrixTableParams) GetSplitByLabel() bool`

GetSplitByLabel returns the SplitByLabel field if non-nil, zero value otherwise.

### GetSplitByLabelOk

`func (o *ConfusionMatrixTableParams) GetSplitByLabelOk() (*bool, bool)`

GetSplitByLabelOk returns a tuple with the SplitByLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitByLabel

`func (o *ConfusionMatrixTableParams) SetSplitByLabel(v bool)`

SetSplitByLabel sets SplitByLabel field to given value.


### GetFilterLabels

`func (o *ConfusionMatrixTableParams) GetFilterLabels() []string`

GetFilterLabels returns the FilterLabels field if non-nil, zero value otherwise.

### GetFilterLabelsOk

`func (o *ConfusionMatrixTableParams) GetFilterLabelsOk() (*[]string, bool)`

GetFilterLabelsOk returns a tuple with the FilterLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLabels

`func (o *ConfusionMatrixTableParams) SetFilterLabels(v []string)`

SetFilterLabels sets FilterLabels field to given value.

### HasFilterLabels

`func (o *ConfusionMatrixTableParams) HasFilterLabels() bool`

HasFilterLabels returns a boolean if a field has been set.

### GetThreshold

`func (o *ConfusionMatrixTableParams) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ConfusionMatrixTableParams) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ConfusionMatrixTableParams) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ConfusionMatrixTableParams) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


