# UnderRepresentationInsight

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
**UnderRepresentationDataset** | [**DataStateType**](DataStateType.md) |  | 
**UnderRepresentationNSamples** | **float64** |  | 
**OverRepresentationDataset** | [**DataStateType**](DataStateType.md) |  | 
**OverRepresentationNSamples** | **float64** |  | 

## Methods

### NewUnderRepresentationInsight

`func NewUnderRepresentationInsight(type_ ScatterInsightType, nSamples float64, mutualInfoElements []MutualInformationElement, blobPath string, severity float64, severityMetrics []SeverityMetricElement, metricsInfo []InsightMetricInfo, index float64, minHash []float64, underRepresentationDataset DataStateType, underRepresentationNSamples float64, overRepresentationDataset DataStateType, overRepresentationNSamples float64, ) *UnderRepresentationInsight`

NewUnderRepresentationInsight instantiates a new UnderRepresentationInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnderRepresentationInsightWithDefaults

`func NewUnderRepresentationInsightWithDefaults() *UnderRepresentationInsight`

NewUnderRepresentationInsightWithDefaults instantiates a new UnderRepresentationInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UnderRepresentationInsight) GetType() ScatterInsightType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnderRepresentationInsight) GetTypeOk() (*ScatterInsightType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnderRepresentationInsight) SetType(v ScatterInsightType)`

SetType sets Type field to given value.


### GetFilters

`func (o *UnderRepresentationInsight) GetFilters() []ScatterFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *UnderRepresentationInsight) GetFiltersOk() (*[]ScatterFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *UnderRepresentationInsight) SetFilters(v []ScatterFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *UnderRepresentationInsight) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetNSamples

`func (o *UnderRepresentationInsight) GetNSamples() float64`

GetNSamples returns the NSamples field if non-nil, zero value otherwise.

### GetNSamplesOk

`func (o *UnderRepresentationInsight) GetNSamplesOk() (*float64, bool)`

GetNSamplesOk returns a tuple with the NSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSamples

`func (o *UnderRepresentationInsight) SetNSamples(v float64)`

SetNSamples sets NSamples field to given value.


### GetMutualInfoElements

`func (o *UnderRepresentationInsight) GetMutualInfoElements() []MutualInformationElement`

GetMutualInfoElements returns the MutualInfoElements field if non-nil, zero value otherwise.

### GetMutualInfoElementsOk

`func (o *UnderRepresentationInsight) GetMutualInfoElementsOk() (*[]MutualInformationElement, bool)`

GetMutualInfoElementsOk returns a tuple with the MutualInfoElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualInfoElements

`func (o *UnderRepresentationInsight) SetMutualInfoElements(v []MutualInformationElement)`

SetMutualInfoElements sets MutualInfoElements field to given value.


### GetBlobPath

`func (o *UnderRepresentationInsight) GetBlobPath() string`

GetBlobPath returns the BlobPath field if non-nil, zero value otherwise.

### GetBlobPathOk

`func (o *UnderRepresentationInsight) GetBlobPathOk() (*string, bool)`

GetBlobPathOk returns a tuple with the BlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobPath

`func (o *UnderRepresentationInsight) SetBlobPath(v string)`

SetBlobPath sets BlobPath field to given value.


### GetInsightSubCategoryDsCuration

`func (o *UnderRepresentationInsight) GetInsightSubCategoryDsCuration() InsightSubCategoryDSCuration`

GetInsightSubCategoryDsCuration returns the InsightSubCategoryDsCuration field if non-nil, zero value otherwise.

### GetInsightSubCategoryDsCurationOk

`func (o *UnderRepresentationInsight) GetInsightSubCategoryDsCurationOk() (*InsightSubCategoryDSCuration, bool)`

GetInsightSubCategoryDsCurationOk returns a tuple with the InsightSubCategoryDsCuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightSubCategoryDsCuration

