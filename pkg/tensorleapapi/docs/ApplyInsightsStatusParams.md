# ApplyInsightsStatusParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InsightId** | **string** |  | 
**Status** | [**InsightStatus**](InsightStatus.md) |  | 
**ProjectId** | **string** |  | 

## Methods

### NewApplyInsightsStatusParams

`func NewApplyInsightsStatusParams(insightId string, status InsightStatus, projectId string, ) *ApplyInsightsStatusParams`

NewApplyInsightsStatusParams instantiates a new ApplyInsightsStatusParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyInsightsStatusParamsWithDefaults

`func NewApplyInsightsStatusParamsWithDefaults() *ApplyInsightsStatusParams`

NewApplyInsightsStatusParamsWithDefaults instantiates a new ApplyInsightsStatusParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsightId

`func (o *ApplyInsightsStatusParams) GetInsightId() string`

GetInsightId returns the InsightId field if non-nil, zero value otherwise.

### GetInsightIdOk

`func (o *ApplyInsightsStatusParams) GetInsightIdOk() (*string, bool)`

GetInsightIdOk returns a tuple with the InsightId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightId

`func (o *ApplyInsightsStatusParams) SetInsightId(v string)`

SetInsightId sets InsightId field to given value.


### GetStatus

`func (o *ApplyInsightsStatusParams) GetStatus() InsightStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplyInsightsStatusParams) GetStatusOk() (*InsightStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplyInsightsStatusParams) SetStatus(v InsightStatus)`

SetStatus sets Status field to given value.


### GetProjectId

`func (o *ApplyInsightsStatusParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ApplyInsightsStatusParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ApplyInsightsStatusParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


