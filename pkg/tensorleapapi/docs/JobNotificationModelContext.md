# JobNotificationModelContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**JobType** | [**JobType**](JobType.md) |  | 
**ProjectName** | **string** |  | 
**ProjectId** | **string** |  | 
**ModelName** | **string** |  | 
**ModelExtId** | **string** |  | 
**IsOverwrite** | **bool** |  | 
**VersionId** | **string** |  | 
**SessionId** | Pointer to **string** |  | [optional] 
**Epoch** | Pointer to **float64** |  | [optional] 

## Methods

### NewJobNotificationModelContext

`func NewJobNotificationModelContext(jobId string, jobType JobType, projectName string, projectId string, modelName string, modelExtId string, isOverwrite bool, versionId string, ) *JobNotificationModelContext`

NewJobNotificationModelContext instantiates a new JobNotificationModelContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobNotificationModelContextWithDefaults

`func NewJobNotificationModelContextWithDefaults() *JobNotificationModelContext`

NewJobNotificationModelContextWithDefaults instantiates a new JobNotificationModelContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *JobNotificationModelContext) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobNotificationModelContext) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobNotificationModelContext) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetJobType

`func (o *JobNotificationModelContext) GetJobType() JobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobNotificationModelContext) GetJobTypeOk() (*JobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobNotificationModelContext) SetJobType(v JobType)`

SetJobType sets JobType field to given value.


### GetProjectName

`func (o *JobNotificationModelContext) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *JobNotificationModelContext) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *JobNotificationModelContext) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetProjectId

`func (o *JobNotificationModelContext) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *JobNotificationModelContext) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *JobNotificationModelContext) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetModelName

`func (o *JobNotificationModelContext) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *JobNotificationModelContext) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *JobNotificationModelContext) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetModelExtId

`func (o *JobNotificationModelContext) GetModelExtId() string`

GetModelExtId returns the ModelExtId field if non-nil, zero value otherwise.

### GetModelExtIdOk

`func (o *JobNotificationModelContext) GetModelExtIdOk() (*string, bool)`

GetModelExtIdOk returns a tuple with the ModelExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelExtId

`func (o *JobNotificationModelContext) SetModelExtId(v string)`

SetModelExtId sets ModelExtId field to given value.


### GetIsOverwrite

`func (o *JobNotificationModelContext) GetIsOverwrite() bool`

GetIsOverwrite returns the IsOverwrite field if non-nil, zero value otherwise.

### GetIsOverwriteOk

`func (o *JobNotificationModelContext) GetIsOverwriteOk() (*bool, bool)`

GetIsOverwriteOk returns a tuple with the IsOverwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverwrite

`func (o *JobNotificationModelContext) SetIsOverwrite(v bool)`

SetIsOverwrite sets IsOverwrite field to given value.


### GetVersionId

`func (o *JobNotificationModelContext) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *JobNotificationModelContext) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *JobNotificationModelContext) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetSessionId

`func (o *JobNotificationModelContext) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *JobNotificationModelContext) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *JobNotificationModelContext) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *JobNotificationModelContext) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetEpoch

`func (o *JobNotificationModelContext) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *JobNotificationModelContext) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *JobNotificationModelContext) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *JobNotificationModelContext) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


