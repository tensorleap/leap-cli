# ProjectMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Tags** | **[]string** |  | 
**HubPublishPolicy** | [**HubPublishPolicy**](HubPublishPolicy.md) |  | 
**BgImageUrl** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewProjectMeta

`func NewProjectMeta(name string, description string, tags []string, hubPublishPolicy HubPublishPolicy, ) *ProjectMeta`

NewProjectMeta instantiates a new ProjectMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectMetaWithDefaults

`func NewProjectMetaWithDefaults() *ProjectMeta`

NewProjectMetaWithDefaults instantiates a new ProjectMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectMeta) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProjectMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectMeta) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *ProjectMeta) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectMeta) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectMeta) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetHubPublishPolicy

`func (o *ProjectMeta) GetHubPublishPolicy() HubPublishPolicy`

GetHubPublishPolicy returns the HubPublishPolicy field if non-nil, zero value otherwise.

### GetHubPublishPolicyOk

`func (o *ProjectMeta) GetHubPublishPolicyOk() (*HubPublishPolicy, bool)`

GetHubPublishPolicyOk returns a tuple with the HubPublishPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubPublishPolicy

`func (o *ProjectMeta) SetHubPublishPolicy(v HubPublishPolicy)`

SetHubPublishPolicy sets HubPublishPolicy field to given value.


### GetBgImageUrl

`func (o *ProjectMeta) GetBgImageUrl() string`

GetBgImageUrl returns the BgImageUrl field if non-nil, zero value otherwise.

### GetBgImageUrlOk

`func (o *ProjectMeta) GetBgImageUrlOk() (*string, bool)`

GetBgImageUrlOk returns a tuple with the BgImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgImageUrl

`func (o *ProjectMeta) SetBgImageUrl(v string)`

SetBgImageUrl sets BgImageUrl field to given value.

### HasBgImageUrl

`func (o *ProjectMeta) HasBgImageUrl() bool`

HasBgImageUrl returns a boolean if a field has been set.

### GetCategories

`func (o *ProjectMeta) GetCategories() map[string]interface{}`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ProjectMeta) GetCategoriesOk() (*map[string]interface{}, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ProjectMeta) SetCategories(v map[string]interface{})`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ProjectMeta) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


