# AnalyzeGraphParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Graph** | [**ModelGraph**](ModelGraph.md) |  | 
**DatasetVersionId** | Pointer to **string** |  | [optional] 
**VersionId** | **string** |  | 
**ProjectId** | **string** |  | 
**RequestToken** | **string** |  | 

## Methods

### NewAnalyzeGraphParams

`func NewAnalyzeGraphParams(graph ModelGraph, versionId string, projectId string, requestToken string, ) *AnalyzeGraphParams`

NewAnalyzeGraphParams instantiates a new AnalyzeGraphParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyzeGraphParamsWithDefaults

`func NewAnalyzeGraphParamsWithDefaults() *AnalyzeGraphParams`

NewAnalyzeGraphParamsWithDefaults instantiates a new AnalyzeGraphParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraph

`func (o *AnalyzeGraphParams) GetGraph() ModelGraph`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *AnalyzeGraphParams) GetGraphOk() (*ModelGraph, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *AnalyzeGraphParams) SetGraph(v ModelGraph)`

SetGraph sets Graph field to given value.


### GetDatasetVersionId

`func (o *AnalyzeGraphParams) GetDatasetVersionId() string`

GetDatasetVersionId returns the DatasetVersionId field if non-nil, zero value otherwise.

### GetDatasetVersionIdOk

`func (o *AnalyzeGraphParams) GetDatasetVersionIdOk() (*string, bool)`

GetDatasetVersionIdOk returns a tuple with the DatasetVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetVersionId

`func (o *AnalyzeGraphParams) SetDatasetVersionId(v string)`

SetDatasetVersionId sets DatasetVersionId field to given value.

### HasDatasetVersionId

`func (o *AnalyzeGraphParams) HasDatasetVersionId() bool`

HasDatasetVersionId returns a boolean if a field has been set.

### GetVersionId

`func (o *AnalyzeGraphParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *AnalyzeGraphParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *AnalyzeGraphParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetProjectId

`func (o *AnalyzeGraphParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AnalyzeGraphParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AnalyzeGraphParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetRequestToken

`func (o *AnalyzeGraphParams) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *AnalyzeGraphParams) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *AnalyzeGraphParams) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


