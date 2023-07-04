# ConfusionMatrixParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**XField** | **string** |  | 
**SessionRunIds** | **[]string** |  | 
**VerticalSplit** | Pointer to **string** |  | [optional] 
**HorizontalSplit** | Pointer to **string** |  | [optional] 
**Threshold** | Pointer to **float64** |  | [optional] 
**CustomMetricName** | **string** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**XFieldAggregationType** | Pointer to [**EsBatchAggregationType**](EsBatchAggregationType.md) |  | [optional] 
**DataDistributionType** | Pointer to [**DataDistributionType**](DataDistributionType.md) |  | [optional] 
**OrderByParam** | Pointer to **string** |  | [optional] 
**OrderParams** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**XAxisSizeInterval** | **float64** |  | 
**LastEpochOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewConfusionMatrixParams

`func NewConfusionMatrixParams(projectId string, xField string, sessionRunIds []string, customMetricName string, xAxisSizeInterval float64, ) *ConfusionMatrixParams`

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


### GetXField

`func (o *ConfusionMatrixParams) GetXField() string`

GetXField returns the XField field if non-nil, zero value otherwise.

### GetXFieldOk

`func (o *ConfusionMatrixParams) GetXFieldOk() (*string, bool)`

GetXFieldOk returns a tuple with the XField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXField

`func (o *ConfusionMatrixParams) SetXField(v string)`

SetXField sets XField field to given value.


### GetSessionRunIds

`func (o *ConfusionMatrixParams) GetSessionRunIds() []string`

GetSessionRunIds returns the SessionRunIds field if non-nil, zero value otherwise.

### GetSessionRunIdsOk

`func (o *ConfusionMatrixParams) GetSessionRunIdsOk() (*[]string, bool)`

GetSessionRunIdsOk returns a tuple with the SessionRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunIds

`func (o *ConfusionMatrixParams) SetSessionRunIds(v []string)`

SetSessionRunIds sets SessionRunIds field to given value.


### GetVerticalSplit

`func (o *ConfusionMatrixParams) GetVerticalSplit() string`

GetVerticalSplit returns the VerticalSplit field if non-nil, zero value otherwise.

### GetVerticalSplitOk

`func (o *ConfusionMatrixParams) GetVerticalSplitOk() (*string, bool)`

GetVerticalSplitOk returns a tuple with the VerticalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalSplit

`func (o *ConfusionMatrixParams) SetVerticalSplit(v string)`

SetVerticalSplit sets VerticalSplit field to given value.

### HasVerticalSplit

`func (o *ConfusionMatrixParams) HasVerticalSplit() bool`

HasVerticalSplit returns a boolean if a field has been set.

### GetHorizontalSplit

`func (o *ConfusionMatrixParams) GetHorizontalSplit() string`

GetHorizontalSplit returns the HorizontalSplit field if non-nil, zero value otherwise.

### GetHorizontalSplitOk

`func (o *ConfusionMatrixParams) GetHorizontalSplitOk() (*string, bool)`

GetHorizontalSplitOk returns a tuple with the HorizontalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalSplit

`func (o *ConfusionMatrixParams) SetHorizontalSplit(v string)`

SetHorizontalSplit sets HorizontalSplit field to given value.

### HasHorizontalSplit

`func (o *ConfusionMatrixParams) HasHorizontalSplit() bool`

HasHorizontalSplit returns a boolean if a field has been set.

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

### GetXFieldAggregationType

`func (o *ConfusionMatrixParams) GetXFieldAggregationType() EsBatchAggregationType`

GetXFieldAggregationType returns the XFieldAggregationType field if non-nil, zero value otherwise.

### GetXFieldAggregationTypeOk

`func (o *ConfusionMatrixParams) GetXFieldAggregationTypeOk() (*EsBatchAggregationType, bool)`

GetXFieldAggregationTypeOk returns a tuple with the XFieldAggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXFieldAggregationType

`func (o *ConfusionMatrixParams) SetXFieldAggregationType(v EsBatchAggregationType)`

SetXFieldAggregationType sets XFieldAggregationType field to given value.

### HasXFieldAggregationType

`func (o *ConfusionMatrixParams) HasXFieldAggregationType() bool`

HasXFieldAggregationType returns a boolean if a field has been set.

### GetDataDistributionType

`func (o *ConfusionMatrixParams) GetDataDistributionType() DataDistributionType`

GetDataDistributionType returns the DataDistributionType field if non-nil, zero value otherwise.

### GetDataDistributionTypeOk

`func (o *ConfusionMatrixParams) GetDataDistributionTypeOk() (*DataDistributionType, bool)`

GetDataDistributionTypeOk returns a tuple with the DataDistributionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDistributionType

`func (o *ConfusionMatrixParams) SetDataDistributionType(v DataDistributionType)`

SetDataDistributionType sets DataDistributionType field to given value.

### HasDataDistributionType

`func (o *ConfusionMatrixParams) HasDataDistributionType() bool`

HasDataDistributionType returns a boolean if a field has been set.

### GetOrderByParam

`func (o *ConfusionMatrixParams) GetOrderByParam() string`

GetOrderByParam returns the OrderByParam field if non-nil, zero value otherwise.

### GetOrderByParamOk

`func (o *ConfusionMatrixParams) GetOrderByParamOk() (*string, bool)`

GetOrderByParamOk returns a tuple with the OrderByParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderByParam

`func (o *ConfusionMatrixParams) SetOrderByParam(v string)`

SetOrderByParam sets OrderByParam field to given value.

### HasOrderByParam

`func (o *ConfusionMatrixParams) HasOrderByParam() bool`

HasOrderByParam returns a boolean if a field has been set.

### GetOrderParams

`func (o *ConfusionMatrixParams) GetOrderParams() OrderType`

GetOrderParams returns the OrderParams field if non-nil, zero value otherwise.

### GetOrderParamsOk

`func (o *ConfusionMatrixParams) GetOrderParamsOk() (*OrderType, bool)`

GetOrderParamsOk returns a tuple with the OrderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderParams

`func (o *ConfusionMatrixParams) SetOrderParams(v OrderType)`

SetOrderParams sets OrderParams field to given value.

### HasOrderParams

`func (o *ConfusionMatrixParams) HasOrderParams() bool`

HasOrderParams returns a boolean if a field has been set.

### GetXAxisSizeInterval

`func (o *ConfusionMatrixParams) GetXAxisSizeInterval() float64`

GetXAxisSizeInterval returns the XAxisSizeInterval field if non-nil, zero value otherwise.

### GetXAxisSizeIntervalOk

`func (o *ConfusionMatrixParams) GetXAxisSizeIntervalOk() (*float64, bool)`

GetXAxisSizeIntervalOk returns a tuple with the XAxisSizeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAxisSizeInterval

`func (o *ConfusionMatrixParams) SetXAxisSizeInterval(v float64)`

SetXAxisSizeInterval sets XAxisSizeInterval field to given value.


### GetLastEpochOnly

`func (o *ConfusionMatrixParams) GetLastEpochOnly() bool`

GetLastEpochOnly returns the LastEpochOnly field if non-nil, zero value otherwise.

### GetLastEpochOnlyOk

`func (o *ConfusionMatrixParams) GetLastEpochOnlyOk() (*bool, bool)`

GetLastEpochOnlyOk returns a tuple with the LastEpochOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEpochOnly

`func (o *ConfusionMatrixParams) SetLastEpochOnly(v bool)`

SetLastEpochOnly sets LastEpochOnly field to given value.

### HasLastEpochOnly

`func (o *ConfusionMatrixParams) HasLastEpochOnly() bool`

HasLastEpochOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


