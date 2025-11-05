# JobNotificationSampleContext

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
**Sample** | [**SampleIdentity**](SampleIdentity.md) |  | 

## Methods

### NewJobNotificationSampleContext

`func NewJobNotificationSampleContext(jobId string, jobType JobType, projectName string, projectId string, modelName string, modelExtId string, isOverwrite bool, versionId string, sample SampleIdentity, ) *JobNotificationSampleContext`

NewJobNotificationSampleContext instantiates a new JobNotificationSampleContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobNotificationSampleContextWithDefaults

`func NewJobNotificationSampleContextWithDefaults() *JobNotificationSampleContext`

NewJobNotificationSampleContextWithDefaults instantiates a new JobNotificationSampleContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *JobNotificationSampleContext) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobNotificationSampleContext) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobNotificationSampleContext) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetJobType

`func (o *JobNotificationSampleContext) GetJobType() JobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobNotificationSampleContext) GetJobTypeOk() (*JobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobNotificationSampleContext) SetJobType(v JobType)`

SetJobType sets JobType field to given value.


### GetProjectName

`func (o *JobNotificationSampleContext) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *JobNotificationSampleContext) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *JobNotificationSampleContext) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetProjectId

`func (o *JobNotificationSampleContext) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *JobNotificationSampleContext) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *JobNotificationSampleContext) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetModelName

`func (o *JobNotificationSampleContext) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *JobNotificationSampleContext) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *JobNotificationSampleContext) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetModelExtId

`func (o *JobNotificationSampleContext) GetModelExtId() string`

GetModelExtId returns the ModelExtId field if non-nil, zero value otherwise.

### GetModelExtIdOk

`func (o *JobNotificationSampleContext) GetModelExtIdOk() (*string, bool)`

GetModelExtIdOk returns a tuple with the ModelExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelExtId

`func (o *JobNotificationSampleContext) SetModelExtId(v string)`

SetModelExtId sets ModelExtId field to given value.


### GetIsOverwrite

`func (o *JobNotificationSampleContext) GetIsOverwrite() bool`

GetIsOverwrite returns the IsOverwrite field if non-nil, zero value otherwise.

### GetIsOverwriteOk

`func (o *JobNotificationSampleContext) GetIsOverwriteOk() (*bool, bool)`

GetIsOverwriteOk returns a tuple with the IsOverwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverwrite

`func (o *JobNotificationSampleContext) SetIsOverwrite(v bool)`

SetIsOverwrite sets IsOverwrite field to given value.


### GetVersionId

`func (o *JobNotificationSampleContext) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *JobNotificationSampleContext) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *JobNotificationSampleContext) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetSessionId

`func (o *JobNotificationSampleContext) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *JobNotificationSampleContext) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *JobNotificationSampleContext) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *JobNotificationSampleContext) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetEpoch

`func (o *JobNotificationSampleContext) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *JobNotificationSampleContext) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *JobNotificationSampleContext) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *JobNotificationSampleContext) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetSample

`func (o *JobNotificationSampleContext) GetSample() SampleIdentity`

GetSample returns the Sample field if non-nil, zero value otherwise.

### GetSampleOk

`func (o *JobNotificationSampleContext) GetSampleOk() (*SampleIdentity, bool)`

GetSampleOk returns a tuple with the Sample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample

`func (o *JobNotificationSampleContext) SetSample(v SampleIdentity)`

SetSample sets Sample field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


