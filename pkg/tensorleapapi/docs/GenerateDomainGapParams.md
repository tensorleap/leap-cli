# GenerateDomainGapParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**VersionId** | **string** |  | 
**GroupAFilters** | [**[]ESFilter**](ESFilter.md) |  | 
**GroupBFilters** | [**[]ESFilter**](ESFilter.md) |  | 

## Methods

### NewGenerateDomainGapParams

`func NewGenerateDomainGapParams(projectId string, versionId string, groupAFilters []ESFilter, groupBFilters []ESFilter, ) *GenerateDomainGapParams`

NewGenerateDomainGapParams instantiates a new GenerateDomainGapParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateDomainGapParamsWithDefaults

`func NewGenerateDomainGapParamsWithDefaults() *GenerateDomainGapParams`

NewGenerateDomainGapParamsWithDefaults instantiates a new GenerateDomainGapParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GenerateDomainGapParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GenerateDomainGapParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GenerateDomainGapParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionId

`func (o *GenerateDomainGapParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *GenerateDomainGapParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *GenerateDomainGapParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetGroupAFilters

`func (o *GenerateDomainGapParams) GetGroupAFilters() []ESFilter`

GetGroupAFilters returns the GroupAFilters field if non-nil, zero value otherwise.

### GetGroupAFiltersOk

`func (o *GenerateDomainGapParams) GetGroupAFiltersOk() (*[]ESFilter, bool)`

GetGroupAFiltersOk returns a tuple with the GroupAFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAFilters

`func (o *GenerateDomainGapParams) SetGroupAFilters(v []ESFilter)`

SetGroupAFilters sets GroupAFilters field to given value.


### GetGroupBFilters

`func (o *GenerateDomainGapParams) GetGroupBFilters() []ESFilter`

GetGroupBFilters returns the GroupBFilters field if non-nil, zero value otherwise.

### GetGroupBFiltersOk

`func (o *GenerateDomainGapParams) GetGroupBFiltersOk() (*[]ESFilter, bool)`

GetGroupBFiltersOk returns a tuple with the GroupBFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBFilters

`func (o *GenerateDomainGapParams) SetGroupBFilters(v []ESFilter)`

SetGroupBFilters sets GroupBFilters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


