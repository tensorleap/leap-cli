# AggFunc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Aggregation** | [**AggregationMethod**](AggregationMethod.md) |  | 

## Methods

### NewAggFunc

`func NewAggFunc(field string, aggregation AggregationMethod, ) *AggFunc`

NewAggFunc instantiates a new AggFunc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggFuncWithDefaults

`func NewAggFuncWithDefaults() *AggFunc`

NewAggFuncWithDefaults instantiates a new AggFunc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AggFunc) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggFunc) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggFunc) SetField(v string)`

SetField sets Field field to given value.


### GetAggregation

`func (o *AggFunc) GetAggregation() AggregationMethod`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *AggFunc) GetAggregationOk() (*AggregationMethod, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *AggFunc) SetAggregation(v AggregationMethod)`

SetAggregation sets Aggregation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


