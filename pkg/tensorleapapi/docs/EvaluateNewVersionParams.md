# EvaluateNewVersionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**InferenceArtifactId** | Pointer to **string** |  | [optional] 
**ProjectId** | **string** |  | 
**BatchSize** | **float64** |  | 
**Monitor** | Pointer to **bool** |  | [optional] 

## Methods

### NewEvaluateNewVersionParams

`func NewEvaluateNewVersionParams(versionId string, projectId string, batchSize float64, ) *EvaluateNewVersionParams`

NewEvaluateNewVersionParams instantiates a new EvaluateNewVersionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluateNewVersionParamsWithDefaults

`func NewEvaluateNewVersionParamsWithDefaults() *EvaluateNewVersionParams`

NewEvaluateNewVersionParamsWithDefaults instantiates a new EvaluateNewVersionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *EvaluateNewVersionParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *EvaluateNewVersionParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *EvaluateNewVersionParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetInferenceArtifactId

`func (o *EvaluateNewVersionParams) GetInferenceArtifactId() string`

GetInferenceArtifactId returns the InferenceArtifactId field if non-nil, zero value otherwise.

### GetInferenceArtifactIdOk

`func (o *EvaluateNewVersionParams) GetInferenceArtifactIdOk() (*string, bool)`

GetInferenceArtifactIdOk returns a tuple with the InferenceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactId

`func (o *EvaluateNewVersionParams) SetInferenceArtifactId(v string)`

SetInferenceArtifactId sets InferenceArtifactId field to given value.

### HasInferenceArtifactId

`func (o *EvaluateNewVersionParams) HasInferenceArtifactId() bool`

HasInferenceArtifactId returns a boolean if a field has been set.

### GetProjectId

`func (o *EvaluateNewVersionParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EvaluateNewVersionParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EvaluateNewVersionParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetBatchSize

`func (o *EvaluateNewVersionParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *EvaluateNewVersionParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *EvaluateNewVersionParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetMonitor

`func (o *EvaluateNewVersionParams) GetMonitor() bool`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *EvaluateNewVersionParams) GetMonitorOk() (*bool, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *EvaluateNewVersionParams) SetMonitor(v bool)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *EvaluateNewVersionParams) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


