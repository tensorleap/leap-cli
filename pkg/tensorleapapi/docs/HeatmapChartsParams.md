# HeatmapChartsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**XField** | **string** |  | 
**XAxisSizeInterval** | **float64** |  | 
**XDataDistributionType** | [**DataDistributionType**](DataDistributionType.md) |  | 
**XOrderParams** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**YField** | **string** |  | 
**YAxisSizeInterval** | **float64** |  | 
**YDataDistributionType** | [**DataDistributionType**](DataDistributionType.md) |  | 
**YOrderParams** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**ColorField** | **string** |  | 
**ColorAggregationMethod** | [**AggregationMethod**](AggregationMethod.md) |  | 
**SessionRunIds** | **[]string** |  | 
**VerticalSplit** | Pointer to **string** |  | [optional] 
**HorizontalSplit** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**LastEpochOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewHeatmapChartsParams

`func NewHeatmapChartsParams(projectId string, xField string, xAxisSizeInterval float64, xDataDistributionType DataDistributionType, yField string, yAxisSizeInterval float64, yDataDistributionType DataDistributionType, colorField string, colorAggregationMethod AggregationMethod, sessionRunIds []string, ) *HeatmapChartsParams`

NewHeatmapChartsParams instantiates a new HeatmapChartsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeatmapChartsParamsWithDefaults

`func NewHeatmapChartsParamsWithDefaults() *HeatmapChartsParams`

NewHeatmapChartsParamsWithDefaults instantiates a new HeatmapChartsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *HeatmapChartsParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *HeatmapChartsParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *HeatmapChartsParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetXField

`func (o *HeatmapChartsParams) GetXField() string`

GetXField returns the XField field if non-nil, zero value otherwise.

### GetXFieldOk

`func (o *HeatmapChartsParams) GetXFieldOk() (*string, bool)`

GetXFieldOk returns a tuple with the XField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXField

`func (o *HeatmapChartsParams) SetXField(v string)`

SetXField sets XField field to given value.


### GetXAxisSizeInterval

`func (o *HeatmapChartsParams) GetXAxisSizeInterval() float64`

GetXAxisSizeInterval returns the XAxisSizeInterval field if non-nil, zero value otherwise.

### GetXAxisSizeIntervalOk

`func (o *HeatmapChartsParams) GetXAxisSizeIntervalOk() (*float64, bool)`

GetXAxisSizeIntervalOk returns a tuple with the XAxisSizeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAxisSizeInterval

`func (o *HeatmapChartsParams) SetXAxisSizeInterval(v float64)`

SetXAxisSizeInterval sets XAxisSizeInterval field to given value.


### GetXDataDistributionType

`func (o *HeatmapChartsParams) GetXDataDistributionType() DataDistributionType`

GetXDataDistributionType returns the XDataDistributionType field if non-nil, zero value otherwise.

### GetXDataDistributionTypeOk

`func (o *HeatmapChartsParams) GetXDataDistributionTypeOk() (*DataDistributionType, bool)`

GetXDataDistributionTypeOk returns a tuple with the XDataDistributionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDataDistributionType

`func (o *HeatmapChartsParams) SetXDataDistributionType(v DataDistributionType)`

SetXDataDistributionType sets XDataDistributionType field to given value.


### GetXOrderParams

`func (o *HeatmapChartsParams) GetXOrderParams() OrderType`

GetXOrderParams returns the XOrderParams field if non-nil, zero value otherwise.

### GetXOrderParamsOk

`func (o *HeatmapChartsParams) GetXOrderParamsOk() (*OrderType, bool)`

GetXOrderParamsOk returns a tuple with the XOrderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXOrderParams

`func (o *HeatmapChartsParams) SetXOrderParams(v OrderType)`

SetXOrderParams sets XOrderParams field to given value.

### HasXOrderParams

`func (o *HeatmapChartsParams) HasXOrderParams() bool`

HasXOrderParams returns a boolean if a field has been set.

### GetYField

`func (o *HeatmapChartsParams) GetYField() string`

GetYField returns the YField field if non-nil, zero value otherwise.

### GetYFieldOk

`func (o *HeatmapChartsParams) GetYFieldOk() (*string, bool)`

GetYFieldOk returns a tuple with the YField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYField

`func (o *HeatmapChartsParams) SetYField(v string)`

SetYField sets YField field to given value.


### GetYAxisSizeInterval

`func (o *HeatmapChartsParams) GetYAxisSizeInterval() float64`

GetYAxisSizeInterval returns the YAxisSizeInterval field if non-nil, zero value otherwise.

### GetYAxisSizeIntervalOk

`func (o *HeatmapChartsParams) GetYAxisSizeIntervalOk() (*float64, bool)`

GetYAxisSizeIntervalOk returns a tuple with the YAxisSizeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisSizeInterval

`func (o *HeatmapChartsParams) SetYAxisSizeInterval(v float64)`

SetYAxisSizeInterval sets YAxisSizeInterval field to given value.


### GetYDataDistributionType

`func (o *HeatmapChartsParams) GetYDataDistributionType() DataDistributionType`

