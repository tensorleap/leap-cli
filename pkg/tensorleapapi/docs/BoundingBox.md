# BoundingBox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | **float64** |  | 
**Y** | **float64** |  | 
**Width** | **float64** |  | 
**Height** | **float64** |  | 
**Confidence** | **float64** |  | 
**Label** | **string** |  | 
**Rotation** | Pointer to **float64** |  | [optional] 

## Methods

### NewBoundingBox

`func NewBoundingBox(x float64, y float64, width float64, height float64, confidence float64, label string, ) *BoundingBox`

NewBoundingBox instantiates a new BoundingBox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundingBoxWithDefaults

`func NewBoundingBoxWithDefaults() *BoundingBox`

NewBoundingBoxWithDefaults instantiates a new BoundingBox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *BoundingBox) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *BoundingBox) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *BoundingBox) SetX(v float64)`

SetX sets X field to given value.


### GetY

`func (o *BoundingBox) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *BoundingBox) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *BoundingBox) SetY(v float64)`

SetY sets Y field to given value.


### GetWidth

`func (o *BoundingBox) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *BoundingBox) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *BoundingBox) SetWidth(v float64)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *BoundingBox) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BoundingBox) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BoundingBox) SetHeight(v float64)`

SetHeight sets Height field to given value.


### GetConfidence

`func (o *BoundingBox) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *BoundingBox) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *BoundingBox) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.


### GetLabel

`func (o *BoundingBox) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BoundingBox) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BoundingBox) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRotation

`func (o *BoundingBox) GetRotation() float64`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *BoundingBox) GetRotationOk() (*float64, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *BoundingBox) SetRotation(v float64)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *BoundingBox) HasRotation() bool`

HasRotation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


