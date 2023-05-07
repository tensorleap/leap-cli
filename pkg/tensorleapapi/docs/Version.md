# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ExtId** | **string** |  | 
**CreatedBy** | **string** |  | 
**Organization** | **string** |  | 
**Project** | **string** |  | 
**Branch** | **string** |  | 
**Tag** | **string** |  | 
**Data** | [**ModelGraph**](ModelGraph.md) |  | 
**CreatedAt** | **time.Time** |  | 
**Notes** | **string** |  | 
**Status** | **string** |  | 
**IsFavourite** | **bool** |  | 
**CodeIntegrationVersionId** | Pointer to **string** |  | [optional] 
**DatasetSetup** | [**DatasetSetup**](DatasetSetup.md) |  | 

## Methods

### NewVersion

`func NewVersion(id string, extId string, createdBy string, organization string, project string, branch string, tag string, data ModelGraph, createdAt time.Time, notes string, status string, isFavourite bool, datasetSetup DatasetSetup, ) *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Version) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Version) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Version) SetId(v string)`

SetId sets Id field to given value.


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


### GetOrganization

`func (o *Version) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Version) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Version) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetProject

`func (o *Version) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Version) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Version) SetProject(v string)`

SetProject sets Project field to given value.


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


### GetCodeIntegrationVersionId

`func (o *Version) GetCodeIntegrationVersionId() string`

GetCodeIntegrationVersionId returns the CodeIntegrationVersionId field if non-nil, zero value otherwise.

### GetCodeIntegrationVersionIdOk

`func (o *Version) GetCodeIntegrationVersionIdOk() (*string, bool)`

GetCodeIntegrationVersionIdOk returns a tuple with the CodeIntegrationVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIntegrationVersionId

`func (o *Version) SetCodeIntegrationVersionId(v string)`

SetCodeIntegrationVersionId sets CodeIntegrationVersionId field to given value.

### HasCodeIntegrationVersionId

`func (o *Version) HasCodeIntegrationVersionId() bool`

HasCodeIntegrationVersionId returns a boolean if a field has been set.

### GetDatasetSetup

`func (o *Version) GetDatasetSetup() DatasetSetup`

GetDatasetSetup returns the DatasetSetup field if non-nil, zero value otherwise.

### GetDatasetSetupOk

`func (o *Version) GetDatasetSetupOk() (*DatasetSetup, bool)`

GetDatasetSetupOk returns a tuple with the DatasetSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetSetup

`func (o *Version) SetDatasetSetup(v DatasetSetup)`

SetDatasetSetup sets DatasetSetup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


