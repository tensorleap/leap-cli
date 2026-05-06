# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**CreatedBy** | **string** |  | 
**TeamId** | **string** |  | 
**ProjectId** | **string** |  | 
**Branch** | **string** |  | 
**Tag** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Notes** | **string** |  | 
**IsFavourite** | **bool** |  | 
**CodeSnapshotId** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **NullableString** |  | [optional] 
**GraphValidationData** | Pointer to [**GraphValidatorData**](GraphValidatorData.md) |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**EpochTags** | Pointer to **map[string]float64** |  | [optional] 
**UploadedModel** | Pointer to [**UploadedModel**](UploadedModel.md) |  | [optional] 
**Resources** | [**VersionResources**](VersionResources.md) |  | 
**ParentVersionId** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**ExperimentId** | Pointer to **string** |  | [optional] 
**UpdateActions** | Pointer to [**[]UpdateAction**](UpdateAction.md) |  | [optional] 
**SerialNumber** | Pointer to **float64** |  | [optional] 
**Data** | Pointer to [**ModelGraph**](ModelGraph.md) |  | [optional] 
**Status** | [**VersionStatus**](VersionStatus.md) |  | 

## Methods

### NewVersion

`func NewVersion(cid string, createdBy string, teamId string, projectId string, branch string, tag string, createdAt time.Time, updatedAt time.Time, notes string, isFavourite bool, resources VersionResources, status VersionStatus, ) *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *Version) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *Version) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *Version) SetCid(v string)`

SetCid sets Cid field to given value.


### GetCreatedBy

`func (o *Version) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Version) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Version) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetTeamId

`func (o *Version) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Version) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Version) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetProjectId

`func (o *Version) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Version) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Version) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetBranch

`func (o *Version) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *Version) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *Version) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetTag

`func (o *Version) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Version) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Version) SetTag(v string)`

SetTag sets Tag field to given value.


### GetCreatedAt

`func (o *Version) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Version) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Version) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Version) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Version) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Version) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetNotes

`func (o *Version) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Version) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Version) SetNotes(v string)`

SetNotes sets Notes field to given value.


### GetIsFavourite

`func (o *Version) GetIsFavourite() bool`

GetIsFavourite returns the IsFavourite field if non-nil, zero value otherwise.

### GetIsFavouriteOk

`func (o *Version) GetIsFavouriteOk() (*bool, bool)`

GetIsFavouriteOk returns a tuple with the IsFavourite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavourite

`func (o *Version) SetIsFavourite(v bool)`

SetIsFavourite sets IsFavourite field to given value.


### GetCodeSnapshotId

`func (o *Version) GetCodeSnapshotId() string`

GetCodeSnapshotId returns the CodeSnapshotId field if non-nil, zero value otherwise.

### GetCodeSnapshotIdOk

`func (o *Version) GetCodeSnapshotIdOk() (*string, bool)`

GetCodeSnapshotIdOk returns a tuple with the CodeSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSnapshotId

`func (o *Version) SetCodeSnapshotId(v string)`

SetCodeSnapshotId sets CodeSnapshotId field to given value.

### HasCodeSnapshotId

`func (o *Version) HasCodeSnapshotId() bool`

HasCodeSnapshotId returns a boolean if a field has been set.

### GetHash

`func (o *Version) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Version) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Version) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Version) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *Version) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *Version) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetGraphValidationData

`func (o *Version) GetGraphValidationData() GraphValidatorData`

GetGraphValidationData returns the GraphValidationData field if non-nil, zero value otherwise.

### GetGraphValidationDataOk

`func (o *Version) GetGraphValidationDataOk() (*GraphValidatorData, bool)`

GetGraphValidationDataOk returns a tuple with the GraphValidationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphValidationData

`func (o *Version) SetGraphValidationData(v GraphValidatorData)`

SetGraphValidationData sets GraphValidationData field to given value.

### HasGraphValidationData

`func (o *Version) HasGraphValidationData() bool`

HasGraphValidationData returns a boolean if a field has been set.

### GetModelId

`func (o *Version) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *Version) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *Version) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *Version) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetEpochTags

`func (o *Version) GetEpochTags() map[string]float64`

GetEpochTags returns the EpochTags field if non-nil, zero value otherwise.

### GetEpochTagsOk

`func (o *Version) GetEpochTagsOk() (*map[string]float64, bool)`

GetEpochTagsOk returns a tuple with the EpochTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochTags

`func (o *Version) SetEpochTags(v map[string]float64)`

SetEpochTags sets EpochTags field to given value.

### HasEpochTags

`func (o *Version) HasEpochTags() bool`

HasEpochTags returns a boolean if a field has been set.

### GetUploadedModel

`func (o *Version) GetUploadedModel() UploadedModel`

GetUploadedModel returns the UploadedModel field if non-nil, zero value otherwise.

### GetUploadedModelOk

`func (o *Version) GetUploadedModelOk() (*UploadedModel, bool)`

GetUploadedModelOk returns a tuple with the UploadedModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedModel

`func (o *Version) SetUploadedModel(v UploadedModel)`

SetUploadedModel sets UploadedModel field to given value.

### HasUploadedModel

`func (o *Version) HasUploadedModel() bool`

HasUploadedModel returns a boolean if a field has been set.

### GetResources

`func (o *Version) GetResources() VersionResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Version) GetResourcesOk() (*VersionResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Version) SetResources(v VersionResources)`

SetResources sets Resources field to given value.


### GetParentVersionId

`func (o *Version) GetParentVersionId() string`

GetParentVersionId returns the ParentVersionId field if non-nil, zero value otherwise.

### GetParentVersionIdOk

`func (o *Version) GetParentVersionIdOk() (*string, bool)`

GetParentVersionIdOk returns a tuple with the ParentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVersionId

`func (o *Version) SetParentVersionId(v string)`

SetParentVersionId sets ParentVersionId field to given value.

### HasParentVersionId

`func (o *Version) HasParentVersionId() bool`

HasParentVersionId returns a boolean if a field has been set.

### GetIsActive

`func (o *Version) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Version) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Version) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Version) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetExperimentId

`func (o *Version) GetExperimentId() string`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *Version) GetExperimentIdOk() (*string, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *Version) SetExperimentId(v string)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *Version) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### GetUpdateActions

`func (o *Version) GetUpdateActions() []UpdateAction`

GetUpdateActions returns the UpdateActions field if non-nil, zero value otherwise.

### GetUpdateActionsOk

`func (o *Version) GetUpdateActionsOk() (*[]UpdateAction, bool)`

GetUpdateActionsOk returns a tuple with the UpdateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateActions

`func (o *Version) SetUpdateActions(v []UpdateAction)`

SetUpdateActions sets UpdateActions field to given value.

### HasUpdateActions

`func (o *Version) HasUpdateActions() bool`

HasUpdateActions returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Version) GetSerialNumber() float64`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Version) GetSerialNumberOk() (*float64, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Version) SetSerialNumber(v float64)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Version) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetData

`func (o *Version) GetData() ModelGraph`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Version) GetDataOk() (*ModelGraph, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Version) SetData(v ModelGraph)`

SetData sets Data field to given value.

### HasData

`func (o *Version) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *Version) GetStatus() VersionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Version) GetStatusOk() (*VersionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Version) SetStatus(v VersionStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


