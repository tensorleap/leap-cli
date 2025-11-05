# InsightType

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
**NSamplesValidation** | **float64** |  | 
**NSamplesTraining** | **float64** |  | 
**UnderRepresentationDataset** | [**DataStateType**](DataStateType.md) |  | 
**UnderRepresentationNSamples** | **float64** |  | 
**OverRepresentationDataset** | [**DataStateType**](DataStateType.md) |  | 
**OverRepresentationNSamples** | **float64** |  | 
**Subset** | [**DataStateType**](DataStateType.md) |  | 
**FirstSubset** | [**DataStateType**](DataStateType.md) |  | 
**SecondSubset** | [**DataStateType**](DataStateType.md) |  | 
**MetadataName** | **string** |  | 
**DomainGapScore** | **float64** |  | 
**DomainA** | **string** |  | 
**DomainB** | **string** |  | 

## Methods

### NewInsightType

`func NewInsightType(type_ ScatterInsightType, nSamples float64, mutualInfoElements []MutualInformationElement, blobPath string, severity float64, severityMetrics []SeverityMetricElement, metricsInfo []InsightMetricInfo, index float64, minHash []float64, nSamplesValidation float64, nSamplesTraining float64, underRepresentationDataset DataStateType, underRepresentationNSamples float64, overRepresentationDataset DataStateType, overRepresentationNSamples float64, subset DataStateType, firstSubset DataStateType, secondSubset DataStateType, metadataName string, domainGapScore float64, domainA string, domainB string, ) *InsightType`

NewInsightType instantiates a new InsightType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightTypeWithDefaults

`func NewInsightTypeWithDefaults() *InsightType`

NewInsightTypeWithDefaults instantiates a new InsightType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InsightType) GetType() ScatterInsightType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InsightType) GetTypeOk() (*ScatterInsightType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InsightType) SetType(v ScatterInsightType)`

SetType sets Type field to given value.


### GetFilters

`func (o *InsightType) GetFilters() []ScatterFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *InsightType) GetFiltersOk() (*[]ScatterFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *InsightType) SetFilters(v []ScatterFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *InsightType) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetNSamples

`func (o *InsightType) GetNSamples() float64`

GetNSamples returns the NSamples field if non-nil, zero value otherwise.

### GetNSamplesOk

`func (o *InsightType) GetNSamplesOk() (*float64, bool)`

GetNSamplesOk returns a tuple with the NSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSamples

`func (o *InsightType) SetNSamples(v float64)`

SetNSamples sets NSamples field to given value.


### GetMutualInfoElements

`func (o *InsightType) GetMutualInfoElements() []MutualInformationElement`

GetMutualInfoElements returns the MutualInfoElements field if non-nil, zero value otherwise.

### GetMutualInfoElementsOk

`func (o *InsightType) GetMutualInfoElementsOk() (*[]MutualInformationElement, bool)`

GetMutualInfoElementsOk returns a tuple with the MutualInfoElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualInfoElements

`func (o *InsightType) SetMutualInfoElements(v []MutualInformationElement)`

SetMutualInfoElements sets MutualInfoElements field to given value.


### GetBlobPath

`func (o *InsightType) GetBlobPath() string`

GetBlobPath returns the BlobPath field if non-nil, zero value otherwise.

### GetBlobPathOk

`func (o *InsightType) GetBlobPathOk() (*string, bool)`

GetBlobPathOk returns a tuple with the BlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobPath

`func (o *InsightType) SetBlobPath(v string)`

SetBlobPath sets BlobPath field to given value.


### GetInsightSubCategoryDsCuration

`func (o *InsightType) GetInsightSubCategoryDsCuration() InsightSubCategoryDSCuration`

GetInsightSubCategoryDsCuration returns the InsightSubCategoryDsCuration field if non-nil, zero value otherwise.

### GetInsightSubCategoryDsCurationOk

`func (o *InsightType) GetInsightSubCategoryDsCurationOk() (*InsightSubCategoryDSCuration, bool)`

GetInsightSubCategoryDsCurationOk returns a tuple with the InsightSubCategoryDsCuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightSubCategoryDsCuration

`func (o *InsightType) SetInsightSubCategoryDsCuration(v InsightSubCategoryDSCuration)`

SetInsightSubCategoryDsCuration sets InsightSubCategoryDsCuration field to given value.

### HasInsightSubCategoryDsCuration

`func (o *InsightType) HasInsightSubCategoryDsCuration() bool`

HasInsightSubCategoryDsCuration returns a boolean if a field has been set.

### GetInsightCategoryPerformance

`func (o *InsightType) GetInsightCategoryPerformance() InsightSubCategoryPerformance`

GetInsightCategoryPerformance returns the InsightCategoryPerformance field if non-nil, zero value otherwise.

