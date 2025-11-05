# DomainGapInsight

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
**MetadataName** | **string** |  | 
**DomainGapScore** | **float64** |  | 
**DomainA** | **string** |  | 
**DomainB** | **string** |  | 

## Methods

### NewDomainGapInsight

`func NewDomainGapInsight(type_ ScatterInsightType, nSamples float64, mutualInfoElements []MutualInformationElement, blobPath string, severity float64, severityMetrics []SeverityMetricElement, metricsInfo []InsightMetricInfo, index float64, minHash []float64, metadataName string, domainGapScore float64, domainA string, domainB string, ) *DomainGapInsight`

NewDomainGapInsight instantiates a new DomainGapInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainGapInsightWithDefaults

`func NewDomainGapInsightWithDefaults() *DomainGapInsight`

NewDomainGapInsightWithDefaults instantiates a new DomainGapInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DomainGapInsight) GetType() ScatterInsightType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainGapInsight) GetTypeOk() (*ScatterInsightType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainGapInsight) SetType(v ScatterInsightType)`

SetType sets Type field to given value.


### GetFilters

`func (o *DomainGapInsight) GetFilters() []ScatterFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DomainGapInsight) GetFiltersOk() (*[]ScatterFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DomainGapInsight) SetFilters(v []ScatterFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DomainGapInsight) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetNSamples

`func (o *DomainGapInsight) GetNSamples() float64`

GetNSamples returns the NSamples field if non-nil, zero value otherwise.

### GetNSamplesOk

`func (o *DomainGapInsight) GetNSamplesOk() (*float64, bool)`

GetNSamplesOk returns a tuple with the NSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSamples

`func (o *DomainGapInsight) SetNSamples(v float64)`

SetNSamples sets NSamples field to given value.


### GetMutualInfoElements

`func (o *DomainGapInsight) GetMutualInfoElements() []MutualInformationElement`

GetMutualInfoElements returns the MutualInfoElements field if non-nil, zero value otherwise.

### GetMutualInfoElementsOk

`func (o *DomainGapInsight) GetMutualInfoElementsOk() (*[]MutualInformationElement, bool)`

GetMutualInfoElementsOk returns a tuple with the MutualInfoElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualInfoElements

`func (o *DomainGapInsight) SetMutualInfoElements(v []MutualInformationElement)`

SetMutualInfoElements sets MutualInfoElements field to given value.


### GetBlobPath

`func (o *DomainGapInsight) GetBlobPath() string`

GetBlobPath returns the BlobPath field if non-nil, zero value otherwise.

### GetBlobPathOk

`func (o *DomainGapInsight) GetBlobPathOk() (*string, bool)`

GetBlobPathOk returns a tuple with the BlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobPath

`func (o *DomainGapInsight) SetBlobPath(v string)`

SetBlobPath sets BlobPath field to given value.


### GetInsightSubCategoryDsCuration

`func (o *DomainGapInsight) GetInsightSubCategoryDsCuration() InsightSubCategoryDSCuration`

GetInsightSubCategoryDsCuration returns the InsightSubCategoryDsCuration field if non-nil, zero value otherwise.

### GetInsightSubCategoryDsCurationOk

`func (o *DomainGapInsight) GetInsightSubCategoryDsCurationOk() (*InsightSubCategoryDSCuration, bool)`

GetInsightSubCategoryDsCurationOk returns a tuple with the InsightSubCategoryDsCuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightSubCategoryDsCuration

`func (o *DomainGapInsight) SetInsightSubCategoryDsCuration(v InsightSubCategoryDSCuration)`

SetInsightSubCategoryDsCuration sets InsightSubCategoryDsCuration field to given value.

### HasInsightSubCategoryDsCuration

`func (o *DomainGapInsight) HasInsightSubCategoryDsCuration() bool`

HasInsightSubCategoryDsCuration returns a boolean if a field has been set.

### GetInsightCategoryPerformance

`func (o *DomainGapInsight) GetInsightCategoryPerformance() InsightSubCategoryPerformance`

GetInsightCategoryPerformance returns the InsightCategoryPerformance field if non-nil, zero value otherwise.

### GetInsightCategoryPerformanceOk

`func (o *DomainGapInsight) GetInsightCategoryPerformanceOk() (*InsightSubCategoryPerformance, bool)`

GetInsightCategoryPerformanceOk returns a tuple with the InsightCategoryPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightCategoryPerformance

`func (o *DomainGapInsight) SetInsightCategoryPerformance(v InsightSubCategoryPerformance)`

