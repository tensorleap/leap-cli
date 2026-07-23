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
**BundleFileUrl** | Pointer to **string** |  | [optional] 
**FilterFileUrl** | Pointer to **string** |  | [optional] 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**Sources** | [**[]SyntheticDataJobParamsSourcesInner**](SyntheticDataJobParamsSourcesInner.md) |  | 
**TargetFilters** | [**[]ESFilter**](ESFilter.md) |  | 
**IsAuto** | Pointer to **bool** |  | [optional] 
**SimulationNames** | Pointer to **[]string** |  | [optional] 
**InitialSimulationFilters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
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

### GetBundleFileUrl

`func (o *SyntheticData) GetBundleFileUrl() string`

GetBundleFileUrl returns the BundleFileUrl field if non-nil, zero value otherwise.

### GetBundleFileUrlOk

`func (o *SyntheticData) GetBundleFileUrlOk() (*string, bool)`

GetBundleFileUrlOk returns a tuple with the BundleFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleFileUrl

`func (o *SyntheticData) SetBundleFileUrl(v string)`

SetBundleFileUrl sets BundleFileUrl field to given value.

### HasBundleFileUrl

`func (o *SyntheticData) HasBundleFileUrl() bool`

HasBundleFileUrl returns a boolean if a field has been set.

### GetFilterFileUrl

`func (o *SyntheticData) GetFilterFileUrl() string`

GetFilterFileUrl returns the FilterFileUrl field if non-nil, zero value otherwise.

### GetFilterFileUrlOk

`func (o *SyntheticData) GetFilterFileUrlOk() (*string, bool)`

GetFilterFileUrlOk returns a tuple with the FilterFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterFileUrl

`func (o *SyntheticData) SetFilterFileUrl(v string)`

SetFilterFileUrl sets FilterFileUrl field to given value.

### HasFilterFileUrl

`func (o *SyntheticData) HasFilterFileUrl() bool`

HasFilterFileUrl returns a boolean if a field has been set.

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


### GetIsAuto

`func (o *SyntheticData) GetIsAuto() bool`

GetIsAuto returns the IsAuto field if non-nil, zero value otherwise.

### GetIsAutoOk

`func (o *SyntheticData) GetIsAutoOk() (*bool, bool)`

GetIsAutoOk returns a tuple with the IsAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuto

`func (o *SyntheticData) SetIsAuto(v bool)`

SetIsAuto sets IsAuto field to given value.

### HasIsAuto

`func (o *SyntheticData) HasIsAuto() bool`

HasIsAuto returns a boolean if a field has been set.

### GetSimulationNames

`func (o *SyntheticData) GetSimulationNames() []string`

GetSimulationNames returns the SimulationNames field if non-nil, zero value otherwise.

### GetSimulationNamesOk

`func (o *SyntheticData) GetSimulationNamesOk() (*[]string, bool)`

GetSimulationNamesOk returns a tuple with the SimulationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationNames

`func (o *SyntheticData) SetSimulationNames(v []string)`

SetSimulationNames sets SimulationNames field to given value.

### HasSimulationNames

`func (o *SyntheticData) HasSimulationNames() bool`

HasSimulationNames returns a boolean if a field has been set.

### GetInitialSimulationFilters

`func (o *SyntheticData) GetInitialSimulationFilters() []ESFilter`

GetInitialSimulationFilters returns the InitialSimulationFilters field if non-nil, zero value otherwise.

### GetInitialSimulationFiltersOk

`func (o *SyntheticData) GetInitialSimulationFiltersOk() (*[]ESFilter, bool)`

GetInitialSimulationFiltersOk returns a tuple with the InitialSimulationFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSimulationFilters

`func (o *SyntheticData) SetInitialSimulationFilters(v []ESFilter)`

SetInitialSimulationFilters sets InitialSimulationFilters field to given value.

### HasInitialSimulationFilters

`func (o *SyntheticData) HasInitialSimulationFilters() bool`

HasInitialSimulationFilters returns a boolean if a field has been set.

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


