# GraphViz

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | **string** |  | 
**SubTitle** | **string** |  | 
**Guid** | **string** |  | 
**Body** | **[][]float64** |  | 
**Heatmap** | Pointer to [**Heatmap**](Heatmap.md) |  | [optional] 

## Methods

### NewGraphViz

`func NewGraphViz(type_ string, title string, subTitle string, guid string, body [][]float64, ) *GraphViz`

NewGraphViz instantiates a new GraphViz object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVizWithDefaults

`func NewGraphVizWithDefaults() *GraphViz`

NewGraphVizWithDefaults instantiates a new GraphViz object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphViz) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphViz) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphViz) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *GraphViz) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GraphViz) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GraphViz) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSubTitle

`func (o *GraphViz) GetSubTitle() string`

GetSubTitle returns the SubTitle field if non-nil, zero value otherwise.

### GetSubTitleOk

`func (o *GraphViz) GetSubTitleOk() (*string, bool)`

GetSubTitleOk returns a tuple with the SubTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTitle

`func (o *GraphViz) SetSubTitle(v string)`

SetSubTitle sets SubTitle field to given value.


### GetGuid

`func (o *GraphViz) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *GraphViz) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *GraphViz) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetBody

`func (o *GraphViz) GetBody() [][]float64`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GraphViz) GetBodyOk() (*[][]float64, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GraphViz) SetBody(v [][]float64)`

SetBody sets Body field to given value.


### GetHeatmap

`func (o *GraphViz) GetHeatmap() Heatmap`

GetHeatmap returns the Heatmap field if non-nil, zero value otherwise.

### GetHeatmapOk

`func (o *GraphViz) GetHeatmapOk() (*Heatmap, bool)`

GetHeatmapOk returns a tuple with the Heatmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmap

`func (o *GraphViz) SetHeatmap(v Heatmap)`

SetHeatmap sets Heatmap field to given value.

### HasHeatmap

`func (o *GraphViz) HasHeatmap() bool`

HasHeatmap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