### GetInsightCategoryPerformanceOk

`func (o *InsightType) GetInsightCategoryPerformanceOk() (*InsightSubCategoryPerformance, bool)`

GetInsightCategoryPerformanceOk returns a tuple with the InsightCategoryPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightCategoryPerformance

`func (o *InsightType) SetInsightCategoryPerformance(v InsightSubCategoryPerformance)`

SetInsightCategoryPerformance sets InsightCategoryPerformance field to given value.

### HasInsightCategoryPerformance

`func (o *InsightType) HasInsightCategoryPerformance() bool`

HasInsightCategoryPerformance returns a boolean if a field has been set.

### GetSeverity

`func (o *InsightType) GetSeverity() float64`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InsightType) GetSeverityOk() (*float64, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InsightType) SetSeverity(v float64)`

SetSeverity sets Severity field to given value.


### GetSeverityMetrics

`func (o *InsightType) GetSeverityMetrics() []SeverityMetricElement`

GetSeverityMetrics returns the SeverityMetrics field if non-nil, zero value otherwise.

### GetSeverityMetricsOk

`func (o *InsightType) GetSeverityMetricsOk() (*[]SeverityMetricElement, bool)`

GetSeverityMetricsOk returns a tuple with the SeverityMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityMetrics

`func (o *InsightType) SetSeverityMetrics(v []SeverityMetricElement)`

SetSeverityMetrics sets SeverityMetrics field to given value.


### GetMetricsInfo

`func (o *InsightType) GetMetricsInfo() []InsightMetricInfo`

GetMetricsInfo returns the MetricsInfo field if non-nil, zero value otherwise.

### GetMetricsInfoOk

`func (o *InsightType) GetMetricsInfoOk() (*[]InsightMetricInfo, bool)`

GetMetricsInfoOk returns a tuple with the MetricsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsInfo

`func (o *InsightType) SetMetricsInfo(v []InsightMetricInfo)`

SetMetricsInfo sets MetricsInfo field to given value.


### GetIndex

`func (o *InsightType) GetIndex() float64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *InsightType) GetIndexOk() (*float64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *InsightType) SetIndex(v float64)`

SetIndex sets Index field to given value.


### GetMinHash

`func (o *InsightType) GetMinHash() []float64`

GetMinHash returns the MinHash field if non-nil, zero value otherwise.

### GetMinHashOk

`func (o *InsightType) GetMinHashOk() (*[]float64, bool)`

GetMinHashOk returns a tuple with the MinHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHash

`func (o *InsightType) SetMinHash(v []float64)`

SetMinHash sets MinHash field to given value.


### GetCsvPath

`func (o *InsightType) GetCsvPath() string`

GetCsvPath returns the CsvPath field if non-nil, zero value otherwise.

### GetCsvPathOk

`func (o *InsightType) GetCsvPathOk() (*string, bool)`

GetCsvPathOk returns a tuple with the CsvPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvPath

`func (o *InsightType) SetCsvPath(v string)`

SetCsvPath sets CsvPath field to given value.

### HasCsvPath

`func (o *InsightType) HasCsvPath() bool`

HasCsvPath returns a boolean if a field has been set.

### GetNSamplesValidation

`func (o *InsightType) GetNSamplesValidation() float64`

GetNSamplesValidation returns the NSamplesValidation field if non-nil, zero value otherwise.

### GetNSamplesValidationOk

`func (o *InsightType) GetNSamplesValidationOk() (*float64, bool)`

GetNSamplesValidationOk returns a tuple with the NSamplesValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSamplesValidation

`func (o *InsightType) SetNSamplesValidation(v float64)`

SetNSamplesValidation sets NSamplesValidation field to given value.


### GetNSamplesTraining

`func (o *InsightType) GetNSamplesTraining() float64`

GetNSamplesTraining returns the NSamplesTraining field if non-nil, zero value otherwise.

### GetNSamplesTrainingOk

`func (o *InsightType) GetNSamplesTrainingOk() (*float64, bool)`

GetNSamplesTrainingOk returns a tuple with the NSamplesTraining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSamplesTraining

`func (o *InsightType) SetNSamplesTraining(v float64)`

SetNSamplesTraining sets NSamplesTraining field to given value.


### GetUnderRepresentationDataset

`func (o *InsightType) GetUnderRepresentationDataset() DataStateType`

GetUnderRepresentationDataset returns the UnderRepresentationDataset field if non-nil, zero value otherwise.

### GetUnderRepresentationDatasetOk

`func (o *InsightType) GetUnderRepresentationDatasetOk() (*DataStateType, bool)`

GetUnderRepresentationDatasetOk returns a tuple with the UnderRepresentationDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderRepresentationDataset

`func (o *InsightType) SetUnderRepresentationDataset(v DataStateType)`

SetUnderRepresentationDataset sets UnderRepresentationDataset field to given value.


### GetUnderRepresentationNSamples

`func (o *InsightType) GetUnderRepresentationNSamples() float64`

GetUnderRepresentationNSamples returns the UnderRepresentationNSamples field if non-nil, zero value otherwise.

### GetUnderRepresentationNSamplesOk

`func (o *InsightType) GetUnderRepresentationNSamplesOk() (*float64, bool)`

GetUnderRepresentationNSamplesOk returns a tuple with the UnderRepresentationNSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderRepresentationNSamples

`func (o *InsightType) SetUnderRepresentationNSamples(v float64)`

SetUnderRepresentationNSamples sets UnderRepresentationNSamples field to given value.


### GetOverRepresentationDataset

`func (o *InsightType) GetOverRepresentationDataset() DataStateType`

GetOverRepresentationDataset returns the OverRepresentationDataset field if non-nil, zero value otherwise.

### GetOverRepresentationDatasetOk

`func (o *InsightType) GetOverRepresentationDatasetOk() (*DataStateType, bool)`

GetOverRepresentationDatasetOk returns a tuple with the OverRepresentationDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverRepresentationDataset

`func (o *InsightType) SetOverRepresentationDataset(v DataStateType)`

SetOverRepresentationDataset sets OverRepresentationDataset field to given value.


### GetOverRepresentationNSamples

`func (o *InsightType) GetOverRepresentationNSamples() float64`

GetOverRepresentationNSamples returns the OverRepresentationNSamples field if non-nil, zero value otherwise.

### GetOverRepresentationNSamplesOk

`func (o *InsightType) GetOverRepresentationNSamplesOk() (*float64, bool)`

GetOverRepresentationNSamplesOk returns a tuple with the OverRepresentationNSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverRepresentationNSamples

`func (o *InsightType) SetOverRepresentationNSamples(v float64)`

SetOverRepresentationNSamples sets OverRepresentationNSamples field to given value.


### GetSubset

`func (o *InsightType) GetSubset() DataStateType`

GetSubset returns the Subset field if non-nil, zero value otherwise.

### GetSubsetOk

`func (o *InsightType) GetSubsetOk() (*DataStateType, bool)`

GetSubsetOk returns a tuple with the Subset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubset

`func (o *InsightType) SetSubset(v DataStateType)`

SetSubset sets Subset field to given value.


### GetFirstSubset

`func (o *InsightType) GetFirstSubset() DataStateType`

GetFirstSubset returns the FirstSubset field if non-nil, zero value otherwise.

### GetFirstSubsetOk

`func (o *InsightType) GetFirstSubsetOk() (*DataStateType, bool)`

GetFirstSubsetOk returns a tuple with the FirstSubset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSubset

`func (o *InsightType) SetFirstSubset(v DataStateType)`

SetFirstSubset sets FirstSubset field to given value.


### GetSecondSubset

`func (o *InsightType) GetSecondSubset() DataStateType`

GetSecondSubset returns the SecondSubset field if non-nil, zero value otherwise.

### GetSecondSubsetOk

`func (o *InsightType) GetSecondSubsetOk() (*DataStateType, bool)`

GetSecondSubsetOk returns a tuple with the SecondSubset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondSubset

`func (o *InsightType) SetSecondSubset(v DataStateType)`

SetSecondSubset sets SecondSubset field to given value.


### GetMetadataName

`func (o *InsightType) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *InsightType) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *InsightType) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.


