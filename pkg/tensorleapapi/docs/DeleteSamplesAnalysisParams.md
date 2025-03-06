# DeleteSamplesAnalysisParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SessionRunIds** | **[]string** |  | 
**SamplesIdentities** | [**[]SampleIdentity**](SampleIdentity.md) |  | 

## Methods

### NewDeleteSamplesAnalysisParams

`func NewDeleteSamplesAnalysisParams(projectId string, sessionRunIds []string, samplesIdentities []SampleIdentity, ) *DeleteSamplesAnalysisParams`

NewDeleteSamplesAnalysisParams instantiates a new DeleteSamplesAnalysisParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSamplesAnalysisParamsWithDefaults

`func NewDeleteSamplesAnalysisParamsWithDefaults() *DeleteSamplesAnalysisParams`

NewDeleteSamplesAnalysisParamsWithDefaults instantiates a new DeleteSamplesAnalysisParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DeleteSamplesAnalysisParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeleteSamplesAnalysisParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeleteSamplesAnalysisParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSessionRunIds

`func (o *DeleteSamplesAnalysisParams) GetSessionRunIds() []string`

GetSessionRunIds returns the SessionRunIds field if non-nil, zero value otherwise.

### GetSessionRunIdsOk

`func (o *DeleteSamplesAnalysisParams) GetSessionRunIdsOk() (*[]string, bool)`

GetSessionRunIdsOk returns a tuple with the SessionRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunIds

`func (o *DeleteSamplesAnalysisParams) SetSessionRunIds(v []string)`

SetSessionRunIds sets SessionRunIds field to given value.


### GetSamplesIdentities

`func (o *DeleteSamplesAnalysisParams) GetSamplesIdentities() []SampleIdentity`

GetSamplesIdentities returns the SamplesIdentities field if non-nil, zero value otherwise.

### GetSamplesIdentitiesOk

`func (o *DeleteSamplesAnalysisParams) GetSamplesIdentitiesOk() (*[]SampleIdentity, bool)`

GetSamplesIdentitiesOk returns a tuple with the SamplesIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplesIdentities

`func (o *DeleteSamplesAnalysisParams) SetSamplesIdentities(v []SampleIdentity)`

SetSamplesIdentities sets SamplesIdentities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


