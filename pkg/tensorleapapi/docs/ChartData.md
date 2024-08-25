# ChartData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HorizontalKey** | Pointer to [**SplitValue**](SplitValue.md) |  | [optional] 
**VerticalKey** | Pointer to [**SplitValue**](SplitValue.md) |  | [optional] 
**Data** | [**GenericDataResponse**](GenericDataResponse.md) |  | 

## Methods

### NewChartData

`func NewChartData(data GenericDataResponse, ) *ChartData`

NewChartData instantiates a new ChartData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartDataWithDefaults

`func NewChartDataWithDefaults() *ChartData`

NewChartDataWithDefaults instantiates a new ChartData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHorizontalKey

`func (o *ChartData) GetHorizontalKey() SplitValue`

GetHorizontalKey returns the HorizontalKey field if non-nil, zero value otherwise.

### GetHorizontalKeyOk

`func (o *ChartData) GetHorizontalKeyOk() (*SplitValue, bool)`

GetHorizontalKeyOk returns a tuple with the HorizontalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalKey

`func (o *ChartData) SetHorizontalKey(v SplitValue)`

SetHorizontalKey sets HorizontalKey field to given value.

### HasHorizontalKey

`func (o *ChartData) HasHorizontalKey() bool`

HasHorizontalKey returns a boolean if a field has been set.

### GetVerticalKey

`func (o *ChartData) GetVerticalKey() SplitValue`

GetVerticalKey returns the VerticalKey field if non-nil, zero value otherwise.

### GetVerticalKeyOk

`func (o *ChartData) GetVerticalKeyOk() (*SplitValue, bool)`

GetVerticalKeyOk returns a tuple with the VerticalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalKey

`func (o *ChartData) SetVerticalKey(v SplitValue)`

SetVerticalKey sets VerticalKey field to given value.

### HasVerticalKey

`func (o *ChartData) HasVerticalKey() bool`

HasVerticalKey returns a boolean if a field has been set.

### GetData

`func (o *ChartData) GetData() GenericDataResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChartData) GetDataOk() (*GenericDataResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChartData) SetData(v GenericDataResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


