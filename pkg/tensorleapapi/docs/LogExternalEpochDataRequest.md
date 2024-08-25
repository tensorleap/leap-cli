# LogExternalEpochDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**ExperimentId** | **string** |  | 
**Epoch** | **float64** |  | 
**Metrics** | [**map[string]EpochMetricsValue**](EpochMetricsValue.md) |  | 
**Force** | Pointer to **bool** |  | [optional] 

## Methods

### NewLogExternalEpochDataRequest

`func NewLogExternalEpochDataRequest(projectId string, experimentId string, epoch float64, metrics map[string]EpochMetricsValue, ) *LogExternalEpochDataRequest`

NewLogExternalEpochDataRequest instantiates a new LogExternalEpochDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogExternalEpochDataRequestWithDefaults

`func NewLogExternalEpochDataRequestWithDefaults() *LogExternalEpochDataRequest`

NewLogExternalEpochDataRequestWithDefaults instantiates a new LogExternalEpochDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *LogExternalEpochDataRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *LogExternalEpochDataRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *LogExternalEpochDataRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetExperimentId

`func (o *LogExternalEpochDataRequest) GetExperimentId() string`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *LogExternalEpochDataRequest) GetExperimentIdOk() (*string, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *LogExternalEpochDataRequest) SetExperimentId(v string)`

SetExperimentId sets ExperimentId field to given value.


### GetEpoch

`func (o *LogExternalEpochDataRequest) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *LogExternalEpochDataRequest) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *LogExternalEpochDataRequest) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetMetrics

`func (o *LogExternalEpochDataRequest) GetMetrics() map[string]EpochMetricsValue`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *LogExternalEpochDataRequest) GetMetricsOk() (*map[string]EpochMetricsValue, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *LogExternalEpochDataRequest) SetMetrics(v map[string]EpochMetricsValue)`

SetMetrics sets Metrics field to given value.


### GetForce

`func (o *LogExternalEpochDataRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *LogExternalEpochDataRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *LogExternalEpochDataRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *LogExternalEpochDataRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


