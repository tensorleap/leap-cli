# SyntheticData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**JobId** | **string** |  | 
**SessionRunId** | **string** |  | 
**SessionRunName** | **string** |  | 
**Epoch** | **float64** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**NextTrialsFileUrl** | Pointer to **string** |  | [optional] 
**SuggestionsFileUrl** | Pointer to **string** |  | [optional] 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**Sources** | [**[]GenerateSyntheticDataParamsSourcesInner**](GenerateSyntheticDataParamsSourcesInner.md) |  | 
**TargetFilters** | [**[]ESFilter**](ESFilter.md) |  | 

## Methods

### NewSyntheticData

`func NewSyntheticData(id string, jobId string, sessionRunId string, sessionRunName string, epoch float64, createdAt time.Time, createdBy string, status JobStatus, sources []GenerateSyntheticDataParamsSourcesInner, targetFilters []ESFilter, ) *SyntheticData`

NewSyntheticData instantiates a new SyntheticData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticDataWithDefaults

`func NewSyntheticDataWithDefaults() *SyntheticData`

NewSyntheticDataWithDefaults instantiates a new SyntheticData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyntheticData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyntheticData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyntheticData) SetId(v string)`

SetId sets Id field to given value.


### GetJobId

`func (o *SyntheticData) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SyntheticData) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SyntheticData) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetSessionRunId

`func (o *SyntheticData) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *SyntheticData) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *SyntheticData) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetSessionRunName

`func (o *SyntheticData) GetSessionRunName() string`

GetSessionRunName returns the SessionRunName field if non-nil, zero value otherwise.

### GetSessionRunNameOk

`func (o *SyntheticData) GetSessionRunNameOk() (*string, bool)`

GetSessionRunNameOk returns a tuple with the SessionRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunName

`func (o *SyntheticData) SetSessionRunName(v string)`

SetSessionRunName sets SessionRunName field to given value.


### GetEpoch

`func (o *SyntheticData) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *SyntheticData) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *SyntheticData) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetCreatedAt

`func (o *SyntheticData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SyntheticData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SyntheticData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *SyntheticData) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SyntheticData) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SyntheticData) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetNextTrialsFileUrl

`func (o *SyntheticData) GetNextTrialsFileUrl() string`

GetNextTrialsFileUrl returns the NextTrialsFileUrl field if non-nil, zero value otherwise.

### GetNextTrialsFileUrlOk

`func (o *SyntheticData) GetNextTrialsFileUrlOk() (*string, bool)`

GetNextTrialsFileUrlOk returns a tuple with the NextTrialsFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTrialsFileUrl

`func (o *SyntheticData) SetNextTrialsFileUrl(v string)`

SetNextTrialsFileUrl sets NextTrialsFileUrl field to given value.

### HasNextTrialsFileUrl

`func (o *SyntheticData) HasNextTrialsFileUrl() bool`

HasNextTrialsFileUrl returns a boolean if a field has been set.

### GetSuggestionsFileUrl

`func (o *SyntheticData) GetSuggestionsFileUrl() string`

GetSuggestionsFileUrl returns the SuggestionsFileUrl field if non-nil, zero value otherwise.

### GetSuggestionsFileUrlOk

`func (o *SyntheticData) GetSuggestionsFileUrlOk() (*string, bool)`

GetSuggestionsFileUrlOk returns a tuple with the SuggestionsFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionsFileUrl

`func (o *SyntheticData) SetSuggestionsFileUrl(v string)`

SetSuggestionsFileUrl sets SuggestionsFileUrl field to given value.

### HasSuggestionsFileUrl

`func (o *SyntheticData) HasSuggestionsFileUrl() bool`

HasSuggestionsFileUrl returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticData) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticData) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticData) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetSources

`func (o *SyntheticData) GetSources() []GenerateSyntheticDataParamsSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SyntheticData) GetSourcesOk() (*[]GenerateSyntheticDataParamsSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SyntheticData) SetSources(v []GenerateSyntheticDataParamsSourcesInner)`

SetSources sets Sources field to given value.


### GetTargetFilters

`func (o *SyntheticData) GetTargetFilters() []ESFilter`

GetTargetFilters returns the TargetFilters field if non-nil, zero value otherwise.

### GetTargetFiltersOk

`func (o *SyntheticData) GetTargetFiltersOk() (*[]ESFilter, bool)`

GetTargetFiltersOk returns a tuple with the TargetFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFilters

`func (o *SyntheticData) SetTargetFilters(v []ESFilter)`

SetTargetFilters sets TargetFilters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


