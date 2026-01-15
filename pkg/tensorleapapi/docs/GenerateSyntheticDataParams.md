# GenerateSyntheticDataParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SessionRunId** | **string** |  | 
**Epoch** | **float64** |  | 
**Sources** | [**[]GenerateSyntheticDataParamsSourcesInner**](GenerateSyntheticDataParamsSourcesInner.md) |  | 
**TargetFilters** | [**[]ESFilter**](ESFilter.md) |  | 

## Methods

### NewGenerateSyntheticDataParams

`func NewGenerateSyntheticDataParams(projectId string, sessionRunId string, epoch float64, sources []GenerateSyntheticDataParamsSourcesInner, targetFilters []ESFilter, ) *GenerateSyntheticDataParams`

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


### GetSessionRunId

`func (o *GenerateSyntheticDataParams) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *GenerateSyntheticDataParams) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *GenerateSyntheticDataParams) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetEpoch

`func (o *GenerateSyntheticDataParams) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *GenerateSyntheticDataParams) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *GenerateSyntheticDataParams) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetSources

`func (o *GenerateSyntheticDataParams) GetSources() []GenerateSyntheticDataParamsSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GenerateSyntheticDataParams) GetSourcesOk() (*[]GenerateSyntheticDataParamsSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GenerateSyntheticDataParams) SetSources(v []GenerateSyntheticDataParamsSourcesInner)`

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


