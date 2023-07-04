# GetJobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processes** | [**[]RunProcess**](RunProcess.md) |  | 

## Methods

### NewGetJobsResponse

`func NewGetJobsResponse(processes []RunProcess, ) *GetJobsResponse`

NewGetJobsResponse instantiates a new GetJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJobsResponseWithDefaults

`func NewGetJobsResponseWithDefaults() *GetJobsResponse`

NewGetJobsResponseWithDefaults instantiates a new GetJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcesses

`func (o *GetJobsResponse) GetProcesses() []RunProcess`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *GetJobsResponse) GetProcessesOk() (*[]RunProcess, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *GetJobsResponse) SetProcesses(v []RunProcess)`

SetProcesses sets Processes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


