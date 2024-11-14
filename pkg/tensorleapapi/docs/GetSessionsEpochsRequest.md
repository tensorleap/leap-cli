# GetSessionsEpochsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SessionIds** | **[]string** |  | 

## Methods

### NewGetSessionsEpochsRequest

`func NewGetSessionsEpochsRequest(projectId string, sessionIds []string, ) *GetSessionsEpochsRequest`

NewGetSessionsEpochsRequest instantiates a new GetSessionsEpochsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSessionsEpochsRequestWithDefaults

`func NewGetSessionsEpochsRequestWithDefaults() *GetSessionsEpochsRequest`

NewGetSessionsEpochsRequestWithDefaults instantiates a new GetSessionsEpochsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GetSessionsEpochsRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetSessionsEpochsRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetSessionsEpochsRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSessionIds

`func (o *GetSessionsEpochsRequest) GetSessionIds() []string`

GetSessionIds returns the SessionIds field if non-nil, zero value otherwise.

### GetSessionIdsOk

`func (o *GetSessionsEpochsRequest) GetSessionIdsOk() (*[]string, bool)`

GetSessionIdsOk returns a tuple with the SessionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIds

`func (o *GetSessionsEpochsRequest) SetSessionIds(v []string)`

SetSessionIds sets SessionIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


