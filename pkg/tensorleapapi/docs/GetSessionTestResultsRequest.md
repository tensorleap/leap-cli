# GetSessionTestResultsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SessionsTestData** | [**[]SessionTestData**](SessionTestData.md) |  | 

## Methods

### NewGetSessionTestResultsRequest

`func NewGetSessionTestResultsRequest(projectId string, sessionsTestData []SessionTestData, ) *GetSessionTestResultsRequest`

NewGetSessionTestResultsRequest instantiates a new GetSessionTestResultsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSessionTestResultsRequestWithDefaults

`func NewGetSessionTestResultsRequestWithDefaults() *GetSessionTestResultsRequest`

NewGetSessionTestResultsRequestWithDefaults instantiates a new GetSessionTestResultsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GetSessionTestResultsRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetSessionTestResultsRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetSessionTestResultsRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSessionsTestData

`func (o *GetSessionTestResultsRequest) GetSessionsTestData() []SessionTestData`

GetSessionsTestData returns the SessionsTestData field if non-nil, zero value otherwise.

### GetSessionsTestDataOk

`func (o *GetSessionTestResultsRequest) GetSessionsTestDataOk() (*[]SessionTestData, bool)`

GetSessionsTestDataOk returns a tuple with the SessionsTestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsTestData

`func (o *GetSessionTestResultsRequest) SetSessionsTestData(v []SessionTestData)`

SetSessionsTestData sets SessionsTestData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


