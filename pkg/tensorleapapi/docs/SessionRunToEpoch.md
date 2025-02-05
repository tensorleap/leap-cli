# SessionRunToEpoch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionRunId** | **string** |  | 
**Epoch** | **float64** |  | 

## Methods

### NewSessionRunToEpoch

`func NewSessionRunToEpoch(sessionRunId string, epoch float64, ) *SessionRunToEpoch`

NewSessionRunToEpoch instantiates a new SessionRunToEpoch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionRunToEpochWithDefaults

`func NewSessionRunToEpochWithDefaults() *SessionRunToEpoch`

NewSessionRunToEpochWithDefaults instantiates a new SessionRunToEpoch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionRunId

`func (o *SessionRunToEpoch) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *SessionRunToEpoch) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *SessionRunToEpoch) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetEpoch

`func (o *SessionRunToEpoch) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *SessionRunToEpoch) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *SessionRunToEpoch) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


