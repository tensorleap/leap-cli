# SessionRunEvaluateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataStates** | [**[]DataStateType**](DataStateType.md) |  | 
**BatchSize** | **float64** |  | 
**EvaluatedEpoch** | **float64** |  | 

## Methods

### NewSessionRunEvaluateParams

`func NewSessionRunEvaluateParams(dataStates []DataStateType, batchSize float64, evaluatedEpoch float64, ) *SessionRunEvaluateParams`

NewSessionRunEvaluateParams instantiates a new SessionRunEvaluateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionRunEvaluateParamsWithDefaults

`func NewSessionRunEvaluateParamsWithDefaults() *SessionRunEvaluateParams`

NewSessionRunEvaluateParamsWithDefaults instantiates a new SessionRunEvaluateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataStates

`func (o *SessionRunEvaluateParams) GetDataStates() []DataStateType`

GetDataStates returns the DataStates field if non-nil, zero value otherwise.

### GetDataStatesOk

`func (o *SessionRunEvaluateParams) GetDataStatesOk() (*[]DataStateType, bool)`

GetDataStatesOk returns a tuple with the DataStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStates

`func (o *SessionRunEvaluateParams) SetDataStates(v []DataStateType)`

SetDataStates sets DataStates field to given value.


### GetBatchSize

`func (o *SessionRunEvaluateParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *SessionRunEvaluateParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *SessionRunEvaluateParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetEvaluatedEpoch

`func (o *SessionRunEvaluateParams) GetEvaluatedEpoch() float64`

GetEvaluatedEpoch returns the EvaluatedEpoch field if non-nil, zero value otherwise.

### GetEvaluatedEpochOk

`func (o *SessionRunEvaluateParams) GetEvaluatedEpochOk() (*float64, bool)`

GetEvaluatedEpochOk returns a tuple with the EvaluatedEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedEpoch

`func (o *SessionRunEvaluateParams) SetEvaluatedEpoch(v float64)`

SetEvaluatedEpoch sets EvaluatedEpoch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


