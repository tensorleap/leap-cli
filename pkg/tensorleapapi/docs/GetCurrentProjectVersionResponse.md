# GetCurrentProjectVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**VersionId** | Pointer to **string** |  | [optional] 

## Methods

### NewGetCurrentProjectVersionResponse

`func NewGetCurrentProjectVersionResponse(projectId string, ) *GetCurrentProjectVersionResponse`

NewGetCurrentProjectVersionResponse instantiates a new GetCurrentProjectVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrentProjectVersionResponseWithDefaults

`func NewGetCurrentProjectVersionResponseWithDefaults() *GetCurrentProjectVersionResponse`

NewGetCurrentProjectVersionResponseWithDefaults instantiates a new GetCurrentProjectVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GetCurrentProjectVersionResponse) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetCurrentProjectVersionResponse) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetCurrentProjectVersionResponse) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionId

`func (o *GetCurrentProjectVersionResponse) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *GetCurrentProjectVersionResponse) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *GetCurrentProjectVersionResponse) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *GetCurrentProjectVersionResponse) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


