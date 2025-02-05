# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** |  | [optional] 
**Cid** | **string** |  | 
**CreatedBy** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 
**Type** | [**JobType**](JobType.md) |  | 
**SubType** | Pointer to [**JobSubType**](JobSubType.md) |  | [optional] 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Params** | Pointer to [**JobParams**](JobParams.md) |  | [optional] 
**SessionRunId** | Pointer to **string** |  | [optional] 
**TeamId** | **string** |  | 
**DatasetVersionInfo** | Pointer to [**DatasetVersionInfo**](DatasetVersionInfo.md) |  | [optional] 
**LogsBlobName** | Pointer to **string** |  | [optional] 

## Methods

### NewJob

`func NewJob(cid string, createdBy string, type_ JobType, status JobStatus, createdAt time.Time, updatedAt time.Time, teamId string, ) *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *Job) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Job) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Job) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Job) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetCid

`func (o *Job) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *Job) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *Job) SetCid(v string)`

SetCid sets Cid field to given value.


### GetCreatedBy

`func (o *Job) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Job) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Job) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetVersion

`func (o *Job) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Job) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Job) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Job) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *Job) GetType() JobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Job) GetTypeOk() (*JobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Job) SetType(v JobType)`

SetType sets Type field to given value.


### GetSubType

`func (o *Job) GetSubType() JobSubType`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *Job) GetSubTypeOk() (*JobSubType, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *Job) SetSubType(v JobSubType)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *Job) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetStatus

`func (o *Job) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *Job) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Job) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Job) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Job) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Job) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Job) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetParams

`func (o *Job) GetParams() JobParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Job) GetParamsOk() (*JobParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Job) SetParams(v JobParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *Job) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetSessionRunId

`func (o *Job) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *Job) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *Job) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.

### HasSessionRunId

`func (o *Job) HasSessionRunId() bool`

HasSessionRunId returns a boolean if a field has been set.

### GetTeamId

`func (o *Job) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Job) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Job) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetDatasetVersionInfo

`func (o *Job) GetDatasetVersionInfo() DatasetVersionInfo`

GetDatasetVersionInfo returns the DatasetVersionInfo field if non-nil, zero value otherwise.

### GetDatasetVersionInfoOk

`func (o *Job) GetDatasetVersionInfoOk() (*DatasetVersionInfo, bool)`

GetDatasetVersionInfoOk returns a tuple with the DatasetVersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetVersionInfo

`func (o *Job) SetDatasetVersionInfo(v DatasetVersionInfo)`

SetDatasetVersionInfo sets DatasetVersionInfo field to given value.

### HasDatasetVersionInfo

`func (o *Job) HasDatasetVersionInfo() bool`

HasDatasetVersionInfo returns a boolean if a field has been set.

### GetLogsBlobName

`func (o *Job) GetLogsBlobName() string`

GetLogsBlobName returns the LogsBlobName field if non-nil, zero value otherwise.

### GetLogsBlobNameOk

`func (o *Job) GetLogsBlobNameOk() (*string, bool)`

GetLogsBlobNameOk returns a tuple with the LogsBlobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsBlobName

`func (o *Job) SetLogsBlobName(v string)`

SetLogsBlobName sets LogsBlobName field to given value.

### HasLogsBlobName

`func (o *Job) HasLogsBlobName() bool`

HasLogsBlobName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


