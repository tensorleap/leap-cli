# Dataset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**Name** | **NullableString** |  | 
**TeamId** | **string** |  | 
**SourceDatasetId** | Pointer to **string** |  | [optional] 
**Access** | [**DatasetAccess**](DatasetAccess.md) |  | 
**CreatedAt** | **string** |  | 
**DefaultBranch** | **string** |  | 
**Pippin** | Pointer to **bool** |  | [optional] 
**LatestVersions** | [**[]DatasetLatestVersionsInner**](DatasetLatestVersionsInner.md) |  | 

## Methods

### NewDataset

`func NewDataset(cid string, name NullableString, teamId string, access DatasetAccess, createdAt string, defaultBranch string, latestVersions []DatasetLatestVersionsInner, ) *Dataset`

NewDataset instantiates a new Dataset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetWithDefaults

`func NewDatasetWithDefaults() *Dataset`

NewDatasetWithDefaults instantiates a new Dataset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *Dataset) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *Dataset) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *Dataset) SetCid(v string)`

SetCid sets Cid field to given value.


### GetName

`func (o *Dataset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dataset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dataset) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Dataset) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Dataset) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTeamId

`func (o *Dataset) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Dataset) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Dataset) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetSourceDatasetId

`func (o *Dataset) GetSourceDatasetId() string`

GetSourceDatasetId returns the SourceDatasetId field if non-nil, zero value otherwise.

### GetSourceDatasetIdOk

`func (o *Dataset) GetSourceDatasetIdOk() (*string, bool)`

GetSourceDatasetIdOk returns a tuple with the SourceDatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDatasetId

`func (o *Dataset) SetSourceDatasetId(v string)`

SetSourceDatasetId sets SourceDatasetId field to given value.

### HasSourceDatasetId

`func (o *Dataset) HasSourceDatasetId() bool`

HasSourceDatasetId returns a boolean if a field has been set.

### GetAccess

`func (o *Dataset) GetAccess() DatasetAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *Dataset) GetAccessOk() (*DatasetAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *Dataset) SetAccess(v DatasetAccess)`

SetAccess sets Access field to given value.


### GetCreatedAt

`func (o *Dataset) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Dataset) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Dataset) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultBranch

`func (o *Dataset) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *Dataset) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *Dataset) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetPippin

`func (o *Dataset) GetPippin() bool`

GetPippin returns the Pippin field if non-nil, zero value otherwise.

### GetPippinOk

`func (o *Dataset) GetPippinOk() (*bool, bool)`

GetPippinOk returns a tuple with the Pippin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPippin

`func (o *Dataset) SetPippin(v bool)`

SetPippin sets Pippin field to given value.

### HasPippin

`func (o *Dataset) HasPippin() bool`

HasPippin returns a boolean if a field has been set.

### GetLatestVersions

`func (o *Dataset) GetLatestVersions() []DatasetLatestVersionsInner`

GetLatestVersions returns the LatestVersions field if non-nil, zero value otherwise.

### GetLatestVersionsOk

`func (o *Dataset) GetLatestVersionsOk() (*[]DatasetLatestVersionsInner, bool)`

GetLatestVersionsOk returns a tuple with the LatestVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersions

`func (o *Dataset) SetLatestVersions(v []DatasetLatestVersionsInner)`

SetLatestVersions sets LatestVersions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


