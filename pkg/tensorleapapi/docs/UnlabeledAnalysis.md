# UnlabeledAnalysis

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
**Status** | [**JobStatus**](JobStatus.md) |  | 
**IsDeleted** | **bool** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**LatentSpaceType** | Pointer to **string** |  | [optional] 
**RunProcess** | Pointer to [**RunProcess**](RunProcess.md) |  | [optional] 

## Methods

### NewUnlabeledAnalysis

`func NewUnlabeledAnalysis(id string, jobId string, versionId string, versionName string, createdAt time.Time, createdBy string, status JobStatus, isDeleted bool, ) *UnlabeledAnalysis`

NewUnlabeledAnalysis instantiates a new UnlabeledAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlabeledAnalysisWithDefaults

`func NewUnlabeledAnalysisWithDefaults() *UnlabeledAnalysis`

NewUnlabeledAnalysisWithDefaults instantiates a new UnlabeledAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnlabeledAnalysis) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnlabeledAnalysis) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnlabeledAnalysis) SetId(v string)`

SetId sets Id field to given value.


### GetJobId

`func (o *UnlabeledAnalysis) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *UnlabeledAnalysis) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *UnlabeledAnalysis) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetVersionId

`func (o *UnlabeledAnalysis) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *UnlabeledAnalysis) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *UnlabeledAnalysis) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersionName

`func (o *UnlabeledAnalysis) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *UnlabeledAnalysis) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *UnlabeledAnalysis) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.


### GetCreatedAt

`func (o *UnlabeledAnalysis) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UnlabeledAnalysis) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UnlabeledAnalysis) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *UnlabeledAnalysis) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UnlabeledAnalysis) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UnlabeledAnalysis) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetStatsFileUrl

`func (o *UnlabeledAnalysis) GetStatsFileUrl() string`

GetStatsFileUrl returns the StatsFileUrl field if non-nil, zero value otherwise.

### GetStatsFileUrlOk

`func (o *UnlabeledAnalysis) GetStatsFileUrlOk() (*string, bool)`

GetStatsFileUrlOk returns a tuple with the StatsFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsFileUrl

`func (o *UnlabeledAnalysis) SetStatsFileUrl(v string)`

SetStatsFileUrl sets StatsFileUrl field to given value.

### HasStatsFileUrl

`func (o *UnlabeledAnalysis) HasStatsFileUrl() bool`

HasStatsFileUrl returns a boolean if a field has been set.

### GetStatus

`func (o *UnlabeledAnalysis) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnlabeledAnalysis) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnlabeledAnalysis) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetIsDeleted

`func (o *UnlabeledAnalysis) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnlabeledAnalysis) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnlabeledAnalysis) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetFilters

`func (o *UnlabeledAnalysis) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *UnlabeledAnalysis) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *UnlabeledAnalysis) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *UnlabeledAnalysis) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLatentSpaceType

`func (o *UnlabeledAnalysis) GetLatentSpaceType() string`

GetLatentSpaceType returns the LatentSpaceType field if non-nil, zero value otherwise.

### GetLatentSpaceTypeOk

`func (o *UnlabeledAnalysis) GetLatentSpaceTypeOk() (*string, bool)`

GetLatentSpaceTypeOk returns a tuple with the LatentSpaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatentSpaceType

`func (o *UnlabeledAnalysis) SetLatentSpaceType(v string)`

SetLatentSpaceType sets LatentSpaceType field to given value.

### HasLatentSpaceType

`func (o *UnlabeledAnalysis) HasLatentSpaceType() bool`

HasLatentSpaceType returns a boolean if a field has been set.

### GetRunProcess

`func (o *UnlabeledAnalysis) GetRunProcess() RunProcess`

GetRunProcess returns the RunProcess field if non-nil, zero value otherwise.

### GetRunProcessOk

`func (o *UnlabeledAnalysis) GetRunProcessOk() (*RunProcess, bool)`

GetRunProcessOk returns a tuple with the RunProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunProcess

`func (o *UnlabeledAnalysis) SetRunProcess(v RunProcess)`

SetRunProcess sets RunProcess field to given value.

### HasRunProcess

`func (o *UnlabeledAnalysis) HasRunProcess() bool`

HasRunProcess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


