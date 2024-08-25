# EvaluateExistingSessionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**ProjectId** | **string** |  | 
**BatchSize** | **float64** |  | 
**DataStates** | [**[]DataStateType**](DataStateType.md) |  | 
**Name** | **string** |  | 
**SkipMetricsEstimation** | **bool** |  | 
**EvaluatedEpoch** | **float64** |  | 
**SessionId** | **string** |  | 

## Methods

### NewEvaluateExistingSessionParams

`func NewEvaluateExistingSessionParams(versionId string, projectId string, batchSize float64, dataStates []DataStateType, name string, skipMetricsEstimation bool, evaluatedEpoch float64, sessionId string, ) *EvaluateExistingSessionParams`

NewEvaluateExistingSessionParams instantiates a new EvaluateExistingSessionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluateExistingSessionParamsWithDefaults

`func NewEvaluateExistingSessionParamsWithDefaults() *EvaluateExistingSessionParams`

NewEvaluateExistingSessionParamsWithDefaults instantiates a new EvaluateExistingSessionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *EvaluateExistingSessionParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *EvaluateExistingSessionParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *EvaluateExistingSessionParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetProjectId

`func (o *EvaluateExistingSessionParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EvaluateExistingSessionParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EvaluateExistingSessionParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetBatchSize

`func (o *EvaluateExistingSessionParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *EvaluateExistingSessionParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *EvaluateExistingSessionParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetDataStates

`func (o *EvaluateExistingSessionParams) GetDataStates() []DataStateType`

GetDataStates returns the DataStates field if non-nil, zero value otherwise.

### GetDataStatesOk

`func (o *EvaluateExistingSessionParams) GetDataStatesOk() (*[]DataStateType, bool)`

GetDataStatesOk returns a tuple with the DataStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStates

`func (o *EvaluateExistingSessionParams) SetDataStates(v []DataStateType)`

SetDataStates sets DataStates field to given value.


### GetName

`func (o *EvaluateExistingSessionParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvaluateExistingSessionParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvaluateExistingSessionParams) SetName(v string)`

SetName sets Name field to given value.


### GetSkipMetricsEstimation

`func (o *EvaluateExistingSessionParams) GetSkipMetricsEstimation() bool`

GetSkipMetricsEstimation returns the SkipMetricsEstimation field if non-nil, zero value otherwise.

### GetSkipMetricsEstimationOk

`func (o *EvaluateExistingSessionParams) GetSkipMetricsEstimationOk() (*bool, bool)`

GetSkipMetricsEstimationOk returns a tuple with the SkipMetricsEstimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipMetricsEstimation

`func (o *EvaluateExistingSessionParams) SetSkipMetricsEstimation(v bool)`

SetSkipMetricsEstimation sets SkipMetricsEstimation field to given value.


### GetEvaluatedEpoch

`func (o *EvaluateExistingSessionParams) GetEvaluatedEpoch() float64`

GetEvaluatedEpoch returns the EvaluatedEpoch field if non-nil, zero value otherwise.

### GetEvaluatedEpochOk

`func (o *EvaluateExistingSessionParams) GetEvaluatedEpochOk() (*float64, bool)`

GetEvaluatedEpochOk returns a tuple with the EvaluatedEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedEpoch

`func (o *EvaluateExistingSessionParams) SetEvaluatedEpoch(v float64)`

SetEvaluatedEpoch sets EvaluatedEpoch field to given value.


### GetSessionId

`func (o *EvaluateExistingSessionParams) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *EvaluateExistingSessionParams) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *EvaluateExistingSessionParams) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


