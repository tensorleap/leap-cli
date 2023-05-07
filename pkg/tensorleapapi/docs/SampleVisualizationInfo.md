# SampleVisualizationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobType** | [**JobTypeEnum**](JobTypeEnum.md) |  | 
**AnalyzeType** | [**AnalyzeTypeEnum**](AnalyzeTypeEnum.md) |  | 
**Guid** | **string** |  | 
**SelectedSample** | [**SampleIdentity**](SampleIdentity.md) |  | 

## Methods

### NewSampleVisualizationInfo

`func NewSampleVisualizationInfo(jobType JobTypeEnum, analyzeType AnalyzeTypeEnum, guid string, selectedSample SampleIdentity, ) *SampleVisualizationInfo`

NewSampleVisualizationInfo instantiates a new SampleVisualizationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleVisualizationInfoWithDefaults

`func NewSampleVisualizationInfoWithDefaults() *SampleVisualizationInfo`

NewSampleVisualizationInfoWithDefaults instantiates a new SampleVisualizationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobType

`func (o *SampleVisualizationInfo) GetJobType() JobTypeEnum`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *SampleVisualizationInfo) GetJobTypeOk() (*JobTypeEnum, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *SampleVisualizationInfo) SetJobType(v JobTypeEnum)`

SetJobType sets JobType field to given value.


### GetAnalyzeType

`func (o *SampleVisualizationInfo) GetAnalyzeType() AnalyzeTypeEnum`

GetAnalyzeType returns the AnalyzeType field if non-nil, zero value otherwise.

### GetAnalyzeTypeOk

`func (o *SampleVisualizationInfo) GetAnalyzeTypeOk() (*AnalyzeTypeEnum, bool)`

GetAnalyzeTypeOk returns a tuple with the AnalyzeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzeType

`func (o *SampleVisualizationInfo) SetAnalyzeType(v AnalyzeTypeEnum)`

SetAnalyzeType sets AnalyzeType field to given value.


### GetGuid

`func (o *SampleVisualizationInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *SampleVisualizationInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *SampleVisualizationInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetSelectedSample

`func (o *SampleVisualizationInfo) GetSelectedSample() SampleIdentity`

GetSelectedSample returns the SelectedSample field if non-nil, zero value otherwise.

### GetSelectedSampleOk

`func (o *SampleVisualizationInfo) GetSelectedSampleOk() (*SampleIdentity, bool)`

GetSelectedSampleOk returns a tuple with the SelectedSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSample

`func (o *SampleVisualizationInfo) SetSelectedSample(v SampleIdentity)`

SetSelectedSample sets SelectedSample field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


