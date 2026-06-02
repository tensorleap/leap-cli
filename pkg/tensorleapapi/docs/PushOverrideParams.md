# PushOverrideParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SecretManagerId** | Pointer to **string** |  | [optional] 
**CodeUrl** | **string** |  | 
**CodeEntryFile** | **string** |  | 
**GenericBaseImageType** | Pointer to **string** |  | [optional] 
**OverwriteVersionId** | **string** |  | 
**BranchName** | Pointer to **string** |  | [optional] 

## Methods

### NewPushOverrideParams

`func NewPushOverrideParams(projectId string, codeUrl string, codeEntryFile string, overwriteVersionId string, ) *PushOverrideParams`

NewPushOverrideParams instantiates a new PushOverrideParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushOverrideParamsWithDefaults

`func NewPushOverrideParamsWithDefaults() *PushOverrideParams`

NewPushOverrideParamsWithDefaults instantiates a new PushOverrideParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *PushOverrideParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PushOverrideParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PushOverrideParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSecretManagerId

`func (o *PushOverrideParams) GetSecretManagerId() string`

GetSecretManagerId returns the SecretManagerId field if non-nil, zero value otherwise.

### GetSecretManagerIdOk

`func (o *PushOverrideParams) GetSecretManagerIdOk() (*string, bool)`

GetSecretManagerIdOk returns a tuple with the SecretManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretManagerId

`func (o *PushOverrideParams) SetSecretManagerId(v string)`

SetSecretManagerId sets SecretManagerId field to given value.

### HasSecretManagerId

`func (o *PushOverrideParams) HasSecretManagerId() bool`

HasSecretManagerId returns a boolean if a field has been set.

### GetCodeUrl

`func (o *PushOverrideParams) GetCodeUrl() string`

GetCodeUrl returns the CodeUrl field if non-nil, zero value otherwise.

### GetCodeUrlOk

`func (o *PushOverrideParams) GetCodeUrlOk() (*string, bool)`

GetCodeUrlOk returns a tuple with the CodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeUrl

`func (o *PushOverrideParams) SetCodeUrl(v string)`

SetCodeUrl sets CodeUrl field to given value.


### GetCodeEntryFile

`func (o *PushOverrideParams) GetCodeEntryFile() string`

GetCodeEntryFile returns the CodeEntryFile field if non-nil, zero value otherwise.

### GetCodeEntryFileOk

`func (o *PushOverrideParams) GetCodeEntryFileOk() (*string, bool)`

GetCodeEntryFileOk returns a tuple with the CodeEntryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeEntryFile

`func (o *PushOverrideParams) SetCodeEntryFile(v string)`

SetCodeEntryFile sets CodeEntryFile field to given value.


### GetGenericBaseImageType

`func (o *PushOverrideParams) GetGenericBaseImageType() string`

GetGenericBaseImageType returns the GenericBaseImageType field if non-nil, zero value otherwise.

### GetGenericBaseImageTypeOk

`func (o *PushOverrideParams) GetGenericBaseImageTypeOk() (*string, bool)`

GetGenericBaseImageTypeOk returns a tuple with the GenericBaseImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericBaseImageType

`func (o *PushOverrideParams) SetGenericBaseImageType(v string)`

SetGenericBaseImageType sets GenericBaseImageType field to given value.

### HasGenericBaseImageType

`func (o *PushOverrideParams) HasGenericBaseImageType() bool`

HasGenericBaseImageType returns a boolean if a field has been set.

### GetOverwriteVersionId

`func (o *PushOverrideParams) GetOverwriteVersionId() string`

GetOverwriteVersionId returns the OverwriteVersionId field if non-nil, zero value otherwise.

### GetOverwriteVersionIdOk

`func (o *PushOverrideParams) GetOverwriteVersionIdOk() (*string, bool)`

GetOverwriteVersionIdOk returns a tuple with the OverwriteVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteVersionId

`func (o *PushOverrideParams) SetOverwriteVersionId(v string)`

SetOverwriteVersionId sets OverwriteVersionId field to given value.


### GetBranchName

`func (o *PushOverrideParams) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *PushOverrideParams) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *PushOverrideParams) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *PushOverrideParams) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


