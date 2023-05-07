# GetSessionInsightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Insights** | [**[]Insight**](Insight.md) |  | 

## Methods

### NewGetSessionInsightsResponse

`func NewGetSessionInsightsResponse(insights []Insight, ) *GetSessionInsightsResponse`

NewGetSessionInsightsResponse instantiates a new GetSessionInsightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSessionInsightsResponseWithDefaults

`func NewGetSessionInsightsResponseWithDefaults() *GetSessionInsightsResponse`

NewGetSessionInsightsResponseWithDefaults instantiates a new GetSessionInsightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsights

`func (o *GetSessionInsightsResponse) GetInsights() []Insight`

GetInsights returns the Insights field if non-nil, zero value otherwise.

### GetInsightsOk

`func (o *GetSessionInsightsResponse) GetInsightsOk() (*[]Insight, bool)`

GetInsightsOk returns a tuple with the Insights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsights

`func (o *GetSessionInsightsResponse) SetInsights(v []Insight)`

SetInsights sets Insights field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


