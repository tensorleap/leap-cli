# SplitAgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderField** | Pointer to **string** |  | [optional] 
**Limit** | **NullableFloat64** |  | 
**Order** | [**OrderType**](OrderType.md) |  | 
**Field** | **string** |  | 
**Interval** | **float64** |  | 
**Distribution** | [**DistributionTypeDistinct**](DistributionTypeDistinct.md) |  | 

## Methods

### NewSplitAgg

`func NewSplitAgg(limit NullableFloat64, order OrderType, field string, interval float64, distribution DistributionTypeDistinct, ) *SplitAgg`

NewSplitAgg instantiates a new SplitAgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitAggWithDefaults

`func NewSplitAggWithDefaults() *SplitAgg`

NewSplitAggWithDefaults instantiates a new SplitAgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderField

`func (o *SplitAgg) GetOrderField() string`

GetOrderField returns the OrderField field if non-nil, zero value otherwise.

### GetOrderFieldOk

`func (o *SplitAgg) GetOrderFieldOk() (*string, bool)`

GetOrderFieldOk returns a tuple with the OrderField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderField

`func (o *SplitAgg) SetOrderField(v string)`

SetOrderField sets OrderField field to given value.

### HasOrderField

`func (o *SplitAgg) HasOrderField() bool`

HasOrderField returns a boolean if a field has been set.

### GetLimit

`func (o *SplitAgg) GetLimit() float64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SplitAgg) GetLimitOk() (*float64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SplitAgg) SetLimit(v float64)`

SetLimit sets Limit field to given value.


### SetLimitNil

`func (o *SplitAgg) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *SplitAgg) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOrder

`func (o *SplitAgg) GetOrder() OrderType`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SplitAgg) GetOrderOk() (*OrderType, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SplitAgg) SetOrder(v OrderType)`

SetOrder sets Order field to given value.


### GetField

`func (o *SplitAgg) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SplitAgg) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SplitAgg) SetField(v string)`

SetField sets Field field to given value.


### GetInterval

`func (o *SplitAgg) GetInterval() float64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SplitAgg) GetIntervalOk() (*float64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SplitAgg) SetInterval(v float64)`

SetInterval sets Interval field to given value.


### GetDistribution

`func (o *SplitAgg) GetDistribution() DistributionTypeDistinct`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *SplitAgg) GetDistributionOk() (*DistributionTypeDistinct, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *SplitAgg) SetDistribution(v DistributionTypeDistinct)`

SetDistribution sets Distribution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


