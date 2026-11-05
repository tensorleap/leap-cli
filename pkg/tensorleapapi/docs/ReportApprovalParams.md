# ReportApprovalParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**InsightId** | **string** |  | 
**Action** | [**ReportApprovalAction**](ReportApprovalAction.md) |  | 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewReportApprovalParams

`func NewReportApprovalParams(projectId string, insightId string, action ReportApprovalAction, ) *ReportApprovalParams`

NewReportApprovalParams instantiates a new ReportApprovalParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportApprovalParamsWithDefaults

`func NewReportApprovalParamsWithDefaults() *ReportApprovalParams`

NewReportApprovalParamsWithDefaults instantiates a new ReportApprovalParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ReportApprovalParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ReportApprovalParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ReportApprovalParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetInsightId

`func (o *ReportApprovalParams) GetInsightId() string`

GetInsightId returns the InsightId field if non-nil, zero value otherwise.

### GetInsightIdOk

`func (o *ReportApprovalParams) GetInsightIdOk() (*string, bool)`

GetInsightIdOk returns a tuple with the InsightId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightId

`func (o *ReportApprovalParams) SetInsightId(v string)`

SetInsightId sets InsightId field to given value.


### GetAction

`func (o *ReportApprovalParams) GetAction() ReportApprovalAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ReportApprovalParams) GetActionOk() (*ReportApprovalAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ReportApprovalParams) SetAction(v ReportApprovalAction)`

SetAction sets Action field to given value.


### GetTitle

`func (o *ReportApprovalParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReportApprovalParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReportApprovalParams) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReportApprovalParams) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