`func (o *UnderRepresentationInsight) SetInsightSubCategoryDsCuration(v InsightSubCategoryDSCuration)`

SetInsightSubCategoryDsCuration sets InsightSubCategoryDsCuration field to given value.

### HasInsightSubCategoryDsCuration

`func (o *UnderRepresentationInsight) HasInsightSubCategoryDsCuration() bool`

HasInsightSubCategoryDsCuration returns a boolean if a field has been set.

### GetInsightCategoryPerformance

`func (o *UnderRepresentationInsight) GetInsightCategoryPerformance() InsightSubCategoryPerformance`

GetInsightCategoryPerformance returns the InsightCategoryPerformance field if non-nil, zero value otherwise.

### GetInsightCategoryPerformanceOk

`func (o *UnderRepresentationInsight) GetInsightCategoryPerformanceOk() (*InsightSubCategoryPerformance, bool)`

GetInsightCategoryPerformanceOk returns a tuple with the InsightCategoryPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightCategoryPerformance

`func (o *UnderRepresentationInsight) SetInsightCategoryPerformance(v InsightSubCategoryPerformance)`

SetInsightCategoryPerformance sets InsightCategoryPerformance field to given value.

### HasInsightCategoryPerformance

`func (o *UnderRepresentationInsight) HasInsightCategoryPerformance() bool`

HasInsightCategoryPerformance returns a boolean if a field has been set.

### GetSeverity

