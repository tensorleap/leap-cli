# ReportApprovalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**InsightStatus**](InsightStatus.md) |  | 
**Applied** | **bool** |  | 

## Methods

### NewReportApprovalResponse

`func NewReportApprovalResponse(status InsightStatus, applied bool, ) *ReportApprovalResponse`

NewReportApprovalResponse instantiates a new ReportApprovalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportApprovalResponseWithDefaults

`func NewReportApprovalResponseWithDefaults() *ReportApprovalResponse`

NewReportApprovalResponseWithDefaults instantiates a new ReportApprovalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ReportApprovalResponse) GetStatus() InsightStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportApprovalResponse) GetStatusOk() (*InsightStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportApprovalResponse) SetStatus(v InsightStatus)`

SetStatus sets Status field to given value.


### GetApplied

`func (o *ReportApprovalResponse) GetApplied() bool`

GetApplied returns the Applied field if non-nil, zero value otherwise.

### GetAppliedOk

`func (o *ReportApprovalResponse) GetAppliedOk() (*bool, bool)`

GetAppliedOk returns a tuple with the Applied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplied

`func (o *ReportApprovalResponse) SetApplied(v bool)`

SetApplied sets Applied field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


