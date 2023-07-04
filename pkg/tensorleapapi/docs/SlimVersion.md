# SlimVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**CreatedBy** | **string** |  | 
**ProjectId** | **string** |  | 
**BranchName** | **string** |  | 
**Tags** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Notes** | **string** |  | 
**CodeIntegrationVersionId** | Pointer to **string** |  | [optional] 
**DatasetSetup** | [**DatasetSetup**](DatasetSetup.md) |  | 
**IsFavourite** | **bool** |  | 
**Hash** | Pointer to **NullableString** |  | [optional] 
**Sessions** | [**[]Session**](Session.md) |  | 

## Methods

### NewSlimVersion

`func NewSlimVersion(cid string, createdBy string, projectId string, branchName string, tags string, createdAt time.Time, notes string, datasetSetup DatasetSetup, isFavourite bool, sessions []Session, ) *SlimVersion`

NewSlimVersion instantiates a new SlimVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlimVersionWithDefaults

`func NewSlimVersionWithDefaults() *SlimVersion`

NewSlimVersionWithDefaults instantiates a new SlimVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *SlimVersion) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *SlimVersion) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *SlimVersion) SetCid(v string)`

SetCid sets Cid field to given value.


### GetCreatedBy

`func (o *SlimVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SlimVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SlimVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetProjectId

`func (o *SlimVersion) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SlimVersion) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SlimVersion) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetBranchName

`func (o *SlimVersion) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *SlimVersion) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *SlimVersion) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetTags

`func (o *SlimVersion) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SlimVersion) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SlimVersion) SetTags(v string)`

SetTags sets Tags field to given value.


### GetCreatedAt

`func (o *SlimVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SlimVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SlimVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetNotes

`func (o *SlimVersion) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SlimVersion) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SlimVersion) SetNotes(v string)`

SetNotes sets Notes field to given value.


### GetCodeIntegrationVersionId

`func (o *SlimVersion) GetCodeIntegrationVersionId() string`

GetCodeIntegrationVersionId returns the CodeIntegrationVersionId field if non-nil, zero value otherwise.

### GetCodeIntegrationVersionIdOk

`func (o *SlimVersion) GetCodeIntegrationVersionIdOk() (*string, bool)`

GetCodeIntegrationVersionIdOk returns a tuple with the CodeIntegrationVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIntegrationVersionId

`func (o *SlimVersion) SetCodeIntegrationVersionId(v string)`

SetCodeIntegrationVersionId sets CodeIntegrationVersionId field to given value.

### HasCodeIntegrationVersionId

`func (o *SlimVersion) HasCodeIntegrationVersionId() bool`

HasCodeIntegrationVersionId returns a boolean if a field has been set.

### GetDatasetSetup

`func (o *SlimVersion) GetDatasetSetup() DatasetSetup`

GetDatasetSetup returns the DatasetSetup field if non-nil, zero value otherwise.

### GetDatasetSetupOk

`func (o *SlimVersion) GetDatasetSetupOk() (*DatasetSetup, bool)`

GetDatasetSetupOk returns a tuple with the DatasetSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetSetup

`func (o *SlimVersion) SetDatasetSetup(v DatasetSetup)`

SetDatasetSetup sets DatasetSetup field to given value.


### GetIsFavourite

`func (o *SlimVersion) GetIsFavourite() bool`

GetIsFavourite returns the IsFavourite field if non-nil, zero value otherwise.

### GetIsFavouriteOk

`func (o *SlimVersion) GetIsFavouriteOk() (*bool, bool)`

GetIsFavouriteOk returns a tuple with the IsFavourite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavourite

`func (o *SlimVersion) SetIsFavourite(v bool)`

SetIsFavourite sets IsFavourite field to given value.


### GetHash

`func (o *SlimVersion) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SlimVersion) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SlimVersion) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *SlimVersion) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *SlimVersion) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *SlimVersion) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetSessions

`func (o *SlimVersion) GetSessions() []Session`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *SlimVersion) GetSessionsOk() (*[]Session, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *SlimVersion) SetSessions(v []Session)`

SetSessions sets Sessions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


