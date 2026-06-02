# PushResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**VersionId** | **string** |  | 
**CodeSnapshot** | [**CodeSnapshot**](CodeSnapshot.md) |  | 

## Methods

### NewPushResponse

`func NewPushResponse(jobId string, versionId string, codeSnapshot CodeSnapshot, ) *PushResponse`

NewPushResponse instantiates a new PushResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushResponseWithDefaults

`func NewPushResponseWithDefaults() *PushResponse`

NewPushResponseWithDefaults instantiates a new PushResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *PushResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *PushResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *PushResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetVersionId

`func (o *PushResponse) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *PushResponse) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *PushResponse) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetCodeSnapshot

`func (o *PushResponse) GetCodeSnapshot() CodeSnapshot`

GetCodeSnapshot returns the CodeSnapshot field if non-nil, zero value otherwise.

### GetCodeSnapshotOk

`func (o *PushResponse) GetCodeSnapshotOk() (*CodeSnapshot, bool)`

GetCodeSnapshotOk returns a tuple with the CodeSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSnapshot

`func (o *PushResponse) SetCodeSnapshot(v CodeSnapshot)`

SetCodeSnapshot sets CodeSnapshot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


