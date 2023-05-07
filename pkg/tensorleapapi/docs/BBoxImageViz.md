# BBoxImageViz

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | **string** |  | 
**SubTitle** | **string** |  | 
**Guid** | **string** |  | 
**BoundingBoxes** | [**[]BoundingBox**](BoundingBox.md) |  | 
**Labels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBBoxImageViz

`func NewBBoxImageViz(type_ string, title string, subTitle string, guid string, boundingBoxes []BoundingBox, ) *BBoxImageViz`

NewBBoxImageViz instantiates a new BBoxImageViz object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBBoxImageVizWithDefaults

`func NewBBoxImageVizWithDefaults() *BBoxImageViz`

NewBBoxImageVizWithDefaults instantiates a new BBoxImageViz object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BBoxImageViz) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BBoxImageViz) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BBoxImageViz) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *BBoxImageViz) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BBoxImageViz) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BBoxImageViz) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSubTitle

`func (o *BBoxImageViz) GetSubTitle() string`

GetSubTitle returns the SubTitle field if non-nil, zero value otherwise.

### GetSubTitleOk

`func (o *BBoxImageViz) GetSubTitleOk() (*string, bool)`

GetSubTitleOk returns a tuple with the SubTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTitle

`func (o *BBoxImageViz) SetSubTitle(v string)`

SetSubTitle sets SubTitle field to given value.


### GetGuid

`func (o *BBoxImageViz) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *BBoxImageViz) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *BBoxImageViz) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetBoundingBoxes

`func (o *BBoxImageViz) GetBoundingBoxes() []BoundingBox`

GetBoundingBoxes returns the BoundingBoxes field if non-nil, zero value otherwise.

### GetBoundingBoxesOk

`func (o *BBoxImageViz) GetBoundingBoxesOk() (*[]BoundingBox, bool)`

GetBoundingBoxesOk returns a tuple with the BoundingBoxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingBoxes

`func (o *BBoxImageViz) SetBoundingBoxes(v []BoundingBox)`

SetBoundingBoxes sets BoundingBoxes field to given value.


### GetLabels

`func (o *BBoxImageViz) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *BBoxImageViz) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *BBoxImageViz) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *BBoxImageViz) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


