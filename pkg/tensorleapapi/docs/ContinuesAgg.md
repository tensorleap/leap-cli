# ContinuesAgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Order** | [**OrderType**](OrderType.md) |  | 
**Limit** | Pointer to **float64** |  | [optional] 
**OrderField** | **string** |  | 
**Interval** | **float64** |  | 
**Distribution** | [**DistributionType**](DistributionType.md) |  | 

## Methods

### NewContinuesAgg

`func NewContinuesAgg(field string, order OrderType, orderField string, interval float64, distribution DistributionType, ) *ContinuesAgg`

NewContinuesAgg instantiates a new ContinuesAgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContinuesAggWithDefaults

`func NewContinuesAggWithDefaults() *ContinuesAgg`

NewContinuesAggWithDefaults instantiates a new ContinuesAgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ContinuesAgg) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ContinuesAgg) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ContinuesAgg) SetField(v string)`

SetField sets Field field to given value.


### GetOrder

`func (o *ContinuesAgg) GetOrder() OrderType`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ContinuesAgg) GetOrderOk() (*OrderType, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ContinuesAgg) SetOrder(v OrderType)`

SetOrder sets Order field to given value.


### GetLimit

`func (o *ContinuesAgg) GetLimit() float64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ContinuesAgg) GetLimitOk() (*float64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ContinuesAgg) SetLimit(v float64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ContinuesAgg) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderField

`func (o *ContinuesAgg) GetOrderField() string`

GetOrderField returns the OrderField field if non-nil, zero value otherwise.

### GetOrderFieldOk

`func (o *ContinuesAgg) GetOrderFieldOk() (*string, bool)`

GetOrderFieldOk returns a tuple with the OrderField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderField

`func (o *ContinuesAgg) SetOrderField(v string)`

SetOrderField sets OrderField field to given value.


### GetInterval

`func (o *ContinuesAgg) GetInterval() float64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ContinuesAgg) GetIntervalOk() (*float64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ContinuesAgg) SetInterval(v float64)`

SetInterval sets Interval field to given value.


### GetDistribution

`func (o *ContinuesAgg) GetDistribution() DistributionType`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *ContinuesAgg) GetDistributionOk() (*DistributionType, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *ContinuesAgg) SetDistribution(v DistributionType)`

SetDistribution sets Distribution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


