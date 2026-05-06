# GetSampleEnrichmentParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | [**[]SampleIdentity**](SampleIdentity.md) | Visible / batch samples to fetch data for. | 
**VersionIds** | **[]string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewGetSampleEnrichmentParams

`func NewGetSampleEnrichmentParams(samples []SampleIdentity, versionIds []string, projectId string, ) *GetSampleEnrichmentParams`

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


### GetVersionIds

`func (o *GetSampleEnrichmentParams) GetVersionIds() []string`

GetVersionIds returns the VersionIds field if non-nil, zero value otherwise.

### GetVersionIdsOk

`func (o *GetSampleEnrichmentParams) GetVersionIdsOk() (*[]string, bool)`

GetVersionIdsOk returns a tuple with the VersionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionIds

`func (o *GetSampleEnrichmentParams) SetVersionIds(v []string)`

SetVersionIds sets VersionIds field to given value.


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