GetYDataDistributionType returns the YDataDistributionType field if non-nil, zero value otherwise.

### GetYDataDistributionTypeOk

`func (o *HeatmapChartsParams) GetYDataDistributionTypeOk() (*DataDistributionType, bool)`

GetYDataDistributionTypeOk returns a tuple with the YDataDistributionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYDataDistributionType

`func (o *HeatmapChartsParams) SetYDataDistributionType(v DataDistributionType)`

SetYDataDistributionType sets YDataDistributionType field to given value.


### GetYOrderParams

`func (o *HeatmapChartsParams) GetYOrderParams() OrderType`

GetYOrderParams returns the YOrderParams field if non-nil, zero value otherwise.

### GetYOrderParamsOk

`func (o *HeatmapChartsParams) GetYOrderParamsOk() (*OrderType, bool)`

GetYOrderParamsOk returns a tuple with the YOrderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYOrderParams

`func (o *HeatmapChartsParams) SetYOrderParams(v OrderType)`

SetYOrderParams sets YOrderParams field to given value.

### HasYOrderParams

`func (o *HeatmapChartsParams) HasYOrderParams() bool`

HasYOrderParams returns a boolean if a field has been set.

### GetColorField

`func (o *HeatmapChartsParams) GetColorField() string`

GetColorField returns the ColorField field if non-nil, zero value otherwise.

### GetColorFieldOk

`func (o *HeatmapChartsParams) GetColorFieldOk() (*string, bool)`

GetColorFieldOk returns a tuple with the ColorField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorField

`func (o *HeatmapChartsParams) SetColorField(v string)`

SetColorField sets ColorField field to given value.


### GetColorAggregationMethod

`func (o *HeatmapChartsParams) GetColorAggregationMethod() AggregationMethod`

GetColorAggregationMethod returns the ColorAggregationMethod field if non-nil, zero value otherwise.

### GetColorAggregationMethodOk

`func (o *HeatmapChartsParams) GetColorAggregationMethodOk() (*AggregationMethod, bool)`

GetColorAggregationMethodOk returns a tuple with the ColorAggregationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorAggregationMethod

`func (o *HeatmapChartsParams) SetColorAggregationMethod(v AggregationMethod)`

SetColorAggregationMethod sets ColorAggregationMethod field to given value.


### GetSessionRunIds

`func (o *HeatmapChartsParams) GetSessionRunIds() []string`

GetSessionRunIds returns the SessionRunIds field if non-nil, zero value otherwise.

### GetSessionRunIdsOk

`func (o *HeatmapChartsParams) GetSessionRunIdsOk() (*[]string, bool)`

GetSessionRunIdsOk returns a tuple with the SessionRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunIds

`func (o *HeatmapChartsParams) SetSessionRunIds(v []string)`

SetSessionRunIds sets SessionRunIds field to given value.


### GetVerticalSplit

`func (o *HeatmapChartsParams) GetVerticalSplit() string`

GetVerticalSplit returns the VerticalSplit field if non-nil, zero value otherwise.

### GetVerticalSplitOk

`func (o *HeatmapChartsParams) GetVerticalSplitOk() (*string, bool)`

GetVerticalSplitOk returns a tuple with the VerticalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalSplit

`func (o *HeatmapChartsParams) SetVerticalSplit(v string)`

SetVerticalSplit sets VerticalSplit field to given value.

### HasVerticalSplit

`func (o *HeatmapChartsParams) HasVerticalSplit() bool`

HasVerticalSplit returns a boolean if a field has been set.

### GetHorizontalSplit

`func (o *HeatmapChartsParams) GetHorizontalSplit() string`

GetHorizontalSplit returns the HorizontalSplit field if non-nil, zero value otherwise.

### GetHorizontalSplitOk

`func (o *HeatmapChartsParams) GetHorizontalSplitOk() (*string, bool)`

GetHorizontalSplitOk returns a tuple with the HorizontalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalSplit

`func (o *HeatmapChartsParams) SetHorizontalSplit(v string)`

SetHorizontalSplit sets HorizontalSplit field to given value.

### HasHorizontalSplit

`func (o *HeatmapChartsParams) HasHorizontalSplit() bool`

HasHorizontalSplit returns a boolean if a field has been set.

### GetFilters

`func (o *HeatmapChartsParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *HeatmapChartsParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *HeatmapChartsParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *HeatmapChartsParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLastEpochOnly

`func (o *HeatmapChartsParams) GetLastEpochOnly() bool`

GetLastEpochOnly returns the LastEpochOnly field if non-nil, zero value otherwise.

### GetLastEpochOnlyOk

`func (o *HeatmapChartsParams) GetLastEpochOnlyOk() (*bool, bool)`

GetLastEpochOnlyOk returns a tuple with the LastEpochOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEpochOnly

`func (o *HeatmapChartsParams) SetLastEpochOnly(v bool)`

SetLastEpochOnly sets LastEpochOnly field to given value.

### HasLastEpochOnly

`func (o *HeatmapChartsParams) HasLastEpochOnly() bool`

HasLastEpochOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


