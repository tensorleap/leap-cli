# SessionSaveData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Epoch** | **float64** |  | 
**GlobalStep** | **float64** |  | 
**TrainRunNumber** | **float64** |  | 

## Methods

### NewSessionSaveData

`func NewSessionSaveData(epoch float64, globalStep float64, trainRunNumber float64, ) *SessionSaveData`

NewSessionSaveData instantiates a new SessionSaveData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionSaveDataWithDefaults

`func NewSessionSaveDataWithDefaults() *SessionSaveData`

NewSessionSaveDataWithDefaults instantiates a new SessionSaveData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpoch

`func (o *SessionSaveData) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *SessionSaveData) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *SessionSaveData) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetGlobalStep

`func (o *SessionSaveData) GetGlobalStep() float64`

GetGlobalStep returns the GlobalStep field if non-nil, zero value otherwise.

### GetGlobalStepOk

`func (o *SessionSaveData) GetGlobalStepOk() (*float64, bool)`

GetGlobalStepOk returns a tuple with the GlobalStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalStep

`func (o *SessionSaveData) SetGlobalStep(v float64)`

SetGlobalStep sets GlobalStep field to given value.


### GetTrainRunNumber

`func (o *SessionSaveData) GetTrainRunNumber() float64`

GetTrainRunNumber returns the TrainRunNumber field if non-nil, zero value otherwise.

### GetTrainRunNumberOk

`func (o *SessionSaveData) GetTrainRunNumberOk() (*float64, bool)`

GetTrainRunNumberOk returns a tuple with the TrainRunNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainRunNumber

`func (o *SessionSaveData) SetTrainRunNumber(v float64)`

SetTrainRunNumber sets TrainRunNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


