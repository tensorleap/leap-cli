# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Organization** | **string** |  | 
**CreatedBy** | **string** |  | 
**CreatedAt** | **NullableTime** |  | 
**Name** | **string** |  | 
**LastAccessed** | **time.Time** |  | 
**Status** | [**ProjectStatus**](ProjectStatus.md) |  | 
**Access** | [**ProjectAccess**](ProjectAccess.md) |  | 
**Tags** | **[]string** |  | 
**BgImagePath** | Pointer to **string** |  | [optional] 
**HubPublishPolicy** | [**HubPublishPolicy**](HubPublishPolicy.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewProject

`func NewProject(id string, organization string, createdBy string, createdAt NullableTime, name string, lastAccessed time.Time, status ProjectStatus, access ProjectAccess, tags []string, hubPublishPolicy HubPublishPolicy, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Project) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v string)`

SetId sets Id field to given value.


### GetOrganization

`func (o *Project) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Project) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Project) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetCreatedBy

`func (o *Project) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Project) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Project) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *Project) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Project) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Project) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Project) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Project) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.


### GetLastAccessed

`func (o *Project) GetLastAccessed() time.Time`

GetLastAccessed returns the LastAccessed field if non-nil, zero value otherwise.

### GetLastAccessedOk

`func (o *Project) GetLastAccessedOk() (*time.Time, bool)`

GetLastAccessedOk returns a tuple with the LastAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessed

`func (o *Project) SetLastAccessed(v time.Time)`

SetLastAccessed sets LastAccessed field to given value.


### GetStatus

`func (o *Project) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Project) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Project) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.


### GetAccess

`func (o *Project) GetAccess() ProjectAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *Project) GetAccessOk() (*ProjectAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *Project) SetAccess(v ProjectAccess)`

SetAccess sets Access field to given value.


### GetTags

`func (o *Project) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Project) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Project) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetBgImagePath

`func (o *Project) GetBgImagePath() string`

GetBgImagePath returns the BgImagePath field if non-nil, zero value otherwise.

### GetBgImagePathOk

`func (o *Project) GetBgImagePathOk() (*string, bool)`

GetBgImagePathOk returns a tuple with the BgImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgImagePath

`func (o *Project) SetBgImagePath(v string)`

SetBgImagePath sets BgImagePath field to given value.

### HasBgImagePath

`func (o *Project) HasBgImagePath() bool`

HasBgImagePath returns a boolean if a field has been set.

### GetHubPublishPolicy

`func (o *Project) GetHubPublishPolicy() HubPublishPolicy`

GetHubPublishPolicy returns the HubPublishPolicy field if non-nil, zero value otherwise.

### GetHubPublishPolicyOk

`func (o *Project) GetHubPublishPolicyOk() (*HubPublishPolicy, bool)`

GetHubPublishPolicyOk returns a tuple with the HubPublishPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubPublishPolicy

`func (o *Project) SetHubPublishPolicy(v HubPublishPolicy)`

SetHubPublishPolicy sets HubPublishPolicy field to given value.


### GetDescription

`func (o *Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Project) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Project) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


