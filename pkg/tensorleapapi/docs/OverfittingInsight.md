# OverfittingInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ScatterInsightType**](ScatterInsightType.md) |  | 
**Filters** | Pointer to [**[]ScatterFilter**](ScatterFilter.md) |  | [optional] 
**NSamples** | **float64** |  | 
**MutualInfoElements** | [**[]MutualInformationElement**](MutualInformationElement.md) |  | 
**BlobPath** | **string** |  | 
**InsightSubCategoryDsCuration** | Pointer to [**InsightSubCategoryDSCuration**](InsightSubCategoryDSCuration.md) |  | [optional] 
**InsightCategoryPerformance** | Pointer to [**InsightSubCategoryPerformance**](InsightSubCategoryPerformance.md) |  | [optional] 
**Severity** | **float64** |  | 
**SeverityMetrics** | [**[]SeverityMetricElement**](SeverityMetricElement.md) |  | 
**MetricsInfo** | [**[]InsightMetricInfo**](InsightMetricInfo.md) |  | 
**Index** | **float64** |  | 
**MinHash** | **[]float64** |  | 
**CsvPath** | Pointer to **string** |  | [optional] 
**AutomaticTests** | Pointer to [**[]InsightAutomaticTest**](InsightAutomaticTest.md) |  | [optional] 
**NSamplesValidation** | **float64** |  | 
**NSamplesTraining** | **float64** |  | 

## Methods

### NewOverfittingInsight

`func NewOverfittingInsight(type_ ScatterInsightType, nSamples float64, mutualInfoElements []MutualInformationElement, blobPath string, severity float64, severityMetrics []SeverityMetricElement, metricsInfo []InsightMetricInfo, index float64, minHash []float64, nSamplesValidation float64, nSamplesTraining float64, ) *OverfittingInsight`

NewOverfittingInsight instantiates a new OverfittingInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverfittingInsightWithDefaults

`func NewOverfittingInsightWithDefaults() *OverfittingInsight`

NewOverfittingInsightWithDefaults instantiates a new OverfittingInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OverfittingInsight) GetType() ScatterInsightType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverfittingInsight) GetTypeOk() (*ScatterInsightType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverfittingInsight) SetType(v ScatterInsightType)`

SetType sets Type field to given value.


### GetFilters

`func (o *OverfittingInsight) GetFilters() []ScatterFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *OverfittingInsight) GetFiltersOk() (*[]ScatterFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *OverfittingInsight) SetFilters(v []ScatterFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *OverfittingInsight) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetNSamples

`func (o *OverfittingInsight) GetNSamples() float64`

GetNSamples returns the NSamples field if non-nil, zero value otherwise.

### GetNSamplesOk

`func (o *OverfittingInsight) GetNSamplesOk() (*float64, bool)`

GetNSamplesOk returns a tuple with the NSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSamples

`func (o *OverfittingInsight) SetNSamples(v float64)`

SetNSamples sets NSamples field to given value.


### GetMutualInfoElements

`func (o *OverfittingInsight) GetMutualInfoElements() []MutualInformationElement`

GetMutualInfoElements returns the MutualInfoElements field if non-nil, zero value otherwise.

### GetMutualInfoElementsOk

`func (o *OverfittingInsight) GetMutualInfoElementsOk() (*[]MutualInformationElement, bool)`

GetMutualInfoElementsOk returns a tuple with the MutualInfoElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualInfoElements

`func (o *OverfittingInsight) SetMutualInfoElements(v []MutualInformationElement)`

SetMutualInfoElements sets MutualInfoElements field to given value.


### GetBlobPath

`func (o *OverfittingInsight) GetBlobPath() string`

GetBlobPath returns the BlobPath field if non-nil, zero value otherwise.

### GetBlobPathOk

`func (o *OverfittingInsight) GetBlobPathOk() (*string, bool)`

GetBlobPathOk returns a tuple with the BlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobPath

`func (o *OverfittingInsight) SetBlobPath(v string)`

SetBlobPath sets BlobPath field to given value.


### GetInsightSubCategoryDsCuration

`func (o *OverfittingInsight) GetInsightSubCategoryDsCuration() InsightSubCategoryDSCuration`

GetInsightSubCategoryDsCuration returns the InsightSubCategoryDsCuration field if non-nil, zero value otherwise.

### GetInsightSubCategoryDsCurationOk

`func (o *OverfittingInsight) GetInsightSubCategoryDsCurationOk() (*InsightSubCategoryDSCuration, bool)`

GetInsightSubCategoryDsCurationOk returns a tuple with the InsightSubCategoryDsCuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightSubCategoryDsCuration

`func (o *OverfittingInsight) SetInsightSubCategoryDsCuration(v InsightSubCategoryDSCuration)`

SetInsightSubCategoryDsCuration sets InsightSubCategoryDsCuration field to given value.

### HasInsightSubCategoryDsCuration

`func (o *OverfittingInsight) HasInsightSubCategoryDsCuration() bool`

HasInsightSubCategoryDsCuration returns a boolean if a field has been set.

### GetInsightCategoryPerformance

`func (o *OverfittingInsight) GetInsightCategoryPerformance() InsightSubCategoryPerformance`

GetInsightCategoryPerformance returns the InsightCategoryPerformance field if non-nil, zero value otherwise.

