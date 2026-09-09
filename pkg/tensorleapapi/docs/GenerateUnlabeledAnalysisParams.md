# GenerateUnlabeledAnalysisParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**VersionId** | **string** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**LatentSpaceType** | Pointer to **string** | Latent space friendly name. Engine falls back to default when unset. | [optional] 
**ElementInstance** | Pointer to **bool** |  | [optional] 

## Methods

### NewGenerateUnlabeledAnalysisParams

`func NewGenerateUnlabeledAnalysisParams(projectId string, versionId string, ) *GenerateUnlabeledAnalysisParams`

NewGenerateUnlabeledAnalysisParams instantiates a new GenerateUnlabeledAnalysisParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateUnlabeledAnalysisParamsWithDefaults

`func NewGenerateUnlabeledAnalysisParamsWithDefaults() *GenerateUnlabeledAnalysisParams`

NewGenerateUnlabeledAnalysisParamsWithDefaults instantiates a new GenerateUnlabeledAnalysisParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GenerateUnlabeledAnalysisParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GenerateUnlabeledAnalysisParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GenerateUnlabeledAnalysisParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionId

`func (o *GenerateUnlabeledAnalysisParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *GenerateUnlabeledAnalysisParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *GenerateUnlabeledAnalysisParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetFilters

`func (o *GenerateUnlabeledAnalysisParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GenerateUnlabeledAnalysisParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GenerateUnlabeledAnalysisParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GenerateUnlabeledAnalysisParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLatentSpaceType

`func (o *GenerateUnlabeledAnalysisParams) GetLatentSpaceType() string`

GetLatentSpaceType returns the LatentSpaceType field if non-nil, zero value otherwise.

### GetLatentSpaceTypeOk

`func (o *GenerateUnlabeledAnalysisParams) GetLatentSpaceTypeOk() (*string, bool)`

GetLatentSpaceTypeOk returns a tuple with the LatentSpaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatentSpaceType

`func (o *GenerateUnlabeledAnalysisParams) SetLatentSpaceType(v string)`

SetLatentSpaceType sets LatentSpaceType field to given value.

### HasLatentSpaceType

`func (o *GenerateUnlabeledAnalysisParams) HasLatentSpaceType() bool`

HasLatentSpaceType returns a boolean if a field has been set.

### GetElementInstance

`func (o *GenerateUnlabeledAnalysisParams) GetElementInstance() bool`

GetElementInstance returns the ElementInstance field if non-nil, zero value otherwise.

### GetElementInstanceOk

`func (o *GenerateUnlabeledAnalysisParams) GetElementInstanceOk() (*bool, bool)`

GetElementInstanceOk returns a tuple with the ElementInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementInstance

`func (o *GenerateUnlabeledAnalysisParams) SetElementInstance(v bool)`

SetElementInstance sets ElementInstance field to given value.

### HasElementInstance

`func (o *GenerateUnlabeledAnalysisParams) HasElementInstance() bool`

HasElementInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


