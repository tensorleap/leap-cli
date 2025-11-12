# CreateSampleVisualizationsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SessionRunId** | **string** |  | 
**Epoch** | **float64** |  | 
**SampleIdentities** | Pointer to [**[]SampleIdentity**](SampleIdentity.md) |  | [optional] 
**Digest** | **string** |  | 
**Refresh** | Pointer to **bool** |  | [optional] 
**Trigger** | Pointer to [**JobTrigger**](JobTrigger.md) |  | [optional] 

## Methods

### NewCreateSampleVisualizationsParams

`func NewCreateSampleVisualizationsParams(projectId string, sessionRunId string, epoch float64, digest string, ) *CreateSampleVisualizationsParams`

NewCreateSampleVisualizationsParams instantiates a new CreateSampleVisualizationsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSampleVisualizationsParamsWithDefaults

`func NewCreateSampleVisualizationsParamsWithDefaults() *CreateSampleVisualizationsParams`

NewCreateSampleVisualizationsParamsWithDefaults instantiates a new CreateSampleVisualizationsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *CreateSampleVisualizationsParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateSampleVisualizationsParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateSampleVisualizationsParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSessionRunId

`func (o *CreateSampleVisualizationsParams) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *CreateSampleVisualizationsParams) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *CreateSampleVisualizationsParams) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetEpoch

`func (o *CreateSampleVisualizationsParams) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *CreateSampleVisualizationsParams) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *CreateSampleVisualizationsParams) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetSampleIdentities

`func (o *CreateSampleVisualizationsParams) GetSampleIdentities() []SampleIdentity`

GetSampleIdentities returns the SampleIdentities field if non-nil, zero value otherwise.

### GetSampleIdentitiesOk

`func (o *CreateSampleVisualizationsParams) GetSampleIdentitiesOk() (*[]SampleIdentity, bool)`

GetSampleIdentitiesOk returns a tuple with the SampleIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIdentities

`func (o *CreateSampleVisualizationsParams) SetSampleIdentities(v []SampleIdentity)`

SetSampleIdentities sets SampleIdentities field to given value.

### HasSampleIdentities

`func (o *CreateSampleVisualizationsParams) HasSampleIdentities() bool`

HasSampleIdentities returns a boolean if a field has been set.

### GetDigest

`func (o *CreateSampleVisualizationsParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *CreateSampleVisualizationsParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *CreateSampleVisualizationsParams) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetRefresh

`func (o *CreateSampleVisualizationsParams) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *CreateSampleVisualizationsParams) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *CreateSampleVisualizationsParams) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *CreateSampleVisualizationsParams) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetTrigger

`func (o *CreateSampleVisualizationsParams) GetTrigger() JobTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CreateSampleVisualizationsParams) GetTriggerOk() (*JobTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CreateSampleVisualizationsParams) SetTrigger(v JobTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *CreateSampleVisualizationsParams) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


