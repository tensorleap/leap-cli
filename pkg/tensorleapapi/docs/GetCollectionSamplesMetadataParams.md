# GetCollectionSamplesMetadataParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**[]CollectionFilterSpec**](CollectionFilterSpec.md) | Active filter spec (same shape used by &#x60;getCollectionSampleOrder&#x60;). Each field&#39;s bucket counts are computed against the other fields&#39; filters, so the sidebar shows the right \&quot;available samples per bucket if you pick this option next\&quot; numbers. | [optional] 
**InferenceArtifactId** | **string** |  | 
**VersionId** | Pointer to **string** |  | [optional] 
**CollectionId** | **string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewGetCollectionSamplesMetadataParams

`func NewGetCollectionSamplesMetadataParams(inferenceArtifactId string, collectionId string, projectId string, ) *GetCollectionSamplesMetadataParams`

NewGetCollectionSamplesMetadataParams instantiates a new GetCollectionSamplesMetadataParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionSamplesMetadataParamsWithDefaults

`func NewGetCollectionSamplesMetadataParamsWithDefaults() *GetCollectionSamplesMetadataParams`

NewGetCollectionSamplesMetadataParamsWithDefaults instantiates a new GetCollectionSamplesMetadataParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *GetCollectionSamplesMetadataParams) GetFilters() []CollectionFilterSpec`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GetCollectionSamplesMetadataParams) GetFiltersOk() (*[]CollectionFilterSpec, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GetCollectionSamplesMetadataParams) SetFilters(v []CollectionFilterSpec)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GetCollectionSamplesMetadataParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetInferenceArtifactId

`func (o *GetCollectionSamplesMetadataParams) GetInferenceArtifactId() string`

GetInferenceArtifactId returns the InferenceArtifactId field if non-nil, zero value otherwise.

### GetInferenceArtifactIdOk

`func (o *GetCollectionSamplesMetadataParams) GetInferenceArtifactIdOk() (*string, bool)`

GetInferenceArtifactIdOk returns a tuple with the InferenceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactId

`func (o *GetCollectionSamplesMetadataParams) SetInferenceArtifactId(v string)`

SetInferenceArtifactId sets InferenceArtifactId field to given value.


### GetVersionId

`func (o *GetCollectionSamplesMetadataParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *GetCollectionSamplesMetadataParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *GetCollectionSamplesMetadataParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *GetCollectionSamplesMetadataParams) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetCollectionId

`func (o *GetCollectionSamplesMetadataParams) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *GetCollectionSamplesMetadataParams) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *GetCollectionSamplesMetadataParams) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.


### GetProjectId

`func (o *GetCollectionSamplesMetadataParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetCollectionSamplesMetadataParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetCollectionSamplesMetadataParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


