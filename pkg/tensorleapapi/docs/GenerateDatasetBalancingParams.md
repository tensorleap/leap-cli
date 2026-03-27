# GenerateDatasetBalancingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**VersionId** | **string** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**MetadataTags** | **[]string** |  | 
**PrioritizedMetadataTags** | Pointer to **[]string** |  | [optional] 
**PercentageOfSamplesToPrune** | Pointer to **float64** |  | [optional] 

## Methods

### NewGenerateDatasetBalancingParams

`func NewGenerateDatasetBalancingParams(projectId string, versionId string, metadataTags []string, ) *GenerateDatasetBalancingParams`

NewGenerateDatasetBalancingParams instantiates a new GenerateDatasetBalancingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateDatasetBalancingParamsWithDefaults

`func NewGenerateDatasetBalancingParamsWithDefaults() *GenerateDatasetBalancingParams`

NewGenerateDatasetBalancingParamsWithDefaults instantiates a new GenerateDatasetBalancingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GenerateDatasetBalancingParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GenerateDatasetBalancingParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GenerateDatasetBalancingParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionId

`func (o *GenerateDatasetBalancingParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *GenerateDatasetBalancingParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *GenerateDatasetBalancingParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetFilters

`func (o *GenerateDatasetBalancingParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GenerateDatasetBalancingParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GenerateDatasetBalancingParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GenerateDatasetBalancingParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMetadataTags

`func (o *GenerateDatasetBalancingParams) GetMetadataTags() []string`

GetMetadataTags returns the MetadataTags field if non-nil, zero value otherwise.

### GetMetadataTagsOk

`func (o *GenerateDatasetBalancingParams) GetMetadataTagsOk() (*[]string, bool)`

GetMetadataTagsOk returns a tuple with the MetadataTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataTags

`func (o *GenerateDatasetBalancingParams) SetMetadataTags(v []string)`

SetMetadataTags sets MetadataTags field to given value.


### GetPrioritizedMetadataTags

`func (o *GenerateDatasetBalancingParams) GetPrioritizedMetadataTags() []string`

GetPrioritizedMetadataTags returns the PrioritizedMetadataTags field if non-nil, zero value otherwise.

### GetPrioritizedMetadataTagsOk

`func (o *GenerateDatasetBalancingParams) GetPrioritizedMetadataTagsOk() (*[]string, bool)`

GetPrioritizedMetadataTagsOk returns a tuple with the PrioritizedMetadataTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritizedMetadataTags

`func (o *GenerateDatasetBalancingParams) SetPrioritizedMetadataTags(v []string)`

SetPrioritizedMetadataTags sets PrioritizedMetadataTags field to given value.

### HasPrioritizedMetadataTags

`func (o *GenerateDatasetBalancingParams) HasPrioritizedMetadataTags() bool`

HasPrioritizedMetadataTags returns a boolean if a field has been set.

### GetPercentageOfSamplesToPrune

`func (o *GenerateDatasetBalancingParams) GetPercentageOfSamplesToPrune() float64`

GetPercentageOfSamplesToPrune returns the PercentageOfSamplesToPrune field if non-nil, zero value otherwise.

### GetPercentageOfSamplesToPruneOk

`func (o *GenerateDatasetBalancingParams) GetPercentageOfSamplesToPruneOk() (*float64, bool)`

GetPercentageOfSamplesToPruneOk returns a tuple with the PercentageOfSamplesToPrune field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfSamplesToPrune

`func (o *GenerateDatasetBalancingParams) SetPercentageOfSamplesToPrune(v float64)`

SetPercentageOfSamplesToPrune sets PercentageOfSamplesToPrune field to given value.

### HasPercentageOfSamplesToPrune

`func (o *GenerateDatasetBalancingParams) HasPercentageOfSamplesToPrune() bool`

HasPercentageOfSamplesToPrune returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


