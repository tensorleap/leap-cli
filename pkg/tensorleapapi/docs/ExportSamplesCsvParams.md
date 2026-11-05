# ExportSamplesCsvParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**[]CollectionFilterSpec**](CollectionFilterSpec.md) |  | [optional] 
**Sort** | Pointer to [**CollectionSortSpec**](CollectionSortSpec.md) |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**InferenceArtifactId** | Pointer to **string** |  | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**ProjectId** | **string** |  | 
**MaxRows** | Pointer to **float64** | Write at most this many rows: the first ones in sort order. The only way to export a view bigger than the cap (&#x60;SAMPLES_CSV_EXPORT_MAX_ROWS&#x60;, 250,000 by default), and clamped to it. | [optional] 
**Columns** | [**[]SampleCsvColumn**](SampleCsvColumn.md) |  | 

## Methods

### NewExportSamplesCsvParams

`func NewExportSamplesCsvParams(projectId string, columns []SampleCsvColumn, ) *ExportSamplesCsvParams`

NewExportSamplesCsvParams instantiates a new ExportSamplesCsvParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportSamplesCsvParamsWithDefaults

`func NewExportSamplesCsvParamsWithDefaults() *ExportSamplesCsvParams`

NewExportSamplesCsvParamsWithDefaults instantiates a new ExportSamplesCsvParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ExportSamplesCsvParams) GetFilters() []CollectionFilterSpec`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ExportSamplesCsvParams) GetFiltersOk() (*[]CollectionFilterSpec, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ExportSamplesCsvParams) SetFilters(v []CollectionFilterSpec)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ExportSamplesCsvParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSort

`func (o *ExportSamplesCsvParams) GetSort() CollectionSortSpec`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ExportSamplesCsvParams) GetSortOk() (*CollectionSortSpec, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ExportSamplesCsvParams) SetSort(v CollectionSortSpec)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ExportSamplesCsvParams) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetVersionId

`func (o *ExportSamplesCsvParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ExportSamplesCsvParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ExportSamplesCsvParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *ExportSamplesCsvParams) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetInferenceArtifactId

`func (o *ExportSamplesCsvParams) GetInferenceArtifactId() string`

GetInferenceArtifactId returns the InferenceArtifactId field if non-nil, zero value otherwise.

### GetInferenceArtifactIdOk

`func (o *ExportSamplesCsvParams) GetInferenceArtifactIdOk() (*string, bool)`

GetInferenceArtifactIdOk returns a tuple with the InferenceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactId

`func (o *ExportSamplesCsvParams) SetInferenceArtifactId(v string)`

SetInferenceArtifactId sets InferenceArtifactId field to given value.

### HasInferenceArtifactId

`func (o *ExportSamplesCsvParams) HasInferenceArtifactId() bool`

HasInferenceArtifactId returns a boolean if a field has been set.

### GetCollectionId

`func (o *ExportSamplesCsvParams) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *ExportSamplesCsvParams) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *ExportSamplesCsvParams) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *ExportSamplesCsvParams) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetProjectId

`func (o *ExportSamplesCsvParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ExportSamplesCsvParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ExportSamplesCsvParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetMaxRows

`func (o *ExportSamplesCsvParams) GetMaxRows() float64`

GetMaxRows returns the MaxRows field if non-nil, zero value otherwise.

### GetMaxRowsOk

`func (o *ExportSamplesCsvParams) GetMaxRowsOk() (*float64, bool)`

GetMaxRowsOk returns a tuple with the MaxRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRows

`func (o *ExportSamplesCsvParams) SetMaxRows(v float64)`

SetMaxRows sets MaxRows field to given value.

### HasMaxRows

`func (o *ExportSamplesCsvParams) HasMaxRows() bool`

HasMaxRows returns a boolean if a field has been set.

### GetColumns

`func (o *ExportSamplesCsvParams) GetColumns() []SampleCsvColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ExportSamplesCsvParams) GetColumnsOk() (*[]SampleCsvColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ExportSamplesCsvParams) SetColumns(v []SampleCsvColumn)`

SetColumns sets Columns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


