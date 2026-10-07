# GetAnalysisJobLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractVersion** | **float64** |  | 
**PodsLogs** | [**[]AnalysisPodLogs**](AnalysisPodLogs.md) |  | 

## Methods

### NewGetAnalysisJobLogsResponse

`func NewGetAnalysisJobLogsResponse(contractVersion float64, podsLogs []AnalysisPodLogs, ) *GetAnalysisJobLogsResponse`

NewGetAnalysisJobLogsResponse instantiates a new GetAnalysisJobLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalysisJobLogsResponseWithDefaults

`func NewGetAnalysisJobLogsResponseWithDefaults() *GetAnalysisJobLogsResponse`

NewGetAnalysisJobLogsResponseWithDefaults instantiates a new GetAnalysisJobLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractVersion

`func (o *GetAnalysisJobLogsResponse) GetContractVersion() float64`

GetContractVersion returns the ContractVersion field if non-nil, zero value otherwise.

### GetContractVersionOk

`func (o *GetAnalysisJobLogsResponse) GetContractVersionOk() (*float64, bool)`

GetContractVersionOk returns a tuple with the ContractVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractVersion

`func (o *GetAnalysisJobLogsResponse) SetContractVersion(v float64)`

SetContractVersion sets ContractVersion field to given value.


### GetPodsLogs

`func (o *GetAnalysisJobLogsResponse) GetPodsLogs() []AnalysisPodLogs`

GetPodsLogs returns the PodsLogs field if non-nil, zero value otherwise.

### GetPodsLogsOk

`func (o *GetAnalysisJobLogsResponse) GetPodsLogsOk() (*[]AnalysisPodLogs, bool)`

GetPodsLogsOk returns a tuple with the PodsLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodsLogs

`func (o *GetAnalysisJobLogsResponse) SetPodsLogs(v []AnalysisPodLogs)`

SetPodsLogs sets PodsLogs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


