# ImageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blob** | **string** |  | 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 

## Methods

### NewImageData

`func NewImageData(blob string, type_ DataTypeEnum, ) *ImageData`

NewImageData instantiates a new ImageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageDataWithDefaults

`func NewImageDataWithDefaults() *ImageData`

NewImageDataWithDefaults instantiates a new ImageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlob

`func (o *ImageData) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *ImageData) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *ImageData) SetBlob(v string)`

SetBlob sets Blob field to given value.


### GetType

`func (o *ImageData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


