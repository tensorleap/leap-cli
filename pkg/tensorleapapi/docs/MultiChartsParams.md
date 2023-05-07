# MultiChartsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**XField** | **string** |  | 
**YField** | **string** |  | 
**AggregationMethod** | [**AggregationMethod**](AggregationMethod.md) |  | 
**SessionRunIds** | **[]string** |  | 
**VerticalSplit** | Pointer to **string** |  | [optional] 
**HorizontalSplit** | Pointer to **string** |  | [optional] 
**InnerSplit** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**DataDistributionType** | [**DataDistributionType**](DataDistributionType.md) |  | 
**OrderByParam** | Pointer to **string** |  | [optional] 
**OrderParams** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**XAxisSizeInterval** | **float64** |  | 
**LastEpochOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewMultiChartsParams

`func NewMultiChartsParams(xField string, yField string, aggregationMethod AggregationMethod, sessionRunIds []string, dataDistributionType DataDistributionType, xAxisSizeInterval float64, ) *MultiChartsParams`

NewMultiChartsParams instantiates a new MultiChartsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiChartsParamsWithDefaults

`func NewMultiChartsParamsWithDefaults() *MultiChartsParams`

NewMultiChartsParamsWithDefaults instantiates a new MultiChartsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXField

`func (o *MultiChartsParams) GetXField() string`

GetXField returns the XField field if non-nil, zero value otherwise.

### GetXFieldOk

`func (o *MultiChartsParams) GetXFieldOk() (*string, bool)`

GetXFieldOk returns a tuple with the XField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXField

`func (o *MultiChartsParams) SetXField(v string)`

SetXField sets XField field to given value.


### GetYField

`func (o *MultiChartsParams) GetYField() string`

GetYField returns the YField field if non-nil, zero value otherwise.

### GetYFieldOk

`func (o *MultiChartsParams) GetYFieldOk() (*string, bool)`

GetYFieldOk returns a tuple with the YField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYField

`func (o *MultiChartsParams) SetYField(v string)`

SetYField sets YField field to given value.


### GetAggregationMethod

`func (o *MultiChartsParams) GetAggregationMethod() AggregationMethod`

GetAggregationMethod returns the AggregationMethod field if non-nil, zero value otherwise.

### GetAggregationMethodOk

`func (o *MultiChartsParams) GetAggregationMethodOk() (*AggregationMethod, bool)`

GetAggregationMethodOk returns a tuple with the AggregationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationMethod

`func (o *MultiChartsParams) SetAggregationMethod(v AggregationMethod)`

SetAggregationMethod sets AggregationMethod field to given value.


### GetSessionRunIds

`func (o *MultiChartsParams) GetSessionRunIds() []string`

GetSessionRunIds returns the SessionRunIds field if non-nil, zero value otherwise.

### GetSessionRunIdsOk

`func (o *MultiChartsParams) GetSessionRunIdsOk() (*[]string, bool)`

GetSessionRunIdsOk returns a tuple with the SessionRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunIds

`func (o *MultiChartsParams) SetSessionRunIds(v []string)`

SetSessionRunIds sets SessionRunIds field to given value.


### GetVerticalSplit

`func (o *MultiChartsParams) GetVerticalSplit() string`

GetVerticalSplit returns the VerticalSplit field if non-nil, zero value otherwise.

### GetVerticalSplitOk

`func (o *MultiChartsParams) GetVerticalSplitOk() (*string, bool)`

GetVerticalSplitOk returns a tuple with the VerticalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalSplit

`func (o *MultiChartsParams) SetVerticalSplit(v string)`

SetVerticalSplit sets VerticalSplit field to given value.

### HasVerticalSplit

`func (o *MultiChartsParams) HasVerticalSplit() bool`

HasVerticalSplit returns a boolean if a field has been set.

### GetHorizontalSplit

`func (o *MultiChartsParams) GetHorizontalSplit() string`

GetHorizontalSplit returns the HorizontalSplit field if non-nil, zero value otherwise.

### GetHorizontalSplitOk

`func (o *MultiChartsParams) GetHorizontalSplitOk() (*string, bool)`

GetHorizontalSplitOk returns a tuple with the HorizontalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalSplit

