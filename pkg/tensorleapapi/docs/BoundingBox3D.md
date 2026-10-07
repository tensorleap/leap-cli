# BoundingBox3D

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | **float64** |  | 
**Y** | **float64** |  | 
**Z** | **float64** |  | 
**Width** | **float64** |  | 
**Length** | **float64** |  | 
**Height** | **float64** |  | 
**Yaw** | **float64** |  | 
**Label** | **string** |  | 
**Confidence** | **float64** |  | 
**IsGroundTruth** | **bool** |  | 
**Metadata** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewBoundingBox3D

`func NewBoundingBox3D(x float64, y float64, z float64, width float64, length float64, height float64, yaw float64, label string, confidence float64, isGroundTruth bool, ) *BoundingBox3D`

NewBoundingBox3D instantiates a new BoundingBox3D object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundingBox3DWithDefaults

`func NewBoundingBox3DWithDefaults() *BoundingBox3D`

NewBoundingBox3DWithDefaults instantiates a new BoundingBox3D object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *BoundingBox3D) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *BoundingBox3D) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *BoundingBox3D) SetX(v float64)`

SetX sets X field to given value.


### GetY

`func (o *BoundingBox3D) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *BoundingBox3D) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *BoundingBox3D) SetY(v float64)`

SetY sets Y field to given value.


### GetZ

`func (o *BoundingBox3D) GetZ() float64`

GetZ returns the Z field if non-nil, zero value otherwise.

### GetZOk

`func (o *BoundingBox3D) GetZOk() (*float64, bool)`

GetZOk returns a tuple with the Z field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZ

`func (o *BoundingBox3D) SetZ(v float64)`

SetZ sets Z field to given value.


### GetWidth

`func (o *BoundingBox3D) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *BoundingBox3D) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *BoundingBox3D) SetWidth(v float64)`

SetWidth sets Width field to given value.


### GetLength

`func (o *BoundingBox3D) GetLength() float64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *BoundingBox3D) GetLengthOk() (*float64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *BoundingBox3D) SetLength(v float64)`

SetLength sets Length field to given value.


### GetHeight

`func (o *BoundingBox3D) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BoundingBox3D) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BoundingBox3D) SetHeight(v float64)`

SetHeight sets Height field to given value.


### GetYaw

`func (o *BoundingBox3D) GetYaw() float64`

GetYaw returns the Yaw field if non-nil, zero value otherwise.

### GetYawOk

`func (o *BoundingBox3D) GetYawOk() (*float64, bool)`

GetYawOk returns a tuple with the Yaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaw

`func (o *BoundingBox3D) SetYaw(v float64)`

SetYaw sets Yaw field to given value.


### GetLabel

`func (o *BoundingBox3D) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BoundingBox3D) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BoundingBox3D) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetConfidence

`func (o *BoundingBox3D) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *BoundingBox3D) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *BoundingBox3D) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.


### GetIsGroundTruth

`func (o *BoundingBox3D) GetIsGroundTruth() bool`

GetIsGroundTruth returns the IsGroundTruth field if non-nil, zero value otherwise.

### GetIsGroundTruthOk

`func (o *BoundingBox3D) GetIsGroundTruthOk() (*bool, bool)`

GetIsGroundTruthOk returns a tuple with the IsGroundTruth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroundTruth

`func (o *BoundingBox3D) SetIsGroundTruth(v bool)`

SetIsGroundTruth sets IsGroundTruth field to given value.


### GetMetadata

`func (o *BoundingBox3D) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BoundingBox3D) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BoundingBox3D) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BoundingBox3D) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


