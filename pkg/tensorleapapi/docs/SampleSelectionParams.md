# SampleSelectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionRunId** | **string** |  | 
**SampleIdentity** | [**SampleIdentity**](SampleIdentity.md) |  | 
**FromDatasetSlice** | [**DataStateType**](DataStateType.md) |  | 
**BatchSize** | **float64** |  | 
**NumOfSamples** | **float64** |  | 
**FromEpoch** | **float64** |  | 

## Methods

### NewSampleSelectionParams

`func NewSampleSelectionParams(sessionRunId string, sampleIdentity SampleIdentity, fromDatasetSlice DataStateType, batchSize float64, numOfSamples float64, fromEpoch float64, ) *SampleSelectionParams`

NewSampleSelectionParams instantiates a new SampleSelectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleSelectionParamsWithDefaults

`func NewSampleSelectionParamsWithDefaults() *SampleSelectionParams`

NewSampleSelectionParamsWithDefaults instantiates a new SampleSelectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionRunId

`func (o *SampleSelectionParams) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *SampleSelectionParams) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *SampleSelectionParams) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetSampleIdentity

`func (o *SampleSelectionParams) GetSampleIdentity() SampleIdentity`

GetSampleIdentity returns the SampleIdentity field if non-nil, zero value otherwise.

### GetSampleIdentityOk

`func (o *SampleSelectionParams) GetSampleIdentityOk() (*SampleIdentity, bool)`

GetSampleIdentityOk returns a tuple with the SampleIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIdentity

`func (o *SampleSelectionParams) SetSampleIdentity(v SampleIdentity)`

SetSampleIdentity sets SampleIdentity field to given value.


### GetFromDatasetSlice

`func (o *SampleSelectionParams) GetFromDatasetSlice() DataStateType`

GetFromDatasetSlice returns the FromDatasetSlice field if non-nil, zero value otherwise.

### GetFromDatasetSliceOk

`func (o *SampleSelectionParams) GetFromDatasetSliceOk() (*DataStateType, bool)`

GetFromDatasetSliceOk returns a tuple with the FromDatasetSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDatasetSlice

`func (o *SampleSelectionParams) SetFromDatasetSlice(v DataStateType)`

SetFromDatasetSlice sets FromDatasetSlice field to given value.


### GetBatchSize

`func (o *SampleSelectionParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *SampleSelectionParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *SampleSelectionParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetNumOfSamples

`func (o *SampleSelectionParams) GetNumOfSamples() float64`

GetNumOfSamples returns the NumOfSamples field if non-nil, zero value otherwise.

### GetNumOfSamplesOk

`func (o *SampleSelectionParams) GetNumOfSamplesOk() (*float64, bool)`

GetNumOfSamplesOk returns a tuple with the NumOfSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfSamples

`func (o *SampleSelectionParams) SetNumOfSamples(v float64)`

SetNumOfSamples sets NumOfSamples field to given value.


### GetFromEpoch

`func (o *SampleSelectionParams) GetFromEpoch() float64`

GetFromEpoch returns the FromEpoch field if non-nil, zero value otherwise.

### GetFromEpochOk

`func (o *SampleSelectionParams) GetFromEpochOk() (*float64, bool)`

GetFromEpochOk returns a tuple with the FromEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEpoch

`func (o *SampleSelectionParams) SetFromEpoch(v float64)`

SetFromEpoch sets FromEpoch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


