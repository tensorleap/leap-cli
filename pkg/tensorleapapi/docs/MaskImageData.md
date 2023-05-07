# MaskImageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Src** | **string** |  | 
**Blob** | **string** |  | 
**MaskSrc** | **string** |  | 
**MaskBlob** | **string** |  | 
**Labels** | **[]string** |  | 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 

## Methods

### NewMaskImageData

`func NewMaskImageData(src string, blob string, maskSrc string, maskBlob string, labels []string, type_ DataTypeEnum, ) *MaskImageData`

NewMaskImageData instantiates a new MaskImageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaskImageDataWithDefaults

`func NewMaskImageDataWithDefaults() *MaskImageData`

NewMaskImageDataWithDefaults instantiates a new MaskImageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrc

`func (o *MaskImageData) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *MaskImageData) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *MaskImageData) SetSrc(v string)`

SetSrc sets Src field to given value.


### GetBlob

`func (o *MaskImageData) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *MaskImageData) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *MaskImageData) SetBlob(v string)`

SetBlob sets Blob field to given value.


### GetMaskSrc

`func (o *MaskImageData) GetMaskSrc() string`

GetMaskSrc returns the MaskSrc field if non-nil, zero value otherwise.

### GetMaskSrcOk

`func (o *MaskImageData) GetMaskSrcOk() (*string, bool)`

GetMaskSrcOk returns a tuple with the MaskSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskSrc

`func (o *MaskImageData) SetMaskSrc(v string)`

SetMaskSrc sets MaskSrc field to given value.


### GetMaskBlob

`func (o *MaskImageData) GetMaskBlob() string`

GetMaskBlob returns the MaskBlob field if non-nil, zero value otherwise.

### GetMaskBlobOk

`func (o *MaskImageData) GetMaskBlobOk() (*string, bool)`

GetMaskBlobOk returns a tuple with the MaskBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBlob

`func (o *MaskImageData) SetMaskBlob(v string)`

SetMaskBlob sets MaskBlob field to given value.


### GetLabels

`func (o *MaskImageData) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MaskImageData) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MaskImageData) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetType

`func (o *MaskImageData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MaskImageData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MaskImageData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


