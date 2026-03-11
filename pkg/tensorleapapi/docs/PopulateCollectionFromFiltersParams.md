# PopulateCollectionFromFiltersParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**Filters** | [**[]ESFilter**](ESFilter.md) |  | 
**InferenceArtifactIds** | **[]string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewPopulateCollectionFromFiltersParams

`func NewPopulateCollectionFromFiltersParams(filters []ESFilter, inferenceArtifactIds []string, projectId string, ) *PopulateCollectionFromFiltersParams`

NewPopulateCollectionFromFiltersParams instantiates a new PopulateCollectionFromFiltersParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPopulateCollectionFromFiltersParamsWithDefaults

`func NewPopulateCollectionFromFiltersParamsWithDefaults() *PopulateCollectionFromFiltersParams`

NewPopulateCollectionFromFiltersParamsWithDefaults instantiates a new PopulateCollectionFromFiltersParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PopulateCollectionFromFiltersParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PopulateCollectionFromFiltersParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PopulateCollectionFromFiltersParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PopulateCollectionFromFiltersParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PopulateCollectionFromFiltersParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PopulateCollectionFromFiltersParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PopulateCollectionFromFiltersParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PopulateCollectionFromFiltersParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCollectionId

`func (o *PopulateCollectionFromFiltersParams) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *PopulateCollectionFromFiltersParams) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *PopulateCollectionFromFiltersParams) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *PopulateCollectionFromFiltersParams) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetFilters

`func (o *PopulateCollectionFromFiltersParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PopulateCollectionFromFiltersParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PopulateCollectionFromFiltersParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.


### GetInferenceArtifactIds

`func (o *PopulateCollectionFromFiltersParams) GetInferenceArtifactIds() []string`

GetInferenceArtifactIds returns the InferenceArtifactIds field if non-nil, zero value otherwise.

### GetInferenceArtifactIdsOk

`func (o *PopulateCollectionFromFiltersParams) GetInferenceArtifactIdsOk() (*[]string, bool)`

GetInferenceArtifactIdsOk returns a tuple with the InferenceArtifactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactIds

`func (o *PopulateCollectionFromFiltersParams) SetInferenceArtifactIds(v []string)`

SetInferenceArtifactIds sets InferenceArtifactIds field to given value.


### GetProjectId

`func (o *PopulateCollectionFromFiltersParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PopulateCollectionFromFiltersParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PopulateCollectionFromFiltersParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


