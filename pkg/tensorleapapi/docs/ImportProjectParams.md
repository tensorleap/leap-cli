# ImportProjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportUrl** | **string** |  | 
**ProjectMeta** | [**ProjectMeta**](ProjectMeta.md) |  | 

## Methods

### NewImportProjectParams

`func NewImportProjectParams(importUrl string, projectMeta ProjectMeta, ) *ImportProjectParams`

NewImportProjectParams instantiates a new ImportProjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportProjectParamsWithDefaults

`func NewImportProjectParamsWithDefaults() *ImportProjectParams`

NewImportProjectParamsWithDefaults instantiates a new ImportProjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportUrl

`func (o *ImportProjectParams) GetImportUrl() string`

GetImportUrl returns the ImportUrl field if non-nil, zero value otherwise.

### GetImportUrlOk

`func (o *ImportProjectParams) GetImportUrlOk() (*string, bool)`

GetImportUrlOk returns a tuple with the ImportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportUrl

`func (o *ImportProjectParams) SetImportUrl(v string)`

SetImportUrl sets ImportUrl field to given value.


### GetProjectMeta

`func (o *ImportProjectParams) GetProjectMeta() ProjectMeta`

GetProjectMeta returns the ProjectMeta field if non-nil, zero value otherwise.

### GetProjectMetaOk

`func (o *ImportProjectParams) GetProjectMetaOk() (*ProjectMeta, bool)`

GetProjectMetaOk returns a tuple with the ProjectMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectMeta

`func (o *ImportProjectParams) SetProjectMeta(v ProjectMeta)`

SetProjectMeta sets ProjectMeta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


