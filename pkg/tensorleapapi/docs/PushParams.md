# PushParams

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
**Model** | [**ImportModelInfo**](ImportModelInfo.md) |  | 

## Methods

### NewPushParams

`func NewPushParams(projectId string, codeUrl string, codeEntryFile string, versionName string, model ImportModelInfo, ) *PushParams`

NewPushParams instantiates a new PushParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushParamsWithDefaults

`func NewPushParamsWithDefaults() *PushParams`

NewPushParamsWithDefaults instantiates a new PushParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *PushParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PushParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PushParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSecretManagerId

`func (o *PushParams) GetSecretManagerId() string`

GetSecretManagerId returns the SecretManagerId field if non-nil, zero value otherwise.

### GetSecretManagerIdOk

`func (o *PushParams) GetSecretManagerIdOk() (*string, bool)`

GetSecretManagerIdOk returns a tuple with the SecretManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretManagerId

`func (o *PushParams) SetSecretManagerId(v string)`

SetSecretManagerId sets SecretManagerId field to given value.

### HasSecretManagerId

`func (o *PushParams) HasSecretManagerId() bool`

HasSecretManagerId returns a boolean if a field has been set.

### GetCodeUrl

`func (o *PushParams) GetCodeUrl() string`

GetCodeUrl returns the CodeUrl field if non-nil, zero value otherwise.

### GetCodeUrlOk

`func (o *PushParams) GetCodeUrlOk() (*string, bool)`

GetCodeUrlOk returns a tuple with the CodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeUrl

`func (o *PushParams) SetCodeUrl(v string)`

SetCodeUrl sets CodeUrl field to given value.


### GetCodeEntryFile

`func (o *PushParams) GetCodeEntryFile() string`

GetCodeEntryFile returns the CodeEntryFile field if non-nil, zero value otherwise.

### GetCodeEntryFileOk

`func (o *PushParams) GetCodeEntryFileOk() (*string, bool)`

GetCodeEntryFileOk returns a tuple with the CodeEntryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeEntryFile

`func (o *PushParams) SetCodeEntryFile(v string)`

SetCodeEntryFile sets CodeEntryFile field to given value.


### GetGenericBaseImageType

`func (o *PushParams) GetGenericBaseImageType() string`

GetGenericBaseImageType returns the GenericBaseImageType field if non-nil, zero value otherwise.

### GetGenericBaseImageTypeOk

`func (o *PushParams) GetGenericBaseImageTypeOk() (*string, bool)`

GetGenericBaseImageTypeOk returns a tuple with the GenericBaseImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericBaseImageType

`func (o *PushParams) SetGenericBaseImageType(v string)`

SetGenericBaseImageType sets GenericBaseImageType field to given value.

### HasGenericBaseImageType

`func (o *PushParams) HasGenericBaseImageType() bool`

HasGenericBaseImageType returns a boolean if a field has been set.

### GetVersionName

`func (o *PushParams) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *PushParams) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *PushParams) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.


### GetBranchName

`func (o *PushParams) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *PushParams) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *PushParams) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *PushParams) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetOverwriteVersionId

`func (o *PushParams) GetOverwriteVersionId() string`

GetOverwriteVersionId returns the OverwriteVersionId field if non-nil, zero value otherwise.

### GetOverwriteVersionIdOk

`func (o *PushParams) GetOverwriteVersionIdOk() (*string, bool)`

GetOverwriteVersionIdOk returns a tuple with the OverwriteVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteVersionId

`func (o *PushParams) SetOverwriteVersionId(v string)`

SetOverwriteVersionId sets OverwriteVersionId field to given value.

### HasOverwriteVersionId

`func (o *PushParams) HasOverwriteVersionId() bool`

HasOverwriteVersionId returns a boolean if a field has been set.

### GetModel

`func (o *PushParams) GetModel() ImportModelInfo`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PushParams) GetModelOk() (*ImportModelInfo, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PushParams) SetModel(v ImportModelInfo)`

SetModel sets Model field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


