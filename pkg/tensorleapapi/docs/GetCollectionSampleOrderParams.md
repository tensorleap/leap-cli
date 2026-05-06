# GetCollectionSampleOrderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageSize** | Pointer to **float64** |  | [optional] 
**Cursor** | Pointer to **string** | Opaque pagination token; here it&#39;s a stringified &#x60;from&#x60; offset. | [optional] 
**Filters** | Pointer to [**[]CollectionFilterSpec**](CollectionFilterSpec.md) |  | [optional] 
**Sort** | Pointer to [**CollectionSortSpec**](CollectionSortSpec.md) |  | [optional] 
**InferenceArtifactId** | **string** |  | 
**CollectionId** | **string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewGetCollectionSampleOrderParams

`func NewGetCollectionSampleOrderParams(inferenceArtifactId string, collectionId string, projectId string, ) *GetCollectionSampleOrderParams`

NewGetCollectionSampleOrderParams instantiates a new GetCollectionSampleOrderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionSampleOrderParamsWithDefaults

`func NewGetCollectionSampleOrderParamsWithDefaults() *GetCollectionSampleOrderParams`

NewGetCollectionSampleOrderParamsWithDefaults instantiates a new GetCollectionSampleOrderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageSize

`func (o *GetCollectionSampleOrderParams) GetPageSize() float64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetCollectionSampleOrderParams) GetPageSizeOk() (*float64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetCollectionSampleOrderParams) SetPageSize(v float64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetCollectionSampleOrderParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetCursor

`func (o *GetCollectionSampleOrderParams) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetCollectionSampleOrderParams) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetCollectionSampleOrderParams) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetCollectionSampleOrderParams) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetFilters

`func (o *GetCollectionSampleOrderParams) GetFilters() []CollectionFilterSpec`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GetCollectionSampleOrderParams) GetFiltersOk() (*[]CollectionFilterSpec, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GetCollectionSampleOrderParams) SetFilters(v []CollectionFilterSpec)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GetCollectionSampleOrderParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSort

`func (o *GetCollectionSampleOrderParams) GetSort() CollectionSortSpec`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GetCollectionSampleOrderParams) GetSortOk() (*CollectionSortSpec, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GetCollectionSampleOrderParams) SetSort(v CollectionSortSpec)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GetCollectionSampleOrderParams) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetInferenceArtifactId

`func (o *GetCollectionSampleOrderParams) GetInferenceArtifactId() string`

GetInferenceArtifactId returns the InferenceArtifactId field if non-nil, zero value otherwise.

### GetInferenceArtifactIdOk

`func (o *GetCollectionSampleOrderParams) GetInferenceArtifactIdOk() (*string, bool)`

GetInferenceArtifactIdOk returns a tuple with the InferenceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactId

`func (o *GetCollectionSampleOrderParams) SetInferenceArtifactId(v string)`

SetInferenceArtifactId sets InferenceArtifactId field to given value.


### GetCollectionId

`func (o *GetCollectionSampleOrderParams) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *GetCollectionSampleOrderParams) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *GetCollectionSampleOrderParams) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.


### GetProjectId

`func (o *GetCollectionSampleOrderParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetCollectionSampleOrderParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetCollectionSampleOrderParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