`func (o *UnderRepresentationInsight) GetSeverity() float64`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *UnderRepresentationInsight) GetSeverityOk() (*float64, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *UnderRepresentationInsight) SetSeverity(v float64)`

SetSeverity sets Severity field to given value.


### GetSeverityMetrics

`func (o *UnderRepresentationInsight) GetSeverityMetrics() []SeverityMetricElement`

GetSeverityMetrics returns the SeverityMetrics field if non-nil, zero value otherwise.

### GetSeverityMetricsOk

`func (o *UnderRepresentationInsight) GetSeverityMetricsOk() (*[]SeverityMetricElement, bool)`

GetSeverityMetricsOk returns a tuple with the SeverityMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityMetrics

`func (o *UnderRepresentationInsight) SetSeverityMetrics(v []SeverityMetricElement)`

SetSeverityMetrics sets SeverityMetrics field to given value.


### GetMetricsInfo

`func (o *UnderRepresentationInsight) GetMetricsInfo() []InsightMetricInfo`

GetMetricsInfo returns the MetricsInfo field if non-nil, zero value otherwise.

### GetMetricsInfoOk

`func (o *UnderRepresentationInsight) GetMetricsInfoOk() (*[]InsightMetricInfo, bool)`

GetMetricsInfoOk returns a tuple with the MetricsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsInfo

`func (o *UnderRepresentationInsight) SetMetricsInfo(v []InsightMetricInfo)`

SetMetricsInfo sets MetricsInfo field to given value.


### GetIndex

`func (o *UnderRepresentationInsight) GetIndex() float64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *UnderRepresentationInsight) GetIndexOk() (*float64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *UnderRepresentationInsight) SetIndex(v float64)`

SetIndex sets Index field to given value.


### GetMinHash

`func (o *UnderRepresentationInsight) GetMinHash() []float64`

GetMinHash returns the MinHash field if non-nil, zero value otherwise.

### GetMinHashOk

`func (o *UnderRepresentationInsight) GetMinHashOk() (*[]float64, bool)`

GetMinHashOk returns a tuple with the MinHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHash

`func (o *UnderRepresentationInsight) SetMinHash(v []float64)`

SetMinHash sets MinHash field to given value.


### GetCsvPath

`func (o *UnderRepresentationInsight) GetCsvPath() string`

GetCsvPath returns the CsvPath field if non-nil, zero value otherwise.

### GetCsvPathOk

`func (o *UnderRepresentationInsight) GetCsvPathOk() (*string, bool)`

GetCsvPathOk returns a tuple with the CsvPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvPath

`func (o *UnderRepresentationInsight) SetCsvPath(v string)`

SetCsvPath sets CsvPath field to given value.

### HasCsvPath

`func (o *UnderRepresentationInsight) HasCsvPath() bool`

HasCsvPath returns a boolean if a field has been set.

### GetAutomaticTests

`func (o *UnderRepresentationInsight) GetAutomaticTests() []InsightAutomaticTest`

GetAutomaticTests returns the AutomaticTests field if non-nil, zero value otherwise.

### GetAutomaticTestsOk

`func (o *UnderRepresentationInsight) GetAutomaticTestsOk() (*[]InsightAutomaticTest, bool)`

GetAutomaticTestsOk returns a tuple with the AutomaticTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticTests

`func (o *UnderRepresentationInsight) SetAutomaticTests(v []InsightAutomaticTest)`

SetAutomaticTests sets AutomaticTests field to given value.

### HasAutomaticTests

`func (o *UnderRepresentationInsight) HasAutomaticTests() bool`

HasAutomaticTests returns a boolean if a field has been set.

### GetUnderRepresentationDataset

`func (o *UnderRepresentationInsight) GetUnderRepresentationDataset() DataStateType`

GetUnderRepresentationDataset returns the UnderRepresentationDataset field if non-nil, zero value otherwise.

### GetUnderRepresentationDatasetOk

`func (o *UnderRepresentationInsight) GetUnderRepresentationDatasetOk() (*DataStateType, bool)`

GetUnderRepresentationDatasetOk returns a tuple with the UnderRepresentationDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderRepresentationDataset

`func (o *UnderRepresentationInsight) SetUnderRepresentationDataset(v DataStateType)`

SetUnderRepresentationDataset sets UnderRepresentationDataset field to given value.


### GetUnderRepresentationNSamples

`func (o *UnderRepresentationInsight) GetUnderRepresentationNSamples() float64`

GetUnderRepresentationNSamples returns the UnderRepresentationNSamples field if non-nil, zero value otherwise.

### GetUnderRepresentationNSamplesOk

`func (o *UnderRepresentationInsight) GetUnderRepresentationNSamplesOk() (*float64, bool)`

GetUnderRepresentationNSamplesOk returns a tuple with the UnderRepresentationNSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderRepresentationNSamples

`func (o *UnderRepresentationInsight) SetUnderRepresentationNSamples(v float64)`

SetUnderRepresentationNSamples sets UnderRepresentationNSamples field to given value.


### GetOverRepresentationDataset

`func (o *UnderRepresentationInsight) GetOverRepresentationDataset() DataStateType`

GetOverRepresentationDataset returns the OverRepresentationDataset field if non-nil, zero value otherwise.

### GetOverRepresentationDatasetOk

`func (o *UnderRepresentationInsight) GetOverRepresentationDatasetOk() (*DataStateType, bool)`

GetOverRepresentationDatasetOk returns a tuple with the OverRepresentationDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverRepresentationDataset

`func (o *UnderRepresentationInsight) SetOverRepresentationDataset(v DataStateType)`

SetOverRepresentationDataset sets OverRepresentationDataset field to given value.


### GetOverRepresentationNSamples

`func (o *UnderRepresentationInsight) GetOverRepresentationNSamples() float64`

GetOverRepresentationNSamples returns the OverRepresentationNSamples field if non-nil, zero value otherwise.

### GetOverRepresentationNSamplesOk

`func (o *UnderRepresentationInsight) GetOverRepresentationNSamplesOk() (*float64, bool)`

GetOverRepresentationNSamplesOk returns a tuple with the OverRepresentationNSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverRepresentationNSamples

`func (o *UnderRepresentationInsight) SetOverRepresentationNSamples(v float64)`

SetOverRepresentationNSamples sets OverRepresentationNSamples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


