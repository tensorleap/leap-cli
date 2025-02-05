# GetConfusionMatrixResultCombinationsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SessionRunsToEpochs** | [**[]SessionRunToEpoch**](SessionRunToEpoch.md) |  | 
**X** | [**SplitAgg**](SplitAgg.md) |  | 
**VerticalSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**HorizontalSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**InnerSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**Threshold** | Pointer to **float64** |  | [optional] 
**CustomMetricName** | **string** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**FilterLabels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetConfusionMatrixResultCombinationsParams

`func NewGetConfusionMatrixResultCombinationsParams(projectId string, sessionRunsToEpochs []SessionRunToEpoch, x SplitAgg, customMetricName string, ) *GetConfusionMatrixResultCombinationsParams`

NewGetConfusionMatrixResultCombinationsParams instantiates a new GetConfusionMatrixResultCombinationsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConfusionMatrixResultCombinationsParamsWithDefaults

`func NewGetConfusionMatrixResultCombinationsParamsWithDefaults() *GetConfusionMatrixResultCombinationsParams`

NewGetConfusionMatrixResultCombinationsParamsWithDefaults instantiates a new GetConfusionMatrixResultCombinationsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GetConfusionMatrixResultCombinationsParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetConfusionMatrixResultCombinationsParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetConfusionMatrixResultCombinationsParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSessionRunsToEpochs

`func (o *GetConfusionMatrixResultCombinationsParams) GetSessionRunsToEpochs() []SessionRunToEpoch`

GetSessionRunsToEpochs returns the SessionRunsToEpochs field if non-nil, zero value otherwise.

### GetSessionRunsToEpochsOk

`func (o *GetConfusionMatrixResultCombinationsParams) GetSessionRunsToEpochsOk() (*[]SessionRunToEpoch, bool)`

GetSessionRunsToEpochsOk returns a tuple with the SessionRunsToEpochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunsToEpochs

`func (o *GetConfusionMatrixResultCombinationsParams) SetSessionRunsToEpochs(v []SessionRunToEpoch)`

SetSessionRunsToEpochs sets SessionRunsToEpochs field to given value.


### GetX

`func (o *GetConfusionMatrixResultCombinationsParams) GetX() SplitAgg`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *GetConfusionMatrixResultCombinationsParams) GetXOk() (*SplitAgg, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *GetConfusionMatrixResultCombinationsParams) SetX(v SplitAgg)`

SetX sets X field to given value.


### GetVerticalSplit

`func (o *GetConfusionMatrixResultCombinationsParams) GetVerticalSplit() SplitAgg`

GetVerticalSplit returns the VerticalSplit field if non-nil, zero value otherwise.

### GetVerticalSplitOk

`func (o *GetConfusionMatrixResultCombinationsParams) GetVerticalSplitOk() (*SplitAgg, bool)`

GetVerticalSplitOk returns a tuple with the VerticalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalSplit

`func (o *GetConfusionMatrixResultCombinationsParams) SetVerticalSplit(v SplitAgg)`

SetVerticalSplit sets VerticalSplit field to given value.

### HasVerticalSplit

`func (o *GetConfusionMatrixResultCombinationsParams) HasVerticalSplit() bool`

HasVerticalSplit returns a boolean if a field has been set.

### GetHorizontalSplit

`func (o *GetConfusionMatrixResultCombinationsParams) GetHorizontalSplit() SplitAgg`

GetHorizontalSplit returns the HorizontalSplit field if non-nil, zero value otherwise.

### GetHorizontalSplitOk

`func (o *GetConfusionMatrixResultCombinationsParams) GetHorizontalSplitOk() (*SplitAgg, bool)`

GetHorizontalSplitOk returns a tuple with the HorizontalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalSplit

`func (o *GetConfusionMatrixResultCombinationsParams) SetHorizontalSplit(v SplitAgg)`

SetHorizontalSplit sets HorizontalSplit field to given value.

### HasHorizontalSplit

`func (o *GetConfusionMatrixResultCombinationsParams) HasHorizontalSplit() bool`

HasHorizontalSplit returns a boolean if a field has been set.

### GetInnerSplit

`func (o *GetConfusionMatrixResultCombinationsParams) GetInnerSplit() SplitAgg`

GetInnerSplit returns the InnerSplit field if non-nil, zero value otherwise.

### GetInnerSplitOk

`func (o *GetConfusionMatrixResultCombinationsParams) GetInnerSplitOk() (*SplitAgg, bool)`

GetInnerSplitOk returns a tuple with the InnerSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerSplit

`func (o *GetConfusionMatrixResultCombinationsParams) SetInnerSplit(v SplitAgg)`

SetInnerSplit sets InnerSplit field to given value.

### HasInnerSplit

`func (o *GetConfusionMatrixResultCombinationsParams) HasInnerSplit() bool`

HasInnerSplit returns a boolean if a field has been set.

### GetThreshold

`func (o *GetConfusionMatrixResultCombinationsParams) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *GetConfusionMatrixResultCombinationsParams) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *GetConfusionMatrixResultCombinationsParams) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *GetConfusionMatrixResultCombinationsParams) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetCustomMetricName

`func (o *GetConfusionMatrixResultCombinationsParams) GetCustomMetricName() string`

GetCustomMetricName returns the CustomMetricName field if non-nil, zero value otherwise.

### GetCustomMetricNameOk

`func (o *GetConfusionMatrixResultCombinationsParams) GetCustomMetricNameOk() (*string, bool)`

GetCustomMetricNameOk returns a tuple with the CustomMetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetricName

`func (o *GetConfusionMatrixResultCombinationsParams) SetCustomMetricName(v string)`

SetCustomMetricName sets CustomMetricName field to given value.


### GetFilters

`func (o *GetConfusionMatrixResultCombinationsParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GetConfusionMatrixResultCombinationsParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GetConfusionMatrixResultCombinationsParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GetConfusionMatrixResultCombinationsParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFilterLabels

`func (o *GetConfusionMatrixResultCombinationsParams) GetFilterLabels() []string`

GetFilterLabels returns the FilterLabels field if non-nil, zero value otherwise.

### GetFilterLabelsOk

`func (o *GetConfusionMatrixResultCombinationsParams) GetFilterLabelsOk() (*[]string, bool)`

GetFilterLabelsOk returns a tuple with the FilterLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLabels

`func (o *GetConfusionMatrixResultCombinationsParams) SetFilterLabels(v []string)`

SetFilterLabels sets FilterLabels field to given value.

### HasFilterLabels

`func (o *GetConfusionMatrixResultCombinationsParams) HasFilterLabels() bool`

HasFilterLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


