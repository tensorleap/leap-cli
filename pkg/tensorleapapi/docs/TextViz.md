# TextViz

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | **string** |  | 
**SubTitle** | **string** |  | 
**Guid** | **string** |  | 
**Body** | **[]string** |  | 
**Clusters** | Pointer to [**Clusters**](Clusters.md) |  | [optional] 
**Heatmap** | Pointer to [**Heatmap**](Heatmap.md) |  | [optional] 

## Methods

### NewTextViz

`func NewTextViz(type_ string, title string, subTitle string, guid string, body []string, ) *TextViz`

NewTextViz instantiates a new TextViz object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextVizWithDefaults

`func NewTextVizWithDefaults() *TextViz`

NewTextVizWithDefaults instantiates a new TextViz object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TextViz) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextViz) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextViz) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *TextViz) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TextViz) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TextViz) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSubTitle

`func (o *TextViz) GetSubTitle() string`

GetSubTitle returns the SubTitle field if non-nil, zero value otherwise.

### GetSubTitleOk

`func (o *TextViz) GetSubTitleOk() (*string, bool)`

GetSubTitleOk returns a tuple with the SubTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTitle

`func (o *TextViz) SetSubTitle(v string)`

SetSubTitle sets SubTitle field to given value.


### GetGuid

`func (o *TextViz) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *TextViz) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *TextViz) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetBody

`func (o *TextViz) GetBody() []string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TextViz) GetBodyOk() (*[]string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TextViz) SetBody(v []string)`

SetBody sets Body field to given value.


### GetClusters

`func (o *TextViz) GetClusters() Clusters`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *TextViz) GetClustersOk() (*Clusters, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *TextViz) SetClusters(v Clusters)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *TextViz) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetHeatmap

`func (o *TextViz) GetHeatmap() Heatmap`

GetHeatmap returns the Heatmap field if non-nil, zero value otherwise.

### GetHeatmapOk

`func (o *TextViz) GetHeatmapOk() (*Heatmap, bool)`

GetHeatmapOk returns a tuple with the Heatmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmap

`func (o *TextViz) SetHeatmap(v Heatmap)`

SetHeatmap sets Heatmap field to given value.

### HasHeatmap

`func (o *TextViz) HasHeatmap() bool`

HasHeatmap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


