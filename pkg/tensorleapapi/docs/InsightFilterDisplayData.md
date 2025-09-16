# InsightFilterDisplayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Insights** | [**[]InsightFilterDisplayDataInsightsInner**](InsightFilterDisplayDataInsightsInner.md) |  | 

## Methods

### NewInsightFilterDisplayData

`func NewInsightFilterDisplayData(type_ string, insights []InsightFilterDisplayDataInsightsInner, ) *InsightFilterDisplayData`

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


### GetInsights

`func (o *InsightFilterDisplayData) GetInsights() []InsightFilterDisplayDataInsightsInner`

GetInsights returns the Insights field if non-nil, zero value otherwise.

### GetInsightsOk

`func (o *InsightFilterDisplayData) GetInsightsOk() (*[]InsightFilterDisplayDataInsightsInner, bool)`

GetInsightsOk returns a tuple with the Insights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsights

`func (o *InsightFilterDisplayData) SetInsights(v []InsightFilterDisplayDataInsightsInner)`

SetInsights sets Insights field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


