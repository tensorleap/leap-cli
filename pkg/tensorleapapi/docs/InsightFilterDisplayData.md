# InsightFilterDisplayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**InsightData** | [**InsightType**](InsightType.md) |  | 
**SessionRun** | [**FilterSessionRun**](FilterSessionRun.md) |  | 

## Methods

### NewInsightFilterDisplayData

`func NewInsightFilterDisplayData(type_ string, insightData InsightType, sessionRun FilterSessionRun, ) *InsightFilterDisplayData`

NewInsightFilterDisplayData instantiates a new InsightFilterDisplayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightFilterDisplayDataWithDefaults

`func NewInsightFilterDisplayDataWithDefaults() *InsightFilterDisplayData`

NewInsightFilterDisplayDataWithDefaults instantiates a new InsightFilterDisplayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InsightFilterDisplayData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InsightFilterDisplayData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InsightFilterDisplayData) SetType(v string)`

SetType sets Type field to given value.


### GetInsightData

`func (o *InsightFilterDisplayData) GetInsightData() InsightType`

GetInsightData returns the InsightData field if non-nil, zero value otherwise.

### GetInsightDataOk

`func (o *InsightFilterDisplayData) GetInsightDataOk() (*InsightType, bool)`

GetInsightDataOk returns a tuple with the InsightData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightData

`func (o *InsightFilterDisplayData) SetInsightData(v InsightType)`

SetInsightData sets InsightData field to given value.


### GetSessionRun

`func (o *InsightFilterDisplayData) GetSessionRun() FilterSessionRun`

GetSessionRun returns the SessionRun field if non-nil, zero value otherwise.

### GetSessionRunOk

`func (o *InsightFilterDisplayData) GetSessionRunOk() (*FilterSessionRun, bool)`

GetSessionRunOk returns a tuple with the SessionRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRun

`func (o *InsightFilterDisplayData) SetSessionRun(v FilterSessionRun)`

SetSessionRun sets SessionRun field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


