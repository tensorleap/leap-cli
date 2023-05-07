# BucketAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**XField** | **string** |  | 
**DataDistributionType** | [**DataDistributionType**](DataDistributionType.md) |  | 
**XAxisSizeInterval** | **float64** |  | 
**OrderParams** | **string** |  | 

## Methods

### NewBucketAggregation

`func NewBucketAggregation(xField string, dataDistributionType DataDistributionType, xAxisSizeInterval float64, orderParams string, ) *BucketAggregation`

NewBucketAggregation instantiates a new BucketAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketAggregationWithDefaults

`func NewBucketAggregationWithDefaults() *BucketAggregation`

NewBucketAggregationWithDefaults instantiates a new BucketAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXField

`func (o *BucketAggregation) GetXField() string`

GetXField returns the XField field if non-nil, zero value otherwise.

### GetXFieldOk

`func (o *BucketAggregation) GetXFieldOk() (*string, bool)`

GetXFieldOk returns a tuple with the XField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXField

`func (o *BucketAggregation) SetXField(v string)`

SetXField sets XField field to given value.


### GetDataDistributionType

`func (o *BucketAggregation) GetDataDistributionType() DataDistributionType`

GetDataDistributionType returns the DataDistributionType field if non-nil, zero value otherwise.

### GetDataDistributionTypeOk

`func (o *BucketAggregation) GetDataDistributionTypeOk() (*DataDistributionType, bool)`

GetDataDistributionTypeOk returns a tuple with the DataDistributionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDistributionType

`func (o *BucketAggregation) SetDataDistributionType(v DataDistributionType)`

SetDataDistributionType sets DataDistributionType field to given value.


### GetXAxisSizeInterval

`func (o *BucketAggregation) GetXAxisSizeInterval() float64`

GetXAxisSizeInterval returns the XAxisSizeInterval field if non-nil, zero value otherwise.

### GetXAxisSizeIntervalOk

`func (o *BucketAggregation) GetXAxisSizeIntervalOk() (*float64, bool)`

GetXAxisSizeIntervalOk returns a tuple with the XAxisSizeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAxisSizeInterval

`func (o *BucketAggregation) SetXAxisSizeInterval(v float64)`

SetXAxisSizeInterval sets XAxisSizeInterval field to given value.


### GetOrderParams

`func (o *BucketAggregation) GetOrderParams() string`

GetOrderParams returns the OrderParams field if non-nil, zero value otherwise.

### GetOrderParamsOk

`func (o *BucketAggregation) GetOrderParamsOk() (*string, bool)`

GetOrderParamsOk returns a tuple with the OrderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderParams

`func (o *BucketAggregation) SetOrderParams(v string)`

SetOrderParams sets OrderParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


