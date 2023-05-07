# ClientFilterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Operator** | [**FilterOperatorType**](FilterOperatorType.md) |  | 
**Value** | **float64** |  | 
**Disable** | Pointer to **bool** |  | [optional] 

## Methods

### NewClientFilterParams

`func NewClientFilterParams(field string, operator FilterOperatorType, value float64, ) *ClientFilterParams`

NewClientFilterParams instantiates a new ClientFilterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientFilterParamsWithDefaults

`func NewClientFilterParamsWithDefaults() *ClientFilterParams`

NewClientFilterParamsWithDefaults instantiates a new ClientFilterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ClientFilterParams) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ClientFilterParams) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ClientFilterParams) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *ClientFilterParams) GetOperator() FilterOperatorType`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ClientFilterParams) GetOperatorOk() (*FilterOperatorType, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ClientFilterParams) SetOperator(v FilterOperatorType)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *ClientFilterParams) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ClientFilterParams) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ClientFilterParams) SetValue(v float64)`

SetValue sets Value field to given value.


### GetDisable

`func (o *ClientFilterParams) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *ClientFilterParams) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *ClientFilterParams) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *ClientFilterParams) HasDisable() bool`

HasDisable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


