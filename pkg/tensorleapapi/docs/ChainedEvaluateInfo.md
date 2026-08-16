# ChainedEvaluateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**JobId** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewChainedEvaluateInfo

`func NewChainedEvaluateInfo(status string, ) *ChainedEvaluateInfo`

NewChainedEvaluateInfo instantiates a new ChainedEvaluateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainedEvaluateInfoWithDefaults

`func NewChainedEvaluateInfoWithDefaults() *ChainedEvaluateInfo`

NewChainedEvaluateInfoWithDefaults instantiates a new ChainedEvaluateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChainedEvaluateInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChainedEvaluateInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChainedEvaluateInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetJobId

`func (o *ChainedEvaluateInfo) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ChainedEvaluateInfo) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ChainedEvaluateInfo) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ChainedEvaluateInfo) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetError

`func (o *ChainedEvaluateInfo) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ChainedEvaluateInfo) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ChainedEvaluateInfo) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ChainedEvaluateInfo) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


