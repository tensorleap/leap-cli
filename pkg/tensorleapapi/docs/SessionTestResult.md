# SessionTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestSucceeded** | **bool** |  | 
**SessionRunId** | **string** |  | 
**Aggregation** | **float64** |  | 
**SuccefullSamples** | **float64** |  | 
**AllSamples** | **float64** |  | 
**Epoch** | **float64** |  | 
**QueryStatus** | **string** |  | 

## Methods

### NewSessionTestResult

`func NewSessionTestResult(testSucceeded bool, sessionRunId string, aggregation float64, succefullSamples float64, allSamples float64, epoch float64, queryStatus string, ) *SessionTestResult`

NewSessionTestResult instantiates a new SessionTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionTestResultWithDefaults

`func NewSessionTestResultWithDefaults() *SessionTestResult`

NewSessionTestResultWithDefaults instantiates a new SessionTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestSucceeded

`func (o *SessionTestResult) GetTestSucceeded() bool`

GetTestSucceeded returns the TestSucceeded field if non-nil, zero value otherwise.

### GetTestSucceededOk

`func (o *SessionTestResult) GetTestSucceededOk() (*bool, bool)`

GetTestSucceededOk returns a tuple with the TestSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSucceeded

`func (o *SessionTestResult) SetTestSucceeded(v bool)`

SetTestSucceeded sets TestSucceeded field to given value.


### GetSessionRunId

`func (o *SessionTestResult) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *SessionTestResult) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *SessionTestResult) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetAggregation

`func (o *SessionTestResult) GetAggregation() float64`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *SessionTestResult) GetAggregationOk() (*float64, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *SessionTestResult) SetAggregation(v float64)`

SetAggregation sets Aggregation field to given value.


### GetSuccefullSamples

`func (o *SessionTestResult) GetSuccefullSamples() float64`

GetSuccefullSamples returns the SuccefullSamples field if non-nil, zero value otherwise.

### GetSuccefullSamplesOk

`func (o *SessionTestResult) GetSuccefullSamplesOk() (*float64, bool)`

GetSuccefullSamplesOk returns a tuple with the SuccefullSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccefullSamples

`func (o *SessionTestResult) SetSuccefullSamples(v float64)`

SetSuccefullSamples sets SuccefullSamples field to given value.


### GetAllSamples

`func (o *SessionTestResult) GetAllSamples() float64`

GetAllSamples returns the AllSamples field if non-nil, zero value otherwise.

### GetAllSamplesOk

`func (o *SessionTestResult) GetAllSamplesOk() (*float64, bool)`

GetAllSamplesOk returns a tuple with the AllSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSamples

`func (o *SessionTestResult) SetAllSamples(v float64)`

SetAllSamples sets AllSamples field to given value.


### GetEpoch

`func (o *SessionTestResult) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *SessionTestResult) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *SessionTestResult) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetQueryStatus

`func (o *SessionTestResult) GetQueryStatus() string`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *SessionTestResult) GetQueryStatusOk() (*string, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *SessionTestResult) SetQueryStatus(v string)`

SetQueryStatus sets QueryStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


