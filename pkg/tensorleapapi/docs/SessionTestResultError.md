# SessionTestResultError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionRunId** | **string** |  | 
**QueryStatus** | **string** |  | 

## Methods

### NewSessionTestResultError

`func NewSessionTestResultError(sessionRunId string, queryStatus string, ) *SessionTestResultError`

NewSessionTestResultError instantiates a new SessionTestResultError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionTestResultErrorWithDefaults

`func NewSessionTestResultErrorWithDefaults() *SessionTestResultError`

NewSessionTestResultErrorWithDefaults instantiates a new SessionTestResultError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionRunId

`func (o *SessionTestResultError) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *SessionTestResultError) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *SessionTestResultError) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetQueryStatus

`func (o *SessionTestResultError) GetQueryStatus() string`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *SessionTestResultError) GetQueryStatusOk() (*string, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *SessionTestResultError) SetQueryStatus(v string)`

SetQueryStatus sets QueryStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


