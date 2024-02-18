# FilterDisplayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**InsightData** | [**InsightType**](InsightType.md) |  | 
**SessionRun** | [**InsightFilterDisplayDataSessionRun**](InsightFilterDisplayDataSessionRun.md) |  | 

## Methods

### NewFilterDisplayData

`func NewFilterDisplayData(type_ string, insightData InsightType, sessionRun InsightFilterDisplayDataSessionRun, ) *FilterDisplayData`

NewFilterDisplayData instantiates a new FilterDisplayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterDisplayDataWithDefaults

`func NewFilterDisplayDataWithDefaults() *FilterDisplayData`

NewFilterDisplayDataWithDefaults instantiates a new FilterDisplayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FilterDisplayData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilterDisplayData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilterDisplayData) SetType(v string)`

SetType sets Type field to given value.


### GetInsightData

`func (o *FilterDisplayData) GetInsightData() InsightType`

GetInsightData returns the InsightData field if non-nil, zero value otherwise.

### GetInsightDataOk

`func (o *FilterDisplayData) GetInsightDataOk() (*InsightType, bool)`

GetInsightDataOk returns a tuple with the InsightData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightData

`func (o *FilterDisplayData) SetInsightData(v InsightType)`

SetInsightData sets InsightData field to given value.


### GetSessionRun

`func (o *FilterDisplayData) GetSessionRun() InsightFilterDisplayDataSessionRun`

GetSessionRun returns the SessionRun field if non-nil, zero value otherwise.

### GetSessionRunOk

`func (o *FilterDisplayData) GetSessionRunOk() (*InsightFilterDisplayDataSessionRun, bool)`

GetSessionRunOk returns a tuple with the SessionRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRun

`func (o *FilterDisplayData) SetSessionRun(v InsightFilterDisplayDataSessionRun)`

SetSessionRun sets SessionRun field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