SetInsightCategoryPerformance sets InsightCategoryPerformance field to given value.

### HasInsightCategoryPerformance

`func (o *DomainGapInsight) HasInsightCategoryPerformance() bool`

HasInsightCategoryPerformance returns a boolean if a field has been set.

### GetSeverity

`func (o *DomainGapInsight) GetSeverity() float64`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *DomainGapInsight) GetSeverityOk() (*float64, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *DomainGapInsight) SetSeverity(v float64)`

SetSeverity sets Severity field to given value.


### GetSeverityMetrics

`func (o *DomainGapInsight) GetSeverityMetrics() []SeverityMetricElement`

GetSeverityMetrics returns the SeverityMetrics field if non-nil, zero value otherwise.

### GetSeverityMetricsOk

`func (o *DomainGapInsight) GetSeverityMetricsOk() (*[]SeverityMetricElement, bool)`

GetSeverityMetricsOk returns a tuple with the SeverityMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityMetrics

`func (o *DomainGapInsight) SetSeverityMetrics(v []SeverityMetricElement)`

SetSeverityMetrics sets SeverityMetrics field to given value.


### GetMetricsInfo

`func (o *DomainGapInsight) GetMetricsInfo() []InsightMetricInfo`

GetMetricsInfo returns the MetricsInfo field if non-nil, zero value otherwise.

### GetMetricsInfoOk

`func (o *DomainGapInsight) GetMetricsInfoOk() (*[]InsightMetricInfo, bool)`

GetMetricsInfoOk returns a tuple with the MetricsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsInfo

`func (o *DomainGapInsight) SetMetricsInfo(v []InsightMetricInfo)`

SetMetricsInfo sets MetricsInfo field to given value.


### GetIndex

`func (o *DomainGapInsight) GetIndex() float64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DomainGapInsight) GetIndexOk() (*float64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DomainGapInsight) SetIndex(v float64)`

SetIndex sets Index field to given value.


### GetMinHash

`func (o *DomainGapInsight) GetMinHash() []float64`

GetMinHash returns the MinHash field if non-nil, zero value otherwise.

### GetMinHashOk

`func (o *DomainGapInsight) GetMinHashOk() (*[]float64, bool)`

GetMinHashOk returns a tuple with the MinHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHash

`func (o *DomainGapInsight) SetMinHash(v []float64)`

SetMinHash sets MinHash field to given value.


### GetCsvPath

`func (o *DomainGapInsight) GetCsvPath() string`

GetCsvPath returns the CsvPath field if non-nil, zero value otherwise.

### GetCsvPathOk

`func (o *DomainGapInsight) GetCsvPathOk() (*string, bool)`

GetCsvPathOk returns a tuple with the CsvPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvPath

`func (o *DomainGapInsight) SetCsvPath(v string)`

SetCsvPath sets CsvPath field to given value.

### HasCsvPath

`func (o *DomainGapInsight) HasCsvPath() bool`

HasCsvPath returns a boolean if a field has been set.

### GetMetadataName

`func (o *DomainGapInsight) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *DomainGapInsight) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *DomainGapInsight) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.


### GetDomainGapScore

`func (o *DomainGapInsight) GetDomainGapScore() float64`

GetDomainGapScore returns the DomainGapScore field if non-nil, zero value otherwise.

### GetDomainGapScoreOk

`func (o *DomainGapInsight) GetDomainGapScoreOk() (*float64, bool)`

GetDomainGapScoreOk returns a tuple with the DomainGapScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGapScore

`func (o *DomainGapInsight) SetDomainGapScore(v float64)`

SetDomainGapScore sets DomainGapScore field to given value.


### GetDomainA

`func (o *DomainGapInsight) GetDomainA() string`

GetDomainA returns the DomainA field if non-nil, zero value otherwise.

### GetDomainAOk

`func (o *DomainGapInsight) GetDomainAOk() (*string, bool)`

GetDomainAOk returns a tuple with the DomainA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainA

`func (o *DomainGapInsight) SetDomainA(v string)`

SetDomainA sets DomainA field to given value.


### GetDomainB

`func (o *DomainGapInsight) GetDomainB() string`

GetDomainB returns the DomainB field if non-nil, zero value otherwise.

### GetDomainBOk

`func (o *DomainGapInsight) GetDomainBOk() (*string, bool)`

GetDomainBOk returns a tuple with the DomainB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainB

`func (o *DomainGapInsight) SetDomainB(v string)`

SetDomainB sets DomainB field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


