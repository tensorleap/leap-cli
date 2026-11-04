# SamplesCsvExportViewParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**[]CollectionFilterSpec**](CollectionFilterSpec.md) |  | [optional] 
**Sort** | Pointer to [**CollectionSortSpec**](CollectionSortSpec.md) |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**InferenceArtifactId** | Pointer to **string** |  | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**ProjectId** | **string** |  | 

## Methods

### NewSamplesCsvExportViewParams

`func NewSamplesCsvExportViewParams(projectId string, ) *SamplesCsvExportViewParams`

NewSamplesCsvExportViewParams instantiates a new SamplesCsvExportViewParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamplesCsvExportViewParamsWithDefaults

`func NewSamplesCsvExportViewParamsWithDefaults() *SamplesCsvExportViewParams`

NewSamplesCsvExportViewParamsWithDefaults instantiates a new SamplesCsvExportViewParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *SamplesCsvExportViewParams) GetFilters() []CollectionFilterSpec`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SamplesCsvExportViewParams) GetFiltersOk() (*[]CollectionFilterSpec, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SamplesCsvExportViewParams) SetFilters(v []CollectionFilterSpec)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SamplesCsvExportViewParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSort

`func (o *SamplesCsvExportViewParams) GetSort() CollectionSortSpec`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SamplesCsvExportViewParams) GetSortOk() (*CollectionSortSpec, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SamplesCsvExportViewParams) SetSort(v CollectionSortSpec)`

SetSort sets Sort field to given value.

### HasSort

`func (o *SamplesCsvExportViewParams) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetVersionId

`func (o *SamplesCsvExportViewParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *SamplesCsvExportViewParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *SamplesCsvExportViewParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *SamplesCsvExportViewParams) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetInferenceArtifactId

`func (o *SamplesCsvExportViewParams) GetInferenceArtifactId() string`

GetInferenceArtifactId returns the InferenceArtifactId field if non-nil, zero value otherwise.

### GetInferenceArtifactIdOk

`func (o *SamplesCsvExportViewParams) GetInferenceArtifactIdOk() (*string, bool)`

GetInferenceArtifactIdOk returns a tuple with the InferenceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactId

`func (o *SamplesCsvExportViewParams) SetInferenceArtifactId(v string)`

SetInferenceArtifactId sets InferenceArtifactId field to given value.

### HasInferenceArtifactId

`func (o *SamplesCsvExportViewParams) HasInferenceArtifactId() bool`

HasInferenceArtifactId returns a boolean if a field has been set.

### GetCollectionId

`func (o *SamplesCsvExportViewParams) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *SamplesCsvExportViewParams) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *SamplesCsvExportViewParams) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *SamplesCsvExportViewParams) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetProjectId

`func (o *SamplesCsvExportViewParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SamplesCsvExportViewParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SamplesCsvExportViewParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


