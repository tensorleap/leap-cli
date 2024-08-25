# TagModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**ExperimentId** | **string** |  | 
**Tags** | **[]string** |  | 
**Epoch** | **float64** |  | 

## Methods

### NewTagModelRequest

`func NewTagModelRequest(projectId string, experimentId string, tags []string, epoch float64, ) *TagModelRequest`

NewTagModelRequest instantiates a new TagModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagModelRequestWithDefaults

`func NewTagModelRequestWithDefaults() *TagModelRequest`

NewTagModelRequestWithDefaults instantiates a new TagModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *TagModelRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TagModelRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TagModelRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetExperimentId

`func (o *TagModelRequest) GetExperimentId() string`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *TagModelRequest) GetExperimentIdOk() (*string, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *TagModelRequest) SetExperimentId(v string)`

SetExperimentId sets ExperimentId field to given value.


### GetTags

`func (o *TagModelRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagModelRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagModelRequest) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetEpoch

`func (o *TagModelRequest) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *TagModelRequest) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *TagModelRequest) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


