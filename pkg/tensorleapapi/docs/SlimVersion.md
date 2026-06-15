# SlimVersion

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
**Color** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 
**HasExternalEpoch** | **bool** |  | 
**IsEvaluate** | **bool** |  | 
**EvaluateParams** | Pointer to [**VersionEvaluateParams**](VersionEvaluateParams.md) |  | [optional] 
**CanContinueEvaluate** | Pointer to **bool** |  | [optional] 
**Jobs** | [**[]Job**](Job.md) |  | 

## Methods

### NewSlimVersion

`func NewSlimVersion(cid string, createdBy string, teamId string, projectId string, branch string, tag string, createdAt time.Time, updatedAt time.Time, notes string, isFavourite bool, resources VersionResources, hasExternalEpoch bool, isEvaluate bool, jobs []Job, ) *SlimVersion`

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


### GetTeamId

`func (o *SlimVersion) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *SlimVersion) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *SlimVersion) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


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


### GetBranch

`func (o *SlimVersion) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *SlimVersion) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *SlimVersion) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetTag

`func (o *SlimVersion) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *SlimVersion) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *SlimVersion) SetTag(v string)`

SetTag sets Tag field to given value.


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


### GetUpdatedAt

`func (o *SlimVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SlimVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SlimVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


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


### GetCodeSnapshotId

`func (o *SlimVersion) GetCodeSnapshotId() string`

GetCodeSnapshotId returns the CodeSnapshotId field if non-nil, zero value otherwise.

### GetCodeSnapshotIdOk

`func (o *SlimVersion) GetCodeSnapshotIdOk() (*string, bool)`

GetCodeSnapshotIdOk returns a tuple with the CodeSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSnapshotId

`func (o *SlimVersion) SetCodeSnapshotId(v string)`

SetCodeSnapshotId sets CodeSnapshotId field to given value.

### HasCodeSnapshotId

`func (o *SlimVersion) HasCodeSnapshotId() bool`

HasCodeSnapshotId returns a boolean if a field has been set.

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
### GetGraphValidationData

`func (o *SlimVersion) GetGraphValidationData() GraphValidatorData`

GetGraphValidationData returns the GraphValidationData field if non-nil, zero value otherwise.

### GetGraphValidationDataOk

`func (o *SlimVersion) GetGraphValidationDataOk() (*GraphValidatorData, bool)`

GetGraphValidationDataOk returns a tuple with the GraphValidationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphValidationData

`func (o *SlimVersion) SetGraphValidationData(v GraphValidatorData)`

SetGraphValidationData sets GraphValidationData field to given value.

### HasGraphValidationData

`func (o *SlimVersion) HasGraphValidationData() bool`

HasGraphValidationData returns a boolean if a field has been set.

### GetModelId

`func (o *SlimVersion) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *SlimVersion) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *SlimVersion) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *SlimVersion) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetEpochTags

`func (o *SlimVersion) GetEpochTags() map[string]float64`

GetEpochTags returns the EpochTags field if non-nil, zero value otherwise.

### GetEpochTagsOk

`func (o *SlimVersion) GetEpochTagsOk() (*map[string]float64, bool)`

GetEpochTagsOk returns a tuple with the EpochTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochTags

`func (o *SlimVersion) SetEpochTags(v map[string]float64)`

SetEpochTags sets EpochTags field to given value.

### HasEpochTags

`func (o *SlimVersion) HasEpochTags() bool`

HasEpochTags returns a boolean if a field has been set.

### GetUploadedModel

`func (o *SlimVersion) GetUploadedModel() UploadedModel`

GetUploadedModel returns the UploadedModel field if non-nil, zero value otherwise.

### GetUploadedModelOk

`func (o *SlimVersion) GetUploadedModelOk() (*UploadedModel, bool)`

GetUploadedModelOk returns a tuple with the UploadedModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedModel

`func (o *SlimVersion) SetUploadedModel(v UploadedModel)`

SetUploadedModel sets UploadedModel field to given value.

### HasUploadedModel

`func (o *SlimVersion) HasUploadedModel() bool`

HasUploadedModel returns a boolean if a field has been set.

### GetResources

`func (o *SlimVersion) GetResources() VersionResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *SlimVersion) GetResourcesOk() (*VersionResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *SlimVersion) SetResources(v VersionResources)`

SetResources sets Resources field to given value.


### GetParentVersionId

`func (o *SlimVersion) GetParentVersionId() string`

GetParentVersionId returns the ParentVersionId field if non-nil, zero value otherwise.

### GetParentVersionIdOk

`func (o *SlimVersion) GetParentVersionIdOk() (*string, bool)`

