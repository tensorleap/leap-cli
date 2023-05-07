# TextData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **[]string** |  | 
**Heatmap** | Pointer to [**Heatmap**](Heatmap.md) |  | [optional] 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 

## Methods

### NewTextData

`func NewTextData(body []string, type_ DataTypeEnum, ) *TextData`

NewTextData instantiates a new TextData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextDataWithDefaults

`func NewTextDataWithDefaults() *TextData`

NewTextDataWithDefaults instantiates a new TextData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *TextData) GetBody() []string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TextData) GetBodyOk() (*[]string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TextData) SetBody(v []string)`

SetBody sets Body field to given value.


### GetHeatmap

`func (o *TextData) GetHeatmap() Heatmap`

GetHeatmap returns the Heatmap field if non-nil, zero value otherwise.

### GetHeatmapOk

`func (o *TextData) GetHeatmapOk() (*Heatmap, bool)`

GetHeatmapOk returns a tuple with the Heatmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmap

`func (o *TextData) SetHeatmap(v Heatmap)`

SetHeatmap sets Heatmap field to given value.

### HasHeatmap

`func (o *TextData) HasHeatmap() bool`

HasHeatmap returns a boolean if a field has been set.

### GetType

`func (o *TextData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


