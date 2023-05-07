# JobNotificationContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**JobType** | [**JobTypeEnum**](JobTypeEnum.md) |  | 
**LeepScriptName** | **string** |  | 
**LeepScriptId** | **string** |  | 
**LeepScriptVersion** | Pointer to **string** |  | [optional] 
**LeepScriptVersionId** | Pointer to **string** |  | [optional] 
**ProjectName** | **string** |  | 
**ProjectId** | **string** |  | 
**ModelName** | **string** |  | 
**ModelExtId** | **string** |  | 
**Sample** | [**SampleIdentity**](SampleIdentity.md) |  | 

## Methods

### NewJobNotificationContext

`func NewJobNotificationContext(jobId string, jobType JobTypeEnum, leepScriptName string, leepScriptId string, projectName string, projectId string, modelName string, modelExtId string, sample SampleIdentity, ) *JobNotificationContext`

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

`func (o *JobNotificationContext) GetJobType() JobTypeEnum`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobNotificationContext) GetJobTypeOk() (*JobTypeEnum, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobNotificationContext) SetJobType(v JobTypeEnum)`

SetJobType sets JobType field to given value.


### GetLeepScriptName

`func (o *JobNotificationContext) GetLeepScriptName() string`

GetLeepScriptName returns the LeepScriptName field if non-nil, zero value otherwise.

### GetLeepScriptNameOk

`func (o *JobNotificationContext) GetLeepScriptNameOk() (*string, bool)`

GetLeepScriptNameOk returns a tuple with the LeepScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeepScriptName

`func (o *JobNotificationContext) SetLeepScriptName(v string)`

SetLeepScriptName sets LeepScriptName field to given value.


### GetLeepScriptId

`func (o *JobNotificationContext) GetLeepScriptId() string`

GetLeepScriptId returns the LeepScriptId field if non-nil, zero value otherwise.

### GetLeepScriptIdOk

`func (o *JobNotificationContext) GetLeepScriptIdOk() (*string, bool)`

GetLeepScriptIdOk returns a tuple with the LeepScriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeepScriptId

`func (o *JobNotificationContext) SetLeepScriptId(v string)`

SetLeepScriptId sets LeepScriptId field to given value.


### GetLeepScriptVersion

`func (o *JobNotificationContext) GetLeepScriptVersion() string`

GetLeepScriptVersion returns the LeepScriptVersion field if non-nil, zero value otherwise.

### GetLeepScriptVersionOk

`func (o *JobNotificationContext) GetLeepScriptVersionOk() (*string, bool)`

GetLeepScriptVersionOk returns a tuple with the LeepScriptVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeepScriptVersion

`func (o *JobNotificationContext) SetLeepScriptVersion(v string)`

SetLeepScriptVersion sets LeepScriptVersion field to given value.

### HasLeepScriptVersion

`func (o *JobNotificationContext) HasLeepScriptVersion() bool`

HasLeepScriptVersion returns a boolean if a field has been set.

### GetLeepScriptVersionId

`func (o *JobNotificationContext) GetLeepScriptVersionId() string`

GetLeepScriptVersionId returns the LeepScriptVersionId field if non-nil, zero value otherwise.

### GetLeepScriptVersionIdOk

`func (o *JobNotificationContext) GetLeepScriptVersionIdOk() (*string, bool)`

GetLeepScriptVersionIdOk returns a tuple with the LeepScriptVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeepScriptVersionId

`func (o *JobNotificationContext) SetLeepScriptVersionId(v string)`

SetLeepScriptVersionId sets LeepScriptVersionId field to given value.

### HasLeepScriptVersionId

`func (o *JobNotificationContext) HasLeepScriptVersionId() bool`

HasLeepScriptVersionId returns a boolean if a field has been set.

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


