# SyntheticData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**JobId** | **string** |  | 
**VersionId** | **string** |  | 
**VersionName** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**NextTrialsFileUrl** | Pointer to **string** |  | [optional] 
**BestTrialsFileUrl** | Pointer to **string** |  | [optional] 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**Sources** | [**[]SyntheticDataJobParamsSourcesInner**](SyntheticDataJobParamsSourcesInner.md) |  | 
**TargetFilters** | [**[]ESFilter**](ESFilter.md) |  | 
**RunProcess** | Pointer to [**RunProcess**](RunProcess.md) |  | [optional] 

## Methods

### NewSyntheticData

`func NewSyntheticData(id string, jobId string, versionId string, versionName string, createdAt time.Time, createdBy string, status JobStatus, sources []SyntheticDataJobParamsSourcesInner, targetFilters []ESFilter, ) *SyntheticData`

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


### GetVersionId

`func (o *SyntheticData) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *SyntheticData) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *SyntheticData) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersionName

`func (o *SyntheticData) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *SyntheticData) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *SyntheticData) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.


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

### GetBestTrialsFileUrl

`func (o *SyntheticData) GetBestTrialsFileUrl() string`

GetBestTrialsFileUrl returns the BestTrialsFileUrl field if non-nil, zero value otherwise.

### GetBestTrialsFileUrlOk

`func (o *SyntheticData) GetBestTrialsFileUrlOk() (*string, bool)`

GetBestTrialsFileUrlOk returns a tuple with the BestTrialsFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestTrialsFileUrl

`func (o *SyntheticData) SetBestTrialsFileUrl(v string)`

SetBestTrialsFileUrl sets BestTrialsFileUrl field to given value.

### HasBestTrialsFileUrl

`func (o *SyntheticData) HasBestTrialsFileUrl() bool`

HasBestTrialsFileUrl returns a boolean if a field has been set.

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

`func (o *SyntheticData) GetSources() []SyntheticDataJobParamsSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SyntheticData) GetSourcesOk() (*[]SyntheticDataJobParamsSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SyntheticData) SetSources(v []SyntheticDataJobParamsSourcesInner)`

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


### GetRunProcess

`func (o *SyntheticData) GetRunProcess() RunProcess`

GetRunProcess returns the RunProcess field if non-nil, zero value otherwise.

### GetRunProcessOk

`func (o *SyntheticData) GetRunProcessOk() (*RunProcess, bool)`

GetRunProcessOk returns a tuple with the RunProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunProcess

`func (o *SyntheticData) SetRunProcess(v RunProcess)`

SetRunProcess sets RunProcess field to given value.

### HasRunProcess

`func (o *SyntheticData) HasRunProcess() bool`

HasRunProcess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


