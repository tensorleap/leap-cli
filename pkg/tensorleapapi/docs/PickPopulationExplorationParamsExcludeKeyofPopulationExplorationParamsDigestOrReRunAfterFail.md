# PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail

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
**NumOfSamples** | **float64** |  | 
**BalanceBy** | **[]string** |  | 
**ShouldFillRemainingWithUnbalanced** | **bool** |  | 
**ReductionAlgorithm** | [**ReductionAlgorithm**](ReductionAlgorithm.md) |  | 
**UseCustomLatentSpace** | Pointer to **bool** |  | [optional] 
**ElementInstance** | Pointer to **bool** |  | [optional] 
**ForceExecute** | Pointer to **bool** |  | [optional] 
**LatentSpaceType** | Pointer to **string** |  | [optional] 

## Methods

### NewPickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail

`func NewPickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail(sessionRunId string, projectId string, batchSize float64, fromEpoch float64, numOfSamples float64, balanceBy []string, shouldFillRemainingWithUnbalanced bool, reductionAlgorithm ReductionAlgorithm, ) *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail`

NewPickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail instantiates a new PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFailWithDefaults

`func NewPickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFailWithDefaults() *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail`

NewPickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFailWithDefaults instantiates a new PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionRunId

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetProjectId

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetBatchSize

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetFromEpoch

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetFromEpoch() float64`

GetFromEpoch returns the FromEpoch field if non-nil, zero value otherwise.

### GetFromEpochOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetFromEpochOk() (*float64, bool)`

GetFromEpochOk returns a tuple with the FromEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEpoch

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetFromEpoch(v float64)`

SetFromEpoch sets FromEpoch field to given value.


### GetFilters

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTimeFilter

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetTimeFilter() ESFilter`

GetTimeFilter returns the TimeFilter field if non-nil, zero value otherwise.

### GetTimeFilterOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetTimeFilterOk() (*ESFilter, bool)`

GetTimeFilterOk returns a tuple with the TimeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFilter

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetTimeFilter(v ESFilter)`

SetTimeFilter sets TimeFilter field to given value.

### HasTimeFilter

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) HasTimeFilter() bool`

HasTimeFilter returns a boolean if a field has been set.

### GetNotApplyTimeFilterOnUnlabeledOnly

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetNotApplyTimeFilterOnUnlabeledOnly() bool`

GetNotApplyTimeFilterOnUnlabeledOnly returns the NotApplyTimeFilterOnUnlabeledOnly field if non-nil, zero value otherwise.

### GetNotApplyTimeFilterOnUnlabeledOnlyOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetNotApplyTimeFilterOnUnlabeledOnlyOk() (*bool, bool)`

GetNotApplyTimeFilterOnUnlabeledOnlyOk returns a tuple with the NotApplyTimeFilterOnUnlabeledOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplyTimeFilterOnUnlabeledOnly

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetNotApplyTimeFilterOnUnlabeledOnly(v bool)`

SetNotApplyTimeFilterOnUnlabeledOnly sets NotApplyTimeFilterOnUnlabeledOnly field to given value.

### HasNotApplyTimeFilterOnUnlabeledOnly

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) HasNotApplyTimeFilterOnUnlabeledOnly() bool`

HasNotApplyTimeFilterOnUnlabeledOnly returns a boolean if a field has been set.

### GetNumOfSamples

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetNumOfSamples() float64`

GetNumOfSamples returns the NumOfSamples field if non-nil, zero value otherwise.

### GetNumOfSamplesOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetNumOfSamplesOk() (*float64, bool)`

GetNumOfSamplesOk returns a tuple with the NumOfSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfSamples

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetNumOfSamples(v float64)`

SetNumOfSamples sets NumOfSamples field to given value.


### GetBalanceBy

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetBalanceBy() []string`

GetBalanceBy returns the BalanceBy field if non-nil, zero value otherwise.

### GetBalanceByOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetBalanceByOk() (*[]string, bool)`

GetBalanceByOk returns a tuple with the BalanceBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceBy

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetBalanceBy(v []string)`

SetBalanceBy sets BalanceBy field to given value.


### GetShouldFillRemainingWithUnbalanced

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetShouldFillRemainingWithUnbalanced() bool`

GetShouldFillRemainingWithUnbalanced returns the ShouldFillRemainingWithUnbalanced field if non-nil, zero value otherwise.

### GetShouldFillRemainingWithUnbalancedOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetShouldFillRemainingWithUnbalancedOk() (*bool, bool)`

GetShouldFillRemainingWithUnbalancedOk returns a tuple with the ShouldFillRemainingWithUnbalanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldFillRemainingWithUnbalanced

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetShouldFillRemainingWithUnbalanced(v bool)`

SetShouldFillRemainingWithUnbalanced sets ShouldFillRemainingWithUnbalanced field to given value.


### GetReductionAlgorithm

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetReductionAlgorithm() ReductionAlgorithm`

GetReductionAlgorithm returns the ReductionAlgorithm field if non-nil, zero value otherwise.

### GetReductionAlgorithmOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetReductionAlgorithmOk() (*ReductionAlgorithm, bool)`

GetReductionAlgorithmOk returns a tuple with the ReductionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReductionAlgorithm

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetReductionAlgorithm(v ReductionAlgorithm)`

SetReductionAlgorithm sets ReductionAlgorithm field to given value.


### GetUseCustomLatentSpace

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetUseCustomLatentSpace() bool`

GetUseCustomLatentSpace returns the UseCustomLatentSpace field if non-nil, zero value otherwise.

### GetUseCustomLatentSpaceOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetUseCustomLatentSpaceOk() (*bool, bool)`

GetUseCustomLatentSpaceOk returns a tuple with the UseCustomLatentSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomLatentSpace

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetUseCustomLatentSpace(v bool)`

SetUseCustomLatentSpace sets UseCustomLatentSpace field to given value.

### HasUseCustomLatentSpace

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) HasUseCustomLatentSpace() bool`

HasUseCustomLatentSpace returns a boolean if a field has been set.

### GetElementInstance

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetElementInstance() bool`

GetElementInstance returns the ElementInstance field if non-nil, zero value otherwise.

### GetElementInstanceOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetElementInstanceOk() (*bool, bool)`

GetElementInstanceOk returns a tuple with the ElementInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementInstance

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetElementInstance(v bool)`

SetElementInstance sets ElementInstance field to given value.

### HasElementInstance

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) HasElementInstance() bool`

HasElementInstance returns a boolean if a field has been set.

### GetForceExecute

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetForceExecute() bool`

GetForceExecute returns the ForceExecute field if non-nil, zero value otherwise.

### GetForceExecuteOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetForceExecuteOk() (*bool, bool)`

GetForceExecuteOk returns a tuple with the ForceExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExecute

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetForceExecute(v bool)`

SetForceExecute sets ForceExecute field to given value.

### HasForceExecute

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) HasForceExecute() bool`

HasForceExecute returns a boolean if a field has been set.

### GetLatentSpaceType

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetLatentSpaceType() string`

GetLatentSpaceType returns the LatentSpaceType field if non-nil, zero value otherwise.

### GetLatentSpaceTypeOk

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) GetLatentSpaceTypeOk() (*string, bool)`

GetLatentSpaceTypeOk returns a tuple with the LatentSpaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatentSpaceType

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) SetLatentSpaceType(v string)`

SetLatentSpaceType sets LatentSpaceType field to given value.

### HasLatentSpaceType

`func (o *PickPopulationExplorationParamsExcludeKeyofPopulationExplorationParamsDigestOrReRunAfterFail) HasLatentSpaceType() bool`

HasLatentSpaceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