### GetInsightCategoryPerformanceOk

`func (o *OverfittingInsight) GetInsightCategoryPerformanceOk() (*InsightSubCategoryPerformance, bool)`

GetInsightCategoryPerformanceOk returns a tuple with the InsightCategoryPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightCategoryPerformance

`func (o *OverfittingInsight) SetInsightCategoryPerformance(v InsightSubCategoryPerformance)`

SetInsightCategoryPerformance sets InsightCategoryPerformance field to given value.

### HasInsightCategoryPerformance

`func (o *OverfittingInsight) HasInsightCategoryPerformance() bool`

HasInsightCategoryPerformance returns a boolean if a field has been set.

### GetSeverity

`func (o *OverfittingInsight) GetSeverity() float64`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *OverfittingInsight) GetSeverityOk() (*float64, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *OverfittingInsight) SetSeverity(v float64)`

SetSeverity sets Severity field to given value.


### GetSeverityMetrics

`func (o *OverfittingInsight) GetSeverityMetrics() []SeverityMetricElement`

GetSeverityMetrics returns the SeverityMetrics field if non-nil, zero value otherwise.

### GetSeverityMetricsOk

`func (o *OverfittingInsight) GetSeverityMetricsOk() (*[]SeverityMetricElement, bool)`

GetSeverityMetricsOk returns a tuple with the SeverityMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityMetrics

`func (o *OverfittingInsight) SetSeverityMetrics(v []SeverityMetricElement)`

SetSeverityMetrics sets SeverityMetrics field to given value.


### GetMetricsInfo

`func (o *OverfittingInsight) GetMetricsInfo() []InsightMetricInfo`

GetMetricsInfo returns the MetricsInfo field if non-nil, zero value otherwise.

### GetMetricsInfoOk

`func (o *OverfittingInsight) GetMetricsInfoOk() (*[]InsightMetricInfo, bool)`

GetMetricsInfoOk returns a tuple with the MetricsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsInfo

`func (o *OverfittingInsight) SetMetricsInfo(v []InsightMetricInfo)`

SetMetricsInfo sets MetricsInfo field to given value.


### GetIndex

`func (o *OverfittingInsight) GetIndex() float64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *OverfittingInsight) GetIndexOk() (*float64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *OverfittingInsight) SetIndex(v float64)`

SetIndex sets Index field to given value.


### GetMinHash

`func (o *OverfittingInsight) GetMinHash() []float64`

GetMinHash returns the MinHash field if non-nil, zero value otherwise.

### GetMinHashOk

`func (o *OverfittingInsight) GetMinHashOk() (*[]float64, bool)`

GetMinHashOk returns a tuple with the MinHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHash

`func (o *OverfittingInsight) SetMinHash(v []float64)`

SetMinHash sets MinHash field to given value.


### GetCsvPath

`func (o *OverfittingInsight) GetCsvPath() string`

GetCsvPath returns the CsvPath field if non-nil, zero value otherwise.

### GetCsvPathOk

`func (o *OverfittingInsight) GetCsvPathOk() (*string, bool)`

GetCsvPathOk returns a tuple with the CsvPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvPath

`func (o *OverfittingInsight) SetCsvPath(v string)`

SetCsvPath sets CsvPath field to given value.

### HasCsvPath

`func (o *OverfittingInsight) HasCsvPath() bool`

HasCsvPath returns a boolean if a field has been set.

### GetAutomaticTests

`func (o *OverfittingInsight) GetAutomaticTests() []InsightAutomaticTest`

GetAutomaticTests returns the AutomaticTests field if non-nil, zero value otherwise.

### GetAutomaticTestsOk

`func (o *OverfittingInsight) GetAutomaticTestsOk() (*[]InsightAutomaticTest, bool)`

GetAutomaticTestsOk returns a tuple with the AutomaticTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticTests

`func (o *OverfittingInsight) SetAutomaticTests(v []InsightAutomaticTest)`

SetAutomaticTests sets AutomaticTests field to given value.

### HasAutomaticTests

`func (o *OverfittingInsight) HasAutomaticTests() bool`

HasAutomaticTests returns a boolean if a field has been set.

### GetNSamplesValidation

`func (o *OverfittingInsight) GetNSamplesValidation() float64`

GetNSamplesValidation returns the NSamplesValidation field if non-nil, zero value otherwise.

### GetNSamplesValidationOk

`func (o *OverfittingInsight) GetNSamplesValidationOk() (*float64, bool)`

GetNSamplesValidationOk returns a tuple with the NSamplesValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSamplesValidation

`func (o *OverfittingInsight) SetNSamplesValidation(v float64)`

SetNSamplesValidation sets NSamplesValidation field to given value.


### GetNSamplesTraining

`func (o *OverfittingInsight) GetNSamplesTraining() float64`

GetNSamplesTraining returns the NSamplesTraining field if non-nil, zero value otherwise.

### GetNSamplesTrainingOk

`func (o *OverfittingInsight) GetNSamplesTrainingOk() (*float64, bool)`

GetNSamplesTrainingOk returns a tuple with the NSamplesTraining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSamplesTraining

`func (o *OverfittingInsight) SetNSamplesTraining(v float64)`

SetNSamplesTraining sets NSamplesTraining field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


