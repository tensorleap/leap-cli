# CodeSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**TeamId** | **string** |  | 
**ProjectId** | **string** |  | 
**CreatedAt** | **string** |  | 
**CreatedBy** | **string** |  | 
**BlobName** | **string** |  | 
**ParseResult** | Pointer to [**CodeSnapshotResult**](CodeSnapshotResult.md) |  | [optional] 
**CodeEntryFile** | **string** |  | 
**GenericBaseImageType** | Pointer to **string** |  | [optional] 
**SecretManagerId** | Pointer to **string** |  | [optional] 

## Methods

### NewCodeSnapshot

`func NewCodeSnapshot(cid string, teamId string, projectId string, createdAt string, createdBy string, blobName string, codeEntryFile string, ) *CodeSnapshot`

NewCodeSnapshot instantiates a new CodeSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeSnapshotWithDefaults

`func NewCodeSnapshotWithDefaults() *CodeSnapshot`

NewCodeSnapshotWithDefaults instantiates a new CodeSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *CodeSnapshot) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *CodeSnapshot) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *CodeSnapshot) SetCid(v string)`

SetCid sets Cid field to given value.


### GetTeamId

`func (o *CodeSnapshot) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CodeSnapshot) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CodeSnapshot) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetProjectId

`func (o *CodeSnapshot) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CodeSnapshot) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CodeSnapshot) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *CodeSnapshot) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CodeSnapshot) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CodeSnapshot) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *CodeSnapshot) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CodeSnapshot) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CodeSnapshot) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetBlobName

`func (o *CodeSnapshot) GetBlobName() string`

GetBlobName returns the BlobName field if non-nil, zero value otherwise.

### GetBlobNameOk

`func (o *CodeSnapshot) GetBlobNameOk() (*string, bool)`

GetBlobNameOk returns a tuple with the BlobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobName

`func (o *CodeSnapshot) SetBlobName(v string)`

SetBlobName sets BlobName field to given value.


### GetParseResult

`func (o *CodeSnapshot) GetParseResult() CodeSnapshotResult`

GetParseResult returns the ParseResult field if non-nil, zero value otherwise.

### GetParseResultOk

`func (o *CodeSnapshot) GetParseResultOk() (*CodeSnapshotResult, bool)`

GetParseResultOk returns a tuple with the ParseResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseResult

`func (o *CodeSnapshot) SetParseResult(v CodeSnapshotResult)`

SetParseResult sets ParseResult field to given value.

### HasParseResult

`func (o *CodeSnapshot) HasParseResult() bool`

HasParseResult returns a boolean if a field has been set.

### GetCodeEntryFile

`func (o *CodeSnapshot) GetCodeEntryFile() string`

GetCodeEntryFile returns the CodeEntryFile field if non-nil, zero value otherwise.

### GetCodeEntryFileOk

`func (o *CodeSnapshot) GetCodeEntryFileOk() (*string, bool)`

GetCodeEntryFileOk returns a tuple with the CodeEntryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeEntryFile

`func (o *CodeSnapshot) SetCodeEntryFile(v string)`

SetCodeEntryFile sets CodeEntryFile field to given value.


### GetGenericBaseImageType

`func (o *CodeSnapshot) GetGenericBaseImageType() string`

GetGenericBaseImageType returns the GenericBaseImageType field if non-nil, zero value otherwise.

### GetGenericBaseImageTypeOk

`func (o *CodeSnapshot) GetGenericBaseImageTypeOk() (*string, bool)`

GetGenericBaseImageTypeOk returns a tuple with the GenericBaseImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericBaseImageType

`func (o *CodeSnapshot) SetGenericBaseImageType(v string)`

SetGenericBaseImageType sets GenericBaseImageType field to given value.

### HasGenericBaseImageType

`func (o *CodeSnapshot) HasGenericBaseImageType() bool`

HasGenericBaseImageType returns a boolean if a field has been set.

### GetSecretManagerId

`func (o *CodeSnapshot) GetSecretManagerId() string`

GetSecretManagerId returns the SecretManagerId field if non-nil, zero value otherwise.

### GetSecretManagerIdOk

`func (o *CodeSnapshot) GetSecretManagerIdOk() (*string, bool)`

GetSecretManagerIdOk returns a tuple with the SecretManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretManagerId

`func (o *CodeSnapshot) SetSecretManagerId(v string)`

SetSecretManagerId sets SecretManagerId field to given value.

### HasSecretManagerId

`func (o *CodeSnapshot) HasSecretManagerId() bool`

HasSecretManagerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


