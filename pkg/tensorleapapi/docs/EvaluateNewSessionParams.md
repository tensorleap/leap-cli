# EvaluateNewSessionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**ProjectId** | **string** |  | 
**BatchSize** | **float64** |  | 
**DataStates** | [**[]DataStateType**](DataStateType.md) |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**SkipMetricsEstimation** | **bool** |  | 

## Methods

### NewEvaluateNewSessionParams

`func NewEvaluateNewSessionParams(versionId string, projectId string, batchSize float64, dataStates []DataStateType, name string, description string, skipMetricsEstimation bool, ) *EvaluateNewSessionParams`

NewEvaluateNewSessionParams instantiates a new EvaluateNewSessionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluateNewSessionParamsWithDefaults

`func NewEvaluateNewSessionParamsWithDefaults() *EvaluateNewSessionParams`

NewEvaluateNewSessionParamsWithDefaults instantiates a new EvaluateNewSessionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *EvaluateNewSessionParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *EvaluateNewSessionParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *EvaluateNewSessionParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetProjectId

`func (o *EvaluateNewSessionParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EvaluateNewSessionParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EvaluateNewSessionParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetBatchSize

`func (o *EvaluateNewSessionParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *EvaluateNewSessionParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *EvaluateNewSessionParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetDataStates

`func (o *EvaluateNewSessionParams) GetDataStates() []DataStateType`

GetDataStates returns the DataStates field if non-nil, zero value otherwise.

### GetDataStatesOk

`func (o *EvaluateNewSessionParams) GetDataStatesOk() (*[]DataStateType, bool)`

GetDataStatesOk returns a tuple with the DataStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStates

`func (o *EvaluateNewSessionParams) SetDataStates(v []DataStateType)`

SetDataStates sets DataStates field to given value.


### GetName

`func (o *EvaluateNewSessionParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvaluateNewSessionParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvaluateNewSessionParams) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EvaluateNewSessionParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EvaluateNewSessionParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EvaluateNewSessionParams) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSkipMetricsEstimation

`func (o *EvaluateNewSessionParams) GetSkipMetricsEstimation() bool`

GetSkipMetricsEstimation returns the SkipMetricsEstimation field if non-nil, zero value otherwise.

### GetSkipMetricsEstimationOk

`func (o *EvaluateNewSessionParams) GetSkipMetricsEstimationOk() (*bool, bool)`

GetSkipMetricsEstimationOk returns a tuple with the SkipMetricsEstimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipMetricsEstimation

`func (o *EvaluateNewSessionParams) SetSkipMetricsEstimation(v bool)`

SetSkipMetricsEstimation sets SkipMetricsEstimation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


