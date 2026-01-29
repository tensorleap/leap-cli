# BaseSplitAgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderField** | Pointer to **string** |  | [optional] 
**Limit** | **NullableFloat64** |  | 
**Order** | [**OrderType**](OrderType.md) |  | 
**Field** | **string** |  | 

## Methods

### NewBaseSplitAgg

`func NewBaseSplitAgg(limit NullableFloat64, order OrderType, field string, ) *BaseSplitAgg`

NewBaseSplitAgg instantiates a new BaseSplitAgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseSplitAggWithDefaults

`func NewBaseSplitAggWithDefaults() *BaseSplitAgg`

NewBaseSplitAggWithDefaults instantiates a new BaseSplitAgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderField

`func (o *BaseSplitAgg) GetOrderField() string`

GetOrderField returns the OrderField field if non-nil, zero value otherwise.

### GetOrderFieldOk

`func (o *BaseSplitAgg) GetOrderFieldOk() (*string, bool)`

GetOrderFieldOk returns a tuple with the OrderField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderField

`func (o *BaseSplitAgg) SetOrderField(v string)`

SetOrderField sets OrderField field to given value.

### HasOrderField

`func (o *BaseSplitAgg) HasOrderField() bool`

HasOrderField returns a boolean if a field has been set.

### GetLimit

`func (o *BaseSplitAgg) GetLimit() float64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BaseSplitAgg) GetLimitOk() (*float64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BaseSplitAgg) SetLimit(v float64)`

SetLimit sets Limit field to given value.


### SetLimitNil

`func (o *BaseSplitAgg) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *BaseSplitAgg) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOrder

`func (o *BaseSplitAgg) GetOrder() OrderType`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *BaseSplitAgg) GetOrderOk() (*OrderType, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *BaseSplitAgg) SetOrder(v OrderType)`

SetOrder sets Order field to given value.


### GetField

`func (o *BaseSplitAgg) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *BaseSplitAgg) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *BaseSplitAgg) SetField(v string)`

SetField sets Field field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


