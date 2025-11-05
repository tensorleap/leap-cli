# ImportNewModelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**VersionId** | **string** |  | 
**Model** | [**ImportModelInfo**](ImportModelInfo.md) |  | 

## Methods

### NewImportNewModelParams

`func NewImportNewModelParams(projectId string, versionId string, model ImportModelInfo, ) *ImportNewModelParams`

NewImportNewModelParams instantiates a new ImportNewModelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportNewModelParamsWithDefaults

`func NewImportNewModelParamsWithDefaults() *ImportNewModelParams`

NewImportNewModelParamsWithDefaults instantiates a new ImportNewModelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ImportNewModelParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ImportNewModelParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ImportNewModelParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionId

`func (o *ImportNewModelParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ImportNewModelParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ImportNewModelParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetModel

`func (o *ImportNewModelParams) GetModel() ImportModelInfo`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ImportNewModelParams) GetModelOk() (*ImportModelInfo, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ImportNewModelParams) SetModel(v ImportModelInfo)`

SetModel sets Model field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


