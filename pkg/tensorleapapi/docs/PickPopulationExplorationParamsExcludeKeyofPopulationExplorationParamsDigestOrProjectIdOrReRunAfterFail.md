# PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionRunId** | **string** |  | 
**BatchSize** | **float64** |  | 
**FromEpoch** | **float64** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**TimeFilter** | Pointer to [**ESFilter**](ESFilter.md) |  | [optional] 
**NotApplyTimeFilterOnUnlabeledOnly** | Pointer to **bool** |  | [optional] 
**ProjectionMetric** | Pointer to **string** |  | [optional] 
**DomainGapMetadata** | Pointer to **string** |  | [optional] 
**NumOfSamples** | **float64** |  | 
**BalanceBy** | **[]string** |  | 
**ShouldFillRemainingWithUnbalanced** | **bool** |  | 
**ReductionAlgorithm** | [**ReductionAlgorithm**](ReductionAlgorithm.md) |  | 
**OptionalAnalysis** | Pointer to [**[]OptionalAnalysis**](OptionalAnalysis.md) |  | [optional] 
**TriggerCreateSampleVisualiztions** | Pointer to **bool** |  | [optional] 
**UseCustomLatentSpace** | Pointer to **bool** |  | [optional] 
**ElementInstance** | Pointer to **bool** |  | [optional] 
**ForceExecute** | Pointer to **bool** |  | [optional] 
**LatentSpaceType** | Pointer to **string** |  | [optional] 

## Methods

### NewPickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail

`func NewPickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail(sessionRunId string, batchSize float64, fromEpoch float64, numOfSamples float64, balanceBy []string, shouldFillRemainingWithUnbalanced bool, reductionAlgorithm ReductionAlgorithm, ) *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail`

NewPickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail instantiates a new PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFailWithDefaults

`func NewPickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFailWithDefaults() *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail`

NewPickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFailWithDefaults instantiates a new PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionRunId

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetBatchSize

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetFromEpoch

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetFromEpoch() float64`

GetFromEpoch returns the FromEpoch field if non-nil, zero value otherwise.

### GetFromEpochOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetFromEpochOk() (*float64, bool)`

GetFromEpochOk returns a tuple with the FromEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEpoch

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetFromEpoch(v float64)`

SetFromEpoch sets FromEpoch field to given value.


### GetFilters

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTimeFilter

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetTimeFilter() ESFilter`

GetTimeFilter returns the TimeFilter field if non-nil, zero value otherwise.

### GetTimeFilterOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetTimeFilterOk() (*ESFilter, bool)`

GetTimeFilterOk returns a tuple with the TimeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFilter

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetTimeFilter(v ESFilter)`

SetTimeFilter sets TimeFilter field to given value.

### HasTimeFilter

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) HasTimeFilter() bool`

HasTimeFilter returns a boolean if a field has been set.

### GetNotApplyTimeFilterOnUnlabeledOnly

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetNotApplyTimeFilterOnUnlabeledOnly() bool`

GetNotApplyTimeFilterOnUnlabeledOnly returns the NotApplyTimeFilterOnUnlabeledOnly field if non-nil, zero value otherwise.

### GetNotApplyTimeFilterOnUnlabeledOnlyOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetNotApplyTimeFilterOnUnlabeledOnlyOk() (*bool, bool)`

GetNotApplyTimeFilterOnUnlabeledOnlyOk returns a tuple with the NotApplyTimeFilterOnUnlabeledOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplyTimeFilterOnUnlabeledOnly

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetNotApplyTimeFilterOnUnlabeledOnly(v bool)`

SetNotApplyTimeFilterOnUnlabeledOnly sets NotApplyTimeFilterOnUnlabeledOnly field to given value.

### HasNotApplyTimeFilterOnUnlabeledOnly

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) HasNotApplyTimeFilterOnUnlabeledOnly() bool`

HasNotApplyTimeFilterOnUnlabeledOnly returns a boolean if a field has been set.

### GetProjectionMetric

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetProjectionMetric() string`

GetProjectionMetric returns the ProjectionMetric field if non-nil, zero value otherwise.

### GetProjectionMetricOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetProjectionMetricOk() (*string, bool)`

GetProjectionMetricOk returns a tuple with the ProjectionMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectionMetric

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetProjectionMetric(v string)`

SetProjectionMetric sets ProjectionMetric field to given value.

### HasProjectionMetric

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) HasProjectionMetric() bool`

HasProjectionMetric returns a boolean if a field has been set.

### GetDomainGapMetadata

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetDomainGapMetadata() string`

GetDomainGapMetadata returns the DomainGapMetadata field if non-nil, zero value otherwise.

### GetDomainGapMetadataOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetDomainGapMetadataOk() (*string, bool)`

GetDomainGapMetadataOk returns a tuple with the DomainGapMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGapMetadata

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetDomainGapMetadata(v string)`

SetDomainGapMetadata sets DomainGapMetadata field to given value.

### HasDomainGapMetadata

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) HasDomainGapMetadata() bool`

HasDomainGapMetadata returns a boolean if a field has been set.

### GetNumOfSamples

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetNumOfSamples() float64`

GetNumOfSamples returns the NumOfSamples field if non-nil, zero value otherwise.

### GetNumOfSamplesOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetNumOfSamplesOk() (*float64, bool)`

GetNumOfSamplesOk returns a tuple with the NumOfSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfSamples

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetNumOfSamples(v float64)`

SetNumOfSamples sets NumOfSamples field to given value.


### GetBalanceBy

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetBalanceBy() []string`

