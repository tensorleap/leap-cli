# BBoxImageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Src** | **string** |  | 
**Blob** | **string** |  | 
**BoundingBox** | [**[]BoundingBox**](BoundingBox.md) |  | 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 

## Methods

### NewBBoxImageData

`func NewBBoxImageData(src string, blob string, boundingBox []BoundingBox, type_ DataTypeEnum, ) *BBoxImageData`

NewBBoxImageData instantiates a new BBoxImageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBBoxImageDataWithDefaults

`func NewBBoxImageDataWithDefaults() *BBoxImageData`

NewBBoxImageDataWithDefaults instantiates a new BBoxImageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrc

`func (o *BBoxImageData) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *BBoxImageData) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *BBoxImageData) SetSrc(v string)`

SetSrc sets Src field to given value.


### GetBlob

`func (o *BBoxImageData) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *BBoxImageData) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *BBoxImageData) SetBlob(v string)`

SetBlob sets Blob field to given value.


### GetBoundingBox

`func (o *BBoxImageData) GetBoundingBox() []BoundingBox`

GetBoundingBox returns the BoundingBox field if non-nil, zero value otherwise.

### GetBoundingBoxOk

`func (o *BBoxImageData) GetBoundingBoxOk() (*[]BoundingBox, bool)`

GetBoundingBoxOk returns a tuple with the BoundingBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingBox

`func (o *BBoxImageData) SetBoundingBox(v []BoundingBox)`

SetBoundingBox sets BoundingBox field to given value.


### GetType

`func (o *BBoxImageData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BBoxImageData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BBoxImageData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


