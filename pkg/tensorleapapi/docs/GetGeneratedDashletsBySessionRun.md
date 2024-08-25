# GetGeneratedDashletsBySessionRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashlets** | [**[]GeneratedDashlet**](GeneratedDashlet.md) |  | 
**SessionRunId** | **string** |  | 
**Status** | [**JobStatus**](JobStatus.md) |  | 

## Methods

### NewGetGeneratedDashletsBySessionRun

`func NewGetGeneratedDashletsBySessionRun(dashlets []GeneratedDashlet, sessionRunId string, status JobStatus, ) *GetGeneratedDashletsBySessionRun`

NewGetGeneratedDashletsBySessionRun instantiates a new GetGeneratedDashletsBySessionRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGeneratedDashletsBySessionRunWithDefaults

`func NewGetGeneratedDashletsBySessionRunWithDefaults() *GetGeneratedDashletsBySessionRun`

NewGetGeneratedDashletsBySessionRunWithDefaults instantiates a new GetGeneratedDashletsBySessionRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashlets

`func (o *GetGeneratedDashletsBySessionRun) GetDashlets() []GeneratedDashlet`

GetDashlets returns the Dashlets field if non-nil, zero value otherwise.

### GetDashletsOk

`func (o *GetGeneratedDashletsBySessionRun) GetDashletsOk() (*[]GeneratedDashlet, bool)`

GetDashletsOk returns a tuple with the Dashlets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashlets

`func (o *GetGeneratedDashletsBySessionRun) SetDashlets(v []GeneratedDashlet)`

SetDashlets sets Dashlets field to given value.


### GetSessionRunId

`func (o *GetGeneratedDashletsBySessionRun) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *GetGeneratedDashletsBySessionRun) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *GetGeneratedDashletsBySessionRun) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetStatus

`func (o *GetGeneratedDashletsBySessionRun) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGeneratedDashletsBySessionRun) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGeneratedDashletsBySessionRun) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


