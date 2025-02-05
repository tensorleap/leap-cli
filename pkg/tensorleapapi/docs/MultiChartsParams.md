# MultiChartsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**X** | [**SplitAgg**](SplitAgg.md) |  | 
**Y** | [**Aggregations**](Aggregations.md) |  | 
**VerticalSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**HorizontalSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**InnerSplit** | Pointer to [**SplitAgg**](SplitAgg.md) |  | [optional] 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**SessionRunsToEpochs** | [**[]SessionRunToEpoch**](SessionRunToEpoch.md) |  | 
**ShowAllEpochs** | **bool** |  | 

## Methods

### NewMultiChartsParams

`func NewMultiChartsParams(projectId string, x SplitAgg, y Aggregations, sessionRunsToEpochs []SessionRunToEpoch, showAllEpochs bool, ) *MultiChartsParams`

NewMultiChartsParams instantiates a new MultiChartsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiChartsParamsWithDefaults

`func NewMultiChartsParamsWithDefaults() *MultiChartsParams`

NewMultiChartsParamsWithDefaults instantiates a new MultiChartsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *MultiChartsParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MultiChartsParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MultiChartsParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetX

`func (o *MultiChartsParams) GetX() SplitAgg`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *MultiChartsParams) GetXOk() (*SplitAgg, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *MultiChartsParams) SetX(v SplitAgg)`

SetX sets X field to given value.


### GetY

`func (o *MultiChartsParams) GetY() Aggregations`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *MultiChartsParams) GetYOk() (*Aggregations, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *MultiChartsParams) SetY(v Aggregations)`

SetY sets Y field to given value.


### GetVerticalSplit

`func (o *MultiChartsParams) GetVerticalSplit() SplitAgg`

GetVerticalSplit returns the VerticalSplit field if non-nil, zero value otherwise.

### GetVerticalSplitOk

`func (o *MultiChartsParams) GetVerticalSplitOk() (*SplitAgg, bool)`

GetVerticalSplitOk returns a tuple with the VerticalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalSplit

`func (o *MultiChartsParams) SetVerticalSplit(v SplitAgg)`

SetVerticalSplit sets VerticalSplit field to given value.

### HasVerticalSplit

`func (o *MultiChartsParams) HasVerticalSplit() bool`

HasVerticalSplit returns a boolean if a field has been set.

### GetHorizontalSplit

`func (o *MultiChartsParams) GetHorizontalSplit() SplitAgg`

GetHorizontalSplit returns the HorizontalSplit field if non-nil, zero value otherwise.

### GetHorizontalSplitOk

`func (o *MultiChartsParams) GetHorizontalSplitOk() (*SplitAgg, bool)`

GetHorizontalSplitOk returns a tuple with the HorizontalSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalSplit

`func (o *MultiChartsParams) SetHorizontalSplit(v SplitAgg)`

SetHorizontalSplit sets HorizontalSplit field to given value.

### HasHorizontalSplit

`func (o *MultiChartsParams) HasHorizontalSplit() bool`

HasHorizontalSplit returns a boolean if a field has been set.

### GetInnerSplit

`func (o *MultiChartsParams) GetInnerSplit() SplitAgg`

GetInnerSplit returns the InnerSplit field if non-nil, zero value otherwise.

### GetInnerSplitOk

`func (o *MultiChartsParams) GetInnerSplitOk() (*SplitAgg, bool)`

GetInnerSplitOk returns a tuple with the InnerSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerSplit

`func (o *MultiChartsParams) SetInnerSplit(v SplitAgg)`

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

### GetSessionRunsToEpochs

`func (o *MultiChartsParams) GetSessionRunsToEpochs() []SessionRunToEpoch`

GetSessionRunsToEpochs returns the SessionRunsToEpochs field if non-nil, zero value otherwise.

### GetSessionRunsToEpochsOk

`func (o *MultiChartsParams) GetSessionRunsToEpochsOk() (*[]SessionRunToEpoch, bool)`

GetSessionRunsToEpochsOk returns a tuple with the SessionRunsToEpochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunsToEpochs

`func (o *MultiChartsParams) SetSessionRunsToEpochs(v []SessionRunToEpoch)`

SetSessionRunsToEpochs sets SessionRunsToEpochs field to given value.


### GetShowAllEpochs

`func (o *MultiChartsParams) GetShowAllEpochs() bool`

GetShowAllEpochs returns the ShowAllEpochs field if non-nil, zero value otherwise.

### GetShowAllEpochsOk

`func (o *MultiChartsParams) GetShowAllEpochsOk() (*bool, bool)`

GetShowAllEpochsOk returns a tuple with the ShowAllEpochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAllEpochs

`func (o *MultiChartsParams) SetShowAllEpochs(v bool)`

SetShowAllEpochs sets ShowAllEpochs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


