# TerminateAllJobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**JobIds** | **[]string** |  | 

## Methods

### NewTerminateAllJobsResponse

`func NewTerminateAllJobsResponse(status string, jobIds []string, ) *TerminateAllJobsResponse`

NewTerminateAllJobsResponse instantiates a new TerminateAllJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminateAllJobsResponseWithDefaults

`func NewTerminateAllJobsResponseWithDefaults() *TerminateAllJobsResponse`

NewTerminateAllJobsResponseWithDefaults instantiates a new TerminateAllJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TerminateAllJobsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TerminateAllJobsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TerminateAllJobsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetJobIds

`func (o *TerminateAllJobsResponse) GetJobIds() []string`

GetJobIds returns the JobIds field if non-nil, zero value otherwise.

### GetJobIdsOk

`func (o *TerminateAllJobsResponse) GetJobIdsOk() (*[]string, bool)`

GetJobIdsOk returns a tuple with the JobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIds

`func (o *TerminateAllJobsResponse) SetJobIds(v []string)`

SetJobIds sets JobIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


