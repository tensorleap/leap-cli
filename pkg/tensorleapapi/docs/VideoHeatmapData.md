# VideoHeatmapData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoBlob** | **string** |  | 
**HeatmapBlob** | **string** |  | 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 

## Methods

### NewVideoHeatmapData

`func NewVideoHeatmapData(videoBlob string, heatmapBlob string, type_ DataTypeEnum, ) *VideoHeatmapData`

NewVideoHeatmapData instantiates a new VideoHeatmapData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoHeatmapDataWithDefaults

`func NewVideoHeatmapDataWithDefaults() *VideoHeatmapData`

NewVideoHeatmapDataWithDefaults instantiates a new VideoHeatmapData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoBlob

`func (o *VideoHeatmapData) GetVideoBlob() string`

GetVideoBlob returns the VideoBlob field if non-nil, zero value otherwise.

### GetVideoBlobOk

`func (o *VideoHeatmapData) GetVideoBlobOk() (*string, bool)`

GetVideoBlobOk returns a tuple with the VideoBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoBlob

`func (o *VideoHeatmapData) SetVideoBlob(v string)`

SetVideoBlob sets VideoBlob field to given value.


### GetHeatmapBlob

`func (o *VideoHeatmapData) GetHeatmapBlob() string`

GetHeatmapBlob returns the HeatmapBlob field if non-nil, zero value otherwise.

### GetHeatmapBlobOk

`func (o *VideoHeatmapData) GetHeatmapBlobOk() (*string, bool)`

GetHeatmapBlobOk returns a tuple with the HeatmapBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmapBlob

`func (o *VideoHeatmapData) SetHeatmapBlob(v string)`

SetHeatmapBlob sets HeatmapBlob field to given value.


### GetType

`func (o *VideoHeatmapData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VideoHeatmapData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VideoHeatmapData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


