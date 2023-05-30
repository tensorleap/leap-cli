# SlimVisualization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**JobId** | **string** |  | 
**SessionRunId** | **string** |  | 
**JobParms** | Pointer to [**JobParams**](JobParams.md) |  | [optional] 
**Type** | [**AnalyzeTypeEnum**](AnalyzeTypeEnum.md) |  | 
**CreatedAt** | **time.Time** |  | 
**Epoch** | **float64** |  | 
**SampleId** | Pointer to **float64** |  | [optional] 
**Layout** | Pointer to [**SizedLayout**](SizedLayout.md) |  | [optional] 
**VisualizationUuid** | **string** |  | 
**Blob** | **string** |  | 

## Methods

### NewSlimVisualization

`func NewSlimVisualization(id string, jobId string, sessionRunId string, type_ AnalyzeTypeEnum, createdAt time.Time, epoch float64, visualizationUuid string, blob string, ) *SlimVisualization`

NewSlimVisualization instantiates a new SlimVisualization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlimVisualizationWithDefaults

`func NewSlimVisualizationWithDefaults() *SlimVisualization`

NewSlimVisualizationWithDefaults instantiates a new SlimVisualization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SlimVisualization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlimVisualization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlimVisualization) SetId(v string)`

SetId sets Id field to given value.


### GetJobId

`func (o *SlimVisualization) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SlimVisualization) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SlimVisualization) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetSessionRunId

`func (o *SlimVisualization) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *SlimVisualization) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *SlimVisualization) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetJobParms

`func (o *SlimVisualization) GetJobParms() JobParams`

GetJobParms returns the JobParms field if non-nil, zero value otherwise.

### GetJobParmsOk

`func (o *SlimVisualization) GetJobParmsOk() (*JobParams, bool)`

GetJobParmsOk returns a tuple with the JobParms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobParms

`func (o *SlimVisualization) SetJobParms(v JobParams)`

SetJobParms sets JobParms field to given value.

### HasJobParms

`func (o *SlimVisualization) HasJobParms() bool`

HasJobParms returns a boolean if a field has been set.

### GetType

`func (o *SlimVisualization) GetType() AnalyzeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlimVisualization) GetTypeOk() (*AnalyzeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlimVisualization) SetType(v AnalyzeTypeEnum)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *SlimVisualization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SlimVisualization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SlimVisualization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEpoch

`func (o *SlimVisualization) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *SlimVisualization) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *SlimVisualization) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetSampleId

`func (o *SlimVisualization) GetSampleId() float64`

GetSampleId returns the SampleId field if non-nil, zero value otherwise.

### GetSampleIdOk

`func (o *SlimVisualization) GetSampleIdOk() (*float64, bool)`

GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleId

`func (o *SlimVisualization) SetSampleId(v float64)`

SetSampleId sets SampleId field to given value.

### HasSampleId

`func (o *SlimVisualization) HasSampleId() bool`

HasSampleId returns a boolean if a field has been set.

### GetLayout

`func (o *SlimVisualization) GetLayout() SizedLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *SlimVisualization) GetLayoutOk() (*SizedLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *SlimVisualization) SetLayout(v SizedLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *SlimVisualization) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetVisualizationUuid

`func (o *SlimVisualization) GetVisualizationUuid() string`

GetVisualizationUuid returns the VisualizationUuid field if non-nil, zero value otherwise.

### GetVisualizationUuidOk

`func (o *SlimVisualization) GetVisualizationUuidOk() (*string, bool)`

GetVisualizationUuidOk returns a tuple with the VisualizationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizationUuid

`func (o *SlimVisualization) SetVisualizationUuid(v string)`

SetVisualizationUuid sets VisualizationUuid field to given value.


### GetBlob

`func (o *SlimVisualization) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *SlimVisualization) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *SlimVisualization) SetBlob(v string)`

SetBlob sets Blob field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


