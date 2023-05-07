# TrainingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Epochs** | **float64** |  | 
**BatchSize** | **float64** |  | 
**EarlyStopParams** | Pointer to [**EarlyStopParams**](EarlyStopParams.md) |  | [optional] 

## Methods

### NewTrainingParams

`func NewTrainingParams(epochs float64, batchSize float64, ) *TrainingParams`

NewTrainingParams instantiates a new TrainingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainingParamsWithDefaults

`func NewTrainingParamsWithDefaults() *TrainingParams`

NewTrainingParamsWithDefaults instantiates a new TrainingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpochs

`func (o *TrainingParams) GetEpochs() float64`

GetEpochs returns the Epochs field if non-nil, zero value otherwise.

### GetEpochsOk

`func (o *TrainingParams) GetEpochsOk() (*float64, bool)`

GetEpochsOk returns a tuple with the Epochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochs

`func (o *TrainingParams) SetEpochs(v float64)`

SetEpochs sets Epochs field to given value.


### GetBatchSize

`func (o *TrainingParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *TrainingParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *TrainingParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetEarlyStopParams

`func (o *TrainingParams) GetEarlyStopParams() EarlyStopParams`

GetEarlyStopParams returns the EarlyStopParams field if non-nil, zero value otherwise.

### GetEarlyStopParamsOk

`func (o *TrainingParams) GetEarlyStopParamsOk() (*EarlyStopParams, bool)`

GetEarlyStopParamsOk returns a tuple with the EarlyStopParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyStopParams

`func (o *TrainingParams) SetEarlyStopParams(v EarlyStopParams)`

SetEarlyStopParams sets EarlyStopParams field to given value.

### HasEarlyStopParams

`func (o *TrainingParams) HasEarlyStopParams() bool`

HasEarlyStopParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


