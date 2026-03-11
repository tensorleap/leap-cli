# EvaluateExistingVersionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**InferenceArtifactId** | Pointer to **string** |  | [optional] 
**ProjectId** | **string** |  | 
**BatchSize** | **float64** |  | 
**Monitor** | Pointer to **bool** |  | [optional] 
**EvaluatedEpoch** | **float64** |  | 

## Methods

### NewEvaluateExistingVersionParams

`func NewEvaluateExistingVersionParams(versionId string, projectId string, batchSize float64, evaluatedEpoch float64, ) *EvaluateExistingVersionParams`

NewEvaluateExistingVersionParams instantiates a new EvaluateExistingVersionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluateExistingVersionParamsWithDefaults

`func NewEvaluateExistingVersionParamsWithDefaults() *EvaluateExistingVersionParams`

NewEvaluateExistingVersionParamsWithDefaults instantiates a new EvaluateExistingVersionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *EvaluateExistingVersionParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *EvaluateExistingVersionParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *EvaluateExistingVersionParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetInferenceArtifactId

`func (o *EvaluateExistingVersionParams) GetInferenceArtifactId() string`

GetInferenceArtifactId returns the InferenceArtifactId field if non-nil, zero value otherwise.

### GetInferenceArtifactIdOk

`func (o *EvaluateExistingVersionParams) GetInferenceArtifactIdOk() (*string, bool)`

GetInferenceArtifactIdOk returns a tuple with the InferenceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactId

`func (o *EvaluateExistingVersionParams) SetInferenceArtifactId(v string)`

SetInferenceArtifactId sets InferenceArtifactId field to given value.

### HasInferenceArtifactId

`func (o *EvaluateExistingVersionParams) HasInferenceArtifactId() bool`

HasInferenceArtifactId returns a boolean if a field has been set.

### GetProjectId

`func (o *EvaluateExistingVersionParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EvaluateExistingVersionParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EvaluateExistingVersionParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetBatchSize

`func (o *EvaluateExistingVersionParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *EvaluateExistingVersionParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *EvaluateExistingVersionParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetMonitor

`func (o *EvaluateExistingVersionParams) GetMonitor() bool`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *EvaluateExistingVersionParams) GetMonitorOk() (*bool, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *EvaluateExistingVersionParams) SetMonitor(v bool)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *EvaluateExistingVersionParams) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetEvaluatedEpoch

`func (o *EvaluateExistingVersionParams) GetEvaluatedEpoch() float64`

GetEvaluatedEpoch returns the EvaluatedEpoch field if non-nil, zero value otherwise.

### GetEvaluatedEpochOk

`func (o *EvaluateExistingVersionParams) GetEvaluatedEpochOk() (*float64, bool)`

GetEvaluatedEpochOk returns a tuple with the EvaluatedEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedEpoch

`func (o *EvaluateExistingVersionParams) SetEvaluatedEpoch(v float64)`

SetEvaluatedEpoch sets EvaluatedEpoch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


