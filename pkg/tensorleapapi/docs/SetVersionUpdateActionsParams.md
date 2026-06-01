# SetVersionUpdateActionsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**ProjectId** | **string** |  | 
**UpdateActions** | [**[]UpdateAction**](UpdateAction.md) |  | 

## Methods

### NewSetVersionUpdateActionsParams

`func NewSetVersionUpdateActionsParams(versionId string, projectId string, updateActions []UpdateAction, ) *SetVersionUpdateActionsParams`

NewSetVersionUpdateActionsParams instantiates a new SetVersionUpdateActionsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetVersionUpdateActionsParamsWithDefaults

`func NewSetVersionUpdateActionsParamsWithDefaults() *SetVersionUpdateActionsParams`

NewSetVersionUpdateActionsParamsWithDefaults instantiates a new SetVersionUpdateActionsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *SetVersionUpdateActionsParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *SetVersionUpdateActionsParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *SetVersionUpdateActionsParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetProjectId

`func (o *SetVersionUpdateActionsParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SetVersionUpdateActionsParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SetVersionUpdateActionsParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetUpdateActions

`func (o *SetVersionUpdateActionsParams) GetUpdateActions() []UpdateAction`

GetUpdateActions returns the UpdateActions field if non-nil, zero value otherwise.

### GetUpdateActionsOk

`func (o *SetVersionUpdateActionsParams) GetUpdateActionsOk() (*[]UpdateAction, bool)`

GetUpdateActionsOk returns a tuple with the UpdateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateActions

`func (o *SetVersionUpdateActionsParams) SetUpdateActions(v []UpdateAction)`

SetUpdateActions sets UpdateActions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


