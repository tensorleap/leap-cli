# JobNotificationContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**JobType** | [**JobType**](JobType.md) |  | 
**CodeSnapshotId** | **string** |  | 
**VersionId** | **string** |  | 
**VersionName** | **string** |  | 
**ProjectName** | **string** |  | 
**ProjectId** | **string** |  | 
**ModelName** | **string** |  | 
**ModelExtId** | **string** |  | 
**IsOverwrite** | **bool** |  | 
**SessionId** | Pointer to **string** |  | [optional] 
**Epoch** | Pointer to **float64** |  | [optional] 
**SessionRunId** | **string** |  | 
**Sample** | [**SampleIdentity**](SampleIdentity.md) |  | 

## Methods

### NewJobNotificationContext

`func NewJobNotificationContext(jobId string, jobType JobType, codeSnapshotId string, versionId string, versionName string, projectName string, projectId string, modelName string, modelExtId string, isOverwrite bool, sessionRunId string, sample SampleIdentity, ) *JobNotificationContext`

NewJobNotificationContext instantiates a new JobNotificationContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobNotificationContextWithDefaults

`func NewJobNotificationContextWithDefaults() *JobNotificationContext`

NewJobNotificationContextWithDefaults instantiates a new JobNotificationContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *JobNotificationContext) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobNotificationContext) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobNotificationContext) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetJobType

`func (o *JobNotificationContext) GetJobType() JobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobNotificationContext) GetJobTypeOk() (*JobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobNotificationContext) SetJobType(v JobType)`

SetJobType sets JobType field to given value.


### GetCodeSnapshotId

`func (o *JobNotificationContext) GetCodeSnapshotId() string`

GetCodeSnapshotId returns the CodeSnapshotId field if non-nil, zero value otherwise.

### GetCodeSnapshotIdOk

`func (o *JobNotificationContext) GetCodeSnapshotIdOk() (*string, bool)`

GetCodeSnapshotIdOk returns a tuple with the CodeSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSnapshotId

`func (o *JobNotificationContext) SetCodeSnapshotId(v string)`

SetCodeSnapshotId sets CodeSnapshotId field to given value.


### GetVersionId

`func (o *JobNotificationContext) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *JobNotificationContext) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *JobNotificationContext) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersionName

`func (o *JobNotificationContext) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *JobNotificationContext) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *JobNotificationContext) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.


### GetProjectName

`func (o *JobNotificationContext) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *JobNotificationContext) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *JobNotificationContext) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetProjectId

`func (o *JobNotificationContext) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *JobNotificationContext) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *JobNotificationContext) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetModelName

`func (o *JobNotificationContext) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *JobNotificationContext) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *JobNotificationContext) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetModelExtId

`func (o *JobNotificationContext) GetModelExtId() string`

GetModelExtId returns the ModelExtId field if non-nil, zero value otherwise.

### GetModelExtIdOk

`func (o *JobNotificationContext) GetModelExtIdOk() (*string, bool)`

GetModelExtIdOk returns a tuple with the ModelExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelExtId

`func (o *JobNotificationContext) SetModelExtId(v string)`

SetModelExtId sets ModelExtId field to given value.


### GetIsOverwrite

`func (o *JobNotificationContext) GetIsOverwrite() bool`

GetIsOverwrite returns the IsOverwrite field if non-nil, zero value otherwise.

### GetIsOverwriteOk

`func (o *JobNotificationContext) GetIsOverwriteOk() (*bool, bool)`

GetIsOverwriteOk returns a tuple with the IsOverwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverwrite

`func (o *JobNotificationContext) SetIsOverwrite(v bool)`

SetIsOverwrite sets IsOverwrite field to given value.


### GetSessionId

`func (o *JobNotificationContext) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *JobNotificationContext) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *JobNotificationContext) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *JobNotificationContext) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetEpoch

`func (o *JobNotificationContext) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *JobNotificationContext) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *JobNotificationContext) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *JobNotificationContext) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetSessionRunId

`func (o *JobNotificationContext) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *JobNotificationContext) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *JobNotificationContext) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetSample

`func (o *JobNotificationContext) GetSample() SampleIdentity`

GetSample returns the Sample field if non-nil, zero value otherwise.

### GetSampleOk

`func (o *JobNotificationContext) GetSampleOk() (*SampleIdentity, bool)`

GetSampleOk returns a tuple with the Sample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample

`func (o *JobNotificationContext) SetSample(v SampleIdentity)`

SetSample sets Sample field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


