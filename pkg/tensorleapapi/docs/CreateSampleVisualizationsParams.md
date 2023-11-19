# CreateSampleVisualizationsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SessionRunId** | **string** |  | 
**Epoch** | **float64** |  | 
**SampleIdentities** | [**[]SampleIdentity**](SampleIdentity.md) |  | 

## Methods

### NewCreateSampleVisualizationsParams

`func NewCreateSampleVisualizationsParams(projectId string, sessionRunId string, epoch float64, sampleIdentities []SampleIdentity, ) *CreateSampleVisualizationsParams`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


