# UpdateProjectMetaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Tags** | **[]string** |  | 
**HubPublishPolicy** | [**HubPublishPolicy**](HubPublishPolicy.md) |  | 
**BgImageUrl** | Pointer to **string** |  | [optional] 
**ProjectId** | **string** |  | 

## Methods

### NewUpdateProjectMetaRequest

`func NewUpdateProjectMetaRequest(name string, description string, tags []string, hubPublishPolicy HubPublishPolicy, projectId string, ) *UpdateProjectMetaRequest`

NewUpdateProjectMetaRequest instantiates a new UpdateProjectMetaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectMetaRequestWithDefaults

`func NewUpdateProjectMetaRequestWithDefaults() *UpdateProjectMetaRequest`

NewUpdateProjectMetaRequestWithDefaults instantiates a new UpdateProjectMetaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateProjectMetaRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProjectMetaRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProjectMetaRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateProjectMetaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProjectMetaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProjectMetaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *UpdateProjectMetaRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateProjectMetaRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateProjectMetaRequest) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetHubPublishPolicy

`func (o *UpdateProjectMetaRequest) GetHubPublishPolicy() HubPublishPolicy`

GetHubPublishPolicy returns the HubPublishPolicy field if non-nil, zero value otherwise.

### GetHubPublishPolicyOk

`func (o *UpdateProjectMetaRequest) GetHubPublishPolicyOk() (*HubPublishPolicy, bool)`

GetHubPublishPolicyOk returns a tuple with the HubPublishPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubPublishPolicy

`func (o *UpdateProjectMetaRequest) SetHubPublishPolicy(v HubPublishPolicy)`

SetHubPublishPolicy sets HubPublishPolicy field to given value.


### GetBgImageUrl

`func (o *UpdateProjectMetaRequest) GetBgImageUrl() string`

GetBgImageUrl returns the BgImageUrl field if non-nil, zero value otherwise.

### GetBgImageUrlOk

`func (o *UpdateProjectMetaRequest) GetBgImageUrlOk() (*string, bool)`

GetBgImageUrlOk returns a tuple with the BgImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgImageUrl

`func (o *UpdateProjectMetaRequest) SetBgImageUrl(v string)`

SetBgImageUrl sets BgImageUrl field to given value.

### HasBgImageUrl

`func (o *UpdateProjectMetaRequest) HasBgImageUrl() bool`

HasBgImageUrl returns a boolean if a field has been set.

### GetProjectId

`func (o *UpdateProjectMetaRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateProjectMetaRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateProjectMetaRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


