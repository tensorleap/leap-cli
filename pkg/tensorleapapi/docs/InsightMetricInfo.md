# InsightMetricInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | **string** |  | 
**MetricStatistics** | [**[]MetricStatisticElement**](MetricStatisticElement.md) |  | 

## Methods

### NewInsightMetricInfo

`func NewInsightMetricInfo(metricName string, metricStatistics []MetricStatisticElement, ) *InsightMetricInfo`

NewInsightMetricInfo instantiates a new InsightMetricInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightMetricInfoWithDefaults

`func NewInsightMetricInfoWithDefaults() *InsightMetricInfo`

NewInsightMetricInfoWithDefaults instantiates a new InsightMetricInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *InsightMetricInfo) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *InsightMetricInfo) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *InsightMetricInfo) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetMetricStatistics

`func (o *InsightMetricInfo) GetMetricStatistics() []MetricStatisticElement`

GetMetricStatistics returns the MetricStatistics field if non-nil, zero value otherwise.

### GetMetricStatisticsOk

`func (o *InsightMetricInfo) GetMetricStatisticsOk() (*[]MetricStatisticElement, bool)`

GetMetricStatisticsOk returns a tuple with the MetricStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricStatistics

`func (o *InsightMetricInfo) SetMetricStatistics(v []MetricStatisticElement)`

SetMetricStatistics sets MetricStatistics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


