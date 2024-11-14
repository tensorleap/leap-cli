# UpdateVersionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**ProjectId** | **string** |  | 
**Data** | [**ModelGraph**](ModelGraph.md) |  | 
**CodeIntegration** | Pointer to [**CodeIntegrationBinder**](CodeIntegrationBinder.md) |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateVersionParams

`func NewUpdateVersionParams(versionId string, projectId string, data ModelGraph, ) *UpdateVersionParams`

NewUpdateVersionParams instantiates a new UpdateVersionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVersionParamsWithDefaults

`func NewUpdateVersionParamsWithDefaults() *UpdateVersionParams`

NewUpdateVersionParamsWithDefaults instantiates a new UpdateVersionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *UpdateVersionParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *UpdateVersionParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *UpdateVersionParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetProjectId

`func (o *UpdateVersionParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateVersionParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateVersionParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetData

`func (o *UpdateVersionParams) GetData() ModelGraph`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateVersionParams) GetDataOk() (*ModelGraph, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateVersionParams) SetData(v ModelGraph)`

SetData sets Data field to given value.


### GetCodeIntegration

`func (o *UpdateVersionParams) GetCodeIntegration() CodeIntegrationBinder`

GetCodeIntegration returns the CodeIntegration field if non-nil, zero value otherwise.

### GetCodeIntegrationOk

`func (o *UpdateVersionParams) GetCodeIntegrationOk() (*CodeIntegrationBinder, bool)`

GetCodeIntegrationOk returns a tuple with the CodeIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIntegration

`func (o *UpdateVersionParams) SetCodeIntegration(v CodeIntegrationBinder)`

SetCodeIntegration sets CodeIntegration field to given value.

### HasCodeIntegration

`func (o *UpdateVersionParams) HasCodeIntegration() bool`

HasCodeIntegration returns a boolean if a field has been set.

### GetHash

`func (o *UpdateVersionParams) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *UpdateVersionParams) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *UpdateVersionParams) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *UpdateVersionParams) HasHash() bool`

HasHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


