# FetchSimilarResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**JobStatus**](JobStatus.md) |  | 
**JobId** | Pointer to **string** |  | [optional] 
**ReadyArtifacts** | [**FetchSimilarResponseReadyArtifacts**](FetchSimilarResponseReadyArtifacts.md) |  | 

## Methods

### NewFetchSimilarResponse

`func NewFetchSimilarResponse(status JobStatus, readyArtifacts FetchSimilarResponseReadyArtifacts, ) *FetchSimilarResponse`

NewFetchSimilarResponse instantiates a new FetchSimilarResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchSimilarResponseWithDefaults

`func NewFetchSimilarResponseWithDefaults() *FetchSimilarResponse`

NewFetchSimilarResponseWithDefaults instantiates a new FetchSimilarResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FetchSimilarResponse) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FetchSimilarResponse) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FetchSimilarResponse) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetJobId

`func (o *FetchSimilarResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *FetchSimilarResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *FetchSimilarResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *FetchSimilarResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetReadyArtifacts

`func (o *FetchSimilarResponse) GetReadyArtifacts() FetchSimilarResponseReadyArtifacts`

GetReadyArtifacts returns the ReadyArtifacts field if non-nil, zero value otherwise.

### GetReadyArtifactsOk

`func (o *FetchSimilarResponse) GetReadyArtifactsOk() (*FetchSimilarResponseReadyArtifacts, bool)`

GetReadyArtifactsOk returns a tuple with the ReadyArtifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyArtifacts

`func (o *FetchSimilarResponse) SetReadyArtifacts(v FetchSimilarResponseReadyArtifacts)`

SetReadyArtifacts sets ReadyArtifacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


