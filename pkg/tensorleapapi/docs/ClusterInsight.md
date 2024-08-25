# ClusterInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ScatterInsightType**](ScatterInsightType.md) |  | 
**Filter** | [**ScatterFilter**](ScatterFilter.md) |  | 
**MutualInfoElements** | Pointer to [**[]MutualInformationElement**](MutualInformationElement.md) |  | [optional] 
**BlobPath** | Pointer to **string** |  | [optional] 
**InsightSubCategoryDsCuration** | Pointer to [**InsightSubCategoryDSCuration**](InsightSubCategoryDSCuration.md) |  | [optional] 
**InsightCategoryPerformance** | Pointer to [**InsightSubCategoryPerformance**](InsightSubCategoryPerformance.md) |  | [optional] 
**Index** | Pointer to **float64** |  | [optional] 
**NSamples** | **float64** |  | 
**AvgMetric** | **float64** |  | 
**MetricName** | **string** |  | 
**Configuration** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 

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


### GetMutualInfoElements

`func (o *ClusterInsight) GetMutualInfoElements() []MutualInformationElement`

GetMutualInfoElements returns the MutualInfoElements field if non-nil, zero value otherwise.

### GetMutualInfoElementsOk

`func (o *ClusterInsight) GetMutualInfoElementsOk() (*[]MutualInformationElement, bool)`

GetMutualInfoElementsOk returns a tuple with the MutualInfoElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualInfoElements

`func (o *ClusterInsight) SetMutualInfoElements(v []MutualInformationElement)`

SetMutualInfoElements sets MutualInfoElements field to given value.

### HasMutualInfoElements

`func (o *ClusterInsight) HasMutualInfoElements() bool`

HasMutualInfoElements returns a boolean if a field has been set.

### GetBlobPath

`func (o *ClusterInsight) GetBlobPath() string`

GetBlobPath returns the BlobPath field if non-nil, zero value otherwise.

### GetBlobPathOk

`func (o *ClusterInsight) GetBlobPathOk() (*string, bool)`

GetBlobPathOk returns a tuple with the BlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobPath

`func (o *ClusterInsight) SetBlobPath(v string)`

SetBlobPath sets BlobPath field to given value.

### HasBlobPath

`func (o *ClusterInsight) HasBlobPath() bool`

HasBlobPath returns a boolean if a field has been set.

### GetInsightSubCategoryDsCuration

`func (o *ClusterInsight) GetInsightSubCategoryDsCuration() InsightSubCategoryDSCuration`

GetInsightSubCategoryDsCuration returns the InsightSubCategoryDsCuration field if non-nil, zero value otherwise.

### GetInsightSubCategoryDsCurationOk

`func (o *ClusterInsight) GetInsightSubCategoryDsCurationOk() (*InsightSubCategoryDSCuration, bool)`

GetInsightSubCategoryDsCurationOk returns a tuple with the InsightSubCategoryDsCuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightSubCategoryDsCuration

`func (o *ClusterInsight) SetInsightSubCategoryDsCuration(v InsightSubCategoryDSCuration)`

SetInsightSubCategoryDsCuration sets InsightSubCategoryDsCuration field to given value.

### HasInsightSubCategoryDsCuration

`func (o *ClusterInsight) HasInsightSubCategoryDsCuration() bool`

HasInsightSubCategoryDsCuration returns a boolean if a field has been set.

### GetInsightCategoryPerformance

`func (o *ClusterInsight) GetInsightCategoryPerformance() InsightSubCategoryPerformance`

GetInsightCategoryPerformance returns the InsightCategoryPerformance field if non-nil, zero value otherwise.

### GetInsightCategoryPerformanceOk

`func (o *ClusterInsight) GetInsightCategoryPerformanceOk() (*InsightSubCategoryPerformance, bool)`

GetInsightCategoryPerformanceOk returns a tuple with the InsightCategoryPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightCategoryPerformance

`func (o *ClusterInsight) SetInsightCategoryPerformance(v InsightSubCategoryPerformance)`

SetInsightCategoryPerformance sets InsightCategoryPerformance field to given value.

### HasInsightCategoryPerformance

`func (o *ClusterInsight) HasInsightCategoryPerformance() bool`

HasInsightCategoryPerformance returns a boolean if a field has been set.

### GetIndex

`func (o *ClusterInsight) GetIndex() float64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ClusterInsight) GetIndexOk() (*float64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ClusterInsight) SetIndex(v float64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ClusterInsight) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

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


### GetConfiguration

`func (o *ClusterInsight) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ClusterInsight) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ClusterInsight) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ClusterInsight) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