GetParentVersionIdOk returns a tuple with the ParentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVersionId

`func (o *SlimVersion) SetParentVersionId(v string)`

SetParentVersionId sets ParentVersionId field to given value.

### HasParentVersionId

`func (o *SlimVersion) HasParentVersionId() bool`

HasParentVersionId returns a boolean if a field has been set.

### GetIsActive

`func (o *SlimVersion) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SlimVersion) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SlimVersion) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *SlimVersion) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetExperimentId

`func (o *SlimVersion) GetExperimentId() string`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *SlimVersion) GetExperimentIdOk() (*string, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *SlimVersion) SetExperimentId(v string)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *SlimVersion) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### GetUpdateActions

`func (o *SlimVersion) GetUpdateActions() []UpdateAction`

GetUpdateActions returns the UpdateActions field if non-nil, zero value otherwise.

### GetUpdateActionsOk

`func (o *SlimVersion) GetUpdateActionsOk() (*[]UpdateAction, bool)`

GetUpdateActionsOk returns a tuple with the UpdateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateActions

`func (o *SlimVersion) SetUpdateActions(v []UpdateAction)`

SetUpdateActions sets UpdateActions field to given value.

### HasUpdateActions

`func (o *SlimVersion) HasUpdateActions() bool`

HasUpdateActions returns a boolean if a field has been set.

### GetSerialNumber

`func (o *SlimVersion) GetSerialNumber() float64`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SlimVersion) GetSerialNumberOk() (*float64, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SlimVersion) SetSerialNumber(v float64)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *SlimVersion) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetColor

`func (o *SlimVersion) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *SlimVersion) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *SlimVersion) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *SlimVersion) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetProperties

`func (o *SlimVersion) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SlimVersion) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SlimVersion) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SlimVersion) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetHasExternalEpoch

`func (o *SlimVersion) GetHasExternalEpoch() bool`

GetHasExternalEpoch returns the HasExternalEpoch field if non-nil, zero value otherwise.

### GetHasExternalEpochOk

`func (o *SlimVersion) GetHasExternalEpochOk() (*bool, bool)`

GetHasExternalEpochOk returns a tuple with the HasExternalEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExternalEpoch

`func (o *SlimVersion) SetHasExternalEpoch(v bool)`

SetHasExternalEpoch sets HasExternalEpoch field to given value.


### GetIsEvaluate

`func (o *SlimVersion) GetIsEvaluate() bool`

GetIsEvaluate returns the IsEvaluate field if non-nil, zero value otherwise.

### GetIsEvaluateOk

`func (o *SlimVersion) GetIsEvaluateOk() (*bool, bool)`

GetIsEvaluateOk returns a tuple with the IsEvaluate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEvaluate

`func (o *SlimVersion) SetIsEvaluate(v bool)`

SetIsEvaluate sets IsEvaluate field to given value.


### GetEvaluateParams

`func (o *SlimVersion) GetEvaluateParams() VersionEvaluateParams`

GetEvaluateParams returns the EvaluateParams field if non-nil, zero value otherwise.

### GetEvaluateParamsOk

`func (o *SlimVersion) GetEvaluateParamsOk() (*VersionEvaluateParams, bool)`

GetEvaluateParamsOk returns a tuple with the EvaluateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluateParams

`func (o *SlimVersion) SetEvaluateParams(v VersionEvaluateParams)`

SetEvaluateParams sets EvaluateParams field to given value.

### HasEvaluateParams

`func (o *SlimVersion) HasEvaluateParams() bool`

HasEvaluateParams returns a boolean if a field has been set.

### GetCanContinueEvaluate

`func (o *SlimVersion) GetCanContinueEvaluate() bool`

GetCanContinueEvaluate returns the CanContinueEvaluate field if non-nil, zero value otherwise.

### GetCanContinueEvaluateOk

`func (o *SlimVersion) GetCanContinueEvaluateOk() (*bool, bool)`

GetCanContinueEvaluateOk returns a tuple with the CanContinueEvaluate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanContinueEvaluate

`func (o *SlimVersion) SetCanContinueEvaluate(v bool)`

SetCanContinueEvaluate sets CanContinueEvaluate field to given value.

### HasCanContinueEvaluate

`func (o *SlimVersion) HasCanContinueEvaluate() bool`

HasCanContinueEvaluate returns a boolean if a field has been set.

### GetJobs

`func (o *SlimVersion) GetJobs() []Job`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *SlimVersion) GetJobsOk() (*[]Job, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *SlimVersion) SetJobs(v []Job)`

SetJobs sets Jobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


