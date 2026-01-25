# PopulationExplorationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionRunId** | **string** |  | 
**ProjectId** | **string** |  | 
**BatchSize** | **float64** |  | 
**FromEpoch** | **float64** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**TimeFilter** | Pointer to [**ESFilter**](ESFilter.md) |  | [optional] 
**NotApplyTimeFilterOnUnlabeledOnly** | Pointer to **bool** |  | [optional] 
**Digest** | **string** |  | 
**NumOfSamples** | **float64** |  | 
**BalanceBy** | **[]string** |  | 
**ShouldFillRemainingWithUnbalanced** | **bool** |  | 
**ReductionAlgorithm** | [**ReductionAlgorithm**](ReductionAlgorithm.md) |  | 
**ReRunAfterFail** | Pointer to **bool** |  | [optional] 
**UseCustomLatentSpace** | Pointer to **bool** |  | [optional] 
**ElementInstance** | Pointer to **bool** |  | [optional] 
**ForceExecute** | Pointer to **bool** |  | [optional] 
**LatentSpaceType** | Pointer to **string** |  | [optional] 

## Methods

### NewPopulationExplorationParams

`func NewPopulationExplorationParams(sessionRunId string, projectId string, batchSize float64, fromEpoch float64, digest string, numOfSamples float64, balanceBy []string, shouldFillRemainingWithUnbalanced bool, reductionAlgorithm ReductionAlgorithm, ) *PopulationExplorationParams`

NewPopulationExplorationParams instantiates a new PopulationExplorationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPopulationExplorationParamsWithDefaults

`func NewPopulationExplorationParamsWithDefaults() *PopulationExplorationParams`

NewPopulationExplorationParamsWithDefaults instantiates a new PopulationExplorationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionRunId

