# VideoData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoBlob** | **string** |  | 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 

## Methods

### NewVideoData

`func NewVideoData(videoBlob string, type_ DataTypeEnum, ) *VideoData`

NewVideoData instantiates a new VideoData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoDataWithDefaults

`func NewVideoDataWithDefaults() *VideoData`

NewVideoDataWithDefaults instantiates a new VideoData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoBlob

`func (o *VideoData) GetVideoBlob() string`

GetVideoBlob returns the VideoBlob field if non-nil, zero value otherwise.

### GetVideoBlobOk

`func (o *VideoData) GetVideoBlobOk() (*string, bool)`

GetVideoBlobOk returns a tuple with the VideoBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoBlob

`func (o *VideoData) SetVideoBlob(v string)`

SetVideoBlob sets VideoBlob field to given value.


### GetType

`func (o *VideoData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VideoData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VideoData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


