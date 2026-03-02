# DatasetBalancing

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
**FileUrl** | Pointer to **string** |  | [optional] 
**FilterFileUrl** | Pointer to **string** |  | [optional] 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**IsDeleted** | **bool** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**MetadataTags** | **[]string** |  | 
**PrioritizedMetadataTags** | Pointer to **[]string** |  | [optional] 
**PercentageOfSamplesToPrune** | Pointer to **float64** |  | [optional] 

## Methods

### NewDatasetBalancing

`func NewDatasetBalancing(id string, jobId string, sessionRunId string, sessionRunName string, epoch float64, createdAt time.Time, createdBy string, status JobStatus, isDeleted bool, metadataTags []string, ) *DatasetBalancing`

NewDatasetBalancing instantiates a new DatasetBalancing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetBalancingWithDefaults

`func NewDatasetBalancingWithDefaults() *DatasetBalancing`

NewDatasetBalancingWithDefaults instantiates a new DatasetBalancing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatasetBalancing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatasetBalancing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatasetBalancing) SetId(v string)`

SetId sets Id field to given value.


### GetJobId

`func (o *DatasetBalancing) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *DatasetBalancing) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *DatasetBalancing) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetSessionRunId

`func (o *DatasetBalancing) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *DatasetBalancing) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *DatasetBalancing) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetSessionRunName

`func (o *DatasetBalancing) GetSessionRunName() string`

GetSessionRunName returns the SessionRunName field if non-nil, zero value otherwise.

### GetSessionRunNameOk

`func (o *DatasetBalancing) GetSessionRunNameOk() (*string, bool)`

GetSessionRunNameOk returns a tuple with the SessionRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunName

`func (o *DatasetBalancing) SetSessionRunName(v string)`

SetSessionRunName sets SessionRunName field to given value.


### GetEpoch

`func (o *DatasetBalancing) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *DatasetBalancing) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *DatasetBalancing) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetCreatedAt

`func (o *DatasetBalancing) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatasetBalancing) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatasetBalancing) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *DatasetBalancing) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DatasetBalancing) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DatasetBalancing) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetFileUrl

`func (o *DatasetBalancing) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *DatasetBalancing) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *DatasetBalancing) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *DatasetBalancing) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetFilterFileUrl

`func (o *DatasetBalancing) GetFilterFileUrl() string`

GetFilterFileUrl returns the FilterFileUrl field if non-nil, zero value otherwise.

### GetFilterFileUrlOk

`func (o *DatasetBalancing) GetFilterFileUrlOk() (*string, bool)`

GetFilterFileUrlOk returns a tuple with the FilterFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterFileUrl

`func (o *DatasetBalancing) SetFilterFileUrl(v string)`

SetFilterFileUrl sets FilterFileUrl field to given value.

### HasFilterFileUrl

`func (o *DatasetBalancing) HasFilterFileUrl() bool`

HasFilterFileUrl returns a boolean if a field has been set.

### GetStatus

`func (o *DatasetBalancing) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatasetBalancing) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatasetBalancing) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetIsDeleted

`func (o *DatasetBalancing) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *DatasetBalancing) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *DatasetBalancing) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetFilters

`func (o *DatasetBalancing) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DatasetBalancing) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DatasetBalancing) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DatasetBalancing) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMetadataTags

`func (o *DatasetBalancing) GetMetadataTags() []string`

GetMetadataTags returns the MetadataTags field if non-nil, zero value otherwise.

### GetMetadataTagsOk

`func (o *DatasetBalancing) GetMetadataTagsOk() (*[]string, bool)`

GetMetadataTagsOk returns a tuple with the MetadataTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataTags

`func (o *DatasetBalancing) SetMetadataTags(v []string)`

SetMetadataTags sets MetadataTags field to given value.


### GetPrioritizedMetadataTags

`func (o *DatasetBalancing) GetPrioritizedMetadataTags() []string`

GetPrioritizedMetadataTags returns the PrioritizedMetadataTags field if non-nil, zero value otherwise.

### GetPrioritizedMetadataTagsOk

`func (o *DatasetBalancing) GetPrioritizedMetadataTagsOk() (*[]string, bool)`

GetPrioritizedMetadataTagsOk returns a tuple with the PrioritizedMetadataTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritizedMetadataTags

`func (o *DatasetBalancing) SetPrioritizedMetadataTags(v []string)`

SetPrioritizedMetadataTags sets PrioritizedMetadataTags field to given value.

### HasPrioritizedMetadataTags

`func (o *DatasetBalancing) HasPrioritizedMetadataTags() bool`

HasPrioritizedMetadataTags returns a boolean if a field has been set.

### GetPercentageOfSamplesToPrune

`func (o *DatasetBalancing) GetPercentageOfSamplesToPrune() float64`

GetPercentageOfSamplesToPrune returns the PercentageOfSamplesToPrune field if non-nil, zero value otherwise.

### GetPercentageOfSamplesToPruneOk

`func (o *DatasetBalancing) GetPercentageOfSamplesToPruneOk() (*float64, bool)`

GetPercentageOfSamplesToPruneOk returns a tuple with the PercentageOfSamplesToPrune field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfSamplesToPrune

`func (o *DatasetBalancing) SetPercentageOfSamplesToPrune(v float64)`

SetPercentageOfSamplesToPrune sets PercentageOfSamplesToPrune field to given value.

### HasPercentageOfSamplesToPrune

`func (o *DatasetBalancing) HasPercentageOfSamplesToPrune() bool`

HasPercentageOfSamplesToPrune returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


