# PushCodeSnapshotParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SecretManagerId** | Pointer to **string** |  | [optional] 
**CodeUrl** | **string** |  | 
**CodeEntryFile** | **string** |  | 
**GenericBaseImageType** | Pointer to **string** |  | [optional] 
**VersionName** | **string** |  | 
**BranchName** | Pointer to **string** |  | [optional] 
**OverwriteVersionId** | Pointer to **string** |  | [optional] 

## Methods

### NewPushCodeSnapshotParams

`func NewPushCodeSnapshotParams(projectId string, codeUrl string, codeEntryFile string, versionName string, ) *PushCodeSnapshotParams`

NewPushCodeSnapshotParams instantiates a new PushCodeSnapshotParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushCodeSnapshotParamsWithDefaults

`func NewPushCodeSnapshotParamsWithDefaults() *PushCodeSnapshotParams`

NewPushCodeSnapshotParamsWithDefaults instantiates a new PushCodeSnapshotParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *PushCodeSnapshotParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PushCodeSnapshotParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PushCodeSnapshotParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSecretManagerId

`func (o *PushCodeSnapshotParams) GetSecretManagerId() string`

GetSecretManagerId returns the SecretManagerId field if non-nil, zero value otherwise.

### GetSecretManagerIdOk

`func (o *PushCodeSnapshotParams) GetSecretManagerIdOk() (*string, bool)`

GetSecretManagerIdOk returns a tuple with the SecretManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretManagerId

`func (o *PushCodeSnapshotParams) SetSecretManagerId(v string)`

SetSecretManagerId sets SecretManagerId field to given value.

### HasSecretManagerId

`func (o *PushCodeSnapshotParams) HasSecretManagerId() bool`

HasSecretManagerId returns a boolean if a field has been set.

### GetCodeUrl

`func (o *PushCodeSnapshotParams) GetCodeUrl() string`

GetCodeUrl returns the CodeUrl field if non-nil, zero value otherwise.

### GetCodeUrlOk

`func (o *PushCodeSnapshotParams) GetCodeUrlOk() (*string, bool)`

GetCodeUrlOk returns a tuple with the CodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeUrl

`func (o *PushCodeSnapshotParams) SetCodeUrl(v string)`

SetCodeUrl sets CodeUrl field to given value.


### GetCodeEntryFile

`func (o *PushCodeSnapshotParams) GetCodeEntryFile() string`

GetCodeEntryFile returns the CodeEntryFile field if non-nil, zero value otherwise.

### GetCodeEntryFileOk

`func (o *PushCodeSnapshotParams) GetCodeEntryFileOk() (*string, bool)`

GetCodeEntryFileOk returns a tuple with the CodeEntryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeEntryFile

`func (o *PushCodeSnapshotParams) SetCodeEntryFile(v string)`

SetCodeEntryFile sets CodeEntryFile field to given value.


### GetGenericBaseImageType

`func (o *PushCodeSnapshotParams) GetGenericBaseImageType() string`

GetGenericBaseImageType returns the GenericBaseImageType field if non-nil, zero value otherwise.

### GetGenericBaseImageTypeOk

`func (o *PushCodeSnapshotParams) GetGenericBaseImageTypeOk() (*string, bool)`

GetGenericBaseImageTypeOk returns a tuple with the GenericBaseImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericBaseImageType

`func (o *PushCodeSnapshotParams) SetGenericBaseImageType(v string)`

SetGenericBaseImageType sets GenericBaseImageType field to given value.

### HasGenericBaseImageType

`func (o *PushCodeSnapshotParams) HasGenericBaseImageType() bool`

HasGenericBaseImageType returns a boolean if a field has been set.

### GetVersionName

`func (o *PushCodeSnapshotParams) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *PushCodeSnapshotParams) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *PushCodeSnapshotParams) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.


### GetBranchName

`func (o *PushCodeSnapshotParams) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *PushCodeSnapshotParams) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *PushCodeSnapshotParams) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *PushCodeSnapshotParams) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetOverwriteVersionId

`func (o *PushCodeSnapshotParams) GetOverwriteVersionId() string`

GetOverwriteVersionId returns the OverwriteVersionId field if non-nil, zero value otherwise.

### GetOverwriteVersionIdOk

`func (o *PushCodeSnapshotParams) GetOverwriteVersionIdOk() (*string, bool)`

GetOverwriteVersionIdOk returns a tuple with the OverwriteVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteVersionId

`func (o *PushCodeSnapshotParams) SetOverwriteVersionId(v string)`

SetOverwriteVersionId sets OverwriteVersionId field to given value.

### HasOverwriteVersionId

`func (o *PushCodeSnapshotParams) HasOverwriteVersionId() bool`

HasOverwriteVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


