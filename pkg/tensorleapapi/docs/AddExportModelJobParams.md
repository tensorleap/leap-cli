# AddExportModelJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SessionWeightId** | **string** |  | 
**ExportModelType** | [**ExportModelTypeEnum**](ExportModelTypeEnum.md) |  | 
**Title** | **string** |  | 
**PruneModel** | **bool** |  | 

## Methods

### NewAddExportModelJobParams

`func NewAddExportModelJobParams(projectId string, sessionWeightId string, exportModelType ExportModelTypeEnum, title string, pruneModel bool, ) *AddExportModelJobParams`

NewAddExportModelJobParams instantiates a new AddExportModelJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddExportModelJobParamsWithDefaults

`func NewAddExportModelJobParamsWithDefaults() *AddExportModelJobParams`

NewAddExportModelJobParamsWithDefaults instantiates a new AddExportModelJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *AddExportModelJobParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AddExportModelJobParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AddExportModelJobParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSessionWeightId

`func (o *AddExportModelJobParams) GetSessionWeightId() string`

GetSessionWeightId returns the SessionWeightId field if non-nil, zero value otherwise.

### GetSessionWeightIdOk

`func (o *AddExportModelJobParams) GetSessionWeightIdOk() (*string, bool)`

GetSessionWeightIdOk returns a tuple with the SessionWeightId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWeightId

`func (o *AddExportModelJobParams) SetSessionWeightId(v string)`

SetSessionWeightId sets SessionWeightId field to given value.


### GetExportModelType

`func (o *AddExportModelJobParams) GetExportModelType() ExportModelTypeEnum`

GetExportModelType returns the ExportModelType field if non-nil, zero value otherwise.

### GetExportModelTypeOk

`func (o *AddExportModelJobParams) GetExportModelTypeOk() (*ExportModelTypeEnum, bool)`

GetExportModelTypeOk returns a tuple with the ExportModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportModelType

`func (o *AddExportModelJobParams) SetExportModelType(v ExportModelTypeEnum)`

SetExportModelType sets ExportModelType field to given value.


### GetTitle

`func (o *AddExportModelJobParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AddExportModelJobParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AddExportModelJobParams) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPruneModel

`func (o *AddExportModelJobParams) GetPruneModel() bool`

GetPruneModel returns the PruneModel field if non-nil, zero value otherwise.

### GetPruneModelOk

`func (o *AddExportModelJobParams) GetPruneModelOk() (*bool, bool)`

GetPruneModelOk returns a tuple with the PruneModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruneModel

`func (o *AddExportModelJobParams) SetPruneModel(v bool)`

SetPruneModel sets PruneModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


