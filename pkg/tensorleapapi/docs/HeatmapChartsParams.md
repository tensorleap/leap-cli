# HeatmapChartsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**X** | [**SplitAgg**](SplitAgg.md) |  | 
**Y** | [**SplitAgg**](SplitAgg.md) |  | 
**Color** | [**Aggregations**](Aggregations.md) |  | 
**SessionRunsToEpochs** | [**[]SessionRunToEpoch**](SessionRunToEpoch.md) |  | 
**ShowAllEpochs** | **bool** |  | 
**VerticalSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**HorizontalSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 

## Methods

### NewHeatmapChartsParams

`func NewHeatmapChartsParams(projectId string, x SplitAgg, y SplitAgg, color Aggregations, sessionRunsToEpochs []SessionRunToEpoch, showAllEpochs bool, ) *HeatmapChartsParams`

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


### GetX

`func (o *HeatmapChartsParams) GetX() SplitAgg`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *HeatmapChartsParams) GetXOk() (*SplitAgg, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *HeatmapChartsParams) SetX(v SplitAgg)`

SetX sets X field to given value.


### GetY

`func (o *HeatmapChartsParams) GetY() SplitAgg`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *HeatmapChartsParams) GetYOk() (*SplitAgg, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *HeatmapChartsParams) SetY(v SplitAgg)`

SetY sets Y field to given value.


### GetColor

`func (o *HeatmapChartsParams) GetColor() Aggregations`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *HeatmapChartsParams) GetColorOk() (*Aggregations, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *HeatmapChartsParams) SetColor(v Aggregations)`

SetColor sets Color field to given value.


### GetSessionRunsToEpochs

`func (o *HeatmapChartsParams) GetSessionRunsToEpochs() []SessionRunToEpoch`

GetSessionRunsToEpochs returns the SessionRunsToEpochs field if non-nil, zero value otherwise.

### GetSessionRunsToEpochsOk

`func (o *HeatmapChartsParams) GetSessionRunsToEpochsOk() (*[]SessionRunToEpoch, bool)`

GetSessionRunsToEpochsOk returns a tuple with the SessionRunsToEpochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunsToEpochs

`func (o *HeatmapChartsParams) SetSessionRunsToEpochs(v []SessionRunToEpoch)`

SetSessionRunsToEpochs sets SessionRunsToEpochs field to given value.


### GetShowAllEpochs

`func (o *HeatmapChartsParams) GetShowAllEpochs() bool`

GetShowAllEpochs returns the ShowAllEpochs field if non-nil, zero value otherwise.

### GetShowAllEpochsOk

`func (o *HeatmapChartsParams) GetShowAllEpochsOk() (*bool, bool)`

GetShowAllEpochsOk returns a tuple with the ShowAllEpochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAllEpochs

`func (o *HeatmapChartsParams) SetShowAllEpochs(v bool)`

SetShowAllEpochs sets ShowAllEpochs field to given value.


### GetVerticalSplit

`func (o *HeatmapChartsParams) GetVerticalSplit() SplitAgg`

GetVerticalSplit returns the VerticalSplit field if non-nil, zero value otherwise.

### GetVerticalSplitOk

`func (o *HeatmapChartsParams) GetVerticalSplitOk() (*SplitAgg, bool)`

GetVerticalSplitOk returns a tuple with the VerticalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalSplit

`func (o *HeatmapChartsParams) SetVerticalSplit(v SplitAgg)`

SetVerticalSplit sets VerticalSplit field to given value.

### HasVerticalSplit

`func (o *HeatmapChartsParams) HasVerticalSplit() bool`

HasVerticalSplit returns a boolean if a field has been set.

### GetHorizontalSplit

`func (o *HeatmapChartsParams) GetHorizontalSplit() SplitAgg`

GetHorizontalSplit returns the HorizontalSplit field if non-nil, zero value otherwise.

### GetHorizontalSplitOk

`func (o *HeatmapChartsParams) GetHorizontalSplitOk() (*SplitAgg, bool)`

GetHorizontalSplitOk returns a tuple with the HorizontalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalSplit

`func (o *HeatmapChartsParams) SetHorizontalSplit(v SplitAgg)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


