# CodeSnapshotResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelSetup** | Pointer to [**ModelSetup**](ModelSetup.md) |  | [optional] 
**SetupStatus** | Pointer to [**CodeSnapshotSetupStatus**](CodeSnapshotSetupStatus.md) |  | [optional] 
**Setup** | Pointer to [**DatasetSetup**](DatasetSetup.md) |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**TestStatus** | [**TestStatus**](TestStatus.md) |  | 

## Methods

### NewCodeSnapshotResult

`func NewCodeSnapshotResult(testStatus TestStatus, ) *CodeSnapshotResult`

NewCodeSnapshotResult instantiates a new CodeSnapshotResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeSnapshotResultWithDefaults

`func NewCodeSnapshotResultWithDefaults() *CodeSnapshotResult`

NewCodeSnapshotResultWithDefaults instantiates a new CodeSnapshotResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelSetup

`func (o *CodeSnapshotResult) GetModelSetup() ModelSetup`

GetModelSetup returns the ModelSetup field if non-nil, zero value otherwise.

### GetModelSetupOk

`func (o *CodeSnapshotResult) GetModelSetupOk() (*ModelSetup, bool)`

GetModelSetupOk returns a tuple with the ModelSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelSetup

`func (o *CodeSnapshotResult) SetModelSetup(v ModelSetup)`

SetModelSetup sets ModelSetup field to given value.

### HasModelSetup

`func (o *CodeSnapshotResult) HasModelSetup() bool`

HasModelSetup returns a boolean if a field has been set.

### GetSetupStatus

`func (o *CodeSnapshotResult) GetSetupStatus() CodeSnapshotSetupStatus`

GetSetupStatus returns the SetupStatus field if non-nil, zero value otherwise.

### GetSetupStatusOk

`func (o *CodeSnapshotResult) GetSetupStatusOk() (*CodeSnapshotSetupStatus, bool)`

GetSetupStatusOk returns a tuple with the SetupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupStatus

`func (o *CodeSnapshotResult) SetSetupStatus(v CodeSnapshotSetupStatus)`

SetSetupStatus sets SetupStatus field to given value.

### HasSetupStatus

`func (o *CodeSnapshotResult) HasSetupStatus() bool`

HasSetupStatus returns a boolean if a field has been set.

### GetSetup

`func (o *CodeSnapshotResult) GetSetup() DatasetSetup`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *CodeSnapshotResult) GetSetupOk() (*DatasetSetup, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *CodeSnapshotResult) SetSetup(v DatasetSetup)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *CodeSnapshotResult) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetJobId

`func (o *CodeSnapshotResult) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *CodeSnapshotResult) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *CodeSnapshotResult) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *CodeSnapshotResult) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetTestStatus

`func (o *CodeSnapshotResult) GetTestStatus() TestStatus`

GetTestStatus returns the TestStatus field if non-nil, zero value otherwise.

### GetTestStatusOk

`func (o *CodeSnapshotResult) GetTestStatusOk() (*TestStatus, bool)`

GetTestStatusOk returns a tuple with the TestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestStatus

`func (o *CodeSnapshotResult) SetTestStatus(v TestStatus)`

SetTestStatus sets TestStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


