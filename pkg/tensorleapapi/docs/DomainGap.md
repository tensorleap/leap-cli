# DomainGap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**JobId** | **string** |  | 
**VersionId** | **string** |  | 
**VersionName** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**StatsFileUrl** | Pointer to **string** |  | [optional] 
**FilterFileUrl** | Pointer to **string** |  | [optional] 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**IsDeleted** | **bool** |  | 
**GroupAFilters** | [**[]ESFilter**](ESFilter.md) |  | 
**GroupBFilters** | [**[]ESFilter**](ESFilter.md) |  | 
**RunProcess** | Pointer to [**RunProcess**](RunProcess.md) |  | [optional] 

## Methods

### NewDomainGap

`func NewDomainGap(id string, jobId string, versionId string, versionName string, createdAt time.Time, createdBy string, status JobStatus, isDeleted bool, groupAFilters []ESFilter, groupBFilters []ESFilter, ) *DomainGap`

NewDomainGap instantiates a new DomainGap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainGapWithDefaults

`func NewDomainGapWithDefaults() *DomainGap`

NewDomainGapWithDefaults instantiates a new DomainGap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DomainGap) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainGap) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainGap) SetId(v string)`

SetId sets Id field to given value.


### GetJobId

`func (o *DomainGap) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *DomainGap) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *DomainGap) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetVersionId

`func (o *DomainGap) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *DomainGap) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *DomainGap) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersionName

`func (o *DomainGap) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *DomainGap) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *DomainGap) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.


### GetCreatedAt

`func (o *DomainGap) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainGap) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainGap) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *DomainGap) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DomainGap) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DomainGap) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetStatsFileUrl

`func (o *DomainGap) GetStatsFileUrl() string`

GetStatsFileUrl returns the StatsFileUrl field if non-nil, zero value otherwise.

### GetStatsFileUrlOk

`func (o *DomainGap) GetStatsFileUrlOk() (*string, bool)`

GetStatsFileUrlOk returns a tuple with the StatsFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsFileUrl

`func (o *DomainGap) SetStatsFileUrl(v string)`

SetStatsFileUrl sets StatsFileUrl field to given value.

### HasStatsFileUrl

`func (o *DomainGap) HasStatsFileUrl() bool`

HasStatsFileUrl returns a boolean if a field has been set.

### GetFilterFileUrl

`func (o *DomainGap) GetFilterFileUrl() string`

GetFilterFileUrl returns the FilterFileUrl field if non-nil, zero value otherwise.

### GetFilterFileUrlOk

`func (o *DomainGap) GetFilterFileUrlOk() (*string, bool)`

GetFilterFileUrlOk returns a tuple with the FilterFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterFileUrl

`func (o *DomainGap) SetFilterFileUrl(v string)`

SetFilterFileUrl sets FilterFileUrl field to given value.

### HasFilterFileUrl

`func (o *DomainGap) HasFilterFileUrl() bool`

HasFilterFileUrl returns a boolean if a field has been set.

### GetStatus

`func (o *DomainGap) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainGap) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainGap) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetIsDeleted

`func (o *DomainGap) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *DomainGap) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *DomainGap) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetGroupAFilters

`func (o *DomainGap) GetGroupAFilters() []ESFilter`

GetGroupAFilters returns the GroupAFilters field if non-nil, zero value otherwise.

### GetGroupAFiltersOk

`func (o *DomainGap) GetGroupAFiltersOk() (*[]ESFilter, bool)`

GetGroupAFiltersOk returns a tuple with the GroupAFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAFilters

`func (o *DomainGap) SetGroupAFilters(v []ESFilter)`

SetGroupAFilters sets GroupAFilters field to given value.


### GetGroupBFilters

`func (o *DomainGap) GetGroupBFilters() []ESFilter`

GetGroupBFilters returns the GroupBFilters field if non-nil, zero value otherwise.

### GetGroupBFiltersOk

`func (o *DomainGap) GetGroupBFiltersOk() (*[]ESFilter, bool)`

GetGroupBFiltersOk returns a tuple with the GroupBFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBFilters

`func (o *DomainGap) SetGroupBFilters(v []ESFilter)`

SetGroupBFilters sets GroupBFilters field to given value.


### GetRunProcess

`func (o *DomainGap) GetRunProcess() RunProcess`

GetRunProcess returns the RunProcess field if non-nil, zero value otherwise.

### GetRunProcessOk

`func (o *DomainGap) GetRunProcessOk() (*RunProcess, bool)`

GetRunProcessOk returns a tuple with the RunProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunProcess

`func (o *DomainGap) SetRunProcess(v RunProcess)`

SetRunProcess sets RunProcess field to given value.

### HasRunProcess

`func (o *DomainGap) HasRunProcess() bool`

HasRunProcess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