`func (o *PopulationExplorationParams) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *PopulationExplorationParams) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *PopulationExplorationParams) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetProjectId

`func (o *PopulationExplorationParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PopulationExplorationParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PopulationExplorationParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetBatchSize

`func (o *PopulationExplorationParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *PopulationExplorationParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *PopulationExplorationParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetFromEpoch

`func (o *PopulationExplorationParams) GetFromEpoch() float64`

GetFromEpoch returns the FromEpoch field if non-nil, zero value otherwise.

### GetFromEpochOk

`func (o *PopulationExplorationParams) GetFromEpochOk() (*float64, bool)`

GetFromEpochOk returns a tuple with the FromEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEpoch

`func (o *PopulationExplorationParams) SetFromEpoch(v float64)`

SetFromEpoch sets FromEpoch field to given value.


### GetFilters

`func (o *PopulationExplorationParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PopulationExplorationParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PopulationExplorationParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PopulationExplorationParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTimeFilter

`func (o *PopulationExplorationParams) GetTimeFilter() ESFilter`

GetTimeFilter returns the TimeFilter field if non-nil, zero value otherwise.

### GetTimeFilterOk

`func (o *PopulationExplorationParams) GetTimeFilterOk() (*ESFilter, bool)`

GetTimeFilterOk returns a tuple with the TimeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFilter

`func (o *PopulationExplorationParams) SetTimeFilter(v ESFilter)`

SetTimeFilter sets TimeFilter field to given value.

### HasTimeFilter

`func (o *PopulationExplorationParams) HasTimeFilter() bool`

HasTimeFilter returns a boolean if a field has been set.

### GetNotApplyTimeFilterOnUnlabeledOnly

`func (o *PopulationExplorationParams) GetNotApplyTimeFilterOnUnlabeledOnly() bool`

GetNotApplyTimeFilterOnUnlabeledOnly returns the NotApplyTimeFilterOnUnlabeledOnly field if non-nil, zero value otherwise.

### GetNotApplyTimeFilterOnUnlabeledOnlyOk

`func (o *PopulationExplorationParams) GetNotApplyTimeFilterOnUnlabeledOnlyOk() (*bool, bool)`

GetNotApplyTimeFilterOnUnlabeledOnlyOk returns a tuple with the NotApplyTimeFilterOnUnlabeledOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplyTimeFilterOnUnlabeledOnly

`func (o *PopulationExplorationParams) SetNotApplyTimeFilterOnUnlabeledOnly(v bool)`

SetNotApplyTimeFilterOnUnlabeledOnly sets NotApplyTimeFilterOnUnlabeledOnly field to given value.

### HasNotApplyTimeFilterOnUnlabeledOnly

`func (o *PopulationExplorationParams) HasNotApplyTimeFilterOnUnlabeledOnly() bool`

HasNotApplyTimeFilterOnUnlabeledOnly returns a boolean if a field has been set.

### GetDigest

`func (o *PopulationExplorationParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *PopulationExplorationParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *PopulationExplorationParams) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetNumOfSamples

`func (o *PopulationExplorationParams) GetNumOfSamples() float64`

GetNumOfSamples returns the NumOfSamples field if non-nil, zero value otherwise.

### GetNumOfSamplesOk

`func (o *PopulationExplorationParams) GetNumOfSamplesOk() (*float64, bool)`

GetNumOfSamplesOk returns a tuple with the NumOfSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfSamples

`func (o *PopulationExplorationParams) SetNumOfSamples(v float64)`

SetNumOfSamples sets NumOfSamples field to given value.


### GetBalanceBy

`func (o *PopulationExplorationParams) GetBalanceBy() []string`

GetBalanceBy returns the BalanceBy field if non-nil, zero value otherwise.

### GetBalanceByOk

`func (o *PopulationExplorationParams) GetBalanceByOk() (*[]string, bool)`

GetBalanceByOk returns a tuple with the BalanceBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceBy

`func (o *PopulationExplorationParams) SetBalanceBy(v []string)`

SetBalanceBy sets BalanceBy field to given value.


### GetShouldFillRemainingWithUnbalanced

`func (o *PopulationExplorationParams) GetShouldFillRemainingWithUnbalanced() bool`

GetShouldFillRemainingWithUnbalanced returns the ShouldFillRemainingWithUnbalanced field if non-nil, zero value otherwise.

### GetShouldFillRemainingWithUnbalancedOk

`func (o *PopulationExplorationParams) GetShouldFillRemainingWithUnbalancedOk() (*bool, bool)`

GetShouldFillRemainingWithUnbalancedOk returns a tuple with the ShouldFillRemainingWithUnbalanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldFillRemainingWithUnbalanced

`func (o *PopulationExplorationParams) SetShouldFillRemainingWithUnbalanced(v bool)`

SetShouldFillRemainingWithUnbalanced sets ShouldFillRemainingWithUnbalanced field to given value.


### GetReductionAlgorithm

`func (o *PopulationExplorationParams) GetReductionAlgorithm() ReductionAlgorithm`

GetReductionAlgorithm returns the ReductionAlgorithm field if non-nil, zero value otherwise.

### GetReductionAlgorithmOk

`func (o *PopulationExplorationParams) GetReductionAlgorithmOk() (*ReductionAlgorithm, bool)`

GetReductionAlgorithmOk returns a tuple with the ReductionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReductionAlgorithm

`func (o *PopulationExplorationParams) SetReductionAlgorithm(v ReductionAlgorithm)`

SetReductionAlgorithm sets ReductionAlgorithm field to given value.


### GetReRunAfterFail

`func (o *PopulationExplorationParams) GetReRunAfterFail() bool`

GetReRunAfterFail returns the ReRunAfterFail field if non-nil, zero value otherwise.

### GetReRunAfterFailOk

`func (o *PopulationExplorationParams) GetReRunAfterFailOk() (*bool, bool)`

GetReRunAfterFailOk returns a tuple with the ReRunAfterFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReRunAfterFail

`func (o *PopulationExplorationParams) SetReRunAfterFail(v bool)`

SetReRunAfterFail sets ReRunAfterFail field to given value.

### HasReRunAfterFail

`func (o *PopulationExplorationParams) HasReRunAfterFail() bool`

HasReRunAfterFail returns a boolean if a field has been set.

### GetUseCustomLatentSpace

`func (o *PopulationExplorationParams) GetUseCustomLatentSpace() bool`

GetUseCustomLatentSpace returns the UseCustomLatentSpace field if non-nil, zero value otherwise.

### GetUseCustomLatentSpaceOk

`func (o *PopulationExplorationParams) GetUseCustomLatentSpaceOk() (*bool, bool)`

GetUseCustomLatentSpaceOk returns a tuple with the UseCustomLatentSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomLatentSpace

`func (o *PopulationExplorationParams) SetUseCustomLatentSpace(v bool)`

SetUseCustomLatentSpace sets UseCustomLatentSpace field to given value.

### HasUseCustomLatentSpace

`func (o *PopulationExplorationParams) HasUseCustomLatentSpace() bool`

HasUseCustomLatentSpace returns a boolean if a field has been set.

### GetElementInstance

`func (o *PopulationExplorationParams) GetElementInstance() bool`

GetElementInstance returns the ElementInstance field if non-nil, zero value otherwise.

### GetElementInstanceOk

`func (o *PopulationExplorationParams) GetElementInstanceOk() (*bool, bool)`

GetElementInstanceOk returns a tuple with the ElementInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementInstance

`func (o *PopulationExplorationParams) SetElementInstance(v bool)`

SetElementInstance sets ElementInstance field to given value.

### HasElementInstance

`func (o *PopulationExplorationParams) HasElementInstance() bool`

HasElementInstance returns a boolean if a field has been set.

### GetForceExecute

`func (o *PopulationExplorationParams) GetForceExecute() bool`

GetForceExecute returns the ForceExecute field if non-nil, zero value otherwise.

### GetForceExecuteOk

`func (o *PopulationExplorationParams) GetForceExecuteOk() (*bool, bool)`

GetForceExecuteOk returns a tuple with the ForceExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExecute

`func (o *PopulationExplorationParams) SetForceExecute(v bool)`

SetForceExecute sets ForceExecute field to given value.

### HasForceExecute

`func (o *PopulationExplorationParams) HasForceExecute() bool`

HasForceExecute returns a boolean if a field has been set.

### GetLatentSpaceType

`func (o *PopulationExplorationParams) GetLatentSpaceType() string`

GetLatentSpaceType returns the LatentSpaceType field if non-nil, zero value otherwise.

### GetLatentSpaceTypeOk

`func (o *PopulationExplorationParams) GetLatentSpaceTypeOk() (*string, bool)`

GetLatentSpaceTypeOk returns a tuple with the LatentSpaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatentSpaceType

`func (o *PopulationExplorationParams) SetLatentSpaceType(v string)`

SetLatentSpaceType sets LatentSpaceType field to given value.

### HasLatentSpaceType

`func (o *PopulationExplorationParams) HasLatentSpaceType() bool`

HasLatentSpaceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


