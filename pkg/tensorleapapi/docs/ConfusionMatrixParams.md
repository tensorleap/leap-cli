# ConfusionMatrixParams

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

## Methods

### NewConfusionMatrixParams

`func NewConfusionMatrixParams(projectId string, sessionRunsToEpochs []SessionRunToEpoch, x SplitAgg, customMetricName string, ) *ConfusionMatrixParams`

NewConfusionMatrixParams instantiates a new ConfusionMatrixParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfusionMatrixParamsWithDefaults

`func NewConfusionMatrixParamsWithDefaults() *ConfusionMatrixParams`

NewConfusionMatrixParamsWithDefaults instantiates a new ConfusionMatrixParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ConfusionMatrixParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ConfusionMatrixParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ConfusionMatrixParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSessionRunsToEpochs

`func (o *ConfusionMatrixParams) GetSessionRunsToEpochs() []SessionRunToEpoch`

GetSessionRunsToEpochs returns the SessionRunsToEpochs field if non-nil, zero value otherwise.

### GetSessionRunsToEpochsOk

`func (o *ConfusionMatrixParams) GetSessionRunsToEpochsOk() (*[]SessionRunToEpoch, bool)`

GetSessionRunsToEpochsOk returns a tuple with the SessionRunsToEpochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunsToEpochs

`func (o *ConfusionMatrixParams) SetSessionRunsToEpochs(v []SessionRunToEpoch)`

SetSessionRunsToEpochs sets SessionRunsToEpochs field to given value.


### GetX

`func (o *ConfusionMatrixParams) GetX() SplitAgg`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ConfusionMatrixParams) GetXOk() (*SplitAgg, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ConfusionMatrixParams) SetX(v SplitAgg)`

SetX sets X field to given value.


### GetVerticalSplit

`func (o *ConfusionMatrixParams) GetVerticalSplit() SplitAgg`

GetVerticalSplit returns the VerticalSplit field if non-nil, zero value otherwise.

### GetVerticalSplitOk

`func (o *ConfusionMatrixParams) GetVerticalSplitOk() (*SplitAgg, bool)`

GetVerticalSplitOk returns a tuple with the VerticalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalSplit

`func (o *ConfusionMatrixParams) SetVerticalSplit(v SplitAgg)`

SetVerticalSplit sets VerticalSplit field to given value.

### HasVerticalSplit

`func (o *ConfusionMatrixParams) HasVerticalSplit() bool`

HasVerticalSplit returns a boolean if a field has been set.

### GetHorizontalSplit

`func (o *ConfusionMatrixParams) GetHorizontalSplit() SplitAgg`

GetHorizontalSplit returns the HorizontalSplit field if non-nil, zero value otherwise.

### GetHorizontalSplitOk

`func (o *ConfusionMatrixParams) GetHorizontalSplitOk() (*SplitAgg, bool)`

GetHorizontalSplitOk returns a tuple with the HorizontalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalSplit

`func (o *ConfusionMatrixParams) SetHorizontalSplit(v SplitAgg)`

SetHorizontalSplit sets HorizontalSplit field to given value.

### HasHorizontalSplit

`func (o *ConfusionMatrixParams) HasHorizontalSplit() bool`

HasHorizontalSplit returns a boolean if a field has been set.

### GetInnerSplit

`func (o *ConfusionMatrixParams) GetInnerSplit() SplitAgg`

GetInnerSplit returns the InnerSplit field if non-nil, zero value otherwise.

### GetInnerSplitOk

`func (o *ConfusionMatrixParams) GetInnerSplitOk() (*SplitAgg, bool)`

GetInnerSplitOk returns a tuple with the InnerSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerSplit

`func (o *ConfusionMatrixParams) SetInnerSplit(v SplitAgg)`

SetInnerSplit sets InnerSplit field to given value.

### HasInnerSplit

`func (o *ConfusionMatrixParams) HasInnerSplit() bool`

HasInnerSplit returns a boolean if a field has been set.

### GetThreshold

`func (o *ConfusionMatrixParams) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ConfusionMatrixParams) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ConfusionMatrixParams) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ConfusionMatrixParams) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetCustomMetricName

`func (o *ConfusionMatrixParams) GetCustomMetricName() string`

GetCustomMetricName returns the CustomMetricName field if non-nil, zero value otherwise.

### GetCustomMetricNameOk

`func (o *ConfusionMatrixParams) GetCustomMetricNameOk() (*string, bool)`

GetCustomMetricNameOk returns a tuple with the CustomMetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetricName

`func (o *ConfusionMatrixParams) SetCustomMetricName(v string)`

SetCustomMetricName sets CustomMetricName field to given value.


### GetFilters

`func (o *ConfusionMatrixParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ConfusionMatrixParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ConfusionMatrixParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ConfusionMatrixParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


