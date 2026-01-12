# DistinctAgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderField** | Pointer to **string** |  | [optional] 
**Limit** | **NullableFloat64** |  | 
**Order** | [**OrderType**](OrderType.md) |  | 
**Field** | **string** |  | 
**Distribution** | [**DistributionTypeDistinct**](DistributionTypeDistinct.md) |  | 

## Methods

### NewDistinctAgg

`func NewDistinctAgg(limit NullableFloat64, order OrderType, field string, distribution DistributionTypeDistinct, ) *DistinctAgg`

NewDistinctAgg instantiates a new DistinctAgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistinctAggWithDefaults

`func NewDistinctAggWithDefaults() *DistinctAgg`

NewDistinctAggWithDefaults instantiates a new DistinctAgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderField

`func (o *DistinctAgg) GetOrderField() string`

GetOrderField returns the OrderField field if non-nil, zero value otherwise.

### GetOrderFieldOk

`func (o *DistinctAgg) GetOrderFieldOk() (*string, bool)`

GetOrderFieldOk returns a tuple with the OrderField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderField

`func (o *DistinctAgg) SetOrderField(v string)`

SetOrderField sets OrderField field to given value.

### HasOrderField

`func (o *DistinctAgg) HasOrderField() bool`

HasOrderField returns a boolean if a field has been set.

### GetLimit

`func (o *DistinctAgg) GetLimit() float64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DistinctAgg) GetLimitOk() (*float64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DistinctAgg) SetLimit(v float64)`

SetLimit sets Limit field to given value.


### SetLimitNil

`func (o *DistinctAgg) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *DistinctAgg) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOrder

`func (o *DistinctAgg) GetOrder() OrderType`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *DistinctAgg) GetOrderOk() (*OrderType, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *DistinctAgg) SetOrder(v OrderType)`

SetOrder sets Order field to given value.


### GetField

`func (o *DistinctAgg) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *DistinctAgg) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *DistinctAgg) SetField(v string)`

SetField sets Field field to given value.


### GetDistribution

`func (o *DistinctAgg) GetDistribution() DistributionTypeDistinct`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *DistinctAgg) GetDistributionOk() (*DistributionTypeDistinct, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *DistinctAgg) SetDistribution(v DistributionTypeDistinct)`

SetDistribution sets Distribution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


