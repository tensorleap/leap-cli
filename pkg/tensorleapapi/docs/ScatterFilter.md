# ScatterFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** |  | 
**Value** | **string** |  | 
**Operator** | [**OperatorEnum**](OperatorEnum.md) |  | 
**DataState** | Pointer to [**DataStateType**](DataStateType.md) |  | [optional] 

## Methods

### NewScatterFilter

`func NewScatterFilter(metric string, value string, operator OperatorEnum, ) *ScatterFilter`

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

`func (o *ScatterFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScatterFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScatterFilter) SetValue(v string)`

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


### GetDataState

`func (o *ScatterFilter) GetDataState() DataStateType`

GetDataState returns the DataState field if non-nil, zero value otherwise.

### GetDataStateOk

`func (o *ScatterFilter) GetDataStateOk() (*DataStateType, bool)`

GetDataStateOk returns a tuple with the DataState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataState

`func (o *ScatterFilter) SetDataState(v DataStateType)`

SetDataState sets DataState field to given value.

### HasDataState

`func (o *ScatterFilter) HasDataState() bool`

HasDataState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


