# ExportProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**CopyToUrl** | Pointer to **string** |  | [optional] 
**NoCache** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to [**ExportOptions**](ExportOptions.md) |  | [optional] 

## Methods

### NewExportProjectRequest

`func NewExportProjectRequest(projectId string, ) *ExportProjectRequest`

NewExportProjectRequest instantiates a new ExportProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportProjectRequestWithDefaults

`func NewExportProjectRequestWithDefaults() *ExportProjectRequest`

NewExportProjectRequestWithDefaults instantiates a new ExportProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ExportProjectRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ExportProjectRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ExportProjectRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCopyToUrl

`func (o *ExportProjectRequest) GetCopyToUrl() string`

GetCopyToUrl returns the CopyToUrl field if non-nil, zero value otherwise.

### GetCopyToUrlOk

`func (o *ExportProjectRequest) GetCopyToUrlOk() (*string, bool)`

GetCopyToUrlOk returns a tuple with the CopyToUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyToUrl

`func (o *ExportProjectRequest) SetCopyToUrl(v string)`

SetCopyToUrl sets CopyToUrl field to given value.

### HasCopyToUrl

`func (o *ExportProjectRequest) HasCopyToUrl() bool`

HasCopyToUrl returns a boolean if a field has been set.

### GetNoCache

`func (o *ExportProjectRequest) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *ExportProjectRequest) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *ExportProjectRequest) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *ExportProjectRequest) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.

### GetOptions

`func (o *ExportProjectRequest) GetOptions() ExportOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ExportProjectRequest) GetOptionsOk() (*ExportOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ExportProjectRequest) SetOptions(v ExportOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ExportProjectRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


