# JobNotificationProjectContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**JobType** | [**JobType**](JobType.md) |  | 
**JobSubType** | Pointer to [**JobSubType**](JobSubType.md) |  | [optional] 
**ProjectName** | **string** |  | 

## Methods

### NewJobNotificationProjectContext

`func NewJobNotificationProjectContext(jobId string, jobType JobType, projectName string, ) *JobNotificationProjectContext`

NewJobNotificationProjectContext instantiates a new JobNotificationProjectContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobNotificationProjectContextWithDefaults

`func NewJobNotificationProjectContextWithDefaults() *JobNotificationProjectContext`

NewJobNotificationProjectContextWithDefaults instantiates a new JobNotificationProjectContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *JobNotificationProjectContext) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobNotificationProjectContext) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobNotificationProjectContext) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetJobType

`func (o *JobNotificationProjectContext) GetJobType() JobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobNotificationProjectContext) GetJobTypeOk() (*JobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobNotificationProjectContext) SetJobType(v JobType)`

SetJobType sets JobType field to given value.


### GetJobSubType

`func (o *JobNotificationProjectContext) GetJobSubType() JobSubType`

GetJobSubType returns the JobSubType field if non-nil, zero value otherwise.

### GetJobSubTypeOk

`func (o *JobNotificationProjectContext) GetJobSubTypeOk() (*JobSubType, bool)`

GetJobSubTypeOk returns a tuple with the JobSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSubType

`func (o *JobNotificationProjectContext) SetJobSubType(v JobSubType)`

SetJobSubType sets JobSubType field to given value.

### HasJobSubType

`func (o *JobNotificationProjectContext) HasJobSubType() bool`

HasJobSubType returns a boolean if a field has been set.

### GetProjectName

`func (o *JobNotificationProjectContext) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *JobNotificationProjectContext) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *JobNotificationProjectContext) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


