# CodeSnapshotSetupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneralError** | Pointer to **string** |  | [optional] 
**PrintLog** | Pointer to **string** |  | [optional] 
**BindersStatus** | [**[]DatasetTestResultPayload**](DatasetTestResultPayload.md) |  | 

## Methods

### NewCodeSnapshotSetupStatus

`func NewCodeSnapshotSetupStatus(bindersStatus []DatasetTestResultPayload, ) *CodeSnapshotSetupStatus`

NewCodeSnapshotSetupStatus instantiates a new CodeSnapshotSetupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeSnapshotSetupStatusWithDefaults

`func NewCodeSnapshotSetupStatusWithDefaults() *CodeSnapshotSetupStatus`

NewCodeSnapshotSetupStatusWithDefaults instantiates a new CodeSnapshotSetupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneralError

`func (o *CodeSnapshotSetupStatus) GetGeneralError() string`

GetGeneralError returns the GeneralError field if non-nil, zero value otherwise.

### GetGeneralErrorOk

`func (o *CodeSnapshotSetupStatus) GetGeneralErrorOk() (*string, bool)`

GetGeneralErrorOk returns a tuple with the GeneralError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralError

`func (o *CodeSnapshotSetupStatus) SetGeneralError(v string)`

SetGeneralError sets GeneralError field to given value.

### HasGeneralError

`func (o *CodeSnapshotSetupStatus) HasGeneralError() bool`

HasGeneralError returns a boolean if a field has been set.

### GetPrintLog

`func (o *CodeSnapshotSetupStatus) GetPrintLog() string`

GetPrintLog returns the PrintLog field if non-nil, zero value otherwise.

### GetPrintLogOk

`func (o *CodeSnapshotSetupStatus) GetPrintLogOk() (*string, bool)`

GetPrintLogOk returns a tuple with the PrintLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintLog

`func (o *CodeSnapshotSetupStatus) SetPrintLog(v string)`

SetPrintLog sets PrintLog field to given value.

### HasPrintLog

`func (o *CodeSnapshotSetupStatus) HasPrintLog() bool`

HasPrintLog returns a boolean if a field has been set.

### GetBindersStatus

`func (o *CodeSnapshotSetupStatus) GetBindersStatus() []DatasetTestResultPayload`

GetBindersStatus returns the BindersStatus field if non-nil, zero value otherwise.

### GetBindersStatusOk

`func (o *CodeSnapshotSetupStatus) GetBindersStatusOk() (*[]DatasetTestResultPayload, bool)`

GetBindersStatusOk returns a tuple with the BindersStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindersStatus

`func (o *CodeSnapshotSetupStatus) SetBindersStatus(v []DatasetTestResultPayload)`

SetBindersStatus sets BindersStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