`func (o *MultiChartsParams) SetHorizontalSplit(v string)`

SetHorizontalSplit sets HorizontalSplit field to given value.

### HasHorizontalSplit

`func (o *MultiChartsParams) HasHorizontalSplit() bool`

HasHorizontalSplit returns a boolean if a field has been set.

### GetInnerSplit

`func (o *MultiChartsParams) GetInnerSplit() string`

GetInnerSplit returns the InnerSplit field if non-nil, zero value otherwise.

### GetInnerSplitOk

`func (o *MultiChartsParams) GetInnerSplitOk() (*string, bool)`

GetInnerSplitOk returns a tuple with the InnerSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerSplit

`func (o *MultiChartsParams) SetInnerSplit(v string)`

SetInnerSplit sets InnerSplit field to given value.

### HasInnerSplit

`func (o *MultiChartsParams) HasInnerSplit() bool`

HasInnerSplit returns a boolean if a field has been set.

### GetFilters

`func (o *MultiChartsParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MultiChartsParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MultiChartsParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MultiChartsParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetDataDistributionType

`func (o *MultiChartsParams) GetDataDistributionType() DataDistributionType`

GetDataDistributionType returns the DataDistributionType field if non-nil, zero value otherwise.

### GetDataDistributionTypeOk

`func (o *MultiChartsParams) GetDataDistributionTypeOk() (*DataDistributionType, bool)`

GetDataDistributionTypeOk returns a tuple with the DataDistributionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDistributionType

`func (o *MultiChartsParams) SetDataDistributionType(v DataDistributionType)`

SetDataDistributionType sets DataDistributionType field to given value.


### GetOrderByParam

`func (o *MultiChartsParams) GetOrderByParam() string`

GetOrderByParam returns the OrderByParam field if non-nil, zero value otherwise.

### GetOrderByParamOk

`func (o *MultiChartsParams) GetOrderByParamOk() (*string, bool)`

GetOrderByParamOk returns a tuple with the OrderByParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderByParam

`func (o *MultiChartsParams) SetOrderByParam(v string)`

SetOrderByParam sets OrderByParam field to given value.

### HasOrderByParam

`func (o *MultiChartsParams) HasOrderByParam() bool`

HasOrderByParam returns a boolean if a field has been set.

### GetOrderParams

`func (o *MultiChartsParams) GetOrderParams() OrderType`

GetOrderParams returns the OrderParams field if non-nil, zero value otherwise.

### GetOrderParamsOk

`func (o *MultiChartsParams) GetOrderParamsOk() (*OrderType, bool)`

GetOrderParamsOk returns a tuple with the OrderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderParams

`func (o *MultiChartsParams) SetOrderParams(v OrderType)`

SetOrderParams sets OrderParams field to given value.

### HasOrderParams

`func (o *MultiChartsParams) HasOrderParams() bool`

HasOrderParams returns a boolean if a field has been set.

### GetXAxisSizeInterval

`func (o *MultiChartsParams) GetXAxisSizeInterval() float64`

GetXAxisSizeInterval returns the XAxisSizeInterval field if non-nil, zero value otherwise.

### GetXAxisSizeIntervalOk

`func (o *MultiChartsParams) GetXAxisSizeIntervalOk() (*float64, bool)`

GetXAxisSizeIntervalOk returns a tuple with the XAxisSizeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAxisSizeInterval

`func (o *MultiChartsParams) SetXAxisSizeInterval(v float64)`

SetXAxisSizeInterval sets XAxisSizeInterval field to given value.


### GetLastEpochOnly

`func (o *MultiChartsParams) GetLastEpochOnly() bool`

GetLastEpochOnly returns the LastEpochOnly field if non-nil, zero value otherwise.

### GetLastEpochOnlyOk

`func (o *MultiChartsParams) GetLastEpochOnlyOk() (*bool, bool)`

GetLastEpochOnlyOk returns a tuple with the LastEpochOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEpochOnly

`func (o *MultiChartsParams) SetLastEpochOnly(v bool)`

SetLastEpochOnly sets LastEpochOnly field to given value.

### HasLastEpochOnly

`func (o *MultiChartsParams) HasLastEpochOnly() bool`

HasLastEpochOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


