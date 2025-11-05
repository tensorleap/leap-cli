# RunProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**ProcessId** | **string** |  | 
**VersionName** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**DatasetName** | Pointer to **string** |  | [optional] 
**JobType** | **string** |  | 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**SessionName** | Pointer to **string** |  | [optional] 
**SessionRunName** | Pointer to **string** |  | [optional] 
**SessionRunId** | Pointer to **string** |  | [optional] 
**Events** | [**[]JobEvent**](JobEvent.md) |  | 
**Params** | Pointer to [**JobParams**](JobParams.md) |  | [optional] 
**MachineType** | Pointer to **string** |  | [optional] 
**BatchSize** | Pointer to **float64** |  | [optional] 
**CodeSnapshotInfo** | Pointer to [**CodeSnapshotInfo**](CodeSnapshotInfo.md) |  | [optional] 
**LogsBlobName** | Pointer to **string** |  | [optional] 

## Methods

### NewRunProcess

`func NewRunProcess(jobId string, processId string, jobType string, status JobStatus, createdAt string, updatedAt string, events []JobEvent, ) *RunProcess`

NewRunProcess instantiates a new RunProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunProcessWithDefaults

`func NewRunProcessWithDefaults() *RunProcess`

NewRunProcessWithDefaults instantiates a new RunProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *RunProcess) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *RunProcess) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *RunProcess) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetProcessId

`func (o *RunProcess) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *RunProcess) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *RunProcess) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetVersionName

`func (o *RunProcess) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *RunProcess) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *RunProcess) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *RunProcess) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### GetVersionId

`func (o *RunProcess) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *RunProcess) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *RunProcess) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *RunProcess) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetProjectId

`func (o *RunProcess) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RunProcess) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RunProcess) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RunProcess) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *RunProcess) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *RunProcess) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *RunProcess) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *RunProcess) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetDatasetName

`func (o *RunProcess) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *RunProcess) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *RunProcess) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.

### HasDatasetName

`func (o *RunProcess) HasDatasetName() bool`

HasDatasetName returns a boolean if a field has been set.

### GetJobType

`func (o *RunProcess) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *RunProcess) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *RunProcess) SetJobType(v string)`

SetJobType sets JobType field to given value.


### GetStatus

`func (o *RunProcess) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RunProcess) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RunProcess) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *RunProcess) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RunProcess) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RunProcess) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RunProcess) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RunProcess) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RunProcess) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSessionName

`func (o *RunProcess) GetSessionName() string`

GetSessionName returns the SessionName field if non-nil, zero value otherwise.

### GetSessionNameOk

`func (o *RunProcess) GetSessionNameOk() (*string, bool)`

GetSessionNameOk returns a tuple with the SessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionName

`func (o *RunProcess) SetSessionName(v string)`

SetSessionName sets SessionName field to given value.

### HasSessionName

`func (o *RunProcess) HasSessionName() bool`

HasSessionName returns a boolean if a field has been set.

### GetSessionRunName

`func (o *RunProcess) GetSessionRunName() string`

GetSessionRunName returns the SessionRunName field if non-nil, zero value otherwise.

### GetSessionRunNameOk

`func (o *RunProcess) GetSessionRunNameOk() (*string, bool)`

GetSessionRunNameOk returns a tuple with the SessionRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunName

`func (o *RunProcess) SetSessionRunName(v string)`

SetSessionRunName sets SessionRunName field to given value.

### HasSessionRunName

`func (o *RunProcess) HasSessionRunName() bool`

HasSessionRunName returns a boolean if a field has been set.

### GetSessionRunId

`func (o *RunProcess) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *RunProcess) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *RunProcess) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.

### HasSessionRunId

`func (o *RunProcess) HasSessionRunId() bool`

HasSessionRunId returns a boolean if a field has been set.

### GetEvents

`func (o *RunProcess) GetEvents() []JobEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RunProcess) GetEventsOk() (*[]JobEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RunProcess) SetEvents(v []JobEvent)`

SetEvents sets Events field to given value.


### GetParams

`func (o *RunProcess) GetParams() JobParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *RunProcess) GetParamsOk() (*JobParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *RunProcess) SetParams(v JobParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *RunProcess) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetMachineType

`func (o *RunProcess) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *RunProcess) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *RunProcess) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *RunProcess) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### GetBatchSize

`func (o *RunProcess) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *RunProcess) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *RunProcess) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *RunProcess) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetCodeSnapshotInfo

`func (o *RunProcess) GetCodeSnapshotInfo() CodeSnapshotInfo`

GetCodeSnapshotInfo returns the CodeSnapshotInfo field if non-nil, zero value otherwise.

### GetCodeSnapshotInfoOk

`func (o *RunProcess) GetCodeSnapshotInfoOk() (*CodeSnapshotInfo, bool)`

GetCodeSnapshotInfoOk returns a tuple with the CodeSnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSnapshotInfo

`func (o *RunProcess) SetCodeSnapshotInfo(v CodeSnapshotInfo)`

SetCodeSnapshotInfo sets CodeSnapshotInfo field to given value.

### HasCodeSnapshotInfo

`func (o *RunProcess) HasCodeSnapshotInfo() bool`

HasCodeSnapshotInfo returns a boolean if a field has been set.

### GetLogsBlobName

`func (o *RunProcess) GetLogsBlobName() string`

GetLogsBlobName returns the LogsBlobName field if non-nil, zero value otherwise.

### GetLogsBlobNameOk

`func (o *RunProcess) GetLogsBlobNameOk() (*string, bool)`

GetLogsBlobNameOk returns a tuple with the LogsBlobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsBlobName

`func (o *RunProcess) SetLogsBlobName(v string)`

SetLogsBlobName sets LogsBlobName field to given value.

### HasLogsBlobName

`func (o *RunProcess) HasLogsBlobName() bool`

HasLogsBlobName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


