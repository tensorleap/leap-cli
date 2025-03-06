# PopulationExplorationJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerCreateSampleVisualiztions** | Pointer to **bool** |  | [optional] 
**DisplayParams** | [**PopulationExplorationDisplayParams**](PopulationExplorationDisplayParams.md) |  | 
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

`func NewPopulationExplorationJobParams(displayParams PopulationExplorationDisplayParams, digest string, fromEpoch float64, batchSize float64, sessionRunId string, type_ string, ) *PopulationExplorationJobParams`

NewPopulationExplorationJobParams instantiates a new PopulationExplorationJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPopulationExplorationJobParamsWithDefaults

`func NewPopulationExplorationJobParamsWithDefaults() *PopulationExplorationJobParams`

NewPopulationExplorationJobParamsWithDefaults instantiates a new PopulationExplorationJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerCreateSampleVisualiztions

`func (o *PopulationExplorationJobParams) GetTriggerCreateSampleVisualiztions() bool`

GetTriggerCreateSampleVisualiztions returns the TriggerCreateSampleVisualiztions field if non-nil, zero value otherwise.

### GetTriggerCreateSampleVisualiztionsOk

`func (o *PopulationExplorationJobParams) GetTriggerCreateSampleVisualiztionsOk() (*bool, bool)`

GetTriggerCreateSampleVisualiztionsOk returns a tuple with the TriggerCreateSampleVisualiztions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCreateSampleVisualiztions

`func (o *PopulationExplorationJobParams) SetTriggerCreateSampleVisualiztions(v bool)`

SetTriggerCreateSampleVisualiztions sets TriggerCreateSampleVisualiztions field to given value.

### HasTriggerCreateSampleVisualiztions

`func (o *PopulationExplorationJobParams) HasTriggerCreateSampleVisualiztions() bool`

HasTriggerCreateSampleVisualiztions returns a boolean if a field has been set.

### GetDisplayParams

`func (o *PopulationExplorationJobParams) GetDisplayParams() PopulationExplorationDisplayParams`

GetDisplayParams returns the DisplayParams field if non-nil, zero value otherwise.

### GetDisplayParamsOk

`func (o *PopulationExplorationJobParams) GetDisplayParamsOk() (*PopulationExplorationDisplayParams, bool)`

GetDisplayParamsOk returns a tuple with the DisplayParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayParams

`func (o *PopulationExplorationJobParams) SetDisplayParams(v PopulationExplorationDisplayParams)`

SetDisplayParams sets DisplayParams field to given value.


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


