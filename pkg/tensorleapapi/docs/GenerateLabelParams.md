# GenerateLabelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**VersionId** | **string** |  | 
**NumOfSamplesToLabel** | Pointer to **float64** |  | [optional] 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 

## Methods

### NewGenerateLabelParams

`func NewGenerateLabelParams(projectId string, versionId string, ) *GenerateLabelParams`

NewGenerateLabelParams instantiates a new GenerateLabelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateLabelParamsWithDefaults

`func NewGenerateLabelParamsWithDefaults() *GenerateLabelParams`

NewGenerateLabelParamsWithDefaults instantiates a new GenerateLabelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GenerateLabelParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GenerateLabelParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GenerateLabelParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionId

`func (o *GenerateLabelParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *GenerateLabelParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *GenerateLabelParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetNumOfSamplesToLabel

`func (o *GenerateLabelParams) GetNumOfSamplesToLabel() float64`

GetNumOfSamplesToLabel returns the NumOfSamplesToLabel field if non-nil, zero value otherwise.

### GetNumOfSamplesToLabelOk

`func (o *GenerateLabelParams) GetNumOfSamplesToLabelOk() (*float64, bool)`

GetNumOfSamplesToLabelOk returns a tuple with the NumOfSamplesToLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfSamplesToLabel

`func (o *GenerateLabelParams) SetNumOfSamplesToLabel(v float64)`

SetNumOfSamplesToLabel sets NumOfSamplesToLabel field to given value.

### HasNumOfSamplesToLabel

`func (o *GenerateLabelParams) HasNumOfSamplesToLabel() bool`

HasNumOfSamplesToLabel returns a boolean if a field has been set.

### GetFilters

`func (o *GenerateLabelParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GenerateLabelParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GenerateLabelParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GenerateLabelParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


