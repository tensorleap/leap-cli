# GenerateAutoSyntheticDataParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**VersionId** | **string** |  | 
**SimulationNames** | **[]string** |  | 
**TargetFilters** | [**[]ESFilter**](ESFilter.md) |  | 
**InitialSimulationFilters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 

## Methods

### NewGenerateAutoSyntheticDataParams

`func NewGenerateAutoSyntheticDataParams(projectId string, versionId string, simulationNames []string, targetFilters []ESFilter, ) *GenerateAutoSyntheticDataParams`

NewGenerateAutoSyntheticDataParams instantiates a new GenerateAutoSyntheticDataParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateAutoSyntheticDataParamsWithDefaults

`func NewGenerateAutoSyntheticDataParamsWithDefaults() *GenerateAutoSyntheticDataParams`

NewGenerateAutoSyntheticDataParamsWithDefaults instantiates a new GenerateAutoSyntheticDataParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GenerateAutoSyntheticDataParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GenerateAutoSyntheticDataParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GenerateAutoSyntheticDataParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionId

`func (o *GenerateAutoSyntheticDataParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *GenerateAutoSyntheticDataParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *GenerateAutoSyntheticDataParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetSimulationNames

`func (o *GenerateAutoSyntheticDataParams) GetSimulationNames() []string`

GetSimulationNames returns the SimulationNames field if non-nil, zero value otherwise.

### GetSimulationNamesOk

`func (o *GenerateAutoSyntheticDataParams) GetSimulationNamesOk() (*[]string, bool)`

GetSimulationNamesOk returns a tuple with the SimulationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationNames

`func (o *GenerateAutoSyntheticDataParams) SetSimulationNames(v []string)`

SetSimulationNames sets SimulationNames field to given value.


### GetTargetFilters

`func (o *GenerateAutoSyntheticDataParams) GetTargetFilters() []ESFilter`

GetTargetFilters returns the TargetFilters field if non-nil, zero value otherwise.

### GetTargetFiltersOk

`func (o *GenerateAutoSyntheticDataParams) GetTargetFiltersOk() (*[]ESFilter, bool)`

GetTargetFiltersOk returns a tuple with the TargetFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFilters

`func (o *GenerateAutoSyntheticDataParams) SetTargetFilters(v []ESFilter)`

SetTargetFilters sets TargetFilters field to given value.


### GetInitialSimulationFilters

`func (o *GenerateAutoSyntheticDataParams) GetInitialSimulationFilters() []ESFilter`

GetInitialSimulationFilters returns the InitialSimulationFilters field if non-nil, zero value otherwise.

### GetInitialSimulationFiltersOk

`func (o *GenerateAutoSyntheticDataParams) GetInitialSimulationFiltersOk() (*[]ESFilter, bool)`

GetInitialSimulationFiltersOk returns a tuple with the InitialSimulationFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSimulationFilters

`func (o *GenerateAutoSyntheticDataParams) SetInitialSimulationFilters(v []ESFilter)`

SetInitialSimulationFilters sets InitialSimulationFilters field to given value.

### HasInitialSimulationFilters

`func (o *GenerateAutoSyntheticDataParams) HasInitialSimulationFilters() bool`

HasInitialSimulationFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


