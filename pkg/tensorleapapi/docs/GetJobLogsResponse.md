# GetJobLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodsLogs** | [**[]PodLogs**](PodLogs.md) |  | 

## Methods

### NewGetJobLogsResponse

`func NewGetJobLogsResponse(podsLogs []PodLogs, ) *GetJobLogsResponse`

NewGetJobLogsResponse instantiates a new GetJobLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJobLogsResponseWithDefaults

`func NewGetJobLogsResponseWithDefaults() *GetJobLogsResponse`

NewGetJobLogsResponseWithDefaults instantiates a new GetJobLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodsLogs

`func (o *GetJobLogsResponse) GetPodsLogs() []PodLogs`

GetPodsLogs returns the PodsLogs field if non-nil, zero value otherwise.

### GetPodsLogsOk

`func (o *GetJobLogsResponse) GetPodsLogsOk() (*[]PodLogs, bool)`

GetPodsLogsOk returns a tuple with the PodsLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodsLogs

`func (o *GetJobLogsResponse) SetPodsLogs(v []PodLogs)`

SetPodsLogs sets PodsLogs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


