# VizInfoType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobType** | [**JobTypeEnum**](JobTypeEnum.md) |  | 
**AnalyzeType** | [**AnalyzeTypeEnum**](AnalyzeTypeEnum.md) |  | 
**Guid** | **string** |  | 
**SelectedSample** | [**SampleIdentity**](SampleIdentity.md) |  | 
**PopulationExplorationNSamples** | **float64** |  | 
**NSelectedSamples** | **float64** |  | 

## Methods

### NewVizInfoType

`func NewVizInfoType(jobType JobTypeEnum, analyzeType AnalyzeTypeEnum, guid string, selectedSample SampleIdentity, populationExplorationNSamples float64, nSelectedSamples float64, ) *VizInfoType`

NewVizInfoType instantiates a new VizInfoType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVizInfoTypeWithDefaults

`func NewVizInfoTypeWithDefaults() *VizInfoType`

NewVizInfoTypeWithDefaults instantiates a new VizInfoType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobType

`func (o *VizInfoType) GetJobType() JobTypeEnum`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *VizInfoType) GetJobTypeOk() (*JobTypeEnum, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *VizInfoType) SetJobType(v JobTypeEnum)`

SetJobType sets JobType field to given value.


### GetAnalyzeType

`func (o *VizInfoType) GetAnalyzeType() AnalyzeTypeEnum`

GetAnalyzeType returns the AnalyzeType field if non-nil, zero value otherwise.

### GetAnalyzeTypeOk

`func (o *VizInfoType) GetAnalyzeTypeOk() (*AnalyzeTypeEnum, bool)`

GetAnalyzeTypeOk returns a tuple with the AnalyzeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzeType

`func (o *VizInfoType) SetAnalyzeType(v AnalyzeTypeEnum)`

SetAnalyzeType sets AnalyzeType field to given value.


### GetGuid

`func (o *VizInfoType) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *VizInfoType) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *VizInfoType) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetSelectedSample

`func (o *VizInfoType) GetSelectedSample() SampleIdentity`

GetSelectedSample returns the SelectedSample field if non-nil, zero value otherwise.

### GetSelectedSampleOk

`func (o *VizInfoType) GetSelectedSampleOk() (*SampleIdentity, bool)`

GetSelectedSampleOk returns a tuple with the SelectedSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSample

`func (o *VizInfoType) SetSelectedSample(v SampleIdentity)`

SetSelectedSample sets SelectedSample field to given value.


### GetPopulationExplorationNSamples

`func (o *VizInfoType) GetPopulationExplorationNSamples() float64`

GetPopulationExplorationNSamples returns the PopulationExplorationNSamples field if non-nil, zero value otherwise.

### GetPopulationExplorationNSamplesOk

`func (o *VizInfoType) GetPopulationExplorationNSamplesOk() (*float64, bool)`

GetPopulationExplorationNSamplesOk returns a tuple with the PopulationExplorationNSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationExplorationNSamples

`func (o *VizInfoType) SetPopulationExplorationNSamples(v float64)`

SetPopulationExplorationNSamples sets PopulationExplorationNSamples field to given value.


### GetNSelectedSamples

`func (o *VizInfoType) GetNSelectedSamples() float64`

GetNSelectedSamples returns the NSelectedSamples field if non-nil, zero value otherwise.

### GetNSelectedSamplesOk

`func (o *VizInfoType) GetNSelectedSamplesOk() (*float64, bool)`

GetNSelectedSamplesOk returns a tuple with the NSelectedSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSelectedSamples

`func (o *VizInfoType) SetNSelectedSamples(v float64)`

SetNSelectedSamples sets NSelectedSamples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


