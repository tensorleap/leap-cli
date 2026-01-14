# MislabeledSamplesInsight

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
**Subset** | [**DataStateType**](DataStateType.md) |  | 

## Methods

### NewMislabeledSamplesInsight

`func NewMislabeledSamplesInsight(type_ ScatterInsightType, nSamples float64, mutualInfoElements []MutualInformationElement, blobPath string, severity float64, severityMetrics []SeverityMetricElement, metricsInfo []InsightMetricInfo, index float64, minHash []float64, subset DataStateType, ) *MislabeledSamplesInsight`

NewMislabeledSamplesInsight instantiates a new MislabeledSamplesInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMislabeledSamplesInsightWithDefaults

`func NewMislabeledSamplesInsightWithDefaults() *MislabeledSamplesInsight`

NewMislabeledSamplesInsightWithDefaults instantiates a new MislabeledSamplesInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MislabeledSamplesInsight) GetType() ScatterInsightType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MislabeledSamplesInsight) GetTypeOk() (*ScatterInsightType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MislabeledSamplesInsight) SetType(v ScatterInsightType)`

SetType sets Type field to given value.


### GetFilters

`func (o *MislabeledSamplesInsight) GetFilters() []ScatterFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MislabeledSamplesInsight) GetFiltersOk() (*[]ScatterFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MislabeledSamplesInsight) SetFilters(v []ScatterFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MislabeledSamplesInsight) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetNSamples

`func (o *MislabeledSamplesInsight) GetNSamples() float64`

GetNSamples returns the NSamples field if non-nil, zero value otherwise.

### GetNSamplesOk

`func (o *MislabeledSamplesInsight) GetNSamplesOk() (*float64, bool)`

GetNSamplesOk returns a tuple with the NSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSamples

`func (o *MislabeledSamplesInsight) SetNSamples(v float64)`

SetNSamples sets NSamples field to given value.


### GetMutualInfoElements

`func (o *MislabeledSamplesInsight) GetMutualInfoElements() []MutualInformationElement`

GetMutualInfoElements returns the MutualInfoElements field if non-nil, zero value otherwise.

### GetMutualInfoElementsOk

`func (o *MislabeledSamplesInsight) GetMutualInfoElementsOk() (*[]MutualInformationElement, bool)`

GetMutualInfoElementsOk returns a tuple with the MutualInfoElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualInfoElements

`func (o *MislabeledSamplesInsight) SetMutualInfoElements(v []MutualInformationElement)`

SetMutualInfoElements sets MutualInfoElements field to given value.


### GetBlobPath

`func (o *MislabeledSamplesInsight) GetBlobPath() string`

GetBlobPath returns the BlobPath field if non-nil, zero value otherwise.

### GetBlobPathOk

`func (o *MislabeledSamplesInsight) GetBlobPathOk() (*string, bool)`

GetBlobPathOk returns a tuple with the BlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobPath

`func (o *MislabeledSamplesInsight) SetBlobPath(v string)`

SetBlobPath sets BlobPath field to given value.


### GetInsightSubCategoryDsCuration

`func (o *MislabeledSamplesInsight) GetInsightSubCategoryDsCuration() InsightSubCategoryDSCuration`

GetInsightSubCategoryDsCuration returns the InsightSubCategoryDsCuration field if non-nil, zero value otherwise.

### GetInsightSubCategoryDsCurationOk

`func (o *MislabeledSamplesInsight) GetInsightSubCategoryDsCurationOk() (*InsightSubCategoryDSCuration, bool)`

GetInsightSubCategoryDsCurationOk returns a tuple with the InsightSubCategoryDsCuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightSubCategoryDsCuration

`func (o *MislabeledSamplesInsight) SetInsightSubCategoryDsCuration(v InsightSubCategoryDSCuration)`

SetInsightSubCategoryDsCuration sets InsightSubCategoryDsCuration field to given value.

### HasInsightSubCategoryDsCuration

`func (o *MislabeledSamplesInsight) HasInsightSubCategoryDsCuration() bool`

HasInsightSubCategoryDsCuration returns a boolean if a field has been set.

### GetInsightCategoryPerformance

`func (o *MislabeledSamplesInsight) GetInsightCategoryPerformance() InsightSubCategoryPerformance`

GetInsightCategoryPerformance returns the InsightCategoryPerformance field if non-nil, zero value otherwise.

### GetInsightCategoryPerformanceOk

`func (o *MislabeledSamplesInsight) GetInsightCategoryPerformanceOk() (*InsightSubCategoryPerformance, bool)`

GetInsightCategoryPerformanceOk returns a tuple with the InsightCategoryPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightCategoryPerformance

