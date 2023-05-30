# BBoxImageScatterLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **[][]float64** |  | 
**Blob** | **string** |  | 
**Type** | **string** |  | 
**BoundingBoxes** | [**[][]BoundingBox**]([]BoundingBox.md) |  | 
**InputName** | Pointer to **string** |  | [optional] 

## Methods

### NewBBoxImageScatterLabel

`func NewBBoxImageScatterLabel(data [][]float64, blob string, type_ string, boundingBoxes [][]BoundingBox, ) *BBoxImageScatterLabel`

NewBBoxImageScatterLabel instantiates a new BBoxImageScatterLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBBoxImageScatterLabelWithDefaults

`func NewBBoxImageScatterLabelWithDefaults() *BBoxImageScatterLabel`

NewBBoxImageScatterLabelWithDefaults instantiates a new BBoxImageScatterLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BBoxImageScatterLabel) GetData() [][]float64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BBoxImageScatterLabel) GetDataOk() (*[][]float64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BBoxImageScatterLabel) SetData(v [][]float64)`

SetData sets Data field to given value.


### GetBlob

`func (o *BBoxImageScatterLabel) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *BBoxImageScatterLabel) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *BBoxImageScatterLabel) SetBlob(v string)`

SetBlob sets Blob field to given value.


### GetType

`func (o *BBoxImageScatterLabel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BBoxImageScatterLabel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BBoxImageScatterLabel) SetType(v string)`

SetType sets Type field to given value.


### GetBoundingBoxes

`func (o *BBoxImageScatterLabel) GetBoundingBoxes() [][]BoundingBox`

GetBoundingBoxes returns the BoundingBoxes field if non-nil, zero value otherwise.

### GetBoundingBoxesOk

`func (o *BBoxImageScatterLabel) GetBoundingBoxesOk() (*[][]BoundingBox, bool)`

GetBoundingBoxesOk returns a tuple with the BoundingBoxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingBoxes

`func (o *BBoxImageScatterLabel) SetBoundingBoxes(v [][]BoundingBox)`

SetBoundingBoxes sets BoundingBoxes field to given value.


### GetInputName

`func (o *BBoxImageScatterLabel) GetInputName() string`

GetInputName returns the InputName field if non-nil, zero value otherwise.

### GetInputNameOk

`func (o *BBoxImageScatterLabel) GetInputNameOk() (*string, bool)`

GetInputNameOk returns a tuple with the InputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputName

`func (o *BBoxImageScatterLabel) SetInputName(v string)`

SetInputName sets InputName field to given value.

### HasInputName

`func (o *BBoxImageScatterLabel) HasInputName() bool`

HasInputName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


