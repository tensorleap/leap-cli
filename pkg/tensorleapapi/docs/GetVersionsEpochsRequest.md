# GetVersionsEpochsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**VersionIds** | **[]string** |  | 

## Methods

### NewGetVersionsEpochsRequest

`func NewGetVersionsEpochsRequest(projectId string, versionIds []string, ) *GetVersionsEpochsRequest`

NewGetVersionsEpochsRequest instantiates a new GetVersionsEpochsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVersionsEpochsRequestWithDefaults

`func NewGetVersionsEpochsRequestWithDefaults() *GetVersionsEpochsRequest`

NewGetVersionsEpochsRequestWithDefaults instantiates a new GetVersionsEpochsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GetVersionsEpochsRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetVersionsEpochsRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetVersionsEpochsRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionIds

`func (o *GetVersionsEpochsRequest) GetVersionIds() []string`

GetVersionIds returns the VersionIds field if non-nil, zero value otherwise.

### GetVersionIdsOk

`func (o *GetVersionsEpochsRequest) GetVersionIdsOk() (*[]string, bool)`

GetVersionIdsOk returns a tuple with the VersionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionIds

`func (o *GetVersionsEpochsRequest) SetVersionIds(v []string)`

SetVersionIds sets VersionIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


