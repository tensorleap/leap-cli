# HorizontalBarData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **[]float64** |  | 
**Labels** | **[]string** |  | 
**Gt** | Pointer to **[]float64** |  | [optional] 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 

## Methods

### NewHorizontalBarData

`func NewHorizontalBarData(body []float64, labels []string, type_ DataTypeEnum, ) *HorizontalBarData`

NewHorizontalBarData instantiates a new HorizontalBarData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHorizontalBarDataWithDefaults

`func NewHorizontalBarDataWithDefaults() *HorizontalBarData`

NewHorizontalBarDataWithDefaults instantiates a new HorizontalBarData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *HorizontalBarData) GetBody() []float64`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *HorizontalBarData) GetBodyOk() (*[]float64, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *HorizontalBarData) SetBody(v []float64)`

SetBody sets Body field to given value.


### GetLabels

`func (o *HorizontalBarData) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HorizontalBarData) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HorizontalBarData) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetGt

`func (o *HorizontalBarData) GetGt() []float64`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *HorizontalBarData) GetGtOk() (*[]float64, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *HorizontalBarData) SetGt(v []float64)`

SetGt sets Gt field to given value.

### HasGt

`func (o *HorizontalBarData) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetType

`func (o *HorizontalBarData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HorizontalBarData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HorizontalBarData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


