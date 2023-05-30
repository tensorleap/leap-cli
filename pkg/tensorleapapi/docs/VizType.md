# VizType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | **string** |  | 
**SubTitle** | **string** |  | 
**Guid** | **string** |  | 
**Blob** | **string** |  | 
**Labels** | **[]string** |  | 
**Body** | **[]float64** |  | 
**Heatmap** | Pointer to [**Heatmap**](Heatmap.md) |  | [optional] 
**Clusters** | Pointer to [**Clusters**](Clusters.md) |  | [optional] 
**ScatterData** | [**ScatterVizDataState**](ScatterVizDataState.md) |  | 
**BoundingBoxes** | [**[]BoundingBox**](BoundingBox.md) |  | 
**VisualizedItems** | [**[]VisualizedItem**](VisualizedItem.md) |  | 
**GradsAnalysis** | [**GradsAnalysis**](GradsAnalysis.md) |  | 
**FeatureImportance** | [**FeatureImportance**](FeatureImportance.md) |  | 
**MetadataMap** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 

## Methods

### NewVizType

`func NewVizType(type_ string, title string, subTitle string, guid string, blob string, labels []string, body []float64, scatterData ScatterVizDataState, boundingBoxes []BoundingBox, visualizedItems []VisualizedItem, gradsAnalysis GradsAnalysis, featureImportance FeatureImportance, metadataMap map[string]interface{}, ) *VizType`

NewVizType instantiates a new VizType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVizTypeWithDefaults

`func NewVizTypeWithDefaults() *VizType`

NewVizTypeWithDefaults instantiates a new VizType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VizType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VizType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VizType) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *VizType) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VizType) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VizType) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSubTitle

`func (o *VizType) GetSubTitle() string`

GetSubTitle returns the SubTitle field if non-nil, zero value otherwise.

### GetSubTitleOk

`func (o *VizType) GetSubTitleOk() (*string, bool)`

GetSubTitleOk returns a tuple with the SubTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTitle

`func (o *VizType) SetSubTitle(v string)`

SetSubTitle sets SubTitle field to given value.


### GetGuid

`func (o *VizType) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *VizType) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *VizType) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetBlob

`func (o *VizType) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *VizType) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *VizType) SetBlob(v string)`

SetBlob sets Blob field to given value.


### GetLabels

`func (o *VizType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *VizType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *VizType) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetBody

`func (o *VizType) GetBody() []float64`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *VizType) GetBodyOk() (*[]float64, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *VizType) SetBody(v []float64)`

SetBody sets Body field to given value.


### GetHeatmap

`func (o *VizType) GetHeatmap() Heatmap`

GetHeatmap returns the Heatmap field if non-nil, zero value otherwise.

### GetHeatmapOk

`func (o *VizType) GetHeatmapOk() (*Heatmap, bool)`

GetHeatmapOk returns a tuple with the Heatmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmap

`func (o *VizType) SetHeatmap(v Heatmap)`

SetHeatmap sets Heatmap field to given value.

### HasHeatmap

`func (o *VizType) HasHeatmap() bool`

HasHeatmap returns a boolean if a field has been set.

### GetClusters

`func (o *VizType) GetClusters() Clusters`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *VizType) GetClustersOk() (*Clusters, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *VizType) SetClusters(v Clusters)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *VizType) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetScatterData

`func (o *VizType) GetScatterData() ScatterVizDataState`

GetScatterData returns the ScatterData field if non-nil, zero value otherwise.

### GetScatterDataOk

`func (o *VizType) GetScatterDataOk() (*ScatterVizDataState, bool)`

GetScatterDataOk returns a tuple with the ScatterData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScatterData

`func (o *VizType) SetScatterData(v ScatterVizDataState)`

SetScatterData sets ScatterData field to given value.


### GetBoundingBoxes

`func (o *VizType) GetBoundingBoxes() []BoundingBox`

GetBoundingBoxes returns the BoundingBoxes field if non-nil, zero value otherwise.

### GetBoundingBoxesOk

`func (o *VizType) GetBoundingBoxesOk() (*[]BoundingBox, bool)`

GetBoundingBoxesOk returns a tuple with the BoundingBoxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingBoxes

`func (o *VizType) SetBoundingBoxes(v []BoundingBox)`

SetBoundingBoxes sets BoundingBoxes field to given value.


### GetVisualizedItems

`func (o *VizType) GetVisualizedItems() []VisualizedItem`

GetVisualizedItems returns the VisualizedItems field if non-nil, zero value otherwise.

### GetVisualizedItemsOk

`func (o *VizType) GetVisualizedItemsOk() (*[]VisualizedItem, bool)`

GetVisualizedItemsOk returns a tuple with the VisualizedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizedItems

`func (o *VizType) SetVisualizedItems(v []VisualizedItem)`

SetVisualizedItems sets VisualizedItems field to given value.


### GetGradsAnalysis

`func (o *VizType) GetGradsAnalysis() GradsAnalysis`

GetGradsAnalysis returns the GradsAnalysis field if non-nil, zero value otherwise.

### GetGradsAnalysisOk

`func (o *VizType) GetGradsAnalysisOk() (*GradsAnalysis, bool)`

GetGradsAnalysisOk returns a tuple with the GradsAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradsAnalysis

`func (o *VizType) SetGradsAnalysis(v GradsAnalysis)`

SetGradsAnalysis sets GradsAnalysis field to given value.


### GetFeatureImportance

`func (o *VizType) GetFeatureImportance() FeatureImportance`

GetFeatureImportance returns the FeatureImportance field if non-nil, zero value otherwise.

### GetFeatureImportanceOk

`func (o *VizType) GetFeatureImportanceOk() (*FeatureImportance, bool)`

GetFeatureImportanceOk returns a tuple with the FeatureImportance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureImportance

`func (o *VizType) SetFeatureImportance(v FeatureImportance)`

SetFeatureImportance sets FeatureImportance field to given value.


### GetMetadataMap

`func (o *VizType) GetMetadataMap() map[string]interface{}`

GetMetadataMap returns the MetadataMap field if non-nil, zero value otherwise.

### GetMetadataMapOk

`func (o *VizType) GetMetadataMapOk() (*map[string]interface{}, bool)`

GetMetadataMapOk returns a tuple with the MetadataMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataMap

`func (o *VizType) SetMetadataMap(v map[string]interface{})`

SetMetadataMap sets MetadataMap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


