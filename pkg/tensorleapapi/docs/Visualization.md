# Visualization

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
**Data** | [**VisualizationResponse**](VisualizationResponse.md) |  | 

## Methods

### NewVisualization

`func NewVisualization(id string, jobId string, sessionRunId string, type_ AnalyzeTypeEnum, createdAt time.Time, epoch float64, visualizationUuid string, blob string, data VisualizationResponse, ) *Visualization`

NewVisualization instantiates a new Visualization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualizationWithDefaults

`func NewVisualizationWithDefaults() *Visualization`

NewVisualizationWithDefaults instantiates a new Visualization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Visualization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Visualization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Visualization) SetId(v string)`

SetId sets Id field to given value.


### GetJobId

`func (o *Visualization) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *Visualization) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *Visualization) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetSessionRunId

`func (o *Visualization) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *Visualization) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *Visualization) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetJobParms

`func (o *Visualization) GetJobParms() JobParams`

GetJobParms returns the JobParms field if non-nil, zero value otherwise.

### GetJobParmsOk

`func (o *Visualization) GetJobParmsOk() (*JobParams, bool)`

GetJobParmsOk returns a tuple with the JobParms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobParms

`func (o *Visualization) SetJobParms(v JobParams)`

SetJobParms sets JobParms field to given value.

### HasJobParms

`func (o *Visualization) HasJobParms() bool`

HasJobParms returns a boolean if a field has been set.

### GetType

`func (o *Visualization) GetType() AnalyzeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Visualization) GetTypeOk() (*AnalyzeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Visualization) SetType(v AnalyzeTypeEnum)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *Visualization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Visualization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Visualization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEpoch

`func (o *Visualization) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *Visualization) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *Visualization) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetSampleId

`func (o *Visualization) GetSampleId() float64`

GetSampleId returns the SampleId field if non-nil, zero value otherwise.

### GetSampleIdOk

`func (o *Visualization) GetSampleIdOk() (*float64, bool)`

GetSampleIdOk returns a tuple with the SampleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleId

`func (o *Visualization) SetSampleId(v float64)`

SetSampleId sets SampleId field to given value.

### HasSampleId

`func (o *Visualization) HasSampleId() bool`

HasSampleId returns a boolean if a field has been set.

### GetLayout

`func (o *Visualization) GetLayout() SizedLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Visualization) GetLayoutOk() (*SizedLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Visualization) SetLayout(v SizedLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *Visualization) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetVisualizationUuid

`func (o *Visualization) GetVisualizationUuid() string`

GetVisualizationUuid returns the VisualizationUuid field if non-nil, zero value otherwise.

### GetVisualizationUuidOk

`func (o *Visualization) GetVisualizationUuidOk() (*string, bool)`

GetVisualizationUuidOk returns a tuple with the VisualizationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizationUuid

`func (o *Visualization) SetVisualizationUuid(v string)`

SetVisualizationUuid sets VisualizationUuid field to given value.


### GetBlob

`func (o *Visualization) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *Visualization) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *Visualization) SetBlob(v string)`

SetBlob sets Blob field to given value.


### GetData

`func (o *Visualization) GetData() VisualizationResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Visualization) GetDataOk() (*VisualizationResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Visualization) SetData(v VisualizationResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


