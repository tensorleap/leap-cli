# UpdateCurationRunDescriptionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**RunType** | [**CurationRunType**](CurationRunType.md) |  | 
**RunId** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateCurationRunDescriptionParams

`func NewUpdateCurationRunDescriptionParams(projectId string, runType CurationRunType, runId string, ) *UpdateCurationRunDescriptionParams`

NewUpdateCurationRunDescriptionParams instantiates a new UpdateCurationRunDescriptionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCurationRunDescriptionParamsWithDefaults

`func NewUpdateCurationRunDescriptionParamsWithDefaults() *UpdateCurationRunDescriptionParams`

NewUpdateCurationRunDescriptionParamsWithDefaults instantiates a new UpdateCurationRunDescriptionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *UpdateCurationRunDescriptionParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateCurationRunDescriptionParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateCurationRunDescriptionParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetRunType

`func (o *UpdateCurationRunDescriptionParams) GetRunType() CurationRunType`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *UpdateCurationRunDescriptionParams) GetRunTypeOk() (*CurationRunType, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *UpdateCurationRunDescriptionParams) SetRunType(v CurationRunType)`

SetRunType sets RunType field to given value.


### GetRunId

`func (o *UpdateCurationRunDescriptionParams) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *UpdateCurationRunDescriptionParams) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *UpdateCurationRunDescriptionParams) SetRunId(v string)`

SetRunId sets RunId field to given value.


### GetDescription

`func (o *UpdateCurationRunDescriptionParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCurationRunDescriptionParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCurationRunDescriptionParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCurationRunDescriptionParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


