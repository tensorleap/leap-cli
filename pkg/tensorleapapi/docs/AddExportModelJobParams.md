# AddExportModelJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetSessionRunId** | **string** |  | 
**ExportModelType** | [**ExportModelTypeEnum**](ExportModelTypeEnum.md) |  | 
**Title** | **string** |  | 
**PruneModel** | **bool** |  | 
**FromEpoch** | **float64** |  | 

## Methods

### NewAddExportModelJobParams

`func NewAddExportModelJobParams(targetSessionRunId string, exportModelType ExportModelTypeEnum, title string, pruneModel bool, fromEpoch float64, ) *AddExportModelJobParams`

NewAddExportModelJobParams instantiates a new AddExportModelJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddExportModelJobParamsWithDefaults

`func NewAddExportModelJobParamsWithDefaults() *AddExportModelJobParams`

NewAddExportModelJobParamsWithDefaults instantiates a new AddExportModelJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetSessionRunId

`func (o *AddExportModelJobParams) GetTargetSessionRunId() string`

GetTargetSessionRunId returns the TargetSessionRunId field if non-nil, zero value otherwise.

### GetTargetSessionRunIdOk

`func (o *AddExportModelJobParams) GetTargetSessionRunIdOk() (*string, bool)`

GetTargetSessionRunIdOk returns a tuple with the TargetSessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSessionRunId

`func (o *AddExportModelJobParams) SetTargetSessionRunId(v string)`

SetTargetSessionRunId sets TargetSessionRunId field to given value.


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


### GetFromEpoch

`func (o *AddExportModelJobParams) GetFromEpoch() float64`

GetFromEpoch returns the FromEpoch field if non-nil, zero value otherwise.

### GetFromEpochOk

`func (o *AddExportModelJobParams) GetFromEpochOk() (*float64, bool)`

GetFromEpochOk returns a tuple with the FromEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEpoch

`func (o *AddExportModelJobParams) SetFromEpoch(v float64)`

SetFromEpoch sets FromEpoch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


