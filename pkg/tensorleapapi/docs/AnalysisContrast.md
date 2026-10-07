# AnalysisContrast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** |  | 
**GroupMedian** | **float64** |  | 
**GroupMean** | **float64** |  | 
**BaselineMedian** | Pointer to **float64** |  | [optional] 
**BaselineMean** | Pointer to **float64** |  | [optional] 
**Baseline** | **string** |  | 

## Methods

### NewAnalysisContrast

`func NewAnalysisContrast(metric string, groupMedian float64, groupMean float64, baseline string, ) *AnalysisContrast`

NewAnalysisContrast instantiates a new AnalysisContrast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisContrastWithDefaults

`func NewAnalysisContrastWithDefaults() *AnalysisContrast`

NewAnalysisContrastWithDefaults instantiates a new AnalysisContrast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *AnalysisContrast) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AnalysisContrast) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AnalysisContrast) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetGroupMedian

`func (o *AnalysisContrast) GetGroupMedian() float64`

GetGroupMedian returns the GroupMedian field if non-nil, zero value otherwise.

### GetGroupMedianOk

`func (o *AnalysisContrast) GetGroupMedianOk() (*float64, bool)`

GetGroupMedianOk returns a tuple with the GroupMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMedian

`func (o *AnalysisContrast) SetGroupMedian(v float64)`

SetGroupMedian sets GroupMedian field to given value.


### GetGroupMean

`func (o *AnalysisContrast) GetGroupMean() float64`

GetGroupMean returns the GroupMean field if non-nil, zero value otherwise.

### GetGroupMeanOk

`func (o *AnalysisContrast) GetGroupMeanOk() (*float64, bool)`

GetGroupMeanOk returns a tuple with the GroupMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMean

`func (o *AnalysisContrast) SetGroupMean(v float64)`

SetGroupMean sets GroupMean field to given value.


### GetBaselineMedian

`func (o *AnalysisContrast) GetBaselineMedian() float64`

GetBaselineMedian returns the BaselineMedian field if non-nil, zero value otherwise.

### GetBaselineMedianOk

`func (o *AnalysisContrast) GetBaselineMedianOk() (*float64, bool)`

GetBaselineMedianOk returns a tuple with the BaselineMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineMedian

`func (o *AnalysisContrast) SetBaselineMedian(v float64)`

SetBaselineMedian sets BaselineMedian field to given value.

### HasBaselineMedian

`func (o *AnalysisContrast) HasBaselineMedian() bool`

HasBaselineMedian returns a boolean if a field has been set.

### GetBaselineMean

`func (o *AnalysisContrast) GetBaselineMean() float64`

GetBaselineMean returns the BaselineMean field if non-nil, zero value otherwise.

### GetBaselineMeanOk

`func (o *AnalysisContrast) GetBaselineMeanOk() (*float64, bool)`

GetBaselineMeanOk returns a tuple with the BaselineMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineMean

`func (o *AnalysisContrast) SetBaselineMean(v float64)`

SetBaselineMean sets BaselineMean field to given value.

### HasBaselineMean

`func (o *AnalysisContrast) HasBaselineMean() bool`

HasBaselineMean returns a boolean if a field has been set.

### GetBaseline

`func (o *AnalysisContrast) GetBaseline() string`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *AnalysisContrast) GetBaselineOk() (*string, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *AnalysisContrast) SetBaseline(v string)`

SetBaseline sets Baseline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


