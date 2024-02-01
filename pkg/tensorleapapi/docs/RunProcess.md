# RunProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**VersionName** | **string** |  | 
**VersionId** | **string** |  | 
**ProjectName** | **string** |  | 
**JobType** | **string** |  | 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**SessionName** | **string** |  | 
**SessionRunName** | **string** |  | 
**SessionRunId** | Pointer to **string** |  | [optional] 
**Events** | [**[]JobEvent**](JobEvent.md) |  | 
**Params** | Pointer to [**JobParams**](JobParams.md) |  | [optional] 

## Methods

### NewRunProcess

`func NewRunProcess(jobId string, versionName string, versionId string, projectName string, jobType string, status JobStatus, createdAt string, updatedAt string, sessionName string, sessionRunName string, events []JobEvent, ) *RunProcess`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


