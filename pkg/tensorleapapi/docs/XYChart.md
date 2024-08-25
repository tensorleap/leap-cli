# XYChart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | [**SplitAgg**](SplitAgg.md) |  | 
**Y** | [**AggFunc**](AggFunc.md) |  | 

## Methods

### NewXYChart

`func NewXYChart(x SplitAgg, y AggFunc, ) *XYChart`

NewXYChart instantiates a new XYChart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXYChartWithDefaults

`func NewXYChartWithDefaults() *XYChart`

NewXYChartWithDefaults instantiates a new XYChart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *XYChart) GetX() SplitAgg`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *XYChart) GetXOk() (*SplitAgg, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *XYChart) SetX(v SplitAgg)`

SetX sets X field to given value.


### GetY

`func (o *XYChart) GetY() AggFunc`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *XYChart) GetYOk() (*AggFunc, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *XYChart) SetY(v AggFunc)`

SetY sets Y field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