### GetDomainGapScore

`func (o *InsightType) GetDomainGapScore() float64`

GetDomainGapScore returns the DomainGapScore field if non-nil, zero value otherwise.

### GetDomainGapScoreOk

`func (o *InsightType) GetDomainGapScoreOk() (*float64, bool)`

GetDomainGapScoreOk returns a tuple with the DomainGapScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGapScore

`func (o *InsightType) SetDomainGapScore(v float64)`

SetDomainGapScore sets DomainGapScore field to given value.


### GetDomainA

`func (o *InsightType) GetDomainA() string`

GetDomainA returns the DomainA field if non-nil, zero value otherwise.

### GetDomainAOk

`func (o *InsightType) GetDomainAOk() (*string, bool)`

GetDomainAOk returns a tuple with the DomainA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainA

`func (o *InsightType) SetDomainA(v string)`

SetDomainA sets DomainA field to given value.


### GetDomainB

`func (o *InsightType) GetDomainB() string`

GetDomainB returns the DomainB field if non-nil, zero value otherwise.

### GetDomainBOk

`func (o *InsightType) GetDomainBOk() (*string, bool)`

GetDomainBOk returns a tuple with the DomainB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainB

`func (o *InsightType) SetDomainB(v string)`

SetDomainB sets DomainB field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


