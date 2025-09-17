# JobMessageParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureInfo** | Pointer to [**FailureInfo**](FailureInfo.md) |  | [optional] 
**Module** | **string** |  | 
**JobStatus** | [**JobStatus**](JobStatus.md) |  | 

## Methods

### NewJobMessageParams

`func NewJobMessageParams(module string, jobStatus JobStatus, ) *JobMessageParams`

NewJobMessageParams instantiates a new JobMessageParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobMessageParamsWithDefaults

`func NewJobMessageParamsWithDefaults() *JobMessageParams`

NewJobMessageParamsWithDefaults instantiates a new JobMessageParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureInfo

`func (o *JobMessageParams) GetFailureInfo() FailureInfo`

GetFailureInfo returns the FailureInfo field if non-nil, zero value otherwise.

### GetFailureInfoOk

`func (o *JobMessageParams) GetFailureInfoOk() (*FailureInfo, bool)`

GetFailureInfoOk returns a tuple with the FailureInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureInfo

`func (o *JobMessageParams) SetFailureInfo(v FailureInfo)`

SetFailureInfo sets FailureInfo field to given value.

### HasFailureInfo

`func (o *JobMessageParams) HasFailureInfo() bool`

HasFailureInfo returns a boolean if a field has been set.

### GetModule

`func (o *JobMessageParams) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *JobMessageParams) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *JobMessageParams) SetModule(v string)`

SetModule sets Module field to given value.


### GetJobStatus

`func (o *JobMessageParams) GetJobStatus() JobStatus`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *JobMessageParams) GetJobStatusOk() (*JobStatus, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *JobMessageParams) SetJobStatus(v JobStatus)`

SetJobStatus sets JobStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


