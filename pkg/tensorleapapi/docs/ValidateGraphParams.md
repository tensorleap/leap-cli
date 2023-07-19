# ValidateGraphParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Graph** | [**ModelGraph**](ModelGraph.md) |  | 
**DatasetVersionId** | **string** |  | 
**VersionId** | **string** |  | 
**ProjectId** | **string** |  | 
**Digest** | **string** |  | 

## Methods

### NewValidateGraphParams

`func NewValidateGraphParams(graph ModelGraph, datasetVersionId string, versionId string, projectId string, digest string, ) *ValidateGraphParams`

NewValidateGraphParams instantiates a new ValidateGraphParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateGraphParamsWithDefaults

`func NewValidateGraphParamsWithDefaults() *ValidateGraphParams`

NewValidateGraphParamsWithDefaults instantiates a new ValidateGraphParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraph

`func (o *ValidateGraphParams) GetGraph() ModelGraph`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *ValidateGraphParams) GetGraphOk() (*ModelGraph, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *ValidateGraphParams) SetGraph(v ModelGraph)`

SetGraph sets Graph field to given value.


### GetDatasetVersionId

`func (o *ValidateGraphParams) GetDatasetVersionId() string`

GetDatasetVersionId returns the DatasetVersionId field if non-nil, zero value otherwise.

### GetDatasetVersionIdOk

`func (o *ValidateGraphParams) GetDatasetVersionIdOk() (*string, bool)`

GetDatasetVersionIdOk returns a tuple with the DatasetVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetVersionId

`func (o *ValidateGraphParams) SetDatasetVersionId(v string)`

SetDatasetVersionId sets DatasetVersionId field to given value.


### GetVersionId

`func (o *ValidateGraphParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ValidateGraphParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ValidateGraphParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetProjectId

`func (o *ValidateGraphParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ValidateGraphParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ValidateGraphParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetDigest

`func (o *ValidateGraphParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ValidateGraphParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ValidateGraphParams) SetDigest(v string)`

SetDigest sets Digest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


