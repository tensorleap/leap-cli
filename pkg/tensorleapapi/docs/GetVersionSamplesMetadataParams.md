# GetVersionSamplesMetadataParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**[]CollectionFilterSpec**](CollectionFilterSpec.md) |  | [optional] 
**VersionId** | **string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewGetVersionSamplesMetadataParams

`func NewGetVersionSamplesMetadataParams(versionId string, projectId string, ) *GetVersionSamplesMetadataParams`

NewGetVersionSamplesMetadataParams instantiates a new GetVersionSamplesMetadataParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVersionSamplesMetadataParamsWithDefaults

`func NewGetVersionSamplesMetadataParamsWithDefaults() *GetVersionSamplesMetadataParams`

NewGetVersionSamplesMetadataParamsWithDefaults instantiates a new GetVersionSamplesMetadataParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *GetVersionSamplesMetadataParams) GetFilters() []CollectionFilterSpec`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GetVersionSamplesMetadataParams) GetFiltersOk() (*[]CollectionFilterSpec, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GetVersionSamplesMetadataParams) SetFilters(v []CollectionFilterSpec)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GetVersionSamplesMetadataParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetVersionId

`func (o *GetVersionSamplesMetadataParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *GetVersionSamplesMetadataParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *GetVersionSamplesMetadataParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetProjectId

`func (o *GetVersionSamplesMetadataParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetVersionSamplesMetadataParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetVersionSamplesMetadataParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


