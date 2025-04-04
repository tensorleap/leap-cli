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
**ProjectionMetric** | Pointer to **string** |  | [optional] 
**DomainGapMetadata** | Pointer to **string** |  | [optional] 
**DisplayParams** | [**PopulationExplorationDisplayParams**](PopulationExplorationDisplayParams.md) |  | 
**TriggerCreateSampleVisualiztions** | Pointer to **bool** |  | [optional] 
**ReRunAfterFail** | Pointer to **bool** |  | [optional] 

## Methods

### NewPopulationExplorationParams

`func NewPopulationExplorationParams(sessionRunId string, projectId string, batchSize float64, fromEpoch float64, digest string, displayParams PopulationExplorationDisplayParams, ) *PopulationExplorationParams`

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


### GetProjectionMetric

`func (o *PopulationExplorationParams) GetProjectionMetric() string`

GetProjectionMetric returns the ProjectionMetric field if non-nil, zero value otherwise.

### GetProjectionMetricOk

`func (o *PopulationExplorationParams) GetProjectionMetricOk() (*string, bool)`

GetProjectionMetricOk returns a tuple with the ProjectionMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectionMetric

`func (o *PopulationExplorationParams) SetProjectionMetric(v string)`

SetProjectionMetric sets ProjectionMetric field to given value.

### HasProjectionMetric

`func (o *PopulationExplorationParams) HasProjectionMetric() bool`

HasProjectionMetric returns a boolean if a field has been set.

### GetDomainGapMetadata

`func (o *PopulationExplorationParams) GetDomainGapMetadata() string`

GetDomainGapMetadata returns the DomainGapMetadata field if non-nil, zero value otherwise.

### GetDomainGapMetadataOk

`func (o *PopulationExplorationParams) GetDomainGapMetadataOk() (*string, bool)`

GetDomainGapMetadataOk returns a tuple with the DomainGapMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGapMetadata

`func (o *PopulationExplorationParams) SetDomainGapMetadata(v string)`

SetDomainGapMetadata sets DomainGapMetadata field to given value.

### HasDomainGapMetadata

`func (o *PopulationExplorationParams) HasDomainGapMetadata() bool`

HasDomainGapMetadata returns a boolean if a field has been set.

### GetDisplayParams

`func (o *PopulationExplorationParams) GetDisplayParams() PopulationExplorationDisplayParams`

GetDisplayParams returns the DisplayParams field if non-nil, zero value otherwise.

### GetDisplayParamsOk

`func (o *PopulationExplorationParams) GetDisplayParamsOk() (*PopulationExplorationDisplayParams, bool)`

GetDisplayParamsOk returns a tuple with the DisplayParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayParams

`func (o *PopulationExplorationParams) SetDisplayParams(v PopulationExplorationDisplayParams)`

SetDisplayParams sets DisplayParams field to given value.


### GetTriggerCreateSampleVisualiztions

`func (o *PopulationExplorationParams) GetTriggerCreateSampleVisualiztions() bool`

GetTriggerCreateSampleVisualiztions returns the TriggerCreateSampleVisualiztions field if non-nil, zero value otherwise.

### GetTriggerCreateSampleVisualiztionsOk

`func (o *PopulationExplorationParams) GetTriggerCreateSampleVisualiztionsOk() (*bool, bool)`

GetTriggerCreateSampleVisualiztionsOk returns a tuple with the TriggerCreateSampleVisualiztions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCreateSampleVisualiztions

`func (o *PopulationExplorationParams) SetTriggerCreateSampleVisualiztions(v bool)`

SetTriggerCreateSampleVisualiztions sets TriggerCreateSampleVisualiztions field to given value.

### HasTriggerCreateSampleVisualiztions

`func (o *PopulationExplorationParams) HasTriggerCreateSampleVisualiztions() bool`

HasTriggerCreateSampleVisualiztions returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


