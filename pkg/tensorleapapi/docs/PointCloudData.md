# PointCloudData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PointsBlob** | **string** |  | 
**NumPoints** | **float64** |  | 
**Bounds** | **[]float64** |  | 
**PreviewBlob** | **string** |  | 
**IntensityBlob** | Pointer to **string** |  | [optional] 
**IntensityRange** | Pointer to **[]float64** |  | [optional] 
**BoundingBoxes** | [**[]BoundingBox3D**](BoundingBox3D.md) |  | 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 

## Methods

### NewPointCloudData

`func NewPointCloudData(pointsBlob string, numPoints float64, bounds []float64, previewBlob string, boundingBoxes []BoundingBox3D, type_ DataTypeEnum, ) *PointCloudData`

NewPointCloudData instantiates a new PointCloudData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointCloudDataWithDefaults

`func NewPointCloudDataWithDefaults() *PointCloudData`

NewPointCloudDataWithDefaults instantiates a new PointCloudData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPointsBlob

`func (o *PointCloudData) GetPointsBlob() string`

GetPointsBlob returns the PointsBlob field if non-nil, zero value otherwise.

### GetPointsBlobOk

`func (o *PointCloudData) GetPointsBlobOk() (*string, bool)`

GetPointsBlobOk returns a tuple with the PointsBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsBlob

`func (o *PointCloudData) SetPointsBlob(v string)`

SetPointsBlob sets PointsBlob field to given value.


### GetNumPoints

`func (o *PointCloudData) GetNumPoints() float64`

GetNumPoints returns the NumPoints field if non-nil, zero value otherwise.

### GetNumPointsOk

`func (o *PointCloudData) GetNumPointsOk() (*float64, bool)`

GetNumPointsOk returns a tuple with the NumPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPoints

`func (o *PointCloudData) SetNumPoints(v float64)`

SetNumPoints sets NumPoints field to given value.


### GetBounds

`func (o *PointCloudData) GetBounds() []float64`

GetBounds returns the Bounds field if non-nil, zero value otherwise.

### GetBoundsOk

`func (o *PointCloudData) GetBoundsOk() (*[]float64, bool)`

GetBoundsOk returns a tuple with the Bounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounds

`func (o *PointCloudData) SetBounds(v []float64)`

SetBounds sets Bounds field to given value.


### GetPreviewBlob

`func (o *PointCloudData) GetPreviewBlob() string`

GetPreviewBlob returns the PreviewBlob field if non-nil, zero value otherwise.

### GetPreviewBlobOk

`func (o *PointCloudData) GetPreviewBlobOk() (*string, bool)`

GetPreviewBlobOk returns a tuple with the PreviewBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewBlob

`func (o *PointCloudData) SetPreviewBlob(v string)`

SetPreviewBlob sets PreviewBlob field to given value.


### GetIntensityBlob

`func (o *PointCloudData) GetIntensityBlob() string`

GetIntensityBlob returns the IntensityBlob field if non-nil, zero value otherwise.

### GetIntensityBlobOk

`func (o *PointCloudData) GetIntensityBlobOk() (*string, bool)`

GetIntensityBlobOk returns a tuple with the IntensityBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntensityBlob

`func (o *PointCloudData) SetIntensityBlob(v string)`

SetIntensityBlob sets IntensityBlob field to given value.

### HasIntensityBlob

`func (o *PointCloudData) HasIntensityBlob() bool`

HasIntensityBlob returns a boolean if a field has been set.

### GetIntensityRange

`func (o *PointCloudData) GetIntensityRange() []float64`

GetIntensityRange returns the IntensityRange field if non-nil, zero value otherwise.

### GetIntensityRangeOk

`func (o *PointCloudData) GetIntensityRangeOk() (*[]float64, bool)`

GetIntensityRangeOk returns a tuple with the IntensityRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntensityRange

`func (o *PointCloudData) SetIntensityRange(v []float64)`

SetIntensityRange sets IntensityRange field to given value.

### HasIntensityRange

`func (o *PointCloudData) HasIntensityRange() bool`

HasIntensityRange returns a boolean if a field has been set.

### GetBoundingBoxes

`func (o *PointCloudData) GetBoundingBoxes() []BoundingBox3D`

GetBoundingBoxes returns the BoundingBoxes field if non-nil, zero value otherwise.

### GetBoundingBoxesOk

`func (o *PointCloudData) GetBoundingBoxesOk() (*[]BoundingBox3D, bool)`

GetBoundingBoxesOk returns a tuple with the BoundingBoxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingBoxes

`func (o *PointCloudData) SetBoundingBoxes(v []BoundingBox3D)`

SetBoundingBoxes sets BoundingBoxes field to given value.


### GetType

`func (o *PointCloudData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PointCloudData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PointCloudData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


