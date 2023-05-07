# ClusterInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ScatterInsightType**](ScatterInsightType.md) |  | 
**Filter** | [**ScatterFilter**](ScatterFilter.md) |  | 
**NSamples** | **float64** |  | 
**AvgMetric** | **float64** |  | 
**MetricName** | **string** |  | 

## Methods

### NewClusterInsight

`func NewClusterInsight(type_ ScatterInsightType, filter ScatterFilter, nSamples float64, avgMetric float64, metricName string, ) *ClusterInsight`

NewClusterInsight instantiates a new ClusterInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInsightWithDefaults

`func NewClusterInsightWithDefaults() *ClusterInsight`

NewClusterInsightWithDefaults instantiates a new ClusterInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClusterInsight) GetType() ScatterInsightType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterInsight) GetTypeOk() (*ScatterInsightType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterInsight) SetType(v ScatterInsightType)`

SetType sets Type field to given value.


### GetFilter

`func (o *ClusterInsight) GetFilter() ScatterFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ClusterInsight) GetFilterOk() (*ScatterFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ClusterInsight) SetFilter(v ScatterFilter)`

SetFilter sets Filter field to given value.


### GetNSamples

`func (o *ClusterInsight) GetNSamples() float64`

GetNSamples returns the NSamples field if non-nil, zero value otherwise.

### GetNSamplesOk

`func (o *ClusterInsight) GetNSamplesOk() (*float64, bool)`

GetNSamplesOk returns a tuple with the NSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSamples

`func (o *ClusterInsight) SetNSamples(v float64)`

SetNSamples sets NSamples field to given value.


### GetAvgMetric

`func (o *ClusterInsight) GetAvgMetric() float64`

GetAvgMetric returns the AvgMetric field if non-nil, zero value otherwise.

### GetAvgMetricOk

`func (o *ClusterInsight) GetAvgMetricOk() (*float64, bool)`

GetAvgMetricOk returns a tuple with the AvgMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgMetric

`func (o *ClusterInsight) SetAvgMetric(v float64)`

SetAvgMetric sets AvgMetric field to given value.


### GetMetricName

`func (o *ClusterInsight) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ClusterInsight) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ClusterInsight) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


