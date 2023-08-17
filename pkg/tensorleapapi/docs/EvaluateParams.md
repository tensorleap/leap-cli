# EvaluateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**ProjectId** | **string** |  | 
**SessionId** | **string** |  | 
**BatchSize** | **float64** |  | 
**DataStates** | [**[]DataStateForEval**](DataStateForEval.md) |  | 
**ShouldRunPopulationExploration** | **bool** |  | 
**EvaluatedEpoch** | **float64** |  | 
**Name** | **string** |  | 

## Methods

### NewEvaluateParams

`func NewEvaluateParams(versionId string, projectId string, sessionId string, batchSize float64, dataStates []DataStateForEval, shouldRunPopulationExploration bool, evaluatedEpoch float64, name string, ) *EvaluateParams`

NewEvaluateParams instantiates a new EvaluateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluateParamsWithDefaults

`func NewEvaluateParamsWithDefaults() *EvaluateParams`

NewEvaluateParamsWithDefaults instantiates a new EvaluateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *EvaluateParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *EvaluateParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *EvaluateParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetProjectId

`func (o *EvaluateParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EvaluateParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EvaluateParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSessionId

`func (o *EvaluateParams) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *EvaluateParams) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *EvaluateParams) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetBatchSize

`func (o *EvaluateParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *EvaluateParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *EvaluateParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetDataStates

`func (o *EvaluateParams) GetDataStates() []DataStateForEval`

GetDataStates returns the DataStates field if non-nil, zero value otherwise.

### GetDataStatesOk

`func (o *EvaluateParams) GetDataStatesOk() (*[]DataStateForEval, bool)`

GetDataStatesOk returns a tuple with the DataStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStates

`func (o *EvaluateParams) SetDataStates(v []DataStateForEval)`

SetDataStates sets DataStates field to given value.


### GetShouldRunPopulationExploration

`func (o *EvaluateParams) GetShouldRunPopulationExploration() bool`

GetShouldRunPopulationExploration returns the ShouldRunPopulationExploration field if non-nil, zero value otherwise.

### GetShouldRunPopulationExplorationOk

`func (o *EvaluateParams) GetShouldRunPopulationExplorationOk() (*bool, bool)`

GetShouldRunPopulationExplorationOk returns a tuple with the ShouldRunPopulationExploration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldRunPopulationExploration

`func (o *EvaluateParams) SetShouldRunPopulationExploration(v bool)`

SetShouldRunPopulationExploration sets ShouldRunPopulationExploration field to given value.


### GetEvaluatedEpoch

`func (o *EvaluateParams) GetEvaluatedEpoch() float64`

GetEvaluatedEpoch returns the EvaluatedEpoch field if non-nil, zero value otherwise.

### GetEvaluatedEpochOk

`func (o *EvaluateParams) GetEvaluatedEpochOk() (*float64, bool)`

GetEvaluatedEpochOk returns a tuple with the EvaluatedEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedEpoch

`func (o *EvaluateParams) SetEvaluatedEpoch(v float64)`

SetEvaluatedEpoch sets EvaluatedEpoch field to given value.


### GetName

`func (o *EvaluateParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvaluateParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvaluateParams) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


