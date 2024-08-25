# ConfusionMatrixTableParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SplitByLabelOrder** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**SplitByLabel** | **bool** |  | 
**FilterLabels** | Pointer to **[]string** |  | [optional] 
**Threshold** | Pointer to **float64** |  | [optional] 

## Methods

### NewConfusionMatrixTableParamsAllOf

`func NewConfusionMatrixTableParamsAllOf(splitByLabel bool, ) *ConfusionMatrixTableParamsAllOf`

NewConfusionMatrixTableParamsAllOf instantiates a new ConfusionMatrixTableParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfusionMatrixTableParamsAllOfWithDefaults

`func NewConfusionMatrixTableParamsAllOfWithDefaults() *ConfusionMatrixTableParamsAllOf`

NewConfusionMatrixTableParamsAllOfWithDefaults instantiates a new ConfusionMatrixTableParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSplitByLabelOrder

`func (o *ConfusionMatrixTableParamsAllOf) GetSplitByLabelOrder() OrderType`

GetSplitByLabelOrder returns the SplitByLabelOrder field if non-nil, zero value otherwise.

### GetSplitByLabelOrderOk

`func (o *ConfusionMatrixTableParamsAllOf) GetSplitByLabelOrderOk() (*OrderType, bool)`

GetSplitByLabelOrderOk returns a tuple with the SplitByLabelOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitByLabelOrder

`func (o *ConfusionMatrixTableParamsAllOf) SetSplitByLabelOrder(v OrderType)`

SetSplitByLabelOrder sets SplitByLabelOrder field to given value.

### HasSplitByLabelOrder

`func (o *ConfusionMatrixTableParamsAllOf) HasSplitByLabelOrder() bool`

HasSplitByLabelOrder returns a boolean if a field has been set.

### GetSplitByLabel

`func (o *ConfusionMatrixTableParamsAllOf) GetSplitByLabel() bool`

GetSplitByLabel returns the SplitByLabel field if non-nil, zero value otherwise.

### GetSplitByLabelOk

`func (o *ConfusionMatrixTableParamsAllOf) GetSplitByLabelOk() (*bool, bool)`

GetSplitByLabelOk returns a tuple with the SplitByLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitByLabel

`func (o *ConfusionMatrixTableParamsAllOf) SetSplitByLabel(v bool)`

SetSplitByLabel sets SplitByLabel field to given value.


### GetFilterLabels

`func (o *ConfusionMatrixTableParamsAllOf) GetFilterLabels() []string`

GetFilterLabels returns the FilterLabels field if non-nil, zero value otherwise.

### GetFilterLabelsOk

`func (o *ConfusionMatrixTableParamsAllOf) GetFilterLabelsOk() (*[]string, bool)`

GetFilterLabelsOk returns a tuple with the FilterLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLabels

`func (o *ConfusionMatrixTableParamsAllOf) SetFilterLabels(v []string)`

SetFilterLabels sets FilterLabels field to given value.

### HasFilterLabels

`func (o *ConfusionMatrixTableParamsAllOf) HasFilterLabels() bool`

HasFilterLabels returns a boolean if a field has been set.

### GetThreshold

`func (o *ConfusionMatrixTableParamsAllOf) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ConfusionMatrixTableParamsAllOf) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ConfusionMatrixTableParamsAllOf) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ConfusionMatrixTableParamsAllOf) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


