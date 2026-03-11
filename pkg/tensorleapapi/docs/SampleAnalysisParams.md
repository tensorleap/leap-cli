# SampleAnalysisParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**ProjectId** | **string** |  | 
**SampleIdentity** | [**SampleIdentity**](SampleIdentity.md) |  | 
**Algo** | [**SampleAnalysisAlgo**](SampleAnalysisAlgo.md) |  | 

## Methods

### NewSampleAnalysisParams

`func NewSampleAnalysisParams(versionId string, projectId string, sampleIdentity SampleIdentity, algo SampleAnalysisAlgo, ) *SampleAnalysisParams`

NewSampleAnalysisParams instantiates a new SampleAnalysisParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleAnalysisParamsWithDefaults

`func NewSampleAnalysisParamsWithDefaults() *SampleAnalysisParams`

NewSampleAnalysisParamsWithDefaults instantiates a new SampleAnalysisParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *SampleAnalysisParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *SampleAnalysisParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *SampleAnalysisParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetProjectId

`func (o *SampleAnalysisParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SampleAnalysisParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SampleAnalysisParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSampleIdentity

`func (o *SampleAnalysisParams) GetSampleIdentity() SampleIdentity`

GetSampleIdentity returns the SampleIdentity field if non-nil, zero value otherwise.

### GetSampleIdentityOk

`func (o *SampleAnalysisParams) GetSampleIdentityOk() (*SampleIdentity, bool)`

GetSampleIdentityOk returns a tuple with the SampleIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIdentity

`func (o *SampleAnalysisParams) SetSampleIdentity(v SampleIdentity)`

SetSampleIdentity sets SampleIdentity field to given value.


### GetAlgo

`func (o *SampleAnalysisParams) GetAlgo() SampleAnalysisAlgo`

GetAlgo returns the Algo field if non-nil, zero value otherwise.

### GetAlgoOk

`func (o *SampleAnalysisParams) GetAlgoOk() (*SampleAnalysisAlgo, bool)`

GetAlgoOk returns a tuple with the Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgo

`func (o *SampleAnalysisParams) SetAlgo(v SampleAnalysisAlgo)`

SetAlgo sets Algo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


