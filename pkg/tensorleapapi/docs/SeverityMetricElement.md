# SeverityMetricElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | **string** |  | 
**Value** | **float64** |  | 
**NormalizedValue** | Pointer to **float64** |  | [optional] 

## Methods

### NewSeverityMetricElement

`func NewSeverityMetricElement(metricName string, value float64, ) *SeverityMetricElement`

NewSeverityMetricElement instantiates a new SeverityMetricElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeverityMetricElementWithDefaults

`func NewSeverityMetricElementWithDefaults() *SeverityMetricElement`

NewSeverityMetricElementWithDefaults instantiates a new SeverityMetricElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *SeverityMetricElement) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *SeverityMetricElement) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *SeverityMetricElement) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetValue

`func (o *SeverityMetricElement) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SeverityMetricElement) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SeverityMetricElement) SetValue(v float64)`

SetValue sets Value field to given value.


### GetNormalizedValue

`func (o *SeverityMetricElement) GetNormalizedValue() float64`

GetNormalizedValue returns the NormalizedValue field if non-nil, zero value otherwise.

### GetNormalizedValueOk

`func (o *SeverityMetricElement) GetNormalizedValueOk() (*float64, bool)`

GetNormalizedValueOk returns a tuple with the NormalizedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedValue

`func (o *SeverityMetricElement) SetNormalizedValue(v float64)`

SetNormalizedValue sets NormalizedValue field to given value.

### HasNormalizedValue

`func (o *SeverityMetricElement) HasNormalizedValue() bool`

HasNormalizedValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


