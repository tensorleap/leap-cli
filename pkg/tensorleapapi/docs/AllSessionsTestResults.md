# AllSessionsTestResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestId** | **string** |  | 
**SessionsResults** | [**[]SessionTestResult**](SessionTestResult.md) |  | 

## Methods

### NewAllSessionsTestResults

`func NewAllSessionsTestResults(testId string, sessionsResults []SessionTestResult, ) *AllSessionsTestResults`

NewAllSessionsTestResults instantiates a new AllSessionsTestResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllSessionsTestResultsWithDefaults

`func NewAllSessionsTestResultsWithDefaults() *AllSessionsTestResults`

NewAllSessionsTestResultsWithDefaults instantiates a new AllSessionsTestResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestId

`func (o *AllSessionsTestResults) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *AllSessionsTestResults) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *AllSessionsTestResults) SetTestId(v string)`

SetTestId sets TestId field to given value.


### GetSessionsResults

`func (o *AllSessionsTestResults) GetSessionsResults() []SessionTestResult`

GetSessionsResults returns the SessionsResults field if non-nil, zero value otherwise.

### GetSessionsResultsOk

`func (o *AllSessionsTestResults) GetSessionsResultsOk() (*[]SessionTestResult, bool)`

GetSessionsResultsOk returns a tuple with the SessionsResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsResults

`func (o *AllSessionsTestResults) SetSessionsResults(v []SessionTestResult)`

SetSessionsResults sets SessionsResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


