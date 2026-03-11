# GenerateSyntheticDataParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**VersionId** | **string** |  | 
**Sources** | [**[]SyntheticDataJobParamsSourcesInner**](SyntheticDataJobParamsSourcesInner.md) |  | 
**TargetFilters** | [**[]ESFilter**](ESFilter.md) |  | 

## Methods

### NewGenerateSyntheticDataParams

`func NewGenerateSyntheticDataParams(projectId string, versionId string, sources []SyntheticDataJobParamsSourcesInner, targetFilters []ESFilter, ) *GenerateSyntheticDataParams`

NewGenerateSyntheticDataParams instantiates a new GenerateSyntheticDataParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateSyntheticDataParamsWithDefaults

`func NewGenerateSyntheticDataParamsWithDefaults() *GenerateSyntheticDataParams`

NewGenerateSyntheticDataParamsWithDefaults instantiates a new GenerateSyntheticDataParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GenerateSyntheticDataParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GenerateSyntheticDataParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GenerateSyntheticDataParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionId

`func (o *GenerateSyntheticDataParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *GenerateSyntheticDataParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *GenerateSyntheticDataParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetSources

`func (o *GenerateSyntheticDataParams) GetSources() []SyntheticDataJobParamsSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GenerateSyntheticDataParams) GetSourcesOk() (*[]SyntheticDataJobParamsSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GenerateSyntheticDataParams) SetSources(v []SyntheticDataJobParamsSourcesInner)`

SetSources sets Sources field to given value.


### GetTargetFilters

`func (o *GenerateSyntheticDataParams) GetTargetFilters() []ESFilter`

GetTargetFilters returns the TargetFilters field if non-nil, zero value otherwise.

### GetTargetFiltersOk

`func (o *GenerateSyntheticDataParams) GetTargetFiltersOk() (*[]ESFilter, bool)`

GetTargetFiltersOk returns a tuple with the TargetFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFilters

`func (o *GenerateSyntheticDataParams) SetTargetFilters(v []ESFilter)`

SetTargetFilters sets TargetFilters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


