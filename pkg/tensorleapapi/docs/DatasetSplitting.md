# DatasetSplitting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**JobId** | **string** |  | 
**VersionId** | **string** |  | 
**VersionName** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**FileUrl** | Pointer to **string** |  | [optional] 
**FilterFileUrl** | Pointer to **string** |  | [optional] 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**IsDeleted** | **bool** |  | 
**SplitsToResplit** | [**[]SplitSubset**](SplitSubset.md) |  | 
**KeepTogetherMetadata** | **[]string** |  | 
**SplitAcrossMetadata** | **[]string** |  | 
**RunProcess** | Pointer to [**RunProcess**](RunProcess.md) |  | [optional] 

## Methods

### NewDatasetSplitting

`func NewDatasetSplitting(id string, jobId string, versionId string, versionName string, createdAt time.Time, createdBy string, status JobStatus, isDeleted bool, splitsToResplit []SplitSubset, keepTogetherMetadata []string, splitAcrossMetadata []string, ) *DatasetSplitting`

NewDatasetSplitting instantiates a new DatasetSplitting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetSplittingWithDefaults

`func NewDatasetSplittingWithDefaults() *DatasetSplitting`

NewDatasetSplittingWithDefaults instantiates a new DatasetSplitting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatasetSplitting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatasetSplitting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatasetSplitting) SetId(v string)`

SetId sets Id field to given value.


### GetJobId

`func (o *DatasetSplitting) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *DatasetSplitting) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *DatasetSplitting) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetVersionId

`func (o *DatasetSplitting) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *DatasetSplitting) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *DatasetSplitting) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersionName

`func (o *DatasetSplitting) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *DatasetSplitting) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *DatasetSplitting) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.


### GetCreatedAt

`func (o *DatasetSplitting) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatasetSplitting) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatasetSplitting) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *DatasetSplitting) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DatasetSplitting) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DatasetSplitting) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetFileUrl

`func (o *DatasetSplitting) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *DatasetSplitting) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *DatasetSplitting) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *DatasetSplitting) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetFilterFileUrl

`func (o *DatasetSplitting) GetFilterFileUrl() string`

GetFilterFileUrl returns the FilterFileUrl field if non-nil, zero value otherwise.

### GetFilterFileUrlOk

`func (o *DatasetSplitting) GetFilterFileUrlOk() (*string, bool)`

GetFilterFileUrlOk returns a tuple with the FilterFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterFileUrl

`func (o *DatasetSplitting) SetFilterFileUrl(v string)`

SetFilterFileUrl sets FilterFileUrl field to given value.

### HasFilterFileUrl

`func (o *DatasetSplitting) HasFilterFileUrl() bool`

HasFilterFileUrl returns a boolean if a field has been set.

### GetStatus

`func (o *DatasetSplitting) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatasetSplitting) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatasetSplitting) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetIsDeleted

`func (o *DatasetSplitting) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *DatasetSplitting) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *DatasetSplitting) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetSplitsToResplit

`func (o *DatasetSplitting) GetSplitsToResplit() []SplitSubset`

GetSplitsToResplit returns the SplitsToResplit field if non-nil, zero value otherwise.

### GetSplitsToResplitOk

`func (o *DatasetSplitting) GetSplitsToResplitOk() (*[]SplitSubset, bool)`

GetSplitsToResplitOk returns a tuple with the SplitsToResplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitsToResplit

`func (o *DatasetSplitting) SetSplitsToResplit(v []SplitSubset)`

SetSplitsToResplit sets SplitsToResplit field to given value.


### GetKeepTogetherMetadata

`func (o *DatasetSplitting) GetKeepTogetherMetadata() []string`

GetKeepTogetherMetadata returns the KeepTogetherMetadata field if non-nil, zero value otherwise.

### GetKeepTogetherMetadataOk

`func (o *DatasetSplitting) GetKeepTogetherMetadataOk() (*[]string, bool)`

GetKeepTogetherMetadataOk returns a tuple with the KeepTogetherMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepTogetherMetadata

`func (o *DatasetSplitting) SetKeepTogetherMetadata(v []string)`

SetKeepTogetherMetadata sets KeepTogetherMetadata field to given value.


### GetSplitAcrossMetadata

`func (o *DatasetSplitting) GetSplitAcrossMetadata() []string`

GetSplitAcrossMetadata returns the SplitAcrossMetadata field if non-nil, zero value otherwise.

### GetSplitAcrossMetadataOk

`func (o *DatasetSplitting) GetSplitAcrossMetadataOk() (*[]string, bool)`

GetSplitAcrossMetadataOk returns a tuple with the SplitAcrossMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitAcrossMetadata

`func (o *DatasetSplitting) SetSplitAcrossMetadata(v []string)`

SetSplitAcrossMetadata sets SplitAcrossMetadata field to given value.


### GetRunProcess

`func (o *DatasetSplitting) GetRunProcess() RunProcess`

GetRunProcess returns the RunProcess field if non-nil, zero value otherwise.

### GetRunProcessOk

`func (o *DatasetSplitting) GetRunProcessOk() (*RunProcess, bool)`

GetRunProcessOk returns a tuple with the RunProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunProcess

`func (o *DatasetSplitting) SetRunProcess(v RunProcess)`

SetRunProcess sets RunProcess field to given value.

### HasRunProcess

`func (o *DatasetSplitting) HasRunProcess() bool`

HasRunProcess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


