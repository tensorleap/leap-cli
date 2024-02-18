# ESFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**FilterOperatorType**](FilterOperatorType.md) |  | 
**Field** | **string** |  | 
**Disable** | Pointer to **bool** |  | [optional] 
**Value** | [**ESFilterValue**](ESFilterValue.md) |  | 
**DisplayData** | Pointer to [**FilterDisplayData**](FilterDisplayData.md) |  | [optional] 

## Methods

### NewESFilter

`func NewESFilter(operator FilterOperatorType, field string, value ESFilterValue, ) *ESFilter`

NewESFilter instantiates a new ESFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESFilterWithDefaults

`func NewESFilterWithDefaults() *ESFilter`

NewESFilterWithDefaults instantiates a new ESFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ESFilter) GetOperator() FilterOperatorType`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ESFilter) GetOperatorOk() (*FilterOperatorType, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ESFilter) SetOperator(v FilterOperatorType)`

SetOperator sets Operator field to given value.


### GetField

`func (o *ESFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ESFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ESFilter) SetField(v string)`

SetField sets Field field to given value.


### GetDisable

`func (o *ESFilter) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *ESFilter) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *ESFilter) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *ESFilter) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetValue

`func (o *ESFilter) GetValue() ESFilterValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ESFilter) GetValueOk() (*ESFilterValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ESFilter) SetValue(v ESFilterValue)`

SetValue sets Value field to given value.


### GetDisplayData

`func (o *ESFilter) GetDisplayData() FilterDisplayData`

GetDisplayData returns the DisplayData field if non-nil, zero value otherwise.

### GetDisplayDataOk

`func (o *ESFilter) GetDisplayDataOk() (*FilterDisplayData, bool)`

GetDisplayDataOk returns a tuple with the DisplayData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayData

`func (o *ESFilter) SetDisplayData(v FilterDisplayData)`

SetDisplayData sets DisplayData field to given value.

### HasDisplayData

`func (o *ESFilter) HasDisplayData() bool`

HasDisplayData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


