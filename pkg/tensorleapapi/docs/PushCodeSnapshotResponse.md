# PushCodeSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeSnapshot** | [**CodeSnapshot**](CodeSnapshot.md) |  | 
**VersionId** | **string** |  | 

## Methods

### NewPushCodeSnapshotResponse

`func NewPushCodeSnapshotResponse(codeSnapshot CodeSnapshot, versionId string, ) *PushCodeSnapshotResponse`

NewPushCodeSnapshotResponse instantiates a new PushCodeSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushCodeSnapshotResponseWithDefaults

`func NewPushCodeSnapshotResponseWithDefaults() *PushCodeSnapshotResponse`

NewPushCodeSnapshotResponseWithDefaults instantiates a new PushCodeSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeSnapshot

`func (o *PushCodeSnapshotResponse) GetCodeSnapshot() CodeSnapshot`

GetCodeSnapshot returns the CodeSnapshot field if non-nil, zero value otherwise.

### GetCodeSnapshotOk

`func (o *PushCodeSnapshotResponse) GetCodeSnapshotOk() (*CodeSnapshot, bool)`

GetCodeSnapshotOk returns a tuple with the CodeSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSnapshot

`func (o *PushCodeSnapshotResponse) SetCodeSnapshot(v CodeSnapshot)`

SetCodeSnapshot sets CodeSnapshot field to given value.


### GetVersionId

`func (o *PushCodeSnapshotResponse) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *PushCodeSnapshotResponse) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *PushCodeSnapshotResponse) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


