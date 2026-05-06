# GetVersionSampleOrderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageSize** | Pointer to **float64** |  | [optional] 
**Cursor** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]CollectionFilterSpec**](CollectionFilterSpec.md) |  | [optional] 
**Sort** | Pointer to [**CollectionSortSpec**](CollectionSortSpec.md) |  | [optional] 
**VersionId** | **string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewGetVersionSampleOrderParams

`func NewGetVersionSampleOrderParams(versionId string, projectId string, ) *GetVersionSampleOrderParams`

NewGetVersionSampleOrderParams instantiates a new GetVersionSampleOrderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVersionSampleOrderParamsWithDefaults

`func NewGetVersionSampleOrderParamsWithDefaults() *GetVersionSampleOrderParams`

NewGetVersionSampleOrderParamsWithDefaults instantiates a new GetVersionSampleOrderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageSize

`func (o *GetVersionSampleOrderParams) GetPageSize() float64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetVersionSampleOrderParams) GetPageSizeOk() (*float64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetVersionSampleOrderParams) SetPageSize(v float64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetVersionSampleOrderParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetCursor

`func (o *GetVersionSampleOrderParams) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetVersionSampleOrderParams) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetVersionSampleOrderParams) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetVersionSampleOrderParams) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetFilters

`func (o *GetVersionSampleOrderParams) GetFilters() []CollectionFilterSpec`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GetVersionSampleOrderParams) GetFiltersOk() (*[]CollectionFilterSpec, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GetVersionSampleOrderParams) SetFilters(v []CollectionFilterSpec)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GetVersionSampleOrderParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSort

`func (o *GetVersionSampleOrderParams) GetSort() CollectionSortSpec`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GetVersionSampleOrderParams) GetSortOk() (*CollectionSortSpec, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GetVersionSampleOrderParams) SetSort(v CollectionSortSpec)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GetVersionSampleOrderParams) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetVersionId

`func (o *GetVersionSampleOrderParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *GetVersionSampleOrderParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *GetVersionSampleOrderParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetProjectId

`func (o *GetVersionSampleOrderParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetVersionSampleOrderParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetVersionSampleOrderParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