`func (o *MislabeledSamplesInsight) SetInsightCategoryPerformance(v InsightSubCategoryPerformance)`

SetInsightCategoryPerformance sets InsightCategoryPerformance field to given value.

### HasInsightCategoryPerformance

`func (o *MislabeledSamplesInsight) HasInsightCategoryPerformance() bool`

HasInsightCategoryPerformance returns a boolean if a field has been set.

### GetSeverity

`func (o *MislabeledSamplesInsight) GetSeverity() float64`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *MislabeledSamplesInsight) GetSeverityOk() (*float64, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *MislabeledSamplesInsight) SetSeverity(v float64)`

SetSeverity sets Severity field to given value.


### GetSeverityMetrics

`func (o *MislabeledSamplesInsight) GetSeverityMetrics() []SeverityMetricElement`

GetSeverityMetrics returns the SeverityMetrics field if non-nil, zero value otherwise.

### GetSeverityMetricsOk

`func (o *MislabeledSamplesInsight) GetSeverityMetricsOk() (*[]SeverityMetricElement, bool)`

GetSeverityMetricsOk returns a tuple with the SeverityMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityMetrics

`func (o *MislabeledSamplesInsight) SetSeverityMetrics(v []SeverityMetricElement)`

SetSeverityMetrics sets SeverityMetrics field to given value.


### GetMetricsInfo

`func (o *MislabeledSamplesInsight) GetMetricsInfo() []InsightMetricInfo`

GetMetricsInfo returns the MetricsInfo field if non-nil, zero value otherwise.

### GetMetricsInfoOk

`func (o *MislabeledSamplesInsight) GetMetricsInfoOk() (*[]InsightMetricInfo, bool)`

GetMetricsInfoOk returns a tuple with the MetricsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsInfo

`func (o *MislabeledSamplesInsight) SetMetricsInfo(v []InsightMetricInfo)`

SetMetricsInfo sets MetricsInfo field to given value.


### GetIndex

`func (o *MislabeledSamplesInsight) GetIndex() float64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MislabeledSamplesInsight) GetIndexOk() (*float64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MislabeledSamplesInsight) SetIndex(v float64)`

SetIndex sets Index field to given value.


### GetMinHash

`func (o *MislabeledSamplesInsight) GetMinHash() []float64`

GetMinHash returns the MinHash field if non-nil, zero value otherwise.

### GetMinHashOk

`func (o *MislabeledSamplesInsight) GetMinHashOk() (*[]float64, bool)`

GetMinHashOk returns a tuple with the MinHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHash

`func (o *MislabeledSamplesInsight) SetMinHash(v []float64)`

SetMinHash sets MinHash field to given value.


### GetCsvPath

`func (o *MislabeledSamplesInsight) GetCsvPath() string`

GetCsvPath returns the CsvPath field if non-nil, zero value otherwise.

### GetCsvPathOk

`func (o *MislabeledSamplesInsight) GetCsvPathOk() (*string, bool)`

GetCsvPathOk returns a tuple with the CsvPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvPath

`func (o *MislabeledSamplesInsight) SetCsvPath(v string)`

SetCsvPath sets CsvPath field to given value.

### HasCsvPath

`func (o *MislabeledSamplesInsight) HasCsvPath() bool`

HasCsvPath returns a boolean if a field has been set.

### GetAutomaticTests

`func (o *MislabeledSamplesInsight) GetAutomaticTests() []InsightAutomaticTest`

GetAutomaticTests returns the AutomaticTests field if non-nil, zero value otherwise.

### GetAutomaticTestsOk

`func (o *MislabeledSamplesInsight) GetAutomaticTestsOk() (*[]InsightAutomaticTest, bool)`

GetAutomaticTestsOk returns a tuple with the AutomaticTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticTests

`func (o *MislabeledSamplesInsight) SetAutomaticTests(v []InsightAutomaticTest)`

SetAutomaticTests sets AutomaticTests field to given value.

### HasAutomaticTests

`func (o *MislabeledSamplesInsight) HasAutomaticTests() bool`

HasAutomaticTests returns a boolean if a field has been set.

### GetSubset

`func (o *MislabeledSamplesInsight) GetSubset() DataStateType`

GetSubset returns the Subset field if non-nil, zero value otherwise.

### GetSubsetOk

`func (o *MislabeledSamplesInsight) GetSubsetOk() (*DataStateType, bool)`

GetSubsetOk returns a tuple with the Subset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubset

`func (o *MislabeledSamplesInsight) SetSubset(v DataStateType)`

SetSubset sets Subset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


