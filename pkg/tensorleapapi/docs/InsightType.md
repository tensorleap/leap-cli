# InsightType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ScatterInsightType**](ScatterInsightType.md) |  | 
**Filter** | [**ScatterFilter**](ScatterFilter.md) |  | 
**MutualInfoElements** | Pointer to [**[]MutualInformationElement**](MutualInformationElement.md) |  | [optional] 
**BlobPath** | Pointer to **string** |  | [optional] 
**InsightSubCategoryDsCuration** | Pointer to [**InsightSubCategoryDSCuration**](InsightSubCategoryDSCuration.md) |  | [optional] 
**InsightCategoryPerformance** | Pointer to [**InsightSubCategoryPerformance**](InsightSubCategoryPerformance.md) |  | [optional] 
**Index** | Pointer to **float64** |  | [optional] 
**NSamplesValidation** | **float64** |  | 
**NSamplesTraining** | **float64** |  | 
**AvgMetricValidation** | **float64** |  | 
**AvgMetricTraining** | **float64** |  | 
**MetricName** | **string** |  | 
**Configuration** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 
**MetricsNames** | Pointer to **[]string** |  | [optional] 
**UnderRepresentationDataset** | [**DataStateType**](DataStateType.md) |  | 
**UnderRepresentationNSamples** | **float64** |  | 
**OverRepresentationDataset** | [**DataStateType**](DataStateType.md) |  | 
**OverRepresentationNSamples** | **float64** |  | 
**NSamples** | **float64** |  | 
**AvgMetric** | **float64** |  | 
**Subset** | [**DataStateType**](DataStateType.md) |  | 
**NDuplicateSamples** | **float64** |  | 
**FirstSubset** | [**DataStateType**](DataStateType.md) |  | 
**SecondSubset** | [**DataStateType**](DataStateType.md) |  | 
**NCrossSubsetsDuplications** | **float64** |  | 

## Methods

### NewInsightType

`func NewInsightType(type_ ScatterInsightType, filter ScatterFilter, nSamplesValidation float64, nSamplesTraining float64, avgMetricValidation float64, avgMetricTraining float64, metricName string, underRepresentationDataset DataStateType, underRepresentationNSamples float64, overRepresentationDataset DataStateType, overRepresentationNSamples float64, nSamples float64, avgMetric float64, subset DataStateType, nDuplicateSamples float64, firstSubset DataStateType, secondSubset DataStateType, nCrossSubsetsDuplications float64, ) *InsightType`

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


### GetFilter

`func (o *InsightType) GetFilter() ScatterFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InsightType) GetFilterOk() (*ScatterFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InsightType) SetFilter(v ScatterFilter)`

SetFilter sets Filter field to given value.


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

### HasMutualInfoElements

`func (o *InsightType) HasMutualInfoElements() bool`

HasMutualInfoElements returns a boolean if a field has been set.

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

### HasBlobPath

`func (o *InsightType) HasBlobPath() bool`

HasBlobPath returns a boolean if a field has been set.

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

### HasIndex

`func (o *InsightType) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

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


### GetAvgMetricValidation

`func (o *InsightType) GetAvgMetricValidation() float64`

GetAvgMetricValidation returns the AvgMetricValidation field if non-nil, zero value otherwise.

### GetAvgMetricValidationOk

`func (o *InsightType) GetAvgMetricValidationOk() (*float64, bool)`

GetAvgMetricValidationOk returns a tuple with the AvgMetricValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgMetricValidation

`func (o *InsightType) SetAvgMetricValidation(v float64)`

SetAvgMetricValidation sets AvgMetricValidation field to given value.


### GetAvgMetricTraining

`func (o *InsightType) GetAvgMetricTraining() float64`

GetAvgMetricTraining returns the AvgMetricTraining field if non-nil, zero value otherwise.

### GetAvgMetricTrainingOk

`func (o *InsightType) GetAvgMetricTrainingOk() (*float64, bool)`

GetAvgMetricTrainingOk returns a tuple with the AvgMetricTraining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgMetricTraining

`func (o *InsightType) SetAvgMetricTraining(v float64)`

SetAvgMetricTraining sets AvgMetricTraining field to given value.


### GetMetricName

`func (o *InsightType) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *InsightType) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *InsightType) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetConfiguration

`func (o *InsightType) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *InsightType) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *InsightType) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *InsightType) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetMetricsNames

`func (o *InsightType) GetMetricsNames() []string`

GetMetricsNames returns the MetricsNames field if non-nil, zero value otherwise.

### GetMetricsNamesOk

`func (o *InsightType) GetMetricsNamesOk() (*[]string, bool)`

GetMetricsNamesOk returns a tuple with the MetricsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsNames

`func (o *InsightType) SetMetricsNames(v []string)`

SetMetricsNames sets MetricsNames field to given value.

### HasMetricsNames

`func (o *InsightType) HasMetricsNames() bool`

HasMetricsNames returns a boolean if a field has been set.

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


### GetAvgMetric

`func (o *InsightType) GetAvgMetric() float64`

GetAvgMetric returns the AvgMetric field if non-nil, zero value otherwise.

### GetAvgMetricOk

`func (o *InsightType) GetAvgMetricOk() (*float64, bool)`

GetAvgMetricOk returns a tuple with the AvgMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgMetric

`func (o *InsightType) SetAvgMetric(v float64)`

SetAvgMetric sets AvgMetric field to given value.


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


### GetNDuplicateSamples

`func (o *InsightType) GetNDuplicateSamples() float64`

GetNDuplicateSamples returns the NDuplicateSamples field if non-nil, zero value otherwise.

### GetNDuplicateSamplesOk

`func (o *InsightType) GetNDuplicateSamplesOk() (*float64, bool)`

GetNDuplicateSamplesOk returns a tuple with the NDuplicateSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDuplicateSamples

`func (o *InsightType) SetNDuplicateSamples(v float64)`

SetNDuplicateSamples sets NDuplicateSamples field to given value.


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


### GetNCrossSubsetsDuplications

`func (o *InsightType) GetNCrossSubsetsDuplications() float64`

GetNCrossSubsetsDuplications returns the NCrossSubsetsDuplications field if non-nil, zero value otherwise.

### GetNCrossSubsetsDuplicationsOk

`func (o *InsightType) GetNCrossSubsetsDuplicationsOk() (*float64, bool)`

GetNCrossSubsetsDuplicationsOk returns a tuple with the NCrossSubsetsDuplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCrossSubsetsDuplications

`func (o *InsightType) SetNCrossSubsetsDuplications(v float64)`

SetNCrossSubsetsDuplications sets NCrossSubsetsDuplications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


