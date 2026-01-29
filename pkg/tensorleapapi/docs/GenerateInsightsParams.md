# GenerateInsightsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SessionRunId** | **string** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**Refresh** | Pointer to **bool** |  | [optional] 

## Methods

### NewGenerateInsightsParams

`func NewGenerateInsightsParams(projectId string, sessionRunId string, ) *GenerateInsightsParams`

NewGenerateInsightsParams instantiates a new GenerateInsightsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateInsightsParamsWithDefaults

`func NewGenerateInsightsParamsWithDefaults() *GenerateInsightsParams`

NewGenerateInsightsParamsWithDefaults instantiates a new GenerateInsightsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GenerateInsightsParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GenerateInsightsParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GenerateInsightsParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSessionRunId

`func (o *GenerateInsightsParams) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *GenerateInsightsParams) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *GenerateInsightsParams) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetFilters

`func (o *GenerateInsightsParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GenerateInsightsParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GenerateInsightsParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GenerateInsightsParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetRefresh

`func (o *GenerateInsightsParams) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *GenerateInsightsParams) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *GenerateInsightsParams) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *GenerateInsightsParams) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


