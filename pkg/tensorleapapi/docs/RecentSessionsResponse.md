# RecentSessionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sessions** | [**[]SessionPopulatedJob**](SessionPopulatedJob.md) |  | 

## Methods

### NewRecentSessionsResponse

`func NewRecentSessionsResponse(sessions []SessionPopulatedJob, ) *RecentSessionsResponse`

NewRecentSessionsResponse instantiates a new RecentSessionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecentSessionsResponseWithDefaults

`func NewRecentSessionsResponseWithDefaults() *RecentSessionsResponse`

NewRecentSessionsResponseWithDefaults instantiates a new RecentSessionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessions

`func (o *RecentSessionsResponse) GetSessions() []SessionPopulatedJob`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *RecentSessionsResponse) GetSessionsOk() (*[]SessionPopulatedJob, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *RecentSessionsResponse) SetSessions(v []SessionPopulatedJob)`

SetSessions sets Sessions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


