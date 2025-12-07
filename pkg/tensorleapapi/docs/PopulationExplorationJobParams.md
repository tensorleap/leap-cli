# PopulationExplorationJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatentSpaceType** | Pointer to **string** |  | [optional] 
**ForceExecute** | Pointer to **bool** |  | [optional] 
**ElementInstance** | Pointer to **bool** |  | [optional] 
**UseCustomLatentSpace** | Pointer to **bool** |  | [optional] 
**OptionalAnalysis** | Pointer to [**[]OptionalAnalysis**](OptionalAnalysis.md) |  | [optional] 
**ReductionAlgorithm** | [**ReductionAlgorithm**](ReductionAlgorithm.md) |  | 
**ShouldFillRemainingWithUnbalanced** | **bool** |  | 
**BalanceBy** | **[]string** |  | 
**NumOfSamples** | **float64** |  | 
**DomainGapMetadata** | Pointer to **string** |  | [optional] 
**ProjectionMetric** | Pointer to **string** |  | [optional] 
**Digest** | **string** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**FromEpoch** | **float64** |  | 
**BatchSize** | **float64** |  | 
**SessionRunId** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewPopulationExplorationJobParams

`func NewPopulationExplorationJobParams(reductionAlgorithm ReductionAlgorithm, shouldFillRemainingWithUnbalanced bool, balanceBy []string, numOfSamples float64, digest string, fromEpoch float64, batchSize float64, sessionRunId string, type_ string, ) *PopulationExplorationJobParams`

NewPopulationExplorationJobParams instantiates a new PopulationExplorationJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPopulationExplorationJobParamsWithDefaults

`func NewPopulationExplorationJobParamsWithDefaults() *PopulationExplorationJobParams`

NewPopulationExplorationJobParamsWithDefaults instantiates a new PopulationExplorationJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatentSpaceType

`func (o *PopulationExplorationJobParams) GetLatentSpaceType() string`

GetLatentSpaceType returns the LatentSpaceType field if non-nil, zero value otherwise.

### GetLatentSpaceTypeOk

`func (o *PopulationExplorationJobParams) GetLatentSpaceTypeOk() (*string, bool)`

GetLatentSpaceTypeOk returns a tuple with the LatentSpaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatentSpaceType

`func (o *PopulationExplorationJobParams) SetLatentSpaceType(v string)`

SetLatentSpaceType sets LatentSpaceType field to given value.

### HasLatentSpaceType

`func (o *PopulationExplorationJobParams) HasLatentSpaceType() bool`

HasLatentSpaceType returns a boolean if a field has been set.

### GetForceExecute

`func (o *PopulationExplorationJobParams) GetForceExecute() bool`

GetForceExecute returns the ForceExecute field if non-nil, zero value otherwise.

### GetForceExecuteOk

`func (o *PopulationExplorationJobParams) GetForceExecuteOk() (*bool, bool)`

GetForceExecuteOk returns a tuple with the ForceExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExecute

`func (o *PopulationExplorationJobParams) SetForceExecute(v bool)`

SetForceExecute sets ForceExecute field to given value.

### HasForceExecute

`func (o *PopulationExplorationJobParams) HasForceExecute() bool`

HasForceExecute returns a boolean if a field has been set.

### GetElementInstance

`func (o *PopulationExplorationJobParams) GetElementInstance() bool`

GetElementInstance returns the ElementInstance field if non-nil, zero value otherwise.

### GetElementInstanceOk

`func (o *PopulationExplorationJobParams) GetElementInstanceOk() (*bool, bool)`

GetElementInstanceOk returns a tuple with the ElementInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementInstance

`func (o *PopulationExplorationJobParams) SetElementInstance(v bool)`

SetElementInstance sets ElementInstance field to given value.

### HasElementInstance

`func (o *PopulationExplorationJobParams) HasElementInstance() bool`

HasElementInstance returns a boolean if a field has been set.

### GetUseCustomLatentSpace

`func (o *PopulationExplorationJobParams) GetUseCustomLatentSpace() bool`

GetUseCustomLatentSpace returns the UseCustomLatentSpace field if non-nil, zero value otherwise.

### GetUseCustomLatentSpaceOk

`func (o *PopulationExplorationJobParams) GetUseCustomLatentSpaceOk() (*bool, bool)`

GetUseCustomLatentSpaceOk returns a tuple with the UseCustomLatentSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomLatentSpace

`func (o *PopulationExplorationJobParams) SetUseCustomLatentSpace(v bool)`

SetUseCustomLatentSpace sets UseCustomLatentSpace field to given value.

### HasUseCustomLatentSpace

`func (o *PopulationExplorationJobParams) HasUseCustomLatentSpace() bool`

HasUseCustomLatentSpace returns a boolean if a field has been set.

### GetOptionalAnalysis

`func (o *PopulationExplorationJobParams) GetOptionalAnalysis() []OptionalAnalysis`

GetOptionalAnalysis returns the OptionalAnalysis field if non-nil, zero value otherwise.

### GetOptionalAnalysisOk

`func (o *PopulationExplorationJobParams) GetOptionalAnalysisOk() (*[]OptionalAnalysis, bool)`

GetOptionalAnalysisOk returns a tuple with the OptionalAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalAnalysis

`func (o *PopulationExplorationJobParams) SetOptionalAnalysis(v []OptionalAnalysis)`

SetOptionalAnalysis sets OptionalAnalysis field to given value.

### HasOptionalAnalysis

