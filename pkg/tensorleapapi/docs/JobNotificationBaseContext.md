# JobNotificationBaseContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**JobType** | [**JobTypeEnum**](JobTypeEnum.md) |  | 

## Methods

### NewJobNotificationBaseContext

`func NewJobNotificationBaseContext(jobId string, jobType JobTypeEnum, ) *JobNotificationBaseContext`

NewJobNotificationBaseContext instantiates a new JobNotificationBaseContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobNotificationBaseContextWithDefaults

`func NewJobNotificationBaseContextWithDefaults() *JobNotificationBaseContext`

NewJobNotificationBaseContextWithDefaults instantiates a new JobNotificationBaseContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *JobNotificationBaseContext) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobNotificationBaseContext) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobNotificationBaseContext) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetJobType

`func (o *JobNotificationBaseContext) GetJobType() JobTypeEnum`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobNotificationBaseContext) GetJobTypeOk() (*JobTypeEnum, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobNotificationBaseContext) SetJobType(v JobTypeEnum)`

SetJobType sets JobType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


