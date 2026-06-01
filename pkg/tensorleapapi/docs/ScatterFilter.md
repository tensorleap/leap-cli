# ScatterFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** |  | 
**Value** | [**ScatterFilterValue**](ScatterFilterValue.md) |  | 
**Operator** | [**OperatorEnum**](OperatorEnum.md) |  | 
**RankBy** | Pointer to **string** |  | [optional] 

## Methods

### NewScatterFilter

`func NewScatterFilter(metric string, value ScatterFilterValue, operator OperatorEnum, ) *ScatterFilter`

NewScatterFilter instantiates a new ScatterFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScatterFilterWithDefaults

`func NewScatterFilterWithDefaults() *ScatterFilter`

NewScatterFilterWithDefaults instantiates a new ScatterFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *ScatterFilter) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ScatterFilter) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ScatterFilter) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetValue

`func (o *ScatterFilter) GetValue() ScatterFilterValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScatterFilter) GetValueOk() (*ScatterFilterValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScatterFilter) SetValue(v ScatterFilterValue)`

SetValue sets Value field to given value.


### GetOperator

`func (o *ScatterFilter) GetOperator() OperatorEnum`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ScatterFilter) GetOperatorOk() (*OperatorEnum, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ScatterFilter) SetOperator(v OperatorEnum)`

SetOperator sets Operator field to given value.


### GetRankBy

`func (o *ScatterFilter) GetRankBy() string`

GetRankBy returns the RankBy field if non-nil, zero value otherwise.

### GetRankByOk

`func (o *ScatterFilter) GetRankByOk() (*string, bool)`

GetRankByOk returns a tuple with the RankBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankBy

`func (o *ScatterFilter) SetRankBy(v string)`

SetRankBy sets RankBy field to given value.

### HasRankBy

`func (o *ScatterFilter) HasRankBy() bool`

HasRankBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