`func (o *PopulationExplorationJobParams) HasOptionalAnalysis() bool`

HasOptionalAnalysis returns a boolean if a field has been set.

### GetReductionAlgorithm

`func (o *PopulationExplorationJobParams) GetReductionAlgorithm() ReductionAlgorithm`

GetReductionAlgorithm returns the ReductionAlgorithm field if non-nil, zero value otherwise.

### GetReductionAlgorithmOk

`func (o *PopulationExplorationJobParams) GetReductionAlgorithmOk() (*ReductionAlgorithm, bool)`

GetReductionAlgorithmOk returns a tuple with the ReductionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReductionAlgorithm

`func (o *PopulationExplorationJobParams) SetReductionAlgorithm(v ReductionAlgorithm)`

SetReductionAlgorithm sets ReductionAlgorithm field to given value.


### GetShouldFillRemainingWithUnbalanced

`func (o *PopulationExplorationJobParams) GetShouldFillRemainingWithUnbalanced() bool`

GetShouldFillRemainingWithUnbalanced returns the ShouldFillRemainingWithUnbalanced field if non-nil, zero value otherwise.

### GetShouldFillRemainingWithUnbalancedOk

`func (o *PopulationExplorationJobParams) GetShouldFillRemainingWithUnbalancedOk() (*bool, bool)`

GetShouldFillRemainingWithUnbalancedOk returns a tuple with the ShouldFillRemainingWithUnbalanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldFillRemainingWithUnbalanced

`func (o *PopulationExplorationJobParams) SetShouldFillRemainingWithUnbalanced(v bool)`

SetShouldFillRemainingWithUnbalanced sets ShouldFillRemainingWithUnbalanced field to given value.


### GetBalanceBy

`func (o *PopulationExplorationJobParams) GetBalanceBy() []string`

GetBalanceBy returns the BalanceBy field if non-nil, zero value otherwise.

### GetBalanceByOk

`func (o *PopulationExplorationJobParams) GetBalanceByOk() (*[]string, bool)`

GetBalanceByOk returns a tuple with the BalanceBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceBy

`func (o *PopulationExplorationJobParams) SetBalanceBy(v []string)`

SetBalanceBy sets BalanceBy field to given value.


### GetNumOfSamples

`func (o *PopulationExplorationJobParams) GetNumOfSamples() float64`

GetNumOfSamples returns the NumOfSamples field if non-nil, zero value otherwise.

### GetNumOfSamplesOk

`func (o *PopulationExplorationJobParams) GetNumOfSamplesOk() (*float64, bool)`

GetNumOfSamplesOk returns a tuple with the NumOfSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfSamples

`func (o *PopulationExplorationJobParams) SetNumOfSamples(v float64)`

SetNumOfSamples sets NumOfSamples field to given value.


### GetDomainGapMetadata

`func (o *PopulationExplorationJobParams) GetDomainGapMetadata() string`

GetDomainGapMetadata returns the DomainGapMetadata field if non-nil, zero value otherwise.

### GetDomainGapMetadataOk

`func (o *PopulationExplorationJobParams) GetDomainGapMetadataOk() (*string, bool)`

GetDomainGapMetadataOk returns a tuple with the DomainGapMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGapMetadata

`func (o *PopulationExplorationJobParams) SetDomainGapMetadata(v string)`

SetDomainGapMetadata sets DomainGapMetadata field to given value.

### HasDomainGapMetadata

`func (o *PopulationExplorationJobParams) HasDomainGapMetadata() bool`

HasDomainGapMetadata returns a boolean if a field has been set.

### GetProjectionMetric

`func (o *PopulationExplorationJobParams) GetProjectionMetric() string`

GetProjectionMetric returns the ProjectionMetric field if non-nil, zero value otherwise.

### GetProjectionMetricOk

`func (o *PopulationExplorationJobParams) GetProjectionMetricOk() (*string, bool)`

GetProjectionMetricOk returns a tuple with the ProjectionMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectionMetric

`func (o *PopulationExplorationJobParams) SetProjectionMetric(v string)`

SetProjectionMetric sets ProjectionMetric field to given value.

### HasProjectionMetric

`func (o *PopulationExplorationJobParams) HasProjectionMetric() bool`

HasProjectionMetric returns a boolean if a field has been set.

### GetDigest

`func (o *PopulationExplorationJobParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *PopulationExplorationJobParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *PopulationExplorationJobParams) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetFilters

`func (o *PopulationExplorationJobParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PopulationExplorationJobParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PopulationExplorationJobParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PopulationExplorationJobParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFromEpoch

`func (o *PopulationExplorationJobParams) GetFromEpoch() float64`

GetFromEpoch returns the FromEpoch field if non-nil, zero value otherwise.

### GetFromEpochOk

`func (o *PopulationExplorationJobParams) GetFromEpochOk() (*float64, bool)`

GetFromEpochOk returns a tuple with the FromEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEpoch

`func (o *PopulationExplorationJobParams) SetFromEpoch(v float64)`

SetFromEpoch sets FromEpoch field to given value.


### GetBatchSize

`func (o *PopulationExplorationJobParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *PopulationExplorationJobParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *PopulationExplorationJobParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetSessionRunId

`func (o *PopulationExplorationJobParams) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *PopulationExplorationJobParams) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *PopulationExplorationJobParams) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetType

`func (o *PopulationExplorationJobParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PopulationExplorationJobParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PopulationExplorationJobParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