GetBalanceBy returns the BalanceBy field if non-nil, zero value otherwise.

### GetBalanceByOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetBalanceByOk() (*[]string, bool)`

GetBalanceByOk returns a tuple with the BalanceBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceBy

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetBalanceBy(v []string)`

SetBalanceBy sets BalanceBy field to given value.


### GetShouldFillRemainingWithUnbalanced

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetShouldFillRemainingWithUnbalanced() bool`

GetShouldFillRemainingWithUnbalanced returns the ShouldFillRemainingWithUnbalanced field if non-nil, zero value otherwise.

### GetShouldFillRemainingWithUnbalancedOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetShouldFillRemainingWithUnbalancedOk() (*bool, bool)`

GetShouldFillRemainingWithUnbalancedOk returns a tuple with the ShouldFillRemainingWithUnbalanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldFillRemainingWithUnbalanced

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetShouldFillRemainingWithUnbalanced(v bool)`

SetShouldFillRemainingWithUnbalanced sets ShouldFillRemainingWithUnbalanced field to given value.


### GetReductionAlgorithm

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetReductionAlgorithm() ReductionAlgorithm`

GetReductionAlgorithm returns the ReductionAlgorithm field if non-nil, zero value otherwise.

### GetReductionAlgorithmOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetReductionAlgorithmOk() (*ReductionAlgorithm, bool)`

GetReductionAlgorithmOk returns a tuple with the ReductionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReductionAlgorithm

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetReductionAlgorithm(v ReductionAlgorithm)`

SetReductionAlgorithm sets ReductionAlgorithm field to given value.


### GetOptionalAnalysis

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetOptionalAnalysis() []OptionalAnalysis`

GetOptionalAnalysis returns the OptionalAnalysis field if non-nil, zero value otherwise.

### GetOptionalAnalysisOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetOptionalAnalysisOk() (*[]OptionalAnalysis, bool)`

GetOptionalAnalysisOk returns a tuple with the OptionalAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalAnalysis

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetOptionalAnalysis(v []OptionalAnalysis)`

SetOptionalAnalysis sets OptionalAnalysis field to given value.

### HasOptionalAnalysis

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) HasOptionalAnalysis() bool`

HasOptionalAnalysis returns a boolean if a field has been set.

### GetTriggerCreateSampleVisualiztions

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetTriggerCreateSampleVisualiztions() bool`

GetTriggerCreateSampleVisualiztions returns the TriggerCreateSampleVisualiztions field if non-nil, zero value otherwise.

### GetTriggerCreateSampleVisualiztionsOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetTriggerCreateSampleVisualiztionsOk() (*bool, bool)`

GetTriggerCreateSampleVisualiztionsOk returns a tuple with the TriggerCreateSampleVisualiztions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCreateSampleVisualiztions

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetTriggerCreateSampleVisualiztions(v bool)`

SetTriggerCreateSampleVisualiztions sets TriggerCreateSampleVisualiztions field to given value.

### HasTriggerCreateSampleVisualiztions

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) HasTriggerCreateSampleVisualiztions() bool`

HasTriggerCreateSampleVisualiztions returns a boolean if a field has been set.

### GetUseCustomLatentSpace

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetUseCustomLatentSpace() bool`

GetUseCustomLatentSpace returns the UseCustomLatentSpace field if non-nil, zero value otherwise.

### GetUseCustomLatentSpaceOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetUseCustomLatentSpaceOk() (*bool, bool)`

GetUseCustomLatentSpaceOk returns a tuple with the UseCustomLatentSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomLatentSpace

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetUseCustomLatentSpace(v bool)`

SetUseCustomLatentSpace sets UseCustomLatentSpace field to given value.

### HasUseCustomLatentSpace

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) HasUseCustomLatentSpace() bool`

HasUseCustomLatentSpace returns a boolean if a field has been set.

### GetElementInstance

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetElementInstance() bool`

GetElementInstance returns the ElementInstance field if non-nil, zero value otherwise.

### GetElementInstanceOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetElementInstanceOk() (*bool, bool)`

GetElementInstanceOk returns a tuple with the ElementInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementInstance

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetElementInstance(v bool)`

SetElementInstance sets ElementInstance field to given value.

### HasElementInstance

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) HasElementInstance() bool`

HasElementInstance returns a boolean if a field has been set.

### GetForceExecute

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetForceExecute() bool`

GetForceExecute returns the ForceExecute field if non-nil, zero value otherwise.

### GetForceExecuteOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetForceExecuteOk() (*bool, bool)`

GetForceExecuteOk returns a tuple with the ForceExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExecute

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetForceExecute(v bool)`

SetForceExecute sets ForceExecute field to given value.

### HasForceExecute

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) HasForceExecute() bool`

HasForceExecute returns a boolean if a field has been set.

### GetLatentSpaceType

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetLatentSpaceType() string`

GetLatentSpaceType returns the LatentSpaceType field if non-nil, zero value otherwise.

### GetLatentSpaceTypeOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) GetLatentSpaceTypeOk() (*string, bool)`

GetLatentSpaceTypeOk returns a tuple with the LatentSpaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatentSpaceType

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) SetLatentSpaceType(v string)`

SetLatentSpaceType sets LatentSpaceType field to given value.

### HasLatentSpaceType

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrProjectIdOrReRunAfterFail) HasLatentSpaceType() bool`

HasLatentSpaceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


