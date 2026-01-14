# InsightAutomaticTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestName** | **string** |  | 
**Filter** | [**ESFilter**](ESFilter.md) |  | 
**MetricName** | **string** |  | 
**MetricValue** | **float64** |  | 
**Operator** | [**TestFilterOperatorType**](TestFilterOperatorType.md) |  | 

## Methods

### NewInsightAutomaticTest

`func NewInsightAutomaticTest(testName string, filter ESFilter, metricName string, metricValue float64, operator TestFilterOperatorType, ) *InsightAutomaticTest`

NewInsightAutomaticTest instantiates a new InsightAutomaticTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightAutomaticTestWithDefaults

`func NewInsightAutomaticTestWithDefaults() *InsightAutomaticTest`

NewInsightAutomaticTestWithDefaults instantiates a new InsightAutomaticTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestName

`func (o *InsightAutomaticTest) GetTestName() string`

GetTestName returns the TestName field if non-nil, zero value otherwise.

### GetTestNameOk

`func (o *InsightAutomaticTest) GetTestNameOk() (*string, bool)`

GetTestNameOk returns a tuple with the TestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestName

`func (o *InsightAutomaticTest) SetTestName(v string)`

SetTestName sets TestName field to given value.


### GetFilter

`func (o *InsightAutomaticTest) GetFilter() ESFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InsightAutomaticTest) GetFilterOk() (*ESFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InsightAutomaticTest) SetFilter(v ESFilter)`

SetFilter sets Filter field to given value.


### GetMetricName

`func (o *InsightAutomaticTest) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *InsightAutomaticTest) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *InsightAutomaticTest) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetMetricValue

`func (o *InsightAutomaticTest) GetMetricValue() float64`

GetMetricValue returns the MetricValue field if non-nil, zero value otherwise.

### GetMetricValueOk

`func (o *InsightAutomaticTest) GetMetricValueOk() (*float64, bool)`

GetMetricValueOk returns a tuple with the MetricValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValue

`func (o *InsightAutomaticTest) SetMetricValue(v float64)`

SetMetricValue sets MetricValue field to given value.


### GetOperator

`func (o *InsightAutomaticTest) GetOperator() TestFilterOperatorType`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *InsightAutomaticTest) GetOperatorOk() (*TestFilterOperatorType, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *InsightAutomaticTest) SetOperator(v TestFilterOperatorType)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


