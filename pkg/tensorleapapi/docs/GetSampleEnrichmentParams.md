# GetSampleEnrichmentParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | [**[]SampleIdentity**](SampleIdentity.md) | Visible / batch samples to fetch data for. | 
**Epoch** | **float64** |  | 
**SessionRunIds** | **[]string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewGetSampleEnrichmentParams

`func NewGetSampleEnrichmentParams(samples []SampleIdentity, epoch float64, sessionRunIds []string, projectId string, ) *GetSampleEnrichmentParams`

NewGetSampleEnrichmentParams instantiates a new GetSampleEnrichmentParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSampleEnrichmentParamsWithDefaults

`func NewGetSampleEnrichmentParamsWithDefaults() *GetSampleEnrichmentParams`

NewGetSampleEnrichmentParamsWithDefaults instantiates a new GetSampleEnrichmentParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *GetSampleEnrichmentParams) GetSamples() []SampleIdentity`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *GetSampleEnrichmentParams) GetSamplesOk() (*[]SampleIdentity, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *GetSampleEnrichmentParams) SetSamples(v []SampleIdentity)`

SetSamples sets Samples field to given value.


### GetEpoch

`func (o *GetSampleEnrichmentParams) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *GetSampleEnrichmentParams) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *GetSampleEnrichmentParams) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetSessionRunIds

`func (o *GetSampleEnrichmentParams) GetSessionRunIds() []string`

GetSessionRunIds returns the SessionRunIds field if non-nil, zero value otherwise.

### GetSessionRunIdsOk

`func (o *GetSampleEnrichmentParams) GetSessionRunIdsOk() (*[]string, bool)`

GetSessionRunIdsOk returns a tuple with the SessionRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunIds

`func (o *GetSampleEnrichmentParams) SetSessionRunIds(v []string)`

SetSessionRunIds sets SessionRunIds field to given value.


### GetProjectId

`func (o *GetSampleEnrichmentParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetSampleEnrichmentParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetSampleEnrichmentParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


