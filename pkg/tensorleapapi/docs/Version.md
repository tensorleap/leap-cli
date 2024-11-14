# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**ExtId** | **string** |  | 
**CreatedBy** | **string** |  | 
**TeamId** | **string** |  | 
**ProjectId** | **string** |  | 
**Branch** | **string** |  | 
**Tag** | **string** |  | 
**Data** | Pointer to [**ModelGraph**](ModelGraph.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Notes** | **string** |  | 
**Status** | **string** |  | 
**IsFavourite** | **bool** |  | 
**CodeIntegration** | Pointer to [**CodeIntegrationBinder**](CodeIntegrationBinder.md) |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 

## Methods

### NewVersion

`func NewVersion(cid string, extId string, createdBy string, teamId string, projectId string, branch string, tag string, createdAt time.Time, notes string, status string, isFavourite bool, ) *Version`

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


### GetExtId

`func (o *Version) GetExtId() string`

GetExtId returns the ExtId field if non-nil, zero value otherwise.

### GetExtIdOk

`func (o *Version) GetExtIdOk() (*string, bool)`

GetExtIdOk returns a tuple with the ExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtId

`func (o *Version) SetExtId(v string)`

SetExtId sets ExtId field to given value.


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


### GetStatus

`func (o *Version) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Version) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Version) SetStatus(v string)`

SetStatus sets Status field to given value.


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


### GetCodeIntegration

`func (o *Version) GetCodeIntegration() CodeIntegrationBinder`

GetCodeIntegration returns the CodeIntegration field if non-nil, zero value otherwise.

### GetCodeIntegrationOk

`func (o *Version) GetCodeIntegrationOk() (*CodeIntegrationBinder, bool)`

GetCodeIntegrationOk returns a tuple with the CodeIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIntegration

`func (o *Version) SetCodeIntegration(v CodeIntegrationBinder)`

SetCodeIntegration sets CodeIntegration field to given value.

### HasCodeIntegration

`func (o *Version) HasCodeIntegration() bool`

HasCodeIntegration returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


