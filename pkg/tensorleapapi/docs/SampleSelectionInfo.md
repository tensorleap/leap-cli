# SampleSelectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobType** | [**JobTypeEnum**](JobTypeEnum.md) |  | 
**AnalyzeType** | [**AnalyzeTypeEnum**](AnalyzeTypeEnum.md) |  | 
**Guid** | **string** |  | 
**NSelectedSamples** | **float64** |  | 
**SelectedSample** | [**SampleIdentity**](SampleIdentity.md) |  | 

## Methods

### NewSampleSelectionInfo

`func NewSampleSelectionInfo(jobType JobTypeEnum, analyzeType AnalyzeTypeEnum, guid string, nSelectedSamples float64, selectedSample SampleIdentity, ) *SampleSelectionInfo`

NewSampleSelectionInfo instantiates a new SampleSelectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleSelectionInfoWithDefaults

`func NewSampleSelectionInfoWithDefaults() *SampleSelectionInfo`

NewSampleSelectionInfoWithDefaults instantiates a new SampleSelectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobType

`func (o *SampleSelectionInfo) GetJobType() JobTypeEnum`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *SampleSelectionInfo) GetJobTypeOk() (*JobTypeEnum, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *SampleSelectionInfo) SetJobType(v JobTypeEnum)`

SetJobType sets JobType field to given value.


### GetAnalyzeType

`func (o *SampleSelectionInfo) GetAnalyzeType() AnalyzeTypeEnum`

GetAnalyzeType returns the AnalyzeType field if non-nil, zero value otherwise.

### GetAnalyzeTypeOk

`func (o *SampleSelectionInfo) GetAnalyzeTypeOk() (*AnalyzeTypeEnum, bool)`

GetAnalyzeTypeOk returns a tuple with the AnalyzeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzeType

`func (o *SampleSelectionInfo) SetAnalyzeType(v AnalyzeTypeEnum)`

SetAnalyzeType sets AnalyzeType field to given value.


### GetGuid

`func (o *SampleSelectionInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *SampleSelectionInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *SampleSelectionInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetNSelectedSamples

`func (o *SampleSelectionInfo) GetNSelectedSamples() float64`

GetNSelectedSamples returns the NSelectedSamples field if non-nil, zero value otherwise.

### GetNSelectedSamplesOk

`func (o *SampleSelectionInfo) GetNSelectedSamplesOk() (*float64, bool)`

GetNSelectedSamplesOk returns a tuple with the NSelectedSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSelectedSamples

`func (o *SampleSelectionInfo) SetNSelectedSamples(v float64)`

SetNSelectedSamples sets NSelectedSamples field to given value.


### GetSelectedSample

`func (o *SampleSelectionInfo) GetSelectedSample() SampleIdentity`

GetSelectedSample returns the SelectedSample field if non-nil, zero value otherwise.

### GetSelectedSampleOk

`func (o *SampleSelectionInfo) GetSelectedSampleOk() (*SampleIdentity, bool)`

GetSelectedSampleOk returns a tuple with the SelectedSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSample

`func (o *SampleSelectionInfo) SetSelectedSample(v SampleIdentity)`

SetSelectedSample sets SelectedSample field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


