# ManualInsightRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**JobId** | **string** |  | 
**VersionId** | **string** |  | 
**VersionName** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**IsDeleted** | **bool** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**RunProcess** | Pointer to [**RunProcess**](RunProcess.md) |  | [optional] 
**InsightFileUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewManualInsightRun

`func NewManualInsightRun(id string, jobId string, versionId string, versionName string, createdAt time.Time, createdBy string, status JobStatus, isDeleted bool, ) *ManualInsightRun`

NewManualInsightRun instantiates a new ManualInsightRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualInsightRunWithDefaults

`func NewManualInsightRunWithDefaults() *ManualInsightRun`

NewManualInsightRunWithDefaults instantiates a new ManualInsightRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualInsightRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualInsightRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualInsightRun) SetId(v string)`

SetId sets Id field to given value.


### GetJobId

`func (o *ManualInsightRun) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ManualInsightRun) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ManualInsightRun) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetVersionId

`func (o *ManualInsightRun) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ManualInsightRun) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ManualInsightRun) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersionName

`func (o *ManualInsightRun) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *ManualInsightRun) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *ManualInsightRun) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.


### GetCreatedAt

`func (o *ManualInsightRun) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManualInsightRun) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManualInsightRun) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ManualInsightRun) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ManualInsightRun) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ManualInsightRun) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetStatus

`func (o *ManualInsightRun) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManualInsightRun) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManualInsightRun) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetIsDeleted

`func (o *ManualInsightRun) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ManualInsightRun) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ManualInsightRun) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetFilters

`func (o *ManualInsightRun) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ManualInsightRun) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ManualInsightRun) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ManualInsightRun) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetRunProcess

`func (o *ManualInsightRun) GetRunProcess() RunProcess`

GetRunProcess returns the RunProcess field if non-nil, zero value otherwise.

### GetRunProcessOk

`func (o *ManualInsightRun) GetRunProcessOk() (*RunProcess, bool)`

GetRunProcessOk returns a tuple with the RunProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunProcess

`func (o *ManualInsightRun) SetRunProcess(v RunProcess)`

SetRunProcess sets RunProcess field to given value.

### HasRunProcess

`func (o *ManualInsightRun) HasRunProcess() bool`

HasRunProcess returns a boolean if a field has been set.

### GetInsightFileUrl

`func (o *ManualInsightRun) GetInsightFileUrl() string`

GetInsightFileUrl returns the InsightFileUrl field if non-nil, zero value otherwise.

### GetInsightFileUrlOk

`func (o *ManualInsightRun) GetInsightFileUrlOk() (*string, bool)`

GetInsightFileUrlOk returns a tuple with the InsightFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightFileUrl

`func (o *ManualInsightRun) SetInsightFileUrl(v string)`

SetInsightFileUrl sets InsightFileUrl field to given value.

### HasInsightFileUrl

`func (o *ManualInsightRun) HasInsightFileUrl() bool`

HasInsightFileUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


