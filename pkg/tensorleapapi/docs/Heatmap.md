# Heatmap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **[]float64** |  | 
**Type** | [**HeatmapType**](HeatmapType.md) |  | 

## Methods

### NewHeatmap

`func NewHeatmap(body []float64, type_ HeatmapType, ) *Heatmap`

NewHeatmap instantiates a new Heatmap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeatmapWithDefaults

`func NewHeatmapWithDefaults() *Heatmap`

NewHeatmapWithDefaults instantiates a new Heatmap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *Heatmap) GetBody() []float64`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Heatmap) GetBodyOk() (*[]float64, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Heatmap) SetBody(v []float64)`

SetBody sets Body field to given value.


### GetType

`func (o *Heatmap) GetType() HeatmapType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Heatmap) GetTypeOk() (*HeatmapType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Heatmap) SetType(v HeatmapType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


