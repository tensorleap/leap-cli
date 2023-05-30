# ImportProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Tags** | **[]string** |  | 
**HubPublishPolicy** | [**HubPublishPolicy**](HubPublishPolicy.md) |  | 
**BgImageUrl** | Pointer to **string** |  | [optional] 
**ImportUrl** | **string** |  | 

## Methods

### NewImportProjectRequest

`func NewImportProjectRequest(name string, description string, tags []string, hubPublishPolicy HubPublishPolicy, importUrl string, ) *ImportProjectRequest`

NewImportProjectRequest instantiates a new ImportProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportProjectRequestWithDefaults

`func NewImportProjectRequestWithDefaults() *ImportProjectRequest`

NewImportProjectRequestWithDefaults instantiates a new ImportProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImportProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ImportProjectRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImportProjectRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImportProjectRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *ImportProjectRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImportProjectRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImportProjectRequest) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetHubPublishPolicy

`func (o *ImportProjectRequest) GetHubPublishPolicy() HubPublishPolicy`

GetHubPublishPolicy returns the HubPublishPolicy field if non-nil, zero value otherwise.

### GetHubPublishPolicyOk

`func (o *ImportProjectRequest) GetHubPublishPolicyOk() (*HubPublishPolicy, bool)`

GetHubPublishPolicyOk returns a tuple with the HubPublishPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubPublishPolicy

`func (o *ImportProjectRequest) SetHubPublishPolicy(v HubPublishPolicy)`

SetHubPublishPolicy sets HubPublishPolicy field to given value.


### GetBgImageUrl

`func (o *ImportProjectRequest) GetBgImageUrl() string`

GetBgImageUrl returns the BgImageUrl field if non-nil, zero value otherwise.

### GetBgImageUrlOk

`func (o *ImportProjectRequest) GetBgImageUrlOk() (*string, bool)`

GetBgImageUrlOk returns a tuple with the BgImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgImageUrl

`func (o *ImportProjectRequest) SetBgImageUrl(v string)`

SetBgImageUrl sets BgImageUrl field to given value.

### HasBgImageUrl

`func (o *ImportProjectRequest) HasBgImageUrl() bool`

HasBgImageUrl returns a boolean if a field has been set.

### GetImportUrl

`func (o *ImportProjectRequest) GetImportUrl() string`

GetImportUrl returns the ImportUrl field if non-nil, zero value otherwise.

### GetImportUrlOk

`func (o *ImportProjectRequest) GetImportUrlOk() (*string, bool)`

GetImportUrlOk returns a tuple with the ImportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportUrl

`func (o *ImportProjectRequest) SetImportUrl(v string)`

SetImportUrl sets ImportUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


